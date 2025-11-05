package fake

import (
	"fmt"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInternet(t *testing.T) {
	tests := []struct {
		name        string
		fn          func(*Fake) string
		requiredFor func(lang string) bool
	}{
		{"UserName", func(f *Fake) string { return f.UserName() }, func(lang string) bool { return true }},
		{"TopLevelDomain", func(f *Fake) string { return f.TopLevelDomain() }, func(lang string) bool { return true }},
		{"DomainName", func(f *Fake) string { return f.DomainName() }, func(lang string) bool { return true }},
		{"EmailAddress", func(f *Fake) string { return f.EmailAddress() }, func(lang string) bool { return true }},
		{"EmailSubject", func(f *Fake) string { return f.EmailSubject() }, func(lang string) bool { return true }},
		{"EmailBody", func(f *Fake) string { return f.EmailBody() }, func(lang string) bool { return true }},
		{"DomainZone", func(f *Fake) string { return f.DomainZone() }, func(lang string) bool { return true }},
		{"IPv4", func(f *Fake) string { return f.IPv4() }, func(lang string) bool { return true }},
		{"IPv6", func(f *Fake) string { return f.IPv6() }, func(lang string) bool { return true }},
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

func containsCyrillic(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Cyrillic, r) {
			return true
		}
	}
	return false
}

func TestUserName_Cyrillic_Russian(t *testing.T) {
	t.Parallel()
	f := New()
	require.NoError(t, f.SetLang("ru"))
	// disable English fallback to ensure we exercise Russian data
	f.EnFallback(false)

	// generate several samples to reduce flakiness
	for i := 0; i < 20; i++ {
		u := f.UserName()
		assert.False(t, containsCyrillic(u), "expected Cyrillic characters in username %q", u)
	}
}
