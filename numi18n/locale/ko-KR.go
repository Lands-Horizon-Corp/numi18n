package locale

import (
	"github.com/shopspring/decimal"
)

// KRLocale is a NumI18NLocale configured for South Korea (ko-KR)
var KRLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "원",
		Plural:   "원",
		Singular: "원",
		Symbol:   "₩",
		FractionUnit: FractionUnit{
			Name:     "전",
			Plural:   "전",
			Singular: "전",
			Symbol:   "전",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "South Korea",
		Currency:       "KRW",
		ISO3166Alpha2:  "KR",
		ISO3166Alpha3:  "KOR",
		ISO3166Numeric: "410",
		Locale:         "ko-KR",
		Timezone:       []string{"Asia/Seoul"},
		Language:       "ko",
		Emoji:          "🇰🇷",
		PhoneCode:      "+82",
		Domain:         ".kr",
	},
	Texts: Texts{
		And:   "그리고",
		Minus: "마이너스",
		Only:  "만",
		Point: "점",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "천조"},
		{Number: 1000000000000, Value: "일조"},
		{Number: 100000000000, Value: "천억"},
		{Number: 10000000000, Value: "백억"},
		{Number: 1000000000, Value: "십억"},
		{Number: 100000000, Value: "일억"},
		{Number: 10000000, Value: "천만"},
		{Number: 1000000, Value: "백만"},
		{Number: 100000, Value: "십만"},
		{Number: 10000, Value: "일만"},
		{Number: 1000, Value: "일천"},
		{Number: 100, Value: "일백"},
		{Number: 90, Value: "구십"},
		{Number: 80, Value: "팔십"},
		{Number: 70, Value: "칠십"},
		{Number: 60, Value: "육십"},
		{Number: 50, Value: "오십"},
		{Number: 40, Value: "사십"},
		{Number: 30, Value: "삼십"},
		{Number: 20, Value: "이십"},
		{Number: 19, Value: "십구"},
		{Number: 18, Value: "십팔"},
		{Number: 17, Value: "십칠"},
		{Number: 16, Value: "십육"},
		{Number: 15, Value: "십오"},
		{Number: 14, Value: "십사"},
		{Number: 13, Value: "십삼"},
		{Number: 12, Value: "십이"},
		{Number: 11, Value: "십일"},
		{Number: 10, Value: "십"},
		{Number: 9, Value: "구"},
		{Number: 8, Value: "팔"},
		{Number: 7, Value: "칠"},
		{Number: 6, Value: "육"},
		{Number: 5, Value: "오"},
		{Number: 4, Value: "사"},
		{Number: 3, Value: "삼"},
		{Number: 2, Value: "이"},
		{Number: 1, Value: "일"},
		{Number: 0, Value: "영"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "일백"},
		{Number: 10000, Value: "일만"},
		{Number: 100000000, Value: "일억"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "첫째", Suffix: "번째", Masculine: "첫째", Feminine: "첫째", Neuter: "첫째"},
		{Number: 2, Word: "둘째", Suffix: "번째", Masculine: "둘째", Feminine: "둘째", Neuter: "둘째"},
		{Number: 3, Word: "셋째", Suffix: "번째", Masculine: "셋째", Feminine: "셋째", Neuter: "셋째"},
		{Number: 4, Word: "넷째", Suffix: "번째", Masculine: "넷째", Feminine: "넷째", Neuter: "넷째"},
		{Number: 5, Word: "다섯째", Suffix: "번째", Masculine: "다섯째", Feminine: "다섯째", Neuter: "다섯째"},
		{Number: 6, Word: "여섯째", Suffix: "번째", Masculine: "여섯째", Feminine: "여섯째", Neuter: "여섯째"},
		{Number: 7, Word: "일곱째", Suffix: "번째", Masculine: "일곱째", Feminine: "일곱째", Neuter: "일곱째"},
		{Number: 8, Word: "여덟째", Suffix: "번째", Masculine: "여덟째", Feminine: "여덟째", Neuter: "여덟째"},
		{Number: 9, Word: "아홉째", Suffix: "번째", Masculine: "아홉째", Feminine: "아홉째", Neuter: "아홉째"},
		{Number: 10, Word: "열째", Suffix: "번째", Masculine: "열째", Feminine: "열째", Neuter: "열째"},
		{Number: 11, Word: "열한째", Suffix: "번째", Masculine: "열한째", Feminine: "열한째", Neuter: "열한째"},
		{Number: 12, Word: "열두째", Suffix: "번째", Masculine: "열두째", Feminine: "열두째", Neuter: "열두째"},
		{Number: 13, Word: "열셋째", Suffix: "번째", Masculine: "열셋째", Feminine: "열셋째", Neuter: "열셋째"},
		{Number: 14, Word: "열넷째", Suffix: "번째", Masculine: "열넷째", Feminine: "열넷째", Neuter: "열넷째"},
		{Number: 15, Word: "열다섯째", Suffix: "번째", Masculine: "열다섯째", Feminine: "열다섯째", Neuter: "열다섯째"},
		{Number: 16, Word: "열여섯째", Suffix: "번째", Masculine: "열여섯째", Feminine: "열여섯째", Neuter: "열여섯째"},
		{Number: 17, Word: "열일곱째", Suffix: "번째", Masculine: "열일곱째", Feminine: "열일곱째", Neuter: "열일곱째"},
		{Number: 18, Word: "열여덟째", Suffix: "번째", Masculine: "열여덟째", Feminine: "열여덟째", Neuter: "열여덟째"},
		{Number: 19, Word: "열아홉째", Suffix: "번째", Masculine: "열아홉째", Feminine: "열아홉째", Neuter: "열아홉째"},
		{Number: 20, Word: "스무째", Suffix: "번째", Masculine: "스무째", Feminine: "스무째", Neuter: "스무째"},
		{Number: 21, Word: "스물한째", Suffix: "번째", Masculine: "스물한째", Feminine: "스물한째", Neuter: "스물한째"},
		{Number: 30, Word: "서른째", Suffix: "번째", Masculine: "서른째", Feminine: "서른째", Neuter: "서른째"},
		{Number: 40, Word: "마흔째", Suffix: "번째", Masculine: "마흔째", Feminine: "마흔째", Neuter: "마흔째"},
		{Number: 50, Word: "쉰째", Suffix: "번째", Masculine: "쉰째", Feminine: "쉰째", Neuter: "쉰째"},
		{Number: 60, Word: "예순째", Suffix: "번째", Masculine: "예순째", Feminine: "예순째", Neuter: "예순째"},
		{Number: 70, Word: "일흔째", Suffix: "번째", Masculine: "일흔째", Feminine: "일흔째", Neuter: "일흔째"},
		{Number: 80, Word: "여든째", Suffix: "번째", Masculine: "여든째", Feminine: "여든째", Neuter: "여든째"},
		{Number: 90, Word: "아흔째", Suffix: "번째", Masculine: "아흔째", Feminine: "아흔째", Neuter: "아흔째"},
		{Number: 100, Word: "백째", Suffix: "번째", Masculine: "백째", Feminine: "백째", Neuter: "백째"},
		{Number: 1000, Word: "천째", Suffix: "번째", Masculine: "천째", Feminine: "천째", Neuter: "천째"},
	},
	LocaleFormatter: &KoreanFormatter{},
}

// KoreanFormatter handles Korean (ko-KR) formatting
type KoreanFormatter struct{}

func (f *KoreanFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *KoreanFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	// Korean doesn't distinguish between singular and plural for currency
	return result + " " + currencyName
}

func (f *KoreanFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *KoreanFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	// Korean doesn't distinguish between singular and plural for fraction units
	return result + " " + fractionName
}

func (f *KoreanFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *KoreanFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 2
	}
	return amount.Truncate(int32(precision))
}

func (f *KoreanFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAsianDecimal(amount)
}

func (f *KoreanFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}
	return FormatAsianCurrency(amount, currencySymbol)
}
