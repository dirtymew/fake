package fake

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCurrencies(t *testing.T) {
	tests := []struct {
		name        string
		fn          func(*Fake) string
		requiredFor func(lang string) bool
	}{
		{"Currency", func(f *Fake) string { return f.Currency() }, func(lang string) bool { return true }},
		{"CurrencyCode", func(f *Fake) string { return f.CurrencyCode() }, func(lang string) bool { return true }},
	}

	for _, lang := range New().GetLangs() {
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
