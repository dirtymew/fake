package fake

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoremIpsum(t *testing.T) {
	tests := []struct {
		name        string
		fn          func(*Fake) string
		requiredFor func(lang string) bool
	}{
		{"Character", func(f *Fake) string { return f.Character() }, func(lang string) bool { return true }},
		{"CharactersN", func(f *Fake) string { return f.CharactersN(2) }, func(lang string) bool { return true }},
		{"Characters", func(f *Fake) string { return f.Characters() }, func(lang string) bool { return true }},
		{"Word", func(f *Fake) string { return f.Word() }, func(lang string) bool { return true }},
		{"WordsN", func(f *Fake) string { return f.WordsN(2) }, func(lang string) bool { return true }},
		{"Words", func(f *Fake) string { return f.Words() }, func(lang string) bool { return true }},
		{"Title", func(f *Fake) string { return f.Title() }, func(lang string) bool { return true }},
		{"Sentence", func(f *Fake) string { return f.Sentence() }, func(lang string) bool { return true }},
		{"SentencesN", func(f *Fake) string { return f.SentencesN(2) }, func(lang string) bool { return true }},
		{"Sentences", func(f *Fake) string { return f.Sentences() }, func(lang string) bool { return true }},
		{"Paragraph", func(f *Fake) string { return f.Paragraph() }, func(lang string) bool { return true }},
		{"ParagraphsN", func(f *Fake) string { return f.ParagraphsN(2) }, func(lang string) bool { return true }},
		{"Paragraphs", func(f *Fake) string { return f.Paragraphs() }, func(lang string) bool { return true }},
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
