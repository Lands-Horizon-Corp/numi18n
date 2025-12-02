package locale

import (
	"strings"
	"github.com/shopspring/decimal"
)

// KHLocale is a NumI18NLocale configured for Cambodia (km-KH)
var KHLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "រៀល",
		Plural:   "រៀល",
		Singular: "រៀល",
		Symbol:   "៛",
		FractionUnit: FractionUnit{
			Name:     "សេន",
			Plural:   "សេន",
			Singular: "សេន",
			Symbol:   "សេន",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Cambodia",
		Currency:       "KHR",
		ISO3166Alpha2:  "KH",
		ISO3166Alpha3:  "KHM",
		ISO3166Numeric: "116",
		Locale:         "km-KH",
		Timezone:       []string{"Asia/Phnom_Penh"},
		Language:       "km",
	},
	Texts: Texts{
		And:   "និង",
		Minus: "ដក",
		Only:  "តែប៉ុណ្ណោះ",
		Point: "ចំណុច",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "មួយ\u200bខ្វាដ្រីលាន"},
		{Number: 1000000000000, Value: "មួយ\u200bទ្រីលាន"},
		{Number: 1000000000, Value: "មួយ\u200bពាន់\u200bលាន"},
		{Number: 1000000, Value: "មួយ\u200bលាន"},
		{Number: 1000, Value: "មួយ\u200bពាន់"},
		{Number: 100, Value: "មួយ\u200bរយ"},
		{Number: 90, Value: "កៅ\u200bស៊ប"},
		{Number: 80, Value: "ប៉ែត\u200bស៊ប"},
		{Number: 70, Value: "ចិត\u200bស៊ប"},
		{Number: 60, Value: "ហុក\u200bស៊ប"},
		{Number: 50, Value: "ហា\u200bស៊ប"},
		{Number: 40, Value: "សែ\u200bស៊ប"},
		{Number: 30, Value: "សាម\u200bស៊ប"},
		{Number: 20, Value: "ម្ភៃ"},
		{Number: 19, Value: "ដប់\u200bកៅ"},
		{Number: 18, Value: "ដប់\u200bប៉ែត"},
		{Number: 17, Value: "ដប់\u200bព្រាំ"},
		{Number: 16, Value: "ដប់\u200bប្រាំ\u200bមួយ"},
		{Number: 15, Value: "ដប់\u200bប្រាំ"},
		{Number: 14, Value: "ដប់\u200bបួន"},
		{Number: 13, Value: "ដប់\u200bបី"},
		{Number: 12, Value: "ដប់\u200bពីរ"},
		{Number: 11, Value: "ដប់\u200bមួយ"},
		{Number: 10, Value: "ដប់"},
		{Number: 9, Value: "កៅ"},
		{Number: 8, Value: "ប៉ែត"},
		{Number: 7, Value: "ព្រាំ"},
		{Number: 6, Value: "ប្រាំ\u200bមួយ"},
		{Number: 5, Value: "ប្រាំ"},
		{Number: 4, Value: "បួន"},
		{Number: 3, Value: "បី"},
		{Number: 2, Value: "ពីរ"},
		{Number: 1, Value: "មួយ"},
		{Number: 0, Value: "សូន្យ"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "មួយ\u200bរយ"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "ទីមួយ", Suffix: ".", Masculine: "ទីមួយ", Feminine: "ទីមួយ", Neuter: "ទីមួយ"},
		{Number: 2, Word: "ទីពីរ", Suffix: ".", Masculine: "ទីពីរ", Feminine: "ទីពីរ", Neuter: "ទីពីរ"},
		{Number: 3, Word: "ទីបី", Suffix: ".", Masculine: "ទីបី", Feminine: "ទីបី", Neuter: "ទីបី"},
		{Number: 4, Word: "ទីបួន", Suffix: ".", Masculine: "ទីបួន", Feminine: "ទីបួន", Neuter: "ទីបួន"},
		{Number: 5, Word: "ទីប្រាំ", Suffix: ".", Masculine: "ទីប្រាំ", Feminine: "ទីប្រាំ", Neuter: "ទីប្រាំ"},
		{Number: 6, Word: "ទីប្រាំមួយ", Suffix: ".", Masculine: "ទីប្រាំមួយ", Feminine: "ទីប្រាំមួយ", Neuter: "ទីប្រាំមួយ"},
		{Number: 7, Word: "ទីព្រាំ", Suffix: ".", Masculine: "ទីព្រាំ", Feminine: "ទីព្រាំ", Neuter: "ទីព្រាំ"},
		{Number: 8, Word: "ទីប៉ែត", Suffix: ".", Masculine: "ទីប៉ែត", Feminine: "ទីប៉ែត", Neuter: "ទីប៉ែត"},
		{Number: 9, Word: "ទីកៅ", Suffix: ".", Masculine: "ទីកៅ", Feminine: "ទីកៅ", Neuter: "ទីកៅ"},
		{Number: 10, Word: "ទីដប់", Suffix: ".", Masculine: "ទីដប់", Feminine: "ទីដប់", Neuter: "ទីដប់"},
		{Number: 11, Word: "ទីដប់មួយ", Suffix: ".", Masculine: "ទីដប់មួយ", Feminine: "ទីដប់មួយ", Neuter: "ទីដប់មួយ"},
		{Number: 12, Word: "ទីដប់ពីរ", Suffix: ".", Masculine: "ទីដប់ពីរ", Feminine: "ទីដប់ពីរ", Neuter: "ទីដប់ពីរ"},
		{Number: 13, Word: "ទីដប់បី", Suffix: ".", Masculine: "ទីដប់បី", Feminine: "ទីដប់បី", Neuter: "ទីដប់បី"},
		{Number: 14, Word: "ទីដប់បួន", Suffix: ".", Masculine: "ទីដប់បួន", Feminine: "ទីដប់បួន", Neuter: "ទីដប់បួន"},
		{Number: 15, Word: "ទីដប់ប្រាំ", Suffix: ".", Masculine: "ទីដប់ប្រាំ", Feminine: "ទីដប់ប្រាំ", Neuter: "ទីដប់ប្រាំ"},
		{Number: 16, Word: "ទីដប់ប្រាំមួយ", Suffix: ".", Masculine: "ទីដប់ប្រាំមួយ", Feminine: "ទីដប់ប្រាំមួយ", Neuter: "ទីដប់ប្រាំមួយ"},
		{Number: 17, Word: "ទីដប់ព្រាំ", Suffix: ".", Masculine: "ទីដប់ព្រាំ", Feminine: "ទីដប់ព្រាំ", Neuter: "ទីដប់ព្រាំ"},
		{Number: 18, Word: "ទីដប់ប៉ែត", Suffix: ".", Masculine: "ទីដប់ប៉ែត", Feminine: "ទីដប់ប៉ែត", Neuter: "ទីដប់ប៉ែត"},
		{Number: 19, Word: "ទីដប់កៅ", Suffix: ".", Masculine: "ទីដប់កៅ", Feminine: "ទីដប់កៅ", Neuter: "ទីដប់កៅ"},
		{Number: 20, Word: "ទីម្ភៃ", Suffix: ".", Masculine: "ទីម្ភៃ", Feminine: "ទីម្ភៃ", Neuter: "ទីម្ភៃ"},
		{Number: 21, Word: "ទីម្ភៃមួយ", Suffix: ".", Masculine: "ទីម្ភៃមួយ", Feminine: "ទីម្ភៃមួយ", Neuter: "ទីម្ភៃមួយ"},
		{Number: 30, Word: "ទីសាមស៊ប", Suffix: ".", Masculine: "ទីសាមស៊ប", Feminine: "ទីសាមស៊ប", Neuter: "ទីសាមស៊ប"},
		{Number: 40, Word: "ទីសែស៊ប", Suffix: ".", Masculine: "ទីសែស៊ប", Feminine: "ទីសែស៊ប", Neuter: "ទីសែស៊ប"},
		{Number: 50, Word: "ទីហាស៊ប", Suffix: ".", Masculine: "ទីហាស៊ប", Feminine: "ទីហាស៊ប", Neuter: "ទីហាស៊ប"},
		{Number: 60, Word: "ទីហុកស៊ប", Suffix: ".", Masculine: "ទីហុកស៊ប", Feminine: "ទីហុកស៊ប", Neuter: "ទីហុកស៊ប"},
		{Number: 70, Word: "ទីចិតស៊ប", Suffix: ".", Masculine: "ទីចិតស៊ប", Feminine: "ទីចិតស៊ប", Neuter: "ទីចិតស៊ប"},
		{Number: 80, Word: "ទីប៉ែតស៊ប", Suffix: ".", Masculine: "ទីប៉ែតស៊ប", Feminine: "ទីប៉ែតស៊ប", Neuter: "ទីប៉ែតស៊ប"},
		{Number: 90, Word: "ទីកៅស៊ប", Suffix: ".", Masculine: "ទីកៅស៊ប", Feminine: "ទីកៅស៊ប", Neuter: "ទីកៅស៊ប"},
		{Number: 100, Word: "ទីមួយរយ", Suffix: ".", Masculine: "ទីមួយរយ", Feminine: "ទីមួយរយ", Neuter: "ទីមួយរយ"},
		{Number: 1000, Word: "ទីមួយពាន់", Suffix: ".", Masculine: "ទីមួយពាន់", Feminine: "ទីមួយពាន់", Neuter: "ទីមួយពាន់"},
	},
	LocaleFormatter: &KhmerFormatter{},
}

// KhmerFormatter handles Khmer formatting
type KhmerFormatter struct{}

func (f *KhmerFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *KhmerFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *KhmerFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *KhmerFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *KhmerFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *KhmerFormatter) ChopDecimal(value decimal.Decimal, places int) decimal.Decimal {
	multiplier := decimal.NewFromInt(1)
	for range places {
		multiplier = multiplier.Mul(decimal.NewFromInt(10))
	}
	return value.Mul(multiplier).Truncate(0).Div(multiplier)
}


func (f *KhmerFormatter) FormatDecimalNumber(amount float64) string {
	return DefaultFormatDecimalNumber(amount, ",", ".")
}
func (f *KhmerFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	formattedNumber := f.FormatDecimalNumber(amount)
	
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}
	
	// Default currency placement for this locale (prefix with symbol)
	if strings.HasPrefix(formattedNumber, "-") {
		formattedNumber = strings.TrimPrefix(formattedNumber, "-")
		return "-" + currencySymbol + formattedNumber
	}
	
	return currencySymbol + formattedNumber
}
