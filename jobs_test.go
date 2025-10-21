package fake

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJobs(t *testing.T) {
	tests := []struct {
		name        string
		fn          func(*Fake) string
		requiredFor func(lang string) bool
	}{
		{"Company", func(f *Fake) string { return f.Company() }, func(lang string) bool { return true }},
		{"JobTitle", func(f *Fake) string { return f.JobTitle() }, func(lang string) bool { return true }},
		{"Industry", func(f *Fake) string { return f.Industry() }, func(lang string) bool { return true }},
	}

	for _, lang := range GetLangs("") {
		for _, tc := range tests {
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
	}
}
