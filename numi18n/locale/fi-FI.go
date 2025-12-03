package locale

import (
	"github.com/shopspring/decimal"
)

// FILocale is a NumI18NLocale configured for Finland (fi-FI)
var FILocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Euro",
		Plural:   "Euroa",
		Singular: "Euro",
		Symbol:   "€",
		FractionUnit: FractionUnit{
			Name:     "Sentti",
			Plural:   "Senttiä",
			Singular: "Sentti",
			Symbol:   "¢",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Finland",
		Currency:       "EUR",
		ISO3166Alpha2:  "FI",
		ISO3166Alpha3:  "FIN",
		ISO3166Numeric: "246",
		Locale:         "fi-FI",
		Timezone:       []string{"Europe/Helsinki"},
		Language:       "fi",
		Emoji:          "🇫🇮",
	},
	Texts: Texts{
		And:   "Ja",
		Minus: "Miinus",
		Only:  "Vain",
		Point: "Pilkku",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Kvadriljoona"},
		{Number: 1000000000000, Value: "Triljoona"},
		{Number: 1000000000, Value: "Miljardi"},
		{Number: 1000000, Value: "Miljoona"},
		{Number: 1000, Value: "Tuhat"},
		{Number: 100, Value: "Sata"},
		{Number: 90, Value: "Yhdeksänkymmentä"},
		{Number: 80, Value: "Kahdeksankymmentä"},
		{Number: 70, Value: "Seitsemänkymmentä"},
		{Number: 60, Value: "Kuusikymmentä"},
		{Number: 50, Value: "Viisikymmentä"},
		{Number: 40, Value: "Neljäkymmentä"},
		{Number: 30, Value: "Kolmekymmentä"},
		{Number: 20, Value: "Kaksikymmentä"},
		{Number: 19, Value: "Yhdeksäntoista"},
		{Number: 18, Value: "Kahdeksantoista"},
		{Number: 17, Value: "Seitsemäntoista"},
		{Number: 16, Value: "Kuusitoista"},
		{Number: 15, Value: "Viisitoista"},
		{Number: 14, Value: "Neljätoista"},
		{Number: 13, Value: "Kolmetoista"},
		{Number: 12, Value: "Kaksitoista"},
		{Number: 11, Value: "Yksitoista"},
		{Number: 10, Value: "Kymmenen"},
		{Number: 9, Value: "Yhdeksän"},
		{Number: 8, Value: "Kahdeksan"},
		{Number: 7, Value: "Seitsemän"},
		{Number: 6, Value: "Kuusi"},
		{Number: 5, Value: "Viisi"},
		{Number: 4, Value: "Neljä"},
		{Number: 3, Value: "Kolme"},
		{Number: 2, Value: "Kaksi"},
		{Number: 1, Value: "Yksi"},
		{Number: 0, Value: "Nolla"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Sata"},
		{Number: 1000, Value: "Tuhat"},
		{Number: 1000000, Value: "Miljoona"},
		{Number: 1000000000, Value: "Miljardi"},
		{Number: 1000000000000, Value: "Triljoona"},
		{Number: 1000000000000000, Value: "Kvadriljoona"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "Ensimmäinen", Suffix: ".", Masculine: "Ensimmäinen", Feminine: "Ensimmäinen", Neuter: "Ensimmäinen"},
		{Number: 2, Word: "Toinen", Suffix: ".", Masculine: "Toinen", Feminine: "Toinen", Neuter: "Toinen"},
		{Number: 3, Word: "Kolmas", Suffix: ".", Masculine: "Kolmas", Feminine: "Kolmas", Neuter: "Kolmas"},
		{Number: 4, Word: "Neljäs", Suffix: ".", Masculine: "Neljäs", Feminine: "Neljäs", Neuter: "Neljäs"},
		{Number: 5, Word: "Viides", Suffix: ".", Masculine: "Viides", Feminine: "Viides", Neuter: "Viides"},
		{Number: 6, Word: "Kuudes", Suffix: ".", Masculine: "Kuudes", Feminine: "Kuudes", Neuter: "Kuudes"},
		{Number: 7, Word: "Seitsemäs", Suffix: ".", Masculine: "Seitsemäs", Feminine: "Seitsemäs", Neuter: "Seitsemäs"},
		{Number: 8, Word: "Kahdeksas", Suffix: ".", Masculine: "Kahdeksas", Feminine: "Kahdeksas", Neuter: "Kahdeksas"},
		{Number: 9, Word: "Yhdeksäs", Suffix: ".", Masculine: "Yhdeksäs", Feminine: "Yhdeksäs", Neuter: "Yhdeksäs"},
		{Number: 10, Word: "Kymmenes", Suffix: ".", Masculine: "Kymmenes", Feminine: "Kymmenes", Neuter: "Kymmenes"},
		{Number: 11, Word: "Yhdestoista", Suffix: ".", Masculine: "Yhdestoista", Feminine: "Yhdestoista", Neuter: "Yhdestoista"},
		{Number: 12, Word: "Kahdestoista", Suffix: ".", Masculine: "Kahdestoista", Feminine: "Kahdestoista", Neuter: "Kahdestoista"},
		{Number: 13, Word: "Kolmastoista", Suffix: ".", Masculine: "Kolmastoista", Feminine: "Kolmastoista", Neuter: "Kolmastoista"},
		{Number: 14, Word: "Neljästoista", Suffix: ".", Masculine: "Neljästoista", Feminine: "Neljästoista", Neuter: "Neljästoista"},
		{Number: 15, Word: "Viidestoista", Suffix: ".", Masculine: "Viidestoista", Feminine: "Viidestoista", Neuter: "Viidestoista"},
		{Number: 16, Word: "Kuudestoista", Suffix: ".", Masculine: "Kuudestoista", Feminine: "Kuudestoista", Neuter: "Kuudestoista"},
		{Number: 17, Word: "Seitsemästoista", Suffix: ".", Masculine: "Seitsemästoista", Feminine: "Seitsemästoista", Neuter: "Seitsemästoista"},
		{Number: 18, Word: "Kahdeksastoista", Suffix: ".", Masculine: "Kahdeksastoista", Feminine: "Kahdeksastoista", Neuter: "Kahdeksastoista"},
		{Number: 19, Word: "Yhdeksästoista", Suffix: ".", Masculine: "Yhdeksästoista", Feminine: "Yhdeksästoista", Neuter: "Yhdeksästoista"},
		{Number: 20, Word: "Kahdeskymmenes", Suffix: ".", Masculine: "Kahdeskymmenes", Feminine: "Kahdeskymmenes", Neuter: "Kahdeskymmenes"},
		{Number: 21, Word: "Kahdeskymmenesensimmäinen", Suffix: ".", Masculine: "Kahdeskymmenesensimmäinen", Feminine: "Kahdeskymmenesensimmäinen", Neuter: "Kahdeskymmenesensimmäinen"},
		{Number: 30, Word: "Kolmaskymmenes", Suffix: ".", Masculine: "Kolmaskymmenes", Feminine: "Kolmaskymmenes", Neuter: "Kolmaskymmenes"},
		{Number: 40, Word: "Neljäskymmenes", Suffix: ".", Masculine: "Neljäskymmenes", Feminine: "Neljäskymmenes", Neuter: "Neljäskymmenes"},
		{Number: 50, Word: "Viideskymmenes", Suffix: ".", Masculine: "Viideskymmenes", Feminine: "Viideskymmenes", Neuter: "Viideskymmenes"},
		{Number: 60, Word: "Kuudeskymmenes", Suffix: ".", Masculine: "Kuudeskymmenes", Feminine: "Kuudeskymmenes", Neuter: "Kuudeskymmenes"},
		{Number: 70, Word: "Seitsemäskymmenes", Suffix: ".", Masculine: "Seitsemäskymmenes", Feminine: "Seitsemäskymmenes", Neuter: "Seitsemäskymmenes"},
		{Number: 80, Word: "Kahdeksaskymmenes", Suffix: ".", Masculine: "Kahdeksaskymmenes", Feminine: "Kahdeksaskymmenes", Neuter: "Kahdeksaskymmenes"},
		{Number: 90, Word: "Yhdeksäskymmenes", Suffix: ".", Masculine: "Yhdeksäskymmenes", Feminine: "Yhdeksäskymmenes", Neuter: "Yhdeksäskymmenes"},
		{Number: 100, Word: "Sadas", Suffix: ".", Masculine: "Sadas", Feminine: "Sadas", Neuter: "Sadas"},
		{Number: 1000, Word: "Tuhannes", Suffix: ".", Masculine: "Tuhannes", Feminine: "Tuhannes", Neuter: "Tuhannes"},
	},
	LocaleFormatter: &FinnishFormatter{},
}

// FinnishFormatter handles Finnish formatting
type FinnishFormatter struct{}

func (f *FinnishFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *FinnishFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *FinnishFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *FinnishFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *FinnishFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *FinnishFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	return amount.Round(int32(precision))
}

func (f *FinnishFormatter) FormatDecimalNumber(amount float64) string {
	return FormatEuropeanDecimal(amount)
}

func (f *FinnishFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with European conventions (period separators, comma decimal, prefix symbol)
	return FormatEuropeanCurrency(amount, currencySymbol)
}
