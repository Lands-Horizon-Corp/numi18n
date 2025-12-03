package locale

import (
	"github.com/shopspring/decimal"
)

// ES001Locale is a NumI18NLocale configured for Spanish (World) - es-001
var ES001Locale = NumI18NLocale{
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
		CountryName:    "World",
		Currency:       "USD",
		ISO3166Alpha2:  "001",
		ISO3166Alpha3:  "WLD",
		ISO3166Numeric: "001",
		Locale:         "es-001",
		Timezone:       []string{"UTC"},
		Language:       "es",
		Emoji:          "🌍",
		PhoneCode:      "+1",
		Domain:         ".com",
	},
	Texts: Texts{
		And:   "Y",
		Minus: "Menos",
		Only:  "Sólo",
		Point: "Punto",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Cuatrillón"},
		{Number: 1000000000000, Value: "Billón"},
		{Number: 1000000000, Value: "Mil millones"},
		{Number: 1000000, Value: "Millón"},
		{Number: 1000, Value: "Mil"},
		{Number: 900, Value: "Novecientos"},
		{Number: 800, Value: "Ochocientos"},
		{Number: 700, Value: "Setecientos"},
		{Number: 600, Value: "Seiscientos"},
		{Number: 500, Value: "Quinientos"},
		{Number: 400, Value: "Cuatrocientos"},
		{Number: 300, Value: "Trescientos"},
		{Number: 200, Value: "Doscientos"},
		{Number: 101, Value: "Ciento uno"},
		{Number: 100, Value: "Ciento"},
		{Number: 90, Value: "Noventa"},
		{Number: 80, Value: "Ochenta"},
		{Number: 70, Value: "Setenta"},
		{Number: 60, Value: "Sesenta"},
		{Number: 50, Value: "Cincuenta"},
		{Number: 40, Value: "Cuarenta"},
		{Number: 30, Value: "Treinta"},
		{Number: 29, Value: "Veintinueve"},
		{Number: 28, Value: "Veintiocho"},
		{Number: 27, Value: "Veintisiete"},
		{Number: 26, Value: "Veintiséis"},
		{Number: 25, Value: "Veinticinco"},
		{Number: 24, Value: "Veinticuatro"},
		{Number: 23, Value: "Veintitrés"},
		{Number: 22, Value: "Veintidós"},
		{Number: 21, Value: "Veintiuno"},
		{Number: 20, Value: "Veinte"},
		{Number: 19, Value: "Diecinueve"},
		{Number: 18, Value: "Dieciocho"},
		{Number: 17, Value: "Diecisiete"},
		{Number: 16, Value: "Dieciséis"},
		{Number: 15, Value: "Quince"},
		{Number: 14, Value: "Catorce"},
		{Number: 13, Value: "Trece"},
		{Number: 12, Value: "Doce"},
		{Number: 11, Value: "Once"},
		{Number: 10, Value: "Diez"},
		{Number: 9, Value: "Nueve"},
		{Number: 8, Value: "Ocho"},
		{Number: 7, Value: "Siete"},
		{Number: 6, Value: "Seis"},
		{Number: 5, Value: "Cinco"},
		{Number: 4, Value: "Cuatro"},
		{Number: 3, Value: "Tres"},
		{Number: 2, Value: "Dos"},
		{Number: 1, Value: "Uno"},
		{Number: 0, Value: "Cero"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 1000000, Value: "Un Millón"},
		{Number: 1000, Value: "Mil"},
		{Number: 900, Value: "Novecientos"},
		{Number: 800, Value: "Ochocientos"},
		{Number: 700, Value: "Setecientos"},
		{Number: 600, Value: "Seiscientos"},
		{Number: 500, Value: "Quinientos"},
		{Number: 400, Value: "Cuatrocientos"},
		{Number: 300, Value: "Trescientos"},
		{Number: 200, Value: "Doscientos"},
		{Number: 100, Value: "Cien"},
		{Number: 29, Value: "Veintinueve"},
		{Number: 28, Value: "Veintiocho"},
		{Number: 27, Value: "Veintisiete"},
		{Number: 26, Value: "Veintiséis"},
		{Number: 25, Value: "Veinticinco"},
		{Number: 24, Value: "Veinticuatro"},
		{Number: 23, Value: "Veintitrés"},
		{Number: 22, Value: "Veintidós"},
		{Number: 21, Value: "Veintiuno"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "Primero", Suffix: "°", Masculine: "Primer", Feminine: "Primera"},
		{Number: 2, Word: "Segundo", Suffix: "°", Masculine: "Segundo", Feminine: "Segunda"},
		{Number: 3, Word: "Tercero", Suffix: "°", Masculine: "Tercer", Feminine: "Tercera"},
		{Number: 4, Word: "Cuarto", Suffix: "°", Masculine: "Cuarto", Feminine: "Cuarta"},
		{Number: 5, Word: "Quinto", Suffix: "°", Masculine: "Quinto", Feminine: "Quinta"},
		{Number: 6, Word: "Sexto", Suffix: "°", Masculine: "Sexto", Feminine: "Sexta"},
		{Number: 7, Word: "Séptimo", Suffix: "°", Masculine: "Séptimo", Feminine: "Séptima"},
		{Number: 8, Word: "Octavo", Suffix: "°", Masculine: "Octavo", Feminine: "Octava"},
		{Number: 9, Word: "Noveno", Suffix: "°", Masculine: "Noveno", Feminine: "Novena"},
		{Number: 10, Word: "Décimo", Suffix: "°", Masculine: "Décimo", Feminine: "Décima"},
		{Number: 11, Word: "Undécimo", Suffix: "°", Masculine: "Undécimo", Feminine: "Undécima"},
		{Number: 12, Word: "Duodécimo", Suffix: "°", Masculine: "Duodécimo", Feminine: "Duodécima"},
		{Number: 13, Word: "Decimotercero", Suffix: "°", Masculine: "Decimotercero", Feminine: "Decimotercera"},
		{Number: 14, Word: "Decimocuarto", Suffix: "°", Masculine: "Decimocuarto", Feminine: "Decimocuarta"},
		{Number: 15, Word: "Decimoquinto", Suffix: "°", Masculine: "Decimoquinto", Feminine: "Decimoquinta"},
		{Number: 16, Word: "Decimosexto", Suffix: "°", Masculine: "Decimosexto", Feminine: "Decimosexta"},
		{Number: 17, Word: "Decimoséptimo", Suffix: "°", Masculine: "Decimoséptimo", Feminine: "Decimoséptima"},
		{Number: 18, Word: "Decimoctavo", Suffix: "°", Masculine: "Decimoctavo", Feminine: "Decimoctava"},
		{Number: 19, Word: "Decimonoveno", Suffix: "°", Masculine: "Decimonoveno", Feminine: "Decimonovena"},
		{Number: 20, Word: "Vigésimo", Suffix: "°", Masculine: "Vigésimo", Feminine: "Vigésima"},
		{Number: 21, Word: "Vigésimo primero", Suffix: "°", Masculine: "Vigésimo primero", Feminine: "Vigésima primera"},
		{Number: 22, Word: "Vigésimo segundo", Suffix: "°", Masculine: "Vigésimo segundo", Feminine: "Vigésima segunda"},
		{Number: 23, Word: "Vigésimo tercero", Suffix: "°", Masculine: "Vigésimo tercero", Feminine: "Vigésima tercera"},
		{Number: 30, Word: "Trigésimo", Suffix: "°", Masculine: "Trigésimo", Feminine: "Trigésima"},
		{Number: 40, Word: "Cuadragésimo", Suffix: "°", Masculine: "Cuadragésimo", Feminine: "Cuadragésima"},
		{Number: 50, Word: "Quincuagésimo", Suffix: "°", Masculine: "Quincuagésimo", Feminine: "Quincuagésima"},
		{Number: 60, Word: "Sexagésimo", Suffix: "°", Masculine: "Sexagésimo", Feminine: "Sexagésima"},
		{Number: 70, Word: "Septuagésimo", Suffix: "°", Masculine: "Septuagésimo", Feminine: "Septuagésima"},
		{Number: 80, Word: "Octogésimo", Suffix: "°", Masculine: "Octogésimo", Feminine: "Octogésima"},
		{Number: 90, Word: "Nonagésimo", Suffix: "°", Masculine: "Nonagésimo", Feminine: "Nonagésima"},
		{Number: 100, Word: "Centésimo", Suffix: "°", Masculine: "Centésimo", Feminine: "Centésima"},
		{Number: 1000, Word: "Milésimo", Suffix: "°", Masculine: "Milésimo", Feminine: "Milésima"},
		{Number: 1000000, Word: "Millonésimo", Suffix: "°", Masculine: "Millonésimo", Feminine: "Millonésima"},
	},
	LocaleFormatter: &SpanishFormatter{},
}

// SpanishFormatter handles Spanish formatting with proper compound number rules
type SpanishFormatter struct{}

func (f *SpanishFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	// First check exact mappings for compound and special numbers
	exactResult := ConvertToWordsWithExactMappingInt64(number, targetLocale)
	if exactResult != ConvertToWordsGenericInt64(number, targetLocale) {
		return exactResult // Found exact mapping
	}

	// Use default conversion for most numbers
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *SpanishFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *SpanishFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *SpanishFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *SpanishFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *SpanishFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 0
	}
	return amount.Truncate(int32(precision))
}

func (f *SpanishFormatter) FormatDecimalNumber(amount float64) string {
	return FormatPolishDecimal(amount)
}
func (f *SpanishFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Polish conventions (comma separators, period decimal, prefix symbol)
	return FormatPolishCurrency(amount, currencySymbol)
}
