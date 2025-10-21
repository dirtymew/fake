package fake

import (
	"testing"
)

func TestProducts(t *testing.T) {
	for _, lang := range GetLangs() {
		f := New()
		_ = f.SetLang(lang)
		f.EnFallback(true)

		v := f.Brand()
		if v == "" {
			t.Errorf("Brand failed with lang %s", lang)
		}

		v = f.ProductName()
		if v == "" {
			t.Errorf("ProductName failed with lang %s", lang)
		}

		v = f.Product()
		if v == "" {
			t.Errorf("Product failed with lang %s", lang)
		}

		v = f.Model()
		if v == "" {
			t.Errorf("Model failed with lang %s", lang)
		}
	}
}
