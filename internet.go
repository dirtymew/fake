package fake

import (
	"net"
	"strings"

	"github.com/corpix/uarand"
)

// UserName generates user name in one of the following forms
// first name + last name, letter + last names or concatenation of from 1 to 3 lowercased words
func (f *Fake) UserName() string {
	gender := f.randGender()
	switch f.r.Intn(3) {
	case 0:
		return f.lookup(LangEnglish, gender+"_first_names", false) + f.lookup(f.lang, gender+"_last_names", false)
	case 1:
		return f.Character() + f.lookup(f.lang, gender+"_last_names", false)
	default:
		return strings.Replace(f.WordsN(f.r.Intn(3)+1), " ", "_", -1)
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
		ip[i] = byte(f.r.Intn(256))
	}
	return net.IP(ip).To4().String()
}

// IPv6 generates IPv6 address
func (f *Fake) IPv6() string {
	size := 16
	ip := make([]byte, size)
	for i := 0; i < size; i++ {
		ip[i] = byte(f.r.Intn(256))
	}
	return net.IP(ip).To16().String()
}

// UserAgent generates a random user agent.
func (f *Fake) UserAgent() string {
	return uarand.GetRandom()
}
