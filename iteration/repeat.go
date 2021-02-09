package iteration

import "strings"

//Repeat prints a character multiple times
func Repeat(character string, times int) string {
	repeated := character
	if times > 0 {
		for i := 1; i < times; i++ {
			repeated += character
		}
	}
	return repeated
}

//EnhancedRepeat prints a character multiple times using strings standard lib
func EnhancedRepeat(character string, times int) string {
	if times > 0 {
		return strings.Repeat(character, times)
	}
	return character
}
