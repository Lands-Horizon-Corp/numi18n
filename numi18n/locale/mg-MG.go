package locale

import (
	"github.com/shopspring/decimal"
)

// MGMGLocale represents the Malagasy (Madagascar) locale
var MGMGLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Ariary",
		Plural:   "Ariary",
		Singular: "Ariary",
		Symbol:   "Ar",
		FractionUnit: FractionUnit{
			Name:     "Iraimbilanja",
			Plural:   "Iraimbilanja",
			Singular: "Iraimbilanja",
			Symbol:   "¢",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Madagascar",
		Currency:       "MGA",
		ISO3166Alpha2:  "MG",
		ISO3166Alpha3:  "MDG",
		ISO3166Numeric: "450",
		Locale:         "mg-MG",
		Timezone:       []string{"Indian/Antananarivo"},
		Language:       "mg",
	},
	Texts: Texts{
		And:   "ary",
		Minus: "miiba",
		Only:  "ihany",
		Point: "teboka",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "arivo tapitrisa"},
		{Number: 1000000000000, Value: "tapitrisa"},
		{Number: 1000000000, Value: "arivo tapitrisa kely"},
		{Number: 1000000, Value: "tapitrisa kely"},
		{Number: 100000, Value: "zato arivo"},
		{Number: 90000, Value: "sivy polo arivo"},
		{Number: 80000, Value: "valo polo arivo"},
		{Number: 70000, Value: "fito polo arivo"},
		{Number: 60000, Value: "enina polo arivo"},
		{Number: 50000, Value: "dimy polo arivo"},
		{Number: 40000, Value: "efatra polo arivo"},
		{Number: 30000, Value: "telo polo arivo"},
		{Number: 20000, Value: "roa polo arivo"},
		{Number: 19000, Value: "sivy ambin'ny folo arivo"},
		{Number: 18000, Value: "valo ambin'ny folo arivo"},
		{Number: 17000, Value: "fito ambin'ny folo arivo"},
		{Number: 16000, Value: "enina ambin'ny folo arivo"},
		{Number: 15000, Value: "dimy ambin'ny folo arivo"},
		{Number: 14000, Value: "efatra ambin'ny folo arivo"},
		{Number: 13000, Value: "telo ambin'ny folo arivo"},
		{Number: 12000, Value: "roa ambin'ny folo arivo"},
		{Number: 11000, Value: "iraika ambin'ny folo arivo"},
		{Number: 10000, Value: "folo arivo"},
		{Number: 9000, Value: "sivy arivo"},
		{Number: 8000, Value: "valo arivo"},
		{Number: 7000, Value: "fito arivo"},
		{Number: 6000, Value: "enina arivo"},
		{Number: 5000, Value: "dimy arivo"},
		{Number: 4000, Value: "efatra arivo"},
		{Number: 3000, Value: "telo arivo"},
		{Number: 2000, Value: "roa arivo"},
		{Number: 1000, Value: "arivo"},
		{Number: 900, Value: "sivy zato"},
		{Number: 800, Value: "valo zato"},
		{Number: 700, Value: "fito zato"},
		{Number: 600, Value: "enina zato"},
		{Number: 500, Value: "dimy zato"},
		{Number: 400, Value: "efatra zato"},
		{Number: 300, Value: "telo zato"},
		{Number: 200, Value: "roa zato"},
		{Number: 100, Value: "zato"},
		{Number: 99, Value: "sivy ambin'ny sivifolo"},
		{Number: 98, Value: "valo ambin'ny sivifolo"},
		{Number: 97, Value: "fito ambin'ny sivifolo"},
		{Number: 96, Value: "enina ambin'ny sivifolo"},
		{Number: 95, Value: "dimy ambin'ny sivifolo"},
		{Number: 94, Value: "efatra ambin'ny sivifolo"},
		{Number: 93, Value: "telo ambin'ny sivifolo"},
		{Number: 92, Value: "roa ambin'ny sivifolo"},
		{Number: 91, Value: "iraika ambin'ny sivifolo"},
		{Number: 90, Value: "sivifolo"},
		{Number: 89, Value: "sivy ambin'ny valopolo"},
		{Number: 88, Value: "valo ambin'ny valopolo"},
		{Number: 87, Value: "fito ambin'ny valopolo"},
		{Number: 86, Value: "enina ambin'ny valopolo"},
		{Number: 85, Value: "dimy ambin'ny valopolo"},
		{Number: 84, Value: "efatra ambin'ny valopolo"},
		{Number: 83, Value: "telo ambin'ny valopolo"},
		{Number: 82, Value: "roa ambin'ny valopolo"},
		{Number: 81, Value: "iraika ambin'ny valopolo"},
		{Number: 80, Value: "valopolo"},
		{Number: 79, Value: "sivy ambin'ny fitopolo"},
		{Number: 78, Value: "valo ambin'ny fitopolo"},
		{Number: 77, Value: "fito ambin'ny fitopolo"},
		{Number: 76, Value: "enina ambin'ny fitopolo"},
		{Number: 75, Value: "dimy ambin'ny fitopolo"},
		{Number: 74, Value: "efatra ambin'ny fitopolo"},
		{Number: 73, Value: "telo ambin'ny fitopolo"},
		{Number: 72, Value: "roa ambin'ny fitopolo"},
		{Number: 71, Value: "iraika ambin'ny fitopolo"},
		{Number: 70, Value: "fitopolo"},
		{Number: 69, Value: "sivy ambin'ny enimpolo"},
		{Number: 68, Value: "valo ambin'ny enimpolo"},
		{Number: 67, Value: "fito ambin'ny enimpolo"},
		{Number: 66, Value: "enina ambin'ny enimpolo"},
		{Number: 65, Value: "dimy ambin'ny enimpolo"},
		{Number: 64, Value: "efatra ambin'ny enimpolo"},
		{Number: 63, Value: "telo ambin'ny enimpolo"},
		{Number: 62, Value: "roa ambin'ny enimpolo"},
		{Number: 61, Value: "iraika ambin'ny enimpolo"},
		{Number: 60, Value: "enimpolo"},
		{Number: 59, Value: "sivy ambin'ny dimampolo"},
		{Number: 58, Value: "valo ambin'ny dimampolo"},
		{Number: 57, Value: "fito ambin'ny dimampolo"},
		{Number: 56, Value: "enina ambin'ny dimampolo"},
		{Number: 55, Value: "dimy ambin'ny dimampolo"},
		{Number: 54, Value: "efatra ambin'ny dimampolo"},
		{Number: 53, Value: "telo ambin'ny dimampolo"},
		{Number: 52, Value: "roa ambin'ny dimampolo"},
		{Number: 51, Value: "iraika ambin'ny dimampolo"},
		{Number: 50, Value: "dimampolo"},
		{Number: 49, Value: "sivy ambin'ny efatrapolo"},
		{Number: 48, Value: "valo ambin'ny efatrapolo"},
		{Number: 47, Value: "fito ambin'ny efatrapolo"},
		{Number: 46, Value: "enina ambin'ny efatrapolo"},
		{Number: 45, Value: "dimy ambin'ny efatrapolo"},
		{Number: 44, Value: "efatra ambin'ny efatrapolo"},
		{Number: 43, Value: "telo ambin'ny efatrapolo"},
		{Number: 42, Value: "roa ambin'ny efatrapolo"},
		{Number: 41, Value: "iraika ambin'ny efatrapolo"},
		{Number: 40, Value: "efatrapolo"},
		{Number: 39, Value: "sivy ambin'ny telopolo"},
		{Number: 38, Value: "valo ambin'ny telopolo"},
		{Number: 37, Value: "fito ambin'ny telopolo"},
		{Number: 36, Value: "enina ambin'ny telopolo"},
		{Number: 35, Value: "dimy ambin'ny telopolo"},
		{Number: 34, Value: "efatra ambin'ny telopolo"},
		{Number: 33, Value: "telo ambin'ny telopolo"},
		{Number: 32, Value: "roa ambin'ny telopolo"},
		{Number: 31, Value: "iraika ambin'ny telopolo"},
		{Number: 30, Value: "telopolo"},
		{Number: 29, Value: "sivy ambin'ny roapolo"},
		{Number: 28, Value: "valo ambin'ny roapolo"},
		{Number: 27, Value: "fito ambin'ny roapolo"},
		{Number: 26, Value: "enina ambin'ny roapolo"},
		{Number: 25, Value: "dimy ambin'ny roapolo"},
		{Number: 24, Value: "efatra ambin'ny roapolo"},
		{Number: 23, Value: "telo ambin'ny roapolo"},
		{Number: 22, Value: "roa ambin'ny roapolo"},
		{Number: 21, Value: "iraika ambin'ny roapolo"},
		{Number: 20, Value: "roapolo"},
		{Number: 19, Value: "sivy ambin'ny folo"},
		{Number: 18, Value: "valo ambin'ny folo"},
		{Number: 17, Value: "fito ambin'ny folo"},
		{Number: 16, Value: "enina ambin'ny folo"},
		{Number: 15, Value: "dimy ambin'ny folo"},
		{Number: 14, Value: "efatra ambin'ny folo"},
		{Number: 13, Value: "telo ambin'ny folo"},
		{Number: 12, Value: "roa ambin'ny folo"},
		{Number: 11, Value: "iraika ambin'ny folo"},
		{Number: 10, Value: "folo"},
		{Number: 9, Value: "sivy"},
		{Number: 8, Value: "valo"},
		{Number: 7, Value: "fito"},
		{Number: 6, Value: "enina"},
		{Number: 5, Value: "dimy"},
		{Number: 4, Value: "efatra"},
		{Number: 3, Value: "telo"},
		{Number: 2, Value: "roa"},
		{Number: 1, Value: "iraika"},
		{Number: 0, Value: "aotra"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Zato iray"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "voalohany", Suffix: "-ny", Masculine: "voalohany", Feminine: "voalohany", Neuter: "voalohany"},
		{Number: 2, Word: "faharoa", Suffix: "-ny", Masculine: "faharoa", Feminine: "faharoa", Neuter: "faharoa"},
		{Number: 3, Word: "fahatelo", Suffix: "-ny", Masculine: "fahatelo", Feminine: "fahatelo", Neuter: "fahatelo"},
		{Number: 4, Word: "fahefatra", Suffix: "-ny", Masculine: "fahefatra", Feminine: "fahefatra", Neuter: "fahefatra"},
		{Number: 5, Word: "fahadimy", Suffix: "-ny", Masculine: "fahadimy", Feminine: "fahadimy", Neuter: "fahadimy"},
		{Number: 6, Word: "fahenina", Suffix: "-ny", Masculine: "fahenina", Feminine: "fahenina", Neuter: "fahenina"},
		{Number: 7, Word: "fahafito", Suffix: "-ny", Masculine: "fahafito", Feminine: "fahafito", Neuter: "fahafito"},
		{Number: 8, Word: "fahavalo", Suffix: "-ny", Masculine: "fahavalo", Feminine: "fahavalo", Neuter: "fahavalo"},
		{Number: 9, Word: "fahasivy", Suffix: "-ny", Masculine: "fahasivy", Feminine: "fahasivy", Neuter: "fahasivy"},
		{Number: 10, Word: "fahafolo", Suffix: "-ny", Masculine: "fahafolo", Feminine: "fahafolo", Neuter: "fahafolo"},
		{Number: 11, Word: "faha-iraika ambin'ny folo", Suffix: "-ny", Masculine: "faha-iraika ambin'ny folo", Feminine: "faha-iraika ambin'ny folo", Neuter: "faha-iraika ambin'ny folo"},
		{Number: 12, Word: "faha-roa ambin'ny folo", Suffix: "-ny", Masculine: "faha-roa ambin'ny folo", Feminine: "faha-roa ambin'ny folo", Neuter: "faha-roa ambin'ny folo"},
		{Number: 20, Word: "faha-roapolo", Suffix: "-ny", Masculine: "faha-roapolo", Feminine: "faha-roapolo", Neuter: "faha-roapolo"},
		{Number: 21, Word: "faha-iraika ambin'ny roapolo", Suffix: "-ny", Masculine: "faha-iraika ambin'ny roapolo", Feminine: "faha-iraika ambin'ny roapolo", Neuter: "faha-iraika ambin'ny roapolo"},
		{Number: 30, Word: "faha-telopolo", Suffix: "-ny", Masculine: "faha-telopolo", Feminine: "faha-telopolo", Neuter: "faha-telopolo"},
		{Number: 100, Word: "faha-zato", Suffix: "-ny", Masculine: "faha-zato", Feminine: "faha-zato", Neuter: "faha-zato"},
		{Number: 1000, Word: "faha-arivo", Suffix: "-ny", Masculine: "faha-arivo", Feminine: "faha-arivo", Neuter: "faha-arivo"},
	},
	LocaleFormatter: &MalagasyFormatter{},
}

// MalagasyFormatter handles Malagasy-specific formatting
type MalagasyFormatter struct{}

func (f *MalagasyFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *MalagasyFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *MalagasyFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *MalagasyFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *MalagasyFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *MalagasyFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 0
	}
	return amount.Truncate(int32(precision))
}

func (f *MalagasyFormatter) FormatDecimalNumber(amount float64) string {
	return FormatFrenchDecimal(amount)
}
func (f *MalagasyFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with French conventions (space separators, comma decimal, prefix symbol)
	return FormatFrenchCurrency(amount, currencySymbol)
}
