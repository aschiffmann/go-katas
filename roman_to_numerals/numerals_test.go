package main

import "testing"

func TestConversionsRomanToArabic(t *testing.T) {
	testData := []struct {
		roman          string
		expectedArabic int
	}{
		{"XLII", 42},
		{"I", 1},
		{"V", 5},
		{"X", 10},
		{"L", 50},
		{"C", 100},
		{"D", 500},
		{"M", 1000},
	}

	for testNr, test := range testData {
		result := NewRomanNumber(test.roman).GetArabicNumber()
		if result != test.expectedArabic {
			t.Errorf("Test %d: Expected '%s' to conver to '%d' but got '%d'", testNr, test.roman, test.expectedArabic, result)
		}
	}
}
