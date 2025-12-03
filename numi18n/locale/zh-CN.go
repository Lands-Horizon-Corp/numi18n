package locale

import (
	"github.com/shopspring/decimal"
)

// ZHCNLocale represents the Chinese (China) locale
var ZHCNLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Chinese Yuan",
		Plural:   "元",
		Singular: "元",
		Symbol:   "¥",
		FractionUnit: FractionUnit{
			Name:     "Jiao",
			Plural:   "角",
			Singular: "角",
			Symbol:   "角",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "China",
		Currency:       "CNY",
		ISO3166Alpha2:  "CN",
		ISO3166Alpha3:  "CHN",
		ISO3166Numeric: "156",
		Locale:         "zh-CN",
		Timezone:       []string{"Asia/Shanghai"},
		Language:       "zh",
		Emoji:          "🇨🇳",
	},
	Texts: Texts{
		And:   "和",
		Minus: "负",
		Only:  "仅",
		Point: "点",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "一千万亿"},
		{Number: 1000000000000, Value: "一万亿"},
		{Number: 1000000000, Value: "十亿"},
		{Number: 100000000, Value: "一亿"},
		{Number: 10000000, Value: "一千万"},
		{Number: 1000000, Value: "一百万"},
		{Number: 100000, Value: "十万"},
		{Number: 10000, Value: "一万"},
		{Number: 9000, Value: "九千"},
		{Number: 8000, Value: "八千"},
		{Number: 7000, Value: "七千"},
		{Number: 6000, Value: "六千"},
		{Number: 5000, Value: "五千"},
		{Number: 4000, Value: "四千"},
		{Number: 3000, Value: "三千"},
		{Number: 2000, Value: "两千"},
		{Number: 1000, Value: "一千"},
		{Number: 900, Value: "九百"},
		{Number: 800, Value: "八百"},
		{Number: 700, Value: "七百"},
		{Number: 600, Value: "六百"},
		{Number: 500, Value: "五百"},
		{Number: 400, Value: "四百"},
		{Number: 300, Value: "三百"},
		{Number: 200, Value: "两百"},
		{Number: 100, Value: "一百"},
		{Number: 99, Value: "九十九"},
		{Number: 98, Value: "九十八"},
		{Number: 97, Value: "九十七"},
		{Number: 96, Value: "九十六"},
		{Number: 95, Value: "九十五"},
		{Number: 94, Value: "九十四"},
		{Number: 93, Value: "九十三"},
		{Number: 92, Value: "九十二"},
		{Number: 91, Value: "九十一"},
		{Number: 90, Value: "九十"},
		{Number: 89, Value: "八十九"},
		{Number: 88, Value: "八十八"},
		{Number: 87, Value: "八十七"},
		{Number: 86, Value: "八十六"},
		{Number: 85, Value: "八十五"},
		{Number: 84, Value: "八十四"},
		{Number: 83, Value: "八十三"},
		{Number: 82, Value: "八十二"},
		{Number: 81, Value: "八十一"},
		{Number: 80, Value: "八十"},
		{Number: 79, Value: "七十九"},
		{Number: 78, Value: "七十八"},
		{Number: 77, Value: "七十七"},
		{Number: 76, Value: "七十六"},
		{Number: 75, Value: "七十五"},
		{Number: 74, Value: "七十四"},
		{Number: 73, Value: "七十三"},
		{Number: 72, Value: "七十二"},
		{Number: 71, Value: "七十一"},
		{Number: 70, Value: "七十"},
		{Number: 69, Value: "六十九"},
		{Number: 68, Value: "六十八"},
		{Number: 67, Value: "六十七"},
		{Number: 66, Value: "六十六"},
		{Number: 65, Value: "六十五"},
		{Number: 64, Value: "六十四"},
		{Number: 63, Value: "六十三"},
		{Number: 62, Value: "六十二"},
		{Number: 61, Value: "六十一"},
		{Number: 60, Value: "六十"},
		{Number: 59, Value: "五十九"},
		{Number: 58, Value: "五十八"},
		{Number: 57, Value: "五十七"},
		{Number: 56, Value: "五十六"},
		{Number: 55, Value: "五十五"},
		{Number: 54, Value: "五十四"},
		{Number: 53, Value: "五十三"},
		{Number: 52, Value: "五十二"},
		{Number: 51, Value: "五十一"},
		{Number: 50, Value: "五十"},
		{Number: 49, Value: "四十九"},
		{Number: 48, Value: "四十八"},
		{Number: 47, Value: "四十七"},
		{Number: 46, Value: "四十六"},
		{Number: 45, Value: "四十五"},
		{Number: 44, Value: "四十四"},
		{Number: 43, Value: "四十三"},
		{Number: 42, Value: "四十二"},
		{Number: 41, Value: "四十一"},
		{Number: 40, Value: "四十"},
		{Number: 39, Value: "三十九"},
		{Number: 38, Value: "三十八"},
		{Number: 37, Value: "三十七"},
		{Number: 36, Value: "三十六"},
		{Number: 35, Value: "三十五"},
		{Number: 34, Value: "三十四"},
		{Number: 33, Value: "三十三"},
		{Number: 32, Value: "三十二"},
		{Number: 31, Value: "三十一"},
		{Number: 30, Value: "三十"},
		{Number: 29, Value: "二十九"},
		{Number: 28, Value: "二十八"},
		{Number: 27, Value: "二十七"},
		{Number: 26, Value: "二十六"},
		{Number: 25, Value: "二十五"},
		{Number: 24, Value: "二十四"},
		{Number: 23, Value: "二十三"},
		{Number: 22, Value: "二十二"},
		{Number: 21, Value: "二十一"},
		{Number: 20, Value: "二十"},
		{Number: 19, Value: "十九"},
		{Number: 18, Value: "十八"},
		{Number: 17, Value: "十七"},
		{Number: 16, Value: "十六"},
		{Number: 15, Value: "十五"},
		{Number: 14, Value: "十四"},
		{Number: 13, Value: "十三"},
		{Number: 12, Value: "十二"},
		{Number: 11, Value: "十一"},
		{Number: 10, Value: "十"},
		{Number: 9, Value: "九"},
		{Number: 8, Value: "八"},
		{Number: 7, Value: "七"},
		{Number: 6, Value: "六"},
		{Number: 5, Value: "五"},
		{Number: 4, Value: "四"},
		{Number: 3, Value: "三"},
		{Number: 2, Value: "二"},
		{Number: 1, Value: "一"},
		{Number: 0, Value: "零"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "百"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "第一", Suffix: "第", Masculine: "第一", Feminine: "第一", Neuter: ""},
		{Number: 2, Word: "第二", Suffix: "第", Masculine: "第二", Feminine: "第二", Neuter: ""},
		{Number: 3, Word: "第三", Suffix: "第", Masculine: "第三", Feminine: "第三", Neuter: ""},
		{Number: 4, Word: "第四", Suffix: "第", Masculine: "第四", Feminine: "第四", Neuter: ""},
		{Number: 5, Word: "第五", Suffix: "第", Masculine: "第五", Feminine: "第五", Neuter: ""},
		{Number: 6, Word: "第六", Suffix: "第", Masculine: "第六", Feminine: "第六", Neuter: ""},
		{Number: 7, Word: "第七", Suffix: "第", Masculine: "第七", Feminine: "第七", Neuter: ""},
		{Number: 8, Word: "第八", Suffix: "第", Masculine: "第八", Feminine: "第八", Neuter: ""},
		{Number: 9, Word: "第九", Suffix: "第", Masculine: "第九", Feminine: "第九", Neuter: ""},
		{Number: 10, Word: "第十", Suffix: "第", Masculine: "第十", Feminine: "第十", Neuter: ""},
		{Number: 11, Word: "第十一", Suffix: "第", Masculine: "第十一", Feminine: "第十一", Neuter: ""},
		{Number: 12, Word: "第十二", Suffix: "第", Masculine: "第十二", Feminine: "第十二", Neuter: ""},
		{Number: 13, Word: "第十三", Suffix: "第", Masculine: "第十三", Feminine: "第十三", Neuter: ""},
		{Number: 14, Word: "第十四", Suffix: "第", Masculine: "第十四", Feminine: "第十四", Neuter: ""},
		{Number: 15, Word: "第十五", Suffix: "第", Masculine: "第十五", Feminine: "第十五", Neuter: ""},
		{Number: 16, Word: "第十六", Suffix: "第", Masculine: "第十六", Feminine: "第十六", Neuter: ""},
		{Number: 17, Word: "第十七", Suffix: "第", Masculine: "第十七", Feminine: "第十七", Neuter: ""},
		{Number: 18, Word: "第十八", Suffix: "第", Masculine: "第十八", Feminine: "第十八", Neuter: ""},
		{Number: 19, Word: "第十九", Suffix: "第", Masculine: "第十九", Feminine: "第十九", Neuter: ""},
		{Number: 20, Word: "第二十", Suffix: "第", Masculine: "第二十", Feminine: "第二十", Neuter: ""},
		{Number: 21, Word: "第二十一", Suffix: "第", Masculine: "第二十一", Feminine: "第二十一", Neuter: ""},
		{Number: 30, Word: "第三十", Suffix: "第", Masculine: "第三十", Feminine: "第三十", Neuter: ""},
		{Number: 40, Word: "第四十", Suffix: "第", Masculine: "第四十", Feminine: "第四十", Neuter: ""},
		{Number: 50, Word: "第五十", Suffix: "第", Masculine: "第五十", Feminine: "第五十", Neuter: ""},
		{Number: 60, Word: "第六十", Suffix: "第", Masculine: "第六十", Feminine: "第六十", Neuter: ""},
		{Number: 70, Word: "第七十", Suffix: "第", Masculine: "第七十", Feminine: "第七十", Neuter: ""},
		{Number: 80, Word: "第八十", Suffix: "第", Masculine: "第八十", Feminine: "第八十", Neuter: ""},
		{Number: 90, Word: "第九十", Suffix: "第", Masculine: "第九十", Feminine: "第九十", Neuter: ""},
		{Number: 100, Word: "第一百", Suffix: "第", Masculine: "第一百", Feminine: "第一百", Neuter: ""},
		{Number: 1000, Word: "第一千", Suffix: "第", Masculine: "第一千", Feminine: "第一千", Neuter: ""},
		{Number: 10000, Word: "第一万", Suffix: "第", Masculine: "第一万", Feminine: "第一万", Neuter: ""},
		{Number: 100000000, Word: "第一亿", Suffix: "第", Masculine: "第一亿", Feminine: "第一亿", Neuter: ""},
	},
	LocaleFormatter: &ChineseChinaFormatter{},
}

// ChineseChinaFormatter handles Chinese (China) formatting
type ChineseChinaFormatter struct{}

func (f *ChineseChinaFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *ChineseChinaFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	// Chinese doesn't have plural forms for currency, always use the same form
	return result + currencyPlural
}

func (f *ChineseChinaFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + andText + fractionalWords
}

func (f *ChineseChinaFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	// Chinese doesn't have plural forms for fractions, always use the same form
	return result + fractionPlural
}

func (f *ChineseChinaFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + result
}

func (f *ChineseChinaFormatter) ChopDecimal(value decimal.Decimal, places int) decimal.Decimal {
	return value.Truncate(int32(places))
}

func (f *ChineseChinaFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAsianDecimal(amount)
}

func (f *ChineseChinaFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}
	return FormatAsianCurrency(amount, currencySymbol)
}
