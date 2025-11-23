package formatter

import (
	"strings"

	"github.com/Lands-Horizon-Corp/numi18n/pkg/locale"
)

// FilipinoFormatter handles Filipino-specific formatting
type FilipinoFormatter struct{}

func (f *FilipinoFormatter) FormatNumber(number int64, targetLocale locale.NumI18NLocale) string {
	return convertToWordsFilipino(number, targetLocale)
}

func (f *FilipinoFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *FilipinoFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *FilipinoFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *FilipinoFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

// convertToWordsFilipino converts number to words with Filipino-specific formatting
func convertToWordsFilipino(number int64, targetLocale locale.NumI18NLocale) string {
	if number == 0 {
		return GetWordForNumber(0, targetLocale)
	}

	result := ""

	// Handle billions
	if number >= 1000000000 {
		billions := number / 1000000000
		result += convertToWordsFilipino(billions, targetLocale) + " " + GetWordForNumber(1000000000, targetLocale)
		number %= 1000000000
		if number > 0 {
			result += " "
		}
	}

	// Handle millions
	if number >= 1000000 {
		millions := number / 1000000
		result += convertToWordsFilipino(millions, targetLocale) + " " + GetWordForNumber(1000000, targetLocale)
		number %= 1000000
		if number > 0 {
			result += " "
		}
	}

	// Handle thousands
	if number >= 1000 {
		thousands := number / 1000
		result += convertToWordsFilipino(thousands, targetLocale) + " " + GetWordForNumber(1000, targetLocale)
		number %= 1000
		if number > 0 {
			result += " "
		}
	}

	// Handle hundreds - check exact mapping first
	if number >= 100 {
		hundreds := number / 100
		if hundreds == 1 {
			// Use exact mapping for "Isang daan"
			for _, mapping := range targetLocale.ExactWordsMapping {
				if mapping.Number == 100 {
					result += mapping.Value
					break
				}
			}
		} else {
			result += convertToWordsFilipino(hundreds, targetLocale) + " " + GetWordForNumber(100, targetLocale)
		}
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
			result += " " + getFilipinoWordForNumber(number, targetLocale)
		}
	} else if number > 0 {
		result += getFilipinoWordForNumber(number, targetLocale)
	}

	return strings.TrimSpace(result)
}

// getFilipinoWordForNumber gets the word for a number with Filipino exact mapping support
func getFilipinoWordForNumber(number int64, targetLocale locale.NumI18NLocale) string {
	// Special case for 1 when it's a standalone number (not part of compound)
	if number == 1 {
		return "Isa"
	}

	// Check for exact mapping first
	for _, mapping := range targetLocale.ExactWordsMapping {
		if mapping.Number == number {
			return mapping.Value
		}
	}
	return GetWordForNumber(number, targetLocale)
}
