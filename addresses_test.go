package fake

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddresses(t *testing.T) {
	tests := []struct {
		name        string
		fn          func(*Fake) string
		requiredFor func(lang string) bool
	}{
		{"Continent", func(f *Fake) string { return f.Continent() }, func(lang string) bool { return true }},
		{"Country", func(f *Fake) string { return f.Country() }, func(lang string) bool { return true }},
		{"City", func(f *Fake) string { return f.City() }, func(lang string) bool { return true }},
		{"State", func(f *Fake) string { return f.State() }, func(lang string) bool { return true }},
		{"StateAbbrev", func(f *Fake) string { return f.StateAbbrev() }, func(lang string) bool { return lang == LangEnglish }},
		{"Street", func(f *Fake) string { return f.Street() }, func(lang string) bool { return true }},
		{"StreetAddress", func(f *Fake) string { return f.StreetAddress() }, func(lang string) bool { return true }},
		{"Zip", func(f *Fake) string { return f.Zip() }, func(lang string) bool { return true }},
		{"Phone", func(f *Fake) string { return f.Phone() }, func(lang string) bool { return true }},
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

func TestPhone_RU_Format(t *testing.T) {
	t.Parallel()
	f := New()
	require.NoError(t, f.SetLang("ru"))

	allowed := regexp.MustCompile(`^[\d\+\-\s\(\)]+$`)
	nonDigits := regexp.MustCompile(`\D`)

	// generate several samples to reduce flakiness
	for i := 0; i < 50; i++ {
		p := f.Phonef("+7-###-###-##-##")
		assert.NotEmpty(t, p, "empty phone")
		assert.True(t, allowed.MatchString(p), "unexpected characters in phone %q", p)
		assert.True(t, strings.HasPrefix(p, "+7"))

		digits := nonDigits.ReplaceAllString(p, "")
		assert.Len(t, digits, 11)
	}
}
