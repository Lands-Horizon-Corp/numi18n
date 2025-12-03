package locale

import (
	"github.com/shopspring/decimal"
)

// BSBALocale is a NumI18NLocale configured for Bosnian (Bosnia and Herzegovina) - bs-BA
var BSBALocale = NumI18NLocale{
	LocaleFormatter: &BosnianFormatter{},
	Currency: Currency{
		Name:     "Konvertibilna marka",
		Plural:   "Konvertibilne marke",
		Singular: "Konvertibilna marka",
		Symbol:   "KM",
		FractionUnit: FractionUnit{
			Name:     "Fening",
			Plural:   "Feninga",
			Singular: "Fening",
			Symbol:   "f",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Bosnia and Herzegovina",
		Currency:       "BAM",
		ISO3166Alpha2:  "BA",
		ISO3166Alpha3:  "BIH",
		ISO3166Numeric: "070",
		Locale:         "bs-BA",
		Timezone:       []string{"Europe/Sarajevo"},
		Language:       "bs",
	},
	Texts: Texts{
		And:   "i",
		Minus: "minus",
		Only:  "samo",
		Point: "tačka",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Kvadrilion"},
		{Number: 1000000000000, Value: "Trilion"},
		{Number: 1000000000, Value: "Milijarda"},
		{Number: 1000000, Value: "Milion"},
		{Number: 1000, Value: "Hiljada"},
		{Number: 100, Value: "Sto"},
		{Number: 90, Value: "Devedeset"},
		{Number: 80, Value: "Osamdeset"},
		{Number: 70, Value: "Sedamdeset"},
		{Number: 60, Value: "Šezdeset"},
		{Number: 50, Value: "Pedeset"},
		{Number: 40, Value: "Četrdeset"},
		{Number: 30, Value: "Trideset"},
		{Number: 20, Value: "Dvadeset"},
		{Number: 19, Value: "Devetnaest"},
		{Number: 18, Value: "Osamnaest"},
		{Number: 17, Value: "Sedamnaest"},
		{Number: 16, Value: "Šesnaest"},
		{Number: 15, Value: "Petnaest"},
		{Number: 14, Value: "Četrnaest"},
		{Number: 13, Value: "Trinaest"},
		{Number: 12, Value: "Dvanaest"},
		{Number: 11, Value: "Jedanaest"},
		{Number: 10, Value: "Deset"},
		{Number: 9, Value: "Devet"},
		{Number: 8, Value: "Osam"},
		{Number: 7, Value: "Sedam"},
		{Number: 6, Value: "Šest"},
		{Number: 5, Value: "Pet"},
		{Number: 4, Value: "Četiri"},
		{Number: 3, Value: "Tri"},
		{Number: 2, Value: "Dva"},
		{Number: 1, Value: "Jedan"},
		{Number: 0, Value: "Nula"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Jedan Sto"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "prvi", Suffix: "", Masculine: "prvi", Feminine: "prva", Neuter: "prvo"},
		{Number: 2, Word: "drugi", Suffix: "", Masculine: "drugi", Feminine: "druga", Neuter: "drugo"},
		{Number: 3, Word: "treći", Suffix: "", Masculine: "treći", Feminine: "treća", Neuter: "treće"},
		{Number: 4, Word: "četvrti", Suffix: "", Masculine: "četvrti", Feminine: "četvrta", Neuter: "četvrto"},
		{Number: 5, Word: "peti", Suffix: "", Masculine: "peti", Feminine: "peta", Neuter: "peto"},
		{Number: 6, Word: "šesti", Suffix: "", Masculine: "šesti", Feminine: "šesta", Neuter: "šesto"},
		{Number: 7, Word: "sedmi", Suffix: "", Masculine: "sedmi", Feminine: "sedma", Neuter: "sedmo"},
		{Number: 8, Word: "osmi", Suffix: "", Masculine: "osmi", Feminine: "osma", Neuter: "osmo"},
		{Number: 9, Word: "deveti", Suffix: "", Masculine: "deveti", Feminine: "deveta", Neuter: "deveto"},
		{Number: 10, Word: "deseti", Suffix: "", Masculine: "deseti", Feminine: "deseta", Neuter: "deseto"},
		{Number: 11, Word: "jedanaesti", Suffix: "", Masculine: "jedanaesti", Feminine: "jedanaesta", Neuter: "jedanaesto"},
		{Number: 12, Word: "dvanaesti", Suffix: "", Masculine: "dvanaesti", Feminine: "dvanaesta", Neuter: "dvanaesto"},
		{Number: 13, Word: "trinaesti", Suffix: "", Masculine: "trinaesti", Feminine: "trinaesta", Neuter: "trinaesto"},
		{Number: 14, Word: "četrnaesti", Suffix: "", Masculine: "četrnaesti", Feminine: "četrnaesta", Neuter: "četrnaesto"},
		{Number: 15, Word: "petnaesti", Suffix: "", Masculine: "petnaesti", Feminine: "petnaesta", Neuter: "petnaesto"},
		{Number: 16, Word: "šesnaesti", Suffix: "", Masculine: "šesnaesti", Feminine: "šesnaesta", Neuter: "šesnaesto"},
		{Number: 17, Word: "sedamnaesti", Suffix: "", Masculine: "sedamnaesti", Feminine: "sedamnaesta", Neuter: "sedamnaesto"},
		{Number: 18, Word: "osamnaesti", Suffix: "", Masculine: "osamnaesti", Feminine: "osamnaesta", Neuter: "osamnaesto"},
		{Number: 19, Word: "devetnaesti", Suffix: "", Masculine: "devetnaesti", Feminine: "devetnaesta", Neuter: "devetnaesto"},
		{Number: 20, Word: "dvadeseti", Suffix: "", Masculine: "dvadeseti", Feminine: "dvadeseta", Neuter: "dvadeseto"},
		{Number: 21, Word: "dvadeset prvi", Suffix: "", Masculine: "dvadeset prvi", Feminine: "dvadeset prva", Neuter: "dvadeset prvo"},
		{Number: 30, Word: "trideseti", Suffix: "", Masculine: "trideseti", Feminine: "trideseta", Neuter: "trideseto"},
		{Number: 40, Word: "četrdeseti", Suffix: "", Masculine: "četrdeseti", Feminine: "četrdeseta", Neuter: "četrdeseto"},
		{Number: 50, Word: "pedeseti", Suffix: "", Masculine: "pedeseti", Feminine: "pedeseta", Neuter: "pedeseto"},
		{Number: 60, Word: "šezdeseti", Suffix: "", Masculine: "šezdeseti", Feminine: "šezdeseta", Neuter: "šezdeseto"},
		{Number: 70, Word: "sedamdeseti", Suffix: "", Masculine: "sedamdeseti", Feminine: "sedamdeseta", Neuter: "sedamdeseto"},
		{Number: 80, Word: "osamdeseti", Suffix: "", Masculine: "osamdeseti", Feminine: "osamdeseta", Neuter: "osamdeseto"},
		{Number: 90, Word: "devedeseti", Suffix: "", Masculine: "devedeseti", Feminine: "devedeseta", Neuter: "devedeseto"},
		{Number: 100, Word: "stoti", Suffix: "", Masculine: "stoti", Feminine: "stota", Neuter: "stoto"},
		{Number: 1000, Word: "hiljaditi", Suffix: "", Masculine: "hiljaditi", Feminine: "hiljadita", Neuter: "hiljadito"},
	},
}

// BosnianFormatter handles Bosnian formatting
type BosnianFormatter struct{}

func (f *BosnianFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *BosnianFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *BosnianFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *BosnianFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *BosnianFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *BosnianFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	return amount.Round(int32(precision))
}

func (f *BosnianFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *BosnianFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
