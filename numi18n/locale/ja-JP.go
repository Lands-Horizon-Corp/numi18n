package locale

import (
	"github.com/shopspring/decimal"
)

// JPLocale is a NumI18NLocale configured for Japan (ja-JP)
var JPLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "円",
		Plural:   "円",
		Singular: "円",
		Symbol:   "¥",
		FractionUnit: FractionUnit{
			Name:     "銭",
			Plural:   "銭",
			Singular: "銭",
			Symbol:   "銭",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Japan",
		Currency:       "JPY",
		ISO3166Alpha2:  "JP",
		ISO3166Alpha3:  "JPN",
		ISO3166Numeric: "392",
		Locale:         "ja-JP",
		Timezone:       []string{"Asia/Tokyo"},
		Language:       "ja",
		Emoji:          "🇯🇵",
		PhoneCode:      "+81",
		Domain:         ".jp",
	},
	Texts: Texts{
		And:   "と",
		Minus: "マイナス",
		Only:  "のみ",
		Point: "点",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "一千兆"},
		{Number: 1000000000000, Value: "一兆"},
		{Number: 100000000, Value: "一億"},
		{Number: 10000000, Value: "一千万"},
		{Number: 1000000, Value: "百万"},
		{Number: 100000, Value: "十万"},
		{Number: 10000, Value: "一万"},
		{Number: 1000, Value: "一千"},
		{Number: 100, Value: "百"},
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
		{Number: 100000000, Value: "一億"},
		{Number: 90000, Value: "九万"},
		{Number: 80000, Value: "八万"},
		{Number: 70000, Value: "七万"},
		{Number: 60000, Value: "六万"},
		{Number: 50000, Value: "五万"},
		{Number: 40000, Value: "四万"},
		{Number: 30000, Value: "三万"},
		{Number: 20000, Value: "二万"},
		{Number: 10000, Value: "一万"},
		{Number: 9000, Value: "九千"},
		{Number: 8000, Value: "八千"},
		{Number: 7000, Value: "七千"},
		{Number: 6000, Value: "六千"},
		{Number: 5000, Value: "五千"},
		{Number: 4000, Value: "四千"},
		{Number: 3000, Value: "三千"},
		{Number: 2000, Value: "二千"},
		{Number: 1000, Value: "千"},
		{Number: 900, Value: "九百"},
		{Number: 800, Value: "八百"},
		{Number: 700, Value: "七百"},
		{Number: 600, Value: "六百"},
		{Number: 500, Value: "五百"},
		{Number: 400, Value: "四百"},
		{Number: 300, Value: "三百"},
		{Number: 200, Value: "二百"},
		{Number: 100, Value: "百"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "第一", Suffix: "番目", Masculine: "第一", Feminine: "第一", Neuter: "第一"},
		{Number: 2, Word: "第二", Suffix: "番目", Masculine: "第二", Feminine: "第二", Neuter: "第二"},
		{Number: 3, Word: "第三", Suffix: "番目", Masculine: "第三", Feminine: "第三", Neuter: "第三"},
		{Number: 4, Word: "第四", Suffix: "番目", Masculine: "第四", Feminine: "第四", Neuter: "第四"},
		{Number: 5, Word: "第五", Suffix: "番目", Masculine: "第五", Feminine: "第五", Neuter: "第五"},
		{Number: 6, Word: "第六", Suffix: "番目", Masculine: "第六", Feminine: "第六", Neuter: "第六"},
		{Number: 7, Word: "第七", Suffix: "番目", Masculine: "第七", Feminine: "第七", Neuter: "第七"},
		{Number: 8, Word: "第八", Suffix: "番目", Masculine: "第八", Feminine: "第八", Neuter: "第八"},
		{Number: 9, Word: "第九", Suffix: "番目", Masculine: "第九", Feminine: "第九", Neuter: "第九"},
		{Number: 10, Word: "第十", Suffix: "番目", Masculine: "第十", Feminine: "第十", Neuter: "第十"},
		{Number: 11, Word: "第十一", Suffix: "番目", Masculine: "第十一", Feminine: "第十一", Neuter: "第十一"},
		{Number: 12, Word: "第十二", Suffix: "番目", Masculine: "第十二", Feminine: "第十二", Neuter: "第十二"},
		{Number: 13, Word: "第十三", Suffix: "番目", Masculine: "第十三", Feminine: "第十三", Neuter: "第十三"},
		{Number: 14, Word: "第十四", Suffix: "番目", Masculine: "第十四", Feminine: "第十四", Neuter: "第十四"},
		{Number: 15, Word: "第十五", Suffix: "番目", Masculine: "第十五", Feminine: "第十五", Neuter: "第十五"},
		{Number: 16, Word: "第十六", Suffix: "番目", Masculine: "第十六", Feminine: "第十六", Neuter: "第十六"},
		{Number: 17, Word: "第十七", Suffix: "番目", Masculine: "第十七", Feminine: "第十七", Neuter: "第十七"},
		{Number: 18, Word: "第十八", Suffix: "番目", Masculine: "第十八", Feminine: "第十八", Neuter: "第十八"},
		{Number: 19, Word: "第十九", Suffix: "番目", Masculine: "第十九", Feminine: "第十九", Neuter: "第十九"},
		{Number: 20, Word: "第二十", Suffix: "番目", Masculine: "第二十", Feminine: "第二十", Neuter: "第二十"},
		{Number: 21, Word: "第二十一", Suffix: "番目", Masculine: "第二十一", Feminine: "第二十一", Neuter: "第二十一"},
		{Number: 30, Word: "第三十", Suffix: "番目", Masculine: "第三十", Feminine: "第三十", Neuter: "第三十"},
		{Number: 40, Word: "第四十", Suffix: "番目", Masculine: "第四十", Feminine: "第四十", Neuter: "第四十"},
		{Number: 50, Word: "第五十", Suffix: "番目", Masculine: "第五十", Feminine: "第五十", Neuter: "第五十"},
		{Number: 60, Word: "第六十", Suffix: "番目", Masculine: "第六十", Feminine: "第六十", Neuter: "第六十"},
		{Number: 70, Word: "第七十", Suffix: "番目", Masculine: "第七十", Feminine: "第七十", Neuter: "第七十"},
		{Number: 80, Word: "第八十", Suffix: "番目", Masculine: "第八十", Feminine: "第八十", Neuter: "第八十"},
		{Number: 90, Word: "第九十", Suffix: "番目", Masculine: "第九十", Feminine: "第九十", Neuter: "第九十"},
		{Number: 100, Word: "第百", Suffix: "番目", Masculine: "第百", Feminine: "第百", Neuter: "第百"},
		{Number: 1000, Word: "第千", Suffix: "番目", Masculine: "第千", Feminine: "第千", Neuter: "第千"},
	},
	LocaleFormatter: &JapaneseFormatter{},
}

