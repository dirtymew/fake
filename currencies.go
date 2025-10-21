package fake

// Currency generates currency name
func (f *Fake) Currency() string {
	return f.lookup(f.lang, "currencies", true)
}

// CurrencyCode generates currency code
func (f *Fake) CurrencyCode() string {
	return f.lookup(f.lang, "currency_codes", true)
}
