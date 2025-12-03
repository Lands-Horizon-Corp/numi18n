package locale

import (
	"github.com/shopspring/decimal"
)

// FOLocale is a NumI18NLocale configured for Faroe Islands (fo-FO)
var FOLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Króna",
		Plural:   "Krónur",
		Singular: "Króna",
		Symbol:   "kr",
		FractionUnit: FractionUnit{
			Name:     "Oyra",
			Plural:   "Oyrur",
			Singular: "Oyra",
			Symbol:   "ø",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Faroe Islands",
		Currency:       "DKK",
		ISO3166Alpha2:  "FO",
		ISO3166Alpha3:  "FRO",
		ISO3166Numeric: "234",
		Locale:         "fo-FO",
		Timezone:       []string{"Atlantic/Faroe"},
		Language:       "fo",
		Emoji:          "🇫🇴",
	},
	Texts: Texts{
		And:   "Og",
		Minus: "Mínus",
		Only:  "Bara",
		Point: "Komma",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Eitt quadrillion"},
		{Number: 1000000000000, Value: "Eitt trillion"},
		{Number: 1000000000, Value: "Ein milliard"},
		{Number: 1000000, Value: "Ein million"},
		{Number: 1000, Value: "Eitt túsund"},
		{Number: 100, Value: "Eitt hundrað"},
		{Number: 90, Value: "Níti"},
		{Number: 80, Value: "Áttati"},
		{Number: 70, Value: "Sjeyti"},
		{Number: 60, Value: "Seksti"},
		{Number: 50, Value: "Hálvtríss"},
		{Number: 40, Value: "Fyrti"},
		{Number: 30, Value: "Tríati"},
		{Number: 20, Value: "Tuttugu"},
		{Number: 19, Value: "Nítjan"},
		{Number: 18, Value: "Átjan"},
		{Number: 17, Value: "Seytan"},
		{Number: 16, Value: "Sekstan"},
		{Number: 15, Value: "Fimtan"},
		{Number: 14, Value: "Fjórtan"},
		{Number: 13, Value: "Trettan"},
		{Number: 12, Value: "Tólv"},
		{Number: 11, Value: "Ellivu"},
		{Number: 10, Value: "Tíggju"},
		{Number: 9, Value: "Níggju"},
		{Number: 8, Value: "Átta"},
		{Number: 7, Value: "Sjey"},
		{Number: 6, Value: "Seks"},
		{Number: 5, Value: "Fimm"},
		{Number: 4, Value: "Fýra"},
		{Number: 3, Value: "Trý"},
		{Number: 2, Value: "Tvey"},
		{Number: 1, Value: "Eitt"},
		{Number: 0, Value: "Null"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Eitt hundrað"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "Fyrsti", Suffix: ".", Masculine: "Fyrsti", Feminine: "Fyrsta", Neuter: "Fyrsta"},
		{Number: 2, Word: "Annar", Suffix: ".", Masculine: "Annar", Feminine: "Onnur", Neuter: "Annað"},
		{Number: 3, Word: "Triðji", Suffix: ".", Masculine: "Triðji", Feminine: "Triðja", Neuter: "Triðja"},
		{Number: 4, Word: "Fjórði", Suffix: ".", Masculine: "Fjórði", Feminine: "Fjórða", Neuter: "Fjórða"},
		{Number: 5, Word: "Fimti", Suffix: ".", Masculine: "Fimti", Feminine: "Fimta", Neuter: "Fimta"},
		{Number: 6, Word: "Sætti", Suffix: ".", Masculine: "Sætti", Feminine: "Sætta", Neuter: "Sætta"},
		{Number: 7, Word: "Sjeyndi", Suffix: ".", Masculine: "Sjeyndi", Feminine: "Sjeynda", Neuter: "Sjeynda"},
		{Number: 8, Word: "Áttundi", Suffix: ".", Masculine: "Áttundi", Feminine: "Áttunda", Neuter: "Áttunda"},
		{Number: 9, Word: "Níundi", Suffix: ".", Masculine: "Níundi", Feminine: "Níunda", Neuter: "Níunda"},
		{Number: 10, Word: "Tíundi", Suffix: ".", Masculine: "Tíundi", Feminine: "Tíunda", Neuter: "Tíunda"},
		{Number: 11, Word: "Ellivti", Suffix: ".", Masculine: "Ellivti", Feminine: "Ellivta", Neuter: "Ellivta"},
		{Number: 12, Word: "Tólvti", Suffix: ".", Masculine: "Tólvti", Feminine: "Tólvta", Neuter: "Tólvta"},
		{Number: 13, Word: "Trettandi", Suffix: ".", Masculine: "Trettandi", Feminine: "Trettanda", Neuter: "Trettanda"},
		{Number: 14, Word: "Fjórtandi", Suffix: ".", Masculine: "Fjórtandi", Feminine: "Fjórtanda", Neuter: "Fjórtanda"},
		{Number: 15, Word: "Fimtandi", Suffix: ".", Masculine: "Fimtandi", Feminine: "Fimtanda", Neuter: "Fimtanda"},
		{Number: 16, Word: "Sekstandi", Suffix: ".", Masculine: "Sekstandi", Feminine: "Sekstanda", Neuter: "Sekstanda"},
		{Number: 17, Word: "Seytandi", Suffix: ".", Masculine: "Seytandi", Feminine: "Seytanda", Neuter: "Seytanda"},
		{Number: 18, Word: "Átjandi", Suffix: ".", Masculine: "Átjandi", Feminine: "Átjanda", Neuter: "Átjanda"},
		{Number: 19, Word: "Nítjandi", Suffix: ".", Masculine: "Nítjandi", Feminine: "Nítjanda", Neuter: "Nítjanda"},
		{Number: 20, Word: "Tuttugundi", Suffix: ".", Masculine: "Tuttugundi", Feminine: "Tuttugunda", Neuter: "Tuttugunda"},
		{Number: 21, Word: "Tuttugu og fyrsti", Suffix: ".", Masculine: "Tuttugu og fyrsti", Feminine: "Tuttugu og fyrsta", Neuter: "Tuttugu og fyrsta"},
		{Number: 30, Word: "Tríatiundi", Suffix: ".", Masculine: "Tríatiundi", Feminine: "Tríatiunda", Neuter: "Tríatiunda"},
		{Number: 40, Word: "Fyrtiundi", Suffix: ".", Masculine: "Fyrtiundi", Feminine: "Fyrtiunda", Neuter: "Fyrtiunda"},
		{Number: 50, Word: "Hálvtríssundi", Suffix: ".", Masculine: "Hálvtríssundi", Feminine: "Hálvtríssunda", Neuter: "Hálvtríssunda"},
		{Number: 60, Word: "Sekstiundi", Suffix: ".", Masculine: "Sekstiundi", Feminine: "Sekstiunda", Neuter: "Sekstiunda"},
		{Number: 70, Word: "Sjeytiundi", Suffix: ".", Masculine: "Sjeytiundi", Feminine: "Sjeytiunda", Neuter: "Sjeytiunda"},
		{Number: 80, Word: "Áttatiundi", Suffix: ".", Masculine: "Áttatiundi", Feminine: "Áttatiunda", Neuter: "Áttatiunda"},
		{Number: 90, Word: "Nítiundi", Suffix: ".", Masculine: "Nítiundi", Feminine: "Nítiunda", Neuter: "Nítiunda"},
		{Number: 100, Word: "Hundrađi", Suffix: ".", Masculine: "Hundrađi", Feminine: "Hundrađa", Neuter: "Hundrađa"},
		{Number: 1000, Word: "Túsundti", Suffix: ".", Masculine: "Túsundti", Feminine: "Túsundta", Neuter: "Túsundta"},
	},
	LocaleFormatter: &FaroeseFormatter{},
}

// FaroeseFormatter handles Faroese formatting
type FaroeseFormatter struct{}

func (f *FaroeseFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *FaroeseFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *FaroeseFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *FaroeseFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *FaroeseFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *FaroeseFormatter) ChopDecimal(value decimal.Decimal, precision int) decimal.Decimal {
	return value.Round(int32(precision))
}

func (f *FaroeseFormatter) FormatDecimalNumber(amount float64) string {
	return FormatPolishDecimal(amount)
}
func (f *FaroeseFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Polish conventions (comma separators, period decimal, prefix symbol)
	return FormatPolishCurrency(amount, currencySymbol)
}
