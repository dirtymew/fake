package fake

import (
	"strings"
)

// Character generates random character in the given language
func (f *Fake) Character() string {
	return f.lookup(f.lang, "characters", true)
}

// CharactersN generates n random characters in the given language
func (f *Fake) CharactersN(n int) string {
	var chars []string
	for i := 0; i < n; i++ {
		chars = append(chars, f.Character())
	}
	return strings.Join(chars, "")
}

// Characters generates from 1 to 5 characters in the given language
func (f *Fake) Characters() string {
	return f.CharactersN(f.r.Intn(5) + 1)
}

// Word generates random word
func (f *Fake) Word() string {
	return f.lookup(f.lang, "words", true)
}

// WordsN generates n random words
func (f *Fake) WordsN(n int) string {
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = f.Word()
	}
	return strings.Join(words, " ")
}

// Words generates from 1 to 5 random words
func (f *Fake) Words() string {
	return f.WordsN(f.r.Intn(5) + 1)
}

// Title generates from 2 to 5 titleized words
func (f *Fake) Title() string {
	return strings.ToTitle(f.WordsN(2 + f.r.Intn(4)))
}

// Sentence generates random sentence
func (f *Fake) Sentence() string {
	var words []string
	for i := 0; i < 3+f.r.Intn(12); i++ {
		word := f.Word()
		if f.r.Intn(5) == 0 {
			word += ","
		}
		words = append(words, word)
	}

	sentence := strings.Join(words, " ")

	if f.r.Intn(8) == 0 {
		sentence += "!"
	} else {
		sentence += "."
	}

	return sentence
}

// SentencesN generates n random sentences
func (f *Fake) SentencesN(n int) string {
	sentences := make([]string, n)
	for i := 0; i < n; i++ {
		sentences[i] = f.Sentence()
	}
	return strings.Join(sentences, " ")
}

// Sentences generates from 1 to 5 random sentences
func (f *Fake) Sentences() string {
	return f.SentencesN(f.r.Intn(5) + 1)
}

// Paragraph generates paragraph
func (f *Fake) Paragraph() string {
	return f.SentencesN(f.r.Intn(10) + 1)
}

// ParagraphsN generates n paragraphs
func (f *Fake) ParagraphsN(n int) string {
	var paragraphs []string
	for i := 0; i < n; i++ {
		paragraphs = append(paragraphs, f.Paragraph())
	}
	return strings.Join(paragraphs, "\t")
}

// Paragraphs generates from 1 to 5 paragraphs
func (f *Fake) Paragraphs() string {
	return f.ParagraphsN(f.r.Intn(5) + 1)
}
