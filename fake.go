package fake

import (
	"fmt"
	"io"
	"io/fs"
	"math/rand"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	LangEnglish = "en"
	LangRussian = "ru"
)

type Fake struct {
	rand             *rand.Rand
	samplesCache     samplesTree
	lang             string
	enFallback       bool
	externalDataPath string
	fs               fs.FS
}

func New() *Fake {
	return &Fake{
		rand:         rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
		samplesCache: make(samplesTree),
	}
}

type samplesTree map[string]map[string][]string

var (
	// ErrNoLanguageFn is the error that indicates that given language is not available
	ErrNoLanguageFn = func(lang string) error { return fmt.Errorf("language passed (%s) is not available", lang) }
	// ErrNoSamplesFn is the error that indicates that there are no samples for the given language
	ErrNoSamplesFn = func(lang string) error { return fmt.Errorf("no samples found for language: %s", lang) }
)

// Seed uses the provided seed value to initialize the internal PRNG to a
// deterministic state.
func (f *Fake) Seed(seed int64) {
	if seed == 0 { // 0 unset to random seed
		f.rand = rand.New(rand.NewSource(time.Now().UnixNano()))
		return
	}
	f.rand.Seed(seed)
}

// GetLangs returns a slice of available languages from the embedded data FS.
func (f *Fake) GetLangs() []string {
	var langs []string
	fsys := f.getFS()
	entries, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return langs
	}
	for _, e := range entries {
		if e.IsDir() {
			langs = append(langs, e.Name())
		}
	}
	return langs
}

// SetLang sets the language in which the data should be generated
// returns error if passed language is not available
func (f *Fake) SetLang(newLang string) error {
	found := false
	for _, l := range f.GetLangs() {
		if newLang == l {
			found = true
			break
		}
	}
	if !found {
		return ErrNoLanguageFn(newLang)
	}
	f.lang = newLang
	return nil
}

// EnFallback sets the flag that allows fake to fallback to englsh samples if the ones for the used languaged are not available
func (f *Fake) EnFallback(flag bool) {
	f.enFallback = flag
}

func (st samplesTree) hasKeyPath(lang, cat string) bool {
	if _, ok := st[lang]; ok {
		if _, ok = st[lang][cat]; ok {
			return true
		}
	}
	return false
}

func join(parts ...string) string {
	var filtered []string
	for _, part := range parts {
		if part != "" {
			filtered = append(filtered, part)
		}
	}
	return strings.Join(filtered, " ")
}

func (f *Fake) generate(lang, cat string, fallback bool) string {
	format := f.lookup(lang, cat+"_format", fallback)
	var result string
	for _, ru := range format {
		if ru != '#' {
			result += string(ru)
		} else {
			result += strconv.Itoa(f.rand.Intn(10))
		}
	}
	return result
}

func (f *Fake) lookup(lang string, cat string, fallback bool) string {
	var samples []string

	if f.samplesCache.hasKeyPath(lang, cat) {
		samples = f.samplesCache[lang][cat]
	} else {
		var err error
		samples, err = f.populateSamples(lang, cat)
		if err != nil {
			if f.lang != LangEnglish && fallback && f.enFallback && err.Error() == ErrNoSamplesFn(f.lang).Error() {
				return f.lookup(LangEnglish, cat, false)
			}
			return ""
		}
	}

	return samples[f.rand.Intn(len(samples))]
}

func (f *Fake) populateSamples(lang, cat string) ([]string, error) {
	data, err := f.readFile(lang, cat)
	if err != nil {
		return nil, err
	}

	if _, ok := f.samplesCache[lang]; !ok {
		f.samplesCache[lang] = make(map[string][]string)
	}

	samples := strings.Split(strings.TrimSpace(string(data)), "\n")

	f.samplesCache[lang][cat] = samples
	return samples, nil
}

func (f *Fake) UseExternalFS(fs fs.FS) {
	f.fs = fs
}

func (f *Fake) readFile(lang, cat string) ([]byte, error) {
	fullpath := path.Join(lang, cat) // no leading slash
	file, err := f.getFS().Open(fullpath)
	if err != nil {
		return nil, ErrNoSamplesFn(lang)
	}
	defer func() {
		_ = file.Close()
	}()

	return io.ReadAll(file)
}
