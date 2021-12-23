package main

import "testing"

func TestConversionsRomanToArabic(t *testing.T) {
	testData := []struct {
		roman          string
		expectedArabic int
	}{
		{"XLII", 42},
	}

	for testNr, test := range testData {
		result := NewRomanNumber(test.roman).GetArabicNumber()
		if result != test.expectedArabic {
			t.Errorf("Test %d: Expected '%s' to conver to '%d' but got '%d'", testNr, test.roman, test.expectedArabic, result)
		}
	}
}
