package fake

// Day generates day of the month
func (f *Fake) Day() int {
	return f.r.Intn(31) + 1
}

// WeekDay generates name ot the week day
func (f *Fake) WeekDay() string {
	return f.lookup(f.lang, "weekdays", true)
}

// WeekDayShort generates abbreviated name of the week day
func (f *Fake) WeekDayShort() string {
	return f.lookup(f.lang, "weekdays_short", true)
}

// WeekdayNum generates number of the day of the week
func (f *Fake) WeekdayNum() int {
	return f.r.Intn(7) + 1
}

// Month generates month name
func (f *Fake) Month() string {
	return f.lookup(f.lang, "months", true)
}

// MonthShort generates abbreviated month name
func (f *Fake) MonthShort() string {
	return f.lookup(f.lang, "months_short", true)
}

// MonthNum generates month number (from 1 to 12)
func (f *Fake) MonthNum() int {
	return f.r.Intn(12) + 1
}

// Year generates year using the given boundaries
func (f *Fake) Year(from, to int) int {
	n := f.r.Intn(to-from) + 1
	return from + n
}
