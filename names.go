package fake

func (f *Fake) randGender() string {
	g := "male"
	if f.rand.Intn(2) == 0 {
		g = "female"
	}
	return g
}

func (f *Fake) firstName(gender string) string {
	return f.lookup(f.lang, gender+"_first_names", true)
}

// MaleFirstName generates male first name
func (f *Fake) MaleFirstName() string {
	return f.firstName("male")
}

// FemaleFirstName generates female first name
func (f *Fake) FemaleFirstName() string {
	return f.firstName("female")
}

// FirstName generates first name
func (f *Fake) FirstName() string {
	return f.firstName(f.randGender())
}

func (f *Fake) lastName(gender string) string {
	return f.lookup(f.lang, gender+"_last_names", true)
}

// MaleLastName generates male last name
func (f *Fake) MaleLastName() string {
	return f.lastName("male")
}

// FemaleLastName generates female last name
func (f *Fake) FemaleLastName() string {
	return f.lastName("female")
}

// LastName generates last name
func (f *Fake) LastName() string {
	return f.lastName(f.randGender())
}

func (f *Fake) patronymic(gender string) string {
	return f.lookup(f.lang, gender+"_patronymics", false)
}

// MalePatronymic generates male patronymic
func (f *Fake) MalePatronymic() string {
	return f.patronymic("male")
}

// FemalePatronymic generates female patronymic
func (f *Fake) FemalePatronymic() string {
	return f.patronymic("female")
}

// Patronymic generates patronymic
func (f *Fake) Patronymic() string {
	return f.patronymic(f.randGender())
}

func (f *Fake) prefix(gender string) string {
	return f.lookup(f.lang, gender+"_name_prefixes", false)
}

func (f *Fake) suffix(gender string) string {
	return f.lookup(f.lang, gender+"_name_suffixes", false)
}

func (f *Fake) fullNameWithPrefix(gender string) string {
	return join(f.prefix(gender), f.firstName(gender), f.lastName(gender))
}

// MaleFullNameWithPrefix generates prefixed male full name
// if prefixes for the given language are available
func (f *Fake) MaleFullNameWithPrefix() string {
	return f.fullNameWithPrefix("male")
}

// FemaleFullNameWithPrefix generates prefixed female full name
// if prefixes for the given language are available
func (f *Fake) FemaleFullNameWithPrefix() string {
	return f.fullNameWithPrefix("female")
}

// FullNameWithPrefix generates prefixed full name
// if prefixes for the given language are available
func (f *Fake) FullNameWithPrefix() string {
	return f.fullNameWithPrefix(f.randGender())
}

func (f *Fake) fullNameWithSuffix(gender string) string {
	return join(f.firstName(gender), f.lastName(gender), f.suffix(gender))
}

// MaleFullNameWithSuffix generates suffixed male full name
// if suffixes for the given language are available
func (f *Fake) MaleFullNameWithSuffix() string {
	return f.fullNameWithPrefix("male")
}

// FemaleFullNameWithSuffix generates suffixed female full name
// if suffixes for the given language are available
func (f *Fake) FemaleFullNameWithSuffix() string {
	return f.fullNameWithPrefix("female")
}

// FullNameWithSuffix generates suffixed full name
// if suffixes for the given language are available
func (f *Fake) FullNameWithSuffix() string {
	return f.fullNameWithPrefix(f.randGender())
}

func (f *Fake) fullName(gender string) string {
	switch f.rand.Intn(10) {
	case 0:
		return f.fullNameWithPrefix(gender)
	case 1:
		return f.fullNameWithSuffix(gender)
	default:
		return join(f.firstName(gender), f.lastName(gender))
	}
}

// MaleFullName generates male full name
// it can occasionally include prefix or suffix
func (f *Fake) MaleFullName() string {
	return f.fullName("male")
}

// FemaleFullName generates female full name
// it can occasionally include prefix or suffix
func (f *Fake) FemaleFullName() string {
	return f.fullName("female")
}

// FullName generates full name
// it can occasionally include prefix or suffix
func (f *Fake) FullName() string {
	return f.fullName(f.randGender())
}
