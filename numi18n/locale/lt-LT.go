package locale

import (
	"github.com/shopspring/decimal"
)

// LTLTLocale represents the Lithuanian (Lithuania) locale
var LTLTLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Euro",
		Plural:   "Eurai",
		Singular: "Euras",
		Symbol:   "€",
		FractionUnit: FractionUnit{
			Name:     "Centas",
			Plural:   "Centai",
			Singular: "Centas",
			Symbol:   "c",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Lithuania",
		Currency:       "EUR",
		ISO3166Alpha2:  "LT",
		ISO3166Alpha3:  "LTU",
		ISO3166Numeric: "440",
		Locale:         "lt-LT",
		Timezone:       []string{"Europe/Vilnius"},
		Language:       "lt",
		Emoji:          "🇱🇹",
		PhoneCode:      "+370",
		Domain:         ".lt",
	},
	LocaleFormatter: &LithuanianFormatter{},
	Texts: Texts{
		And:   "ir",
		Minus: "minus",
		Only:  "tik",
		Point: "kablelis",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "kvadrilijonas"},
		{Number: 1000000000000, Value: "trilijonas"},
		{Number: 1000000000, Value: "milijardas"},
		{Number: 1000000, Value: "milijonas"},
		{Number: 100000, Value: "šimtas tūkstančių"},
		{Number: 90000, Value: "devyniasdešimt tūkstančių"},
		{Number: 80000, Value: "aštuoniasdešimt tūkstančių"},
		{Number: 70000, Value: "septyniasdešimt tūkstančių"},
		{Number: 60000, Value: "šešiasdešimt tūkstančių"},
		{Number: 50000, Value: "penkiasdešimt tūkstančių"},
		{Number: 40000, Value: "keturiasdešimt tūkstančių"},
		{Number: 30000, Value: "trisdešimt tūkstančių"},
		{Number: 20000, Value: "dvidešimt tūkstančių"},
		{Number: 19000, Value: "devyniolika tūkstančių"},
		{Number: 18000, Value: "aštuoniolika tūkstančių"},
		{Number: 17000, Value: "septyniolika tūkstančių"},
		{Number: 16000, Value: "šešiolika tūkstančių"},
		{Number: 15000, Value: "penkiolika tūkstančių"},
		{Number: 14000, Value: "keturiolika tūkstančių"},
		{Number: 13000, Value: "trylika tūkstančių"},
		{Number: 12000, Value: "dvylika tūkstančių"},
		{Number: 11000, Value: "vienuolika tūkstančių"},
		{Number: 10000, Value: "dešimt tūkstančių"},
		{Number: 9000, Value: "devyni tūkstančiai"},
		{Number: 8000, Value: "aštuoni tūkstančiai"},
		{Number: 7000, Value: "septyni tūkstančiai"},
		{Number: 6000, Value: "šeši tūkstančiai"},
		{Number: 5000, Value: "penki tūkstančiai"},
		{Number: 4000, Value: "keturi tūkstančiai"},
		{Number: 3000, Value: "trys tūkstančiai"},
		{Number: 2000, Value: "du tūkstančiai"},
		{Number: 1000, Value: "tūkstantis"},
		{Number: 900, Value: "devyni šimtai"},
		{Number: 800, Value: "aštuoni šimtai"},
		{Number: 700, Value: "septyni šimtai"},
		{Number: 600, Value: "šeši šimtai"},
		{Number: 500, Value: "penki šimtai"},
		{Number: 400, Value: "keturi šimtai"},
		{Number: 300, Value: "trys šimtai"},
		{Number: 200, Value: "du šimtai"},
		{Number: 100, Value: "šimtas"},
		{Number: 99, Value: "devyniasdešimt devyni"},
		{Number: 98, Value: "devyniasdešimt aštuoni"},
		{Number: 97, Value: "devyniasdešimt septyni"},
		{Number: 96, Value: "devyniasdešimt šeši"},
		{Number: 95, Value: "devyniasdešimt penki"},
		{Number: 94, Value: "devyniasdešimt keturi"},
		{Number: 93, Value: "devyniasdešimt trys"},
		{Number: 92, Value: "devyniasdešimt du"},
		{Number: 91, Value: "devyniasdešimt vienas"},
		{Number: 90, Value: "devyniasdešimt"},
		{Number: 89, Value: "aštuoniasdešimt devyni"},
		{Number: 88, Value: "aštuoniasdešimt aštuoni"},
		{Number: 87, Value: "aštuoniasdešimt septyni"},
		{Number: 86, Value: "aštuoniasdešimt šeši"},
		{Number: 85, Value: "aštuoniasdešimt penki"},
		{Number: 84, Value: "aštuoniasdešimt keturi"},
		{Number: 83, Value: "aštuoniasdešimt trys"},
		{Number: 82, Value: "aštuoniasdešimt du"},
		{Number: 81, Value: "aštuoniasdešimt vienas"},
		{Number: 80, Value: "aštuoniasdešimt"},
		{Number: 79, Value: "septyniasdešimt devyni"},
		{Number: 78, Value: "septyniasdešimt aštuoni"},
		{Number: 77, Value: "septyniasdešimt septyni"},
		{Number: 76, Value: "septyniasdešimt šeši"},
		{Number: 75, Value: "septyniasdešimt penki"},
		{Number: 74, Value: "septyniasdešimt keturi"},
		{Number: 73, Value: "septyniasdešimt trys"},
		{Number: 72, Value: "septyniasdešimt du"},
		{Number: 71, Value: "septyniasdešimt vienas"},
		{Number: 70, Value: "septyniasdešimt"},
		{Number: 69, Value: "šešiasdešimt devyni"},
		{Number: 68, Value: "šešiasdešimt aštuoni"},
		{Number: 67, Value: "šešiasdešimt septyni"},
		{Number: 66, Value: "šešiasdešimt šeši"},
		{Number: 65, Value: "šešiasdešimt penki"},
		{Number: 64, Value: "šešiasdešimt keturi"},
		{Number: 63, Value: "šešiasdešimt trys"},
		{Number: 62, Value: "šešiasdešimt du"},
		{Number: 61, Value: "šešiasdešimt vienas"},
		{Number: 60, Value: "šešiasdešimt"},
		{Number: 59, Value: "penkiasdešimt devyni"},
		{Number: 58, Value: "penkiasdešimt aštuoni"},
		{Number: 57, Value: "penkiasdešimt septyni"},
		{Number: 56, Value: "penkiasdešimt šeši"},
		{Number: 55, Value: "penkiasdešimt penki"},
		{Number: 54, Value: "penkiasdešimt keturi"},
		{Number: 53, Value: "penkiasdešimt trys"},
		{Number: 52, Value: "penkiasdešimt du"},
		{Number: 51, Value: "penkiasdešimt vienas"},
		{Number: 50, Value: "penkiasdešimt"},
		{Number: 49, Value: "keturiasdešimt devyni"},
		{Number: 48, Value: "keturiasdešimt aštuoni"},
		{Number: 47, Value: "keturiasdešimt septyni"},
		{Number: 46, Value: "keturiasdešimt šeši"},
		{Number: 45, Value: "keturiasdešimt penki"},
		{Number: 44, Value: "keturiasdešimt keturi"},
		{Number: 43, Value: "keturiasdešimt trys"},
		{Number: 42, Value: "keturiasdešimt du"},
		{Number: 41, Value: "keturiasdešimt vienas"},
		{Number: 40, Value: "keturiasdešimt"},
		{Number: 39, Value: "trisdešimt devyni"},
		{Number: 38, Value: "trisdešimt aštuoni"},
		{Number: 37, Value: "trisdešimt septyni"},
		{Number: 36, Value: "trisdešimt šeši"},
		{Number: 35, Value: "trisdešimt penki"},
		{Number: 34, Value: "trisdešimt keturi"},
		{Number: 33, Value: "trisdešimt trys"},
		{Number: 32, Value: "trisdešimt du"},
		{Number: 31, Value: "trisdešimt vienas"},
		{Number: 30, Value: "trisdešimt"},
		{Number: 29, Value: "dvidešimt devyni"},
		{Number: 28, Value: "dvidešimt aštuoni"},
		{Number: 27, Value: "dvidešimt septyni"},
		{Number: 26, Value: "dvidešimt šeši"},
		{Number: 25, Value: "dvidešimt penki"},
		{Number: 24, Value: "dvidešimt keturi"},
		{Number: 23, Value: "dvidešimt trys"},
		{Number: 22, Value: "dvidešimt du"},
		{Number: 21, Value: "dvidešimt vienas"},
		{Number: 20, Value: "dvidešimt"},
		{Number: 19, Value: "devyniolika"},
		{Number: 18, Value: "aštuoniolika"},
		{Number: 17, Value: "septyniolika"},
		{Number: 16, Value: "šešiolika"},
		{Number: 15, Value: "penkiolika"},
		{Number: 14, Value: "keturiolika"},
		{Number: 13, Value: "trylika"},
		{Number: 12, Value: "dvylika"},
		{Number: 11, Value: "vienuolika"},
		{Number: 10, Value: "dešimt"},
		{Number: 9, Value: "devyni"},
		{Number: 8, Value: "aštuoni"},
		{Number: 7, Value: "septyni"},
		{Number: 6, Value: "šeši"},
		{Number: 5, Value: "penki"},
		{Number: 4, Value: "keturi"},
		{Number: 3, Value: "trys"},
		{Number: 2, Value: "du"},
		{Number: 1, Value: "vienas"},
		{Number: 0, Value: "nulis"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Vienas šimtas"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "pirmas", Suffix: "-as", Masculine: "pirmas", Feminine: "pirma", Neuter: "pirma"},
		{Number: 2, Word: "antras", Suffix: "-as", Masculine: "antras", Feminine: "antra", Neuter: "antra"},
		{Number: 3, Word: "trečias", Suffix: "-ias", Masculine: "trečias", Feminine: "trečia", Neuter: "trečia"},
		{Number: 4, Word: "ketvirtas", Suffix: "-as", Masculine: "ketvirtas", Feminine: "ketvirta", Neuter: "ketvirta"},
		{Number: 5, Word: "penktas", Suffix: "-as", Masculine: "penktas", Feminine: "penkta", Neuter: "penkta"},
		{Number: 6, Word: "šeštas", Suffix: "-as", Masculine: "šeštas", Feminine: "šešta", Neuter: "šešta"},
		{Number: 7, Word: "septintas", Suffix: "-as", Masculine: "septintas", Feminine: "septinta", Neuter: "septinta"},
		{Number: 8, Word: "aštuntas", Suffix: "-as", Masculine: "aštuntas", Feminine: "aštunta", Neuter: "aštunta"},
		{Number: 9, Word: "devintas", Suffix: "-as", Masculine: "devintas", Feminine: "devinta", Neuter: "devinta"},
		{Number: 10, Word: "dešimtas", Suffix: "-as", Masculine: "dešimtas", Feminine: "dešimta", Neuter: "dešimta"},
		{Number: 11, Word: "vienuoliktas", Suffix: "-as", Masculine: "vienuoliktas", Feminine: "vienuolikta", Neuter: "vienuolikta"},
		{Number: 12, Word: "dvyliktas", Suffix: "-as", Masculine: "dvyliktas", Feminine: "dvylikta", Neuter: "dvylikta"},
		{Number: 20, Word: "dvidešimtas", Suffix: "-as", Masculine: "dvidešimtas", Feminine: "dvidešimta", Neuter: "dvidešimta"},
		{Number: 21, Word: "dvidešimt pirmas", Suffix: "-as", Masculine: "dvidešimt pirmas", Feminine: "dvidešimt pirma", Neuter: "dvidešimt pirma"},
		{Number: 30, Word: "trisdešimtas", Suffix: "-as", Masculine: "trisdešimtas", Feminine: "trisdešimta", Neuter: "trisdešimta"},
		{Number: 100, Word: "šimtasis", Suffix: "-asis", Masculine: "šimtasis", Feminine: "šimtoji", Neuter: "šimtasis"},
		{Number: 1000, Word: "tūkstantasis", Suffix: "-asis", Masculine: "tūkstantasis", Feminine: "tūkstantoji", Neuter: "tūkstantasis"},
	},
}

// LithuanianFormatter handles Lithuanian formatting
type LithuanianFormatter struct{}

func (f *LithuanianFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *LithuanianFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *LithuanianFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *LithuanianFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *LithuanianFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *LithuanianFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	return amount.Truncate(int32(precision))
}

func (f *LithuanianFormatter) FormatDecimalNumber(amount float64) string {
	return FormatEuropeanDecimal(amount)
}
func (f *LithuanianFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with European conventions (period separators, comma decimal, prefix symbol)
	return FormatEuropeanCurrency(amount, currencySymbol)
}
