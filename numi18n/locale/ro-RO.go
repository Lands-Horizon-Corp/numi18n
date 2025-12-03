package locale

import (
	"github.com/shopspring/decimal"
)

// ROROLocale represents the Romanian (Romania) locale
var ROROLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Leu românesc",
		Plural:   "Lei românești",
		Singular: "Leu românesc",
		Symbol:   "RON",
		FractionUnit: FractionUnit{
			Name:     "Ban",
			Plural:   "Bani",
			Singular: "Ban",
			Symbol:   "b",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Romania",
		Currency:       "RON",
		ISO3166Alpha2:  "RO",
		ISO3166Alpha3:  "ROU",
		ISO3166Numeric: "642",
		Locale:         "ro-RO",
		Timezone:       []string{"Europe/Bucharest"},
		Language:       "ro",
		Emoji:          "🇷🇴",
		PhoneCode:      "+40",
		Domain:         ".ro",
	},
	Texts: Texts{
		And:   "și",
		Minus: "minus",
		Only:  "numai",
		Point: "virgulă",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "un biliard"},
		{Number: 1000000000000, Value: "un bilion"},
		{Number: 1000000000, Value: "un miliard"},
		{Number: 1000000, Value: "un milion"},
		{Number: 100000, Value: "o sută de mii"},
		{Number: 90000, Value: "nouăzeci de mii"},
		{Number: 80000, Value: "optzeci de mii"},
		{Number: 70000, Value: "șaptezeci de mii"},
		{Number: 60000, Value: "șaizeci de mii"},
		{Number: 50000, Value: "cincizeci de mii"},
		{Number: 40000, Value: "patruzeci de mii"},
		{Number: 30000, Value: "treizeci de mii"},
		{Number: 20000, Value: "douăzeci de mii"},
		{Number: 19000, Value: "nouăsprezece mii"},
		{Number: 18000, Value: "optsprezece mii"},
		{Number: 17000, Value: "șaptesprezece mii"},
		{Number: 16000, Value: "șaisprezece mii"},
		{Number: 15000, Value: "cincisprezece mii"},
		{Number: 14000, Value: "paisprezece mii"},
		{Number: 13000, Value: "treisprezece mii"},
		{Number: 12000, Value: "doisprezece mii"},
		{Number: 11000, Value: "unsprezece mii"},
		{Number: 10000, Value: "zece mii"},
		{Number: 9000, Value: "nouă mii"},
		{Number: 8000, Value: "opt mii"},
		{Number: 7000, Value: "șapte mii"},
		{Number: 6000, Value: "șase mii"},
		{Number: 5000, Value: "cinci mii"},
		{Number: 4000, Value: "patru mii"},
		{Number: 3000, Value: "trei mii"},
		{Number: 2000, Value: "două mii"},
		{Number: 1000, Value: "o mie"},
		{Number: 900, Value: "nouă sute"},
		{Number: 800, Value: "opt sute"},
		{Number: 700, Value: "șapte sute"},
		{Number: 600, Value: "șase sute"},
		{Number: 500, Value: "cinci sute"},
		{Number: 400, Value: "patru sute"},
		{Number: 300, Value: "trei sute"},
		{Number: 200, Value: "două sute"},
		{Number: 100, Value: "o sută"},
		{Number: 99, Value: "nouăzeci și nouă"},
		{Number: 98, Value: "nouăzeci și opt"},
		{Number: 97, Value: "nouăzeci și șapte"},
		{Number: 96, Value: "nouăzeci și șase"},
		{Number: 95, Value: "nouăzeci și cinci"},
		{Number: 94, Value: "nouăzeci și patru"},
		{Number: 93, Value: "nouăzeci și trei"},
		{Number: 92, Value: "nouăzeci și doi"},
		{Number: 91, Value: "nouăzeci și unu"},
		{Number: 90, Value: "nouăzeci"},
		{Number: 89, Value: "optzeci și nouă"},
		{Number: 88, Value: "optzeci și opt"},
		{Number: 87, Value: "optzeci și șapte"},
		{Number: 86, Value: "optzeci și șase"},
		{Number: 85, Value: "optzeci și cinci"},
		{Number: 84, Value: "optzeci și patru"},
		{Number: 83, Value: "optzeci și trei"},
		{Number: 82, Value: "optzeci și doi"},
		{Number: 81, Value: "optzeci și unu"},
		{Number: 80, Value: "optzeci"},
		{Number: 79, Value: "șaptezeci și nouă"},
		{Number: 78, Value: "șaptezeci și opt"},
		{Number: 77, Value: "șaptezeci și șapte"},
		{Number: 76, Value: "șaptezeci și șase"},
		{Number: 75, Value: "șaptezeci și cinci"},
		{Number: 74, Value: "șaptezeci și patru"},
		{Number: 73, Value: "șaptezeci și trei"},
		{Number: 72, Value: "șaptezeci și doi"},
		{Number: 71, Value: "șaptezeci și unu"},
		{Number: 70, Value: "șaptezeci"},
		{Number: 69, Value: "șaizeci și nouă"},
		{Number: 68, Value: "șaizeci și opt"},
		{Number: 67, Value: "șaizeci și șapte"},
		{Number: 66, Value: "șaizeci și șase"},
		{Number: 65, Value: "șaizeci și cinci"},
		{Number: 64, Value: "șaizeci și patru"},
		{Number: 63, Value: "șaizeci și trei"},
		{Number: 62, Value: "șaizeci și doi"},
		{Number: 61, Value: "șaizeci și unu"},
		{Number: 60, Value: "șaizeci"},
		{Number: 59, Value: "cincizeci și nouă"},
		{Number: 58, Value: "cincizeci și opt"},
		{Number: 57, Value: "cincizeci și șapte"},
		{Number: 56, Value: "cincizeci și șase"},
		{Number: 55, Value: "cincizeci și cinci"},
		{Number: 54, Value: "cincizeci și patru"},
		{Number: 53, Value: "cincizeci și trei"},
		{Number: 52, Value: "cincizeci și doi"},
		{Number: 51, Value: "cincizeci și unu"},
		{Number: 50, Value: "cincizeci"},
		{Number: 49, Value: "patruzeci și nouă"},
		{Number: 48, Value: "patruzeci și opt"},
		{Number: 47, Value: "patruzeci și șapte"},
		{Number: 46, Value: "patruzeci și șase"},
		{Number: 45, Value: "patruzeci și cinci"},
		{Number: 44, Value: "patruzeci și patru"},
		{Number: 43, Value: "patruzeci și trei"},
		{Number: 42, Value: "patruzeci și doi"},
		{Number: 41, Value: "patruzeci și unu"},
		{Number: 40, Value: "patruzeci"},
		{Number: 39, Value: "treizeci și nouă"},
		{Number: 38, Value: "treizeci și opt"},
		{Number: 37, Value: "treizeci și șapte"},
		{Number: 36, Value: "treizeci și șase"},
		{Number: 35, Value: "treizeci și cinci"},
		{Number: 34, Value: "treizeci și patru"},
		{Number: 33, Value: "treizeci și trei"},
		{Number: 32, Value: "treizeci și doi"},
		{Number: 31, Value: "treizeci și unu"},
		{Number: 30, Value: "treizeci"},
		{Number: 29, Value: "douăzeci și nouă"},
		{Number: 28, Value: "douăzeci și opt"},
		{Number: 27, Value: "douăzeci și șapte"},
		{Number: 26, Value: "douăzeci și șase"},
		{Number: 25, Value: "douăzeci și cinci"},
		{Number: 24, Value: "douăzeci și patru"},
		{Number: 23, Value: "douăzeci și trei"},
		{Number: 22, Value: "douăzeci și doi"},
		{Number: 21, Value: "douăzeci și unu"},
		{Number: 20, Value: "douăzeci"},
		{Number: 19, Value: "nouăsprezece"},
		{Number: 18, Value: "optsprezece"},
		{Number: 17, Value: "șaptesprezece"},
		{Number: 16, Value: "șaisprezece"},
		{Number: 15, Value: "cincisprezece"},
		{Number: 14, Value: "paisprezece"},
		{Number: 13, Value: "treisprezece"},
		{Number: 12, Value: "doisprezece"},
		{Number: 11, Value: "unsprezece"},
		{Number: 10, Value: "zece"},
		{Number: 9, Value: "nouă"},
		{Number: 8, Value: "opt"},
		{Number: 7, Value: "șapte"},
		{Number: 6, Value: "șase"},
		{Number: 5, Value: "cinci"},
		{Number: 4, Value: "patru"},
		{Number: 3, Value: "trei"},
		{Number: 2, Value: "doi"},
		{Number: 1, Value: "unu"},
		{Number: 0, Value: "zero"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "O sută"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "primul", Suffix: "-lea", Masculine: "primul", Feminine: "prima", Neuter: "primul"},
		{Number: 2, Word: "al doilea", Suffix: "-lea", Masculine: "al doilea", Feminine: "a doua", Neuter: "al doilea"},
		{Number: 3, Word: "al treilea", Suffix: "-lea", Masculine: "al treilea", Feminine: "a treia", Neuter: "al treilea"},
		{Number: 4, Word: "al patrulea", Suffix: "-lea", Masculine: "al patrulea", Feminine: "a patra", Neuter: "al patrulea"},
		{Number: 5, Word: "al cincilea", Suffix: "-lea", Masculine: "al cincilea", Feminine: "a cincea", Neuter: "al cincilea"},
		{Number: 6, Word: "al șaselea", Suffix: "-lea", Masculine: "al șaselea", Feminine: "a șasea", Neuter: "al șaselea"},
		{Number: 7, Word: "al șaptelea", Suffix: "-lea", Masculine: "al șaptelea", Feminine: "a șaptea", Neuter: "al șaptelea"},
		{Number: 8, Word: "al optulea", Suffix: "-lea", Masculine: "al optulea", Feminine: "a opta", Neuter: "al optulea"},
		{Number: 9, Word: "al nouălea", Suffix: "-lea", Masculine: "al nouălea", Feminine: "a noua", Neuter: "al nouălea"},
		{Number: 10, Word: "al zecelea", Suffix: "-lea", Masculine: "al zecelea", Feminine: "a zecea", Neuter: "al zecelea"},
		{Number: 11, Word: "al unsprezecelea", Suffix: "-lea", Masculine: "al unsprezecelea", Feminine: "a unsprezecea", Neuter: "al unsprezecelea"},
		{Number: 12, Word: "al doisprezecelea", Suffix: "-lea", Masculine: "al doisprezecelea", Feminine: "a doisprezecea", Neuter: "al doisprezecelea"},
		{Number: 20, Word: "al douăzecilea", Suffix: "-lea", Masculine: "al douăzecilea", Feminine: "a douăzecea", Neuter: "al douăzecilea"},
		{Number: 21, Word: "al douăzeci și unulea", Suffix: "-lea", Masculine: "al douăzeci și unulea", Feminine: "a douăzeci și una", Neuter: "al douăzeci și unulea"},
		{Number: 30, Word: "al treizecilea", Suffix: "-lea", Masculine: "al treizecilea", Feminine: "a treizecea", Neuter: "al treizecilea"},
		{Number: 50, Word: "al cincizecilea", Suffix: "-lea", Masculine: "al cincizecilea", Feminine: "a cincizecea", Neuter: "al cincizecilea"},
		{Number: 100, Word: "al o sutelea", Suffix: "-lea", Masculine: "al o sutelea", Feminine: "a o suta", Neuter: "al o sutelea"},
		{Number: 1000, Word: "al o miilea", Suffix: "-lea", Masculine: "al o miilea", Feminine: "a o mia", Neuter: "al o miilea"},
	},
	LocaleFormatter: &RomanianRomaniaFormatter{},
}

// RomanianRomaniaFormatter handles Romanian (Romania) formatting
type RomanianRomaniaFormatter struct{}

func (f *RomanianRomaniaFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *RomanianRomaniaFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *RomanianRomaniaFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *RomanianRomaniaFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *RomanianRomaniaFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *RomanianRomaniaFormatter) ChopDecimal(value decimal.Decimal, precision int) decimal.Decimal {
	return value.Truncate(int32(precision))
}

func (f *RomanianRomaniaFormatter) FormatDecimalNumber(amount float64) string {
	return FormatPolishDecimal(amount)
}
func (f *RomanianRomaniaFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Polish conventions (comma separators, period decimal, prefix symbol)
	return FormatPolishCurrency(amount, currencySymbol)
}
