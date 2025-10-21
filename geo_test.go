package fake

import (
	"testing"
)

func TestGeo(t *testing.T) {
	for _, lang := range GetLangs() {
		fk := New()
		_ = fk.SetLang(lang)

		f := fk.Latitude()
		if f < -90 || f > 90 {
			t.Errorf("Latitude failed with lang %s", lang)
		}

		i := fk.LatitudeDegrees()
		if i < -180 || i > 180 {
			t.Errorf("LatitudeDegrees failed with lang %s", lang)
		}

		i = fk.LatitudeMinutes()
		if i < 0 || i >= 60 {
			t.Errorf("LatitudeMinutes failed with lang %s", lang)
		}

		i = fk.LatitudeSeconds()
		if i < 0 || i >= 60 {
			t.Errorf("LatitudeSeconds failed with lang %s", lang)
		}

		s := fk.LatitudeDirection()
		if s != "N" && s != "S" {
			t.Errorf("LatitudeDirection failed with lang %s", lang)
		}

		f = fk.Longitude()
		if f < -180 || f > 180 {
			t.Errorf("Longitude failed with lang %s", lang)
		}

		i = fk.LongitudeDegrees()
		if i < -180 || i > 180 {
			t.Errorf("LongitudeDegrees failed with lang %s", lang)
		}

		i = fk.LongitudeMinutes()
		if i < 0 || i >= 60 {
			t.Errorf("LongitudeMinutes failed with lang %s", lang)
		}

		i = fk.LongitudeSeconds()
		if i < 0 || i >= 60 {
			t.Errorf("LongitudeSeconds failed with lang %s", lang)
		}

		s = fk.LongitudeDirection()
		if s != "W" && s != "E" {
			t.Errorf("LongitudeDirection failed with lang %s", lang)
		}
	}
}
