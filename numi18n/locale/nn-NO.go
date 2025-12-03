package locale

import (
	"github.com/shopspring/decimal"
)

// NNNOLocale represents the Norwegian Nynorsk (Norway) locale
var NNNOLocale = NumI18NLocale{
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
		Locale:         "nn-NO",
		Timezone:       []string{"Europe/Oslo"},
		Language:       "nn",
		Emoji:          "🇳🇴",
		PhoneCode:      "+47",
		Domain:         ".no",
	},
	Texts: Texts{
		And:   "og",
		Minus: "minus",
		Only:  "berre",
		Point: "komma",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "tusen billionar"},
		{Number: 1000000000000, Value: "billion"},
		{Number: 1000000000, Value: "milliard"},
		{Number: 1000000, Value: "million"},
		{Number: 100000, Value: "hundretusen"},
		{Number: 90000, Value: "nittitusen"},
		{Number: 80000, Value: "åttitusen"},
		{Number: 70000, Value: "sytitusen"},
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
		{Number: 7000, Value: "sju tusen"},
		{Number: 6000, Value: "seks tusen"},
		{Number: 5000, Value: "fem tusen"},
		{Number: 4000, Value: "fire tusen"},
		{Number: 3000, Value: "tre tusen"},
		{Number: 2000, Value: "to tusen"},
		{Number: 1000, Value: "tusen"},
		{Number: 900, Value: "ni hundre"},
		{Number: 800, Value: "åtte hundre"},
		{Number: 700, Value: "sju hundre"},
		{Number: 600, Value: "seks hundre"},
		{Number: 500, Value: "fem hundre"},
		{Number: 400, Value: "fire hundre"},
		{Number: 300, Value: "tre hundre"},
		{Number: 200, Value: "to hundre"},
		{Number: 100, Value: "hundre"},
		{Number: 99, Value: "nittini"},
		{Number: 98, Value: "nittiåtte"},
		{Number: 97, Value: "nittisju"},
		{Number: 96, Value: "nittiseks"},
		{Number: 95, Value: "nittifem"},
		{Number: 94, Value: "nittifire"},
		{Number: 93, Value: "nittitre"},
		{Number: 92, Value: "nittito"},
		{Number: 91, Value: "nittiein"},
		{Number: 90, Value: "nitti"},
		{Number: 89, Value: "åttini"},
		{Number: 88, Value: "åttiåtte"},
		{Number: 87, Value: "åttisju"},
		{Number: 86, Value: "åttiseks"},
		{Number: 85, Value: "åttifem"},
		{Number: 84, Value: "åttifire"},
		{Number: 83, Value: "åttitre"},
		{Number: 82, Value: "åttito"},
		{Number: 81, Value: "åttiein"},
		{Number: 80, Value: "åtti"},
		{Number: 79, Value: "sytini"},
		{Number: 78, Value: "sytiåtte"},
		{Number: 77, Value: "sytisju"},
		{Number: 76, Value: "sytiseks"},
		{Number: 75, Value: "sytifem"},
		{Number: 74, Value: "sytifire"},
		{Number: 73, Value: "sytitre"},
		{Number: 72, Value: "sytito"},
		{Number: 71, Value: "sytiein"},
		{Number: 70, Value: "syti"},
		{Number: 69, Value: "seksini"},
		{Number: 68, Value: "seksiåtte"},
		{Number: 67, Value: "seksisju"},
		{Number: 66, Value: "seksiseks"},
		{Number: 65, Value: "seksifem"},
		{Number: 64, Value: "seksifire"},
		{Number: 63, Value: "seksitre"},
		{Number: 62, Value: "seksito"},
		{Number: 61, Value: "seksiein"},
		{Number: 60, Value: "seksti"},
		{Number: 59, Value: "femtini"},
		{Number: 58, Value: "femtiåtte"},
		{Number: 57, Value: "femtisju"},
		{Number: 56, Value: "femtiseks"},
		{Number: 55, Value: "femtifem"},
		{Number: 54, Value: "femtifire"},
		{Number: 53, Value: "femtitre"},
		{Number: 52, Value: "femtito"},
		{Number: 51, Value: "femtiein"},
		{Number: 50, Value: "femti"},
		{Number: 49, Value: "førtini"},
		{Number: 48, Value: "førtiåtte"},
		{Number: 47, Value: "førtisju"},
		{Number: 46, Value: "førtiseks"},
		{Number: 45, Value: "førtifem"},
		{Number: 44, Value: "førtifire"},
		{Number: 43, Value: "førtitre"},
		{Number: 42, Value: "førtito"},
		{Number: 41, Value: "førtiein"},
		{Number: 40, Value: "førti"},
		{Number: 39, Value: "trettini"},
		{Number: 38, Value: "trettiåtte"},
		{Number: 37, Value: "trettisju"},
		{Number: 36, Value: "trettiseks"},
		{Number: 35, Value: "trettifem"},
		{Number: 34, Value: "trettifire"},
		{Number: 33, Value: "trettitre"},
		{Number: 32, Value: "trettito"},
		{Number: 31, Value: "trettiein"},
		{Number: 30, Value: "tretti"},
		{Number: 29, Value: "tjueni"},
		{Number: 28, Value: "tjueåtte"},
		{Number: 27, Value: "tjuesju"},
		{Number: 26, Value: "tjueseks"},
		{Number: 25, Value: "tjuefem"},
		{Number: 24, Value: "tjuefire"},
		{Number: 23, Value: "tjuetre"},
		{Number: 22, Value: "tjueto"},
		{Number: 21, Value: "tjueein"},
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
		{Number: 7, Value: "sju"},
		{Number: 6, Value: "seks"},
		{Number: 5, Value: "fem"},
		{Number: 4, Value: "fire"},
		{Number: 3, Value: "tre"},
		{Number: 2, Value: "to"},
		{Number: 1, Value: "ein"},
		{Number: 0, Value: "null"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Eitt hundre"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "første", Suffix: "-te", Masculine: "første", Feminine: "første", Neuter: "første"},
		{Number: 2, Word: "andre", Suffix: "-te", Masculine: "andre", Feminine: "andre", Neuter: "andre"},
		{Number: 3, Word: "tredje", Suffix: "-te", Masculine: "tredje", Feminine: "tredje", Neuter: "tredje"},
		{Number: 4, Word: "fjerde", Suffix: "-te", Masculine: "fjerde", Feminine: "fjerde", Neuter: "fjerde"},
		{Number: 5, Word: "femte", Suffix: "-te", Masculine: "femte", Feminine: "femte", Neuter: "femte"},
		{Number: 6, Word: "sjette", Suffix: "-te", Masculine: "sjette", Feminine: "sjette", Neuter: "sjette"},
		{Number: 7, Word: "sjuande", Suffix: "-ande", Masculine: "sjuande", Feminine: "sjuande", Neuter: "sjuande"},
		{Number: 8, Word: "åttande", Suffix: "-ande", Masculine: "åttande", Feminine: "åttande", Neuter: "åttande"},
		{Number: 9, Word: "niande", Suffix: "-ande", Masculine: "niande", Feminine: "niande", Neuter: "niande"},
		{Number: 10, Word: "tiande", Suffix: "-ande", Masculine: "tiande", Feminine: "tiande", Neuter: "tiande"},
		{Number: 11, Word: "ellevte", Suffix: "-te", Masculine: "ellevte", Feminine: "ellevte", Neuter: "ellevte"},
		{Number: 12, Word: "tolvte", Suffix: "-te", Masculine: "tolvte", Feminine: "tolvte", Neuter: "tolvte"},
		{Number: 20, Word: "tjuande", Suffix: "-ande", Masculine: "tjuande", Feminine: "tjuande", Neuter: "tjuande"},
		{Number: 21, Word: "tjueførste", Suffix: "-te", Masculine: "tjueførste", Feminine: "tjueførste", Neuter: "tjueførste"},
		{Number: 30, Word: "trettiande", Suffix: "-ande", Masculine: "trettiande", Feminine: "trettiande", Neuter: "trettiande"},
		{Number: 100, Word: "hundrade", Suffix: "-de", Masculine: "hundrade", Feminine: "hundrade", Neuter: "hundrade"},
		{Number: 1000, Word: "tusande", Suffix: "-ande", Masculine: "tusande", Feminine: "tusande", Neuter: "tusande"},
	},
	LocaleFormatter: &NorwegianNynorskFormatter{},
}

// NorwegianNynorskFormatter handles Norwegian Nynorsk (nn-NO) formatting
type NorwegianNynorskFormatter struct{}

func (f *NorwegianNynorskFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *NorwegianNynorskFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *NorwegianNynorskFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *NorwegianNynorskFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *NorwegianNynorskFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *NorwegianNynorskFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 2
	}
	return amount.Truncate(int32(precision))
}

func (f *NorwegianNynorskFormatter) FormatDecimalNumber(amount float64) string {
	return FormatPolishDecimal(amount)
}
func (f *NorwegianNynorskFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Polish conventions (comma separators, period decimal, prefix symbol)
	return FormatPolishCurrency(amount, currencySymbol)
}
