package fake

// Brand generates brand name
func (f *Fake) Brand() string {
	return f.Company()
}

// ProductName generates product name
func (f *Fake) ProductName() string {
	productName := f.lookup(f.lang, "adjectives", true) + " " + f.lookup(f.lang, "nouns", true)
	if f.r.Intn(2) == 1 {
		productName = f.lookup(f.lang, "adjectives", true) + " " + productName
	}
	return productName
}

// Product generates product title as brand + product name
func (f *Fake) Product() string {
	return f.Brand() + " " + f.ProductName()
}

// Model generates model name that consists of letters and digits, optionally with a hyphen between them
func (f *Fake) Model() string {
	seps := []string{"", " ", "-"}
	return f.CharactersN(f.r.Intn(3)+1) + seps[f.r.Intn(len(seps))] + f.Digits()
}
