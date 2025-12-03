package locale

import (
	"github.com/shopspring/decimal"
)

// AR001Locale is a NumI18NLocale configured for Arabic (World) - ar-001
var AR001Locale = NumI18NLocale{
	Currency: Currency{
		Name:     "دولار",
		Plural:   "دولارات",
		Singular: "دولار",
		Symbol:   "$",
		FractionUnit: FractionUnit{
			Name:     "سنت",
			Plural:   "سنتات",
			Singular: "سنت",
			Symbol:   "¢",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "World",
		Currency:       "USD",
		ISO3166Alpha2:  "001",
		ISO3166Alpha3:  "WLD",
		ISO3166Numeric: "001",
		Locale:         "ar-001",
		Timezone:       []string{"UTC"},
		Language:       "ar",
		Emoji:          "🌍",
	},
	Texts: Texts{
		And:   "و",
		Minus: "ناقص",
		Only:  "فقط",
		Point: "فاصل",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "كوادريليون"},
		{Number: 1000000000000, Value: "تريليون"},
		{Number: 1000000000, Value: "مليار"},
		{Number: 1000000, Value: "مليون"},
		{Number: 1000, Value: "ألف"},
		{Number: 100, Value: "مئة"},
		{Number: 90, Value: "تسعون"},
		{Number: 80, Value: "ثمانون"},
		{Number: 70, Value: "سبعون"},
		{Number: 60, Value: "ستون"},
		{Number: 50, Value: "خمسون"},
		{Number: 40, Value: "أربعون"},
		{Number: 30, Value: "ثلاثون"},
		{Number: 20, Value: "عشرون"},
		{Number: 19, Value: "تسعة عشر"},
		{Number: 18, Value: "ثمانية عشر"},
		{Number: 17, Value: "سبعة عشر"},
		{Number: 16, Value: "ستة عشر"},
		{Number: 15, Value: "خمسة عشر"},
		{Number: 14, Value: "أربعة عشر"},
		{Number: 13, Value: "ثلاثة عشر"},
		{Number: 12, Value: "اثنا عشر"},
		{Number: 11, Value: "أحد عشر"},
		{Number: 10, Value: "عشرة"},
		{Number: 9, Value: "تسعة"},
		{Number: 8, Value: "ثمانية"},
		{Number: 7, Value: "سبعة"},
		{Number: 6, Value: "ستة"},
		{Number: 5, Value: "خمسة"},
		{Number: 4, Value: "أربعة"},
		{Number: 3, Value: "ثلاثة"},
		{Number: 2, Value: "اثنان"},
		{Number: 1, Value: "واحد"},
		{Number: 0, Value: "صفر"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "مئة"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "الأول", Suffix: "", Masculine: "الأول", Feminine: "الأولى", Neuter: ""},
		{Number: 2, Word: "الثاني", Suffix: "", Masculine: "الثاني", Feminine: "الثانية", Neuter: ""},
		{Number: 3, Word: "الثالث", Suffix: "", Masculine: "الثالث", Feminine: "الثالثة", Neuter: ""},
		{Number: 4, Word: "الرابع", Suffix: "", Masculine: "الرابع", Feminine: "الرابعة", Neuter: ""},
		{Number: 5, Word: "الخامس", Suffix: "", Masculine: "الخامس", Feminine: "الخامسة", Neuter: ""},
		{Number: 6, Word: "السادس", Suffix: "", Masculine: "السادس", Feminine: "السادسة", Neuter: ""},
		{Number: 7, Word: "السابع", Suffix: "", Masculine: "السابع", Feminine: "السابعة", Neuter: ""},
		{Number: 8, Word: "الثامن", Suffix: "", Masculine: "الثامن", Feminine: "الثامنة", Neuter: ""},
		{Number: 9, Word: "التاسع", Suffix: "", Masculine: "التاسع", Feminine: "التاسعة", Neuter: ""},
		{Number: 10, Word: "العاشر", Suffix: "", Masculine: "العاشر", Feminine: "العاشرة", Neuter: ""},
		{Number: 11, Word: "الحادي عشر", Suffix: "", Masculine: "الحادي عشر", Feminine: "الحادية عشرة", Neuter: ""},
		{Number: 12, Word: "الثاني عشر", Suffix: "", Masculine: "الثاني عشر", Feminine: "الثانية عشرة", Neuter: ""},
		{Number: 13, Word: "الثالث عشر", Suffix: "", Masculine: "الثالث عشر", Feminine: "الثالثة عشرة", Neuter: ""},
		{Number: 14, Word: "الرابع عشر", Suffix: "", Masculine: "الرابع عشر", Feminine: "الرابعة عشرة", Neuter: ""},
		{Number: 15, Word: "الخامس عشر", Suffix: "", Masculine: "الخامس عشر", Feminine: "الخامسة عشرة", Neuter: ""},
		{Number: 16, Word: "السادس عشر", Suffix: "", Masculine: "السادس عشر", Feminine: "السادسة عشرة", Neuter: ""},
		{Number: 17, Word: "السابع عشر", Suffix: "", Masculine: "السابع عشر", Feminine: "السابعة عشرة", Neuter: ""},
		{Number: 18, Word: "الثامن عشر", Suffix: "", Masculine: "الثامن عشر", Feminine: "الثامنة عشرة", Neuter: ""},
		{Number: 19, Word: "التاسع عشر", Suffix: "", Masculine: "التاسع عشر", Feminine: "التاسعة عشرة", Neuter: ""},
		{Number: 20, Word: "العشرون", Suffix: "", Masculine: "العشرون", Feminine: "العشرون", Neuter: ""},
		{Number: 21, Word: "الحادي والعشرون", Suffix: "", Masculine: "الحادي والعشرون", Feminine: "الحادية والعشرون", Neuter: ""},
		{Number: 30, Word: "الثلاثون", Suffix: "", Masculine: "الثلاثون", Feminine: "الثلاثون", Neuter: ""},
		{Number: 40, Word: "الأربعون", Suffix: "", Masculine: "الأربعون", Feminine: "الأربعون", Neuter: ""},
		{Number: 50, Word: "الخمسون", Suffix: "", Masculine: "الخمسون", Feminine: "الخمسون", Neuter: ""},
		{Number: 60, Word: "الستون", Suffix: "", Masculine: "الستون", Feminine: "الستون", Neuter: ""},
		{Number: 70, Word: "السبعون", Suffix: "", Masculine: "السبعون", Feminine: "السبعون", Neuter: ""},
		{Number: 80, Word: "الثمانون", Suffix: "", Masculine: "الثمانون", Feminine: "الثمانون", Neuter: ""},
		{Number: 90, Word: "التسعون", Suffix: "", Masculine: "التسعون", Feminine: "التسعون", Neuter: ""},
		{Number: 100, Word: "المئة", Suffix: "", Masculine: "المئة", Feminine: "المئة", Neuter: ""},
		{Number: 1000, Word: "الألف", Suffix: "", Masculine: "الألف", Feminine: "الألف", Neuter: ""},
	},
	LocaleFormatter: &ArabicFormatter{},
}

// ArabicFormatter handles Arabic (ar-001) formatting
type ArabicFormatter struct{}

func (f *ArabicFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *ArabicFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	// In Arabic, currency formatting depends on the number
	// 1: واحد دولار (one dollar)
	// 2: دولاران (two dollars - dual form)
	// 3-10: ثلاثة دولارات (three dollars - plural)
	// 11+: أحد عشر دولارًا (eleven dollars - accusative plural)
	switch wholePart {
	case 1:
		return result + " " + currencyName
	case 2:
		// Arabic dual form - but we'll use plural for simplicity
		return result + " " + currencyPlural
	default:
		return result + " " + currencyPlural
	}
}

func (f *ArabicFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	// In Arabic, "و" (and) connects the whole and fractional parts
	return result + " " + andText + " " + fractionalWords
}

func (f *ArabicFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	// Similar to currency formatting
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *ArabicFormatter) FormatNegative(result, negativeWord string) string {
	// In Arabic, negative word (ناقص) comes before the number
	return negativeWord + " " + result
}

func (f *ArabicFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 2
	}
	return amount.Truncate(int32(precision))
}

func (f *ArabicFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}

func (f *ArabicFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma thousands, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
