package fake

import (
	"testing"
)

func TestDates(t *testing.T) {
	for _, lang := range GetLangs() {
		f := New()
		_ = f.SetLang(lang)
		f.EnFallback(true)

		v := f.WeekDay()
		if v == "" {
			t.Errorf("WeekDay failed with lang %s", lang)
		}

		v = f.WeekDayShort()
		if v == "" {
			t.Errorf("WeekDayShort failed with lang %s", lang)
		}

		n := f.WeekdayNum()
		if n < 0 || n > 7 {
			t.Errorf("WeekdayNum failed with lang %s", lang)
		}

		v = f.Month()
		if v == "" {
			t.Errorf("Month failed with lang %s", lang)
		}

		v = f.MonthShort()
		if v == "" {
			t.Errorf("MonthShort failed with lang %s", lang)
		}

		n = f.MonthNum()
		if n < 0 || n > 31 {
			t.Errorf("MonthNum failed with lang %s", lang)
		}

		n = f.Year(1950, 2020)
		if n < 1950 || n > 2020 {
			t.Errorf("Year failed with lang %s", lang)
		}
	}
}
