package fake

import (
	"testing"
)

func TestSetLang(t *testing.T) {
	f := New()
	err := f.SetLang(LangRussian)
	if err != nil {
		t.Error("SetLang should successfully set lang")
	}

	err = f.SetLang("sd")
	if err == nil {
		t.Error("SetLang with nonexistent lang should return error")
	}
}

func TestFakerRuWithoutCallback(t *testing.T) {
	f := New()
	_ = f.SetLang(LangRussian)
	f.EnFallback(false)
	brand := f.Brand()
	if brand != "" {
		t.Error("Fake call with no samples should return blank string")
	}
}

func TestFakerRuWithCallback(t *testing.T) {
	f := New()
	_ = f.SetLang(LangRussian)
	f.EnFallback(true)
	brand := f.Brand()
	if brand == "" {
		t.Error("Fake call for name with no samples with callback should not return blank string")
	}
}

// TestConcurrentSafety runs fake methods in multiple go routines concurrently.
// This test should be run with the race detector enabled.
func TestConcurrentSafety(t *testing.T) {
	workerCount := 10
	doneChan := make(chan struct{})

	for i := 0; i < workerCount; i++ {
		go func() {
			f := New()
			for j := 0; j < 1000; j++ {
				f.FirstName()
				f.LastName()
				f.Gender()
				f.FullName()
				f.Day()
				f.Country()
				f.Company()
				f.Industry()
				f.Street()
			}
			doneChan <- struct{}{}
		}()
	}

	for i := 0; i < workerCount; i++ {
		<-doneChan
	}
}

func BenchmarkNewFastFaker_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		fk := New()
		for pb.Next() {
			fk.UserName()
		}
	})
}
