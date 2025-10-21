package fake

const (
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numeric      = "0123456789"
	specialChars = `!'@#$%^&*()_+-=[]{};:",./?`
	hexDigits    = "0123456789abcdef"
)

func (f *Fake) text(atLeast, atMost int, allowLower, allowUpper, allowNumeric, allowSpecial bool) string {
	if atLeast < 0 {
		atLeast = 0
	}

	if atMost < atLeast {
		atMost = atLeast
	}

	allowedChars := []rune{}
	if allowLower {
		allowedChars = append(allowedChars, []rune(lowerLetters)...)
	}
	if allowUpper {
		allowedChars = append(allowedChars, []rune(upperLetters)...)
	}
	if allowNumeric {
		allowedChars = append(allowedChars, []rune(numeric)...)
	}
	if allowSpecial {
		allowedChars = append(allowedChars, []rune(specialChars)...)
	}

	if len(allowedChars) == 0 {
		return ""
	}

	nTimes := f.rand.Intn(atMost-atLeast+1) + atLeast
	result := make([]rune, nTimes)
	for i := 0; i < nTimes; i++ {
		result[i] = allowedChars[f.rand.Intn(len(allowedChars))]
	}
	return string(result)
}

const (
	DefaultPasswordMinLength = 6
	DefaultPasswordMaxLength = 12
	HexColorLength           = 6
	HexColorShortLength      = 3
)

// Password generates password with the length from atLeast to atMost characters,
// allow* parameters specify whether corresponding symbols can be used
//
// Example:
//
//	f := fake.New()
//	password := f.Password(8, 16, true, true, true)
//	// Generates password like "Abc123!@"
func (f *Fake) Password(atLeast, atMost int, allowUpper, allowNumeric, allowSpecial bool) string {
	return f.text(atLeast, atMost, true, allowUpper, allowNumeric, allowSpecial)
}

// SimplePassword is a wrapper around Password,
// it generates password with the length from 6 to 12 symbols, with upper characters and numeric symbols allowed
func (f *Fake) SimplePassword() string {
	return f.Password(DefaultPasswordMinLength, DefaultPasswordMaxLength, true, true, false)
}

// Color generates color name
func (f *Fake) Color() string {
	return f.lookup(f.lang, "colors", true)
}

// DigitsN returns n digits as a string
func (f *Fake) DigitsN(n int) string {
	if n <= 0 {
		return ""
	}

	digits := make([]rune, n)
	for i := 0; i < n; i++ {
		digits[i] = []rune(numeric)[f.rand.Intn(len([]rune(numeric)))]
	}
	return string(digits)
}

// Digits returns from 1 to 5 digits as a string
func (f *Fake) Digits() string {
	return f.DigitsN(f.rand.Intn(5) + 1)
}

func (f *Fake) hexDigitsStr(n int) string {
	if n <= 0 {
		return ""
	}

	num := make([]rune, n)
	for i := 0; i < n; i++ {
		num[i] = []rune(hexDigits)[f.rand.Intn(len([]rune(hexDigits)))]
	}
	return string(num)
}

// HexColor generates hex color name
func (f *Fake) HexColor() string {
	return f.hexDigitsStr(HexColorLength)
}

// HexColorShort generates short hex color name
func (f *Fake) HexColorShort() string {
	return f.hexDigitsStr(HexColorShortLength)
}
