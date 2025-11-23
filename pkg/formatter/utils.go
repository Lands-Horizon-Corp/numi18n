package formatter

import (
	"strings"

	"github.com/Lands-Horizon-Corp/numi18n/pkg/locale"
)

// getWordForNumber finds the word representation for a number
func GetWordForNumber(number int64, targetLocale locale.NumI18NLocale) string {
	for _, mapping := range targetLocale.NumberWordsMapping {
		if mapping.Number == number {
			return mapping.Value
		}
	}
	return ""
}

// getTensWord finds the tens word representation for a number (20, 30, 40, etc.)
func GetTensWord(tens int, targetLocale locale.NumI18NLocale) string {
	tensNumber := int64(tens * 10)
	return GetWordForNumber(tensNumber, targetLocale)
}

// convertToWordsGeneric converts number to words using standard algorithm
func ConvertToWordsGeneric(number int64, targetLocale locale.NumI18NLocale) string {
	if number == 0 {
		return GetWordForNumber(0, targetLocale)
	}

	result := ""

	// Handle billions
	if number >= 1000000000 {
		billions := number / 1000000000
		result += ConvertToWordsGeneric(billions, targetLocale) + " " + GetWordForNumber(1000000000, targetLocale)
		number %= 1000000000
		if number > 0 {
			result += " "
		}
	}

	// Handle millions
	if number >= 1000000 {
		millions := number / 1000000
		result += ConvertToWordsGeneric(millions, targetLocale) + " " + GetWordForNumber(1000000, targetLocale)
		number %= 1000000
		if number > 0 {
			result += " "
		}
	}

	// Handle thousands
	if number >= 1000 {
		thousands := number / 1000
		result += ConvertToWordsGeneric(thousands, targetLocale) + " " + GetWordForNumber(1000, targetLocale)
		number %= 1000
		if number > 0 {
			result += " "
		}
	}

	// Handle hundreds
	if number >= 100 {
		hundreds := number / 100
		result += ConvertToWordsGeneric(hundreds, targetLocale) + " " + GetWordForNumber(100, targetLocale)
		number %= 100
		if number > 0 {
			result += " "
		}
	}

	// Handle tens and ones
	if number >= 20 {
		tens := number / 10
		result += GetTensWord(int(tens), targetLocale)
		number %= 10
		if number > 0 {
			result += " " + GetWordForNumber(number, targetLocale)
		}
	} else if number > 0 {
		result += GetWordForNumber(number, targetLocale)
	}

	return strings.TrimSpace(result)
}

// convertToWordsWithExactMapping converts number to words using exact mapping when available
func ConvertToWordsWithExactMapping(number int64, targetLocale locale.NumI18NLocale) string {
	// Check for exact mapping first (used by Filipino locale)
	for _, mapping := range targetLocale.ExactWordsMapping {
		if mapping.Number == number {
			return mapping.Value
		}
	}
	return ConvertToWordsGeneric(number, targetLocale)
}
