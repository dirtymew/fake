package fake

import (
	"testing"
)

func TestInternet(t *testing.T) {
	for _, lang := range GetLangs() {
		f := New()
		_ = f.SetLang(lang)
		f.EnFallback(true)

		v := f.UserName()
		if v == "" {
			t.Errorf("UserName failed with lang %s", lang)
		}

		v = f.TopLevelDomain()
		if v == "" {
			t.Errorf("TopLevelDomain failed with lang %s", lang)
		}

		v = f.DomainName()
		if v == "" {
			t.Errorf("DomainName failed with lang %s", lang)
		}

		v = f.EmailAddress()
		if v == "" {
			t.Errorf("EmailAddress failed with lang %s", lang)
		}

		v = f.EmailSubject()
		if v == "" {
			t.Errorf("EmailSubject failed with lang %s", lang)
		}

		v = f.EmailBody()
		if v == "" {
			t.Errorf("EmailBody failed with lang %s", lang)
		}

		v = f.DomainZone()
		if v == "" {
			t.Errorf("DomainZone failed with lang %s", lang)
		}

		v = f.IPv4()
		if v == "" {
			t.Errorf("IPv4 failed with lang %s", lang)
		}

		v = f.UserAgent()
		if v == "" {
			t.Errorf("UserAgent failed with lang %s", lang)
		}

		v = f.IPv6()
		if v == "" {
			t.Errorf("IPv6 failed with lang %s", lang)
		}
	}
}
