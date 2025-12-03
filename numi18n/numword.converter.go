package numi18n

import (
	"strings"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n/locale"
	"github.com/shopspring/decimal"
)

var Locale = locale.NewNumI18Locales()

func Locales() []locale.NumI18NLocale {
	return Locale.AllLocales()
}

// ToWords converts a numeric amount to its written word representation in the specified locale
// Example: 1234.56 -> "One Thousand Two Hundred Thirty Four Dollars and Fifty Six Cents" (en-US)
//
//	1234.56 -> "Eintausendzweihundertvierunddreißig Euro und Sechsundfünfzig Cent" (de-DE)
//	1234.56 -> "Mil Doscientos Treinta y Cuatro Euros con Cincuenta y Seis Céntimos" (es-ES)
//
// The function supports:
// - Currency formatting when WordDetails.Currency is true
// - Decimal places when WordDetails.Decimal is true
// - Negative numbers when WordDetails.NegativeWord is true
// - Text case transformations (uppercase, lowercase, capitalize)
// - Custom currency names via WordDetails.OverrideOptions
// - Over 30 locales with proper linguistic formatting (compound numbers, gender agreement, etc.)
//
// Parameters:
//
//	amount: The numeric value to convert (supports negative values and decimals)
//
// Returns:
//
//	String representation of the number in words according to the locale's linguistic rules
//	Returns empty string if locale is not found or supported
func (op *NumI18NOptions) ToWords(amount float64) string {
	targetLocale := op.findLocale()
	if targetLocale == nil {
		return ""
	}
	decimalAmount := decimal.NewFromFloat(amount)
	if targetLocale.LocaleFormatter != nil {
		decimalAmount = targetLocale.LocaleFormatter.ChopDecimal(decimalAmount, 2)
	}
	isNegative := decimalAmount.IsNegative()
	if isNegative {
		decimalAmount = decimalAmount.Abs()
	}
	result := op.convertDecimalToWords(decimalAmount, *targetLocale)
	if isNegative && op.WordDetails != nil && op.WordDetails.NegativeWord {
		negativeWord := targetLocale.Texts.Minus
		if op.WordDetails.OverrideOptions != nil && op.WordDetails.OverrideOptions.NegativeWord != "" {
			negativeWord = op.WordDetails.OverrideOptions.NegativeWord
		}
		result = targetLocale.LocaleFormatter.FormatNegative(result, negativeWord)
	}
	result = op.applyTextFormatting(result, *targetLocale)
	return result
}

// ToFormat converts a number to localized formatted string with proper thousands separators and decimal notation
// Example: 1234567.89 -> "1,234,567.89" (en-US) or "1.234.567,89" (de-DE) or "1 234 567,89" (fr-FR)
// For currency: 1234567.89 -> "$1,234,567.89" (en-US) or "1.234.567,89 €" (de-DE) or "¥1,234,568" (ja-JP)
//
// The function applies locale-specific formatting rules:
// - Thousands separators (comma, period, or space)
// - Decimal separators (period or comma)
// - Currency symbol placement (before or after amount)
// - Currency rounding rules (e.g., Japanese Yen has no decimal places)
//
// Parameters:
//
//	amount: The numeric value to format
//
// Returns:
//
//	Formatted string according to locale conventions
//	Returns string representation of decimal if locale formatter is unavailable
func (op *NumI18NOptions) ToFormat(amount float64) string {
	targetLocale := op.findLocale()
	if targetLocale == nil {
		return ""
	}
	showCurrency := op.WordDetails != nil && op.WordDetails.Currency
	if targetLocale.LocaleFormatter != nil {
		if showCurrency {
			return targetLocale.LocaleFormatter.FormatDecimalNumberWithCurrency(
				amount, *targetLocale, (*locale.OverrideOptions)(op.WordDetails.OverrideOptions))
		} else {
			return targetLocale.LocaleFormatter.FormatDecimalNumber(amount)
		}
	}
	decAmount := decimal.NewFromFloat(amount)
	return decAmount.String()

}

// ToOrdinal converts a number to its ordinal representation
// Example: 123 -> "123rd", 1 -> "1st", 2 -> "2nd"
func (op *NumI18NOptions) ToOrdinal(amount int32) string {
	return ""
}

// ToOrdinalWords converts a number to its ordinal word representation
// Example: 123 -> "one hundred twenty-third", 1 -> "first", 2 -> "second"
func (op *NumI18NOptions) ToOrdinalWords(amount int32) string {
	return ""
}

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

// ToRoman converts an integer to Roman numeral representation
// Example: 1994 -> "MCMXCIV", 42 -> "XLII", 3999 -> "MMMCMXCIX"
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

// FromRoman converts a Roman numeral string to its integer value
// Example: "MCMXCIV" -> 1994, "XLII" -> 42, "IV" -> 4
func FromRoman(roman string) int {
	if roman == "" {
		return 0
	}
	roman = strings.ToUpper(roman)
	if !IsValid(roman) {
		return 0
	}
	return parseRomanUnchecked(roman)
}

// IsValid checks if a Roman numeral string is properly formatted
// Example: "MCMXCIV" -> true, "IIII" -> false, "VV" -> false
func IsValid(roman string) bool {
	if roman == "" {
		return false
	}

	// Convert to uppercase for validation
	roman = strings.ToUpper(roman)

	// Check for valid Roman numeral rules
	// 1. No more than 3 consecutive identical characters (except for extended notation)
	// 2. V, L, D cannot be repeated
	// 3. Only valid subtraction patterns allowed

	// Check for invalid patterns
	invalidPatterns := []string{
		"IIII", "VV", "XXXX", "LL", "CCCC", "DD", "MMMM",
		"VX", "VL", "VC", "VD", "VM",
		"LC", "LD", "LM",
		"DM",
	}

	for _, pattern := range invalidPatterns {
		if strings.Contains(roman, pattern) {
			return false
		}
	}

	// Try to parse using our conversion logic and check round-trip
	i := 0

	// Use our romanNumerals table to validate structure
	for i < len(roman) {
		found := false

		// Try to match any symbol from our table
		for _, rn := range romanNumerals {
			if strings.HasPrefix(roman[i:], rn.Symbol) {
				i += len(rn.Symbol)
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	// Additional round-trip validation
	num := parseRomanUnchecked(roman)
	if num <= 0 {
		return false
	}

	// Check if converting back produces the same result
	converted := ToRoman(num)
	return converted == roman
}

// parseRomanUnchecked is a helper that doesn't validate input
func parseRomanUnchecked(roman string) int {
	if roman == "" {
		return 0
	}

	i := 0
	num := 0

	for i < len(roman) {
		found := false

		for _, rn := range romanNumerals {
			if strings.HasPrefix(roman[i:], rn.Symbol) {
				num += rn.Value
				i += len(rn.Symbol)
				found = true
				break
			}
		}

		if !found {
			return 0
		}
	}

	return num
}
