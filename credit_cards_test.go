package fake

import (
	"testing"
)

func TestCreditCards(t *testing.T) {
	for _, lang := range GetLangs() {
		f := New()
		_ = f.SetLang(lang)

		v := f.CreditCardType()
		if v == "" {
			t.Errorf("CreditCardType failed with lang %s", lang)
		}

		v = f.CreditCardNum("")
		if v == "" {
			t.Errorf("CreditCardNum failed with lang %s", lang)
		}

		v = f.CreditCardNum("visa")
		if v == "" {
			t.Errorf("CreditCardNum failed with lang %s", lang)
		}
	}
}
