package locale

import (
	"github.com/shopspring/decimal"
)

// MAFRLocale is a NumI18NLocale configured for Morocco (fr-MA)
var MAFRLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Dirham marocain",
		Plural:   "Dirhams marocains",
		Singular: "Dirham marocain",
		Symbol:   "MAD",
		FractionUnit: FractionUnit{
			Name:     "Centime",
			Plural:   "Centimes",
			Singular: "Centime",
			Symbol:   "¢",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Morocco",
		Currency:       "MAD",
		ISO3166Alpha2:  "MA",
		ISO3166Alpha3:  "MAR",
		ISO3166Numeric: "504",
		Locale:         "fr-MA",
		Timezone:       []string{"Africa/Casablanca"},
		Language:       "fr",
		Emoji:          "🇲🇦",
	},
	Texts: Texts{
		And:   "Et",
		Minus: "Moins",
		Only:  "Seulement",
		Point: "Virgule",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Un billiard"},
		{Number: 1000000000000, Value: "Un billion"},
		{Number: 1000000000, Value: "Un milliard"},
		{Number: 1000000, Value: "Un million"},
		{Number: 1000, Value: "Mille"},
		{Number: 100, Value: "Cent"},
		{Number: 90, Value: "Quatre-vingt-dix"},
		{Number: 80, Value: "Quatre-vingts"},
		{Number: 70, Value: "Soixante-dix"},
		{Number: 60, Value: "Soixante"},
		{Number: 50, Value: "Cinquante"},
		{Number: 40, Value: "Quarante"},
		{Number: 30, Value: "Trente"},
		{Number: 20, Value: "Vingt"},
		{Number: 19, Value: "Dix-neuf"},
		{Number: 18, Value: "Dix-huit"},
		{Number: 17, Value: "Dix-sept"},
		{Number: 16, Value: "Seize"},
		{Number: 15, Value: "Quinze"},
		{Number: 14, Value: "Quatorze"},
		{Number: 13, Value: "Treize"},
		{Number: 12, Value: "Douze"},
		{Number: 11, Value: "Onze"},
		{Number: 10, Value: "Dix"},
		{Number: 9, Value: "Neuf"},
		{Number: 8, Value: "Huit"},
		{Number: 7, Value: "Sept"},
		{Number: 6, Value: "Six"},
		{Number: 5, Value: "Cinq"},
		{Number: 4, Value: "Quatre"},
		{Number: 3, Value: "Trois"},
		{Number: 2, Value: "Deux"},
		{Number: 1, Value: "Un"},
		{Number: 0, Value: "Zéro"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Cent"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "Premier", Suffix: "er", Masculine: "Premier", Feminine: "Première", Neuter: "Premier"},
		{Number: 2, Word: "Deuxième", Suffix: "e", Masculine: "Deuxième", Feminine: "Deuxième", Neuter: "Deuxième"},
		{Number: 3, Word: "Troisième", Suffix: "e", Masculine: "Troisième", Feminine: "Troisième", Neuter: "Troisième"},
		{Number: 4, Word: "Quatrième", Suffix: "e", Masculine: "Quatrième", Feminine: "Quatrième", Neuter: "Quatrième"},
		{Number: 5, Word: "Cinquième", Suffix: "e", Masculine: "Cinquième", Feminine: "Cinquième", Neuter: "Cinquième"},
		{Number: 6, Word: "Sixième", Suffix: "e", Masculine: "Sixième", Feminine: "Sixième", Neuter: "Sixième"},
		{Number: 7, Word: "Septième", Suffix: "e", Masculine: "Septième", Feminine: "Septième", Neuter: "Septième"},
		{Number: 8, Word: "Huitième", Suffix: "e", Masculine: "Huitième", Feminine: "Huitième", Neuter: "Huitième"},
		{Number: 9, Word: "Neuvième", Suffix: "e", Masculine: "Neuvième", Feminine: "Neuvième", Neuter: "Neuvième"},
		{Number: 10, Word: "Dixième", Suffix: "e", Masculine: "Dixième", Feminine: "Dixième", Neuter: "Dixième"},
		{Number: 11, Word: "Onzième", Suffix: "e", Masculine: "Onzième", Feminine: "Onzième", Neuter: "Onzième"},
		{Number: 12, Word: "Douzième", Suffix: "e", Masculine: "Douzième", Feminine: "Douzième", Neuter: "Douzième"},
		{Number: 13, Word: "Treizième", Suffix: "e", Masculine: "Treizième", Feminine: "Treizième", Neuter: "Treizième"},
		{Number: 14, Word: "Quatorzième", Suffix: "e", Masculine: "Quatorzième", Feminine: "Quatorzième", Neuter: "Quatorzième"},
		{Number: 15, Word: "Quinzième", Suffix: "e", Masculine: "Quinzième", Feminine: "Quinzième", Neuter: "Quinzième"},
		{Number: 16, Word: "Seizième", Suffix: "e", Masculine: "Seizième", Feminine: "Seizième", Neuter: "Seizième"},
		{Number: 17, Word: "Dix-septième", Suffix: "e", Masculine: "Dix-septième", Feminine: "Dix-septième", Neuter: "Dix-septième"},
		{Number: 18, Word: "Dix-huitième", Suffix: "e", Masculine: "Dix-huitième", Feminine: "Dix-huitième", Neuter: "Dix-huitième"},
		{Number: 19, Word: "Dix-neuvième", Suffix: "e", Masculine: "Dix-neuvième", Feminine: "Dix-neuvième", Neuter: "Dix-neuvième"},
		{Number: 20, Word: "Vingtième", Suffix: "e", Masculine: "Vingtième", Feminine: "Vingtième", Neuter: "Vingtième"},
		{Number: 21, Word: "Vingt et unième", Suffix: "e", Masculine: "Vingt et unième", Feminine: "Vingt et unième", Neuter: "Vingt et unième"},
		{Number: 30, Word: "Trentième", Suffix: "e", Masculine: "Trentième", Feminine: "Trentième", Neuter: "Trentième"},
		{Number: 40, Word: "Quarantième", Suffix: "e", Masculine: "Quarantième", Feminine: "Quarantième", Neuter: "Quarantième"},
		{Number: 50, Word: "Cinquantième", Suffix: "e", Masculine: "Cinquantième", Feminine: "Cinquantième", Neuter: "Cinquantième"},
		{Number: 60, Word: "Soixantième", Suffix: "e", Masculine: "Soixantième", Feminine: "Soixantième", Neuter: "Soixantième"},
		{Number: 70, Word: "Soixante-dixième", Suffix: "e", Masculine: "Soixante-dixième", Feminine: "Soixante-dixième", Neuter: "Soixante-dixième"},
		{Number: 80, Word: "Quatre-vingtième", Suffix: "e", Masculine: "Quatre-vingtième", Feminine: "Quatre-vingtième", Neuter: "Quatre-vingtième"},
		{Number: 90, Word: "Quatre-vingt-dixième", Suffix: "e", Masculine: "Quatre-vingt-dixième", Feminine: "Quatre-vingt-dixième", Neuter: "Quatre-vingt-dixième"},
		{Number: 100, Word: "Centième", Suffix: "e", Masculine: "Centième", Feminine: "Centième", Neuter: "Centième"},
		{Number: 1000, Word: "Millième", Suffix: "e", Masculine: "Millième", Feminine: "Millième", Neuter: "Millième"},
	},
	LocaleFormatter: &FrenchMoroccoFormatter{},
}

// FrenchMoroccoFormatter handles French Morocco formatting
type FrenchMoroccoFormatter struct{}

func (f *FrenchMoroccoFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *FrenchMoroccoFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *FrenchMoroccoFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *FrenchMoroccoFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *FrenchMoroccoFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *FrenchMoroccoFormatter) ChopDecimal(value decimal.Decimal, precision int) decimal.Decimal {
	return value.Round(int32(precision))
}

func (f *FrenchMoroccoFormatter) FormatDecimalNumber(amount float64) string {
	return FormatFrenchDecimal(amount)
}

func (f *FrenchMoroccoFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with French conventions (space separators, comma decimal, prefix symbol)
	return FormatFrenchCurrency(amount, currencySymbol)
}
