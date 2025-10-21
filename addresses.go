package fake

import "strconv"

// Continent generates random continent
func (f *Fake) Continent() string {
	return f.lookup(f.lang, "continents", true)
}

// Country generates random country
func (f *Fake) Country() string {
	return f.lookup(f.lang, "countries", true)
}

// City generates random city
func (f *Fake) City() string {
	city := f.lookup(f.lang, "cities", true)
	switch f.r.Intn(5) {
	case 0:
		return join(f.cityPrefix(), city)
	case 1:
		return join(city, f.citySuffix())
	default:
		return city
	}
}

func (f *Fake) cityPrefix() string {
	return f.lookup(f.lang, "city_prefixes", false)
}

func (f *Fake) citySuffix() string {
	return f.lookup(f.lang, "city_suffixes", false)
}

// State generates random state
func (f *Fake) State() string {
	return f.lookup(f.lang, "states", false)
}

// StateAbbrev generates random state abbreviation
func (f *Fake) StateAbbrev() string {
	return f.lookup(f.lang, "state_abbrevs", false)
}

// Street generates random street name
func (f *Fake) Street() string {
	street := f.lookup(f.lang, "streets", true)
	return join(street, f.streetSuffix())
}

// StreetAddress generates random street name along with building number
func (f *Fake) StreetAddress() string {
	return join(f.Street(), strconv.Itoa(f.r.Intn(100)))
}

func (f *Fake) streetSuffix() string {
	return f.lookup(f.lang, "street_suffixes", true)
}

// Zip generates random zip code using one of the formats specifies in zip_format file
func (f *Fake) Zip() string {
	return f.generate(f.lang, "zips", true)
}

// Phone generates random phone number using one of the formats format specified in phone_format file
func (f *Fake) Phone() string {
	return f.generate(f.lang, "phones", true)
}
