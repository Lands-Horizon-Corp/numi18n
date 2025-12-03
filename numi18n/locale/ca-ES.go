package locale

import (
	"github.com/shopspring/decimal"
)

// CAESLocale is a NumI18NLocale configured for Catalan (Spain) - ca-ES
var CAESLocale = NumI18NLocale{
	LocaleFormatter: &CatalanFormatter{},
	Currency: Currency{
		Name:     "Euro",
		Plural:   "Euros",
		Singular: "Euro",
		Symbol:   "€",
		FractionUnit: FractionUnit{
			Name:     "Cèntim",
			Plural:   "Cèntims",
			Singular: "Cèntim",
			Symbol:   "c",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Spain",
		Currency:       "EUR",
		ISO3166Alpha2:  "ES",
		ISO3166Alpha3:  "ESP",
		ISO3166Numeric: "724",
		Locale:         "ca-ES",
		Timezone:       []string{"Europe/Madrid"},
		Language:       "ca",
		Emoji:          "🇪🇸",
	},
	Texts: Texts{
		And:   "i",
		Minus: "Menys",
		Only:  "Només",
		Point: "Punt",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Quadrilió"},
		{Number: 1000000000000, Value: "Trilió"},
		{Number: 1000000000, Value: "Mil milions"},
		{Number: 1000000, Value: "Milió"},
		{Number: 1000, Value: "Mil"},
		{Number: 100, Value: "Cent"},
		{Number: 90, Value: "Noranta"},
		{Number: 80, Value: "Vuitanta"},
		{Number: 70, Value: "Setanta"},
		{Number: 60, Value: "Seixanta"},
		{Number: 50, Value: "Cinquanta"},
		{Number: 40, Value: "Quaranta"},
		{Number: 30, Value: "Trenta"},
		{Number: 20, Value: "Vint"},
		{Number: 19, Value: "Dinou"},
		{Number: 18, Value: "Divuit"},
		{Number: 17, Value: "Disset"},
		{Number: 16, Value: "Setze"},
		{Number: 15, Value: "Quinze"},
		{Number: 14, Value: "Catorze"},
		{Number: 13, Value: "Tretze"},
		{Number: 12, Value: "Dotze"},
		{Number: 11, Value: "Onze"},
		{Number: 10, Value: "Deu"},
		{Number: 9, Value: "Nou"},
		{Number: 8, Value: "Vuit"},
		{Number: 7, Value: "Set"},
		{Number: 6, Value: "Sis"},
		{Number: 5, Value: "Cinc"},
		{Number: 4, Value: "Quatre"},
		{Number: 3, Value: "Tres"},
		{Number: 2, Value: "Dos"},
		{Number: 1, Value: "Un"},
		{Number: 0, Value: "Zero"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Cent"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "primer", Suffix: "r", Masculine: "primer", Feminine: "primera", Neuter: ""},
		{Number: 2, Word: "segon", Suffix: "n", Masculine: "segon", Feminine: "segona", Neuter: ""},
		{Number: 3, Word: "tercer", Suffix: "r", Masculine: "tercer", Feminine: "tercera", Neuter: ""},
		{Number: 4, Word: "quart", Suffix: "t", Masculine: "quart", Feminine: "quarta", Neuter: ""},
		{Number: 5, Word: "cinquè", Suffix: "è", Masculine: "cinquè", Feminine: "cinquena", Neuter: ""},
		{Number: 6, Word: "sisè", Suffix: "è", Masculine: "sisè", Feminine: "sisena", Neuter: ""},
		{Number: 7, Word: "setè", Suffix: "è", Masculine: "setè", Feminine: "setena", Neuter: ""},
		{Number: 8, Word: "vuitè", Suffix: "è", Masculine: "vuitè", Feminine: "vuitena", Neuter: ""},
		{Number: 9, Word: "novè", Suffix: "è", Masculine: "novè", Feminine: "novena", Neuter: ""},
		{Number: 10, Word: "desè", Suffix: "è", Masculine: "desè", Feminine: "desena", Neuter: ""},
		{Number: 11, Word: "onzè", Suffix: "è", Masculine: "onzè", Feminine: "onzena", Neuter: ""},
		{Number: 12, Word: "dotzè", Suffix: "è", Masculine: "dotzè", Feminine: "dotzena", Neuter: ""},
		{Number: 13, Word: "tretzè", Suffix: "è", Masculine: "tretzè", Feminine: "tretzena", Neuter: ""},
		{Number: 14, Word: "catorzè", Suffix: "è", Masculine: "catorzè", Feminine: "catorzena", Neuter: ""},
		{Number: 15, Word: "quinzè", Suffix: "è", Masculine: "quinzè", Feminine: "quinzena", Neuter: ""},
		{Number: 16, Word: "setzè", Suffix: "è", Masculine: "setzè", Feminine: "setzena", Neuter: ""},
		{Number: 17, Word: "dissetè", Suffix: "è", Masculine: "dissetè", Feminine: "dissetena", Neuter: ""},
		{Number: 18, Word: "divuitè", Suffix: "è", Masculine: "divuitè", Feminine: "divuitena", Neuter: ""},
		{Number: 19, Word: "dinovè", Suffix: "è", Masculine: "dinovè", Feminine: "dinovena", Neuter: ""},
		{Number: 20, Word: "vintè", Suffix: "è", Masculine: "vintè", Feminine: "vintena", Neuter: ""},
		{Number: 21, Word: "vint-i-unè", Suffix: "è", Masculine: "vint-i-unè", Feminine: "vint-i-unena", Neuter: ""},
		{Number: 30, Word: "trentè", Suffix: "è", Masculine: "trentè", Feminine: "trentena", Neuter: ""},
		{Number: 40, Word: "quarantè", Suffix: "è", Masculine: "quarantè", Feminine: "quarantena", Neuter: ""},
		{Number: 50, Word: "cinquantè", Suffix: "è", Masculine: "cinquantè", Feminine: "cinquantena", Neuter: ""},
		{Number: 60, Word: "seixantè", Suffix: "è", Masculine: "seixantè", Feminine: "seixantena", Neuter: ""},
		{Number: 70, Word: "setantè", Suffix: "è", Masculine: "setantè", Feminine: "setantena", Neuter: ""},
		{Number: 80, Word: "vuitantè", Suffix: "è", Masculine: "vuitantè", Feminine: "vuitantena", Neuter: ""},
		{Number: 90, Word: "norantè", Suffix: "è", Masculine: "norantè", Feminine: "norantena", Neuter: ""},
		{Number: 100, Word: "centè", Suffix: "è", Masculine: "centè", Feminine: "centena", Neuter: ""},
		{Number: 1000, Word: "mil·lè", Suffix: "è", Masculine: "mil·lè", Feminine: "mil·lena", Neuter: ""},
	},
}

// CatalanFormatter handles Catalan formatting
type CatalanFormatter struct{}

func (f *CatalanFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *CatalanFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *CatalanFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *CatalanFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *CatalanFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *CatalanFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	return amount.Round(int32(precision))
}

func (f *CatalanFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *CatalanFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
