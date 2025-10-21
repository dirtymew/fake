package fake

import (
	"testing"
)

func TestPersonal(t *testing.T) {
	for _, lang := range GetLangs() {
		f := New()
		_ = f.SetLang(lang)

		v := f.Gender()
		if v == "" {
			t.Errorf("Gender failed with lang %s", lang)
		}

		v = f.GenderAbbrev()
		if v == "" {
			t.Errorf("GenderAbbrev failed with lang %s", lang)
		}

		v = f.Language()
		if v == "" {
			t.Errorf("Language failed with lang %s", lang)
		}
	}
}
