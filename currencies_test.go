package fake

import (
	"testing"
)

func TestCurrencies(t *testing.T) {
	for _, lang := range GetLangs() {
		f := New()
		_ = f.SetLang(lang)
		f.EnFallback(true)

		v := f.Currency()
		if v == "" {
			t.Errorf("Currency failed with lang %s", lang)
		}

		v = f.CurrencyCode()
		if v == "" {
			t.Errorf("CurrencyCode failed with lang %s", lang)
		}
	}
}
