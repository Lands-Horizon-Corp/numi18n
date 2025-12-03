package locale

import (
	"github.com/shopspring/decimal"
)

// PTAOLocale represents the Portuguese (Angola) locale
var PTAOLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Kwanza",
		Plural:   "Kwanzas",
		Singular: "Kwanza",
		Symbol:   "Kz",
		FractionUnit: FractionUnit{
			Name:     "Cêntimo",
			Plural:   "Cêntimos",
			Singular: "Cêntimo",
			Symbol:   "c",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Angola",
		Currency:       "AOA",
		ISO3166Alpha2:  "AO",
		ISO3166Alpha3:  "AGO",
		ISO3166Numeric: "024",
		Locale:         "pt-AO",
		Timezone:       []string{"Africa/Luanda"},
		Language:       "pt",
		Emoji:          "🇦🇴",
		PhoneCode:      "+244",
		Domain:         ".ao",
	},
	Texts: Texts{
		And:   "e",
		Minus: "menos",
		Only:  "apenas",
		Point: "vírgula",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "mil biliões"},
		{Number: 1000000000000, Value: "bilião"},
		{Number: 1000000000, Value: "mil milhões"},
		{Number: 1000000, Value: "milhão"},
		{Number: 100000, Value: "cem mil"},
		{Number: 90000, Value: "noventa mil"},
		{Number: 80000, Value: "oitenta mil"},
		{Number: 70000, Value: "setenta mil"},
		{Number: 60000, Value: "sessenta mil"},
		{Number: 50000, Value: "cinquenta mil"},
		{Number: 40000, Value: "quarenta mil"},
		{Number: 30000, Value: "trinta mil"},
		{Number: 20000, Value: "vinte mil"},
		{Number: 19000, Value: "dezanove mil"},
		{Number: 18000, Value: "dezoito mil"},
		{Number: 17000, Value: "dezassete mil"},
		{Number: 16000, Value: "dezasseis mil"},
		{Number: 15000, Value: "quinze mil"},
		{Number: 14000, Value: "catorze mil"},
		{Number: 13000, Value: "treze mil"},
		{Number: 12000, Value: "doze mil"},
		{Number: 11000, Value: "onze mil"},
		{Number: 10000, Value: "dez mil"},
		{Number: 9000, Value: "nove mil"},
		{Number: 8000, Value: "oito mil"},
		{Number: 7000, Value: "sete mil"},
		{Number: 6000, Value: "seis mil"},
		{Number: 5000, Value: "cinco mil"},
		{Number: 4000, Value: "quatro mil"},
		{Number: 3000, Value: "três mil"},
		{Number: 2000, Value: "dois mil"},
		{Number: 1000, Value: "mil"},
		{Number: 900, Value: "novecentos"},
		{Number: 800, Value: "oitocentos"},
		{Number: 700, Value: "setecentos"},
		{Number: 600, Value: "seiscentos"},
		{Number: 500, Value: "quinhentos"},
		{Number: 400, Value: "quatrocentos"},
		{Number: 300, Value: "trezentos"},
		{Number: 200, Value: "duzentos"},
		{Number: 100, Value: "cem"},
		{Number: 99, Value: "noventa e nove"},
		{Number: 98, Value: "noventa e oito"},
		{Number: 97, Value: "noventa e sete"},
		{Number: 96, Value: "noventa e seis"},
		{Number: 95, Value: "noventa e cinco"},
		{Number: 94, Value: "noventa e quatro"},
		{Number: 93, Value: "noventa e três"},
		{Number: 92, Value: "noventa e dois"},
		{Number: 91, Value: "noventa e um"},
		{Number: 90, Value: "noventa"},
		{Number: 89, Value: "oitenta e nove"},
		{Number: 88, Value: "oitenta e oito"},
		{Number: 87, Value: "oitenta e sete"},
		{Number: 86, Value: "oitenta e seis"},
		{Number: 85, Value: "oitenta e cinco"},
		{Number: 84, Value: "oitenta e quatro"},
		{Number: 83, Value: "oitenta e três"},
		{Number: 82, Value: "oitenta e dois"},
		{Number: 81, Value: "oitenta e um"},
		{Number: 80, Value: "oitenta"},
		{Number: 79, Value: "setenta e nove"},
		{Number: 78, Value: "setenta e oito"},
		{Number: 77, Value: "setenta e sete"},
		{Number: 76, Value: "setenta e seis"},
		{Number: 75, Value: "setenta e cinco"},
		{Number: 74, Value: "setenta e quatro"},
		{Number: 73, Value: "setenta e três"},
		{Number: 72, Value: "setenta e dois"},
		{Number: 71, Value: "setenta e um"},
		{Number: 70, Value: "setenta"},
		{Number: 69, Value: "sessenta e nove"},
		{Number: 68, Value: "sessenta e oito"},
		{Number: 67, Value: "sessenta e sete"},
		{Number: 66, Value: "sessenta e seis"},
		{Number: 65, Value: "sessenta e cinco"},
		{Number: 64, Value: "sessenta e quatro"},
		{Number: 63, Value: "sessenta e três"},
		{Number: 62, Value: "sessenta e dois"},
		{Number: 61, Value: "sessenta e um"},
		{Number: 60, Value: "sessenta"},
		{Number: 59, Value: "cinquenta e nove"},
		{Number: 58, Value: "cinquenta e oito"},
		{Number: 57, Value: "cinquenta e sete"},
		{Number: 56, Value: "cinquenta e seis"},
		{Number: 55, Value: "cinquenta e cinco"},
		{Number: 54, Value: "cinquenta e quatro"},
		{Number: 53, Value: "cinquenta e três"},
		{Number: 52, Value: "cinquenta e dois"},
		{Number: 51, Value: "cinquenta e um"},
		{Number: 50, Value: "cinquenta"},
		{Number: 49, Value: "quarenta e nove"},
		{Number: 48, Value: "quarenta e oito"},
		{Number: 47, Value: "quarenta e sete"},
		{Number: 46, Value: "quarenta e seis"},
		{Number: 45, Value: "quarenta e cinco"},
		{Number: 44, Value: "quarenta e quatro"},
		{Number: 43, Value: "quarenta e três"},
		{Number: 42, Value: "quarenta e dois"},
		{Number: 41, Value: "quarenta e um"},
		{Number: 40, Value: "quarenta"},
		{Number: 39, Value: "trinta e nove"},
		{Number: 38, Value: "trinta e oito"},
		{Number: 37, Value: "trinta e sete"},
		{Number: 36, Value: "trinta e seis"},
		{Number: 35, Value: "trinta e cinco"},
		{Number: 34, Value: "trinta e quatro"},
		{Number: 33, Value: "trinta e três"},
		{Number: 32, Value: "trinta e dois"},
		{Number: 31, Value: "trinta e um"},
		{Number: 30, Value: "trinta"},
		{Number: 29, Value: "vinte e nove"},
		{Number: 28, Value: "vinte e oito"},
		{Number: 27, Value: "vinte e sete"},
		{Number: 26, Value: "vinte e seis"},
		{Number: 25, Value: "vinte e cinco"},
		{Number: 24, Value: "vinte e quatro"},
		{Number: 23, Value: "vinte e três"},
		{Number: 22, Value: "vinte e dois"},
		{Number: 21, Value: "vinte e um"},
		{Number: 20, Value: "vinte"},
		{Number: 19, Value: "dezanove"},
		{Number: 18, Value: "dezoito"},
		{Number: 17, Value: "dezassete"},
		{Number: 16, Value: "dezasseis"},
		{Number: 15, Value: "quinze"},
		{Number: 14, Value: "catorze"},
		{Number: 13, Value: "treze"},
		{Number: 12, Value: "doze"},
		{Number: 11, Value: "onze"},
		{Number: 10, Value: "dez"},
		{Number: 9, Value: "nove"},
		{Number: 8, Value: "oito"},
		{Number: 7, Value: "sete"},
		{Number: 6, Value: "seis"},
		{Number: 5, Value: "cinco"},
		{Number: 4, Value: "quatro"},
		{Number: 3, Value: "três"},
		{Number: 2, Value: "dois"},
		{Number: 1, Value: "um"},
		{Number: 0, Value: "zero"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Cem"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "primeiro", Suffix: "º", Masculine: "primeiro", Feminine: "primeira", Neuter: "primeiro"},
		{Number: 2, Word: "segundo", Suffix: "º", Masculine: "segundo", Feminine: "segunda", Neuter: "segundo"},
		{Number: 3, Word: "terceiro", Suffix: "º", Masculine: "terceiro", Feminine: "terceira", Neuter: "terceiro"},
		{Number: 4, Word: "quarto", Suffix: "º", Masculine: "quarto", Feminine: "quarta", Neuter: "quarto"},
		{Number: 5, Word: "quinto", Suffix: "º", Masculine: "quinto", Feminine: "quinta", Neuter: "quinto"},
		{Number: 6, Word: "sexto", Suffix: "º", Masculine: "sexto", Feminine: "sexta", Neuter: "sexto"},
		{Number: 7, Word: "sétimo", Suffix: "º", Masculine: "sétimo", Feminine: "sétima", Neuter: "sétimo"},
		{Number: 8, Word: "oitavo", Suffix: "º", Masculine: "oitavo", Feminine: "oitava", Neuter: "oitavo"},
		{Number: 9, Word: "nono", Suffix: "º", Masculine: "nono", Feminine: "nona", Neuter: "nono"},
		{Number: 10, Word: "décimo", Suffix: "º", Masculine: "décimo", Feminine: "décima", Neuter: "décimo"},
		{Number: 11, Word: "décimo primeiro", Suffix: "º", Masculine: "décimo primeiro", Feminine: "décima primeira", Neuter: "décimo primeiro"},
		{Number: 12, Word: "décimo segundo", Suffix: "º", Masculine: "décimo segundo", Feminine: "décima segunda", Neuter: "décimo segundo"},
		{Number: 20, Word: "vigésimo", Suffix: "º", Masculine: "vigésimo", Feminine: "vigésima", Neuter: "vigésimo"},
		{Number: 21, Word: "vigésimo primeiro", Suffix: "º", Masculine: "vigésimo primeiro", Feminine: "vigésima primeira", Neuter: "vigésimo primeiro"},
		{Number: 30, Word: "trigésimo", Suffix: "º", Masculine: "trigésimo", Feminine: "trigésima", Neuter: "trigésimo"},
		{Number: 50, Word: "quinquagésimo", Suffix: "º", Masculine: "quinquagésimo", Feminine: "quinquagésima", Neuter: "quinquagésimo"},
		{Number: 100, Word: "centésimo", Suffix: "º", Masculine: "centésimo", Feminine: "centésima", Neuter: "centésimo"},
		{Number: 1000, Word: "milésimo", Suffix: "º", Masculine: "milésimo", Feminine: "milésima", Neuter: "milésimo"},
	},
	LocaleFormatter: &PortugueseAngolaFormatter{},
}

// PortugueseAngolaFormatter handles Portuguese (Angola) formatting
type PortugueseAngolaFormatter struct{}

func (f *PortugueseAngolaFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *PortugueseAngolaFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *PortugueseAngolaFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *PortugueseAngolaFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *PortugueseAngolaFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *PortugueseAngolaFormatter) ChopDecimal(value decimal.Decimal, precision int) decimal.Decimal {
	return value.Truncate(int32(precision))
}

func (f *PortugueseAngolaFormatter) FormatDecimalNumber(amount float64) string {
	return FormatPolishDecimal(amount)
}
func (f *PortugueseAngolaFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Polish conventions (comma separators, period decimal, prefix symbol)
	return FormatPolishCurrency(amount, currencySymbol)
}
