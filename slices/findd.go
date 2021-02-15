package slices

import (
	"regexp"
	"strconv"
	"strings"
)

var digitRegexp = regexp.MustCompile("[0-9]+")

//FindFirstConsecutiveDigits returns an slice containg the first sequence of integers in the input string slice
func FindFirstConsecutiveDigits(sequence []string) []int {
	result := []int{}
	jointString := strings.Join(sequence, "")
	onlyDigitsString := strings.Split(digitRegexp.FindString(jointString), "")
	for _, i := range onlyDigitsString {
		j, _ := strconv.Atoi(i)
		result = append(result, j)
	}
	return result
}
