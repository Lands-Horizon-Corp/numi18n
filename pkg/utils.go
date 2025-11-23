package pkg

import (
	"strings"

	"github.com/Lands-Horizon-Corp/numi18n/pkg/locale"
	"github.com/shopspring/decimal"
)

func (op *NumI18NOptions) findLocale() *locale.NumI18NLocale {
	// Create an identifier from the options
	identifier := locale.NumI18Identifier{
		CountryName:    op.CountryName,
		Currency:       op.Currency,
		ISO3166Alpha2:  op.ISO3166Alpha2,
		ISO3166Alpha3:  op.ISO3166Alpha3,
		ISO3166Numeric: op.ISO3166Numeric,
		Locale:         op.Locale,
		Language:       op.Language,
		Timezone:       []string{op.Timezone},
	}
	foundLocale := Locale.FindMostMatchedLocale(identifier)
	if foundLocale == nil {
		fallbackIdentifier := locale.NumI18Identifier{
			Locale:   "en-US",
			Language: "en",
			Currency: "USD",
		}
		foundLocale = Locale.FindMostMatchedLocale(fallbackIdentifier)
	}

	return foundLocale
}
func (op *NumI18NOptions) convertDecimalToWords(amount decimal.Decimal, targetLocale locale.NumI18NLocale) string {

	wholeNumber := amount.Floor().IntPart()
	fractionalPart := amount.Sub(amount.Floor())

	var result string
	if targetLocale.LocaleFormatter != nil {
		result = targetLocale.LocaleFormatter.FormatNumber(wholeNumber, targetLocale)
	} else {
		result = locale.ConvertToWordsGenericInt64(wholeNumber, targetLocale)
	}
	if op.WordDetails != nil && op.WordDetails.Currency {
		result = op.formatCurrencyPart(result, wholeNumber, targetLocale)
	}
	if op.WordDetails != nil && op.WordDetails.Decimal && !fractionalPart.IsZero() {
		result = op.formatDecimalPart(result, fractionalPart, targetLocale)
	}

	return result
}
func (op *NumI18NOptions) formatCurrencyPart(result string, wholePart int64, targetLocale locale.NumI18NLocale) string {

	currencyName := targetLocale.Currency.Singular
	currencyPlural := targetLocale.Currency.Plural

	if op.WordDetails.OverrideOptions != nil {
		if op.WordDetails.OverrideOptions.Name != "" {
			currencyName = op.WordDetails.OverrideOptions.Name
		}
		if op.WordDetails.OverrideOptions.Plural != "" {
			currencyPlural = op.WordDetails.OverrideOptions.Plural
		}
	}
	if targetLocale.LocaleFormatter != nil {
		return targetLocale.LocaleFormatter.FormatCurrency(result, wholePart, currencyName, currencyPlural)
	}
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}
func (op *NumI18NOptions) formatDecimalPart(result string, fractionalPart decimal.Decimal, targetLocale locale.NumI18NLocale) string {

	multiplier := decimal.NewFromInt(100)
	fractionalValue := fractionalPart.Mul(multiplier).IntPart()

	if fractionalValue == 0 {
		return result
	}
	var fractionalWords string
	if targetLocale.LocaleFormatter != nil {
		fractionalWords = targetLocale.LocaleFormatter.FormatNumber(fractionalValue, targetLocale)
	} else {
		fractionalWords = locale.ConvertToWordsGenericInt64(fractionalValue, targetLocale)
	}

	// Add "And" connector
	andText := targetLocale.Texts.And
	result = targetLocale.LocaleFormatter.FormatFractional(result, fractionalWords, andText)
	if op.WordDetails.Currency {
		fractionName := targetLocale.Currency.FractionUnit.Singular
		fractionPlural := targetLocale.Currency.FractionUnit.Plural
		if op.WordDetails.OverrideOptions != nil {
			if op.WordDetails.OverrideOptions.FractionUnitName != "" {
				fractionName = op.WordDetails.OverrideOptions.FractionUnitName
			}
			if op.WordDetails.OverrideOptions.FractionPlural != "" {
				fractionPlural = op.WordDetails.OverrideOptions.FractionPlural
			}
		}
		result = targetLocale.LocaleFormatter.FormatFractionalCurrency(result, fractionalValue, fractionName, fractionPlural)
	}

	return result
}
func (op *NumI18NOptions) applyTextFormatting(result string, targetLocale locale.NumI18NLocale) string {
	if op.WordDetails == nil {
		return result
	}
	if op.WordDetails.Only {
		onlyText := targetLocale.Texts.Only
		// Japanese and some other languages don't use spaces
		if targetLocale.NumI18Identifier.Language == "ja" ||
			targetLocale.NumI18Identifier.Language == "zh" ||
			targetLocale.NumI18Identifier.Language == "ko" {
			result = result + onlyText
		} else {
			result = result + " " + onlyText
		}
	}
	if op.WordDetails.Uppercase {
		return strings.ToUpper(result)
	} else if op.WordDetails.Lowercase {
		return strings.ToLower(result)
	} else if op.WordDetails.Capitalize {
		return op.capitalizeFirst(result)
	}

	return result
}
func (op *NumI18NOptions) capitalizeFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
