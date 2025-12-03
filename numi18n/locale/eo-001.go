package locale

import (
	"github.com/shopspring/decimal"
)

// EOLocale is a NumI18NLocale configured for Esperanto (eo-001)
var EOLocale = NumI18NLocale{
	LocaleFormatter: &EsperantoFormatter{},
	Currency: Currency{
		Name:     "Euro",
		Plural:   "Eŭroj",
		Singular: "Eŭro",
		Symbol:   "€",
		FractionUnit: FractionUnit{
			Name:     "Cento",
			Plural:   "Centoj",
			Singular: "Cento",
			Symbol:   "¢",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "World",
		Currency:       "USD",
		ISO3166Alpha2:  "001",
		ISO3166Alpha3:  "WLD",
		ISO3166Numeric: "001",
		Locale:         "eo-001",
		Timezone:       []string{"UTC"},
		Language:       "eo",
		Emoji:          "🌍",
		PhoneCode:      "+1",
		Domain:         ".com",
	},
	Texts: Texts{
		And:   "Kaj",
		Minus: "Minus",
		Only:  "Nur",
		Point: "Punkto",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Kvadriliono"},
		{Number: 1000000000000, Value: "Triliono"},
		{Number: 1000000000, Value: "Bilio"},
		{Number: 1000000, Value: "Miliono"},
		{Number: 1000, Value: "Mil"},
		{Number: 100, Value: "Cent"},
		{Number: 90, Value: "Naŭdek"},
		{Number: 80, Value: "Okdek"},
		{Number: 70, Value: "Sepdek"},
		{Number: 60, Value: "Sesdek"},
		{Number: 50, Value: "Kvindek"},
		{Number: 40, Value: "Kvardek"},
		{Number: 30, Value: "Tridek"},
		{Number: 20, Value: "Dudek"},
		{Number: 19, Value: "Dek naŭ"},
		{Number: 18, Value: "Dek ok"},
		{Number: 17, Value: "Dek sep"},
		{Number: 16, Value: "Dek ses"},
		{Number: 15, Value: "Dek kvin"},
		{Number: 14, Value: "Dek kvar"},
		{Number: 13, Value: "Dek tri"},
		{Number: 12, Value: "Dek du"},
		{Number: 11, Value: "Dek unu"},
		{Number: 10, Value: "Dek"},
		{Number: 9, Value: "Naŭ"},
		{Number: 8, Value: "Ok"},
		{Number: 7, Value: "Sep"},
		{Number: 6, Value: "Ses"},
		{Number: 5, Value: "Kvin"},
		{Number: 4, Value: "Kvar"},
		{Number: 3, Value: "Tri"},
		{Number: 2, Value: "Du"},
		{Number: 1, Value: "Unu"},
		{Number: 0, Value: "Nulo"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Unu Cent"},
		{Number: 1000, Value: "Unu Mil"},
		{Number: 1000000, Value: "Unu Miliono"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "Unua", Suffix: "-a"},
		{Number: 2, Word: "Dua", Suffix: "-a"},
		{Number: 3, Word: "Tria", Suffix: "-a"},
		{Number: 4, Word: "Kvara", Suffix: "-a"},
		{Number: 5, Word: "Kvina", Suffix: "-a"},
		{Number: 6, Word: "Sesa", Suffix: "-a"},
		{Number: 7, Word: "Sepa", Suffix: "-a"},
		{Number: 8, Word: "Oka", Suffix: "-a"},
		{Number: 9, Word: "Naŭa", Suffix: "-a"},
		{Number: 10, Word: "Deka", Suffix: "-a"},
		{Number: 11, Word: "Dek-unua", Suffix: "-a"},
		{Number: 12, Word: "Dek-dua", Suffix: "-a"},
		{Number: 13, Word: "Dek-tria", Suffix: "-a"},
		{Number: 14, Word: "Dek-kvara", Suffix: "-a"},
		{Number: 15, Word: "Dek-kvina", Suffix: "-a"},
		{Number: 16, Word: "Dek-sesa", Suffix: "-a"},
		{Number: 17, Word: "Dek-sepa", Suffix: "-a"},
		{Number: 18, Word: "Dek-oka", Suffix: "-a"},
		{Number: 19, Word: "Dek-naŭa", Suffix: "-a"},
		{Number: 20, Word: "Dudeka", Suffix: "-a"},
		{Number: 21, Word: "Dudek-unua", Suffix: "-a"},
		{Number: 22, Word: "Dudek-dua", Suffix: "-a"},
		{Number: 23, Word: "Dudek-tria", Suffix: "-a"},
		{Number: 30, Word: "Trideka", Suffix: "-a"},
		{Number: 40, Word: "Kvardeka", Suffix: "-a"},
		{Number: 50, Word: "Kvindeka", Suffix: "-a"},
		{Number: 60, Word: "Sesdeka", Suffix: "-a"},
		{Number: 70, Word: "Sepdeka", Suffix: "-a"},
		{Number: 80, Word: "Okdeka", Suffix: "-a"},
		{Number: 90, Word: "Naŭdeka", Suffix: "-a"},
		{Number: 100, Word: "Centa", Suffix: "-a"},
		{Number: 1000, Word: "Mila", Suffix: "-a"},
		{Number: 1000000, Word: "Miliona", Suffix: "-a"},
	},
}

// EsperantoFormatter handles Esperanto formatting
type EsperantoFormatter struct{}

func (f *EsperantoFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *EsperantoFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *EsperantoFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *EsperantoFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *EsperantoFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *EsperantoFormatter) ChopDecimal(d decimal.Decimal, precision int) decimal.Decimal {
	return d.Truncate(int32(precision))
}

func (f *EsperantoFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *EsperantoFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
