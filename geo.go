package fake

// Latitude generates latitude (from -90.0 to 90.0)
func (f *Fake) Latitude() float32 {
	return f.rand.Float32()*180 - 90
}

// LatitudeDegrees generates latitude degrees (from -90 to 90)
func (f *Fake) LatitudeDegrees() int {
	return f.rand.IntN(180) - 90
}

// LatitudeMinutes generates latitude minutes (from 0 to 60)
func (f *Fake) LatitudeMinutes() int {
	return f.rand.IntN(60)
}

// LatitudeSeconds generates latitude seconds (from 0 to 60)
func (f *Fake) LatitudeSeconds() int {
	return f.rand.IntN(60)
}

// LatitudeDirection generates latitude direction (N(orth) o S(outh))
func (f *Fake) LatitudeDirection() string {
	if f.rand.IntN(2) == 0 {
		return "N"
	}
	return "S"
}

// Longitude generates longitude (from -180 to 180)
func (f *Fake) Longitude() float32 {
	return f.rand.Float32()*360 - 180
}

// LongitudeDegrees generates longitude degrees (from -180 to 180)
func (f *Fake) LongitudeDegrees() int {
	return f.rand.IntN(360) - 180
}

// LongitudeMinutes generates (from 0 to 60)
func (f *Fake) LongitudeMinutes() int {
	return f.rand.IntN(60)
}

// LongitudeSeconds generates (from 0 to 60)
func (f *Fake) LongitudeSeconds() int {
	return f.rand.IntN(60)
}

// LongitudeDirection generates (W(est) or E(ast))
func (f *Fake) LongitudeDirection() string {
	if f.rand.IntN(2) == 0 {
		return "W"
	}
	return "E"
}
