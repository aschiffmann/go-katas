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
	return 42
}

func (roman *RomanNumber) String() string {
	return roman.romanLetters
}

func NewRomanNumber(romanLetters string) RomanNumber {
	return RomanNumber{romanLetters: romanLetters}
}
