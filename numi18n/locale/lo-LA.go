package locale

import (
	"github.com/shopspring/decimal"
)

// LOLALocale represents the Lao (Laos) locale
var LOLALocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Kip",
		Plural:   "Kip",
		Singular: "Kip",
		Symbol:   "₭",
		FractionUnit: FractionUnit{
			Name:     "Att",
			Plural:   "Att",
			Singular: "Att",
			Symbol:   "att",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Laos",
		Currency:       "LAK",
		ISO3166Alpha2:  "LA",
		ISO3166Alpha3:  "LAO",
		ISO3166Numeric: "418",
		Locale:         "lo-LA",
		Timezone:       []string{"Asia/Vientiane"},
		Language:       "lo",
		Emoji:          "🇱🇦",
	},
	LocaleFormatter: &LaoFormatter{},
	Texts: Texts{
		And:   "ແລະ",
		Minus: "ລົບ",
		Only:  "ພຽງ",
		Point: "ຈຸດ",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "ພັນລ້ານລ້ານ"},
		{Number: 1000000000000, Value: "ລ້ານລ້ານ"},
		{Number: 1000000000, Value: "ພັນລ້ານ"},
		{Number: 1000000, Value: "ລ້ານ"},
		{Number: 100000, Value: "ແສນ"},
		{Number: 10000, Value: "ໝື່ນ"},
		{Number: 9000, Value: "ເກົ້າພັນ"},
		{Number: 8000, Value: "ແປດພັນ"},
		{Number: 7000, Value: "ເຈັດພັນ"},
		{Number: 6000, Value: "ຫົກພັນ"},
		{Number: 5000, Value: "ຫ້າພັນ"},
		{Number: 4000, Value: "ສີ່ພັນ"},
		{Number: 3000, Value: "ສາມພັນ"},
		{Number: 2000, Value: "ສອງພັນ"},
		{Number: 1000, Value: "ພັນ"},
		{Number: 900, Value: "ເກົ້າຮ້ອຍ"},
		{Number: 800, Value: "ແປດຮ້ອຍ"},
		{Number: 700, Value: "ເຈັດຮ້ອຍ"},
		{Number: 600, Value: "ຫົກຮ້ອຍ"},
		{Number: 500, Value: "ຫ້າຮ້ອຍ"},
		{Number: 400, Value: "ສີ່ຮ້ອຍ"},
		{Number: 300, Value: "ສາມຮ້ອຍ"},
		{Number: 200, Value: "ສອງຮ້ອຍ"},
		{Number: 100, Value: "ຮ້ອຍ"},
		{Number: 99, Value: "ເກົ້າສິບເກົ້າ"},
		{Number: 98, Value: "ເກົ້າສິບແປດ"},
		{Number: 97, Value: "ເກົ້າສິບເຈັດ"},
		{Number: 96, Value: "ເກົ້າສິບຫົກ"},
		{Number: 95, Value: "ເກົ້າສິບຫ້າ"},
		{Number: 94, Value: "ເກົ້າສິບສີ່"},
		{Number: 93, Value: "ເກົ້າສິບສາມ"},
		{Number: 92, Value: "ເກົ້າສິບສອງ"},
		{Number: 91, Value: "ເກົ້າສິບເອັດ"},
		{Number: 90, Value: "ເກົ້າສິບ"},
		{Number: 89, Value: "ແປດສິບເກົ້າ"},
		{Number: 88, Value: "ແປດສິບແປດ"},
		{Number: 87, Value: "ແປດສິບເຈັດ"},
		{Number: 86, Value: "ແປດສິບຫົກ"},
		{Number: 85, Value: "ແປດສິບຫ້າ"},
		{Number: 84, Value: "ແປດສິບສີ່"},
		{Number: 83, Value: "ແປດສິບສາມ"},
		{Number: 82, Value: "ແປດສິບສອງ"},
		{Number: 81, Value: "ແປດສິບເອັດ"},
		{Number: 80, Value: "ແປດສິບ"},
		{Number: 79, Value: "ເຈັດສິບເກົ້າ"},
		{Number: 78, Value: "ເຈັດສິບແປດ"},
		{Number: 77, Value: "ເຈັດສິບເຈັດ"},
		{Number: 76, Value: "ເຈັດສິບຫົກ"},
		{Number: 75, Value: "ເຈັດສິບຫ້າ"},
		{Number: 74, Value: "ເຈັດສິບສີ່"},
		{Number: 73, Value: "ເຈັດສິບສາມ"},
		{Number: 72, Value: "ເຈັດສິບສອງ"},
		{Number: 71, Value: "ເຈັດສິບເອັດ"},
		{Number: 70, Value: "ເຈັດສິບ"},
		{Number: 69, Value: "ຫົກສິບເກົ້າ"},
		{Number: 68, Value: "ຫົກສິບແປດ"},
		{Number: 67, Value: "ຫົກສິບເຈັດ"},
		{Number: 66, Value: "ຫົກສິບຫົກ"},
		{Number: 65, Value: "ຫົກສິບຫ້າ"},
		{Number: 64, Value: "ຫົກສິບສີ່"},
		{Number: 63, Value: "ຫົກສິບສາມ"},
		{Number: 62, Value: "ຫົກສິບສອງ"},
		{Number: 61, Value: "ຫົກສິບເອັດ"},
		{Number: 60, Value: "ຫົກສິບ"},
		{Number: 59, Value: "ຫ້າສິບເກົ້າ"},
		{Number: 58, Value: "ຫ້າສິບແປດ"},
		{Number: 57, Value: "ຫ້າສິບເຈັດ"},
		{Number: 56, Value: "ຫ້າສິບຫົກ"},
		{Number: 55, Value: "ຫ້າສິບຫ້າ"},
		{Number: 54, Value: "ຫ້າສິບສີ່"},
		{Number: 53, Value: "ຫ້າສິບສາມ"},
		{Number: 52, Value: "ຫ້າສິບສອງ"},
		{Number: 51, Value: "ຫ້າສິບເອັດ"},
		{Number: 50, Value: "ຫ້າສິບ"},
		{Number: 49, Value: "ສີ່ສິບເກົ້າ"},
		{Number: 48, Value: "ສີ່ສິບແປດ"},
		{Number: 47, Value: "ສີ່ສິບເຈັດ"},
		{Number: 46, Value: "ສີ່ສິບຫົກ"},
		{Number: 45, Value: "ສີ່ສິບຫ້າ"},
		{Number: 44, Value: "ສີ່ສິບສີ່"},
		{Number: 43, Value: "ສີ່ສິບສາມ"},
		{Number: 42, Value: "ສີ່ສິບສອງ"},
		{Number: 41, Value: "ສີ່ສິບເອັດ"},
		{Number: 40, Value: "ສີ່ສິບ"},
		{Number: 39, Value: "ສາມສິບເກົ້າ"},
		{Number: 38, Value: "ສາມສິບແປດ"},
		{Number: 37, Value: "ສາມສິບເຈັດ"},
		{Number: 36, Value: "ສາມສິບຫົກ"},
		{Number: 35, Value: "ສາມສິບຫ້າ"},
		{Number: 34, Value: "ສາມສິບສີ່"},
		{Number: 33, Value: "ສາມສິບສາມ"},
		{Number: 32, Value: "ສາມສິບສອງ"},
		{Number: 31, Value: "ສາມສິບເອັດ"},
		{Number: 30, Value: "ສາມສິບ"},
		{Number: 29, Value: "ຊາວເກົ້າ"},
		{Number: 28, Value: "ຊາວແປດ"},
		{Number: 27, Value: "ຊາວເຈັດ"},
		{Number: 26, Value: "ຊາວຫົກ"},
		{Number: 25, Value: "ຊາວຫ້າ"},
		{Number: 24, Value: "ຊາວສີ່"},
		{Number: 23, Value: "ຊາວສາມ"},
		{Number: 22, Value: "ຊາວສອງ"},
		{Number: 21, Value: "ຊາວເອັດ"},
		{Number: 20, Value: "ຊາວ"},
		{Number: 19, Value: "ສິບເກົ້າ"},
		{Number: 18, Value: "ສິບແປດ"},
		{Number: 17, Value: "ສິບເຈັດ"},
		{Number: 16, Value: "ສິບຫົກ"},
		{Number: 15, Value: "ສິບຫ້າ"},
		{Number: 14, Value: "ສິບສີ່"},
		{Number: 13, Value: "ສິບສາມ"},
		{Number: 12, Value: "ສິບສອງ"},
		{Number: 11, Value: "ສິບເອັດ"},
		{Number: 10, Value: "ສິບ"},
		{Number: 9, Value: "ເກົ້າ"},
		{Number: 8, Value: "ແປດ"},
		{Number: 7, Value: "ເຈັດ"},
		{Number: 6, Value: "ຫົກ"},
		{Number: 5, Value: "ຫ້າ"},
		{Number: 4, Value: "ສີ່"},
		{Number: 3, Value: "ສາມ"},
		{Number: 2, Value: "ສອງ"},
		{Number: 1, Value: "ຫນຶ່ງ"},
		{Number: 0, Value: "ສູນ"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "ນຶ່ງຮ້ອຍ"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "ທີຫນຶ່ງ", Suffix: ".", Masculine: "ທີຫນຶ່ງ", Feminine: "ທີຫນຶ່ງ", Neuter: "ທີຫນຶ່ງ"},
		{Number: 2, Word: "ທີສອງ", Suffix: ".", Masculine: "ທີສອງ", Feminine: "ທີສອງ", Neuter: "ທີສອງ"},
		{Number: 3, Word: "ທີສາມ", Suffix: ".", Masculine: "ທີສາມ", Feminine: "ທີສາມ", Neuter: "ທີສາມ"},
		{Number: 4, Word: "ທີສີ່", Suffix: ".", Masculine: "ທີສີ່", Feminine: "ທີສີ່", Neuter: "ທີສີ່"},
		{Number: 5, Word: "ທີຫ້າ", Suffix: ".", Masculine: "ທີຫ້າ", Feminine: "ທີຫ້າ", Neuter: "ທີຫ້າ"},
		{Number: 6, Word: "ທີຫົກ", Suffix: ".", Masculine: "ທີຫົກ", Feminine: "ທີຫົກ", Neuter: "ທີຫົກ"},
		{Number: 7, Word: "ທີເຈັດ", Suffix: ".", Masculine: "ທີເຈັດ", Feminine: "ທີເຈັດ", Neuter: "ທີເຈັດ"},
		{Number: 8, Word: "ທີແປດ", Suffix: ".", Masculine: "ທີແປດ", Feminine: "ທີແປດ", Neuter: "ທີແປດ"},
		{Number: 9, Word: "ທີເກົ້າ", Suffix: ".", Masculine: "ທີເກົ້າ", Feminine: "ທີເກົ້າ", Neuter: "ທີເກົ້າ"},
		{Number: 10, Word: "ທີສິບ", Suffix: ".", Masculine: "ທີສິບ", Feminine: "ທີສິບ", Neuter: "ທີສິບ"},
		{Number: 11, Word: "ທີສິບເອັດ", Suffix: ".", Masculine: "ທີສິບເອັດ", Feminine: "ທີສິບເອັດ", Neuter: "ທີສິບເອັດ"},
		{Number: 12, Word: "ທີສິບສອງ", Suffix: ".", Masculine: "ທີສິບສອງ", Feminine: "ທີສິບສອງ", Neuter: "ທີສິບສອງ"},
		{Number: 13, Word: "ທີສິບສາມ", Suffix: ".", Masculine: "ທີສິບສາມ", Feminine: "ທີສິບສາມ", Neuter: "ທີສິບສາມ"},
		{Number: 14, Word: "ທີສິບສີ່", Suffix: ".", Masculine: "ທີສິບສີ່", Feminine: "ທີສິບສີ່", Neuter: "ທີສິບສີ່"},
		{Number: 15, Word: "ທີສິບຫ້າ", Suffix: ".", Masculine: "ທີສິບຫ້າ", Feminine: "ທີສິບຫ້າ", Neuter: "ທີສິບຫ້າ"},
		{Number: 16, Word: "ທີສິບຫົກ", Suffix: ".", Masculine: "ທີສິບຫົກ", Feminine: "ທີສິບຫົກ", Neuter: "ທີສິບຫົກ"},
		{Number: 17, Word: "ທີສິບເຈັດ", Suffix: ".", Masculine: "ທີສິບເຈັດ", Feminine: "ທີສິບເຈັດ", Neuter: "ທີສິບເຈັດ"},
		{Number: 18, Word: "ທີສິບແປດ", Suffix: ".", Masculine: "ທີສິບແປດ", Feminine: "ທີສິບແປດ", Neuter: "ທີສິບແປດ"},
		{Number: 19, Word: "ທີສິບເກົ້າ", Suffix: ".", Masculine: "ທີສິບເກົ້າ", Feminine: "ທີສິບເກົ້າ", Neuter: "ທີສິບເກົ້າ"},
		{Number: 20, Word: "ທີຊາວ", Suffix: ".", Masculine: "ທີຊາວ", Feminine: "ທີຊາວ", Neuter: "ທີຊາວ"},
		{Number: 21, Word: "ທີຊາວເອັດ", Suffix: ".", Masculine: "ທີຊາວເອັດ", Feminine: "ທີຊາວເອັດ", Neuter: "ທີຊາວເອັດ"},
		{Number: 30, Word: "ທີສາມສິບ", Suffix: ".", Masculine: "ທີສາມສິບ", Feminine: "ທີສາມສິບ", Neuter: "ທີສາມສິບ"},
		{Number: 40, Word: "ທີສີ່ສິບ", Suffix: ".", Masculine: "ທີສີ່ສິບ", Feminine: "ທີສີ່ສິບ", Neuter: "ທີສີ່ສິບ"},
		{Number: 50, Word: "ທີຫ້າສິບ", Suffix: ".", Masculine: "ທີຫ້າສິບ", Feminine: "ທີຫ້າສິບ", Neuter: "ທີຫ້າສິບ"},
		{Number: 60, Word: "ທີຫົກສິບ", Suffix: ".", Masculine: "ທີຫົກສິບ", Feminine: "ທີຫົກສິບ", Neuter: "ທີຫົກສິບ"},
		{Number: 70, Word: "ທີເຈັດສິບ", Suffix: ".", Masculine: "ທີເຈັດສິບ", Feminine: "ທີເຈັດສິບ", Neuter: "ທີເຈັດສິບ"},
		{Number: 80, Word: "ທີແປດສິບ", Suffix: ".", Masculine: "ທີແປດສິບ", Feminine: "ທີແປດສິບ", Neuter: "ທີແປດສິບ"},
		{Number: 90, Word: "ທີເກົ້າສິບ", Suffix: ".", Masculine: "ທີເກົ້າສິບ", Feminine: "ທີເກົ້າສິບ", Neuter: "ທີເກົ້າສິບ"},
		{Number: 100, Word: "ທີຮ້ອຍ", Suffix: ".", Masculine: "ທີຮ້ອຍ", Feminine: "ທີຮ້ອຍ", Neuter: "ທີຮ້ອຍ"},
		{Number: 1000, Word: "ທีພັນ", Suffix: ".", Masculine: "ທีພັນ", Feminine: "ທีພັນ", Neuter: "ທีພັນ"},
	},
}

// LaoFormatter handles Lao formatting
type LaoFormatter struct{}

func (f *LaoFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *LaoFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	// In Lao, Kip doesn't change form for singular/plural
	return result + " " + currencyName
}

func (f *LaoFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *LaoFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	// In Lao, Att doesn't change form for singular/plural
	return result + " " + fractionName
}

func (f *LaoFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *LaoFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	return amount.Truncate(int32(precision))
}

func (f *LaoFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *LaoFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	return FormatAngloCurrency(amount, currencySymbol)
}