// JapaneseFormatter handles Japanese-specific formatting
type JapaneseFormatter struct{}

func (f *JapaneseFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	// First check exact mappings
	exactResult := ConvertToWordsWithExactMappingInt64(number, targetLocale)
	if exactResult != ConvertToWordsGenericInt64(number, targetLocale) {
		return exactResult // Found exact mapping
	}

	// If no exact mapping, use the custom Japanese logic
	decimalNumber := decimal.NewFromInt(number)
	if decimalNumber.Equal(decimal.Zero) {
		return GetWordForNumber(decimal.Zero, targetLocale)
	}

	result := ""
	oku := decimal.NewFromInt(100000000) // 億 (oku)
	man := decimal.NewFromInt(10000)     // 万 (man)

	// Handle large numbers first
	result, decimalNumber = f.handleOku(result, decimalNumber, oku, targetLocale)
	result, decimalNumber = f.handleMan(result, decimalNumber, man, targetLocale)
	result = f.handleRemainingDigits(result, decimalNumber, targetLocale)

	return result
}

func (f *JapaneseFormatter) handleOku(result string, decimalNumber, oku decimal.Decimal, targetLocale NumI18NLocale) (string, decimal.Decimal) {
	if decimalNumber.GreaterThanOrEqual(oku) {
		okuPart := decimalNumber.Div(oku).Floor()
		if okuPart.GreaterThan(decimal.Zero) {
			result += f.FormatNumber(okuPart.IntPart(), targetLocale) + GetWordForNumber(oku, targetLocale)
		}
		decimalNumber = decimalNumber.Mod(oku)
	}
	return result, decimalNumber
}

func (f *JapaneseFormatter) handleMan(result string, decimalNumber, man decimal.Decimal, targetLocale NumI18NLocale) (string, decimal.Decimal) {
	if decimalNumber.GreaterThanOrEqual(man) {
		manPart := decimalNumber.Div(man).Floor()
		if manPart.GreaterThan(decimal.Zero) {
			if manPart.Equal(decimal.NewFromInt(1)) {
				result += "万" // Just "万" for 10000
			} else {
				// For complex numbers like 123万, we should get "百二十三万", not "百二十三一万"
				manWords := f.formatNumberRecursive(manPart.IntPart(), targetLocale)
				result += manWords + "万"
			}
		}
		decimalNumber = decimalNumber.Mod(man)
	}
	return result, decimalNumber
}

func (f *JapaneseFormatter) handleRemainingDigits(result string, decimalNumber decimal.Decimal, targetLocale NumI18NLocale) string {
	thousand := decimal.NewFromInt(1000) // 千 (sen)
	hundred := decimal.NewFromInt(100)   // 百 (hyaku)
	ten := decimal.NewFromInt(10)        // 十 (ju)

	result = f.handleThousands(result, decimalNumber, thousand, targetLocale)
	decimalNumber = decimalNumber.Mod(thousand)

	result = f.handleHundreds(result, decimalNumber, hundred, targetLocale)
	decimalNumber = decimalNumber.Mod(hundred)

	result = f.handleTens(result, decimalNumber, ten, targetLocale)
	decimalNumber = decimalNumber.Mod(ten)

	result = f.handleOnes(result, decimalNumber, targetLocale)

	return result
}

