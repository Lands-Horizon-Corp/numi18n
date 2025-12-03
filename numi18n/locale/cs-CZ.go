package locale

import (
	"github.com/shopspring/decimal"
)

// CSCZLocale is a NumI18NLocale configured for Czech (Czech Republic) - cs-CZ
var CSCZLocale = NumI18NLocale{
	LocaleFormatter: &CzechFormatter{},
	Currency: Currency{
		Name:     "Koruna",
		Plural:   "Koruny",
		Singular: "Koruna",
		Symbol:   "Kč",
		FractionUnit: FractionUnit{
			Name:     "Haléř",
			Plural:   "Haléře",
			Singular: "Haléř",
			Symbol:   "h",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Czech Republic",
		Currency:       "CZK",
		ISO3166Alpha2:  "CZ",
		ISO3166Alpha3:  "CZE",
		ISO3166Numeric: "203",
		Locale:         "cs-CZ",
		Timezone:       []string{"Europe/Prague"},
		Language:       "cs",
		Emoji:          "🇨🇿",
	},
	Texts: Texts{
		And:   "a",
		Minus: "mínus",
		Only:  "pouze",
		Point: "tečka",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Kvadrilion"},
		{Number: 1000000000000, Value: "Bilion"},
		{Number: 1000000000, Value: "Miliarda"},
		{Number: 1000000, Value: "Milion"},
		{Number: 1000, Value: "Tisíc"},
		{Number: 100, Value: "Sto"},
		{Number: 90, Value: "Devadesát"},
		{Number: 80, Value: "Osmdesát"},
		{Number: 70, Value: "Sedmdesát"},
		{Number: 60, Value: "Šedesát"},
		{Number: 50, Value: "Padesát"},
		{Number: 40, Value: "Čtyřicet"},
		{Number: 30, Value: "Třicet"},
		{Number: 20, Value: "Dvacet"},
		{Number: 19, Value: "Devatenáct"},
		{Number: 18, Value: "Osmnáct"},
		{Number: 17, Value: "Sedmnáct"},
		{Number: 16, Value: "Šestnáct"},
		{Number: 15, Value: "Patnáct"},
		{Number: 14, Value: "Čtrnáct"},
		{Number: 13, Value: "Třináct"},
		{Number: 12, Value: "Dvanáct"},
		{Number: 11, Value: "Jedenáct"},
		{Number: 10, Value: "Deset"},
		{Number: 9, Value: "Devět"},
		{Number: 8, Value: "Osm"},
		{Number: 7, Value: "Sedm"},
		{Number: 6, Value: "Šest"},
		{Number: 5, Value: "Pět"},
		{Number: 4, Value: "Čtyři"},
		{Number: 3, Value: "Tři"},
		{Number: 2, Value: "Dva"},
		{Number: 1, Value: "Jedna"},
		{Number: 0, Value: "Nula"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Sto"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "první", Suffix: ".", Masculine: "první", Feminine: "první", Neuter: "první"},
		{Number: 2, Word: "druhý", Suffix: ".", Masculine: "druhý", Feminine: "druhá", Neuter: "druhé"},
		{Number: 3, Word: "třetí", Suffix: ".", Masculine: "třetí", Feminine: "třetí", Neuter: "třetí"},
		{Number: 4, Word: "čtvrtý", Suffix: ".", Masculine: "čtvrtý", Feminine: "čtvrtá", Neuter: "čtvrté"},
		{Number: 5, Word: "pátý", Suffix: ".", Masculine: "pátý", Feminine: "pátá", Neuter: "páté"},
		{Number: 6, Word: "šestý", Suffix: ".", Masculine: "šestý", Feminine: "šestá", Neuter: "šesté"},
		{Number: 7, Word: "sedmý", Suffix: ".", Masculine: "sedmý", Feminine: "sedmá", Neuter: "sedmé"},
		{Number: 8, Word: "osmý", Suffix: ".", Masculine: "osmý", Feminine: "osmá", Neuter: "osmé"},
		{Number: 9, Word: "devátý", Suffix: ".", Masculine: "devátý", Feminine: "devátá", Neuter: "deváté"},
		{Number: 10, Word: "desátý", Suffix: ".", Masculine: "desátý", Feminine: "desátá", Neuter: "desáté"},
		{Number: 11, Word: "jedenáctý", Suffix: ".", Masculine: "jedenáctý", Feminine: "jedenáctá", Neuter: "jedenácté"},
		{Number: 12, Word: "dvanáctý", Suffix: ".", Masculine: "dvanáctý", Feminine: "dvanáctá", Neuter: "dvanácté"},
		{Number: 13, Word: "třináctý", Suffix: ".", Masculine: "třináctý", Feminine: "třináctá", Neuter: "třinácté"},
		{Number: 14, Word: "čtrnáctý", Suffix: ".", Masculine: "čtrnáctý", Feminine: "čtrnáctá", Neuter: "čtrnácté"},
		{Number: 15, Word: "patnáctý", Suffix: ".", Masculine: "patnáctý", Feminine: "patnáctá", Neuter: "patnácté"},
		{Number: 16, Word: "šestnáctý", Suffix: ".", Masculine: "šestnáctý", Feminine: "šestnáctá", Neuter: "šestnácté"},
		{Number: 17, Word: "sedmnáctý", Suffix: ".", Masculine: "sedmnáctý", Feminine: "sedmnáctá", Neuter: "sedmnácté"},
		{Number: 18, Word: "osmnáctý", Suffix: ".", Masculine: "osmnáctý", Feminine: "osmnáctá", Neuter: "osmnácté"},
		{Number: 19, Word: "devatenáctý", Suffix: ".", Masculine: "devatenáctý", Feminine: "devatenáctá", Neuter: "devatenácté"},
		{Number: 20, Word: "dvacátý", Suffix: ".", Masculine: "dvacátý", Feminine: "dvacátá", Neuter: "dvacáté"},
		{Number: 21, Word: "dvacátý první", Suffix: ".", Masculine: "dvacátý první", Feminine: "dvacátá první", Neuter: "dvacáté první"},
		{Number: 30, Word: "třicátý", Suffix: ".", Masculine: "třicátý", Feminine: "třicátá", Neuter: "třicáté"},
		{Number: 40, Word: "čtyřicátý", Suffix: ".", Masculine: "čtyřicátý", Feminine: "čtyřicátá", Neuter: "čtyřicáté"},
		{Number: 50, Word: "padesátý", Suffix: ".", Masculine: "padesátý", Feminine: "padesátá", Neuter: "padesáté"},
		{Number: 60, Word: "šedesátý", Suffix: ".", Masculine: "šedesátý", Feminine: "šedesátá", Neuter: "šedesáté"},
		{Number: 70, Word: "sedmdesátý", Suffix: ".", Masculine: "sedmdesátý", Feminine: "sedmdesátá", Neuter: "sedmdesáté"},
		{Number: 80, Word: "osmdesátý", Suffix: ".", Masculine: "osmdesátý", Feminine: "osmdesátá", Neuter: "osmdesáté"},
		{Number: 90, Word: "devadesátý", Suffix: ".", Masculine: "devadesátý", Feminine: "devadesátá", Neuter: "devadesáté"},
		{Number: 100, Word: "stý", Suffix: ".", Masculine: "stý", Feminine: "stá", Neuter: "sté"},
		{Number: 1000, Word: "tisící", Suffix: ".", Masculine: "tisící", Feminine: "tisící", Neuter: "tisící"},
	},
}

// CzechFormatter handles Czech formatting
type CzechFormatter struct{}

func (f *CzechFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *CzechFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *CzechFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *CzechFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *CzechFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *CzechFormatter) ChopDecimal(value decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 2
	}
	return value.Truncate(int32(precision))
}

func (f *CzechFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *CzechFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
