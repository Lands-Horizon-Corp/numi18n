package formatter

import (
	"github.com/Lands-Horizon-Corp/numi18n/pkg/locale"
)

// EnglishFormatter handles English (and default) formatting
type EnglishFormatter struct{}

func (f *EnglishFormatter) FormatNumber(number int64, targetLocale locale.NumI18NLocale) string {
	return ConvertToWordsWithExactMapping(number, targetLocale)
}

func (f *EnglishFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *EnglishFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *EnglishFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *EnglishFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}
