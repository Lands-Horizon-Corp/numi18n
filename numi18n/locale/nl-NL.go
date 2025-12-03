package locale

import (
	"strings"

	"github.com/shopspring/decimal"
)

// NLNLLocale represents the Dutch (Netherlands) locale
var NLNLLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "euro",
		Plural:   "euro",
		Singular: "euro",
		Symbol:   "€",
		FractionUnit: FractionUnit{
			Name:     "cent",
			Plural:   "cent",
			Singular: "cent",
			Symbol:   "c",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Netherlands",
		Currency:       "EUR",
		ISO3166Alpha2:  "NL",
		ISO3166Alpha3:  "NLD",
		ISO3166Numeric: "528",
		Locale:         "nl-NL",
		Timezone:       []string{"Europe/Amsterdam"},
		Language:       "nl",
		Emoji:          "🇳🇱",
	},
	Texts: Texts{
		And:   "en",
		Minus: "min",
		Only:  "alleen",
		Point: "komma",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "biljard"},
		{Number: 1000000000000, Value: "biljoen"},
		{Number: 1000000000, Value: "miljard"},
		{Number: 1000000, Value: "miljoen"},
		{Number: 1000, Value: "duizend"},
		{Number: 100, Value: "honderd"},
		{Number: 90, Value: "negentig"},
		{Number: 80, Value: "tachtig"},
		{Number: 70, Value: "zeventig"},
		{Number: 60, Value: "zestig"},
		{Number: 50, Value: "vijftig"},
		{Number: 40, Value: "veertig"},
		{Number: 30, Value: "dertig"},
		{Number: 20, Value: "twintig"},
		{Number: 19, Value: "negentien"},
		{Number: 18, Value: "achttien"},
		{Number: 17, Value: "zeventien"},
		{Number: 16, Value: "zestien"},
		{Number: 15, Value: "vijftien"},
		{Number: 14, Value: "veertien"},
		{Number: 13, Value: "dertien"},
		{Number: 12, Value: "twaalf"},
		{Number: 11, Value: "elf"},
		{Number: 10, Value: "tien"},
		{Number: 9, Value: "negen"},
		{Number: 8, Value: "acht"},
		{Number: 7, Value: "zeven"},
		{Number: 6, Value: "zes"},
		{Number: 5, Value: "vijf"},
		{Number: 4, Value: "vier"},
		{Number: 3, Value: "drie"},
		{Number: 2, Value: "twee"},
		{Number: 1, Value: "een"},
		{Number: 0, Value: "nul"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 1000000, Value: "One Million"},
		{Number: 1000, Value: "One Thousand"},
		{Number: 100, Value: "One Hundred"},
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
	LocaleFormatter: &DutchFormatter{},
}

// DutchFormatter handles Dutch (Netherlands) formatting
type DutchFormatter struct{}

func (f *DutchFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return f.convertDutchNumber(number, targetLocale)
}

// convertDutchNumber converts numbers to Dutch compound words
func (f *DutchFormatter) convertDutchNumber(number int64, targetLocale NumI18NLocale) string {
	if number == 0 {
		return "nul"
	}

	if number < 0 {
		return "min " + f.convertDutchNumber(-number, targetLocale)
	}

	// Handle different number ranges
	if number <= 19 {
		return f.handleSingleDigits(number, targetLocale)
	}

	if number < 100 {
		return f.handleTens(number, targetLocale)
	}

	if number < 1000 {
		return f.handleHundreds(number, targetLocale)
	}

	if number < 1000000 {
		return f.handleThousands(number, targetLocale)
	}

	// For larger numbers, fall back to generic conversion but lowercase
	return strings.ToLower(ConvertToWordsWithExactMappingInt64(number, targetLocale))
}

func (f *DutchFormatter) handleSingleDigits(number int64, targetLocale NumI18NLocale) string {
	for _, mapping := range targetLocale.NumberWordsMapping {
		if mapping.Number == number {
			return strings.ToLower(mapping.Value)
		}
	}
	return ""
}

func (f *DutchFormatter) handleTens(number int64, targetLocale NumI18NLocale) string {
	tens := (number / 10) * 10
	ones := number % 10

	tensWord := f.findNumberWord(tens, targetLocale)
	if ones == 0 {
		return tensWord
	}

	onesWord := f.findNumberWord(ones, targetLocale)
	return onesWord + "en" + tensWord
}

func (f *DutchFormatter) handleHundreds(number int64, targetLocale NumI18NLocale) string {
	hundreds := number / 100
	remainder := number % 100

	var result string
	if hundreds == 1 {
		result = "honderd"
	} else {
		hundredsWord := f.convertDutchNumber(hundreds, targetLocale)
		result = hundredsWord + "honderd"
	}

	if remainder > 0 {
		result += f.convertDutchNumber(remainder, targetLocale)
	}

	return result
}

func (f *DutchFormatter) handleThousands(number int64, targetLocale NumI18NLocale) string {
	thousands := number / 1000
	remainder := number % 1000

	var result string
	if thousands == 1 {
		result = "duizend"
	} else {
		thousandsWord := f.convertDutchNumber(thousands, targetLocale)
		result = thousandsWord + "duizend"
	}

	if remainder > 0 {
		result += f.convertDutchNumber(remainder, targetLocale)
	}

	return result
}

func (f *DutchFormatter) findNumberWord(number int64, targetLocale NumI18NLocale) string {
	for _, mapping := range targetLocale.NumberWordsMapping {
		if mapping.Number == number {
			return strings.ToLower(mapping.Value)
		}
	}
	return ""
}

func (f *DutchFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *DutchFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *DutchFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *DutchFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *DutchFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 0
	}
	return amount.Truncate(int32(precision))
}

func (f *DutchFormatter) FormatDecimalNumber(amount float64) string {
	return FormatEuropeanDecimal(amount)
}

func (f *DutchFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with European conventions (period separators, comma decimal, prefix symbol)
	return FormatEuropeanCurrency(amount, currencySymbol)
}
