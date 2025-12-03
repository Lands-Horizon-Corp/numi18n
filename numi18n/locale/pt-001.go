package locale

import (
	"github.com/shopspring/decimal"
)

// PT001Locale is a NumI18NLocale configured for Portuguese (World) - pt-001
var PT001Locale = NumI18NLocale{
	Currency: Currency{
		Name:     "Euro",
		Plural:   "Euros",
		Singular: "Euro",
		Symbol:   "€",
		FractionUnit: FractionUnit{
			Name:     "Cêntimo",
			Plural:   "Cêntimos",
			Singular: "Cêntimo",
			Symbol:   "¢",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "World",
		Currency:       "USD",
		ISO3166Alpha2:  "001",
		ISO3166Alpha3:  "WLD",
		ISO3166Numeric: "001",
		Locale:         "pt-001",
		Timezone:       []string{"UTC"},
		Language:       "pt",
		Emoji:          "🌍",
	},
	Texts: Texts{
		And:   "E",
		Minus: "Menos",
		Only:  "Apenas",
		Point: "Vírgula",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Quadrilhão"},
		{Number: 1000000000000, Value: "Bilião"},
		{Number: 1000000000, Value: "Mil milhões"},
		{Number: 1000000, Value: "Milhão"},
		{Number: 1000, Value: "Mil"},
		{Number: 100, Value: "Cento"},
		{Number: 90, Value: "Noventa"},
		{Number: 80, Value: "Oitenta"},
		{Number: 70, Value: "Setenta"},
		{Number: 60, Value: "Sessenta"},
		{Number: 50, Value: "Cinquenta"},
		{Number: 40, Value: "Quarenta"},
		{Number: 30, Value: "Trinta"},
		{Number: 20, Value: "Vinte"},
		{Number: 19, Value: "Dezanove"},
		{Number: 18, Value: "Dezoito"},
		{Number: 17, Value: "Dezessete"},
		{Number: 16, Value: "Dezasseis"},
		{Number: 15, Value: "Quinze"},
		{Number: 14, Value: "Catorze"},
		{Number: 13, Value: "Treze"},
		{Number: 12, Value: "Doze"},
		{Number: 11, Value: "Onze"},
		{Number: 10, Value: "Dez"},
		{Number: 9, Value: "Nove"},
		{Number: 8, Value: "Oito"},
		{Number: 7, Value: "Sete"},
		{Number: 6, Value: "Seis"},
		{Number: 5, Value: "Cinco"},
		{Number: 4, Value: "Quatro"},
		{Number: 3, Value: "Três"},
		{Number: 2, Value: "Dois"},
		{Number: 1, Value: "Um"},
		{Number: 0, Value: "Zero"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Cento"},
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
	LocaleFormatter: &PortugueseFormatter{},
}

// PortugueseFormatter handles Portuguese formatting
type PortugueseFormatter struct{}

func (f *PortugueseFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *PortugueseFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *PortugueseFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *PortugueseFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *PortugueseFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *PortugueseFormatter) ChopDecimal(value decimal.Decimal, precision int) decimal.Decimal {
	return value.Truncate(int32(precision))
}

func (f *PortugueseFormatter) FormatDecimalNumber(amount float64) string {
	return FormatPolishDecimal(amount)
}
func (f *PortugueseFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Polish conventions (comma separators, period decimal, prefix symbol)
	return FormatPolishCurrency(amount, currencySymbol)
}
