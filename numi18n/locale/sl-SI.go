package locale

import (
	"github.com/shopspring/decimal"
)

// SLSILocale represents the Slovenian (Slovenia) locale
var SLSILocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Euro",
		Plural:   "Evri",
		Singular: "Evro",
		Symbol:   "€",
		FractionUnit: FractionUnit{
			Name:     "Cent",
			Plural:   "Centi",
			Singular: "Cent",
			Symbol:   "c",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Slovenia",
		Currency:       "EUR",
		ISO3166Alpha2:  "SI",
		ISO3166Alpha3:  "SVN",
		ISO3166Numeric: "705",
		Locale:         "sl-SI",
		Timezone:       []string{"Europe/Ljubljana"},
		Language:       "sl",
		Emoji:          "🇸🇮",
	},
	Texts: Texts{
		And:   "in",
		Minus: "minus",
		Only:  "samo",
		Point: "pika",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "tisoč bilijonov"},
		{Number: 1000000000000, Value: "bilijon"},
		{Number: 1000000000, Value: "milijarda"},
		{Number: 1000000, Value: "milijon"},
		{Number: 100000, Value: "sto tisoč"},
		{Number: 90000, Value: "devetdeset tisoč"},
		{Number: 80000, Value: "osemdeset tisoč"},
		{Number: 70000, Value: "sedemdeset tisoč"},
		{Number: 60000, Value: "šestdeset tisoč"},
		{Number: 50000, Value: "petdeset tisoč"},
		{Number: 40000, Value: "štirideset tisoč"},
		{Number: 30000, Value: "trideset tisoč"},
		{Number: 20000, Value: "dvajset tisoč"},
		{Number: 19000, Value: "devetnajst tisoč"},
		{Number: 18000, Value: "osemnajst tisoč"},
		{Number: 17000, Value: "sedemnajst tisoč"},
		{Number: 16000, Value: "šestnajst tisoč"},
		{Number: 15000, Value: "petnajst tisoč"},
		{Number: 14000, Value: "štirinajst tisoč"},
		{Number: 13000, Value: "trinajst tisoč"},
		{Number: 12000, Value: "dvanajst tisoč"},
		{Number: 11000, Value: "enajst tisoč"},
		{Number: 10000, Value: "deset tisoč"},
		{Number: 9000, Value: "devet tisoč"},
		{Number: 8000, Value: "osem tisoč"},
		{Number: 7000, Value: "sedem tisoč"},
		{Number: 6000, Value: "šest tisoč"},
		{Number: 5000, Value: "pet tisoč"},
		{Number: 4000, Value: "štiri tisoč"},
		{Number: 3000, Value: "tri tisoč"},
		{Number: 2000, Value: "dva tisoč"},
		{Number: 1000, Value: "tisoč"},
		{Number: 900, Value: "devetsto"},
		{Number: 800, Value: "osemsto"},
		{Number: 700, Value: "sedemsto"},
		{Number: 600, Value: "šeststo"},
		{Number: 500, Value: "petsto"},
		{Number: 400, Value: "štiristo"},
		{Number: 300, Value: "tristo"},
		{Number: 200, Value: "dvesto"},
		{Number: 100, Value: "sto"},
		{Number: 99, Value: "devetdeset devet"},
		{Number: 98, Value: "devetdeset osem"},
		{Number: 97, Value: "devetdeset sedem"},
		{Number: 96, Value: "devetdeset šest"},
		{Number: 95, Value: "devetdeset pet"},
		{Number: 94, Value: "devetdeset štiri"},
		{Number: 93, Value: "devetdeset tri"},
		{Number: 92, Value: "devetdeset dva"},
		{Number: 91, Value: "devetdeset ena"},
		{Number: 90, Value: "devetdeset"},
		{Number: 89, Value: "osemdeset devet"},
		{Number: 88, Value: "osemdeset osem"},
		{Number: 87, Value: "osemdeset sedem"},
		{Number: 86, Value: "osemdeset šest"},
		{Number: 85, Value: "osemdeset pet"},
		{Number: 84, Value: "osemdeset štiri"},
		{Number: 83, Value: "osemdeset tri"},
		{Number: 82, Value: "osemdeset dva"},
		{Number: 81, Value: "osemdeset ena"},
		{Number: 80, Value: "osemdeset"},
		{Number: 79, Value: "sedemdeset devet"},
		{Number: 78, Value: "sedemdeset osem"},
		{Number: 77, Value: "sedemdeset sedem"},
		{Number: 76, Value: "sedemdeset šest"},
		{Number: 75, Value: "sedemdeset pet"},
		{Number: 74, Value: "sedemdeset štiri"},
		{Number: 73, Value: "sedemdeset tri"},
		{Number: 72, Value: "sedemdeset dva"},
		{Number: 71, Value: "sedemdeset ena"},
		{Number: 70, Value: "sedemdeset"},
		{Number: 69, Value: "šestdeset devet"},
		{Number: 68, Value: "šestdeset osem"},
		{Number: 67, Value: "šestdeset sedem"},
		{Number: 66, Value: "šestdeset šest"},
		{Number: 65, Value: "šestdeset pet"},
		{Number: 64, Value: "šestdeset štiri"},
		{Number: 63, Value: "šestdeset tri"},
		{Number: 62, Value: "šestdeset dva"},
		{Number: 61, Value: "šestdeset ena"},
		{Number: 60, Value: "šestdeset"},
		{Number: 59, Value: "petdeset devet"},
		{Number: 58, Value: "petdeset osem"},
		{Number: 57, Value: "petdeset sedem"},
		{Number: 56, Value: "petdeset šest"},
		{Number: 55, Value: "petdeset pet"},
		{Number: 54, Value: "petdeset štiri"},
		{Number: 53, Value: "petdeset tri"},
		{Number: 52, Value: "petdeset dva"},
		{Number: 51, Value: "petdeset ena"},
		{Number: 50, Value: "petdeset"},
		{Number: 49, Value: "štirideset devet"},
		{Number: 48, Value: "štirideset osem"},
		{Number: 47, Value: "štirideset sedem"},
		{Number: 46, Value: "štirideset šest"},
		{Number: 45, Value: "štirideset pet"},
		{Number: 44, Value: "štirideset štiri"},
		{Number: 43, Value: "štirideset tri"},
		{Number: 42, Value: "štirideset dva"},
		{Number: 41, Value: "štirideset ena"},
		{Number: 40, Value: "štirideset"},
		{Number: 39, Value: "trideset devet"},
		{Number: 38, Value: "trideset osem"},
		{Number: 37, Value: "trideset sedem"},
		{Number: 36, Value: "trideset šest"},
		{Number: 35, Value: "trideset pet"},
		{Number: 34, Value: "trideset štiri"},
		{Number: 33, Value: "trideset tri"},
		{Number: 32, Value: "trideset dva"},
		{Number: 31, Value: "trideset ena"},
		{Number: 30, Value: "trideset"},
		{Number: 29, Value: "dvajset devet"},
		{Number: 28, Value: "dvajset osem"},
		{Number: 27, Value: "dvajset sedem"},
		{Number: 26, Value: "dvajset šest"},
		{Number: 25, Value: "dvajset pet"},
		{Number: 24, Value: "dvajset štiri"},
		{Number: 23, Value: "dvajset tri"},
		{Number: 22, Value: "dvajset dva"},
		{Number: 21, Value: "dvajset ena"},
		{Number: 20, Value: "dvajset"},
		{Number: 19, Value: "devetnajst"},
		{Number: 18, Value: "osemnajst"},
		{Number: 17, Value: "sedemnajst"},
		{Number: 16, Value: "šestnajst"},
		{Number: 15, Value: "petnajst"},
		{Number: 14, Value: "štirinajst"},
		{Number: 13, Value: "trinajst"},
		{Number: 12, Value: "dvanajst"},
		{Number: 11, Value: "enajst"},
		{Number: 10, Value: "deset"},
		{Number: 9, Value: "devet"},
		{Number: 8, Value: "osem"},
		{Number: 7, Value: "sedem"},
		{Number: 6, Value: "šest"},
		{Number: 5, Value: "pet"},
		{Number: 4, Value: "štiri"},
		{Number: 3, Value: "tri"},
		{Number: 2, Value: "dva"},
		{Number: 1, Value: "ena"},
		{Number: 0, Value: "nič"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Sto"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "prvi", Suffix: ".", Masculine: "prvi", Feminine: "prva", Neuter: "prvo"},
		{Number: 2, Word: "drugi", Suffix: ".", Masculine: "drugi", Feminine: "druga", Neuter: "drugo"},
		{Number: 3, Word: "tretji", Suffix: ".", Masculine: "tretji", Feminine: "tretja", Neuter: "tretje"},
		{Number: 4, Word: "četrti", Suffix: ".", Masculine: "četrti", Feminine: "četrta", Neuter: "četrto"},
		{Number: 5, Word: "peti", Suffix: ".", Masculine: "peti", Feminine: "peta", Neuter: "peto"},
		{Number: 6, Word: "šesti", Suffix: ".", Masculine: "šesti", Feminine: "šesta", Neuter: "šesto"},
		{Number: 7, Word: "sedmi", Suffix: ".", Masculine: "sedmi", Feminine: "sedma", Neuter: "sedmo"},
		{Number: 8, Word: "osmi", Suffix: ".", Masculine: "osmi", Feminine: "osma", Neuter: "osmo"},
		{Number: 9, Word: "deveti", Suffix: ".", Masculine: "deveti", Feminine: "deveta", Neuter: "deveto"},
		{Number: 10, Word: "deseti", Suffix: ".", Masculine: "deseti", Feminine: "deseta", Neuter: "deseto"},
		{Number: 11, Word: "enajsti", Suffix: ".", Masculine: "enajsti", Feminine: "enajsta", Neuter: "enajsto"},
		{Number: 12, Word: "dvanajsti", Suffix: ".", Masculine: "dvanajsti", Feminine: "dvanajsta", Neuter: "dvanajsto"},
		{Number: 20, Word: "dvajseti", Suffix: ".", Masculine: "dvajseti", Feminine: "dvajseta", Neuter: "dvajseto"},
		{Number: 21, Word: "enaindvajseti", Suffix: ".", Masculine: "enaindvajseti", Feminine: "enaindvajseta", Neuter: "enaindvajseto"},
		{Number: 30, Word: "trideseti", Suffix: ".", Masculine: "trideseti", Feminine: "trideseta", Neuter: "trideseto"},
		{Number: 50, Word: "petdeseti", Suffix: ".", Masculine: "petdeseti", Feminine: "petdeseta", Neuter: "petdeseto"},
		{Number: 100, Word: "stoti", Suffix: ".", Masculine: "stoti", Feminine: "stota", Neuter: "stoto"},
		{Number: 1000, Word: "tisoči", Suffix: ".", Masculine: "tisoči", Feminine: "tisoča", Neuter: "tisoče"},
	},
	LocaleFormatter: &SlovenianSloveniaFormatter{},
}

// SlovenianSloveniaFormatter handles Slovenian (Slovenia) formatting
type SlovenianSloveniaFormatter struct{}

func (f *SlovenianSloveniaFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *SlovenianSloveniaFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *SlovenianSloveniaFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *SlovenianSloveniaFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *SlovenianSloveniaFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *SlovenianSloveniaFormatter) ChopDecimal(d decimal.Decimal, precision int) decimal.Decimal {
	return d.Truncate(int32(precision))
}

func (f *SlovenianSloveniaFormatter) FormatDecimalNumber(amount float64) string {
	return FormatEuropeanDecimal(amount)
}
func (f *SlovenianSloveniaFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with European conventions (period separators, comma decimal, prefix symbol)
	return FormatEuropeanCurrency(amount, currencySymbol)
}
