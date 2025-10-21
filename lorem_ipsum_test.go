package fake

import (
	"testing"
)

func TestLoremIpsum(t *testing.T) {
	for _, lang := range GetLangs() {
		f := New()
		_ = f.SetLang(lang)
		f.EnFallback(true)

		v := f.Character()
		if v == "" {
			t.Errorf("Character failed with lang %s", lang)
		}

		v = f.CharactersN(2)
		if v == "" {
			t.Errorf("CharactersN failed with lang %s", lang)
		}

		v = f.Characters()
		if v == "" {
			t.Errorf("Characters failed with lang %s", lang)
		}

		v = f.Word()
		if v == "" {
			t.Errorf("Word failed with lang %s", lang)
		}

		v = f.WordsN(2)
		if v == "" {
			t.Errorf("WordsN failed with lang %s", lang)
		}

		v = f.Words()
		if v == "" {
			t.Errorf("Words failed with lang %s", lang)
		}

		v = f.Title()
		if v == "" {
			t.Errorf("Title failed with lang %s", lang)
		}

		v = f.Sentence()
		if v == "" {
			t.Errorf("Sentence failed with lang %s", lang)
		}

		v = f.SentencesN(2)
		if v == "" {
			t.Errorf("SentencesN failed with lang %s", lang)
		}

		v = f.Sentences()
		if v == "" {
			t.Errorf("Sentences failed with lang %s", lang)
		}

		v = f.Paragraph()
		if v == "" {
			t.Errorf("Paragraph failed with lang %s", lang)
		}

		v = f.ParagraphsN(2)
		if v == "" {
			t.Errorf("ParagraphsN failed with lang %s", lang)
		}

		v = f.Paragraphs()
		if v == "" {
			t.Errorf("Paragraphs failed with lang %s", lang)
		}
	}
}
