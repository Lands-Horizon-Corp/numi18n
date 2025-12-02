package locale

import (
	"strings"
	"github.com/shopspring/decimal"
)

// TNBWLocale represents the Tswana (Botswana) locale
var TNBWLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Botswana Pula",
		Plural:   "Dipula",
		Singular: "Pula",
		Symbol:   "P",
		FractionUnit: FractionUnit{
			Name:     "Thebe",
			Plural:   "Dithebe",
			Singular: "Thebe",
			Symbol:   "t",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Botswana",
		Currency:       "BWP",
		ISO3166Alpha2:  "BW",
		ISO3166Alpha3:  "BWA",
		ISO3166Numeric: "072",
		Locale:         "tn-BW",
		Timezone:       []string{"Africa/Gaborone"},
		Language:       "tn",
	},
	Texts: Texts{
		And:   "le",
		Minus: "tlosa",
		Only:  "fela",
		Point: "tshupo",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "sekete sa dibiliyone"},
		{Number: 1000000000000, Value: "biliyone"},
		{Number: 1000000000, Value: "sekete sa dimiliyone"},
		{Number: 1000000, Value: "miliyone"},
		{Number: 100000, Value: "lekgolo la dikete"},
		{Number: 90000, Value: "masome a robongwe sekete"},
		{Number: 80000, Value: "masome a robedi sekete"},
		{Number: 70000, Value: "masome a supa sekete"},
		{Number: 60000, Value: "masome a thataro le masome sekete"},
		{Number: 50000, Value: "masome a tlhano sekete"},
		{Number: 40000, Value: "masome a ne sekete"},
		{Number: 30000, Value: "masome a tharo sekete"},
		{Number: 20000, Value: "masome a mabedi sekete"},
		{Number: 19000, Value: "lesome le robongwe sekete"},
		{Number: 18000, Value: "lesome le robedi sekete"},
		{Number: 17000, Value: "lesome le supa sekete"},
		{Number: 16000, Value: "lesome le thataro le masome sekete"},
		{Number: 15000, Value: "lesome le tlhano sekete"},
		{Number: 14000, Value: "lesome le ne sekete"},
		{Number: 13000, Value: "lesome le tharo sekete"},
		{Number: 12000, Value: "lesome le pedi sekete"},
		{Number: 11000, Value: "lesome le nngwe sekete"},
		{Number: 10000, Value: "lesome sekete"},
		{Number: 9000, Value: "robongwe sekete"},
		{Number: 8000, Value: "robedi sekete"},
		{Number: 7000, Value: "supa sekete"},
		{Number: 6000, Value: "thataro sekete"},
		{Number: 5000, Value: "tlhano sekete"},
		{Number: 4000, Value: "ne sekete"},
		{Number: 3000, Value: "tharo sekete"},
		{Number: 2000, Value: "pedi sekete"},
		{Number: 1000, Value: "sekete"},
		{Number: 900, Value: "makgolo a robongwe"},
		{Number: 800, Value: "makgolo a robedi"},
		{Number: 700, Value: "makgolo a supa"},
		{Number: 600, Value: "makgolo a thataro"},
		{Number: 500, Value: "makgolo a tlhano"},
		{Number: 400, Value: "makgolo a ne"},
		{Number: 300, Value: "makgolo a tharo"},
		{Number: 200, Value: "makgolo a mabedi"},
		{Number: 100, Value: "lekgolo"},
		{Number: 99, Value: "masome a robongwe le robongwe"},
		{Number: 98, Value: "masome a robongwe le robedi"},
		{Number: 97, Value: "masome a robongwe le supa"},
		{Number: 96, Value: "masome a robongwe le thataro"},
		{Number: 95, Value: "masome a robongwe le tlhano"},
		{Number: 94, Value: "masome a robongwe le ne"},
		{Number: 93, Value: "masome a robongwe le tharo"},
		{Number: 92, Value: "masome a robongwe le pedi"},
		{Number: 91, Value: "masome a robongwe le nngwe"},
		{Number: 90, Value: "masome a robongwe"},
		{Number: 89, Value: "masome a robedi le robongwe"},
		{Number: 88, Value: "masome a robedi le robedi"},
		{Number: 87, Value: "masome a robedi le supa"},
		{Number: 86, Value: "masome a robedi le thataro"},
		{Number: 85, Value: "masome a robedi le tlhano"},
		{Number: 84, Value: "masome a robedi le ne"},
		{Number: 83, Value: "masome a robedi le tharo"},
		{Number: 82, Value: "masome a robedi le pedi"},
		{Number: 81, Value: "masome a robedi le nngwe"},
		{Number: 80, Value: "masome a robedi"},
		{Number: 79, Value: "masome a supa le robongwe"},
		{Number: 78, Value: "masome a supa le robedi"},
		{Number: 77, Value: "masome a supa le supa"},
		{Number: 76, Value: "masome a supa le thataro"},
		{Number: 75, Value: "masome a supa le tlhano"},
		{Number: 74, Value: "masome a supa le ne"},
		{Number: 73, Value: "masome a supa le tharo"},
		{Number: 72, Value: "masome a supa le pedi"},
		{Number: 71, Value: "masome a supa le nngwe"},
		{Number: 70, Value: "masome a supa"},
		{Number: 69, Value: "masome a thataro le masome le robongwe"},
		{Number: 68, Value: "masome a thataro le masome le robedi"},
		{Number: 67, Value: "masome a thataro le masome le supa"},
		{Number: 66, Value: "masome a thataro le masome le thataro"},
		{Number: 65, Value: "masome a thataro le masome le tlhano"},
		{Number: 64, Value: "masome a thataro le masome le ne"},
		{Number: 63, Value: "masome a thataro le masome le tharo"},
		{Number: 62, Value: "masome a thataro le masome le pedi"},
		{Number: 61, Value: "masome a thataro le masome le nngwe"},
		{Number: 60, Value: "masome a thataro le masome"},
		{Number: 59, Value: "masome a tlhano le robongwe"},
		{Number: 58, Value: "masome a tlhano le robedi"},
		{Number: 57, Value: "masome a tlhano le supa"},
		{Number: 56, Value: "masome a tlhano le thataro"},
		{Number: 55, Value: "masome a tlhano le tlhano"},
		{Number: 54, Value: "masome a tlhano le ne"},
		{Number: 53, Value: "masome a tlhano le tharo"},
		{Number: 52, Value: "masome a tlhano le pedi"},
		{Number: 51, Value: "masome a tlhano le nngwe"},
		{Number: 50, Value: "masome a tlhano"},
		{Number: 49, Value: "masome a ne le robongwe"},
		{Number: 48, Value: "masome a ne le robedi"},
		{Number: 47, Value: "masome a ne le supa"},
		{Number: 46, Value: "masome a ne le thataro"},
		{Number: 45, Value: "masome a ne le tlhano"},
		{Number: 44, Value: "masome a ne le ne"},
		{Number: 43, Value: "masome a ne le tharo"},
		{Number: 42, Value: "masome a ne le pedi"},
		{Number: 41, Value: "masome a ne le nngwe"},
		{Number: 40, Value: "masome a ne"},
		{Number: 39, Value: "masome a tharo le robongwe"},
		{Number: 38, Value: "masome a tharo le robedi"},
		{Number: 37, Value: "masome a tharo le supa"},
		{Number: 36, Value: "masome a tharo le thataro"},
		{Number: 35, Value: "masome a tharo le tlhano"},
		{Number: 34, Value: "masome a tharo le ne"},
		{Number: 33, Value: "masome a tharo le tharo"},
		{Number: 32, Value: "masome a tharo le pedi"},
		{Number: 31, Value: "masome a tharo le nngwe"},
		{Number: 30, Value: "masome a tharo"},
		{Number: 29, Value: "masome a mabedi le robongwe"},
		{Number: 28, Value: "masome a mabedi le robedi"},
		{Number: 27, Value: "masome a mabedi le supa"},
		{Number: 26, Value: "masome a mabedi le thataro"},
		{Number: 25, Value: "masome a mabedi le tlhano"},
		{Number: 24, Value: "masome a mabedi le ne"},
		{Number: 23, Value: "masome a mabedi le tharo"},
		{Number: 22, Value: "masome a mabedi le pedi"},
		{Number: 21, Value: "masome a mabedi le nngwe"},
		{Number: 20, Value: "masome a mabedi"},
		{Number: 19, Value: "lesome le robongwe"},
		{Number: 18, Value: "lesome le robedi"},
		{Number: 17, Value: "lesome le supa"},
		{Number: 16, Value: "lesome le thataro"},
		{Number: 15, Value: "lesome le tlhano"},
		{Number: 14, Value: "lesome le ne"},
		{Number: 13, Value: "lesome le tharo"},
		{Number: 12, Value: "lesome le pedi"},
		{Number: 11, Value: "lesome le nngwe"},
		{Number: 10, Value: "lesome"},
		{Number: 9, Value: "robongwe"},
		{Number: 8, Value: "robedi"},
		{Number: 7, Value: "supa"},
		{Number: 6, Value: "thataro"},
		{Number: 5, Value: "tlhano"},
		{Number: 4, Value: "ne"},
		{Number: 3, Value: "tharo"},
		{Number: 2, Value: "pedi"},
		{Number: 1, Value: "nngwe"},
		{Number: 0, Value: "lefela"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Lekgolo"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "wa ntlha", Suffix: ".", Masculine: "wa ntlha", Feminine: "wa ntlha", Neuter: "wa ntlha"},
		{Number: 2, Word: "wa bobedi", Suffix: ".", Masculine: "wa bobedi", Feminine: "wa bobedi", Neuter: "wa bobedi"},
		{Number: 3, Word: "wa boraro", Suffix: ".", Masculine: "wa boraro", Feminine: "wa boraro", Neuter: "wa boraro"},
		{Number: 4, Word: "wa bone", Suffix: ".", Masculine: "wa bone", Feminine: "wa bone", Neuter: "wa bone"},
		{Number: 5, Word: "wa botlhano", Suffix: ".", Masculine: "wa botlhano", Feminine: "wa botlhano", Neuter: "wa botlhano"},
		{Number: 6, Word: "wa borataro", Suffix: ".", Masculine: "wa borataro", Feminine: "wa borataro", Neuter: "wa borataro"},
		{Number: 7, Word: "wa bosupa", Suffix: ".", Masculine: "wa bosupa", Feminine: "wa bosupa", Neuter: "wa bosupa"},
		{Number: 8, Word: "wa borobedi", Suffix: ".", Masculine: "wa borobedi", Feminine: "wa borobedi", Neuter: "wa borobedi"},
		{Number: 9, Word: "wa borobongwe", Suffix: ".", Masculine: "wa borobongwe", Feminine: "wa borobongwe", Neuter: "wa borobongwe"},
		{Number: 10, Word: "wa bolesome", Suffix: ".", Masculine: "wa bolesome", Feminine: "wa bolesome", Neuter: "wa bolesome"},
		{Number: 11, Word: "wa lesome le nngwe", Suffix: ".", Masculine: "wa lesome le nngwe", Feminine: "wa lesome le nngwe", Neuter: "wa lesome le nngwe"},
		{Number: 12, Word: "wa lesome le pedi", Suffix: ".", Masculine: "wa lesome le pedi", Feminine: "wa lesome le pedi", Neuter: "wa lesome le pedi"},
		{Number: 20, Word: "wa masome a mabedi", Suffix: ".", Masculine: "wa masome a mabedi", Feminine: "wa masome a mabedi", Neuter: "wa masome a mabedi"},
		{Number: 21, Word: "wa masome a mabedi le nngwe", Suffix: ".", Masculine: "wa masome a mabedi le nngwe", Feminine: "wa masome a mabedi le nngwe", Neuter: "wa masome a mabedi le nngwe"},
		{Number: 30, Word: "wa masome a tharo", Suffix: ".", Masculine: "wa masome a tharo", Feminine: "wa masome a tharo", Neuter: "wa masome a tharo"},
		{Number: 50, Word: "wa masome a tlhano", Suffix: ".", Masculine: "wa masome a tlhano", Feminine: "wa masome a tlhano", Neuter: "wa masome a tlhano"},
		{Number: 100, Word: "wa lekgolo", Suffix: ".", Masculine: "wa lekgolo", Feminine: "wa lekgolo", Neuter: "wa lekgolo"},
		{Number: 1000, Word: "wa sekete", Suffix: ".", Masculine: "wa sekete", Feminine: "wa sekete", Neuter: "wa sekete"},
	},
	LocaleFormatter: &TswanaFormatter{},
}

// TswanaFormatter handles Tswana (Botswana) formatting
type TswanaFormatter struct{}

func (f *TswanaFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *TswanaFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *TswanaFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *TswanaFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *TswanaFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *TswanaFormatter) ChopDecimal(value decimal.Decimal, decimalPlaces int) decimal.Decimal {
	return value.Truncate(int32(decimalPlaces))
}


func (f *TswanaFormatter) FormatDecimalNumber(amount float64) string {
	return DefaultFormatDecimalNumber(amount, ",", ".")
}
func (f *TswanaFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
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
