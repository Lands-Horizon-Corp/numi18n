package pkg

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Lands-Horizon-Corp/numi18n/pkg/locale"
)

var Locale = locale.NewNumI18Locales()

// ToWords converts a number to its word representation
// Example: 123.45 -> "one hundred twenty-three and forty-five hundredths"
func (op *NumI18NOptions) ToWords(amount float64) string {
	// Find the appropriate locale
	var targetLocale *locale.NumI18NLocale

	// Try to match locale first
	if op.Locale != "" {
		matches := Locale.FindByLocale(op.Locale)
		if len(matches) > 0 {
			targetLocale = &matches[0]
		}
	}

	// Fallback to currency matching
	if targetLocale == nil && op.Currency != "" {
		matches := Locale.FindByCurrency(op.Currency)
		if len(matches) > 0 {
			// Prefer US locale for USD currency
			for _, match := range matches {
				if match.NumI18Identifier.Locale == "en-US" {
					targetLocale = &match
					break
				}
			}
			if targetLocale == nil {
				targetLocale = &matches[0]
			}
		}
	}

	// Fallback to country name
	if targetLocale == nil && op.CountryName != "" {
		matches := Locale.FindByCountryName(op.CountryName)
		if len(matches) > 0 {
			targetLocale = &matches[0]
		}
	}

	// Default to US locale
	if targetLocale == nil {
		matches := Locale.FindByLocale("en-US")
		if len(matches) > 0 {
			targetLocale = &matches[0]
		}
	}

	if targetLocale == nil {
		return ""
	}

	// Handle negative numbers
	isNegative := amount < 0
	if isNegative {
		amount = math.Abs(amount)
	}

	// Split into whole and fractional parts
	wholePart := int64(amount)
	fractionalPart := amount - float64(wholePart)

	// Convert whole part to words
	wholeWords := convertToWords(wholePart, *targetLocale)

	// Handle different formatting options
	result := wholeWords

	// Add currency if requested
	if op.WordDetails != nil && op.WordDetails.Currency {
		currencyName := targetLocale.Currency.Singular
		currencyPlural := targetLocale.Currency.Plural

		// Apply overrides if provided
		if op.WordDetails.OverrideOptions != nil {
			if op.WordDetails.OverrideOptions.Name != "" {
				currencyName = op.WordDetails.OverrideOptions.Name
			}
			if op.WordDetails.OverrideOptions.Plural != "" {
				currencyPlural = op.WordDetails.OverrideOptions.Plural
			}
		}

		// Handle currency spacing based on locale
		if targetLocale.NumI18Identifier.Language == "ja" {
			// Japanese: no space before currency
			if wholePart == 1 {
				result += currencyName
			} else {
				result += currencyPlural
			}
		} else {
			// Other languages: add space before currency
			if wholePart == 1 {
				result += " " + currencyName
			} else {
				result += " " + currencyPlural
			}
		}
	} // Handle fractional part
	if fractionalPart > 0.001 { // Avoid floating point precision issues
		fractionalStr := fmt.Sprintf("%.2f", fractionalPart)
		fractionalStr = fractionalStr[2:] // Remove "0."
		fractionalValue, _ := strconv.ParseInt(fractionalStr, 10, 64)

		if fractionalValue > 0 {
			fractionalWords := convertToWords(fractionalValue, *targetLocale)

			if op.WordDetails != nil && op.WordDetails.Decimal {
				// Handle "and" text spacing based on locale
				if targetLocale.NumI18Identifier.Language == "ja" {
					result += targetLocale.Texts.And + fractionalWords
				} else {
					result += " " + targetLocale.Texts.And + " " + fractionalWords
				}

				if op.WordDetails.Currency {
					fractionName := targetLocale.Currency.FractionUnit.Singular
					fractionPlural := targetLocale.Currency.FractionUnit.Plural

					// Apply fraction overrides if provided
					if op.WordDetails.OverrideOptions != nil {
						if op.WordDetails.OverrideOptions.FractionUnitName != "" {
							fractionName = op.WordDetails.OverrideOptions.FractionUnitName
						}
						if op.WordDetails.OverrideOptions.FractionPlural != "" {
							fractionPlural = op.WordDetails.OverrideOptions.FractionPlural
						}
					}

					// Handle fraction currency spacing based on locale
					if targetLocale.NumI18Identifier.Language == "ja" {
						// Japanese: no space before fraction currency
						if fractionalValue == 1 {
							result += fractionName
						} else {
							result += fractionPlural
						}
					} else {
						// Other languages: add space before fraction currency
						if fractionalValue == 1 {
							result += " " + fractionName
						} else {
							result += " " + fractionPlural
						}
					}
				}
			}
		}
	}

	// Add negative word
	if isNegative && op.WordDetails != nil && op.WordDetails.NegativeWord {
		negativeWord := targetLocale.Texts.Minus
		// Apply negative word override if provided
		if op.WordDetails.OverrideOptions != nil && op.WordDetails.OverrideOptions.NegativeWord != "" {
			negativeWord = op.WordDetails.OverrideOptions.NegativeWord
		}

		// Handle negative word spacing based on locale
		if targetLocale.NumI18Identifier.Language == "ja" {
			// Japanese: no space after negative word
			result = negativeWord + result
		} else {
			// Other languages: add space after negative word
			result = negativeWord + " " + result
		}
	}

	// Apply text formatting
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

// convertToWords converts an integer to words using the locale's number mapping
func convertToWords(number int64, targetLocale locale.NumI18NLocale) string {
	if number == 0 {
		for _, mapping := range targetLocale.NumberWordsMapping {
			if mapping.Number == 0 {
				return mapping.Value
			}
		}
		return "Zero"
	}

	// Check exact word mappings first
	for _, exact := range targetLocale.ExactWordsMapping {
		if exact.Number == number {
			return exact.Value
		}
	}

	var result []string
	remaining := number

	// Process each number word mapping in order (largest to smallest)
	for _, mapping := range targetLocale.NumberWordsMapping {
		if mapping.Number <= remaining && mapping.Number > 0 {
			count := remaining / mapping.Number
			remaining = remaining % mapping.Number

			if count > 0 {
				if mapping.Number >= 100 {
					// For hundreds, thousands, millions, etc., include the count
					if targetLocale.NumI18Identifier.Language == "ja" {
						// Japanese: for exact matches in mapping, use the mapping value directly
						// Check if this exact number has a direct mapping
						exactMatch := false
						for _, exact := range targetLocale.NumberWordsMapping {
							if exact.Number == mapping.Number*count {
								result = append(result, exact.Value)
								exactMatch = true
								break
							}
						}
						if !exactMatch {
							countWords := convertToWords(count, targetLocale)
							if count == 1 && mapping.Number == 100 {
								result = append(result, mapping.Value) // Just "百" not "一百"
							} else {
								result = append(result, countWords+mapping.Value)
							}
						}
					} else {
						// For other languages, check if exact mapping exists for count * mapping
						exactMatch := false
						for _, exact := range targetLocale.ExactWordsMapping {
							if exact.Number == mapping.Number*count {
								result = append(result, exact.Value)
								exactMatch = true
								break
							}
						}
						if !exactMatch {
							countWords := convertToWords(count, targetLocale)
							result = append(result, countWords+" "+mapping.Value)
						}
					}
				} else {
					// For tens and units, just use the mapping value
					result = append(result, mapping.Value)
				}
			}
		}
	}

	// Handle locale-specific formatting
	if targetLocale.NumI18Identifier.Language == "ja" {
		// Japanese doesn't use spaces between number components
		return strings.Join(result, "")
	}
	return strings.Join(result, " ")
}
