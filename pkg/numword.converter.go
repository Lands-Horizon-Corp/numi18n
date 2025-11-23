package pkg

import (
	"github.com/Lands-Horizon-Corp/numi18n/pkg/locale"
	"github.com/shopspring/decimal"
)

var Locale = locale.NewNumI18Locales()

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
