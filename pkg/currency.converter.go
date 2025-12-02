package pkg

import (
	"github.com/Lands-Horizon-Corp/numi18n/pkg/locale"
	"github.com/shopspring/decimal"
)

// ToFormat converts a number to localized formatted string
// Example: 1234567.89 -> "1,234,567.89" (US) or "1.234.567,89" (DE)
// For currency: 1234567.89 -> "$1,234,567.89" (US) or "¥1234567" (JP) or "₱1,234,567.89" (PH)
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
