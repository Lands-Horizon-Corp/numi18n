package locale

import (
	"github.com/shopspring/decimal"
)

// RMCHLocale represents the Romansh (Switzerland) locale
var RMCHLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Franc svizzer",
		Plural:   "Francs svizzers",
		Singular: "Franc svizzer",
		Symbol:   "CHF",
		FractionUnit: FractionUnit{
			Name:     "Centesim",
			Plural:   "Centesims",
			Singular: "Centesim",
			Symbol:   "c",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Switzerland",
		Currency:       "CHF",
		ISO3166Alpha2:  "CH",
		ISO3166Alpha3:  "CHE",
		ISO3166Numeric: "756",
		Locale:         "rm-CH",
		Timezone:       []string{"Europe/Zurich"},
		Language:       "rm",
	},
	Texts: Texts{
		And:   "e",
		Minus: "minus",
		Only:  "mo",
		Point: "comma",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "in biliun milliuns"},
		{Number: 1000000000000, Value: "in biliun"},
		{Number: 1000000000, Value: "in milliard"},
		{Number: 1000000, Value: "in milliun"},
		{Number: 100000, Value: "tschient milli"},
		{Number: 90000, Value: "novanta milli"},
		{Number: 80000, Value: "otganta milli"},
		{Number: 70000, Value: "settanta milli"},
		{Number: 60000, Value: "sessanta milli"},
		{Number: 50000, Value: "tschuncanta milli"},
		{Number: 40000, Value: "quaranta milli"},
		{Number: 30000, Value: "trenta milli"},
		{Number: 20000, Value: "ventg milli"},
		{Number: 19000, Value: "dischnov milli"},
		{Number: 18000, Value: "dischot milli"},
		{Number: 17000, Value: "dischset milli"},
		{Number: 16000, Value: "sedisch milli"},
		{Number: 15000, Value: "quindisch milli"},
		{Number: 14000, Value: "quattordisch milli"},
		{Number: 13000, Value: "tredesch milli"},
		{Number: 12000, Value: "dudesch milli"},
		{Number: 11000, Value: "indesch milli"},
		{Number: 10000, Value: "desch milli"},
		{Number: 9000, Value: "nov milli"},
		{Number: 8000, Value: "otg milli"},
		{Number: 7000, Value: "set milli"},
		{Number: 6000, Value: "sis milli"},
		{Number: 5000, Value: "tschun milli"},
		{Number: 4000, Value: "quatter milli"},
		{Number: 3000, Value: "trais milli"},
		{Number: 2000, Value: "dus milli"},
		{Number: 1000, Value: "milli"},
		{Number: 900, Value: "novtschient"},
		{Number: 800, Value: "otgtschient"},
		{Number: 700, Value: "settschient"},
		{Number: 600, Value: "sistschient"},
		{Number: 500, Value: "tschuntschient"},
		{Number: 400, Value: "quattertschient"},
		{Number: 300, Value: "traistschient"},
		{Number: 200, Value: "dustschient"},
		{Number: 100, Value: "tschient"},
		{Number: 99, Value: "novanta e nov"},
		{Number: 98, Value: "novanta e otg"},
		{Number: 97, Value: "novanta e set"},
		{Number: 96, Value: "novanta e sis"},
		{Number: 95, Value: "novanta e tschun"},
		{Number: 94, Value: "novanta e quatter"},
		{Number: 93, Value: "novanta e trais"},
		{Number: 92, Value: "novanta e dus"},
		{Number: 91, Value: "novanta e in"},
		{Number: 90, Value: "novanta"},
		{Number: 89, Value: "otganta e nov"},
		{Number: 88, Value: "otganta e otg"},
		{Number: 87, Value: "otganta e set"},
		{Number: 86, Value: "otganta e sis"},
		{Number: 85, Value: "otganta e tschun"},
		{Number: 84, Value: "otganta e quatter"},
		{Number: 83, Value: "otganta e trais"},
		{Number: 82, Value: "otganta e dus"},
		{Number: 81, Value: "otganta e in"},
		{Number: 80, Value: "otganta"},
		{Number: 79, Value: "settanta e nov"},
		{Number: 78, Value: "settanta e otg"},
		{Number: 77, Value: "settanta e set"},
		{Number: 76, Value: "settanta e sis"},
		{Number: 75, Value: "settanta e tschun"},
		{Number: 74, Value: "settanta e quatter"},
		{Number: 73, Value: "settanta e trais"},
		{Number: 72, Value: "settanta e dus"},
		{Number: 71, Value: "settanta e in"},
		{Number: 70, Value: "settanta"},
		{Number: 69, Value: "sessanta e nov"},
		{Number: 68, Value: "sessanta e otg"},
		{Number: 67, Value: "sessanta e set"},
		{Number: 66, Value: "sessanta e sis"},
		{Number: 65, Value: "sessanta e tschun"},
		{Number: 64, Value: "sessanta e quatter"},
		{Number: 63, Value: "sessanta e trais"},
		{Number: 62, Value: "sessanta e dus"},
		{Number: 61, Value: "sessanta e in"},
		{Number: 60, Value: "sessanta"},
		{Number: 59, Value: "tschuncanta e nov"},
		{Number: 58, Value: "tschuncanta e otg"},
		{Number: 57, Value: "tschuncanta e set"},
		{Number: 56, Value: "tschuncanta e sis"},
		{Number: 55, Value: "tschuncanta e tschun"},
		{Number: 54, Value: "tschuncanta e quatter"},
		{Number: 53, Value: "tschuncanta e trais"},
		{Number: 52, Value: "tschuncanta e dus"},
		{Number: 51, Value: "tschuncanta e in"},
		{Number: 50, Value: "tschuncanta"},
		{Number: 49, Value: "quaranta e nov"},
		{Number: 48, Value: "quaranta e otg"},
		{Number: 47, Value: "quaranta e set"},
		{Number: 46, Value: "quaranta e sis"},
		{Number: 45, Value: "quaranta e tschun"},
		{Number: 44, Value: "quaranta e quatter"},
		{Number: 43, Value: "quaranta e trais"},
		{Number: 42, Value: "quaranta e dus"},
		{Number: 41, Value: "quaranta e in"},
		{Number: 40, Value: "quaranta"},
		{Number: 39, Value: "trenta e nov"},
		{Number: 38, Value: "trenta e otg"},
		{Number: 37, Value: "trenta e set"},
		{Number: 36, Value: "trenta e sis"},
		{Number: 35, Value: "trenta e tschun"},
		{Number: 34, Value: "trenta e quatter"},
		{Number: 33, Value: "trenta e trais"},
		{Number: 32, Value: "trenta e dus"},
		{Number: 31, Value: "trenta e in"},
		{Number: 30, Value: "trenta"},
		{Number: 29, Value: "ventg e nov"},
		{Number: 28, Value: "ventg e otg"},
		{Number: 27, Value: "ventg e set"},
		{Number: 26, Value: "ventg e sis"},
		{Number: 25, Value: "ventg e tschun"},
		{Number: 24, Value: "ventg e quatter"},
		{Number: 23, Value: "ventg e trais"},
		{Number: 22, Value: "ventg e dus"},
		{Number: 21, Value: "ventg e in"},
		{Number: 20, Value: "ventg"},
		{Number: 19, Value: "dischnov"},
		{Number: 18, Value: "dischot"},
		{Number: 17, Value: "dischset"},
		{Number: 16, Value: "sedisch"},
		{Number: 15, Value: "quindisch"},
		{Number: 14, Value: "quattordisch"},
		{Number: 13, Value: "tredesch"},
		{Number: 12, Value: "dudesch"},
		{Number: 11, Value: "indesch"},
		{Number: 10, Value: "desch"},
		{Number: 9, Value: "nov"},
		{Number: 8, Value: "otg"},
		{Number: 7, Value: "set"},
		{Number: 6, Value: "sis"},
		{Number: 5, Value: "tschun"},
		{Number: 4, Value: "quatter"},
		{Number: 3, Value: "trais"},
		{Number: 2, Value: "dus"},
		{Number: 1, Value: "in"},
		{Number: 0, Value: "zero"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Tschient"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "emprim", Suffix: ".", Masculine: "emprim", Feminine: "emprima", Neuter: "emprim"},
		{Number: 2, Word: "segund", Suffix: ".", Masculine: "segund", Feminine: "segunda", Neuter: "segund"},
		{Number: 3, Word: "terz", Suffix: ".", Masculine: "terz", Feminine: "terza", Neuter: "terz"},
		{Number: 4, Word: "quart", Suffix: ".", Masculine: "quart", Feminine: "quarta", Neuter: "quart"},
		{Number: 5, Word: "tschuncavel", Suffix: ".", Masculine: "tschuncavel", Feminine: "tschuncavla", Neuter: "tschuncavel"},
		{Number: 6, Word: "sisavel", Suffix: ".", Masculine: "sisavel", Feminine: "sisavla", Neuter: "sisavel"},
		{Number: 7, Word: "setavel", Suffix: ".", Masculine: "setavel", Feminine: "setavla", Neuter: "setavel"},
		{Number: 8, Word: "otgavel", Suffix: ".", Masculine: "otgavel", Feminine: "otgavla", Neuter: "otgavel"},
		{Number: 9, Word: "novavel", Suffix: ".", Masculine: "novavel", Feminine: "novavla", Neuter: "novavel"},
		{Number: 10, Word: "deschavel", Suffix: ".", Masculine: "deschavel", Feminine: "deschavla", Neuter: "deschavel"},
		{Number: 11, Word: "indeschavel", Suffix: ".", Masculine: "indeschavel", Feminine: "indeschavla", Neuter: "indeschavel"},
		{Number: 12, Word: "dudeschavel", Suffix: ".", Masculine: "dudeschavel", Feminine: "dudeschavla", Neuter: "dudeschavel"},
		{Number: 20, Word: "ventgavel", Suffix: ".", Masculine: "ventgavel", Feminine: "ventgavla", Neuter: "ventgavel"},
		{Number: 21, Word: "ventg e emprimavel", Suffix: ".", Masculine: "ventg e emprimavel", Feminine: "ventg e emprimavla", Neuter: "ventg e emprimavel"},
		{Number: 30, Word: "trentavel", Suffix: ".", Masculine: "trentavel", Feminine: "trentavla", Neuter: "trentavel"},
		{Number: 50, Word: "tschuncantavel", Suffix: ".", Masculine: "tschuncantavel", Feminine: "tschuncantavla", Neuter: "tschuncantavel"},
		{Number: 100, Word: "tschientavel", Suffix: ".", Masculine: "tschientavel", Feminine: "tschientavla", Neuter: "tschientavel"},
		{Number: 1000, Word: "milliavel", Suffix: ".", Masculine: "milliavel", Feminine: "milliavla", Neuter: "milliavel"},
	},
	LocaleFormatter: &RomanshSwitzerlandFormatter{},
}

// RomanshSwitzerlandFormatter handles Romansh (Switzerland) formatting
type RomanshSwitzerlandFormatter struct{}

func (f *RomanshSwitzerlandFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *RomanshSwitzerlandFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *RomanshSwitzerlandFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *RomanshSwitzerlandFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *RomanshSwitzerlandFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *RomanshSwitzerlandFormatter) ChopDecimal(value decimal.Decimal, precision int) decimal.Decimal {
	return value.Truncate(int32(precision))
}
