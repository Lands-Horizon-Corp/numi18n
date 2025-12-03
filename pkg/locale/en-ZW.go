package locale

import (
	"github.com/shopspring/decimal"
)

// ZWLocale is a NumI18NLocale configured for Zimbabwe (en-ZW)
var ZWLocale = NumI18NLocale{
	LocaleFormatter: &EnglishZimbabweFormatter{},
	Currency: Currency{
		Name:     "Zimbabwean Dollar",
		Plural:   "Zimbabwean Dollars",
		Singular: "Zimbabwean Dollar",
		Symbol:   "ZWL$",
		FractionUnit: FractionUnit{
			Name:     "Cent",
			Plural:   "Cents",
			Singular: "Cent",
			Symbol:   "c",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Zimbabwe",
		Currency:       "ZWL",
		ISO3166Alpha2:  "ZW",
		ISO3166Alpha3:  "ZWE",
		ISO3166Numeric: "716",
		Locale:         "en-ZW",
		Timezone:       []string{"Africa/Harare"},
		Language:       "en",
	},
	Texts: Texts{
		And:   "And",
		Minus: "Minus",
		Only:  "Only",
		Point: "Point",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Quadrillion"},
		{Number: 1000000000000, Value: "Trillion"},
		{Number: 1000000000, Value: "Billion"},
		{Number: 1000000, Value: "Million"},
		{Number: 1000, Value: "Thousand"},
		{Number: 100, Value: "Hundred"},
		{Number: 90, Value: "Ninety"},
		{Number: 80, Value: "Eighty"},
		{Number: 70, Value: "Seventy"},
		{Number: 60, Value: "Sixty"},
		{Number: 50, Value: "Fifty"},
		{Number: 40, Value: "Forty"},
		{Number: 30, Value: "Thirty"},
		{Number: 20, Value: "Twenty"},
		{Number: 19, Value: "Nineteen"},
		{Number: 18, Value: "Eighteen"},
		{Number: 17, Value: "Seventeen"},
		{Number: 16, Value: "Sixteen"},
		{Number: 15, Value: "Fifteen"},
		{Number: 14, Value: "Fourteen"},
		{Number: 13, Value: "Thirteen"},
		{Number: 12, Value: "Twelve"},
		{Number: 11, Value: "Eleven"},
		{Number: 10, Value: "Ten"},
		{Number: 9, Value: "Nine"},
		{Number: 8, Value: "Eight"},
		{Number: 7, Value: "Seven"},
		{Number: 6, Value: "Six"},
		{Number: 5, Value: "Five"},
		{Number: 4, Value: "Four"},
		{Number: 3, Value: "Three"},
		{Number: 2, Value: "Two"},
		{Number: 1, Value: "One"},
		{Number: 0, Value: "Zero"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "One Hundred"},
		{Number: 1000, Value: "One Thousand"},
		{Number: 1000000, Value: "One Million"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "First", Suffix: "st"},
		{Number: 2, Word: "Second", Suffix: "nd"},
		{Number: 3, Word: "Third", Suffix: "rd"},
		{Number: 4, Word: "Fourth", Suffix: "th"},
		{Number: 5, Word: "Fifth", Suffix: "th"},
		{Number: 6, Word: "Sixth", Suffix: "th"},
		{Number: 7, Word: "Seventh", Suffix: "th"},
		{Number: 8, Word: "Eighth", Suffix: "th"},
		{Number: 9, Word: "Ninth", Suffix: "th"},
		{Number: 10, Word: "Tenth", Suffix: "th"},
		{Number: 11, Word: "Eleventh", Suffix: "th"},
		{Number: 12, Word: "Twelfth", Suffix: "th"},
		{Number: 13, Word: "Thirteenth", Suffix: "th"},
		{Number: 14, Word: "Fourteenth", Suffix: "th"},
		{Number: 15, Word: "Fifteenth", Suffix: "th"},
		{Number: 16, Word: "Sixteenth", Suffix: "th"},
		{Number: 17, Word: "Seventeenth", Suffix: "th"},
		{Number: 18, Word: "Eighteenth", Suffix: "th"},
		{Number: 19, Word: "Nineteenth", Suffix: "th"},
		{Number: 20, Word: "Twentieth", Suffix: "th"},
		{Number: 21, Word: "Twenty-First", Suffix: "st"},
		{Number: 22, Word: "Twenty-Second", Suffix: "nd"},
		{Number: 23, Word: "Twenty-Third", Suffix: "rd"},
		{Number: 30, Word: "Thirtieth", Suffix: "th"},
		{Number: 40, Word: "Fortieth", Suffix: "th"},
		{Number: 50, Word: "Fiftieth", Suffix: "th"},
		{Number: 60, Word: "Sixtieth", Suffix: "th"},
		{Number: 70, Word: "Seventieth", Suffix: "th"},
		{Number: 80, Word: "Eightieth", Suffix: "th"},
		{Number: 90, Word: "Ninetieth", Suffix: "th"},
		{Number: 100, Word: "One Hundredth", Suffix: "th"},
		{Number: 1000, Word: "One Thousandth", Suffix: "th"},
		{Number: 1000000, Word: "One Millionth", Suffix: "th"},
	},
}

// EnglishZimbabweFormatter handles English (Zimbabwe) formatting
type EnglishZimbabweFormatter struct{}

func (f *EnglishZimbabweFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *EnglishZimbabweFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *EnglishZimbabweFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *EnglishZimbabweFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *EnglishZimbabweFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *EnglishZimbabweFormatter) ChopDecimal(d decimal.Decimal, precision int) decimal.Decimal {
	return d.Truncate(int32(precision))
}

func (f *EnglishZimbabweFormatter) FormatDecimalNumber(amount float64) string {
	return FormatDecimalWithSeparators(amount, ",")
}

func (f *EnglishZimbabweFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with thousand separators, currency symbol, and symbol prefix (true for Zimbabwe)
	return FormatCurrencyWithSeparators(amount, ",", currencySymbol, true)
}
