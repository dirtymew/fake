package fake

import (
	"testing"
)

func TestGeneral(t *testing.T) {
	for _, lang := range GetLangs() {
		f := New()
		_ = f.SetLang(lang)

		v := f.Password(4, 10, true, true, true)
		if v == "" {
			t.Errorf("Password failed with lang %s", lang)
		}

		v = f.SimplePassword()
		if v == "" {
			t.Errorf("SimplePassword failed with lang %s", lang)
		}

		v = f.Color()
		if v == "" {
			t.Errorf("Color failed with lang %s", lang)
		}

		v = f.HexColor()
		if v == "" {
			t.Errorf("HexColor failed with lang %s", lang)
		}

		v = f.HexColorShort()
		if v == "" {
			t.Errorf("HexColorShort failed with lang %s", lang)
		}

		v = f.DigitsN(2)
		if v == "" {
			t.Errorf("DigitsN failed with lang %s", lang)
		}

		v = f.Digits()
		if v == "" {
			t.Errorf("Digits failed with lang %s", lang)
		}
	}
}
