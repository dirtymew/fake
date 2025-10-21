package fake

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNames(t *testing.T) {
	tests := []struct {
		name        string
		fn          func(*Fake) string
		requiredFor func(lang string) bool
	}{
		{"MaleFirstName", func(f *Fake) string { return f.MaleFirstName() }, func(lang string) bool { return true }},
		{"FemaleFirstName", func(f *Fake) string { return f.FemaleFirstName() }, func(lang string) bool { return true }},
		{"FirstName", func(f *Fake) string { return f.FirstName() }, func(lang string) bool { return true }},
		{"MaleLastName", func(f *Fake) string { return f.MaleLastName() }, func(lang string) bool { return true }},
		{"FemaleLastName", func(f *Fake) string { return f.FemaleLastName() }, func(lang string) bool { return true }},
		{"LastName", func(f *Fake) string { return f.LastName() }, func(lang string) bool { return true }},
		{"MalePatronymic", func(f *Fake) string { return f.MalePatronymic() }, func(lang string) bool { return true }},
		{"FemalePatronymic", func(f *Fake) string { return f.FemalePatronymic() }, func(lang string) bool { return true }},
		{"Patronymic", func(f *Fake) string { return f.Patronymic() }, func(lang string) bool { return true }},
		{"MaleFullNameWithPrefix", func(f *Fake) string { return f.MaleFullNameWithPrefix() }, func(lang string) bool { return true }},
		{"FemaleFullNameWithPrefix", func(f *Fake) string { return f.FemaleFullNameWithPrefix() }, func(lang string) bool { return true }},
		{"FullNameWithPrefix", func(f *Fake) string { return f.FullNameWithPrefix() }, func(lang string) bool { return true }},
		{"MaleFullNameWithSuffix", func(f *Fake) string { return f.MaleFullNameWithSuffix() }, func(lang string) bool { return true }},
		{"FemaleFullNameWithSuffix", func(f *Fake) string { return f.FemaleFullNameWithSuffix() }, func(lang string) bool { return true }},
		{"FullNameWithSuffix", func(f *Fake) string { return f.FullNameWithSuffix() }, func(lang string) bool { return true }},
		{"MaleFullName", func(f *Fake) string { return f.MaleFullName() }, func(lang string) bool { return true }},
		{"FemaleFullName", func(f *Fake) string { return f.FemaleFullName() }, func(lang string) bool { return true }},
		{"FullName", func(f *Fake) string { return f.FullName() }, func(lang string) bool { return true }},
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
