package locale

import (
	"github.com/shopspring/decimal"
)

// ASINLocale is a NumI18NLocale configured for Assamese (India) - as-IN
var ASINLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "ৰুপি",
		Plural:   "ৰুপি",
		Singular: "ৰুপি",
		Symbol:   "₹",
		FractionUnit: FractionUnit{
			Name:     "পাইছা",
			Plural:   "পাইছা",
			Singular: "পাইছা",
			Symbol:   "p",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "India",
		Currency:       "INR",
		ISO3166Alpha2:  "IN",
		ISO3166Alpha3:  "IND",
		ISO3166Numeric: "356",
		Locale:         "as-IN",
		Timezone:       []string{"Asia/Kolkata"},
		Language:       "as",
		Emoji:          "🇮🇳",
	},
	Texts: Texts{
		And:   "আৰু",
		Minus: "মাইনাছ",
		Only:  "মাত্ৰ",
		Point: "ডট",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "কোয়াড্ৰিলিয়ন"},
		{Number: 1000000000000, Value: "ট্ৰিলিয়ন"},
		{Number: 1000000000, Value: "বিলিয়ন"},
		{Number: 1000000, Value: "মিলিয়ন"},
		{Number: 1000, Value: "হাজাৰ"},
		{Number: 100, Value: "শত"},
		{Number: 90, Value: "নৱ্বই"},
		{Number: 80, Value: "আশি"},
		{Number: 70, Value: "সত্তৰ"},
		{Number: 60, Value: "ষাট"},
		{Number: 50, Value: "পঞ্চাশ"},
		{Number: 40, Value: "চাৰিশ"},
		{Number: 30, Value: "ত্ৰিশ"},
		{Number: 20, Value: "বিশ"},
		{Number: 19, Value: "উনিশ"},
		{Number: 18, Value: "অঠাৰ"},
		{Number: 17, Value: "সতের"},
		{Number: 16, Value: "ষোল"},
		{Number: 15, Value: "পন্ধৰ"},
		{Number: 14, Value: "চৌদ্দ"},
		{Number: 13, Value: "তেৰ"},
		{Number: 12, Value: "বার"},
		{Number: 11, Value: "এঘাৰ"},
		{Number: 10, Value: "দহ"},
		{Number: 9, Value: "নয়"},
		{Number: 8, Value: "আঠ"},
		{Number: 7, Value: "সাত"},
		{Number: 6, Value: "ছয়"},
		{Number: 5, Value: "পাঁচ"},
		{Number: 4, Value: "চাৰ"},
		{Number: 3, Value: "তিন"},
		{Number: 2, Value: "দুই"},
		{Number: 1, Value: "এটা"},
		{Number: 0, Value: "শূন্য"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "এক শত"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "প্ৰথম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 2, Word: "দ্বিতীয়", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 3, Word: "তৃতীয়", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 4, Word: "চতুৰ্থ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 5, Word: "পঞ্চম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 6, Word: "ষষ্ঠ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 7, Word: "সপ্তম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 8, Word: "অষ্টম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 9, Word: "নৱম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 10, Word: "দশম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 11, Word: "একাদশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 12, Word: "দ্বাদশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 13, Word: "ত্ৰয়োদশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 14, Word: "চতুৰ্দশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 15, Word: "পঞ্চদশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 16, Word: "ষোড়শ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 17, Word: "সপ্তদশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 18, Word: "অষ্টাদশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 19, Word: "একোনবিংশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 20, Word: "বিংশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 21, Word: "একবিংশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 30, Word: "ত্ৰিংশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 40, Word: "চত্বাৰিংশ", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 50, Word: "পঞ্চাশত্তম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 60, Word: "ষষ্ঠিত্তম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 70, Word: "সপ্ততিত্তম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 80, Word: "অশীতিত্তম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 90, Word: "নৱতিত্তম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 100, Word: "শততম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 1000, Word: "সহস্ৰত্তম", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
	},
	LocaleFormatter: &AssameseFormatter{},
}

// AssameseFormatter handles Assamese (as-IN) formatting
type AssameseFormatter struct{}

func (f *AssameseFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *AssameseFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *AssameseFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *AssameseFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *AssameseFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *AssameseFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 2
	}
	return amount.Truncate(int32(precision))
}

func (f *AssameseFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *AssameseFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
