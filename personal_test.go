package fake

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPersonal(t *testing.T) {
	tests := []struct {
		name        string
		fn          func(*Fake) string
		requiredFor func(lang string) bool
	}{
		{"Gender", func(f *Fake) string { return f.Gender() }, func(lang string) bool { return true }},
		{"GenderAbbrev", func(f *Fake) string { return f.GenderAbbrev() }, func(lang string) bool { return true }},
		{"Language", func(f *Fake) string { return f.Language() }, func(lang string) bool { return true }},
	}

	for _, lang := range GetLangs() {
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
