package locale

import (
	"strings"
	"github.com/shopspring/decimal"
)

// LVLVLocale represents the Latvian (Latvia) locale
var LVLVLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Euro",
		Plural:   "Eiro",
		Singular: "Eiro",
		Symbol:   "€",
		FractionUnit: FractionUnit{
			Name:     "Cents",
			Plural:   "Centi",
			Singular: "Cents",
			Symbol:   "c",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Latvia",
		Currency:       "EUR",
		ISO3166Alpha2:  "LV",
		ISO3166Alpha3:  "LVA",
		ISO3166Numeric: "428",
		Locale:         "lv-LV",
		Timezone:       []string{"Europe/Riga"},
		Language:       "lv",
	},
	Texts: Texts{
		And:   "un",
		Minus: "mīnus",
		Only:  "tikai",
		Point: "komats",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "kvadriljons"},
		{Number: 1000000000000, Value: "triljons"},
		{Number: 1000000000, Value: "miljards"},
		{Number: 1000000, Value: "miljons"},
		{Number: 100000, Value: "simts tūkstoši"},
		{Number: 90000, Value: "deviņdesmit tūkstoši"},
		{Number: 80000, Value: "astoņdesmit tūkstoši"},
		{Number: 70000, Value: "septiņdesmit tūkstoši"},
		{Number: 60000, Value: "sešdesmit tūkstoši"},
		{Number: 50000, Value: "piecdesmit tūkstoši"},
		{Number: 40000, Value: "četrdesmit tūkstoši"},
		{Number: 30000, Value: "trīsdesmit tūkstoši"},
		{Number: 20000, Value: "divdesmit tūkstoši"},
		{Number: 19000, Value: "deviņpadsmit tūkstoši"},
		{Number: 18000, Value: "astoņpadsmit tūkstoši"},
		{Number: 17000, Value: "septiņpadsmit tūkstoši"},
		{Number: 16000, Value: "sešpadsmit tūkstoši"},
		{Number: 15000, Value: "piecpadsmit tūkstoši"},
		{Number: 14000, Value: "četrpadsmit tūkstoši"},
		{Number: 13000, Value: "trīspadsmit tūkstoši"},
		{Number: 12000, Value: "divpadsmit tūkstoši"},
		{Number: 11000, Value: "vienpadsmit tūkstoši"},
		{Number: 10000, Value: "desmit tūkstoši"},
		{Number: 9000, Value: "deviņi tūkstoši"},
		{Number: 8000, Value: "astoņi tūkstoši"},
		{Number: 7000, Value: "septiņi tūkstoši"},
		{Number: 6000, Value: "seši tūkstoši"},
		{Number: 5000, Value: "pieci tūkstoši"},
		{Number: 4000, Value: "četri tūkstoši"},
		{Number: 3000, Value: "trīs tūkstoši"},
		{Number: 2000, Value: "divi tūkstoši"},
		{Number: 1000, Value: "tūkstotis"},
		{Number: 900, Value: "deviņi simti"},
		{Number: 800, Value: "astoņi simti"},
		{Number: 700, Value: "septiņi simti"},
		{Number: 600, Value: "seši simti"},
		{Number: 500, Value: "pieci simti"},
		{Number: 400, Value: "četri simti"},
		{Number: 300, Value: "trīs simti"},
		{Number: 200, Value: "divi simti"},
		{Number: 100, Value: "simts"},
		{Number: 99, Value: "deviņdesmit deviņi"},
		{Number: 98, Value: "deviņdesmit astoņi"},
		{Number: 97, Value: "deviņdesmit septiņi"},
		{Number: 96, Value: "deviņdesmit seši"},
		{Number: 95, Value: "deviņdesmit pieci"},
		{Number: 94, Value: "deviņdesmit četri"},
		{Number: 93, Value: "deviņdesmit trīs"},
		{Number: 92, Value: "deviņdesmit divi"},
		{Number: 91, Value: "deviņdesmit viens"},
		{Number: 90, Value: "deviņdesmit"},
		{Number: 89, Value: "astoņdesmit deviņi"},
		{Number: 88, Value: "astoņdesmit astoņi"},
		{Number: 87, Value: "astoņdesmit septiņi"},
		{Number: 86, Value: "astoņdesmit seši"},
		{Number: 85, Value: "astoņdesmit pieci"},
		{Number: 84, Value: "astoņdesmit četri"},
		{Number: 83, Value: "astoņdesmit trīs"},
		{Number: 82, Value: "astoņdesmit divi"},
		{Number: 81, Value: "astoņdesmit viens"},
		{Number: 80, Value: "astoņdesmit"},
		{Number: 79, Value: "septiņdesmit deviņi"},
		{Number: 78, Value: "septiņdesmit astoņi"},
		{Number: 77, Value: "septiņdesmit septiņi"},
		{Number: 76, Value: "septiņdesmit seši"},
		{Number: 75, Value: "septiņdesmit pieci"},
		{Number: 74, Value: "septiņdesmit četri"},
		{Number: 73, Value: "septiņdesmit trīs"},
		{Number: 72, Value: "septiņdesmit divi"},
		{Number: 71, Value: "septiņdesmit viens"},
		{Number: 70, Value: "septiņdesmit"},
		{Number: 69, Value: "sešdesmit deviņi"},
		{Number: 68, Value: "sešdesmit astoņi"},
		{Number: 67, Value: "sešdesmit septiņi"},
		{Number: 66, Value: "sešdesmit seši"},
		{Number: 65, Value: "sešdesmit pieci"},
		{Number: 64, Value: "sešdesmit četri"},
		{Number: 63, Value: "sešdesmit trīs"},
		{Number: 62, Value: "sešdesmit divi"},
		{Number: 61, Value: "sešdesmit viens"},
		{Number: 60, Value: "sešdesmit"},
		{Number: 59, Value: "piecdesmit deviņi"},
		{Number: 58, Value: "piecdesmit astoņi"},
		{Number: 57, Value: "piecdesmit septiņi"},
		{Number: 56, Value: "piecdesmit seši"},
		{Number: 55, Value: "piecdesmit pieci"},
		{Number: 54, Value: "piecdesmit četri"},
		{Number: 53, Value: "piecdesmit trīs"},
		{Number: 52, Value: "piecdesmit divi"},
		{Number: 51, Value: "piecdesmit viens"},
		{Number: 50, Value: "piecdesmit"},
		{Number: 49, Value: "četrdesmit deviņi"},
		{Number: 48, Value: "četrdesmit astoņi"},
		{Number: 47, Value: "četrdesmit septiņi"},
		{Number: 46, Value: "četrdesmit seši"},
		{Number: 45, Value: "četrdesmit pieci"},
		{Number: 44, Value: "četrdesmit četri"},
		{Number: 43, Value: "četrdesmit trīs"},
		{Number: 42, Value: "četrdesmit divi"},
		{Number: 41, Value: "četrdesmit viens"},
		{Number: 40, Value: "četrdesmit"},
		{Number: 39, Value: "trīsdesmit deviņi"},
		{Number: 38, Value: "trīsdesmit astoņi"},
		{Number: 37, Value: "trīsdesmit septiņi"},
		{Number: 36, Value: "trīsdesmit seši"},
		{Number: 35, Value: "trīsdesmit pieci"},
		{Number: 34, Value: "trīsdesmit četri"},
		{Number: 33, Value: "trīsdesmit trīs"},
		{Number: 32, Value: "trīsdesmit divi"},
		{Number: 31, Value: "trīsdesmit viens"},
		{Number: 30, Value: "trīsdesmit"},
		{Number: 29, Value: "divdesmit deviņi"},
		{Number: 28, Value: "divdesmit astoņi"},
		{Number: 27, Value: "divdesmit septiņi"},
		{Number: 26, Value: "divdesmit seši"},
		{Number: 25, Value: "divdesmit pieci"},
		{Number: 24, Value: "divdesmit četri"},
		{Number: 23, Value: "divdesmit trīs"},
		{Number: 22, Value: "divdesmit divi"},
		{Number: 21, Value: "divdesmit viens"},
		{Number: 20, Value: "divdesmit"},
		{Number: 19, Value: "deviņpadsmit"},
		{Number: 18, Value: "astoņpadsmit"},
		{Number: 17, Value: "septiņpadsmit"},
		{Number: 16, Value: "sešpadsmit"},
		{Number: 15, Value: "piecpadsmit"},
		{Number: 14, Value: "četrpadsmit"},
		{Number: 13, Value: "trīspadsmit"},
		{Number: 12, Value: "divpadsmit"},
		{Number: 11, Value: "vienpadsmit"},
		{Number: 10, Value: "desmit"},
		{Number: 9, Value: "deviņi"},
		{Number: 8, Value: "astoņi"},
		{Number: 7, Value: "septiņi"},
		{Number: 6, Value: "seši"},
		{Number: 5, Value: "pieci"},
		{Number: 4, Value: "četri"},
		{Number: 3, Value: "trīs"},
		{Number: 2, Value: "divi"},
		{Number: 1, Value: "viens"},
		{Number: 0, Value: "nulle"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Viens simts"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "pirmais", Suffix: "-ais", Masculine: "pirmais", Feminine: "pirmā", Neuter: "pirmais"},
		{Number: 2, Word: "otrais", Suffix: "-ais", Masculine: "otrais", Feminine: "otrā", Neuter: "otrais"},
		{Number: 3, Word: "trešais", Suffix: "-ais", Masculine: "trešais", Feminine: "trešā", Neuter: "trešais"},
		{Number: 4, Word: "ceturtais", Suffix: "-ais", Masculine: "ceturtais", Feminine: "ceturtā", Neuter: "ceturtais"},
		{Number: 5, Word: "piektais", Suffix: "-ais", Masculine: "piektais", Feminine: "piektā", Neuter: "piektais"},
		{Number: 6, Word: "sestais", Suffix: "-ais", Masculine: "sestais", Feminine: "sestā", Neuter: "sestais"},
		{Number: 7, Word: "septītais", Suffix: "-ais", Masculine: "septītais", Feminine: "septītā", Neuter: "septītais"},
		{Number: 8, Word: "astotais", Suffix: "-ais", Masculine: "astotais", Feminine: "astotā", Neuter: "astotais"},
		{Number: 9, Word: "devītais", Suffix: "-ais", Masculine: "devītais", Feminine: "devītā", Neuter: "devītais"},
		{Number: 10, Word: "desmitais", Suffix: "-ais", Masculine: "desmitais", Feminine: "desmitā", Neuter: "desmitais"},
		{Number: 11, Word: "vienpadsmitais", Suffix: "-ais", Masculine: "vienpadsmitais", Feminine: "vienpadsmitā", Neuter: "vienpadsmitais"},
		{Number: 12, Word: "divpadsmitais", Suffix: "-ais", Masculine: "divpadsmitais", Feminine: "divpadsmitā", Neuter: "divpadsmitais"},
		{Number: 20, Word: "divdesmitais", Suffix: "-ais", Masculine: "divdesmitais", Feminine: "divdesmitā", Neuter: "divdesmitais"},
		{Number: 21, Word: "divdesmit pirmais", Suffix: "-ais", Masculine: "divdesmit pirmais", Feminine: "divdesmit pirmā", Neuter: "divdesmit pirmais"},
		{Number: 30, Word: "trīsdesmitais", Suffix: "-ais", Masculine: "trīsdesmitais", Feminine: "trīsdesmitā", Neuter: "trīsdesmitais"},
		{Number: 100, Word: "simtais", Suffix: "-ais", Masculine: "simtais", Feminine: "simtā", Neuter: "simtais"},
		{Number: 1000, Word: "tūkstošais", Suffix: "-ais", Masculine: "tūkstošais", Feminine: "tūkstošā", Neuter: "tūkstošais"},
	},
	LocaleFormatter: &LatvianFormatter{},
}

// LatvianFormatter handles Latvian-specific formatting
type LatvianFormatter struct{}

func (f *LatvianFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *LatvianFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *LatvianFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *LatvianFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *LatvianFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *LatvianFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 0
	}
	return amount.Truncate(int32(precision))
}


func (f *LatvianFormatter) FormatDecimalNumber(amount float64) string {
	return DefaultFormatDecimalNumber(amount, ",", ".")
}
func (f *LatvianFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
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
