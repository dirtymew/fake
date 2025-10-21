package fake

import (
	"testing"
)

func TestNames(t *testing.T) {
	for _, lang := range GetLangs() {
		f := New()
		_ = f.SetLang(lang)

		v := f.MaleFirstName()
		if v == "" {
			t.Errorf("MaleFirstName failed with lang %s", lang)
		}

		v = f.FemaleFirstName()
		if v == "" {
			t.Errorf("FemaleFirstName failed with lang %s", lang)
		}

		v = f.FirstName()
		if v == "" {
			t.Errorf("FirstName failed with lang %s", lang)
		}

		v = f.MaleLastName()
		if v == "" {
			t.Errorf("MaleLastName failed with lang %s", lang)
		}

		v = f.FemaleLastName()
		if v == "" {
			t.Errorf("FemaleLastName failed with lang %s", lang)
		}

		v = f.LastName()
		if v == "" {
			t.Errorf("LastName failed with lang %s", lang)
		}

		v = f.MalePatronymic()
		if v == "" {
			t.Errorf("MalePatronymic failed with lang %s", lang)
		}

		v = f.FemalePatronymic()
		if v == "" {
			t.Errorf("FemalePatronymic failed with lang %s", lang)
		}

		v = f.Patronymic()
		if v == "" {
			t.Errorf("Patronymic failed with lang %s", lang)
		}

		v = f.MaleFullNameWithPrefix()
		if v == "" {
			t.Errorf("MaleFullNameWithPrefix failed with lang %s", lang)
		}

		v = f.FemaleFullNameWithPrefix()
		if v == "" {
			t.Errorf("FemaleFullNameWithPrefix failed with lang %s", lang)
		}

		v = f.FullNameWithPrefix()
		if v == "" {
			t.Errorf("FullNameWithPrefix failed with lang %s", lang)
		}

		v = f.MaleFullNameWithSuffix()
		if v == "" {
			t.Errorf("MaleFullNameWithSuffix failed with lang %s", lang)
		}

		v = f.FemaleFullNameWithSuffix()
		if v == "" {
			t.Errorf("FemaleFullNameWithSuffix failed with lang %s", lang)
		}

		v = f.FullNameWithSuffix()
		if v == "" {
			t.Errorf("FullNameWithSuffix failed with lang %s", lang)
		}

		v = f.MaleFullName()
		if v == "" {
			t.Errorf("MaleFullName failed with lang %s", lang)
		}

		v = f.FemaleFullName()
		if v == "" {
			t.Errorf("FemaleFullName failed with lang %s", lang)
		}

		v = f.FullName()
		if v == "" {
			t.Errorf("FullName failed with lang %s", lang)
		}
	}
}
