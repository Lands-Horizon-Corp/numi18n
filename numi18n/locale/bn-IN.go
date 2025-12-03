package locale

import (
	"github.com/shopspring/decimal"
)

// BNINLocale is a NumI18NLocale configured for Bengali (India) - bn-IN
var BNINLocale = NumI18NLocale{
	LocaleFormatter: &BengaliIndiaFormatter{},
	Currency: Currency{
		Name:     "রুপি",
		Plural:   "রুপি",
		Singular: "রুপি",
		Symbol:   "₹",
		FractionUnit: FractionUnit{
			Name:     "পয়সা",
			Plural:   "পয়সা",
			Singular: "পয়সা",
			Symbol:   "p",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "India",
		Currency:       "INR",
		ISO3166Alpha2:  "IN",
		ISO3166Alpha3:  "IND",
		ISO3166Numeric: "356",
		Locale:         "bn-IN",
		Timezone:       []string{"Asia/Kolkata"},
		Language:       "bn",
		Emoji:          "🇮🇳",
		PhoneCode:      "+91",
		Domain:         ".in",
	},
	Texts: Texts{
		And:   "এবং",
		Minus: "মাইনাস",
		Only:  "শুধুমাত্র",
		Point: "দশমিক",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "কোয়াড্রিলিয়ন"},
		{Number: 1000000000000, Value: "ট্রিলিয়ন"},
		{Number: 1000000000, Value: "বিলিয়ন"},
		{Number: 1000000, Value: "মিলিয়ন"},
		{Number: 1000, Value: "হাজার"},
		{Number: 100, Value: "শত"},
		{Number: 90, Value: "নব্বই"},
		{Number: 80, Value: "আশি"},
		{Number: 70, Value: "সত্তর"},
		{Number: 60, Value: "ষাট"},
		{Number: 50, Value: "পঞ্চাশ"},
		{Number: 40, Value: "চল্লিশ"},
		{Number: 30, Value: "ত্রিশ"},
		{Number: 20, Value: "বিশ"},
		{Number: 19, Value: "উনিশ"},
		{Number: 18, Value: "অষ্টাদশ"},
		{Number: 17, Value: "সতের"},
		{Number: 16, Value: "ষোল"},
		{Number: 15, Value: "পনের"},
		{Number: 14, Value: "চৌদ্দ"},
		{Number: 13, Value: "তের"},
		{Number: 12, Value: "বার"},
		{Number: 11, Value: "এঘার"},
		{Number: 10, Value: "দশ"},
		{Number: 9, Value: "নয়"},
		{Number: 8, Value: "আট"},
		{Number: 7, Value: "সাত"},
		{Number: 6, Value: "ছয়"},
		{Number: 5, Value: "পাঁচ"},
		{Number: 4, Value: "চার"},
		{Number: 3, Value: "তিন"},
		{Number: 2, Value: "দুই"},
		{Number: 1, Value: "এক"},
		{Number: 0, Value: "শূন্য"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "একশত"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "প্রথম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 2, Word: "দ্বিতীয়", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 3, Word: "তৃতীয়", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 4, Word: "চতুর্থ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 5, Word: "পঞ্চম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 6, Word: "ষষ্ঠ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 7, Word: "সপ্তম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 8, Word: "অষ্টম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 9, Word: "নবম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 10, Word: "দশম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 11, Word: "একাদশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 12, Word: "দ্বাদশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 13, Word: "ত্রয়োদশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 14, Word: "চতুর্দশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 15, Word: "পঞ্চদশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 16, Word: "ষোড়শ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 17, Word: "সপ্তদশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 18, Word: "অষ্টাদশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 19, Word: "ঊনবিংশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 20, Word: "বিংশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 21, Word: "একবিংশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 30, Word: "ত্রিংশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 40, Word: "চত্বারিংশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 50, Word: "পঞ্চাশত্তম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 60, Word: "ষষ্ঠিত্তম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 70, Word: "সপ্ততিত্তম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 80, Word: "অশীতিত্তম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 90, Word: "নবতিত্তম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 100, Word: "শততম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 1000, Word: "সহস্রত্তম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
	},
}

// BengaliIndiaFormatter handles Bengali (India) formatting
type BengaliIndiaFormatter struct{}

func (f *BengaliIndiaFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *BengaliIndiaFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	// Bengali rupee doesn't change for singular/plural
	return result + " " + currencyName
}

func (f *BengaliIndiaFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *BengaliIndiaFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	// Bengali paisa doesn't change for singular/plural
	return result + " " + fractionName
}

func (f *BengaliIndiaFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *BengaliIndiaFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	return amount.Round(int32(precision))
}

func (f *BengaliIndiaFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *BengaliIndiaFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
