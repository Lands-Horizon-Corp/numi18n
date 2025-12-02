package locale

import (
	"strings"
	"github.com/shopspring/decimal"
)

// ZALocale is a NumI18NLocale configured for South Africa (en-ZA)
var ZALocale = NumI18NLocale{
	LocaleFormatter: &EnglishSouthAfricaFormatter{},
	Currency: Currency{
		Name:     "South African Rand",
		Plural:   "South African Rand",
		Singular: "Rand",
		Symbol:   "R",
		FractionUnit: FractionUnit{
			Name:     "Cent",
			Plural:   "Cents",
			Singular: "Cent",
			Symbol:   "c",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "South Africa",
		Currency:       "ZAR",
		ISO3166Alpha2:  "ZA",
		ISO3166Alpha3:  "ZAF",
		ISO3166Numeric: "710",
		Locale:         "en-ZA",
		Timezone:       []string{"Africa/Johannesburg"},
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

// EnglishSouthAfricaFormatter handles English (South Africa) formatting
type EnglishSouthAfricaFormatter struct{}

func (f *EnglishSouthAfricaFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *EnglishSouthAfricaFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *EnglishSouthAfricaFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *EnglishSouthAfricaFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *EnglishSouthAfricaFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *EnglishSouthAfricaFormatter) ChopDecimal(d decimal.Decimal, precision int) decimal.Decimal {
	return d.Truncate(int32(precision))
}


func (f *EnglishSouthAfricaFormatter) FormatDecimalNumber(amount float64) string {
	return DefaultFormatDecimalNumber(amount, " ", ",")
}
func (f *EnglishSouthAfricaFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	formattedNumber := f.FormatDecimalNumber(amount)
	
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}
	
	// Default currency placement for this locale (prefix with symbol)
	if strings.HasPrefix(formattedNumber, "-") {
		formattedNumber = strings.TrimPrefix(formattedNumber, "-")
		return "-" + currencySymbol + formattedNumber
	}
	
	return currencySymbol + formattedNumber
}
