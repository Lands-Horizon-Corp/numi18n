package locale

import (
	"github.com/shopspring/decimal"
)

// ZH001Locale is a NumI18NLocale configured for Chinese (World) - zh-001
var ZH001Locale = NumI18NLocale{
	Currency: Currency{
		Name:     "元",
		Plural:   "元",
		Singular: "元",
		Symbol:   "¥",
		FractionUnit: FractionUnit{
			Name:     "分",
			Plural:   "分",
			Singular: "分",
			Symbol:   "¢",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "World",
		Currency:       "USD",
		ISO3166Alpha2:  "001",
		ISO3166Alpha3:  "WLD",
		ISO3166Numeric: "001",
		Locale:         "zh-001",
		Timezone:       []string{"UTC"},
		Language:       "zh",
		Emoji:          "🌍",
		PhoneCode:      "+1",
		Domain:         ".com",
	},
	Texts: Texts{
		And:   "和",
		Minus: "负",
		Only:  "仅",
		Point: "点",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "千万亿"}, // Quadrillion
		{Number: 1000000000000, Value: "万亿"},     // Trillion
		{Number: 1000000000, Value: "十亿"},        // Billion
		{Number: 1000000, Value: "百万"},           // Million
		{Number: 1000, Value: "千"},               // Thousand
		{Number: 100, Value: "百"},                // Hundred
		{Number: 90, Value: "九十"},
		{Number: 80, Value: "八十"},
		{Number: 70, Value: "七十"},
		{Number: 60, Value: "六十"},
		{Number: 50, Value: "五十"},
		{Number: 40, Value: "四十"},
		{Number: 30, Value: "三十"},
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
		{Number: 100, Value: "一百"},
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
		{Number: 1000000, Word: "第一百万", Suffix: "第", Masculine: "第一百万", Feminine: "第一百万", Neuter: ""},
		{Number: 1000000000, Word: "第十亿", Suffix: "第", Masculine: "第十亿", Feminine: "第十亿", Neuter: ""},
	},
	LocaleFormatter: &ChineseFormatter{},
}

// ChineseFormatter handles Chinese-specific formatting
type ChineseFormatter struct{}

func (f *ChineseFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *ChineseFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	// Chinese doesn't have plural forms for currency, always use the same form
	return result + currencyPlural
}

func (f *ChineseFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + andText + fractionalWords
}

func (f *ChineseFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	// Chinese doesn't have plural forms for fractions, always use the same form
	return result + fractionPlural
}

func (f *ChineseFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + result
}

func (f *ChineseFormatter) ChopDecimal(value decimal.Decimal, places int) decimal.Decimal {
	return value.Truncate(int32(places))
}

func (f *ChineseFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAsianDecimal(amount)
}
func (f *ChineseFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Asian conventions (no separators, period decimal, prefix symbol)
	return FormatAsianCurrency(amount, currencySymbol)
}
