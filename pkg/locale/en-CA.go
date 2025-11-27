package locale

import "github.com/shopspring/decimal"

// CALocale is a NumI18NLocale configured for Canada (en-CA)
var CALocale = NumI18NLocale{
	LocaleFormatter: &EnglishCanadaFormatter{},
	Currency: Currency{
		Name:     "Canadian Dollar",
		Plural:   "Canadian Dollars",
		Singular: "Canadian Dollar",
		Symbol:   "C$",
		FractionUnit: FractionUnit{
			Name:     "Cent",
			Plural:   "Cents",
			Singular: "Cent",
			Symbol:   "¢",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Canada",
		Currency:       "CAD",
		ISO3166Alpha2:  "CA",
		ISO3166Alpha3:  "CAN",
		ISO3166Numeric: "124",
		Locale:         "en-CA",
		Timezone:       []string{"America/Toronto"},
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
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "First", Suffix: "st", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 2, Word: "Second", Suffix: "nd", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 3, Word: "Third", Suffix: "rd", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 4, Word: "Fourth", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 5, Word: "Fifth", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 6, Word: "Sixth", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 7, Word: "Seventh", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 8, Word: "Eighth", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 9, Word: "Ninth", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 10, Word: "Tenth", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 11, Word: "Eleventh", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 12, Word: "Twelfth", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 13, Word: "Thirteenth", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 14, Word: "Fourteenth", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 15, Word: "Fifteenth", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 16, Word: "Sixteenth", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 17, Word: "Seventeenth", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 18, Word: "Eighteenth", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 19, Word: "Nineteenth", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 20, Word: "Twentieth", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 21, Word: "Twenty-First", Suffix: "st", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 22, Word: "Twenty-Second", Suffix: "nd", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 23, Word: "Twenty-Third", Suffix: "rd", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 30, Word: "Thirtieth", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 40, Word: "Fortieth", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 50, Word: "Fiftieth", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 100, Word: "Hundredth", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 1000, Word: "Thousandth", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 1000000, Word: "Millionth", Suffix: "th", Masculine: "", Feminine: "", Neuter: ""},
	},
}

// EnglishCanadaFormatter handles English (Canada) formatting
type EnglishCanadaFormatter struct{}

func (f *EnglishCanadaFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *EnglishCanadaFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *EnglishCanadaFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *EnglishCanadaFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *EnglishCanadaFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *EnglishCanadaFormatter) ChopDecimal(d decimal.Decimal, precision int) decimal.Decimal {
	return d.Truncate(int32(precision))
}
