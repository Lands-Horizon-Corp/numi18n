package locale

import (
	"github.com/shopspring/decimal"
)

// THTHLocale represents the Thai (Thailand) locale
var THTHLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Thai Baht",
		Plural:   "บาท",
		Singular: "บาท",
		Symbol:   "฿",
		FractionUnit: FractionUnit{
			Name:     "Satang",
			Plural:   "สตางค์",
			Singular: "สตางค์",
			Symbol:   "ส.",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Thailand",
		Currency:       "THB",
		ISO3166Alpha2:  "TH",
		ISO3166Alpha3:  "THA",
		ISO3166Numeric: "764",
		Locale:         "th-TH",
		Timezone:       []string{"Asia/Bangkok"},
		Language:       "th",
		Emoji:          "🇹🇭",
	},
	Texts: Texts{
		And:   "และ",
		Minus: "ลบ",
		Only:  "เท่านั้น",
		Point: "จุด",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "หนึ่งพันล้านล้าน"},
		{Number: 1000000000000, Value: "หนึ่งล้านล้าน"},
		{Number: 1000000000, Value: "หนึ่งพันล้าน"},
		{Number: 1000000, Value: "หนึ่งล้าน"},
		{Number: 100000, Value: "หนึ่งแสน"},
		{Number: 90000, Value: "เก้าหมื่น"},
		{Number: 80000, Value: "แปดหมื่น"},
		{Number: 70000, Value: "เจ็ดหมื่น"},
		{Number: 60000, Value: "หกหมื่น"},
		{Number: 50000, Value: "ห้าหมื่น"},
		{Number: 40000, Value: "สี่หมื่น"},
		{Number: 30000, Value: "สามหมื่น"},
		{Number: 20000, Value: "สองหมื่น"},
		{Number: 19000, Value: "หนึ่งหมื่นเก้าพัน"},
		{Number: 18000, Value: "หนึ่งหมื่นแปดพัน"},
		{Number: 17000, Value: "หนึ่งหมื่นเจ็ดพัน"},
		{Number: 16000, Value: "หนึ่งหมื่นหกพัน"},
		{Number: 15000, Value: "หนึ่งหมื่นห้าพัน"},
		{Number: 14000, Value: "หนึ่งหมื่นสี่พัน"},
		{Number: 13000, Value: "หนึ่งหมื่นสามพัน"},
		{Number: 12000, Value: "หนึ่งหมื่นสองพัน"},
		{Number: 11000, Value: "หนึ่งหมื่นหนึ่งพัน"},
		{Number: 10000, Value: "หนึ่งหมื่น"},
		{Number: 9000, Value: "เก้าพัน"},
		{Number: 8000, Value: "แปดพัน"},
		{Number: 7000, Value: "เจ็ดพัน"},
		{Number: 6000, Value: "หกพัน"},
		{Number: 5000, Value: "ห้าพัน"},
		{Number: 4000, Value: "สี่พัน"},
		{Number: 3000, Value: "สามพัน"},
		{Number: 2000, Value: "สองพัน"},
		{Number: 1000, Value: "หนึ่งพัน"},
		{Number: 900, Value: "เก้าร้อย"},
		{Number: 800, Value: "แปดร้อย"},
		{Number: 700, Value: "เจ็ดร้อย"},
		{Number: 600, Value: "หกร้อย"},
		{Number: 500, Value: "ห้าร้อย"},
		{Number: 400, Value: "สี่ร้อย"},
		{Number: 300, Value: "สามร้อย"},
		{Number: 200, Value: "สองร้อย"},
		{Number: 100, Value: "หนึ่งร้อย"},
		{Number: 99, Value: "เก้าสิบเก้า"},
		{Number: 98, Value: "เก้าสิบแปด"},
		{Number: 97, Value: "เก้าสิบเจ็ด"},
		{Number: 96, Value: "เก้าสิบหก"},
		{Number: 95, Value: "เก้าสิบห้า"},
		{Number: 94, Value: "เก้าสิบสี่"},
		{Number: 93, Value: "เก้าสิบสาม"},
		{Number: 92, Value: "เก้าสิบสอง"},
		{Number: 91, Value: "เก้าสิบเอ็ด"},
		{Number: 90, Value: "เก้าสิบ"},
		{Number: 89, Value: "แปดสิบเก้า"},
		{Number: 88, Value: "แปดสิบแปด"},
		{Number: 87, Value: "แปดสิบเจ็ด"},
		{Number: 86, Value: "แปดสิบหก"},
		{Number: 85, Value: "แปดสิบห้า"},
		{Number: 84, Value: "แปดสิบสี่"},
		{Number: 83, Value: "แปดสิบสาม"},
		{Number: 82, Value: "แปดสิบสอง"},
		{Number: 81, Value: "แปดสิบเอ็ด"},
		{Number: 80, Value: "แปดสิบ"},
		{Number: 79, Value: "เจ็ดสิบเก้า"},
		{Number: 78, Value: "เจ็ดสิบแปด"},
		{Number: 77, Value: "เจ็ดสิบเจ็ด"},
		{Number: 76, Value: "เจ็ดสิบหก"},
		{Number: 75, Value: "เจ็ดสิบห้า"},
		{Number: 74, Value: "เจ็ดสิบสี่"},
		{Number: 73, Value: "เจ็ดสิบสาม"},
		{Number: 72, Value: "เจ็ดสิบสอง"},
		{Number: 71, Value: "เจ็ดสิบเอ็ด"},
		{Number: 70, Value: "เจ็ดสิบ"},
		{Number: 69, Value: "หกสิบเก้า"},
		{Number: 68, Value: "หกสิบแปด"},
		{Number: 67, Value: "หกสิบเจ็ด"},
		{Number: 66, Value: "หกสิบหก"},
		{Number: 65, Value: "หกสิบห้า"},
		{Number: 64, Value: "หกสิบสี่"},
		{Number: 63, Value: "หกสิบสาม"},
		{Number: 62, Value: "หกสิบสอง"},
		{Number: 61, Value: "หกสิบเอ็ด"},
		{Number: 60, Value: "หกสิบ"},
		{Number: 59, Value: "ห้าสิบเก้า"},
		{Number: 58, Value: "ห้าสิบแปด"},
		{Number: 57, Value: "ห้าสิบเจ็ด"},
		{Number: 56, Value: "ห้าสิบหก"},
		{Number: 55, Value: "ห้าสิบห้า"},
		{Number: 54, Value: "ห้าสิบสี่"},
		{Number: 53, Value: "ห้าสิบสาม"},
		{Number: 52, Value: "ห้าสิบสอง"},
		{Number: 51, Value: "ห้าสิบเอ็ด"},
		{Number: 50, Value: "ห้าสิบ"},
		{Number: 49, Value: "สี่สิบเก้า"},
		{Number: 48, Value: "สี่สิบแปด"},
		{Number: 47, Value: "สี่สิบเจ็ด"},
		{Number: 46, Value: "สี่สิบหก"},
		{Number: 45, Value: "สี่สิบห้า"},
		{Number: 44, Value: "สี่สิบสี่"},
		{Number: 43, Value: "สี่สิบสาม"},
		{Number: 42, Value: "สี่สิบสอง"},
		{Number: 41, Value: "สี่สิบเอ็ด"},
		{Number: 40, Value: "สี่สิบ"},
		{Number: 39, Value: "สามสิบเก้า"},
		{Number: 38, Value: "สามสิบแปด"},
		{Number: 37, Value: "สามสิบเจ็ด"},
		{Number: 36, Value: "สามสิบหก"},
		{Number: 35, Value: "สามสิบห้า"},
		{Number: 34, Value: "สามสิบสี่"},
		{Number: 33, Value: "สามสิบสาม"},
		{Number: 32, Value: "สามสิบสอง"},
		{Number: 31, Value: "สามสิบเอ็ด"},
		{Number: 30, Value: "สามสิบ"},
		{Number: 29, Value: "ยี่สิบเก้า"},
		{Number: 28, Value: "ยี่สิบแปด"},
		{Number: 27, Value: "ยี่สิบเจ็ด"},
		{Number: 26, Value: "ยี่สิบหก"},
		{Number: 25, Value: "ยี่สิบห้า"},
		{Number: 24, Value: "ยี่สิบสี่"},
		{Number: 23, Value: "ยี่สิบสาม"},
		{Number: 22, Value: "ยี่สิบสอง"},
		{Number: 21, Value: "ยี่สิบเอ็ด"},
		{Number: 20, Value: "ยี่สิบ"},
		{Number: 19, Value: "สิบเก้า"},
		{Number: 18, Value: "สิบแปด"},
		{Number: 17, Value: "สิบเจ็ด"},
		{Number: 16, Value: "สิบหก"},
		{Number: 15, Value: "สิบห้า"},
		{Number: 14, Value: "สิบสี่"},
		{Number: 13, Value: "สิบสาม"},
		{Number: 12, Value: "สิบสอง"},
		{Number: 11, Value: "สิบเอ็ด"},
		{Number: 10, Value: "สิบ"},
		{Number: 9, Value: "เก้า"},
		{Number: 8, Value: "แปด"},
		{Number: 7, Value: "เจ็ด"},
		{Number: 6, Value: "หก"},
		{Number: 5, Value: "ห้า"},
		{Number: 4, Value: "สี่"},
		{Number: 3, Value: "สาม"},
		{Number: 2, Value: "สอง"},
		{Number: 1, Value: "หนึ่ง"},
		{Number: 0, Value: "ศูนย์"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "หนึ่งร้อย"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "ที่หนึ่ง", Suffix: ".", Masculine: "ที่หนึ่ง", Feminine: "ที่หนึ่ง", Neuter: "ที่หนึ่ง"},
		{Number: 2, Word: "ที่สอง", Suffix: ".", Masculine: "ที่สอง", Feminine: "ที่สอง", Neuter: "ที่สอง"},
		{Number: 3, Word: "ที่สาม", Suffix: ".", Masculine: "ที่สาม", Feminine: "ที่สาม", Neuter: "ที่สาม"},
		{Number: 4, Word: "ที่สี่", Suffix: ".", Masculine: "ที่สี่", Feminine: "ที่สี่", Neuter: "ที่สี่"},
		{Number: 5, Word: "ที่ห้า", Suffix: ".", Masculine: "ที่ห้า", Feminine: "ที่ห้า", Neuter: "ที่ห้า"},
		{Number: 6, Word: "ที่หก", Suffix: ".", Masculine: "ที่หก", Feminine: "ที่หก", Neuter: "ที่หก"},
		{Number: 7, Word: "ที่เจ็ด", Suffix: ".", Masculine: "ที่เจ็ด", Feminine: "ที่เจ็ด", Neuter: "ที่เจ็ด"},
		{Number: 8, Word: "ที่แปด", Suffix: ".", Masculine: "ที่แปด", Feminine: "ที่แปด", Neuter: "ที่แปด"},
		{Number: 9, Word: "ที่เก้า", Suffix: ".", Masculine: "ที่เก้า", Feminine: "ที่เก้า", Neuter: "ที่เก้า"},
		{Number: 10, Word: "ที่สิบ", Suffix: ".", Masculine: "ที่สิบ", Feminine: "ที่สิบ", Neuter: "ที่สิบ"},
		{Number: 11, Word: "ที่สิบเอ็ด", Suffix: ".", Masculine: "ที่สิบเอ็ด", Feminine: "ที่สิบเอ็ด", Neuter: "ที่สิบเอ็ด"},
		{Number: 12, Word: "ที่สิบสอง", Suffix: ".", Masculine: "ที่สิบสอง", Feminine: "ที่สิบสอง", Neuter: "ที่สิบสอง"},
		{Number: 20, Word: "ที่ยี่สิบ", Suffix: ".", Masculine: "ที่ยี่สิบ", Feminine: "ที่ยี่สิบ", Neuter: "ที่ยี่สิบ"},
		{Number: 21, Word: "ที่ยี่สิบเอ็ด", Suffix: ".", Masculine: "ที่ยี่สิบเอ็ด", Feminine: "ที่ยี่สิบเอ็ด", Neuter: "ที่ยี่สิบเอ็ด"},
		{Number: 30, Word: "ที่สามสิบ", Suffix: ".", Masculine: "ที่สามสิบ", Feminine: "ที่สามสิบ", Neuter: "ที่สามสิบ"},
		{Number: 50, Word: "ที่ห้าสิบ", Suffix: ".", Masculine: "ที่ห้าสิบ", Feminine: "ที่ห้าสิบ", Neuter: "ที่ห้าสิบ"},
		{Number: 100, Word: "ที่หนึ่งร้อย", Suffix: ".", Masculine: "ที่หนึ่งร้อย", Feminine: "ที่หนึ่งร้อย", Neuter: "ที่หนึ่งร้อย"},
		{Number: 1000, Word: "ที่หนึ่งพัน", Suffix: ".", Masculine: "ที่หนึ่งพัน", Feminine: "ที่หนึ่งพัน", Neuter: "ที่หนึ่งพัน"},
	},
	LocaleFormatter: &ThaiFormatter{},
}

// ThaiFormatter handles Thai (Thailand) formatting
type ThaiFormatter struct{}

func (f *ThaiFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *ThaiFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	return result + currencyName
}

func (f *ThaiFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + andText + fractionalWords
}

func (f *ThaiFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	return result + fractionName
}

func (f *ThaiFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + result
}

func (f *ThaiFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 2
	}
	return amount.Truncate(int32(precision))
}

func (f *ThaiFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAsianDecimal(amount)
}

func (f *ThaiFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}
	return FormatAsianCurrency(amount, currencySymbol)
}
