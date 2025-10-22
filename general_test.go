package fake

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGeneral(t *testing.T) {
	tests := []struct {
		name        string
		fn          func(*Fake) string
		requiredFor func(lang string) bool
	}{
		{"Password", func(f *Fake) string { return f.Password(4, 10, true, true, true) }, func(lang string) bool { return true }},
		{"SimplePassword", func(f *Fake) string { return f.SimplePassword() }, func(lang string) bool { return true }},
		{"Color", func(f *Fake) string { return f.Color() }, func(lang string) bool { return true }},
		{"HexColor", func(f *Fake) string { return f.HexColor() }, func(lang string) bool { return true }},
		{"HexColorShort", func(f *Fake) string { return f.HexColorShort() }, func(lang string) bool { return true }},
		{"DigitsN", func(f *Fake) string { return f.DigitsN(2) }, func(lang string) bool { return true }},
		{"Digits", func(f *Fake) string { return f.Digits() }, func(lang string) bool { return true }},
	}

	for _, lang := range New().GetLangs() {
		for _, tc := range tests {
			t.Run(fmt.Sprintf("%s/%s", lang, tc.name), func(t *testing.T) {
				t.Parallel()
				f := New()
				require.NoError(t, f.SetLang(lang))

				v := tc.fn(f)
				if tc.requiredFor != nil && tc.requiredFor(lang) {
					assert.NotEmpty(t, v, "%s failed with lang %s", tc.name, lang)
				}
			})
		}
	}
}
