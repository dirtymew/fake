package fake

import (
	"strings"
)

// Company generates company name
func (f *Fake) Company() string {
	return f.lookup(f.lang, "companies", true)
}

// JobTitle generates job title
func (f *Fake) JobTitle() string {
	job := f.lookup(f.lang, "jobs", true)
	return strings.Replace(job, "#{N}", f.jobTitleSuffix(), 1)
}

func (f *Fake) jobTitleSuffix() string {
	return f.lookup(f.lang, "jobs_suffixes", false)
}

// Industry generates industry name
func (f *Fake) Industry() string {
	return f.lookup(f.lang, "industries", true)
}
