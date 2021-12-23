package main

import "testing"

type testData struct {
	roman          string
	expectedArabic int
}

func TestConversionsRomanToArabic_sameLetter(t *testing.T) {
	tests := []testData{
		{"I", 1},
		{"V", 5},
		{"X", 10},
		{"L", 50},
		{"C", 100},
		{"D", 500},
		{"M", 1000},
		{"II", 2},
		{"III", 3},
		{"XX", 20},
		{"XXX", 30},
		{"CC", 200},
		{"CCC", 300},
		{"MM", 2000},
		{"MMM", 3000},
	}

	execTests(t, tests)
}

func TestConversionsRomanToArabic_differentLetters(t *testing.T) {
	tests := []testData{
		{"VI", 6},
		{"VII", 7},
		{"XI", 11},
		{"CXI", 111},
		{"MCXII", 1112},
		{"DLV", 555},
		{"MDCLXI", 1661},
	}

	execTests(t, tests)
}

func execTests(t *testing.T, tests []testData) {
	for testNr, test := range tests {
		result := NewRomanNumber(test.roman).GetArabicNumber()
		if result != test.expectedArabic {
			t.Errorf("Test %d: Expected '%s' to conver to '%d' but got '%d'", testNr, test.roman, test.expectedArabic, result)
		}
	}
}
