package locale

import (
	"github.com/shopspring/decimal"
)

// HULocale is a NumI18NLocale configured for Hungary (hu-HU)
var HULocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Forint",
		Plural:   "Forint",
		Singular: "Forint",
		Symbol:   "Ft",
		FractionUnit: FractionUnit{
			Name:     "Fillér",
			Plural:   "Fillér",
			Singular: "Fillér",
			Symbol:   "f",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Hungary",
		Currency:       "HUF",
		ISO3166Alpha2:  "HU",
		ISO3166Alpha3:  "HUN",
		ISO3166Numeric: "348",
		Locale:         "hu-HU",
		Timezone:       []string{"Europe/Budapest"},
		Language:       "hu",
	},
	Texts: Texts{
		And:   "És",
		Minus: "Mínusz",
		Only:  "Csak",
		Point: "Pont",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Egy billiárd"},
		{Number: 1000000000000, Value: "Egy billió"},
		{Number: 1000000000, Value: "Egy milliárd"},
		{Number: 1000000, Value: "Egy millió"},
		{Number: 1000, Value: "Egy ezer"},
		{Number: 100, Value: "Egy száz"},
		{Number: 90, Value: "Kilencven"},
		{Number: 80, Value: "Nyolcvan"},
		{Number: 70, Value: "Hetven"},
		{Number: 60, Value: "Hatvan"},
		{Number: 50, Value: "Ötven"},
		{Number: 40, Value: "Negyven"},
		{Number: 30, Value: "Harminc"},
		{Number: 20, Value: "Húsz"},
		{Number: 19, Value: "Tizenkilenc"},
		{Number: 18, Value: "Tizennyolc"},
		{Number: 17, Value: "Tizenhét"},
		{Number: 16, Value: "Tizenhat"},
		{Number: 15, Value: "Tizenöt"},
		{Number: 14, Value: "Tizennégy"},
		{Number: 13, Value: "Tizenhárom"},
		{Number: 12, Value: "Tizenkettő"},
		{Number: 11, Value: "Tizenegy"},
		{Number: 10, Value: "Tíz"},
		{Number: 9, Value: "Kilenc"},
		{Number: 8, Value: "Nyolc"},
		{Number: 7, Value: "Hét"},
		{Number: 6, Value: "Hat"},
		{Number: 5, Value: "Öt"},
		{Number: 4, Value: "Négy"},
		{Number: 3, Value: "Három"},
		{Number: 2, Value: "Kettő"},
		{Number: 1, Value: "Egy"},
		{Number: 0, Value: "Nulla"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Egy száz"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "első", Suffix: ".", Masculine: "első", Feminine: "első", Neuter: "első"},
		{Number: 2, Word: "második", Suffix: ".", Masculine: "második", Feminine: "második", Neuter: "második"},
		{Number: 3, Word: "harmadik", Suffix: ".", Masculine: "harmadik", Feminine: "harmadik", Neuter: "harmadik"},
		{Number: 4, Word: "negyedik", Suffix: ".", Masculine: "negyedik", Feminine: "negyedik", Neuter: "negyedik"},
		{Number: 5, Word: "ötödik", Suffix: ".", Masculine: "ötödik", Feminine: "ötödik", Neuter: "ötödik"},
		{Number: 6, Word: "hatodik", Suffix: ".", Masculine: "hatodik", Feminine: "hatodik", Neuter: "hatodik"},
		{Number: 7, Word: "hetedik", Suffix: ".", Masculine: "hetedik", Feminine: "hetedik", Neuter: "hetedik"},
		{Number: 8, Word: "nyolcadik", Suffix: ".", Masculine: "nyolcadik", Feminine: "nyolcadik", Neuter: "nyolcadik"},
		{Number: 9, Word: "kilencedik", Suffix: ".", Masculine: "kilencedik", Feminine: "kilencedik", Neuter: "kilencedik"},
		{Number: 10, Word: "tizedik", Suffix: ".", Masculine: "tizedik", Feminine: "tizedik", Neuter: "tizedik"},
		{Number: 11, Word: "tizenegyedik", Suffix: ".", Masculine: "tizenegyedik", Feminine: "tizenegyedik", Neuter: "tizenegyedik"},
		{Number: 12, Word: "tizenkettedik", Suffix: ".", Masculine: "tizenkettedik", Feminine: "tizenkettedik", Neuter: "tizenkettedik"},
		{Number: 13, Word: "tizenharmadik", Suffix: ".", Masculine: "tizenharmadik", Feminine: "tizenharmadik", Neuter: "tizenharmadik"},
		{Number: 14, Word: "tizennegyedik", Suffix: ".", Masculine: "tizennegyedik", Feminine: "tizennegyedik", Neuter: "tizennegyedik"},
		{Number: 15, Word: "tizenötödik", Suffix: ".", Masculine: "tizenötödik", Feminine: "tizenötödik", Neuter: "tizenötödik"},
		{Number: 16, Word: "tizenhatodik", Suffix: ".", Masculine: "tizenhatodik", Feminine: "tizenhatodik", Neuter: "tizenhatodik"},
		{Number: 17, Word: "tizenhetedik", Suffix: ".", Masculine: "tizenhetedik", Feminine: "tizenhetedik", Neuter: "tizenhetedik"},
		{Number: 18, Word: "tizennyolcadik", Suffix: ".", Masculine: "tizennyolcadik", Feminine: "tizennyolcadik", Neuter: "tizennyolcadik"},
		{Number: 19, Word: "tizenkilencedik", Suffix: ".", Masculine: "tizenkilencedik", Feminine: "tizenkilencedik", Neuter: "tizenkilencedik"},
		{Number: 20, Word: "huszadik", Suffix: ".", Masculine: "huszadik", Feminine: "huszadik", Neuter: "huszadik"},
		{Number: 21, Word: "huszonegyedik", Suffix: ".", Masculine: "huszonegyedik", Feminine: "huszonegyedik", Neuter: "huszonegyedik"},
		{Number: 30, Word: "harmincadik", Suffix: ".", Masculine: "harmincadik", Feminine: "harmincadik", Neuter: "harmincadik"},
		{Number: 40, Word: "negyvenedik", Suffix: ".", Masculine: "negyvenedik", Feminine: "negyvenedik", Neuter: "negyvenedik"},
		{Number: 50, Word: "ötvenedik", Suffix: ".", Masculine: "ötvenedik", Feminine: "ötvenedik", Neuter: "ötvenedik"},
		{Number: 60, Word: "hatvanadik", Suffix: ".", Masculine: "hatvanadik", Feminine: "hatvanadik", Neuter: "hatvanadik"},
		{Number: 70, Word: "hetvenedik", Suffix: ".", Masculine: "hetvenedik", Feminine: "hetvenedik", Neuter: "hetvenedik"},
		{Number: 80, Word: "nyolcvanadik", Suffix: ".", Masculine: "nyolcvanadik", Feminine: "nyolcvanadik", Neuter: "nyolcvanadik"},
		{Number: 90, Word: "kilencvenedik", Suffix: ".", Masculine: "kilencvenedik", Feminine: "kilencvenedik", Neuter: "kilencvenedik"},
		{Number: 100, Word: "századik", Suffix: ".", Masculine: "századik", Feminine: "századik", Neuter: "századik"},
		{Number: 1000, Word: "ezredik", Suffix: ".", Masculine: "ezredik", Feminine: "ezredik", Neuter: "ezredik"},
	},
	LocaleFormatter: &HungarianFormatter{},
}

// HungarianFormatter handles Hungarian formatting
type HungarianFormatter struct{}

func (f *HungarianFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *HungarianFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *HungarianFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *HungarianFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *HungarianFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *HungarianFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	return amount.Truncate(int32(precision))
}

func (f *HungarianFormatter) FormatDecimalNumber(amount float64) string {
	return FormatEuropeanDecimal(amount)
}
func (f *HungarianFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with European conventions (period separators, comma decimal, prefix symbol)
	return FormatEuropeanCurrency(amount, currencySymbol)
}
