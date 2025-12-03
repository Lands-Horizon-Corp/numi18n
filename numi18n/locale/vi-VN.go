package locale

import (
	"github.com/shopspring/decimal"
)

// VIVNLocale represents the Vietnamese (Vietnam) locale
var VIVNLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Vietnamese Dong",
		Plural:   "đồng",
		Singular: "đồng",
		Symbol:   "₫",
		FractionUnit: FractionUnit{
			Name:     "Xu",
			Plural:   "xu",
			Singular: "xu",
			Symbol:   "xu",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Vietnam",
		Currency:       "VND",
		ISO3166Alpha2:  "VN",
		ISO3166Alpha3:  "VNM",
		ISO3166Numeric: "704",
		Locale:         "vi-VN",
		Timezone:       []string{"Asia/Ho_Chi_Minh"},
		Language:       "vi",
		Emoji:          "🇻🇳",
	},
	Texts: Texts{
		And:   "và",
		Minus: "âm",
		Only:  "chỉ",
		Point: "phẩy",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "một triệu tỷ"},
		{Number: 1000000000000, Value: "một nghìn tỷ"},
		{Number: 1000000000, Value: "một tỷ"},
		{Number: 1000000, Value: "một triệu"},
		{Number: 100000, Value: "một trăm nghìn"},
		{Number: 90000, Value: "chín mươi nghìn"},
		{Number: 80000, Value: "tám mươi nghìn"},
		{Number: 70000, Value: "bảy mươi nghìn"},
		{Number: 60000, Value: "sáu mươi nghìn"},
		{Number: 50000, Value: "năm mươi nghìn"},
		{Number: 40000, Value: "bốn mươi nghìn"},
		{Number: 30000, Value: "ba mươi nghìn"},
		{Number: 20000, Value: "hai mươi nghìn"},
		{Number: 19000, Value: "mười chín nghìn"},
		{Number: 18000, Value: "mười tám nghìn"},
		{Number: 17000, Value: "mười bảy nghìn"},
		{Number: 16000, Value: "mười sáu nghìn"},
		{Number: 15000, Value: "mười lăm nghìn"},
		{Number: 14000, Value: "mười bốn nghìn"},
		{Number: 13000, Value: "mười ba nghìn"},
		{Number: 12000, Value: "mười hai nghìn"},
		{Number: 11000, Value: "mười một nghìn"},
		{Number: 10000, Value: "mười nghìn"},
		{Number: 9000, Value: "chín nghìn"},
		{Number: 8000, Value: "tám nghìn"},
		{Number: 7000, Value: "bảy nghìn"},
		{Number: 6000, Value: "sáu nghìn"},
		{Number: 5000, Value: "năm nghìn"},
		{Number: 4000, Value: "bốn nghìn"},
		{Number: 3000, Value: "ba nghìn"},
		{Number: 2000, Value: "hai nghìn"},
		{Number: 1000, Value: "một nghìn"},
		{Number: 900, Value: "chín trăm"},
		{Number: 800, Value: "tám trăm"},
		{Number: 700, Value: "bảy trăm"},
		{Number: 600, Value: "sáu trăm"},
		{Number: 500, Value: "năm trăm"},
		{Number: 400, Value: "bốn trăm"},
		{Number: 300, Value: "ba trăm"},
		{Number: 200, Value: "hai trăm"},
		{Number: 100, Value: "một trăm"},
		{Number: 99, Value: "chín mươi chín"},
		{Number: 98, Value: "chín mươi tám"},
		{Number: 97, Value: "chín mươi bảy"},
		{Number: 96, Value: "chín mươi sáu"},
		{Number: 95, Value: "chín mươi lăm"},
		{Number: 94, Value: "chín mươi bốn"},
		{Number: 93, Value: "chín mươi ba"},
		{Number: 92, Value: "chín mươi hai"},
		{Number: 91, Value: "chín mươi một"},
		{Number: 90, Value: "chín mươi"},
		{Number: 89, Value: "tám mươi chín"},
		{Number: 88, Value: "tám mươi tám"},
		{Number: 87, Value: "tám mươi bảy"},
		{Number: 86, Value: "tám mươi sáu"},
		{Number: 85, Value: "tám mươi lăm"},
		{Number: 84, Value: "tám mươi bốn"},
		{Number: 83, Value: "tám mươi ba"},
		{Number: 82, Value: "tám mươi hai"},
		{Number: 81, Value: "tám mươi một"},
		{Number: 80, Value: "tám mươi"},
		{Number: 79, Value: "bảy mươi chín"},
		{Number: 78, Value: "bảy mươi tám"},
		{Number: 77, Value: "bảy mươi bảy"},
		{Number: 76, Value: "bảy mươi sáu"},
		{Number: 75, Value: "bảy mươi lăm"},
		{Number: 74, Value: "bảy mươi bốn"},
		{Number: 73, Value: "bảy mươi ba"},
		{Number: 72, Value: "bảy mươi hai"},
		{Number: 71, Value: "bảy mươi một"},
		{Number: 70, Value: "bảy mươi"},
		{Number: 69, Value: "sáu mươi chín"},
		{Number: 68, Value: "sáu mươi tám"},
		{Number: 67, Value: "sáu mươi bảy"},
		{Number: 66, Value: "sáu mươi sáu"},
		{Number: 65, Value: "sáu mươi lăm"},
		{Number: 64, Value: "sáu mươi bốn"},
		{Number: 63, Value: "sáu mươi ba"},
		{Number: 62, Value: "sáu mươi hai"},
		{Number: 61, Value: "sáu mươi một"},
		{Number: 60, Value: "sáu mươi"},
		{Number: 59, Value: "năm mươi chín"},
		{Number: 58, Value: "năm mươi tám"},
		{Number: 57, Value: "năm mươi bảy"},
		{Number: 56, Value: "năm mươi sáu"},
		{Number: 55, Value: "năm mươi lăm"},
		{Number: 54, Value: "năm mươi bốn"},
		{Number: 53, Value: "năm mươi ba"},
		{Number: 52, Value: "năm mươi hai"},
		{Number: 51, Value: "năm mươi một"},
		{Number: 50, Value: "năm mươi"},
		{Number: 49, Value: "bốn mươi chín"},
		{Number: 48, Value: "bốn mươi tám"},
		{Number: 47, Value: "bốn mươi bảy"},
		{Number: 46, Value: "bốn mươi sáu"},
		{Number: 45, Value: "bốn mươi lăm"},
		{Number: 44, Value: "bốn mươi bốn"},
		{Number: 43, Value: "bốn mươi ba"},
		{Number: 42, Value: "bốn mươi hai"},
		{Number: 41, Value: "bốn mươi một"},
		{Number: 40, Value: "bốn mươi"},
		{Number: 39, Value: "ba mươi chín"},
		{Number: 38, Value: "ba mươi tám"},
		{Number: 37, Value: "ba mươi bảy"},
		{Number: 36, Value: "ba mươi sáu"},
		{Number: 35, Value: "ba mươi lăm"},
		{Number: 34, Value: "ba mươi bốn"},
		{Number: 33, Value: "ba mươi ba"},
		{Number: 32, Value: "ba mươi hai"},
		{Number: 31, Value: "ba mươi một"},
		{Number: 30, Value: "ba mươi"},
		{Number: 29, Value: "hai mươi chín"},
		{Number: 28, Value: "hai mươi tám"},
		{Number: 27, Value: "hai mươi bảy"},
		{Number: 26, Value: "hai mươi sáu"},
		{Number: 25, Value: "hai mươi lăm"},
		{Number: 24, Value: "hai mươi bốn"},
		{Number: 23, Value: "hai mươi ba"},
		{Number: 22, Value: "hai mươi hai"},
		{Number: 21, Value: "hai mươi một"},
		{Number: 20, Value: "hai mươi"},
		{Number: 19, Value: "mười chín"},
		{Number: 18, Value: "mười tám"},
		{Number: 17, Value: "mười bảy"},
		{Number: 16, Value: "mười sáu"},
		{Number: 15, Value: "mười lăm"},
		{Number: 14, Value: "mười bốn"},
		{Number: 13, Value: "mười ba"},
		{Number: 12, Value: "mười hai"},
		{Number: 11, Value: "mười một"},
		{Number: 10, Value: "mười"},
		{Number: 9, Value: "chín"},
		{Number: 8, Value: "tám"},
		{Number: 7, Value: "bảy"},
		{Number: 6, Value: "sáu"},
		{Number: 5, Value: "năm"},
		{Number: 4, Value: "bốn"},
		{Number: 3, Value: "ba"},
		{Number: 2, Value: "hai"},
		{Number: 1, Value: "một"},
		{Number: 0, Value: "không"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Trăm"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "thứ nhất", Suffix: ".", Masculine: "thứ nhất", Feminine: "thứ nhất", Neuter: ""},
		{Number: 2, Word: "thứ hai", Suffix: ".", Masculine: "thứ hai", Feminine: "thứ hai", Neuter: ""},
		{Number: 3, Word: "thứ ba", Suffix: ".", Masculine: "thứ ba", Feminine: "thứ ba", Neuter: ""},
		{Number: 4, Word: "thứ tư", Suffix: ".", Masculine: "thứ tư", Feminine: "thứ tư", Neuter: ""},
		{Number: 5, Word: "thứ năm", Suffix: ".", Masculine: "thứ năm", Feminine: "thứ năm", Neuter: ""},
		{Number: 6, Word: "thứ sáu", Suffix: ".", Masculine: "thứ sáu", Feminine: "thứ sáu", Neuter: ""},
		{Number: 7, Word: "thứ bảy", Suffix: ".", Masculine: "thứ bảy", Feminine: "thứ bảy", Neuter: ""},
		{Number: 8, Word: "thứ tám", Suffix: ".", Masculine: "thứ tám", Feminine: "thứ tám", Neuter: ""},
		{Number: 9, Word: "thứ chín", Suffix: ".", Masculine: "thứ chín", Feminine: "thứ chín", Neuter: ""},
		{Number: 10, Word: "thứ mười", Suffix: ".", Masculine: "thứ mười", Feminine: "thứ mười", Neuter: ""},
		{Number: 11, Word: "thứ mười một", Suffix: ".", Masculine: "thứ mười một", Feminine: "thứ mười một", Neuter: ""},
		{Number: 12, Word: "thứ mười hai", Suffix: ".", Masculine: "thứ mười hai", Feminine: "thứ mười hai", Neuter: ""},
		{Number: 13, Word: "thứ mười ba", Suffix: ".", Masculine: "thứ mười ba", Feminine: "thứ mười ba", Neuter: ""},
		{Number: 14, Word: "thứ mười bốn", Suffix: ".", Masculine: "thứ mười bốn", Feminine: "thứ mười bốn", Neuter: ""},
		{Number: 15, Word: "thứ mười lăm", Suffix: ".", Masculine: "thứ mười lăm", Feminine: "thứ mười lăm", Neuter: ""},
		{Number: 16, Word: "thứ mười sáu", Suffix: ".", Masculine: "thứ mười sáu", Feminine: "thứ mười sáu", Neuter: ""},
		{Number: 17, Word: "thứ mười bảy", Suffix: ".", Masculine: "thứ mười bảy", Feminine: "thứ mười bảy", Neuter: ""},
		{Number: 18, Word: "thứ mười tám", Suffix: ".", Masculine: "thứ mười tám", Feminine: "thứ mười tám", Neuter: ""},
		{Number: 19, Word: "thứ mười chín", Suffix: ".", Masculine: "thứ mười chín", Feminine: "thứ mười chín", Neuter: ""},
		{Number: 20, Word: "thứ hai mười", Suffix: ".", Masculine: "thứ hai mười", Feminine: "thứ hai mười", Neuter: ""},
		{Number: 21, Word: "thứ hai mười một", Suffix: ".", Masculine: "thứ hai mười một", Feminine: "thứ hai mười một", Neuter: ""},
		{Number: 30, Word: "thứ ba mười", Suffix: ".", Masculine: "thứ ba mười", Feminine: "thứ ba mười", Neuter: ""},
		{Number: 40, Word: "thứ bốn mười", Suffix: ".", Masculine: "thứ bốn mười", Feminine: "thứ bốn mười", Neuter: ""},
		{Number: 50, Word: "thứ năm mười", Suffix: ".", Masculine: "thứ năm mười", Feminine: "thứ năm mười", Neuter: ""},
		{Number: 60, Word: "thứ sáu mười", Suffix: ".", Masculine: "thứ sáu mười", Feminine: "thứ sáu mười", Neuter: ""},
		{Number: 70, Word: "thứ bảy mười", Suffix: ".", Masculine: "thứ bảy mười", Feminine: "thứ bảy mười", Neuter: ""},
		{Number: 80, Word: "thứ tám mười", Suffix: ".", Masculine: "thứ tám mười", Feminine: "thứ tám mười", Neuter: ""},
		{Number: 90, Word: "thứ chín mười", Suffix: ".", Masculine: "thứ chín mười", Feminine: "thứ chín mười", Neuter: ""},
		{Number: 100, Word: "thứ một trăm", Suffix: ".", Masculine: "thứ một trăm", Feminine: "thứ một trăm", Neuter: ""},
		{Number: 1000, Word: "thứ một nghìn", Suffix: ".", Masculine: "thứ một nghìn", Feminine: "thứ một nghìn", Neuter: ""},
		{Number: 1000000, Word: "thứ một triệu", Suffix: ".", Masculine: "thứ một triệu", Feminine: "thứ một triệu", Neuter: ""},
		{Number: 1000000000, Word: "thứ một tỷ", Suffix: ".", Masculine: "thứ một tỷ", Feminine: "thứ một tỷ", Neuter: ""},
	},
	LocaleFormatter: &VietnameseFormatter{},
}

// VietnameseFormatter handles Vietnamese (Vietnam) formatting
type VietnameseFormatter struct{}

func (f *VietnameseFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *VietnameseFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *VietnameseFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *VietnameseFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *VietnameseFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *VietnameseFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	return amount.Truncate(int32(precision))
}

func (f *VietnameseFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *VietnameseFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
