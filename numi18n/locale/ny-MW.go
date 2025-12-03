package locale

import (
	"github.com/shopspring/decimal"
)

// NYMWLocale represents the Chichewa (Malawi) locale
var NYMWLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Malawi Kwacha",
		Plural:   "Malawi Kwacha",
		Singular: "Malawi Kwacha",
		Symbol:   "MK",
		FractionUnit: FractionUnit{
			Name:     "Tambala",
			Plural:   "Tambala",
			Singular: "Tambala",
			Symbol:   "t",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Malawi",
		Currency:       "MWK",
		ISO3166Alpha2:  "MW",
		ISO3166Alpha3:  "MWI",
		ISO3166Numeric: "454",
		Locale:         "ny-MW",
		Timezone:       []string{"Africa/Blantyre"},
		Language:       "ny",
		Emoji:          "🇲🇼",
	},
	Texts: Texts{
		And:   "ndi",
		Minus: "chotsitsa",
		Only:  "kokha",
		Point: "dothi",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "chikwi chambiri cha biliyoni"},
		{Number: 1000000000000, Value: "chikwi cha biliyoni"},
		{Number: 1000000000, Value: "biliyoni"},
		{Number: 1000000, Value: "miliyoni"},
		{Number: 100000, Value: "chikwi"},
		{Number: 90000, Value: "makumi asanu ndi anayi chikwi"},
		{Number: 80000, Value: "makumi asanu ndi atatu chikwi"},
		{Number: 70000, Value: "makumi asanu ndi awiri chikwi"},
		{Number: 60000, Value: "makumi asanu ndi chimodzi chikwi"},
		{Number: 50000, Value: "makumi asanu chikwi"},
		{Number: 40000, Value: "makumi anayi chikwi"},
		{Number: 30000, Value: "makumi atatu chikwi"},
		{Number: 20000, Value: "makumi awiri chikwi"},
		{Number: 19000, Value: "khumi ndi zisanu ndi zinayi chikwi"},
		{Number: 18000, Value: "khumi ndi zisanu ndi zitatu chikwi"},
		{Number: 17000, Value: "khumi ndi zisanu ndi ziwiri chikwi"},
		{Number: 16000, Value: "khumi ndi zisanu ndi chimodzi chikwi"},
		{Number: 15000, Value: "khumi ndi zisanu chikwi"},
		{Number: 14000, Value: "khumi ndi zinayi chikwi"},
		{Number: 13000, Value: "khumi ndi zitatu chikwi"},
		{Number: 12000, Value: "khumi ndi ziwiri chikwi"},
		{Number: 11000, Value: "khumi ndi chimodzi chikwi"},
		{Number: 10000, Value: "khumi chikwi"},
		{Number: 9000, Value: "zisanu ndi zinayi chikwi"},
		{Number: 8000, Value: "zisanu ndi zitatu chikwi"},
		{Number: 7000, Value: "zisanu ndi ziwiri chikwi"},
		{Number: 6000, Value: "zisanu ndi chimodzi chikwi"},
		{Number: 5000, Value: "zisanu chikwi"},
		{Number: 4000, Value: "zinayi chikwi"},
		{Number: 3000, Value: "zitatu chikwi"},
		{Number: 2000, Value: "ziwiri chikwi"},
		{Number: 1000, Value: "chikwi"},
		{Number: 900, Value: "zisanu ndi zinayi mazana"},
		{Number: 800, Value: "zisanu ndi zitatu mazana"},
		{Number: 700, Value: "zisanu ndi ziwiri mazana"},
		{Number: 600, Value: "zisanu ndi chimodzi mazana"},
		{Number: 500, Value: "zisanu mazana"},
		{Number: 400, Value: "zinayi mazana"},
		{Number: 300, Value: "zitatu mazana"},
		{Number: 200, Value: "ziwiri mazana"},
		{Number: 100, Value: "zana"},
		{Number: 99, Value: "makumi asanu ndi zinayi ndi zisanu ndi zinayi"},
		{Number: 98, Value: "makumi asanu ndi zinayi ndi zisanu ndi zitatu"},
		{Number: 97, Value: "makumi asanu ndi zinayi ndi zisanu ndi ziwiri"},
		{Number: 96, Value: "makumi asanu ndi zinayi ndi zisanu ndi chimodzi"},
		{Number: 95, Value: "makumi asanu ndi zinayi ndi zisanu"},
		{Number: 94, Value: "makumi asanu ndi zinayi ndi zinayi"},
		{Number: 93, Value: "makumi asanu ndi zinayi ndi zitatu"},
		{Number: 92, Value: "makumi asanu ndi zinayi ndi ziwiri"},
		{Number: 91, Value: "makumi asanu ndi zinayi ndi chimodzi"},
		{Number: 90, Value: "makumi asanu ndi zinayi"},
		{Number: 89, Value: "makumi asanu ndi atatu ndi zisanu ndi zinayi"},
		{Number: 88, Value: "makumi asanu ndi atatu ndi zisanu ndi zitatu"},
		{Number: 87, Value: "makumi asanu ndi atatu ndi zisanu ndi ziwiri"},
		{Number: 86, Value: "makumi asanu ndi atatu ndi zisanu ndi chimodzi"},
		{Number: 85, Value: "makumi asanu ndi atatu ndi zisanu"},
		{Number: 84, Value: "makumi asanu ndi atatu ndi zinayi"},
		{Number: 83, Value: "makumi asanu ndi atatu ndi zitatu"},
		{Number: 82, Value: "makumi asanu ndi atatu ndi ziwiri"},
		{Number: 81, Value: "makumi asanu ndi atatu ndi chimodzi"},
		{Number: 80, Value: "makumi asanu ndi atatu"},
		{Number: 79, Value: "makumi asanu ndi awiri ndi zisanu ndi zinayi"},
		{Number: 78, Value: "makumi asanu ndi awiri ndi zisanu ndi zitatu"},
		{Number: 77, Value: "makumi asanu ndi awiri ndi zisanu ndi ziwiri"},
		{Number: 76, Value: "makumi asanu ndi awiri ndi zisanu ndi chimodzi"},
		{Number: 75, Value: "makumi asanu ndi awiri ndi zisanu"},
		{Number: 74, Value: "makumi asanu ndi awiri ndi zinayi"},
		{Number: 73, Value: "makumi asanu ndi awiri ndi zitatu"},
		{Number: 72, Value: "makumi asanu ndi awiri ndi ziwiri"},
		{Number: 71, Value: "makumi asanu ndi awiri ndi chimodzi"},
		{Number: 70, Value: "makumi asanu ndi awiri"},
		{Number: 69, Value: "makumi asanu ndi chimodzi ndi zisanu ndi zinayi"},
		{Number: 68, Value: "makumi asanu ndi chimodzi ndi zisanu ndi zitatu"},
		{Number: 67, Value: "makumi asanu ndi chimodzi ndi zisanu ndi ziwiri"},
		{Number: 66, Value: "makumi asanu ndi chimodzi ndi zisanu ndi chimodzi"},
		{Number: 65, Value: "makumi asanu ndi chimodzi ndi zisanu"},
		{Number: 64, Value: "makumi asanu ndi chimodzi ndi zinayi"},
		{Number: 63, Value: "makumi asanu ndi chimodzi ndi zitatu"},
		{Number: 62, Value: "makumi asanu ndi chimodzi ndi ziwiri"},
		{Number: 61, Value: "makumi asanu ndi chimodzi ndi chimodzi"},
		{Number: 60, Value: "makumi asanu ndi chimodzi"},
		{Number: 59, Value: "makumi asanu ndi zisanu ndi zinayi"},
		{Number: 58, Value: "makumi asanu ndi zisanu ndi zitatu"},
		{Number: 57, Value: "makumi asanu ndi zisanu ndi ziwiri"},
		{Number: 56, Value: "makumi asanu ndi zisanu ndi chimodzi"},
		{Number: 55, Value: "makumi asanu ndi zisanu"},
		{Number: 54, Value: "makumi asanu ndi zinayi"},
		{Number: 53, Value: "makumi asanu ndi zitatu"},
		{Number: 52, Value: "makumi asanu ndi ziwiri"},
		{Number: 51, Value: "makumi asanu ndi chimodzi"},
		{Number: 50, Value: "makumi asanu"},
		{Number: 49, Value: "makumi anayi ndi zisanu ndi zinayi"},
		{Number: 48, Value: "makumi anayi ndi zisanu ndi zitatu"},
		{Number: 47, Value: "makumi anayi ndi zisanu ndi ziwiri"},
		{Number: 46, Value: "makumi anayi ndi zisanu ndi chimodzi"},
		{Number: 45, Value: "makumi anayi ndi zisanu"},
		{Number: 44, Value: "makumi anayi ndi zinayi"},
		{Number: 43, Value: "makumi anayi ndi zitatu"},
		{Number: 42, Value: "makumi anayi ndi ziwiri"},
		{Number: 41, Value: "makumi anayi ndi chimodzi"},
		{Number: 40, Value: "makumi anayi"},
		{Number: 39, Value: "makumi atatu ndi zisanu ndi zinayi"},
		{Number: 38, Value: "makumi atatu ndi zisanu ndi zitatu"},
		{Number: 37, Value: "makumi atatu ndi zisanu ndi ziwiri"},
		{Number: 36, Value: "makumi atatu ndi zisanu ndi chimodzi"},
		{Number: 35, Value: "makumi atatu ndi zisanu"},
		{Number: 34, Value: "makumi atatu ndi zinayi"},
		{Number: 33, Value: "makumi atatu ndi zitatu"},
		{Number: 32, Value: "makumi atatu ndi ziwiri"},
		{Number: 31, Value: "makumi atatu ndi chimodzi"},
		{Number: 30, Value: "makumi atatu"},
		{Number: 29, Value: "makumi awiri ndi zisanu ndi zinayi"},
		{Number: 28, Value: "makumi awiri ndi zisanu ndi zitatu"},
		{Number: 27, Value: "makumi awiri ndi zisanu ndi ziwiri"},
		{Number: 26, Value: "makumi awiri ndi zisanu ndi chimodzi"},
		{Number: 25, Value: "makumi awiri ndi zisanu"},
		{Number: 24, Value: "makumi awiri ndi zinayi"},
		{Number: 23, Value: "makumi awiri ndi zitatu"},
		{Number: 22, Value: "makumi awiri ndi ziwiri"},
		{Number: 21, Value: "makumi awiri ndi chimodzi"},
		{Number: 20, Value: "makumi awiri"},
		{Number: 19, Value: "khumi ndi zisanu ndi zinayi"},
		{Number: 18, Value: "khumi ndi zisanu ndi zitatu"},
		{Number: 17, Value: "khumi ndi zisanu ndi ziwiri"},
		{Number: 16, Value: "khumi ndi zisanu ndi chimodzi"},
		{Number: 15, Value: "khumi ndi zisanu"},
		{Number: 14, Value: "khumi ndi zinayi"},
		{Number: 13, Value: "khumi ndi zitatu"},
		{Number: 12, Value: "khumi ndi ziwiri"},
		{Number: 11, Value: "khumi ndi chimodzi"},
		{Number: 10, Value: "khumi"},
		{Number: 9, Value: "zisanu ndi zinayi"},
		{Number: 8, Value: "zisanu ndi zitatu"},
		{Number: 7, Value: "zisanu ndi ziwiri"},
		{Number: 6, Value: "zisanu ndi chimodzi"},
		{Number: 5, Value: "zisanu"},
		{Number: 4, Value: "zinayi"},
		{Number: 3, Value: "zitatu"},
		{Number: 2, Value: "ziwiri"},
		{Number: 1, Value: "chimodzi"},
		{Number: 0, Value: "ziro"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Zana limodzi"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "choyamba", Suffix: "-oyamba", Masculine: "choyamba", Feminine: "choyamba", Neuter: "choyamba"},
		{Number: 2, Word: "chachiwiri", Suffix: "-achiwiri", Masculine: "chachiwiri", Feminine: "chachiwiri", Neuter: "chachiwiri"},
		{Number: 3, Word: "chachitatu", Suffix: "-achitatu", Masculine: "chachitatu", Feminine: "chachitatu", Neuter: "chachitatu"},
		{Number: 4, Word: "chachinayi", Suffix: "-achinayi", Masculine: "chachinayi", Feminine: "chachinayi", Neuter: "chachinayi"},
		{Number: 5, Word: "chachisanu", Suffix: "-achisanu", Masculine: "chachisanu", Feminine: "chachisanu", Neuter: "chachisanu"},
		{Number: 6, Word: "chachisanu ndi chimodzi", Suffix: "-achisanu ndi chimodzi", Masculine: "chachisanu ndi chimodzi", Feminine: "chachisanu ndi chimodzi", Neuter: "chachisanu ndi chimodzi"},
		{Number: 7, Word: "chachisanu ndi chiwiri", Suffix: "-achisanu ndi chiwiri", Masculine: "chachisanu ndi chiwiri", Feminine: "chachisanu ndi chiwiri", Neuter: "chachisanu ndi chiwiri"},
		{Number: 8, Word: "chachisanu ndi chitatu", Suffix: "-achisanu ndi chitatu", Masculine: "chachisanu ndi chitatu", Feminine: "chachisanu ndi chitatu", Neuter: "chachisanu ndi chitatu"},
		{Number: 9, Word: "chachisanu ndi chinayi", Suffix: "-achisanu ndi chinayi", Masculine: "chachisanu ndi chinayi", Feminine: "chachisanu ndi chinayi", Neuter: "chachisanu ndi chinayi"},
		{Number: 10, Word: "chakhumi", Suffix: "-akhumi", Masculine: "chakhumi", Feminine: "chakhumi", Neuter: "chakhumi"},
		{Number: 11, Word: "chakhumi ndi chimodzi", Suffix: "-akhumi ndi chimodzi", Masculine: "chakhumi ndi chimodzi", Feminine: "chakhumi ndi chimodzi", Neuter: "chakhumi ndi chimodzi"},
		{Number: 12, Word: "chakhumi ndi chiwiri", Suffix: "-akhumi ndi chiwiri", Masculine: "chakhumi ndi chiwiri", Feminine: "chakhumi ndi chiwiri", Neuter: "chakhumi ndi chiwiri"},
		{Number: 20, Word: "chamakumi awiri", Suffix: "-amakumi awiri", Masculine: "chamakumi awiri", Feminine: "chamakumi awiri", Neuter: "chamakumi awiri"},
		{Number: 21, Word: "chamakumi awiri ndi chimodzi", Suffix: "-amakumi awiri ndi chimodzi", Masculine: "chamakumi awiri ndi chimodzi", Feminine: "chamakumi awiri ndi chimodzi", Neuter: "chamakumi awiri ndi chimodzi"},
		{Number: 30, Word: "chamakumi atatu", Suffix: "-amakumi atatu", Masculine: "chamakumi atatu", Feminine: "chamakumi atatu", Neuter: "chamakumi atatu"},
		{Number: 100, Word: "chazana", Suffix: "-azana", Masculine: "chazana", Feminine: "chazana", Neuter: "chazana"},
		{Number: 1000, Word: "chachikwi", Suffix: "-achikwi", Masculine: "chachikwi", Feminine: "chachikwi", Neuter: "chachikwi"},
	},
	LocaleFormatter: &ChichewaFormatter{},
}

// ChichewaFormatter handles Chichewa (ny-MW) formatting
type ChichewaFormatter struct{}

func (f *ChichewaFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *ChichewaFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *ChichewaFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *ChichewaFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *ChichewaFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *ChichewaFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 2
	}
	return amount.Truncate(int32(precision))
}

func (f *ChichewaFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *ChichewaFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
