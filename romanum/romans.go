package romanum

import (
	"errors"
	"strings"
)

type romanNumeral struct {
	Value  uint16
	Symbol string
}

type romanNumerals []romanNumeral

type windowedRoman string

//ErrOutOfRomanNumeralRange is thronw when try to convert to roman an arabic greater than 3999
var ErrOutOfRomanNumeralRange = errors.New("roman numerals cannot go above 3999")

const maxRomanNumeral = 3999

var allRomanNumerals = romanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func (r romanNumerals) valueOf(symbols ...byte) uint16 {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

func (r romanNumerals) exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return true
		}
	}
	return false
}

func (w windowedRoman) symbols() (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		symbol := w[i]
		notAtEnd := i+1 < len(w)

		if notAtEnd && isSubtractive(symbol) && allRomanNumerals.exists(symbol, w[i+1]) {
			symbols = append(symbols, []byte{symbol, w[i+1]})
			i++
		} else {
			symbols = append(symbols, []byte{symbol})
		}
	}
	return
}

//ConvertToRoman transform an arabic/decimal number to a roman numeral string
func ConvertToRoman(arabic uint16) (string, error) {

	if arabic > maxRomanNumeral {
		return "", ErrOutOfRomanNumeralRange
	}

	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}

	}

	return result.String(), nil
}

//ConvertToArabic transform a roman numeral string to a decimal/arabic number
func ConvertToArabic(roman string) (total uint16) {
	for _, symbols := range windowedRoman(roman).symbols() {
		total += allRomanNumerals.valueOf(symbols...)
	}
	return
}

func isSubtractive(symbol uint8) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}
