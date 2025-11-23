package formatter

import "github.com/Lands-Horizon-Corp/numi18n/pkg/locale"

// LocaleFormatter defines the interface for locale-specific formatting
type LocaleFormatter interface {
	FormatNumber(number int64, targetLocale locale.NumI18NLocale) string
	FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string
	FormatFractional(result, fractionalWords string, andText string) string
	FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string
	FormatNegative(result, negativeWord string) string
}
