package romanum

import (
	"fmt"
	"regexp"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

var (
	cases = []struct {
		Arabic uint16
		Roman  string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{9, "IX"},
		{10, "X"},
		{11, "XI"},
		{12, "XII"},
		{14, "XIV"},
		{18, "XVIII"},
		{20, "XX"},
		{39, "XXXIX"},
		{40, "XL"},
		{47, "XLVII"},
		{49, "XLIX"},
		{50, "L"},
		{51, "LI"},
		{54, "LIV"},
		{55, "LV"},
		{100, "C"},
		{90, "XC"},
		{102, "CII"},
		{91, "XCI"},
		{400, "CD"},
		{500, "D"},
		{900, "CM"},
		{1000, "M"},
		{1984, "MCMLXXXIV"},
		{3999, "MMMCMXCIX"},
		{2014, "MMXIV"},
		{1006, "MVI"},
		{798, "DCCXCVIII"},
	}
)

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got, _ := ConvertToRoman(test.Arabic)
			require.Equal(t, test.Roman, got)
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			require.Equal(t, test.Arabic, got)
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		roman, _ := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	assertPropertyBasedTest(t, assertion)
}

func TestPropertiesOfConversionWhenError(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic < 4000 {
			return true
		}
		_, err := ConvertToRoman(arabic)
		return err == ErrOutOfRomanNumeralRange
	}

	assertPropertyBasedTest(t, assertion)
}

func TestPropertiesOfNumeralRepetition(t *testing.T) {
	romanNumeralIlegalRepetitionRegex, _ := regexp.Compile("IIII|XXXX|VVVV|DDDD|CCCC|MMMM|LLLL")
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		roman, _ := ConvertToRoman(arabic)
		return !romanNumeralIlegalRepetitionRegex.MatchString(roman)
	}

	assertPropertyBasedTest(t, assertion)
}

func assertPropertyBasedTest(t *testing.T, assertion interface{}) {
	err := quick.Check(assertion, &quick.Config{
		MaxCount: 10000,
	})
	require.NoError(t, err, err)
}
