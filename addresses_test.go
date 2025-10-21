package fake

import (
	"testing"
)

func TestAddresses(t *testing.T) {
	for _, lang := range GetLangs() {
		f := New()
		_ = f.SetLang(lang)

		v := f.Continent()
		if v == "" {
			t.Errorf("Continent failed with lang %s", lang)
		}

		v = f.Country()
		if v == "" {
			t.Errorf("Country failed with lang %s", lang)
		}

		v = f.City()
		if v == "" {
			t.Errorf("City failed with lang %s", lang)
		}

		v = f.State()
		if v == "" {
			t.Errorf("State failed with lang %s", lang)
		}

		v = f.StateAbbrev()
		if v == "" && lang == LangEnglish {
			t.Errorf("StateAbbrev failed with lang %s", lang)
		}

		v = f.Street()
		if v == "" {
			t.Errorf("Street failed with lang %s", lang)
		}

		v = f.StreetAddress()
		if v == "" {
			t.Errorf("StreetAddress failed with lang %s", lang)
		}

		v = f.Zip()
		if v == "" {
			t.Errorf("Zip failed with lang %s", lang)
		}

		v = f.Phone()
		if v == "" {
			t.Errorf("Phone failed with lang %s", lang)
		}
	}
}
