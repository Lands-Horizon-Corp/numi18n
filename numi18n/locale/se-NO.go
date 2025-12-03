package locale

import (
	"github.com/shopspring/decimal"
)

// SENOLocale represents the Northern Sami (Norway) locale
var SENOLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Norgga kruvdno",
		Plural:   "Norgga kruvdnot",
		Singular: "Norgga kruvdno",
		Symbol:   "NOK",
		FractionUnit: FractionUnit{
			Name:     "Øre",
			Plural:   "Øret",
			Singular: "Øre",
			Symbol:   "ø",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Norway",
		Currency:       "NOK",
		ISO3166Alpha2:  "NO",
		ISO3166Alpha3:  "NOR",
		ISO3166Numeric: "578",
		Locale:         "se-NO",
		Timezone:       []string{"Europe/Oslo"},
		Language:       "se",
		Emoji:          "🇳🇴",
		PhoneCode:      "+47",
		Domain:         ".no",
	},
	Texts: Texts{
		And:   "ja",
		Minus: "miinus",
		Only:  "dušše",
		Point: "čuokkis",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "tuhat biljonat"},
		{Number: 1000000000000, Value: "biljovdna"},
		{Number: 1000000000, Value: "miljarda"},
		{Number: 1000000, Value: "miljonat"},
		{Number: 100000, Value: "čuođi duhát"},
		{Number: 90000, Value: "ovcci duhát"},
		{Number: 80000, Value: "gávcci duhát"},
		{Number: 70000, Value: "čieža duhát"},
		{Number: 60000, Value: "guhtta duhát"},
		{Number: 50000, Value: "viđa duhát"},
		{Number: 40000, Value: "njealje duhát"},
		{Number: 30000, Value: "golbma duhát"},
		{Number: 20000, Value: "guokte duhát"},
		{Number: 19000, Value: "ovccinuppelohkái duhát"},
		{Number: 18000, Value: "gávccinuppelohkái duhát"},
		{Number: 17000, Value: "čiežanuppelohkái duhát"},
		{Number: 16000, Value: "guhttanuppelohkái duhát"},
		{Number: 15000, Value: "viđanuppelohkái duhát"},
		{Number: 14000, Value: "njeljenuppelohkái duhát"},
		{Number: 13000, Value: "golbmanuppelohkái duhát"},
		{Number: 12000, Value: "guoktenuppelohkái duhát"},
		{Number: 11000, Value: "oktnuppelohkái duhát"},
		{Number: 10000, Value: "logi duhát"},
		{Number: 9000, Value: "ovcci duhát"},
		{Number: 8000, Value: "gávcci duhát"},
		{Number: 7000, Value: "čieža duhát"},
		{Number: 6000, Value: "guhtta duhát"},
		{Number: 5000, Value: "viđa duhát"},
		{Number: 4000, Value: "njealje duhát"},
		{Number: 3000, Value: "golbma duhát"},
		{Number: 2000, Value: "guokte duhát"},
		{Number: 1000, Value: "duhát"},
		{Number: 900, Value: "ovcci čuođi"},
		{Number: 800, Value: "gávcci čuođi"},
		{Number: 700, Value: "čieža čuođi"},
		{Number: 600, Value: "guhtta čuođi"},
		{Number: 500, Value: "viđa čuođi"},
		{Number: 400, Value: "njealje čuođi"},
		{Number: 300, Value: "golbma čuođi"},
		{Number: 200, Value: "guokte čuođi"},
		{Number: 100, Value: "čuođi"},
		{Number: 99, Value: "ovccilogi ja ovcci"},
		{Number: 98, Value: "ovccilogi ja gávcci"},
		{Number: 97, Value: "ovccilogi ja čieža"},
		{Number: 96, Value: "ovccilogi ja guhtta"},
		{Number: 95, Value: "ovccilogi ja viđa"},
		{Number: 94, Value: "ovccilogi ja njealje"},
		{Number: 93, Value: "ovccilogi ja golbma"},
		{Number: 92, Value: "ovccilogi ja guokte"},
		{Number: 91, Value: "ovccilogi ja okta"},
		{Number: 90, Value: "ovccilogi"},
		{Number: 89, Value: "gávccilogi ja ovcci"},
		{Number: 88, Value: "gávccilogi ja gávcci"},
		{Number: 87, Value: "gávccilogi ja čieža"},
		{Number: 86, Value: "gávccilogi ja guhtta"},
		{Number: 85, Value: "gávccilogi ja viđa"},
		{Number: 84, Value: "gávccilogi ja njealje"},
		{Number: 83, Value: "gávccilogi ja golbma"},
		{Number: 82, Value: "gávccilogi ja guokte"},
		{Number: 81, Value: "gávccilogi ja okta"},
		{Number: 80, Value: "gávccilogi"},
		{Number: 79, Value: "čiežalogi ja ovcci"},
		{Number: 78, Value: "čiežalogi ja gávcci"},
		{Number: 77, Value: "čiežalogi ja čieža"},
		{Number: 76, Value: "čiežalogi ja guhtta"},
		{Number: 75, Value: "čiežalogi ja viđa"},
		{Number: 74, Value: "čiežalogi ja njealje"},
		{Number: 73, Value: "čiežalogi ja golbma"},
		{Number: 72, Value: "čiežalogi ja guokte"},
		{Number: 71, Value: "čiežalogi ja okta"},
		{Number: 70, Value: "čiežalogi"},
		{Number: 69, Value: "guhttálogi ja ovcci"},
		{Number: 68, Value: "guhttálogi ja gávcci"},
		{Number: 67, Value: "guhttálogi ja čieža"},
		{Number: 66, Value: "guhttálogi ja guhtta"},
		{Number: 65, Value: "guhttálogi ja viđa"},
		{Number: 64, Value: "guhttálogi ja njealje"},
		{Number: 63, Value: "guhttálogi ja golbma"},
		{Number: 62, Value: "guhttálogi ja guokte"},
		{Number: 61, Value: "guhttálogi ja okta"},
		{Number: 60, Value: "guhttálogi"},
		{Number: 59, Value: "viđalogi ja ovcci"},
		{Number: 58, Value: "viđalogi ja gávcci"},
		{Number: 57, Value: "viđalogi ja čieža"},
		{Number: 56, Value: "viđalogi ja guhtta"},
		{Number: 55, Value: "viđalogi ja viđa"},
		{Number: 54, Value: "viđalogi ja njealje"},
		{Number: 53, Value: "viđalogi ja golbma"},
		{Number: 52, Value: "viđalogi ja guokte"},
		{Number: 51, Value: "viđalogi ja okta"},
		{Number: 50, Value: "viđalogi"},
		{Number: 49, Value: "njeljelogi ja ovcci"},
		{Number: 48, Value: "njeljelogi ja gávcci"},
		{Number: 47, Value: "njeljelogi ja čieža"},
		{Number: 46, Value: "njeljelogi ja guhtta"},
		{Number: 45, Value: "njeljelogi ja viđa"},
		{Number: 44, Value: "njeljelogi ja njealje"},
		{Number: 43, Value: "njeljelogi ja golbma"},
		{Number: 42, Value: "njeljelogi ja guokte"},
		{Number: 41, Value: "njeljelogi ja okta"},
		{Number: 40, Value: "njeljelogi"},
		{Number: 39, Value: "golbmalogi ja ovcci"},
		{Number: 38, Value: "golbmalogi ja gávcci"},
		{Number: 37, Value: "golbmalogi ja čieža"},
		{Number: 36, Value: "golbmalogi ja guhtta"},
		{Number: 35, Value: "golbmalogi ja viđa"},
		{Number: 34, Value: "golbmalogi ja njealje"},
		{Number: 33, Value: "golbmalogi ja golbma"},
		{Number: 32, Value: "golbmalogi ja guokte"},
		{Number: 31, Value: "golbmalogi ja okta"},
		{Number: 30, Value: "golbmalogi"},
		{Number: 29, Value: "guoktelogi ja ovcci"},
		{Number: 28, Value: "guoktelogi ja gávcci"},
		{Number: 27, Value: "guoktelogi ja čieža"},
		{Number: 26, Value: "guoktelogi ja guhtta"},
		{Number: 25, Value: "guoktelogi ja viđa"},
		{Number: 24, Value: "guoktelogi ja njealje"},
		{Number: 23, Value: "guoktelogi ja golbma"},
		{Number: 22, Value: "guoktelogi ja guokte"},
		{Number: 21, Value: "guoktelogi ja okta"},
		{Number: 20, Value: "guoktelogi"},
		{Number: 19, Value: "ovccinuppelohkái"},
		{Number: 18, Value: "gávccinuppelohkái"},
		{Number: 17, Value: "čiežanuppelohkái"},
		{Number: 16, Value: "guhttanuppelohkái"},
		{Number: 15, Value: "viđanuppelohkái"},
		{Number: 14, Value: "njeljenuppelohkái"},
		{Number: 13, Value: "golbmanuppelohkái"},
		{Number: 12, Value: "guoktenuppelohkái"},
		{Number: 11, Value: "oktnuppelohkái"},
		{Number: 10, Value: "logi"},
		{Number: 9, Value: "ovcci"},
		{Number: 8, Value: "gávcci"},
		{Number: 7, Value: "čieža"},
		{Number: 6, Value: "guhtta"},
		{Number: 5, Value: "viđa"},
		{Number: 4, Value: "njealje"},
		{Number: 3, Value: "golbma"},
		{Number: 2, Value: "guokte"},
		{Number: 1, Value: "okta"},
		{Number: 0, Value: "nolla"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Čuođi"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "vuosttaš", Suffix: ".", Masculine: "vuosttaš", Feminine: "vuosttaš", Neuter: "vuosttaš"},
		{Number: 2, Word: "nubbi", Suffix: ".", Masculine: "nubbi", Feminine: "nubbi", Neuter: "nubbi"},
		{Number: 3, Word: "golmmaš", Suffix: ".", Masculine: "golmmaš", Feminine: "golmmaš", Neuter: "golmmaš"},
		{Number: 4, Word: "njeljaš", Suffix: ".", Masculine: "njeljaš", Feminine: "njeljaš", Neuter: "njeljaš"},
		{Number: 5, Word: "viđaš", Suffix: ".", Masculine: "viđaš", Feminine: "viđaš", Neuter: "viđaš"},
		{Number: 6, Word: "guhttaš", Suffix: ".", Masculine: "guhttaš", Feminine: "guhttaš", Neuter: "guhttaš"},
		{Number: 7, Word: "čiežaš", Suffix: ".", Masculine: "čiežaš", Feminine: "čiežaš", Neuter: "čiežaš"},
		{Number: 8, Word: "gávcaš", Suffix: ".", Masculine: "gávcaš", Feminine: "gávcaš", Neuter: "gávcaš"},
		{Number: 9, Word: "ovccaš", Suffix: ".", Masculine: "ovccaš", Feminine: "ovccaš", Neuter: "ovccaš"},
		{Number: 10, Word: "logaš", Suffix: ".", Masculine: "logaš", Feminine: "logaš", Neuter: "logaš"},
		{Number: 11, Word: "oktanuppelohkaš", Suffix: ".", Masculine: "oktanuppelohkaš", Feminine: "oktanuppelohkaš", Neuter: "oktanuppelohkaš"},
		{Number: 12, Word: "guoktenuppelohkaš", Suffix: ".", Masculine: "guoktenuppelohkaš", Feminine: "guoktenuppelohkaš", Neuter: "guoktenuppelohkaš"},
		{Number: 20, Word: "guoktelogaš", Suffix: ".", Masculine: "guoktelogaš", Feminine: "guoktelogaš", Neuter: "guoktelogaš"},
		{Number: 21, Word: "guoktelogi ja vuosttaš", Suffix: ".", Masculine: "guoktelogi ja vuosttaš", Feminine: "guoktelogi ja vuosttaš", Neuter: "guoktelogi ja vuosttaš"},
		{Number: 30, Word: "golbmalogaš", Suffix: ".", Masculine: "golbmalogaš", Feminine: "golbmalogaš", Neuter: "golbmalogaš"},
		{Number: 50, Word: "viđalogaš", Suffix: ".", Masculine: "viđalogaš", Feminine: "viđalogaš", Neuter: "viđalogaš"},
		{Number: 100, Word: "čuođaš", Suffix: ".", Masculine: "čuođaš", Feminine: "čuođaš", Neuter: "čuođaš"},
		{Number: 1000, Word: "duháttaš", Suffix: ".", Masculine: "duháttaš", Feminine: "duháttaš", Neuter: "duháttaš"},
	},
	LocaleFormatter: &NorthernSamiNorwayFormatter{},
}

// NorthernSamiNorwayFormatter handles Northern Sami (Norway) formatting
type NorthernSamiNorwayFormatter struct{}

func (f *NorthernSamiNorwayFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *NorthernSamiNorwayFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *NorthernSamiNorwayFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *NorthernSamiNorwayFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *NorthernSamiNorwayFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *NorthernSamiNorwayFormatter) ChopDecimal(d decimal.Decimal, precision int) decimal.Decimal {
	return d.Truncate(int32(precision))
}

func (f *NorthernSamiNorwayFormatter) FormatDecimalNumber(amount float64) string {
	return FormatPolishDecimal(amount)
}
func (f *NorthernSamiNorwayFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Polish conventions (comma separators, period decimal, prefix symbol)
	return FormatPolishCurrency(amount, currencySymbol)
}
