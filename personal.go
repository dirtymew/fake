package fake

import (
	"strings"
)

// Gender generates random gender
func (f *Fake) Gender() string {
	return f.lookup(f.lang, "genders", true)
}

// GenderAbbrev returns first downcased letter of the random gender
func (f *Fake) GenderAbbrev() string {
	g := f.Gender()
	if g != "" {
		return strings.ToLower(string(g[0]))
	}
	return ""
}

// Language generates random human language
func (f *Fake) Language() string {
	return f.lookup(f.lang, "languages", true)
}
