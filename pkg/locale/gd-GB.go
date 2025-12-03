package locale

import (
	"github.com/shopspring/decimal"
)

// GBGDLocale is a NumI18NLocale configured for Great Britain (gd-GB)
var GBGDLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Punnd Sasannach",
		Plural:   "Puinnd Sasannach",
		Singular: "Punnd Sasannach",
		Symbol:   "£",
		FractionUnit: FractionUnit{
			Name:     "Sgillinn",
			Plural:   "Sgillinnean",
			Singular: "Sgillinn",
			Symbol:   "p",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "United Kingdom",
		Currency:       "GBP",
		ISO3166Alpha2:  "GB",
		ISO3166Alpha3:  "GBR",
		ISO3166Numeric: "826",
		Locale:         "gd-GB",
		Timezone:       []string{"Europe/London"},
		Language:       "gd",
	},
	Texts: Texts{
		And:   "Agus",
		Minus: "As aonais",
		Only:  "A-mhàin",
		Point: "Puing",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Mìle billean billean"},
		{Number: 1000000000000, Value: "Billean billean"},
		{Number: 1000000000, Value: "Billean"},
		{Number: 1000000, Value: "Millean"},
		{Number: 1000, Value: "Mìle"},
		{Number: 100, Value: "Ceud"},
		{Number: 90, Value: "Naochad"},
		{Number: 80, Value: "Ochdad"},
		{Number: 70, Value: "Seachdad"},
		{Number: 60, Value: "Seasgad"},
		{Number: 50, Value: "Leth-cheud"},
		{Number: 40, Value: "Dà fhichead"},
		{Number: 30, Value: "Trithead"},
		{Number: 20, Value: "Fichead"},
		{Number: 19, Value: "Naoi deug"},
		{Number: 18, Value: "Ochd deug"},
		{Number: 17, Value: "Seachd deug"},
		{Number: 16, Value: "Sia deug"},
		{Number: 15, Value: "Còig deug"},
		{Number: 14, Value: "Ceithir deug"},
		{Number: 13, Value: "Trì deug"},
		{Number: 12, Value: "Dà dheug"},
		{Number: 11, Value: "Aon deug"},
		{Number: 10, Value: "Deich"},
		{Number: 9, Value: "Naoi"},
		{Number: 8, Value: "Ochd"},
		{Number: 7, Value: "Seachd"},
		{Number: 6, Value: "Sia"},
		{Number: 5, Value: "Còig"},
		{Number: 4, Value: "Ceithir"},
		{Number: 3, Value: "Trì"},
		{Number: 2, Value: "Dà"},
		{Number: 1, Value: "Aon"},
		{Number: 0, Value: "Neoni"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Ceud"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "Ciad", Suffix: "mh", Masculine: "Ciad", Feminine: "Chiad", Neuter: "Ciad"},
		{Number: 2, Word: "Darna", Suffix: "mh", Masculine: "Darna", Feminine: "Darna", Neuter: "Darna"},
		{Number: 3, Word: "Treas", Suffix: "mh", Masculine: "Treas", Feminine: "Threas", Neuter: "Treas"},
		{Number: 4, Word: "Ceathramh", Suffix: "mh", Masculine: "Ceathramh", Feminine: "Cheathramh", Neuter: "Ceathramh"},
		{Number: 5, Word: "Còigeamh", Suffix: "mh", Masculine: "Còigeamh", Feminine: "Chòigeamh", Neuter: "Còigeamh"},
		{Number: 6, Word: "Siathamh", Suffix: "mh", Masculine: "Siathamh", Feminine: "Shiathamh", Neuter: "Siathamh"},
		{Number: 7, Word: "Seachdamh", Suffix: "mh", Masculine: "Seachdamh", Feminine: "Sheachdamh", Neuter: "Seachdamh"},
		{Number: 8, Word: "Ochdamh", Suffix: "mh", Masculine: "Ochdamh", Feminine: "Ochdamh", Neuter: "Ochdamh"},
		{Number: 9, Word: "Naoidheamh", Suffix: "mh", Masculine: "Naoidheamh", Feminine: "Naoidheamh", Neuter: "Naoidheamh"},
		{Number: 10, Word: "Deicheamh", Suffix: "mh", Masculine: "Deicheamh", Feminine: "Dheicheamh", Neuter: "Deicheamh"},
		{Number: 11, Word: "Aon deug", Suffix: "mh", Masculine: "Aon deug", Feminine: "Aon deug", Neuter: "Aon deug"},
		{Number: 12, Word: "Dà dheug", Suffix: "mh", Masculine: "Dà dheug", Feminine: "Dà dheug", Neuter: "Dà dheug"},
		{Number: 13, Word: "Trì deug", Suffix: "mh", Masculine: "Trì deug", Feminine: "Trì deug", Neuter: "Trì deug"},
		{Number: 14, Word: "Ceithir deug", Suffix: "mh", Masculine: "Ceithir deug", Feminine: "Ceithir deug", Neuter: "Ceithir deug"},
		{Number: 15, Word: "Còig deug", Suffix: "mh", Masculine: "Còig deug", Feminine: "Còig deug", Neuter: "Còig deug"},
		{Number: 16, Word: "Sia deug", Suffix: "mh", Masculine: "Sia deug", Feminine: "Sia deug", Neuter: "Sia deug"},
		{Number: 17, Word: "Seachd deug", Suffix: "mh", Masculine: "Seachd deug", Feminine: "Seachd deug", Neuter: "Seachd deug"},
		{Number: 18, Word: "Ochd deug", Suffix: "mh", Masculine: "Ochd deug", Feminine: "Ochd deug", Neuter: "Ochd deug"},
		{Number: 19, Word: "Naoi deug", Suffix: "mh", Masculine: "Naoi deug", Feminine: "Naoi deug", Neuter: "Naoi deug"},
		{Number: 20, Word: "Ficheadamh", Suffix: "mh", Masculine: "Ficheadamh", Feminine: "Fhicheadamh", Neuter: "Ficheadamh"},
		{Number: 21, Word: "Fichead 's a h-aon", Suffix: "mh", Masculine: "Fichead 's a h-aon", Feminine: "Fichead 's a h-aon", Neuter: "Fichead 's a h-aon"},
		{Number: 30, Word: "Trithead", Suffix: "mh", Masculine: "Trithead", Feminine: "Thrithead", Neuter: "Trithead"},
		{Number: 40, Word: "Dà fhicheadamh", Suffix: "mh", Masculine: "Dà fhicheadamh", Feminine: "Dà fhicheadamh", Neuter: "Dà fhicheadamh"},
		{Number: 50, Word: "Leth-cheudamh", Suffix: "mh", Masculine: "Leth-cheudamh", Feminine: "Leth-cheudamh", Neuter: "Leth-cheudamh"},
		{Number: 60, Word: "Seasgadamh", Suffix: "mh", Masculine: "Seasgadamh", Feminine: "Sheasgadamh", Neuter: "Seasgadamh"},
		{Number: 70, Word: "Seachdadamh", Suffix: "mh", Masculine: "Seachdadamh", Feminine: "Sheachdadamh", Neuter: "Seachdadamh"},
		{Number: 80, Word: "Ochdadamh", Suffix: "mh", Masculine: "Ochdadamh", Feminine: "Ochdadamh", Neuter: "Ochdadamh"},
		{Number: 90, Word: "Naochdadamh", Suffix: "mh", Masculine: "Naochdadamh", Feminine: "Naochdadamh", Neuter: "Naochdadamh"},
		{Number: 100, Word: "Ceudamh", Suffix: "mh", Masculine: "Ceudamh", Feminine: "Cheudamh", Neuter: "Ceudamh"},
		{Number: 1000, Word: "Mìleamh", Suffix: "mh", Masculine: "Mìleamh", Feminine: "Mhìleamh", Neuter: "Mìleamh"},
	},
	LocaleFormatter: &ScottishGaelicFormatter{},
}

// ScottishGaelicFormatter handles Scottish Gaelic (gd-GB) formatting
type ScottishGaelicFormatter struct{}

func (f *ScottishGaelicFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *ScottishGaelicFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *ScottishGaelicFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *ScottishGaelicFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *ScottishGaelicFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *ScottishGaelicFormatter) ChopDecimal(num decimal.Decimal, precision int) decimal.Decimal {
	return num.Truncate(int32(precision))
}

func (f *ScottishGaelicFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *ScottishGaelicFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
