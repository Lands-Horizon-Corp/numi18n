package pkg

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Lands-Horizon-Corp/numi18n/pkg/formatter"
	"github.com/Lands-Horizon-Corp/numi18n/pkg/locale"
)

var Locale = locale.NewNumI18Locales()

// getLocaleFormatter returns the appropriate formatter for the given locale
func getLocaleFormatter(language string) formatter.LocaleFormatter {
	switch language {
	case "ja":
		return &formatter.JapaneseFormatter{}
	case "ph":
		return &formatter.FilipinoFormatter{}
	default:
		return &formatter.EnglishFormatter{}
	}
}

// ToWords converts a number to its word representation
// Example: 123.45 -> "one hundred twenty-three and forty-five hundredths"
func (op *NumI18NOptions) ToWords(amount float64) string {
	// Find the appropriate locale
	targetLocale := op.findTargetLocale()
	if targetLocale == nil {
		return ""
	}

	// Get the appropriate formatter for this locale
	formatter := getLocaleFormatter(targetLocale.NumI18Identifier.Language)

	// Handle negative numbers
	isNegative := amount < 0
	if isNegative {
		amount = math.Abs(amount)
	}

	// Split into whole and fractional parts
	wholePart := int64(amount)
	fractionalPart := amount - float64(wholePart)

	// Convert whole part to words using formatter
	result := formatter.FormatNumber(wholePart, *targetLocale)

	// Add currency if requested
	if op.WordDetails != nil && op.WordDetails.Currency {
		currencyName, currencyPlural := op.getCurrencyNames(targetLocale)
		result = formatter.FormatCurrency(result, wholePart, currencyName, currencyPlural)
	}

	// Handle fractional part
	if fractionalPart > 0.001 { // Avoid floating point precision issues
		fractionalValue, fractionalWords := op.getFractionalWords(fractionalPart, *targetLocale, formatter)
		if fractionalValue > 0 && op.WordDetails != nil && op.WordDetails.Decimal {
			result = formatter.FormatFractional(result, fractionalWords, targetLocale.Texts.And)

			if op.WordDetails.Currency {
				fractionName, fractionPlural := op.getFractionNames(targetLocale)
				result = formatter.FormatFractionalCurrency(result, fractionalValue, fractionName, fractionPlural)
			}
		}
	}

	// Add negative word
	if isNegative && op.WordDetails != nil && op.WordDetails.NegativeWord {
		negativeWord := op.getNegativeWord(targetLocale)
		result = formatter.FormatNegative(result, negativeWord)
	}

	// Apply text formatting
	return op.applyTextFormatting(result, targetLocale)
}

// findTargetLocale finds the appropriate locale based on options
func (op *NumI18NOptions) findTargetLocale() *locale.NumI18NLocale {
	// Try to match locale first
	if op.Locale != "" {
		matches := Locale.FindByLocale(op.Locale)
		if len(matches) > 0 {
			return &matches[0]
		}
	}

	// Fallback to currency matching
	if op.Currency != "" {
		matches := Locale.FindByCurrency(op.Currency)
		if len(matches) > 0 {
			// Prefer US locale for USD currency
			for _, match := range matches {
				if match.NumI18Identifier.Locale == "en-US" {
					return &match
				}
			}
			return &matches[0]
		}
	}

	// Fallback to country name
	if op.CountryName != "" {
		matches := Locale.FindByCountryName(op.CountryName)
		if len(matches) > 0 {
			return &matches[0]
		}
	}

	// Default to US locale
	matches := Locale.FindByLocale("en-US")
	if len(matches) > 0 {
		return &matches[0]
	}

	return nil
}

// getCurrencyNames returns the currency names with overrides applied
func (op *NumI18NOptions) getCurrencyNames(targetLocale *locale.NumI18NLocale) (string, string) {
	currencyName := targetLocale.Currency.Singular
	currencyPlural := targetLocale.Currency.Plural

	// Apply overrides if provided
	if op.WordDetails != nil && op.WordDetails.OverrideOptions != nil {
		if op.WordDetails.OverrideOptions.Name != "" {
			currencyName = op.WordDetails.OverrideOptions.Name
		}
		if op.WordDetails.OverrideOptions.Plural != "" {
			currencyPlural = op.WordDetails.OverrideOptions.Plural
		}
	}

	return currencyName, currencyPlural
}

// getFractionNames returns the fraction names with overrides applied
func (op *NumI18NOptions) getFractionNames(targetLocale *locale.NumI18NLocale) (string, string) {
	fractionName := targetLocale.Currency.FractionUnit.Singular
	fractionPlural := targetLocale.Currency.FractionUnit.Plural

	// Apply fraction overrides if provided
	if op.WordDetails != nil && op.WordDetails.OverrideOptions != nil {
		if op.WordDetails.OverrideOptions.FractionUnitName != "" {
			fractionName = op.WordDetails.OverrideOptions.FractionUnitName
		}
		if op.WordDetails.OverrideOptions.FractionPlural != "" {
			fractionPlural = op.WordDetails.OverrideOptions.FractionPlural
		}
	}

	return fractionName, fractionPlural
}

// getNegativeWord returns the negative word with overrides applied
func (op *NumI18NOptions) getNegativeWord(targetLocale *locale.NumI18NLocale) string {
	negativeWord := targetLocale.Texts.Minus

	// Apply negative word override if provided
	if op.WordDetails != nil && op.WordDetails.OverrideOptions != nil && op.WordDetails.OverrideOptions.NegativeWord != "" {
		negativeWord = op.WordDetails.OverrideOptions.NegativeWord
	}

	return negativeWord
}

// getFractionalWords converts fractional part to words
func (op *NumI18NOptions) getFractionalWords(fractionalPart float64, targetLocale locale.NumI18NLocale, formatter formatter.LocaleFormatter) (int64, string) {
	fractionalStr := fmt.Sprintf("%.2f", fractionalPart)
	fractionalStr = fractionalStr[2:] // Remove "0."
	fractionalValue, _ := strconv.ParseInt(fractionalStr, 10, 64)

	if fractionalValue > 0 {
		fractionalWords := formatter.FormatNumber(fractionalValue, targetLocale)
		return fractionalValue, fractionalWords
	}

	return 0, ""
}

// applyTextFormatting applies text formatting options
func (op *NumI18NOptions) applyTextFormatting(result string, targetLocale *locale.NumI18NLocale) string {
	if op.WordDetails != nil {
		if op.WordDetails.Capitalize {
			if len(result) > 0 {
				result = strings.ToUpper(string(result[0])) + result[1:]
			}
		} else if op.WordDetails.Uppercase {
			result = strings.ToUpper(result)
		} else if op.WordDetails.Lowercase {
			result = strings.ToLower(result)
		}
		if op.WordDetails.Only {
			result += " " + targetLocale.Texts.Only
		}
	}
	return result
}
