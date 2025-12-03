package locale

import (
	"github.com/shopspring/decimal"
)

// NONOLocale represents the Norwegian (Norway) locale
var NONOLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Norsk krone",
		Plural:   "Norske kroner",
		Singular: "Norsk krone",
		Symbol:   "kr",
		FractionUnit: FractionUnit{
			Name:     "Øre",
			Plural:   "Øre",
			Singular: "Øre",
			Symbol:   "øre",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Norway",
		Currency:       "NOK",
		ISO3166Alpha2:  "NO",
		ISO3166Alpha3:  "NOR",
		ISO3166Numeric: "578",
		Locale:         "no-NO",
		Timezone:       []string{"Europe/Oslo"},
		Language:       "no",
		Emoji:          "🇳🇴",
	},
	Texts: Texts{
		And:   "og",
		Minus: "minus",
		Only:  "bare",
		Point: "komma",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "tusen billioner"},
		{Number: 1000000000000, Value: "billion"},
		{Number: 1000000000, Value: "milliard"},
		{Number: 1000000, Value: "million"},
		{Number: 100000, Value: "hundretusen"},
		{Number: 90000, Value: "nittitusen"},
		{Number: 80000, Value: "åttitusen"},
		{Number: 70000, Value: "syttitusen"},
		{Number: 60000, Value: "sekstusen"},
		{Number: 50000, Value: "femtitusen"},
		{Number: 40000, Value: "førtitusen"},
		{Number: 30000, Value: "trettitusen"},
		{Number: 20000, Value: "tjuetusen"},
		{Number: 19000, Value: "nittentusen"},
		{Number: 18000, Value: "attentusen"},
		{Number: 17000, Value: "syttentusen"},
		{Number: 16000, Value: "sekstentusen"},
		{Number: 15000, Value: "femtentusen"},
		{Number: 14000, Value: "fjortentusen"},
		{Number: 13000, Value: "trettentusen"},
		{Number: 12000, Value: "tolvtusen"},
		{Number: 11000, Value: "ellevetusen"},
		{Number: 10000, Value: "titusen"},
		{Number: 9000, Value: "nitusen"},
		{Number: 8000, Value: "åttetusen"},
		{Number: 7000, Value: "syv tusen"},
		{Number: 6000, Value: "seks tusen"},
		{Number: 5000, Value: "fem tusen"},
		{Number: 4000, Value: "fire tusen"},
		{Number: 3000, Value: "tre tusen"},
		{Number: 2000, Value: "to tusen"},
		{Number: 1000, Value: "tusen"},
		{Number: 900, Value: "ni hundre"},
		{Number: 800, Value: "åtte hundre"},
		{Number: 700, Value: "syv hundre"},
		{Number: 600, Value: "seks hundre"},
		{Number: 500, Value: "fem hundre"},
		{Number: 400, Value: "fire hundre"},
		{Number: 300, Value: "tre hundre"},
		{Number: 200, Value: "to hundre"},
		{Number: 100, Value: "hundre"},
		{Number: 99, Value: "nittini"},
		{Number: 98, Value: "nittiåtte"},
		{Number: 97, Value: "nittisyv"},
		{Number: 96, Value: "nittiseks"},
		{Number: 95, Value: "nittifem"},
		{Number: 94, Value: "nittifire"},
		{Number: 93, Value: "nittitre"},
		{Number: 92, Value: "nittito"},
		{Number: 91, Value: "nittien"},
		{Number: 90, Value: "nitti"},
		{Number: 89, Value: "åttini"},
		{Number: 88, Value: "åttiåtte"},
		{Number: 87, Value: "åttisyv"},
		{Number: 86, Value: "åttiseks"},
		{Number: 85, Value: "åttifem"},
		{Number: 84, Value: "åttifire"},
		{Number: 83, Value: "åttitre"},
		{Number: 82, Value: "åttito"},
		{Number: 81, Value: "åttien"},
		{Number: 80, Value: "åtti"},
		{Number: 79, Value: "syttini"},
		{Number: 78, Value: "syttiåtte"},
		{Number: 77, Value: "syttisyv"},
		{Number: 76, Value: "syttiseks"},
		{Number: 75, Value: "syttifem"},
		{Number: 74, Value: "syttifire"},
		{Number: 73, Value: "syttitre"},
		{Number: 72, Value: "syttito"},
		{Number: 71, Value: "syttien"},
		{Number: 70, Value: "sytti"},
		{Number: 69, Value: "seksini"},
		{Number: 68, Value: "seksiåtte"},
		{Number: 67, Value: "seksisyv"},
		{Number: 66, Value: "seksiseks"},
		{Number: 65, Value: "seksifem"},
		{Number: 64, Value: "seksifire"},
		{Number: 63, Value: "seksitre"},
		{Number: 62, Value: "seksito"},
		{Number: 61, Value: "seksien"},
		{Number: 60, Value: "seksti"},
		{Number: 59, Value: "femtini"},
		{Number: 58, Value: "femtiåtte"},
		{Number: 57, Value: "femtisyv"},
		{Number: 56, Value: "femtiseks"},
		{Number: 55, Value: "femtifem"},
		{Number: 54, Value: "femtifire"},
		{Number: 53, Value: "femtitre"},
		{Number: 52, Value: "femtito"},
		{Number: 51, Value: "femtien"},
		{Number: 50, Value: "femti"},
		{Number: 49, Value: "førtini"},
		{Number: 48, Value: "førtiåtte"},
		{Number: 47, Value: "førtisyv"},
		{Number: 46, Value: "førtiseks"},
		{Number: 45, Value: "førtifem"},
		{Number: 44, Value: "førtifire"},
		{Number: 43, Value: "førtitre"},
		{Number: 42, Value: "førtito"},
		{Number: 41, Value: "førtien"},
		{Number: 40, Value: "førti"},
		{Number: 39, Value: "trettini"},
		{Number: 38, Value: "trettiåtte"},
		{Number: 37, Value: "trettisyv"},
		{Number: 36, Value: "trettiseks"},
		{Number: 35, Value: "trettifem"},
		{Number: 34, Value: "trettifire"},
		{Number: 33, Value: "trettitre"},
		{Number: 32, Value: "trettito"},
		{Number: 31, Value: "trettien"},
		{Number: 30, Value: "tretti"},
		{Number: 29, Value: "tjueni"},
		{Number: 28, Value: "tjueåtte"},
		{Number: 27, Value: "tjuesyv"},
		{Number: 26, Value: "tjueseks"},
		{Number: 25, Value: "tjuefem"},
		{Number: 24, Value: "tjuefire"},
		{Number: 23, Value: "tjuetre"},
		{Number: 22, Value: "tjueto"},
		{Number: 21, Value: "tjueen"},
		{Number: 20, Value: "tjue"},
		{Number: 19, Value: "nitten"},
		{Number: 18, Value: "atten"},
		{Number: 17, Value: "sytten"},
		{Number: 16, Value: "seksten"},
		{Number: 15, Value: "femten"},
		{Number: 14, Value: "fjorten"},
		{Number: 13, Value: "tretten"},
		{Number: 12, Value: "tolv"},
		{Number: 11, Value: "elleve"},
		{Number: 10, Value: "ti"},
		{Number: 9, Value: "ni"},
		{Number: 8, Value: "åtte"},
		{Number: 7, Value: "syv"},
		{Number: 6, Value: "seks"},
		{Number: 5, Value: "fem"},
		{Number: 4, Value: "fire"},
		{Number: 3, Value: "tre"},
		{Number: 2, Value: "to"},
		{Number: 1, Value: "en"},
		{Number: 0, Value: "null"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Ett hundre"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "første", Suffix: "-te", Masculine: "første", Feminine: "første", Neuter: "første"},
		{Number: 2, Word: "andre", Suffix: "-te", Masculine: "andre", Feminine: "andre", Neuter: "andre"},
		{Number: 3, Word: "tredje", Suffix: "-te", Masculine: "tredje", Feminine: "tredje", Neuter: "tredje"},
		{Number: 4, Word: "fjerde", Suffix: "-te", Masculine: "fjerde", Feminine: "fjerde", Neuter: "fjerde"},
		{Number: 5, Word: "femte", Suffix: "-te", Masculine: "femte", Feminine: "femte", Neuter: "femte"},
		{Number: 6, Word: "sjette", Suffix: "-te", Masculine: "sjette", Feminine: "sjette", Neuter: "sjette"},
		{Number: 7, Word: "syvende", Suffix: "-ende", Masculine: "syvende", Feminine: "syvende", Neuter: "syvende"},
		{Number: 8, Word: "åttende", Suffix: "-ende", Masculine: "åttende", Feminine: "åttende", Neuter: "åttende"},
		{Number: 9, Word: "niende", Suffix: "-ende", Masculine: "niende", Feminine: "niende", Neuter: "niende"},
		{Number: 10, Word: "tiende", Suffix: "-ende", Masculine: "tiende", Feminine: "tiende", Neuter: "tiende"},
		{Number: 11, Word: "ellevte", Suffix: "-te", Masculine: "ellevte", Feminine: "ellevte", Neuter: "ellevte"},
		{Number: 12, Word: "tolvte", Suffix: "-te", Masculine: "tolvte", Feminine: "tolvte", Neuter: "tolvte"},
		{Number: 20, Word: "tjuende", Suffix: "-ende", Masculine: "tjuende", Feminine: "tjuende", Neuter: "tjuende"},
		{Number: 21, Word: "tjueførste", Suffix: "-te", Masculine: "tjueførste", Feminine: "tjueførste", Neuter: "tjueførste"},
		{Number: 30, Word: "tredvende", Suffix: "-ende", Masculine: "tredvende", Feminine: "tredvende", Neuter: "tredvende"},
		{Number: 100, Word: "hundrede", Suffix: "-de", Masculine: "hundrede", Feminine: "hundrede", Neuter: "hundrede"},
		{Number: 1000, Word: "tusende", Suffix: "-ende", Masculine: "tusende", Feminine: "tusende", Neuter: "tusende"},
	},
	LocaleFormatter: &NorwegianBokmålFormatter{},
}

// NorwegianBokmålFormatter handles Norwegian Bokmål (no-NO) formatting
type NorwegianBokmålFormatter struct{}

func (f *NorwegianBokmålFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *NorwegianBokmålFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *NorwegianBokmålFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *NorwegianBokmålFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *NorwegianBokmålFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *NorwegianBokmålFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 2
	}
	return amount.Truncate(int32(precision))
}

func (f *NorwegianBokmålFormatter) FormatDecimalNumber(amount float64) string {
	return FormatEuropeanDecimal(amount)
}

func (f *NorwegianBokmålFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with European conventions (period separators, comma decimal, prefix symbol)
	return FormatEuropeanCurrency(amount, currencySymbol)
}
