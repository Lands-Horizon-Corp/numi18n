package pkg

import (
	"regexp"
	"strings"
)

var romanNumerals = []struct {
	Value  int
	Symbol string
}{
	{1000000, "_M"},
	{900000, "_C_M"},
	{500000, "_D"},
	{400000, "_C_D"},
	{100000, "_C"},
	{90000, "_X_C"},
	{50000, "_L"},
	{40000, "_X_L"},
	{10000, "_X"},
	{9000, "M_X"},
	{5000, "_V"},
	{4000, "M_V"},
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

func ToRoman(amount int) string {
	if amount <= 0 {
		return ""
	}

	var sb strings.Builder
	for _, rn := range romanNumerals {
		for amount >= rn.Value {
			sb.WriteString(rn.Symbol)
			amount -= rn.Value
		}
	}
	return sb.String()
}

func FromRoman(roman string) int {
	if !IsValid(roman) {
		return 0
	}
	roman = strings.ToUpper(roman)
	i := 0
	num := 0
	for i < len(roman) {
		if i+1 < len(roman) {
			twoChar := roman[i : i+2]
			for _, rn := range romanNumerals {
				if rn.Symbol == twoChar {
					num += rn.Value
					i += 2
					goto Next
				}
			}
		}
		for _, rn := range romanNumerals {
			if rn.Symbol == string(roman[i]) {
				num += rn.Value
				i++
				break
			}
		}
	Next:
	}
	return num
}

func IsValid(roman string) bool {
	re := `^(_?M{0,3}|M_V|_V|_X_C|_C_D|_D|_C_M|_M)(_?C{0,3}|CM|CD|D?C{0,3})(_?X{0,3}|XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`
	matched, _ := regexp.MatchString(re, roman)
	return matched
}
