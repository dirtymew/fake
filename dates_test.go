package fake

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDates(t *testing.T) {
	stringTests := []struct {
		name        string
		fn          func(*Fake) string
		requiredFor func(lang string) bool
	}{
		{"WeekDay", func(f *Fake) string { return f.WeekDay() }, func(lang string) bool { return true }},
		{"WeekDayShort", func(f *Fake) string { return f.WeekDayShort() }, func(lang string) bool { return true }},
		{"Month", func(f *Fake) string { return f.Month() }, func(lang string) bool { return true }},
		{"MonthShort", func(f *Fake) string { return f.MonthShort() }, func(lang string) bool { return true }},
	}

	intTests := []struct {
		name string
		fn   func(*Fake) int
		min  int
		max  int
	}{
		{"Day", func(f *Fake) int { return f.Day() }, 1, 31},
		{"WeekdayNum", func(f *Fake) int { return f.WeekdayNum() }, 1, 7},
		{"MonthNum", func(f *Fake) int { return f.MonthNum() }, 1, 12},
		{"Year", func(f *Fake) int { return f.Year(1950, 2020) }, 1950, 2020},
	}

	for _, lang := range GetLangs("") {
		for _, tc := range stringTests {
			t.Run(fmt.Sprintf("%s/%s", lang, tc.name), func(t *testing.T) {
				t.Parallel()
				f := New()
				require.NoError(t, f.SetLang(lang))
				f.EnFallback(true)

				v := tc.fn(f)
				if tc.requiredFor != nil && tc.requiredFor(lang) {
					assert.NotEmpty(t, v, "%s failed with lang %s", tc.name, lang)
				}
			})
		}

		for _, tc := range intTests {
			t.Run(fmt.Sprintf("%s/%s", lang, tc.name), func(t *testing.T) {
				t.Parallel()
				f := New()
				require.NoError(t, f.SetLang(lang))
				f.EnFallback(true)

				n := tc.fn(f)
				assert.True(t, n >= tc.min && n <= tc.max, "%s failed with lang %s: got %d, want between %d and %d", tc.name, lang, n, tc.min, tc.max)
			})
		}
	}
}
