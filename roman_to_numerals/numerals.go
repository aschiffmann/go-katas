package main

import "fmt"

func main() {
	roman := NewRomanNumber("XXX")
	arabic := roman.GetArabicNumber()
	fmt.Printf("the roman number '%s' means '%d' in arabic numbers", roman.String(), arabic)
}

type RomanNumber struct {
	romanLetters string
}

func (roman RomanNumber) GetArabicNumber() int {
	if len(roman.romanLetters) == 1 {
		return convertLetter(roman.romanLetters)
	}
	return 42
}

func convertLetter(l string) int {
	switch l {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}

func (roman *RomanNumber) String() string {
	return roman.romanLetters
}

func NewRomanNumber(romanLetters string) RomanNumber {
	return RomanNumber{romanLetters: romanLetters}
}
