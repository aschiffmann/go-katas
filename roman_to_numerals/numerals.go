package main

import "fmt"

func main() {
	roman := NewRomanNumber("XXX")
	arabic := roman.GetArabicNumber()
	fmt.Printf("the roman number '%s' means '%d' in arabic numbers", roman.String(), arabic)
}

type RomanNumber struct {
	romanLetters []rune
}

func (roman RomanNumber) GetArabicNumber() int {
	result := 0
	for l := range roman.romanLetters {
		result += convertLetter(roman.romanLetters[l])
	}
	return result
}

func convertLetter(l rune) int {
	switch l {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}

func (roman *RomanNumber) String() string {
	return string(roman.romanLetters)
}

func NewRomanNumber(romanLetters string) RomanNumber {
	return RomanNumber{romanLetters: []rune(romanLetters)}
}
