package locale

import "github.com/shopspring/decimal"

// DADKLocale is a NumI18NLocale configured for Danish (Denmark) - da-DK
var DADKLocale = NumI18NLocale{
	LocaleFormatter: &DanishFormatter{},
	Currency: Currency{
		Name:     "Krone",
		Plural:   "Kroner",
		Singular: "Krone",
		Symbol:   "kr",
		FractionUnit: FractionUnit{
			Name:     "Øre",
			Plural:   "Øre",
			Singular: "Øre",
			Symbol:   "ø",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Denmark",
		Currency:       "DKK",
		ISO3166Alpha2:  "DK",
		ISO3166Alpha3:  "DNK",
		ISO3166Numeric: "208",
		Locale:         "da-DK",
		Timezone:       []string{"Europe/Copenhagen"},
		Language:       "da",
	},
	Texts: Texts{
		And:   "og",
		Minus: "minus",
		Only:  "kun",
		Point: "punktum",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Kvadrillion"},
		{Number: 1000000000000, Value: "Billion"},
		{Number: 1000000000, Value: "Milliard"},
		{Number: 1000000, Value: "Million"},
		{Number: 1000, Value: "Tusind"},
		{Number: 100, Value: "Hundrede"},
		{Number: 90, Value: "Halvfems"},
		{Number: 80, Value: "Firs"},
		{Number: 70, Value: "Halvfjerds"},
		{Number: 60, Value: "Tres"},
		{Number: 50, Value: "Halvtreds"},
		{Number: 40, Value: "Fyrre"},
		{Number: 30, Value: "Tredive"},
		{Number: 20, Value: "Tyve"},
		{Number: 19, Value: "Nitten"},
		{Number: 18, Value: "Atten"},
		{Number: 17, Value: "Sytten"},
		{Number: 16, Value: "Seksten"},
		{Number: 15, Value: "Femten"},
		{Number: 14, Value: "Fjorten"},
		{Number: 13, Value: "Tretten"},
		{Number: 12, Value: "Tolv"},
		{Number: 11, Value: "Elleve"},
		{Number: 10, Value: "Ti"},
		{Number: 9, Value: "Ni"},
		{Number: 8, Value: "Otte"},
		{Number: 7, Value: "Syv"},
		{Number: 6, Value: "Seks"},
		{Number: 5, Value: "Fem"},
		{Number: 4, Value: "Fire"},
		{Number: 3, Value: "Tre"},
		{Number: 2, Value: "To"},
		{Number: 1, Value: "En"},
		{Number: 0, Value: "Nul"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Hundrede"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "første", Suffix: ".", Masculine: "første", Feminine: "første", Neuter: "første"},
		{Number: 2, Word: "anden", Suffix: ".", Masculine: "anden", Feminine: "anden", Neuter: "andet"},
		{Number: 3, Word: "tredje", Suffix: ".", Masculine: "tredje", Feminine: "tredje", Neuter: "tredje"},
		{Number: 4, Word: "fjerde", Suffix: ".", Masculine: "fjerde", Feminine: "fjerde", Neuter: "fjerde"},
		{Number: 5, Word: "femte", Suffix: ".", Masculine: "femte", Feminine: "femte", Neuter: "femte"},
		{Number: 6, Word: "sjette", Suffix: ".", Masculine: "sjette", Feminine: "sjette", Neuter: "sjette"},
		{Number: 7, Word: "syvende", Suffix: ".", Masculine: "syvende", Feminine: "syvende", Neuter: "syvende"},
		{Number: 8, Word: "ottende", Suffix: ".", Masculine: "ottende", Feminine: "ottende", Neuter: "ottende"},
		{Number: 9, Word: "niende", Suffix: ".", Masculine: "niende", Feminine: "niende", Neuter: "niende"},
		{Number: 10, Word: "tiende", Suffix: ".", Masculine: "tiende", Feminine: "tiende", Neuter: "tiende"},
		{Number: 11, Word: "ellevte", Suffix: ".", Masculine: "ellevte", Feminine: "ellevte", Neuter: "ellevte"},
		{Number: 12, Word: "tolvte", Suffix: ".", Masculine: "tolvte", Feminine: "tolvte", Neuter: "tolvte"},
		{Number: 13, Word: "trettende", Suffix: ".", Masculine: "trettende", Feminine: "trettende", Neuter: "trettende"},
		{Number: 14, Word: "fjortende", Suffix: ".", Masculine: "fjortende", Feminine: "fjortende", Neuter: "fjortende"},
		{Number: 15, Word: "femtende", Suffix: ".", Masculine: "femtende", Feminine: "femtende", Neuter: "femtende"},
		{Number: 16, Word: "sekstende", Suffix: ".", Masculine: "sekstende", Feminine: "sekstende", Neuter: "sekstende"},
		{Number: 17, Word: "syttende", Suffix: ".", Masculine: "syttende", Feminine: "syttende", Neuter: "syttende"},
		{Number: 18, Word: "attende", Suffix: ".", Masculine: "attende", Feminine: "attende", Neuter: "attende"},
		{Number: 19, Word: "nittende", Suffix: ".", Masculine: "nittende", Feminine: "nittende", Neuter: "nittende"},
		{Number: 20, Word: "tyvende", Suffix: ".", Masculine: "tyvende", Feminine: "tyvende", Neuter: "tyvende"},
		{Number: 21, Word: "enogtyvende", Suffix: ".", Masculine: "enogtyvende", Feminine: "enogtyvende", Neuter: "enogtyvende"},
		{Number: 30, Word: "tredivte", Suffix: ".", Masculine: "tredivte", Feminine: "tredivte", Neuter: "tredivte"},
		{Number: 40, Word: "fyrrende", Suffix: ".", Masculine: "fyrrende", Feminine: "fyrrende", Neuter: "fyrrende"},
		{Number: 50, Word: "halvtredsindstyvende", Suffix: ".", Masculine: "halvtredsindstyvende", Feminine: "halvtredsindstyvende", Neuter: "halvtredsindstyvende"},
		{Number: 60, Word: "tresindstyvende", Suffix: ".", Masculine: "tresindstyvende", Feminine: "tresindstyvende", Neuter: "tresindstyvende"},
		{Number: 70, Word: "halvfjerdsindstyvende", Suffix: ".", Masculine: "halvfjerdsindstyvende", Feminine: "halvfjerdsindstyvende", Neuter: "halvfjerdsindstyvende"},
		{Number: 80, Word: "firsindstyvende", Suffix: ".", Masculine: "firsindstyvende", Feminine: "firsindstyvende", Neuter: "firsindstyvende"},
		{Number: 90, Word: "halvfemsindstyvende", Suffix: ".", Masculine: "halvfemsindstyvende", Feminine: "halvfemsindstyvende", Neuter: "halvfemsindstyvende"},
		{Number: 100, Word: "hundredende", Suffix: ".", Masculine: "hundredende", Feminine: "hundredende", Neuter: "hundredende"},
		{Number: 1000, Word: "tusindende", Suffix: ".", Masculine: "tusindende", Feminine: "tusindende", Neuter: "tusindende"},
	},
}

// DanishFormatter handles Danish formatting
type DanishFormatter struct{}

func (f *DanishFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *DanishFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *DanishFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *DanishFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *DanishFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *DanishFormatter) ChopDecimal(value decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 2
	}
	return value.Truncate(int32(precision))
}
