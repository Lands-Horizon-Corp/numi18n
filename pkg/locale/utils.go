package locale

import (
	"fmt"
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

// DefaultFormatDecimalNumber provides a basic implementation that can be used by any formatter
func DefaultFormatDecimalNumber(amount float64, thousandSeparator string, decimalSeparator string) string {
	decAmount := decimal.NewFromFloat(amount)

	// Handle negative numbers
	isNegative := decAmount.IsNegative()
	if isNegative {
		decAmount = decAmount.Abs()
	}

	// Separate whole and fractional parts
	wholePart := decAmount.Floor()
	fractionalPart := decAmount.Sub(wholePart)

	// Format whole part with separators if provided
	wholeStr := wholePart.String()
	if thousandSeparator != "" && len(wholeStr) > 3 {
		wholeStr = formatWithSeparators(wholeStr, thousandSeparator)
	}

	// Format fractional part if it exists
	result := wholeStr
	if !fractionalPart.IsZero() {
		// Get fractional part as string (remove "0." prefix) and remove trailing zeros
		fractionalStr := fractionalPart.String()
		if len(fractionalStr) > 2 {
			fractionalStr = fractionalStr[2:] // Remove "0."
			// Remove trailing zeros
			fractionalStr = strings.TrimRight(fractionalStr, "0")
			if fractionalStr != "" {
				result += decimalSeparator + fractionalStr
			}
		}
	}

	// Add negative sign if needed
	if isNegative {
		result = "-" + result
	}
	return result
}

// formatWithSeparators adds separators every three digits from right to left
func formatWithSeparators(numStr string, separator string) string {
	if len(numStr) <= 3 {
		return numStr
	}

	var result string
	for i, digit := range numStr {
		if i > 0 && (len(numStr)-i)%3 == 0 {
			result += separator
		}
		result += string(digit)
	}

	return result
}

// FormatDecimalWithSeparators formats a decimal number with thousand separators
func FormatDecimalWithSeparators(amount float64, thousandSeparator string) string {
	decAmount := decimal.NewFromFloat(amount)

	// Handle negative numbers
	isNegative := decAmount.IsNegative()
	if isNegative {
		decAmount = decAmount.Abs()
	}

	// Separate whole and fractional parts
	wholePart := decAmount.Floor()
	fractionalPart := decAmount.Sub(wholePart)

	// Format whole part with thousand separators
	wholeStr := formatWithSeparators(wholePart.String(), thousandSeparator)

	// Format fractional part if it exists
	result := wholeStr
	if !fractionalPart.IsZero() {
		// Get fractional part as string (remove "0." prefix)
		fractionalStr := fractionalPart.String()
		if len(fractionalStr) > 2 {
			fractionalStr = fractionalStr[2:] // Remove "0."
			result += "." + fractionalStr
		}
	}

	// Add negative sign if needed
	if isNegative {
		result = "-" + result
	}

	return result
}

// FormatCurrencyWithSeparators formats currency with thousand separators and removes trailing zeros
func FormatCurrencyWithSeparators(amount float64, thousandSeparator string, symbol string, symbolPrefix bool) string {
	// Use the decimal formatting function which removes trailing zeros
	formattedNumber := DefaultFormatDecimalNumber(amount, thousandSeparator, ".")

	// Handle negative numbers - move negative sign before currency symbol
	if strings.HasPrefix(formattedNumber, "-") {
		formattedNumber = strings.TrimPrefix(formattedNumber, "-")
		if symbolPrefix {
			return "-" + symbol + formattedNumber
		}
		return "-" + formattedNumber + symbol
	}

	if symbolPrefix {
		return symbol + formattedNumber
	}
	return formattedNumber + symbol
}

// FormatUSCurrencyWithSeparators formats US currency with thousand separators and always shows 2 decimal places
func FormatUSCurrencyWithSeparators(amount float64, thousandSeparator string, symbol string, symbolPrefix bool) string {
	// Format number with exactly 2 decimal places for US currency
	formattedNumber := fmt.Sprintf("%.2f", amount)

	// Apply thousands separators if needed (>=1000 or <=-1000)
	if amount >= 1000 || amount <= -1000 {
		formattedNumber = addThousandSeparatorsToFormattedNumber(formattedNumber, thousandSeparator)
	}

	// Handle negative numbers - move negative sign before currency symbol
	if strings.HasPrefix(formattedNumber, "-") {
		formattedNumber = strings.TrimPrefix(formattedNumber, "-")
		if symbolPrefix {
			return "-" + symbol + formattedNumber
		}
		return "-" + formattedNumber + symbol
	}

	if symbolPrefix {
		return symbol + formattedNumber
	}
	return formattedNumber + symbol
} // addThousandSeparatorsToFormattedNumber adds thousand separators to a formatted number string
func addThousandSeparatorsToFormattedNumber(formattedNumber string, separator string) string {
	// Replace decimal point temporarily
	parts := strings.Split(formattedNumber, ".")
	if len(parts) == 2 {
		// Add thousands separators to integer part
		integerPart := parts[0]
		// Handle negative sign
		negative := strings.HasPrefix(integerPart, "-")
		if negative {
			integerPart = strings.TrimPrefix(integerPart, "-")
		}

		formattedInteger := formatWithSeparators(integerPart, separator)

		if negative {
			formattedInteger = "-" + formattedInteger
		}
		return formattedInteger + "." + parts[1]
	}
	return formattedNumber
}

// FormatJapaneseCurrency formats currency for Japanese locale (no separators, no decimals)
func FormatJapaneseCurrency(amount float64, symbol string) string {
	// Format number without decimals for Japanese Yen (no fractional currency)
	formattedNumber := fmt.Sprintf("%.0f", amount)

	// Handle negative numbers - move negative sign before currency symbol
	if strings.HasPrefix(formattedNumber, "-") {
		formattedNumber = strings.TrimPrefix(formattedNumber, "-")
		return "-" + symbol + formattedNumber
	}

	return symbol + formattedNumber
}

// FormatDecimalWithoutSeparators formats a decimal number without thousand separators (for Japanese)
func FormatDecimalWithoutSeparators(amount float64) string {
	return DefaultFormatDecimalNumber(amount, "", ".")
}

// FormatAngloDecimal formats a decimal number with Anglo conventions (comma thousands, period decimal)
func FormatAngloDecimal(amount float64) string {
	return DefaultFormatDecimalNumber(amount, ",", ".")
}

// FormatAngloCurrency formats currency with Anglo conventions (comma thousands, period decimal, prefix symbol)
func FormatAngloCurrency(amount float64, symbol string) string {
	return FormatCurrencyWithSeparators(amount, ",", symbol, true)
}

// FormatAsianDecimal formats a decimal number with Asian conventions (no thousands separators, period decimal)
func FormatAsianDecimal(amount float64) string {
	return DefaultFormatDecimalNumber(amount, "", ".")
}

// FormatAsianCurrency formats currency with Asian conventions (no thousands separators, period decimal, prefix symbol)
func FormatAsianCurrency(amount float64, symbol string) string {
	// Use decimal formatting to remove trailing zeros
	formattedNumber := DefaultFormatDecimalNumber(amount, "", ".")

	// Handle negative numbers - move negative sign before currency symbol
	if strings.HasPrefix(formattedNumber, "-") {
		formattedNumber = strings.TrimPrefix(formattedNumber, "-")
		return "-" + symbol + formattedNumber
	}

	return symbol + formattedNumber
}

// FormatFrenchDecimal formats a number with French locale conventions (space as thousands separator, comma as decimal separator)
func FormatFrenchDecimal(amount float64) string {
	return DefaultFormatDecimalNumber(amount, " ", ",")
}

// FormatFrenchCurrency formats currency with French locale conventions (space separators, comma decimal, prefix currency symbol)
func FormatFrenchCurrency(amount float64, currencySymbol string) string {
	// Use French decimal formatting which removes trailing zeros
	formattedNumber := FormatFrenchDecimal(amount)

	// Handle negative numbers
	if strings.HasPrefix(formattedNumber, "-") {
		formattedNumber = strings.TrimPrefix(formattedNumber, "-")
		return "-" + currencySymbol + formattedNumber
	}

	return currencySymbol + formattedNumber
} // addThousandSeparatorsToFormattedNumberFrench adds space separators for French formatting
func addThousandSeparatorsToFormattedNumberFrench(formattedNumber string) string {
	parts := strings.Split(formattedNumber, ",")
	if len(parts) == 2 {
		integerPart := parts[0]
		// Handle negative sign
		negative := strings.HasPrefix(integerPart, "-")
		if negative {
			integerPart = strings.TrimPrefix(integerPart, "-")
		}

		formattedInteger := formatWithSeparators(integerPart, " ")

		if negative {
			formattedInteger = "-" + formattedInteger
		}
		return formattedInteger + "," + parts[1]
	}
	return formattedNumber
}

// FormatEuropeanDecimal formats a number with European locale conventions (period as thousands separator, comma as decimal separator)
func FormatEuropeanDecimal(amount float64) string {
	return DefaultFormatDecimalNumber(amount, ".", ",")
}

// FormatEuropeanCurrency formats currency with European locale conventions (period separators, comma decimal, prefix currency symbol)
func FormatEuropeanCurrency(amount float64, currencySymbol string) string {
	// Use European decimal formatting which removes trailing zeros
	formattedNumber := FormatEuropeanDecimal(amount)

	// Handle negative numbers
	if strings.HasPrefix(formattedNumber, "-") {
		formattedNumber = strings.TrimPrefix(formattedNumber, "-")
		return "-" + currencySymbol + formattedNumber
	}

	return currencySymbol + formattedNumber
} // addThousandSeparatorsToFormattedNumberEuropean adds period separators for European formatting
func addThousandSeparatorsToFormattedNumberEuropean(formattedNumber string) string {
	parts := strings.Split(formattedNumber, ",")
	if len(parts) == 2 {
		integerPart := parts[0]
		// Handle negative sign
		negative := strings.HasPrefix(integerPart, "-")
		if negative {
			integerPart = strings.TrimPrefix(integerPart, "-")
		}

		formattedInteger := formatWithSeparators(integerPart, ".")

		if negative {
			formattedInteger = "-" + formattedInteger
		}
		return formattedInteger + "," + parts[1]
	}
	return formattedNumber
}

// FormatPolishDecimal formats a number with Polish locale conventions (comma as thousands separator, period as decimal separator)
func FormatPolishDecimal(amount float64) string {
	return DefaultFormatDecimalNumber(amount, ",", ".")
}

// FormatPolishCurrency formats currency with Polish locale conventions (comma separators, period decimal, prefix currency symbol)
func FormatPolishCurrency(amount float64, currencySymbol string) string {
	formattedNumber := FormatPolishDecimal(amount)

	// Handle negative numbers
	if strings.HasPrefix(formattedNumber, "-") {
		formattedNumber = strings.TrimPrefix(formattedNumber, "-")
		return "-" + currencySymbol + formattedNumber
	}

	return currencySymbol + formattedNumber
}
