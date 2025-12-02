package locale

import (
	"strings"
	"github.com/shopspring/decimal"
)

// ESEULocale is a NumI18NLocale configured for Basque - Spain (eu-ES)
var ESEULocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Euro",
		Plural:   "Euroak",
		Singular: "Euro",
		Symbol:   "€",
		FractionUnit: FractionUnit{
			Name:     "Zentimo",
			Plural:   "Zentimoak",
			Singular: "Zentimo",
			Symbol:   "¢",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Spain",
		Currency:       "EUR",
		ISO3166Alpha2:  "ES",
		ISO3166Alpha3:  "ESP",
		ISO3166Numeric: "724",
		Locale:         "eu-ES",
		Timezone:       []string{"Europe/Madrid"},
		Language:       "eu",
	},
	Texts: Texts{
		And:   "Eta",
		Minus: "Minus",
		Only:  "Soilik",
		Point: "Puntua",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Kualdrilioi"},
		{Number: 1000000000000, Value: "Trilioi"},
		{Number: 1000000000, Value: "Mila milioi"},
		{Number: 1000000, Value: "Milioi"},
		{Number: 1000, Value: "Mila"},
		{Number: 100, Value: "Ehun"},
		{Number: 90, Value: "Laurogeita Hamar"},
		{Number: 80, Value: "Laurogei"},
		{Number: 70, Value: "Hirurogeita Hamar"},
		{Number: 60, Value: "Hirurogei"},
		{Number: 50, Value: "Berrogeita Hamar"},
		{Number: 40, Value: "Berrogei"},
		{Number: 30, Value: "Hogeita Hamar"},
		{Number: 29, Value: "Hogeita Bederatzi"},
		{Number: 28, Value: "Hogeita Zortzi"},
		{Number: 27, Value: "Hogeita Zazpi"},
		{Number: 26, Value: "Hogeita Sei"},
		{Number: 25, Value: "Hogeita Bost"},
		{Number: 24, Value: "Hogeita Lau"},
		{Number: 23, Value: "Hogeita Hiru"},
		{Number: 22, Value: "Hogeita Bi"},
		{Number: 21, Value: "Hogeita Bat"},
		{Number: 20, Value: "Hogei"},
		{Number: 19, Value: "Hemeretzi"},
		{Number: 18, Value: "Hemezortzi"},
		{Number: 17, Value: "Hamazazpi"},
		{Number: 16, Value: "Hamasei"},
		{Number: 15, Value: "Hamabost"},
		{Number: 14, Value: "Hamalau"},
		{Number: 13, Value: "Hamahiru"},
		{Number: 12, Value: "Hamabi"},
		{Number: 11, Value: "Hamaika"},
		{Number: 10, Value: "Hamar"},
		{Number: 9, Value: "Bederatzi"},
		{Number: 8, Value: "Zortzi"},
		{Number: 7, Value: "Zazpi"},
		{Number: 6, Value: "Sei"},
		{Number: 5, Value: "Bost"},
		{Number: 4, Value: "Lau"},
		{Number: 3, Value: "Hiru"},
		{Number: 2, Value: "Bi"},
		{Number: 1, Value: "Bat"},
		{Number: 0, Value: "Zero"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Ehun"},
		{Number: 1000, Value: "Mila"},
		{Number: 1000000, Value: "Milioi Bat"},
		{Number: 1000000000, Value: "Mila Milioi"},
		{Number: 47, Value: "Berrogeita Zazpi"},
		{Number: 56, Value: "Berrogeita Hamasei"},
		{Number: 99, Value: "Laurogeita Hemeretzi"},
		{Number: 75, Value: "Hirurogeita Hamabost"},
		{Number: 256, Value: "Bi Ehun Berrogeita Hamasei"},
		{Number: 34, Value: "Hogeita Hamalau"},
		{Number: 123, Value: "Ehun Hogei Hiru"},
		{Number: 101, Value: "Ehun Bat"},
		{Number: 1001, Value: "Mila Bat"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "Lehenengoa", Suffix: ".", Masculine: "Lehenengoa", Feminine: "Lehenengoa", Neuter: "Lehenengoa"},
		{Number: 2, Word: "Bigarren", Suffix: ".", Masculine: "Bigarren", Feminine: "Bigarren", Neuter: "Bigarren"},
		{Number: 3, Word: "Hirugarren", Suffix: ".", Masculine: "Hirugarren", Feminine: "Hirugarren", Neuter: "Hirugarren"},
		{Number: 4, Word: "Laugarren", Suffix: ".", Masculine: "Laugarren", Feminine: "Laugarren", Neuter: "Laugarren"},
		{Number: 5, Word: "Bosgarren", Suffix: ".", Masculine: "Bosgarren", Feminine: "Bosgarren", Neuter: "Bosgarren"},
		{Number: 6, Word: "Seigarren", Suffix: ".", Masculine: "Seigarren", Feminine: "Seigarren", Neuter: "Seigarren"},
		{Number: 7, Word: "Zazpigarren", Suffix: ".", Masculine: "Zazpigarren", Feminine: "Zazpigarren", Neuter: "Zazpigarren"},
		{Number: 8, Word: "Zortzigarren", Suffix: ".", Masculine: "Zortzigarren", Feminine: "Zortzigarren", Neuter: "Zortzigarren"},
		{Number: 9, Word: "Bederatzigarren", Suffix: ".", Masculine: "Bederatzigarren", Feminine: "Bederatzigarren", Neuter: "Bederatzigarren"},
		{Number: 10, Word: "Hamargarren", Suffix: ".", Masculine: "Hamargarren", Feminine: "Hamargarren", Neuter: "Hamargarren"},
		{Number: 11, Word: "Hamaikgarren", Suffix: ".", Masculine: "Hamaikgarren", Feminine: "Hamaikgarren", Neuter: "Hamaikgarren"},
		{Number: 12, Word: "Hamabigarren", Suffix: ".", Masculine: "Hamabigarren", Feminine: "Hamabigarren", Neuter: "Hamabigarren"},
		{Number: 13, Word: "Hamahirugarren", Suffix: ".", Masculine: "Hamahirugarren", Feminine: "Hamahirugarren", Neuter: "Hamahirugarren"},
		{Number: 14, Word: "Hamalaugarren", Suffix: ".", Masculine: "Hamalaugarren", Feminine: "Hamalaugarren", Neuter: "Hamalaugarren"},
		{Number: 15, Word: "Hamabostgarren", Suffix: ".", Masculine: "Hamabostgarren", Feminine: "Hamabostgarren", Neuter: "Hamabostgarren"},
		{Number: 16, Word: "Hamaseigarren", Suffix: ".", Masculine: "Hamaseigarren", Feminine: "Hamaseigarren", Neuter: "Hamaseigarren"},
		{Number: 17, Word: "Hamazazpigarren", Suffix: ".", Masculine: "Hamazazpigarren", Feminine: "Hamazazpigarren", Neuter: "Hamazazpigarren"},
		{Number: 18, Word: "Hemezortzigarren", Suffix: ".", Masculine: "Hemezortzigarren", Feminine: "Hemezortzigarren", Neuter: "Hemezortzigarren"},
		{Number: 19, Word: "Hemeretzigarren", Suffix: ".", Masculine: "Hemeretzigarren", Feminine: "Hemeretzigarren", Neuter: "Hemeretzigarren"},
		{Number: 20, Word: "Hogeigarren", Suffix: ".", Masculine: "Hogeigarren", Feminine: "Hogeigarren", Neuter: "Hogeigarren"},
		{Number: 21, Word: "Hogeita batgarren", Suffix: ".", Masculine: "Hogeita batgarren", Feminine: "Hogeita batgarren", Neuter: "Hogeita batgarren"},
		{Number: 30, Word: "Hogeita hamargarren", Suffix: ".", Masculine: "Hogeita hamargarren", Feminine: "Hogeita hamargarren", Neuter: "Hogeita hamargarren"},
		{Number: 40, Word: "Berrogeigarren", Suffix: ".", Masculine: "Berrogeigarren", Feminine: "Berrogeigarren", Neuter: "Berrogeigarren"},
		{Number: 50, Word: "Berrogeita hamargarren", Suffix: ".", Masculine: "Berrogeita hamargarren", Feminine: "Berrogeita hamargarren", Neuter: "Berrogeita hamargarren"},
		{Number: 60, Word: "Hirurogeigarren", Suffix: ".", Masculine: "Hirurogeigarren", Feminine: "Hirurogeigarren", Neuter: "Hirurogeigarren"},
		{Number: 70, Word: "Hirurogeita hamargarren", Suffix: ".", Masculine: "Hirurogeita hamargarren", Feminine: "Hirurogeita hamargarren", Neuter: "Hirurogeita hamargarren"},
		{Number: 80, Word: "Laurogeigarren", Suffix: ".", Masculine: "Laurogeigarren", Feminine: "Laurogeigarren", Neuter: "Laurogeigarren"},
		{Number: 90, Word: "Laurogeita hamargarren", Suffix: ".", Masculine: "Laurogeita hamargarren", Feminine: "Laurogeita hamargarren", Neuter: "Laurogeita hamargarren"},
		{Number: 100, Word: "Ehungarren", Suffix: ".", Masculine: "Ehungarren", Feminine: "Ehungarren", Neuter: "Ehungarren"},
		{Number: 1000, Word: "Milagarren", Suffix: ".", Masculine: "Milagarren", Feminine: "Milagarren", Neuter: "Milagarren"},
	},
	LocaleFormatter: &BasqueFormatter{},
}

// BasqueFormatter handles Basque formatting
type BasqueFormatter struct{}

func (f *BasqueFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *BasqueFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *BasqueFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *BasqueFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *BasqueFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *BasqueFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	return amount.Truncate(int32(precision))
}


func (f *BasqueFormatter) FormatDecimalNumber(amount float64) string {
	return DefaultFormatDecimalNumber(amount, ",", ".")
}
func (f *BasqueFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	formattedNumber := f.FormatDecimalNumber(amount)
	
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}
	
	// Default currency placement for this locale (prefix with symbol)
	if strings.HasPrefix(formattedNumber, "-") {
		formattedNumber = strings.TrimPrefix(formattedNumber, "-")
		return "-" + currencySymbol + formattedNumber
	}
	
	return currencySymbol + formattedNumber
}
