package locale

import (
	"github.com/shopspring/decimal"
)

// IRLocale is a NumI18NLocale configured for Iran (fa-IR)
var IRLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "ریال",
		Plural:   "ریال",
		Singular: "ریال",
		Symbol:   "﷼",
		FractionUnit: FractionUnit{
			Name:     "دینار",
			Plural:   "دینار",
			Singular: "دینار",
			Symbol:   "دینار",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Iran",
		Currency:       "IRR",
		ISO3166Alpha2:  "IR",
		ISO3166Alpha3:  "IRN",
		ISO3166Numeric: "364",
		Locale:         "fa-IR",
		Timezone:       []string{"Asia/Tehran"},
		Language:       "fa",
		Emoji:          "🇮🇷",
	},
	Texts: Texts{
		And:   "و",
		Minus: "منفی",
		Only:  "فقط",
		Point: "نقطه",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "کوادریلیون"},
		{Number: 1000000000000, Value: "تریلیون"},
		{Number: 1000000000, Value: "میلیارد"},
		{Number: 1000000, Value: "میلیون"},
		{Number: 1000, Value: "هزار"},
		{Number: 100, Value: "صد"},
		{Number: 90, Value: "نود"},
		{Number: 80, Value: "هشتاد"},
		{Number: 70, Value: "هفتاد"},
		{Number: 60, Value: "شصت"},
		{Number: 50, Value: "پنجاه"},
		{Number: 40, Value: "چهل"},
		{Number: 30, Value: "سی"},
		{Number: 20, Value: "بیست"},
		{Number: 19, Value: "نوزده"},
		{Number: 18, Value: "هجده"},
		{Number: 17, Value: "هفده"},
		{Number: 16, Value: "شانزده"},
		{Number: 15, Value: "پانزده"},
		{Number: 14, Value: "چهارده"},
		{Number: 13, Value: "سیزده"},
		{Number: 12, Value: "دوازده"},
		{Number: 11, Value: "یازده"},
		{Number: 10, Value: "ده"},
		{Number: 9, Value: "نه"},
		{Number: 8, Value: "هشت"},
		{Number: 7, Value: "هفت"},
		{Number: 6, Value: "شش"},
		{Number: 5, Value: "پنج"},
		{Number: 4, Value: "چهار"},
		{Number: 3, Value: "سه"},
		{Number: 2, Value: "دو"},
		{Number: 1, Value: "یک"},
		{Number: 0, Value: "صفر"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "یک صد"},
		{Number: 1000, Value: "یک هزار"},
		{Number: 1000000, Value: "یک میلیون"},
		{Number: 1000000000, Value: "یک میلیارد"},
		{Number: 1000000000000, Value: "یک تریلیون"},
		{Number: 1000000000000000, Value: "یک کوادریلیون"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "اول", Suffix: "م", Masculine: "اول", Feminine: "اول", Neuter: "اول"},
		{Number: 2, Word: "دوم", Suffix: "م", Masculine: "دوم", Feminine: "دوم", Neuter: "دوم"},
		{Number: 3, Word: "سوم", Suffix: "م", Masculine: "سوم", Feminine: "سوم", Neuter: "سوم"},
		{Number: 4, Word: "چهارم", Suffix: "م", Masculine: "چهارم", Feminine: "چهارم", Neuter: "چهارم"},
		{Number: 5, Word: "پنجم", Suffix: "م", Masculine: "پنجم", Feminine: "پنجم", Neuter: "پنجم"},
		{Number: 6, Word: "ششم", Suffix: "م", Masculine: "ششم", Feminine: "ششم", Neuter: "ششم"},
		{Number: 7, Word: "هفتم", Suffix: "م", Masculine: "هفتم", Feminine: "هفتم", Neuter: "هفتم"},
		{Number: 8, Word: "هشتم", Suffix: "م", Masculine: "هشتم", Feminine: "هشتم", Neuter: "هشتم"},
		{Number: 9, Word: "نهم", Suffix: "م", Masculine: "نهم", Feminine: "نهم", Neuter: "نهم"},
		{Number: 10, Word: "دهم", Suffix: "م", Masculine: "دهم", Feminine: "دهم", Neuter: "دهم"},
		{Number: 11, Word: "یازدهم", Suffix: "م", Masculine: "یازدهم", Feminine: "یازدهم", Neuter: "یازدهم"},
		{Number: 12, Word: "دوازدهم", Suffix: "م", Masculine: "دوازدهم", Feminine: "دوازدهم", Neuter: "دوازدهم"},
		{Number: 13, Word: "سیزدهم", Suffix: "م", Masculine: "سیزدهم", Feminine: "سیزدهم", Neuter: "سیزدهم"},
		{Number: 14, Word: "چهاردهم", Suffix: "م", Masculine: "چهاردهم", Feminine: "چهاردهم", Neuter: "چهاردهم"},
		{Number: 15, Word: "پانزدهم", Suffix: "م", Masculine: "پانزدهم", Feminine: "پانزدهم", Neuter: "پانزدهم"},
		{Number: 16, Word: "شانزدهم", Suffix: "م", Masculine: "شانزدهم", Feminine: "شانزدهم", Neuter: "شانزدهم"},
		{Number: 17, Word: "هفدهم", Suffix: "م", Masculine: "هفدهم", Feminine: "هفدهم", Neuter: "هفدهم"},
		{Number: 18, Word: "هجدهم", Suffix: "م", Masculine: "هجدهم", Feminine: "هجدهم", Neuter: "هجدهم"},
		{Number: 19, Word: "نوزدهم", Suffix: "م", Masculine: "نوزدهم", Feminine: "نوزدهم", Neuter: "نوزدهم"},
		{Number: 20, Word: "بیستم", Suffix: "م", Masculine: "بیستم", Feminine: "بیستم", Neuter: "بیستم"},
		{Number: 21, Word: "بیست و یکم", Suffix: "م", Masculine: "بیست و یکم", Feminine: "بیست و یکم", Neuter: "بیست و یکم"},
		{Number: 30, Word: "سی\u200cام", Suffix: "م", Masculine: "سی\u200cام", Feminine: "سی\u200cام", Neuter: "سی\u200cام"},
		{Number: 40, Word: "چهلم", Suffix: "م", Masculine: "چهلم", Feminine: "چهلم", Neuter: "چهلم"},
		{Number: 50, Word: "پنجاهم", Suffix: "م", Masculine: "پنجاهم", Feminine: "پنجاهم", Neuter: "پنجاهم"},
		{Number: 60, Word: "شصتم", Suffix: "م", Masculine: "شصتم", Feminine: "شصتم", Neuter: "شصتم"},
		{Number: 70, Word: "هفتادم", Suffix: "م", Masculine: "هفتادم", Feminine: "هفتادم", Neuter: "هفتادم"},
		{Number: 80, Word: "هشتادم", Suffix: "م", Masculine: "هشتادم", Feminine: "هشتادم", Neuter: "هشتادم"},
		{Number: 90, Word: "نودم", Suffix: "م", Masculine: "نودم", Feminine: "نودم", Neuter: "نودم"},
		{Number: 100, Word: "صدم", Suffix: "م", Masculine: "صدم", Feminine: "صدم", Neuter: "صدم"},
		{Number: 1000, Word: "هزارم", Suffix: "م", Masculine: "هزارم", Feminine: "هزارم", Neuter: "هزارم"},
	},
	LocaleFormatter: &PersianIranFormatter{},
}

// PersianIranFormatter handles Persian (Iran) formatting
type PersianIranFormatter struct{}

func (f *PersianIranFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *PersianIranFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	// In Persian, currency name comes after the number
	return result + " " + currencyPlural
}

func (f *PersianIranFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *PersianIranFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	// In Persian, fraction unit name comes after the number
	return result + " " + fractionPlural
}

func (f *PersianIranFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *PersianIranFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	return amount.Round(int32(precision))
}

func (f *PersianIranFormatter) FormatDecimalNumber(amount float64) string {
	return FormatPolishDecimal(amount)
}
func (f *PersianIranFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Polish conventions (comma separators, period decimal, prefix symbol)
	return FormatPolishCurrency(amount, currencySymbol)
}
