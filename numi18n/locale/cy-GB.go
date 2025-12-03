package locale

import (
	"github.com/shopspring/decimal"
)

// CYGBLocale is a NumI18NLocale configured for Welsh (United Kingdom) - cy-GB
var CYGBLocale = NumI18NLocale{
	LocaleFormatter: &WelshFormatter{},
	Currency: Currency{
		Name:     "Punt",
		Plural:   "Puntiau",
		Singular: "Punt",
		Symbol:   "£",
		FractionUnit: FractionUnit{
			Name:     "Ceiniog",
			Plural:   "Ceiniogau",
			Singular: "Ceiniog",
			Symbol:   "p",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "United Kingdom",
		Currency:       "GBP",
		ISO3166Alpha2:  "GB",
		ISO3166Alpha3:  "GBR",
		ISO3166Numeric: "826",
		Locale:         "cy-GB",
		Timezone:       []string{"Europe/London"},
		Language:       "cy",
		Emoji:          "🇬🇧",
	},
	Texts: Texts{
		And:   "a",
		Minus: "minus",
		Only:  "yn unig",
		Point: "pwynt",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Quadriliwn"},
		{Number: 1000000000000, Value: "Triliwn"},
		{Number: 1000000000, Value: "Biliwn"},
		{Number: 1000000, Value: "Miliwn"},
		{Number: 1000, Value: "Mil"},
		{Number: 100, Value: "Can"},
		{Number: 90, Value: "Nawdeg"},
		{Number: 80, Value: "Wythdeg"},
		{Number: 70, Value: "Saithdeg"},
		{Number: 60, Value: "Chwedeg"},
		{Number: 50, Value: "Hanner cant"},
		{Number: 40, Value: "Pedwar deg"},
		{Number: 30, Value: "Deg ar hugain"},
		{Number: 20, Value: "Ugain"},
		{Number: 19, Value: "Pedwar ar bymtheg"},
		{Number: 18, Value: "Deunaw"},
		{Number: 17, Value: "Saith ar bymtheg"},
		{Number: 16, Value: "Un ar bymtheg"},
		{Number: 15, Value: "Pymtheg"},
		{Number: 14, Value: "Pedwar ar ddeg"},
		{Number: 13, Value: "Tri ar ddeg"},
		{Number: 12, Value: "Deuddeg"},
		{Number: 11, Value: "Un ar ddeg"},
		{Number: 10, Value: "Deg"},
		{Number: 9, Value: "Naw"},
		{Number: 8, Value: "Wyth"},
		{Number: 7, Value: "Saith"},
		{Number: 6, Value: "Chwech"},
		{Number: 5, Value: "Pum"},
		{Number: 4, Value: "Pedwar"},
		{Number: 3, Value: "Tri"},
		{Number: 2, Value: "Dau"},
		{Number: 1, Value: "Un"},
		{Number: 0, Value: "Sero"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Can"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "cyntaf", Suffix: "af", Masculine: "cyntaf", Feminine: "gyntaf", Neuter: ""},
		{Number: 2, Word: "ail", Suffix: "ail", Masculine: "ail", Feminine: "ail", Neuter: ""},
		{Number: 3, Word: "trydydd", Suffix: "ydd", Masculine: "trydydd", Feminine: "drydedd", Neuter: ""},
		{Number: 4, Word: "pedwerydd", Suffix: "ydd", Masculine: "pedwerydd", Feminine: "bedwaredd", Neuter: ""},
		{Number: 5, Word: "pumed", Suffix: "ed", Masculine: "pumed", Feminine: "bumed", Neuter: ""},
		{Number: 6, Word: "chweched", Suffix: "ed", Masculine: "chweched", Feminine: "chweched", Neuter: ""},
		{Number: 7, Word: "seithfed", Suffix: "fed", Masculine: "seithfed", Feminine: "seithfed", Neuter: ""},
		{Number: 8, Word: "wythfed", Suffix: "fed", Masculine: "wythfed", Feminine: "wythfed", Neuter: ""},
		{Number: 9, Word: "nawfed", Suffix: "fed", Masculine: "nawfed", Feminine: "nawfed", Neuter: ""},
		{Number: 10, Word: "degfed", Suffix: "fed", Masculine: "degfed", Feminine: "degfed", Neuter: ""},
		{Number: 11, Word: "unfed ar ddeg", Suffix: "fed", Masculine: "unfed ar ddeg", Feminine: "unfed ar ddeg", Neuter: ""},
		{Number: 12, Word: "deuddegfed", Suffix: "fed", Masculine: "deuddegfed", Feminine: "deuddegfed", Neuter: ""},
		{Number: 13, Word: "trydydd ar ddeg", Suffix: "ydd", Masculine: "trydydd ar ddeg", Feminine: "drydedd ar ddeg", Neuter: ""},
		{Number: 14, Word: "pedwerydd ar ddeg", Suffix: "ydd", Masculine: "pedwerydd ar ddeg", Feminine: "bedwaredd ar ddeg", Neuter: ""},
		{Number: 15, Word: "pymthegfed", Suffix: "fed", Masculine: "pymthegfed", Feminine: "pymthegfed", Neuter: ""},
		{Number: 16, Word: "unfed ar bymtheg", Suffix: "fed", Masculine: "unfed ar bymtheg", Feminine: "unfed ar bymtheg", Neuter: ""},
		{Number: 17, Word: "seithfed ar bymtheg", Suffix: "fed", Masculine: "seithfed ar bymtheg", Feminine: "seithfed ar bymtheg", Neuter: ""},
		{Number: 18, Word: "deunawfed", Suffix: "fed", Masculine: "deunawfed", Feminine: "deunawfed", Neuter: ""},
		{Number: 19, Word: "pedwerydd ar bymtheg", Suffix: "ydd", Masculine: "pedwerydd ar bymtheg", Feminine: "bedwaredd ar bymtheg", Neuter: ""},
		{Number: 20, Word: "ugainfed", Suffix: "fed", Masculine: "ugainfed", Feminine: "ugainfed", Neuter: ""},
		{Number: 21, Word: "unfed ar hugain", Suffix: "fed", Masculine: "unfed ar hugain", Feminine: "unfed ar hugain", Neuter: ""},
		{Number: 30, Word: "degfed ar hugain", Suffix: "fed", Masculine: "degfed ar hugain", Feminine: "degfed ar hugain", Neuter: ""},
		{Number: 40, Word: "degfed ar drideg", Suffix: "fed", Masculine: "degfed ar drideg", Feminine: "degfed ar drideg", Neuter: ""},
		{Number: 50, Word: "hanner canfed", Suffix: "fed", Masculine: "hanner canfed", Feminine: "hanner canfed", Neuter: ""},
		{Number: 60, Word: "chweched deg", Suffix: "ed", Masculine: "chweched deg", Feminine: "chweched deg", Neuter: ""},
		{Number: 70, Word: "seithfed deg", Suffix: "fed", Masculine: "seithfed deg", Feminine: "seithfed deg", Neuter: ""},
		{Number: 80, Word: "wythfed deg", Suffix: "fed", Masculine: "wythfed deg", Feminine: "wythfed deg", Neuter: ""},
		{Number: 90, Word: "nawfed deg", Suffix: "fed", Masculine: "nawfed deg", Feminine: "nawfed deg", Neuter: ""},
		{Number: 100, Word: "canfed", Suffix: "fed", Masculine: "canfed", Feminine: "canfed", Neuter: ""},
		{Number: 1000, Word: "milfed", Suffix: "fed", Masculine: "milfed", Feminine: "milfed", Neuter: ""},
	},
}

// WelshFormatter handles Welsh formatting
type WelshFormatter struct{}

func (f *WelshFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *WelshFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *WelshFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *WelshFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *WelshFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *WelshFormatter) ChopDecimal(value decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 2
	}
	return value.Truncate(int32(precision))
}

func (f *WelshFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *WelshFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
