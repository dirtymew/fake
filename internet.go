package fake

import (
	"net"
	"strings"
)

// UserName generates user name in one of the following forms
// first name + last name, letter + last names or concatenation of from 1 to 3 lowercased words
func (f *Fake) UserName() string {
	gender := f.randGender()
	w := f.rand.IntN(3)
	switch w {
	case 0:
		return f.lookup(LangEnglish, gender+"_first_names", false) + f.lookup(LangEnglish, gender+"_last_names", false)
	case 1:
		return f.lookup(LangEnglish, "characters", true) + f.lookup(LangEnglish, gender+"_last_names", false)
	default:
		words := make([]string, w+1)
		for i := 0; i < len(words); i++ {
			words[i] = f.lookup(f.lang, "words", true)
		}
		return strings.Join(words, "_")
	}
}

// TopLevelDomain generates random top level domain
func (f *Fake) TopLevelDomain() string {
	return f.lookup(f.lang, "top_level_domains", true)
}

// DomainName generates random domain name
func (f *Fake) DomainName() string {
	return f.Company() + "." + f.TopLevelDomain()
}

// EmailAddress generates email address
func (f *Fake) EmailAddress() string {
	return f.UserName() + "@" + f.DomainName()
}

// EmailSubject generates random email subject
func (f *Fake) EmailSubject() string {
	return f.Sentence()
}

// EmailBody generates random email body
func (f *Fake) EmailBody() string {
	return f.Paragraphs()
}

// DomainZone generates random domain zone
func (f *Fake) DomainZone() string {
	return f.lookup(f.lang, "domain_zones", true)
}

// IPv4 generates IPv4 address
func (f *Fake) IPv4() string {
	size := 4
	ip := make([]byte, size)
	for i := 0; i < size; i++ {
		ip[i] = byte(f.rand.IntN(256))
	}
	return net.IP(ip).To4().String()
}

// IPv6 generates IPv6 address
func (f *Fake) IPv6() string {
	size := 16
	ip := make([]byte, size)
	for i := 0; i < size; i++ {
		ip[i] = byte(f.rand.IntN(256))
	}
	return net.IP(ip).To16().String()
}
