package locale

import (
	"github.com/shopspring/decimal"
)

// RWRWLocale represents the Kinyarwanda (Rwanda) locale
var RWRWLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Ifaranga ry'u Rwanda",
		Plural:   "Amafaranga y'u Rwanda",
		Singular: "Ifaranga ry'u Rwanda",
		Symbol:   "RWF",
		FractionUnit: FractionUnit{
			Name:     "Santima",
			Plural:   "Santima",
			Singular: "Santima",
			Symbol:   "s",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Rwanda",
		Currency:       "RWF",
		ISO3166Alpha2:  "RW",
		ISO3166Alpha3:  "RWA",
		ISO3166Numeric: "646",
		Locale:         "rw-RW",
		Timezone:       []string{"Africa/Kigali"},
		Language:       "rw",
		Emoji:          "🇷🇼",
	},
	Texts: Texts{
		And:   "na",
		Minus: "kubana",
		Only:  "gusa",
		Point: "akadomo",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "umuyobozi mukuru"},
		{Number: 1000000000000, Value: "umuyobozi"},
		{Number: 1000000000, Value: "umuyagutse"},
		{Number: 1000000, Value: "miliyoni"},
		{Number: 100000, Value: "ijana b'amakumi"},
		{Number: 90000, Value: "mirongo icyenda n'amakumi"},
		{Number: 80000, Value: "mirongo inani n'amakumi"},
		{Number: 70000, Value: "mirongo irindwi n'amakumi"},
		{Number: 60000, Value: "mirongo itandatu n'amakumi"},
		{Number: 50000, Value: "mirongo itanu n'amakumi"},
		{Number: 40000, Value: "mirongo ine n'amakumi"},
		{Number: 30000, Value: "mirongo itatu n'amakumi"},
		{Number: 20000, Value: "mirongo ibiri n'amakumi"},
		{Number: 19000, Value: "icumi n'icyenda mu makumi"},
		{Number: 18000, Value: "icumi n'inani mu makumi"},
		{Number: 17000, Value: "icumi n'irindwi mu makumi"},
		{Number: 16000, Value: "icumi n'itandatu mu makumi"},
		{Number: 15000, Value: "icumi n'itanu mu makumi"},
		{Number: 14000, Value: "icumi n'ine mu makumi"},
		{Number: 13000, Value: "icumi n'itatu mu makumi"},
		{Number: 12000, Value: "icumi n'ibiri mu makumi"},
		{Number: 11000, Value: "icumi n'imwe mu makumi"},
		{Number: 10000, Value: "icumi mu makumi"},
		{Number: 9000, Value: "icyenda mu makumi"},
		{Number: 8000, Value: "inani mu makumi"},
		{Number: 7000, Value: "irindwi mu makumi"},
		{Number: 6000, Value: "itandatu mu makumi"},
		{Number: 5000, Value: "itanu mu makumi"},
		{Number: 4000, Value: "ine mu makumi"},
		{Number: 3000, Value: "itatu mu makumi"},
		{Number: 2000, Value: "ibiri mu makumi"},
		{Number: 1000, Value: "ijana"},
		{Number: 900, Value: "ijana icyenda"},
		{Number: 800, Value: "ijana inani"},
		{Number: 700, Value: "ijana irindwi"},
		{Number: 600, Value: "ijana itandatu"},
		{Number: 500, Value: "ijana itanu"},
		{Number: 400, Value: "ijana ine"},
		{Number: 300, Value: "ijana itatu"},
		{Number: 200, Value: "ijana ibiri"},
		{Number: 100, Value: "ijana"},
		{Number: 99, Value: "mirongo icyenda n'icyenda"},
		{Number: 98, Value: "mirongo icyenda n'inani"},
		{Number: 97, Value: "mirongo icyenda n'irindwi"},
		{Number: 96, Value: "mirongo icyenda n'itandatu"},
		{Number: 95, Value: "mirongo icyenda n'itanu"},
		{Number: 94, Value: "mirongo icyenda n'ine"},
		{Number: 93, Value: "mirongo icyenda n'itatu"},
		{Number: 92, Value: "mirongo icyenda n'ibiri"},
		{Number: 91, Value: "mirongo icyenda n'imwe"},
		{Number: 90, Value: "mirongo icyenda"},
		{Number: 89, Value: "mirongo inani n'icyenda"},
		{Number: 88, Value: "mirongo inani n'inani"},
		{Number: 87, Value: "mirongo inani n'irindwi"},
		{Number: 86, Value: "mirongo inani n'itandatu"},
		{Number: 85, Value: "mirongo inani n'itanu"},
		{Number: 84, Value: "mirongo inani n'ine"},
		{Number: 83, Value: "mirongo inani n'itatu"},
		{Number: 82, Value: "mirongo inani n'ibiri"},
		{Number: 81, Value: "mirongo inani n'imwe"},
		{Number: 80, Value: "mirongo inani"},
		{Number: 79, Value: "mirongo irindwi n'icyenda"},
		{Number: 78, Value: "mirongo irindwi n'inani"},
		{Number: 77, Value: "mirongo irindwi n'irindwi"},
		{Number: 76, Value: "mirongo irindwi n'itandatu"},
		{Number: 75, Value: "mirongo irindwi n'itanu"},
		{Number: 74, Value: "mirongo irindwi n'ine"},
		{Number: 73, Value: "mirongo irindwi n'itatu"},
		{Number: 72, Value: "mirongo irindwi n'ibiri"},
		{Number: 71, Value: "mirongo irindwi n'imwe"},
		{Number: 70, Value: "mirongo irindwi"},
		{Number: 69, Value: "mirongo itandatu n'icyenda"},
		{Number: 68, Value: "mirongo itandatu n'inani"},
		{Number: 67, Value: "mirongo itandatu n'irindwi"},
		{Number: 66, Value: "mirongo itandatu n'itandatu"},
		{Number: 65, Value: "mirongo itandatu n'itanu"},
		{Number: 64, Value: "mirongo itandatu n'ine"},
		{Number: 63, Value: "mirongo itandatu n'itatu"},
		{Number: 62, Value: "mirongo itandatu n'ibiri"},
		{Number: 61, Value: "mirongo itandatu n'imwe"},
		{Number: 60, Value: "mirongo itandatu"},
		{Number: 59, Value: "mirongo itanu n'icyenda"},
		{Number: 58, Value: "mirongo itanu n'inani"},
		{Number: 57, Value: "mirongo itanu n'irindwi"},
		{Number: 56, Value: "mirongo itanu n'itandatu"},
		{Number: 55, Value: "mirongo itanu n'itanu"},
		{Number: 54, Value: "mirongo itanu n'ine"},
		{Number: 53, Value: "mirongo itanu n'itatu"},
		{Number: 52, Value: "mirongo itanu n'ibiri"},
		{Number: 51, Value: "mirongo itanu n'imwe"},
		{Number: 50, Value: "mirongo itanu"},
		{Number: 49, Value: "mirongo ine n'icyenda"},
		{Number: 48, Value: "mirongo ine n'inani"},
		{Number: 47, Value: "mirongo ine n'irindwi"},
		{Number: 46, Value: "mirongo ine n'itandatu"},
		{Number: 45, Value: "mirongo ine n'itanu"},
		{Number: 44, Value: "mirongo ine n'ine"},
		{Number: 43, Value: "mirongo ine n'itatu"},
		{Number: 42, Value: "mirongo ine n'ibiri"},
		{Number: 41, Value: "mirongo ine n'imwe"},
		{Number: 40, Value: "mirongo ine"},
		{Number: 39, Value: "mirongo itatu n'icyenda"},
		{Number: 38, Value: "mirongo itatu n'inani"},
		{Number: 37, Value: "mirongo itatu n'irindwi"},
		{Number: 36, Value: "mirongo itatu n'itandatu"},
		{Number: 35, Value: "mirongo itatu n'itanu"},
		{Number: 34, Value: "mirongo itatu n'ine"},
		{Number: 33, Value: "mirongo itatu n'itatu"},
		{Number: 32, Value: "mirongo itatu n'ibiri"},
		{Number: 31, Value: "mirongo itatu n'imwe"},
		{Number: 30, Value: "mirongo itatu"},
		{Number: 29, Value: "mirongo ibiri n'icyenda"},
		{Number: 28, Value: "mirongo ibiri n'inani"},
		{Number: 27, Value: "mirongo ibiri n'irindwi"},
		{Number: 26, Value: "mirongo ibiri n'itandatu"},
		{Number: 25, Value: "mirongo ibiri n'itanu"},
		{Number: 24, Value: "mirongo ibiri n'ine"},
		{Number: 23, Value: "mirongo ibiri n'itatu"},
		{Number: 22, Value: "mirongo ibiri n'ibiri"},
		{Number: 21, Value: "mirongo ibiri n'imwe"},
		{Number: 20, Value: "mirongo ibiri"},
		{Number: 19, Value: "icumi n'icyenda"},
		{Number: 18, Value: "icumi n'inani"},
		{Number: 17, Value: "icumi n'irindwi"},
		{Number: 16, Value: "icumi n'itandatu"},
		{Number: 15, Value: "icumi n'itanu"},
		{Number: 14, Value: "icumi n'ine"},
		{Number: 13, Value: "icumi n'itatu"},
		{Number: 12, Value: "icumi n'ibiri"},
		{Number: 11, Value: "icumi n'imwe"},
		{Number: 10, Value: "icumi"},
		{Number: 9, Value: "icyenda"},
		{Number: 8, Value: "inani"},
		{Number: 7, Value: "irindwi"},
		{Number: 6, Value: "itandatu"},
		{Number: 5, Value: "itanu"},
		{Number: 4, Value: "ine"},
		{Number: 3, Value: "itatu"},
		{Number: 2, Value: "ibiri"},
		{Number: 1, Value: "rimwe"},
		{Number: 0, Value: "ubusa"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Ijana"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "wa mbere", Suffix: "-a", Masculine: "wa mbere", Feminine: "wa mbere", Neuter: "wa mbere"},
		{Number: 2, Word: "wa kabiri", Suffix: "-a", Masculine: "wa kabiri", Feminine: "wa kabiri", Neuter: "wa kabiri"},
		{Number: 3, Word: "wa gatatu", Suffix: "-a", Masculine: "wa gatatu", Feminine: "wa gatatu", Neuter: "wa gatatu"},
		{Number: 4, Word: "wa kane", Suffix: "-a", Masculine: "wa kane", Feminine: "wa kane", Neuter: "wa kane"},
		{Number: 5, Word: "wa gatanu", Suffix: "-a", Masculine: "wa gatanu", Feminine: "wa gatanu", Neuter: "wa gatanu"},
		{Number: 6, Word: "wa gatandatu", Suffix: "-a", Masculine: "wa gatandatu", Feminine: "wa gatandatu", Neuter: "wa gatandatu"},
		{Number: 7, Word: "wa karindwi", Suffix: "-a", Masculine: "wa karindwi", Feminine: "wa karindwi", Neuter: "wa karindwi"},
		{Number: 8, Word: "wa munani", Suffix: "-a", Masculine: "wa munani", Feminine: "wa munani", Neuter: "wa munani"},
		{Number: 9, Word: "wa cyenda", Suffix: "-a", Masculine: "wa cyenda", Feminine: "wa cyenda", Neuter: "wa cyenda"},
		{Number: 10, Word: "wa cumi", Suffix: "-a", Masculine: "wa cumi", Feminine: "wa cumi", Neuter: "wa cumi"},
		{Number: 11, Word: "wa cumi na rimwe", Suffix: "-a", Masculine: "wa cumi na rimwe", Feminine: "wa cumi na rimwe", Neuter: "wa cumi na rimwe"},
		{Number: 12, Word: "wa cumi na kabiri", Suffix: "-a", Masculine: "wa cumi na kabiri", Feminine: "wa cumi na kabiri", Neuter: "wa cumi na kabiri"},
		{Number: 20, Word: "wa makumyabiri", Suffix: "-a", Masculine: "wa makumyabiri", Feminine: "wa makumyabiri", Neuter: "wa makumyabiri"},
		{Number: 21, Word: "wa makumyabiri na rimwe", Suffix: "-a", Masculine: "wa makumyabiri na rimwe", Feminine: "wa makumyabiri na rimwe", Neuter: "wa makumyabiri na rimwe"},
		{Number: 30, Word: "wa makumyatatu", Suffix: "-a", Masculine: "wa makumyatatu", Feminine: "wa makumyatatu", Neuter: "wa makumyatatu"},
		{Number: 50, Word: "wa makumyatanu", Suffix: "-a", Masculine: "wa makumyatanu", Feminine: "wa makumyatanu", Neuter: "wa makumyatanu"},
		{Number: 100, Word: "wa ijana", Suffix: "-a", Masculine: "wa ijana", Feminine: "wa ijana", Neuter: "wa ijana"},
		{Number: 1000, Word: "wa rukumi", Suffix: "-a", Masculine: "wa rukumi", Feminine: "wa rukumi", Neuter: "wa rukumi"},
	},
	LocaleFormatter: &KinyarwandaRwandaFormatter{},
}

// KinyarwandaRwandaFormatter handles Kinyarwanda (Rwanda) formatting
type KinyarwandaRwandaFormatter struct{}

func (f *KinyarwandaRwandaFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *KinyarwandaRwandaFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *KinyarwandaRwandaFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *KinyarwandaRwandaFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return " " + fractionName
	}
	return " " + fractionPlural
}

func (f *KinyarwandaRwandaFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *KinyarwandaRwandaFormatter) ChopDecimal(value decimal.Decimal, precision int) decimal.Decimal {
	return value.Truncate(int32(precision))
}

func (f *KinyarwandaRwandaFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *KinyarwandaRwandaFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
