package locale

import (
	"strings"

	"github.com/shopspring/decimal"
)

// getWordForNumber finds the word representation for a number
func GetWordForNumber(number decimal.Decimal, targetLocale NumI18NLocale) string {
	for _, mapping := range targetLocale.NumberWordsMapping {
		if decimal.NewFromInt(mapping.Number).Equal(number) {
			return mapping.Value
		}
	}
	return ""
}

// GetWordForNumberInt64 finds the word representation for an int64 number (backward compatibility)
func GetWordForNumberInt64(number int64, targetLocale NumI18NLocale) string {
	return GetWordForNumber(decimal.NewFromInt(number), targetLocale)
}

// getTensWord finds the tens word representation for a number (20, 30, 40, etc.)
func GetTensWord(tens decimal.Decimal, targetLocale NumI18NLocale) string {
	tensNumber := tens.Mul(decimal.NewFromInt(10))
	return GetWordForNumber(tensNumber, targetLocale)
}

// GetTensWordInt finds the tens word representation for an int number (backward compatibility)
func GetTensWordInt(tens int, targetLocale NumI18NLocale) string {
	return GetTensWord(decimal.NewFromInt(int64(tens)), targetLocale)
}

// convertToWordsGeneric converts number to words using standard algorithm
func ConvertToWordsGeneric(number decimal.Decimal, targetLocale NumI18NLocale) string {
	if number.Equal(decimal.Zero) {
		return GetWordForNumber(decimal.Zero, targetLocale)
	}

	result := ""
	billion := decimal.NewFromInt(1000000000)
	million := decimal.NewFromInt(1000000)
	thousand := decimal.NewFromInt(1000)
	hundred := decimal.NewFromInt(100)
	twenty := decimal.NewFromInt(20)
	ten := decimal.NewFromInt(10)

	// Handle billions
	if number.GreaterThanOrEqual(billion) {
		billions := number.Div(billion).Floor()
		result += ConvertToWordsGeneric(billions, targetLocale) + " " + GetWordForNumber(billion, targetLocale)
		number = number.Mod(billion)
		if number.GreaterThan(decimal.Zero) {
			result += " "
		}
	}

	// Handle millions
	if number.GreaterThanOrEqual(million) {
		millions := number.Div(million).Floor()
		result += ConvertToWordsGeneric(millions, targetLocale) + " " + GetWordForNumber(million, targetLocale)
		number = number.Mod(million)
		if number.GreaterThan(decimal.Zero) {
			result += " "
		}
	}

	// Handle thousands
	if number.GreaterThanOrEqual(thousand) {
		thousands := number.Div(thousand).Floor()
		result += ConvertToWordsGeneric(thousands, targetLocale) + " " + GetWordForNumber(thousand, targetLocale)
		number = number.Mod(thousand)
		if number.GreaterThan(decimal.Zero) {
			result += " "
		}
	}

	// Handle hundreds
	if number.GreaterThanOrEqual(hundred) {
		hundreds := number.Div(hundred).Floor()
		result += ConvertToWordsGeneric(hundreds, targetLocale) + " " + GetWordForNumber(hundred, targetLocale)
		number = number.Mod(hundred)
		if number.GreaterThan(decimal.Zero) {
			result += " "
		}
	}

	// Handle tens and ones
	if number.GreaterThanOrEqual(twenty) {
		tens := number.Div(ten).Floor()
		result += GetTensWord(tens, targetLocale)
		number = number.Mod(ten)
		if number.GreaterThan(decimal.Zero) {
			result += " " + GetWordForNumber(number, targetLocale)
		}
	} else if number.GreaterThan(decimal.Zero) {
		result += GetWordForNumber(number, targetLocale)
	}

	return strings.TrimSpace(result)
}

// ConvertToWordsGenericInt64 converts int64 number to words (backward compatibility)
func ConvertToWordsGenericInt64(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsGeneric(decimal.NewFromInt(number), targetLocale)
}

// convertToWordsWithExactMapping converts number to words using exact mapping when available
func ConvertToWordsWithExactMapping(number decimal.Decimal, targetLocale NumI18NLocale) string {
	// Check for exact word mapping first (special cases like "One Hundred")
	for _, mapping := range targetLocale.ExactWordsMapping {
		if decimal.NewFromInt(mapping.Number).Equal(number) {
			return mapping.Value
		}
	}

	// Check for regular number mapping (thousands, millions, etc.)
	for _, mapping := range targetLocale.NumberWordsMapping {
		if decimal.NewFromInt(mapping.Number).Equal(number) {
			return mapping.Value
		}
	}

	return ConvertToWordsGeneric(number, targetLocale)
}

// ConvertToWordsWithExactMappingInt64 converts int64 number to words using exact mapping (backward compatibility)
func ConvertToWordsWithExactMappingInt64(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMapping(decimal.NewFromInt(number), targetLocale)
}
