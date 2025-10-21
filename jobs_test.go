package fake

import (
	"testing"
)

func TestJobs(t *testing.T) {
	for _, lang := range GetLangs() {
		f := New()
		_ = f.SetLang(lang)
		f.EnFallback(true)

		v := f.Company()
		if v == "" {
			t.Errorf("Company failed with lang %s", lang)
		}

		v = f.JobTitle()
		if v == "" {
			t.Errorf("JobTitle failed with lang %s", lang)
		}

		v = f.Industry()
		if v == "" {
			t.Errorf("Industry failed with lang %s", lang)
		}
	}
}