func (f *JapaneseFormatter) handleThousands(result string, decimalNumber, thousand decimal.Decimal, targetLocale NumI18NLocale) string {
	if decimalNumber.GreaterThanOrEqual(thousand) {
		thousandsPart := decimalNumber.Div(thousand).Floor()
		if thousandsPart.GreaterThan(decimal.Zero) {
			if thousandsPart.Equal(decimal.NewFromInt(1)) {
				result += "千" // Just "千" for 1000
			} else {
				result += f.formatNumberRecursive(thousandsPart.IntPart(), targetLocale) + "千"
			}
		}
	}
	return result
}

func (f *JapaneseFormatter) handleHundreds(result string, decimalNumber, hundred decimal.Decimal, targetLocale NumI18NLocale) string {
	if decimalNumber.GreaterThanOrEqual(hundred) {
		hundredsPart := decimalNumber.Div(hundred).Floor()
		if hundredsPart.GreaterThan(decimal.Zero) {
			if hundredsPart.Equal(decimal.NewFromInt(1)) {
				result += "百" // Just "百" for 100
			} else {
				result += f.formatNumberRecursive(hundredsPart.IntPart(), targetLocale) + "百"
			}
		}
	}
	return result
}

func (f *JapaneseFormatter) handleTens(result string, decimalNumber, ten decimal.Decimal, targetLocale NumI18NLocale) string {
	if decimalNumber.GreaterThanOrEqual(ten) {
		tensPart := decimalNumber.Div(ten).Floor()
		if tensPart.GreaterThan(decimal.Zero) {
			if tensPart.Equal(decimal.NewFromInt(1)) {
				result += "十" // Just "十" for 10
			} else {
				result += f.formatNumberRecursive(tensPart.IntPart(), targetLocale) + "十"
			}
		}
	}
	return result
}

func (f *JapaneseFormatter) handleOnes(result string, decimalNumber decimal.Decimal, targetLocale NumI18NLocale) string {
	if decimalNumber.GreaterThan(decimal.Zero) {
		result += GetWordForNumber(decimalNumber, targetLocale)
	}
	return result
}

// formatNumberRecursive handles recursive formatting for Japanese numbers without causing infinite recursion
func (f *JapaneseFormatter) formatNumberRecursive(number int64, targetLocale NumI18NLocale) string {
	decimalNumber := decimal.NewFromInt(number)
	if decimalNumber.Equal(decimal.Zero) {
		return GetWordForNumber(decimal.Zero, targetLocale)
	}

	result := ""
	thousand := decimal.NewFromInt(1000) // 千 (sen)
	hundred := decimal.NewFromInt(100)   // 百 (hyaku)
	ten := decimal.NewFromInt(10)        // 十 (ju)

	// Handle thousands (1000-9999)
	if decimalNumber.GreaterThanOrEqual(thousand) {
		thousandsPart := decimalNumber.Div(thousand).Floor()
		if thousandsPart.GreaterThan(decimal.Zero) {
			if thousandsPart.Equal(decimal.NewFromInt(1)) {
				result += "千" // Just "千" for 1000
			} else {
				result += GetWordForNumber(thousandsPart, targetLocale) + "千"
			}
		}
		decimalNumber = decimalNumber.Mod(thousand)
	}

	// Handle hundreds (100-999)
	if decimalNumber.GreaterThanOrEqual(hundred) {
		hundredsPart := decimalNumber.Div(hundred).Floor()
		if hundredsPart.GreaterThan(decimal.Zero) {
			if hundredsPart.Equal(decimal.NewFromInt(1)) {
				result += "百" // Just "百" for 100
			} else {
				result += GetWordForNumber(hundredsPart, targetLocale) + "百"
			}
		}
		decimalNumber = decimalNumber.Mod(hundred)
	}

	// Handle tens (10-99)
	if decimalNumber.GreaterThanOrEqual(ten) {
		tensPart := decimalNumber.Div(ten).Floor()
		if tensPart.GreaterThan(decimal.Zero) {
			if tensPart.Equal(decimal.NewFromInt(1)) {
				result += "十" // Just "十" for 10
			} else {
				result += GetWordForNumber(tensPart, targetLocale) + "十"
			}
		}
		decimalNumber = decimalNumber.Mod(ten)
	}

	// Handle ones (1-9)
	if decimalNumber.GreaterThan(decimal.Zero) {
		result += GetWordForNumber(decimalNumber, targetLocale)
	}

	return result
}

func (f *JapaneseFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + currencyName
	}
	return result + currencyPlural
}

func (f *JapaneseFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + andText + fractionalWords
}

func (f *JapaneseFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + fractionName
	}
	return result + fractionPlural
}

func (f *JapaneseFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + result
}

func (f *JapaneseFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 0
	}
	return amount.Truncate(int32(precision))
}

func (f *JapaneseFormatter) FormatDecimalNumber(amount float64) string {
	return FormatDecimalWithoutSeparators(amount)
}

func (f *JapaneseFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Japanese currency style (no separators, no decimals)
	return FormatJapaneseCurrency(amount, currencySymbol)
}
