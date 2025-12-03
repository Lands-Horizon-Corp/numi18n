package locale

import (
	"github.com/shopspring/decimal"
)

// SVSELocale represents the Swedish (Sweden) locale
var SVSELocale = NumI18NLocale{
	LocaleFormatter: &SwedishSwedenFormatter{},
	Currency: Currency{
		Name:     "Swedish Krona",
		Plural:   "Kronor",
		Singular: "Krona",
		Symbol:   "kr",
		FractionUnit: FractionUnit{
			Name:     "Öre",
			Plural:   "Öre",
			Singular: "Öre",
			Symbol:   "ö",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Sweden",
		Currency:       "SEK",
		ISO3166Alpha2:  "SE",
		ISO3166Alpha3:  "SWE",
		ISO3166Numeric: "752",
		Locale:         "sv-SE",
		Timezone:       []string{"Europe/Stockholm"},
		Language:       "sv",
		Emoji:          "🇸🇪",
		PhoneCode:      "+46",
		Domain:         ".se",
	},
	Texts: Texts{
		And:   "och",
		Minus: "minus",
		Only:  "endast",
		Point: "komma",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "en biljard"},
		{Number: 1000000000000, Value: "en biljon"},
		{Number: 1000000000, Value: "en miljard"},
		{Number: 1000000, Value: "en miljon"},
		{Number: 100000, Value: "etthundratusen"},
		{Number: 90000, Value: "nittiotusen"},
		{Number: 80000, Value: "åttiotusen"},
		{Number: 70000, Value: "sjuttiotusen"},
		{Number: 60000, Value: "sextiotusen"},
		{Number: 50000, Value: "femtiotusen"},
		{Number: 40000, Value: "fyrtiotusen"},
		{Number: 30000, Value: "trettiotusen"},
		{Number: 20000, Value: "tjugotusen"},
		{Number: 19000, Value: "nittontusen"},
		{Number: 18000, Value: "artontusen"},
		{Number: 17000, Value: "sjuttontusen"},
		{Number: 16000, Value: "sextontusen"},
		{Number: 15000, Value: "femtontusen"},
		{Number: 14000, Value: "fjortontusen"},
		{Number: 13000, Value: "trettontusen"},
		{Number: 12000, Value: "tolvtusen"},
		{Number: 11000, Value: "elvatusen"},
		{Number: 10000, Value: "tiotusen"},
		{Number: 9000, Value: "niotusen"},
		{Number: 8000, Value: "åttatusen"},
		{Number: 7000, Value: "sjutusen"},
		{Number: 6000, Value: "sextusen"},
		{Number: 5000, Value: "femtusen"},
		{Number: 4000, Value: "fyratusen"},
		{Number: 3000, Value: "tretusen"},
		{Number: 2000, Value: "tvåtusen"},
		{Number: 1000, Value: "ettusen"},
		{Number: 900, Value: "niohundra"},
		{Number: 800, Value: "åttahundra"},
		{Number: 700, Value: "sjuhundra"},
		{Number: 600, Value: "sexhundra"},
		{Number: 500, Value: "femhundra"},
		{Number: 400, Value: "fyrahundra"},
		{Number: 300, Value: "trehundra"},
		{Number: 200, Value: "tvåhundra"},
		{Number: 100, Value: "etthundra"},
		{Number: 99, Value: "nittionio"},
		{Number: 98, Value: "nittioåtta"},
		{Number: 97, Value: "nittiosju"},
		{Number: 96, Value: "nittiosex"},
		{Number: 95, Value: "nittiofem"},
		{Number: 94, Value: "nittiofyra"},
		{Number: 93, Value: "nittiotre"},
		{Number: 92, Value: "nittiotvå"},
		{Number: 91, Value: "nittioett"},
		{Number: 90, Value: "nittio"},
		{Number: 89, Value: "åttionio"},
		{Number: 88, Value: "åttioåtta"},
		{Number: 87, Value: "åttiosju"},
		{Number: 86, Value: "åttiosex"},
		{Number: 85, Value: "åttiofem"},
		{Number: 84, Value: "åttiofyra"},
		{Number: 83, Value: "åttiotre"},
		{Number: 82, Value: "åttiotvå"},
		{Number: 81, Value: "åttioett"},
		{Number: 80, Value: "åttio"},
		{Number: 79, Value: "sjuttionio"},
		{Number: 78, Value: "sjuttioåtta"},
		{Number: 77, Value: "sjuttiosju"},
		{Number: 76, Value: "sjuttiosex"},
		{Number: 75, Value: "sjuttiofem"},
		{Number: 74, Value: "sjuttiofyra"},
		{Number: 73, Value: "sjuttiotre"},
		{Number: 72, Value: "sjuttiotvå"},
		{Number: 71, Value: "sjuttioett"},
		{Number: 70, Value: "sjuttio"},
		{Number: 69, Value: "sextionio"},
		{Number: 68, Value: "sextioåtta"},
		{Number: 67, Value: "sextiosju"},
		{Number: 66, Value: "sextiosex"},
		{Number: 65, Value: "sextiofem"},
		{Number: 64, Value: "sextiofyra"},
		{Number: 63, Value: "sextiotre"},
		{Number: 62, Value: "sextiotvå"},
		{Number: 61, Value: "sextioett"},
		{Number: 60, Value: "sextio"},
		{Number: 59, Value: "femtionio"},
		{Number: 58, Value: "femtioåtta"},
		{Number: 57, Value: "femtiosju"},
		{Number: 56, Value: "femtiosex"},
		{Number: 55, Value: "femtiofem"},
		{Number: 54, Value: "femtiofyra"},
		{Number: 53, Value: "femtiotre"},
		{Number: 52, Value: "femtiotvå"},
		{Number: 51, Value: "femtioett"},
		{Number: 50, Value: "femtio"},
		{Number: 49, Value: "fyrtionio"},
		{Number: 48, Value: "fyrtioåtta"},
		{Number: 47, Value: "fyrtiosju"},
		{Number: 46, Value: "fyrtiosex"},
		{Number: 45, Value: "fyrtiofem"},
		{Number: 44, Value: "fyrtiofyra"},
		{Number: 43, Value: "fyrtiotre"},
		{Number: 42, Value: "fyrtiotvå"},
		{Number: 41, Value: "fyrtioett"},
		{Number: 40, Value: "fyrtio"},
		{Number: 39, Value: "trettionio"},
		{Number: 38, Value: "trettioåtta"},
		{Number: 37, Value: "trettiosju"},
		{Number: 36, Value: "trettiosex"},
		{Number: 35, Value: "trettiofem"},
		{Number: 34, Value: "trettiofyra"},
		{Number: 33, Value: "trettiotre"},
		{Number: 32, Value: "trettiotvå"},
		{Number: 31, Value: "trettioett"},
		{Number: 30, Value: "trettio"},
		{Number: 29, Value: "tjugonio"},
		{Number: 28, Value: "tjugoåtta"},
		{Number: 27, Value: "tjugosju"},
		{Number: 26, Value: "tjugosex"},
		{Number: 25, Value: "tjugofem"},
		{Number: 24, Value: "tjugofyra"},
		{Number: 23, Value: "tjugotre"},
		{Number: 22, Value: "tjugotvå"},
		{Number: 21, Value: "tjugoett"},
		{Number: 20, Value: "tjugo"},
		{Number: 19, Value: "nitton"},
		{Number: 18, Value: "arton"},
		{Number: 17, Value: "sjutton"},
		{Number: 16, Value: "sexton"},
		{Number: 15, Value: "femton"},
		{Number: 14, Value: "fjorton"},
		{Number: 13, Value: "tretton"},
		{Number: 12, Value: "tolv"},
		{Number: 11, Value: "elva"},
		{Number: 10, Value: "tio"},
		{Number: 9, Value: "nio"},
		{Number: 8, Value: "åtta"},
		{Number: 7, Value: "sju"},
		{Number: 6, Value: "sex"},
		{Number: 5, Value: "fem"},
		{Number: 4, Value: "fyra"},
		{Number: 3, Value: "tre"},
		{Number: 2, Value: "två"},
		{Number: 1, Value: "ett"},
		{Number: 0, Value: "noll"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Etthundra"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "första", Suffix: ":a", Masculine: "första", Feminine: "första", Neuter: "första"},
		{Number: 2, Word: "andra", Suffix: ":a", Masculine: "andra", Feminine: "andra", Neuter: "andra"},
		{Number: 3, Word: "tredje", Suffix: ":e", Masculine: "tredje", Feminine: "tredje", Neuter: "tredje"},
		{Number: 4, Word: "fjärde", Suffix: ":e", Masculine: "fjärde", Feminine: "fjärde", Neuter: "fjärde"},
		{Number: 5, Word: "femte", Suffix: ":e", Masculine: "femte", Feminine: "femte", Neuter: "femte"},
		{Number: 6, Word: "sjätte", Suffix: ":e", Masculine: "sjätte", Feminine: "sjätte", Neuter: "sjätte"},
		{Number: 7, Word: "sjunde", Suffix: ":e", Masculine: "sjunde", Feminine: "sjunde", Neuter: "sjunde"},
		{Number: 8, Word: "åttonde", Suffix: ":e", Masculine: "åttonde", Feminine: "åttonde", Neuter: "åttonde"},
		{Number: 9, Word: "nionde", Suffix: ":e", Masculine: "nionde", Feminine: "nionde", Neuter: "nionde"},
		{Number: 10, Word: "tionde", Suffix: ":e", Masculine: "tionde", Feminine: "tionde", Neuter: "tionde"},
		{Number: 11, Word: "elfte", Suffix: ":e", Masculine: "elfte", Feminine: "elfte", Neuter: "elfte"},
		{Number: 12, Word: "tolfte", Suffix: ":e", Masculine: "tolfte", Feminine: "tolfte", Neuter: "tolfte"},
		{Number: 20, Word: "tjugonde", Suffix: ":e", Masculine: "tjugonde", Feminine: "tjugonde", Neuter: "tjugonde"},
		{Number: 21, Word: "tjugoförsta", Suffix: ":a", Masculine: "tjugoförsta", Feminine: "tjugoförsta", Neuter: "tjugoförsta"},
		{Number: 30, Word: "trettionde", Suffix: ":e", Masculine: "trettionde", Feminine: "trettionde", Neuter: "trettionde"},
		{Number: 50, Word: "femtionde", Suffix: ":e", Masculine: "femtionde", Feminine: "femtionde", Neuter: "femtionde"},
		{Number: 100, Word: "hundrade", Suffix: ":e", Masculine: "hundrade", Feminine: "hundrade", Neuter: "hundrade"},
		{Number: 1000, Word: "tusende", Suffix: ":e", Masculine: "tusende", Feminine: "tusende", Neuter: "tusende"},
	},
}

// SwedishSwedenFormatter handles Swedish (Sweden) formatting
type SwedishSwedenFormatter struct{}

func (f *SwedishSwedenFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *SwedishSwedenFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *SwedishSwedenFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *SwedishSwedenFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *SwedishSwedenFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *SwedishSwedenFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	return amount.Truncate(int32(precision))
}

func (f *SwedishSwedenFormatter) FormatDecimalNumber(amount float64) string {
	return FormatEuropeanDecimal(amount)
}

func (f *SwedishSwedenFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with European conventions (period separators, comma decimal, prefix symbol)
	return FormatEuropeanCurrency(amount, currencySymbol)
}
