package fake

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGeo(t *testing.T) {
	floatTests := []struct {
		name string
		fn   func(*Fake) float32
		min  float32
		max  float32
	}{
		{"Latitude", func(f *Fake) float32 { return f.Latitude() }, -90, 90},
		{"Longitude", func(f *Fake) float32 { return f.Longitude() }, -180, 180},
	}

	intTests := []struct {
		name string
		fn   func(*Fake) int
		min  int
		max  int
	}{
		{"LatitudeDegrees", func(f *Fake) int { return f.LatitudeDegrees() }, -90, 90},
		{"LatitudeMinutes", func(f *Fake) int { return f.LatitudeMinutes() }, 0, 59},
		{"LatitudeSeconds", func(f *Fake) int { return f.LatitudeSeconds() }, 0, 59},
		{"LongitudeDegrees", func(f *Fake) int { return f.LongitudeDegrees() }, -180, 180},
		{"LongitudeMinutes", func(f *Fake) int { return f.LongitudeMinutes() }, 0, 59},
		{"LongitudeSeconds", func(f *Fake) int { return f.LongitudeSeconds() }, 0, 59},
	}

	strTests := []struct {
		name string
		fn   func(*Fake) string
		ok   func(string) bool
	}{
		{"LatitudeDirection", func(f *Fake) string { return f.LatitudeDirection() }, func(s string) bool { return s == "N" || s == "S" }},
		{"LongitudeDirection", func(f *Fake) string { return f.LongitudeDirection() }, func(s string) bool { return s == "W" || s == "E" }},
	}

	for _, lang := range New().GetLangs() {
		for _, tc := range floatTests {
			t.Run(fmt.Sprintf("%s/%s", lang, tc.name), func(t *testing.T) {
				t.Parallel()
				f := New()
				require.NoError(t, f.SetLang(lang))

				v := tc.fn(f)
				assert.True(t, v >= tc.min && v <= tc.max, "%s failed with lang %s: got %v, want between %v and %v", tc.name, lang, v, tc.min, tc.max)
			})
		}

		for _, tc := range intTests {
			t.Run(fmt.Sprintf("%s/%s", lang, tc.name), func(t *testing.T) {
				t.Parallel()
				f := New()
				require.NoError(t, f.SetLang(lang))

				n := tc.fn(f)
				assert.True(t, n >= tc.min && n <= tc.max, "%s failed with lang %s: got %d, want between %d and %d", tc.name, lang, n, tc.min, tc.max)
			})
		}

		for _, tc := range strTests {
			t.Run(fmt.Sprintf("%s/%s", lang, tc.name), func(t *testing.T) {
				t.Parallel()
				f := New()
				require.NoError(t, f.SetLang(lang))

				s := tc.fn(f)
				assert.True(t, tc.ok(s), "%s failed with lang %s: got %q", tc.name, lang, s)
			})
		}
	}
}
