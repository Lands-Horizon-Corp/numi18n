package locale

import (
	"github.com/shopspring/decimal"
)

// URINLocale represents the Urdu (India) locale
var URINLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Indian Rupee",
		Plural:   "روپے",
		Singular: "روپیہ",
		Symbol:   "₹",
		FractionUnit: FractionUnit{
			Name:     "Paisa",
			Plural:   "پیسے",
			Singular: "پیسہ",
			Symbol:   "p",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "India",
		Currency:       "INR",
		ISO3166Alpha2:  "IN",
		ISO3166Alpha3:  "IND",
		ISO3166Numeric: "356",
		Locale:         "ur-IN",
		Timezone:       []string{"Asia/Kolkata"},
		Language:       "ur",
		Emoji:          "🇮🇳",
		PhoneCode:      "+91",
		Domain:         ".in",
	},
	Texts: Texts{
		And:   "اور",
		Minus: "منفی",
		Only:  "صرف",
		Point: "نقطہ",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "ایک کھرب"},
		{Number: 1000000000000, Value: "ایک ارب"},
		{Number: 1000000000, Value: "ایک کروڑ"},
		{Number: 10000000, Value: "ایک کروڑ"},
		{Number: 1000000, Value: "دس لاکھ"},
		{Number: 100000, Value: "ایک لاکھ"},
		{Number: 90000, Value: "نوے ہزار"},
		{Number: 80000, Value: "اسی ہزار"},
		{Number: 70000, Value: "ستر ہزار"},
		{Number: 60000, Value: "ساٹھ ہزار"},
		{Number: 50000, Value: "پچاس ہزار"},
		{Number: 40000, Value: "چالیس ہزار"},
		{Number: 30000, Value: "تیس ہزار"},
		{Number: 20000, Value: "بیس ہزار"},
		{Number: 19000, Value: "انیس ہزار"},
		{Number: 18000, Value: "اٹھارہ ہزار"},
		{Number: 17000, Value: "سترہ ہزار"},
		{Number: 16000, Value: "سولہ ہزار"},
		{Number: 15000, Value: "پندرہ ہزار"},
		{Number: 14000, Value: "چودہ ہزار"},
		{Number: 13000, Value: "تیرہ ہزار"},
		{Number: 12000, Value: "بارہ ہزار"},
		{Number: 11000, Value: "گیارہ ہزار"},
		{Number: 10000, Value: "دس ہزار"},
		{Number: 9000, Value: "نو ہزار"},
		{Number: 8000, Value: "آٹھ ہزار"},
		{Number: 7000, Value: "سات ہزار"},
		{Number: 6000, Value: "چھ ہزار"},
		{Number: 5000, Value: "پانچ ہزار"},
		{Number: 4000, Value: "چار ہزار"},
		{Number: 3000, Value: "تین ہزار"},
		{Number: 2000, Value: "دو ہزار"},
		{Number: 1000, Value: "ایک ہزار"},
		{Number: 900, Value: "نو سو"},
		{Number: 800, Value: "آٹھ سو"},
		{Number: 700, Value: "سات سو"},
		{Number: 600, Value: "چھ سو"},
		{Number: 500, Value: "پانچ سو"},
		{Number: 400, Value: "چار سو"},
		{Number: 300, Value: "تین سو"},
		{Number: 200, Value: "دو سو"},
		{Number: 100, Value: "ایک سو"},
		{Number: 99, Value: "ننانوے"},
		{Number: 98, Value: "اٹھانوے"},
		{Number: 97, Value: "ستانوے"},
		{Number: 96, Value: "چھیانوے"},
		{Number: 95, Value: "پچانوے"},
		{Number: 94, Value: "چورانوے"},
		{Number: 93, Value: "ترانوے"},
		{Number: 92, Value: "بانوے"},
		{Number: 91, Value: "اکیانوے"},
		{Number: 90, Value: "نوے"},
		{Number: 89, Value: "نواسی"},
		{Number: 88, Value: "اٹھاسی"},
		{Number: 87, Value: "ستاسی"},
		{Number: 86, Value: "چھیاسی"},
		{Number: 85, Value: "پچاسی"},
		{Number: 84, Value: "چوراسی"},
		{Number: 83, Value: "تراسی"},
		{Number: 82, Value: "باسی"},
		{Number: 81, Value: "اکیاسی"},
		{Number: 80, Value: "اسی"},
		{Number: 79, Value: "انہتر"},
		{Number: 78, Value: "اٹھہتر"},
		{Number: 77, Value: "ستہتر"},
		{Number: 76, Value: "چھہتر"},
		{Number: 75, Value: "پچہتر"},
		{Number: 74, Value: "چوہتر"},
		{Number: 73, Value: "تہتر"},
		{Number: 72, Value: "بہتر"},
		{Number: 71, Value: "اکہتر"},
		{Number: 70, Value: "ستر"},
		{Number: 69, Value: "انہتر"},
		{Number: 68, Value: "اڑسٹھ"},
		{Number: 67, Value: "سڑسٹھ"},
		{Number: 66, Value: "چھیاسٹھ"},
		{Number: 65, Value: "پینسٹھ"},
		{Number: 64, Value: "چوسٹھ"},
		{Number: 63, Value: "ترسٹھ"},
		{Number: 62, Value: "باسٹھ"},
		{Number: 61, Value: "اکسٹھ"},
		{Number: 60, Value: "ساٹھ"},
		{Number: 59, Value: "انسٹھ"},
		{Number: 58, Value: "اٹھاون"},
		{Number: 57, Value: "ستاون"},
		{Number: 56, Value: "چھپن"},
		{Number: 55, Value: "پچپن"},
		{Number: 54, Value: "چون"},
		{Number: 53, Value: "ترپن"},
		{Number: 52, Value: "باون"},
		{Number: 51, Value: "اکاون"},
		{Number: 50, Value: "پچاس"},
		{Number: 49, Value: "انچاس"},
		{Number: 48, Value: "اڑتالیس"},
		{Number: 47, Value: "سینتالیس"},
		{Number: 46, Value: "چھیالیس"},
		{Number: 45, Value: "پینتالیس"},
		{Number: 44, Value: "چوالیس"},
		{Number: 43, Value: "تینتالیس"},
		{Number: 42, Value: "بیالیس"},
		{Number: 41, Value: "اکتالیس"},
		{Number: 40, Value: "چالیس"},
		{Number: 39, Value: "انتالیس"},
		{Number: 38, Value: "اڑتیس"},
		{Number: 37, Value: "سینتیس"},
		{Number: 36, Value: "چھتیس"},
		{Number: 35, Value: "پینتیس"},
		{Number: 34, Value: "چونتیس"},
		{Number: 33, Value: "تینتیس"},
		{Number: 32, Value: "بتیس"},
		{Number: 31, Value: "اکتیس"},
		{Number: 30, Value: "تیس"},
		{Number: 29, Value: "انتیس"},
		{Number: 28, Value: "اٹھائیس"},
		{Number: 27, Value: "ستائیس"},
		{Number: 26, Value: "چھبیس"},
		{Number: 25, Value: "پچیس"},
		{Number: 24, Value: "چوبیس"},
		{Number: 23, Value: "تیئیس"},
		{Number: 22, Value: "بائیس"},
		{Number: 21, Value: "اکیس"},
		{Number: 20, Value: "بیس"},
		{Number: 19, Value: "انیس"},
		{Number: 18, Value: "اٹھارہ"},
		{Number: 17, Value: "سترہ"},
		{Number: 16, Value: "سولہ"},
		{Number: 15, Value: "پندرہ"},
		{Number: 14, Value: "چودہ"},
		{Number: 13, Value: "تیرہ"},
		{Number: 12, Value: "بارہ"},
		{Number: 11, Value: "گیارہ"},
		{Number: 10, Value: "دس"},
		{Number: 9, Value: "نو"},
		{Number: 8, Value: "آٹھ"},
		{Number: 7, Value: "سات"},
		{Number: 6, Value: "چھ"},
		{Number: 5, Value: "پانچ"},
		{Number: 4, Value: "چار"},
		{Number: 3, Value: "تین"},
		{Number: 2, Value: "دو"},
		{Number: 1, Value: "ایک"},
		{Number: 0, Value: "صفر"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "سو"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "پہلا", Suffix: "ویں", Masculine: "پہلا", Feminine: "پہلی", Neuter: ""},
		{Number: 2, Word: "دوسرا", Suffix: "ویں", Masculine: "دوسرا", Feminine: "دوسری", Neuter: ""},
		{Number: 3, Word: "تیسرا", Suffix: "ویں", Masculine: "تیسرا", Feminine: "تیسری", Neuter: ""},
		{Number: 4, Word: "چوتھا", Suffix: "ویں", Masculine: "چوتھا", Feminine: "چوتھی", Neuter: ""},
		{Number: 5, Word: "پانچواں", Suffix: "ویں", Masculine: "پانچواں", Feminine: "پانچویں", Neuter: ""},
		{Number: 6, Word: "چھٹا", Suffix: "ویں", Masculine: "چھٹا", Feminine: "چھٹی", Neuter: ""},
		{Number: 7, Word: "ساتواں", Suffix: "ویں", Masculine: "ساتواں", Feminine: "ساتویں", Neuter: ""},
		{Number: 8, Word: "آٹھواں", Suffix: "ویں", Masculine: "آٹھواں", Feminine: "آٹھویں", Neuter: ""},
		{Number: 9, Word: "نواں", Suffix: "ویں", Masculine: "نواں", Feminine: "نویں", Neuter: ""},
		{Number: 10, Word: "دسواں", Suffix: "ویں", Masculine: "دسواں", Feminine: "دسویں", Neuter: ""},
		{Number: 11, Word: "گیارہواں", Suffix: "ویں", Masculine: "گیارہواں", Feminine: "گیارہویں", Neuter: ""},
		{Number: 12, Word: "بارہواں", Suffix: "ویں", Masculine: "بارہواں", Feminine: "بارہویں", Neuter: ""},
		{Number: 13, Word: "تیرہواں", Suffix: "ویں", Masculine: "تیرہواں", Feminine: "تیرہویں", Neuter: ""},
		{Number: 14, Word: "چودہواں", Suffix: "ویں", Masculine: "چودہواں", Feminine: "چودہویں", Neuter: ""},
		{Number: 15, Word: "پندرہواں", Suffix: "ویں", Masculine: "پندرہواں", Feminine: "پندرہویں", Neuter: ""},
		{Number: 16, Word: "سولہواں", Suffix: "ویں", Masculine: "سولہواں", Feminine: "سولہویں", Neuter: ""},
		{Number: 17, Word: "سترہواں", Suffix: "ویں", Masculine: "سترہواں", Feminine: "سترہویں", Neuter: ""},
		{Number: 18, Word: "اٹھارہواں", Suffix: "ویں", Masculine: "اٹھارہواں", Feminine: "اٹھارہویں", Neuter: ""},
		{Number: 19, Word: "انیسواں", Suffix: "ویں", Masculine: "انیسواں", Feminine: "انیسویں", Neuter: ""},
		{Number: 20, Word: "بیسواں", Suffix: "ویں", Masculine: "بیسواں", Feminine: "بیسویں", Neuter: ""},
		{Number: 21, Word: "اکیسواں", Suffix: "ویں", Masculine: "اکیسواں", Feminine: "اکیسویں", Neuter: ""},
		{Number: 30, Word: "تیسواں", Suffix: "ویں", Masculine: "تیسواں", Feminine: "تیسویں", Neuter: ""},
		{Number: 40, Word: "چالیسواں", Suffix: "ویں", Masculine: "چالیسواں", Feminine: "چالیسویں", Neuter: ""},
		{Number: 50, Word: "پچاسواں", Suffix: "ویں", Masculine: "پچاسواں", Feminine: "پچاسویں", Neuter: ""},
		{Number: 60, Word: "ساٹھواں", Suffix: "ویں", Masculine: "ساٹھواں", Feminine: "ساٹھویں", Neuter: ""},
		{Number: 70, Word: "ستّرواں", Suffix: "ویں", Masculine: "ستّرواں", Feminine: "ستّرویں", Neuter: ""},
		{Number: 80, Word: "اسّیواں", Suffix: "ویں", Masculine: "اسّیواں", Feminine: "اسّیویں", Neuter: ""},
		{Number: 90, Word: "نوّیواں", Suffix: "ویں", Masculine: "نوّیواں", Feminine: "نوّیویں", Neuter: ""},
		{Number: 100, Word: "سواں", Suffix: "ویں", Masculine: "سواں", Feminine: "سویں", Neuter: ""},
		{Number: 1000, Word: "ہزارواں", Suffix: "ویں", Masculine: "ہزارواں", Feminine: "ہزارویں", Neuter: ""},
		{Number: 100000, Word: "لاکھواں", Suffix: "ویں", Masculine: "لاکھواں", Feminine: "لاکھویں", Neuter: ""},
		{Number: 10000000, Word: "کروڑواں", Suffix: "ویں", Masculine: "کروڑواں", Feminine: "کروڑویں", Neuter: ""},
	},
	LocaleFormatter: &UrduFormatter{},
}

// UrduFormatter handles Urdu (India) formatting
type UrduFormatter struct{}

func (f *UrduFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *UrduFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *UrduFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *UrduFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *UrduFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *UrduFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	return amount.Truncate(int32(precision))
}

func (f *UrduFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *UrduFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
