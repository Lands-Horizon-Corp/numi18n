package locale

import (
	"github.com/shopspring/decimal"
)

// ESGLLocale is a NumI18NLocale configured for Spain (gl-ES)
var ESGLLocale = NumI18NLocale{
	LocaleFormatter: &GalicianFormatter{},
	Currency: Currency{
		Name:     "Euro",
		Plural:   "Euros",
		Singular: "Euro",
		Symbol:   "€",
		FractionUnit: FractionUnit{
			Name:     "Céntimo",
			Plural:   "Céntimos",
			Singular: "Céntimo",
			Symbol:   "¢",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Spain",
		Currency:       "EUR",
		ISO3166Alpha2:  "ES",
		ISO3166Alpha3:  "ESP",
		ISO3166Numeric: "724",
		Locale:         "gl-ES",
		Timezone:       []string{"Europe/Madrid"},
		Language:       "gl",
	},
	Texts: Texts{
		And:   "E",
		Minus: "Menos",
		Only:  "Só",
		Point: "Punto",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Un billón de billóns"},
		{Number: 1000000000000, Value: "Un billón"},
		{Number: 1000000000, Value: "Mil millóns"},
		{Number: 1000000, Value: "Un millón"},
		{Number: 1000, Value: "Mil"},
		{Number: 100, Value: "Cen"},
		{Number: 90, Value: "Noventa"},
		{Number: 80, Value: "Oitenta"},
		{Number: 70, Value: "Setenta"},
		{Number: 60, Value: "Sesenta"},
		{Number: 50, Value: "Cincuenta"},
		{Number: 40, Value: "Corenta"},
		{Number: 30, Value: "Trinta"},
		{Number: 20, Value: "Vinte"},
		{Number: 19, Value: "Dezanove"},
		{Number: 18, Value: "Dezaoito"},
		{Number: 17, Value: "Dezasete"},
		{Number: 16, Value: "Dezaseis"},
		{Number: 15, Value: "Quince"},
		{Number: 14, Value: "Catorce"},
		{Number: 13, Value: "Trece"},
		{Number: 12, Value: "Doce"},
		{Number: 11, Value: "Once"},
		{Number: 10, Value: "Dez"},
		{Number: 9, Value: "Nove"},
		{Number: 8, Value: "Oito"},
		{Number: 7, Value: "Sete"},
		{Number: 6, Value: "Seis"},
		{Number: 5, Value: "Cinco"},
		{Number: 4, Value: "Catro"},
		{Number: 3, Value: "Tres"},
		{Number: 2, Value: "Dous"},
		{Number: 1, Value: "Un"},
		{Number: 0, Value: "Zero"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Cen"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "Primeiro", Suffix: "º", Masculine: "Primeiro", Feminine: "Primeira", Neuter: "Primeiro"},
		{Number: 2, Word: "Segundo", Suffix: "º", Masculine: "Segundo", Feminine: "Segunda", Neuter: "Segundo"},
		{Number: 3, Word: "Terceiro", Suffix: "º", Masculine: "Terceiro", Feminine: "Terceira", Neuter: "Terceiro"},
		{Number: 4, Word: "Cuarto", Suffix: "º", Masculine: "Cuarto", Feminine: "Cuarta", Neuter: "Cuarto"},
		{Number: 5, Word: "Quinto", Suffix: "º", Masculine: "Quinto", Feminine: "Quinta", Neuter: "Quinto"},
		{Number: 6, Word: "Sexto", Suffix: "º", Masculine: "Sexto", Feminine: "Sexta", Neuter: "Sexto"},
		{Number: 7, Word: "Sétimo", Suffix: "º", Masculine: "Sétimo", Feminine: "Sétima", Neuter: "Sétimo"},
		{Number: 8, Word: "Oitavo", Suffix: "º", Masculine: "Oitavo", Feminine: "Oitava", Neuter: "Oitavo"},
		{Number: 9, Word: "Noveno", Suffix: "º", Masculine: "Noveno", Feminine: "Novena", Neuter: "Noveno"},
		{Number: 10, Word: "Décimo", Suffix: "º", Masculine: "Décimo", Feminine: "Décima", Neuter: "Décimo"},
		{Number: 11, Word: "Undécimo", Suffix: "º", Masculine: "Undécimo", Feminine: "Undécima", Neuter: "Undécimo"},
		{Number: 12, Word: "Duodécimo", Suffix: "º", Masculine: "Duodécimo", Feminine: "Duodécima", Neuter: "Duodécimo"},
		{Number: 13, Word: "Décimo terceiro", Suffix: "º", Masculine: "Décimo terceiro", Feminine: "Décima terceira", Neuter: "Décimo terceiro"},
		{Number: 14, Word: "Décimo cuarto", Suffix: "º", Masculine: "Décimo cuarto", Feminine: "Décima cuarta", Neuter: "Décimo cuarto"},
		{Number: 15, Word: "Décimo quinto", Suffix: "º", Masculine: "Décimo quinto", Feminine: "Décima quinta", Neuter: "Décimo quinto"},
		{Number: 16, Word: "Décimo sexto", Suffix: "º", Masculine: "Décimo sexto", Feminine: "Décima sexta", Neuter: "Décimo sexto"},
		{Number: 17, Word: "Décimo sétimo", Suffix: "º", Masculine: "Décimo sétimo", Feminine: "Décima sétima", Neuter: "Décimo sétimo"},
		{Number: 18, Word: "Décimo oitavo", Suffix: "º", Masculine: "Décimo oitavo", Feminine: "Décima oitava", Neuter: "Décimo oitavo"},
		{Number: 19, Word: "Décimo noveno", Suffix: "º", Masculine: "Décimo noveno", Feminine: "Décima novena", Neuter: "Décimo noveno"},
		{Number: 20, Word: "Vixésimo", Suffix: "º", Masculine: "Vixésimo", Feminine: "Vixésima", Neuter: "Vixésimo"},
		{Number: 21, Word: "Vixésimo primeiro", Suffix: "º", Masculine: "Vixésimo primeiro", Feminine: "Vixésima primeira", Neuter: "Vixésimo primeiro"},
		{Number: 30, Word: "Trixésimo", Suffix: "º", Masculine: "Trixésimo", Feminine: "Trixésima", Neuter: "Trixésimo"},
		{Number: 40, Word: "Cuadraxésimo", Suffix: "º", Masculine: "Cuadraxésimo", Feminine: "Cuadraxésima", Neuter: "Cuadraxésimo"},
		{Number: 50, Word: "Quincuaxésimo", Suffix: "º", Masculine: "Quincuaxésimo", Feminine: "Quincuaxésima", Neuter: "Quincuaxésimo"},
		{Number: 60, Word: "Sexaxésimo", Suffix: "º", Masculine: "Sexaxésimo", Feminine: "Sexaxésima", Neuter: "Sexaxésimo"},
		{Number: 70, Word: "Septuaxésimo", Suffix: "º", Masculine: "Septuaxésimo", Feminine: "Septuaxésima", Neuter: "Septuaxésimo"},
		{Number: 80, Word: "Octoxésimo", Suffix: "º", Masculine: "Octoxésimo", Feminine: "Octoxésima", Neuter: "Octoxésimo"},
		{Number: 90, Word: "Nonaxésimo", Suffix: "º", Masculine: "Nonaxésimo", Feminine: "Nonaxésima", Neuter: "Nonaxésimo"},
		{Number: 100, Word: "Centésimo", Suffix: "º", Masculine: "Centésimo", Feminine: "Centésima", Neuter: "Centésimo"},
		{Number: 1000, Word: "Milésimo", Suffix: "º", Masculine: "Milésimo", Feminine: "Milésima", Neuter: "Milésimo"},
	},
}

// GalicianFormatter handles Galician (gl-ES) formatting
type GalicianFormatter struct{}

func (f *GalicianFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *GalicianFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *GalicianFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *GalicianFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *GalicianFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *GalicianFormatter) ChopDecimal(d decimal.Decimal, precision int) decimal.Decimal {
	return d.Truncate(int32(precision))
}

func (f *GalicianFormatter) FormatDecimalNumber(amount float64) string {
	return FormatEuropeanDecimal(amount)
}
func (f *GalicianFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with European conventions (period separators, comma decimal, prefix symbol)
	return FormatEuropeanCurrency(amount, currencySymbol)
}
