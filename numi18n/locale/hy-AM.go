package locale

import (
	"github.com/shopspring/decimal"
)

// AMLocale is a NumI18NLocale configured for Armenia (hy-AM)
var AMLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Դրամ",
		Plural:   "Դրամ",
		Singular: "Դրամ",
		Symbol:   "֏",
		FractionUnit: FractionUnit{
			Name:     "Լումա",
			Plural:   "Լումա",
			Singular: "Լումա",
			Symbol:   "լ",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Armenia",
		Currency:       "AMD",
		ISO3166Alpha2:  "AM",
		ISO3166Alpha3:  "ARM",
		ISO3166Numeric: "051",
		Locale:         "hy-AM",
		Timezone:       []string{"Asia/Yerevan"},
		Language:       "hy",
		Emoji:          "🇦🇲",
	},
	Texts: Texts{
		And:   "Եվ",
		Minus: "Մինուս",
		Only:  "Միայն",
		Point: "Կետ",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Քվադրիլիոն"},
		{Number: 1000000000000, Value: "Տրիլիոն"},
		{Number: 1000000000, Value: "Միլիարդ"},
		{Number: 1000000, Value: "Միլիոն"},
		{Number: 1000, Value: "Հազար"},
		{Number: 100, Value: "Հարյուր"},
		{Number: 90, Value: "Իննսուն"},
		{Number: 80, Value: "Ութսուն"},
		{Number: 70, Value: "Յոթանասուն"},
		{Number: 60, Value: "Վաթսուն"},
		{Number: 50, Value: "Հիսուն"},
		{Number: 40, Value: "Քառասուն"},
		{Number: 30, Value: "Երեսուն"},
		{Number: 20, Value: "Քսան"},
		{Number: 19, Value: "Տասնինը"},
		{Number: 18, Value: "Տասնութ"},
		{Number: 17, Value: "Տասնյոթ"},
		{Number: 16, Value: "Տասնվեց"},
		{Number: 15, Value: "Տասնհինգ"},
		{Number: 14, Value: "Տասնչորս"},
		{Number: 13, Value: "Տասներեք"},
		{Number: 12, Value: "Տասներկու"},
		{Number: 11, Value: "Տասնմեկ"},
		{Number: 10, Value: "Տասը"},
		{Number: 9, Value: "Ինը"},
		{Number: 8, Value: "Ութ"},
		{Number: 7, Value: "Յոթ"},
		{Number: 6, Value: "Վեց"},
		{Number: 5, Value: "Հինգ"},
		{Number: 4, Value: "Չորս"},
		{Number: 3, Value: "Երեք"},
		{Number: 2, Value: "Երկու"},
		{Number: 1, Value: "Մեկ"},
		{Number: 0, Value: "Զրո"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Մեկ Հարյուր"},
		{Number: 1000, Value: "Մեկ Հազար"},
		{Number: 1000000, Value: "Մեկ Միլիոն"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "առաջին", Suffix: "-րդ", Masculine: "առաջին", Feminine: "առաջին", Neuter: "առաջին"},
		{Number: 2, Word: "երկրորդ", Suffix: "-րդ", Masculine: "երկրորդ", Feminine: "երկրորդ", Neuter: "երկրորդ"},
		{Number: 3, Word: "երրորդ", Suffix: "-րդ", Masculine: "երրորդ", Feminine: "երրորդ", Neuter: "երրորդ"},
		{Number: 4, Word: "չորրորդ", Suffix: "-րդ", Masculine: "չորրորդ", Feminine: "չորրորդ", Neuter: "չորրորդ"},
		{Number: 5, Word: "հինգերորդ", Suffix: "-րդ", Masculine: "հինգերորդ", Feminine: "հինգերորդ", Neuter: "հինգերորդ"},
		{Number: 6, Word: "վեցերորդ", Suffix: "-րդ", Masculine: "վեցերորդ", Feminine: "վեցերորդ", Neuter: "վեցերորդ"},
		{Number: 7, Word: "յոթերորդ", Suffix: "-րդ", Masculine: "յոթերորդ", Feminine: "յոթերորդ", Neuter: "յոթերորդ"},
		{Number: 8, Word: "ութերորդ", Suffix: "-րդ", Masculine: "ութերորդ", Feminine: "ութերորդ", Neuter: "ութերորդ"},
		{Number: 9, Word: "իններորդ", Suffix: "-րդ", Masculine: "իններորդ", Feminine: "իններորդ", Neuter: "իններորդ"},
		{Number: 10, Word: "տասներորդ", Suffix: "-րդ", Masculine: "տասներորդ", Feminine: "տասներորդ", Neuter: "տասներորդ"},
		{Number: 11, Word: "տասնմեկերորդ", Suffix: "-րդ", Masculine: "տասնմեկերորդ", Feminine: "տասնմեկերորդ", Neuter: "տասնմեկերորդ"},
		{Number: 12, Word: "տասներկրորդ", Suffix: "-րդ", Masculine: "տասներկրորդ", Feminine: "տասներկրորդ", Neuter: "տասներկրորդ"},
		{Number: 13, Word: "տասներեքերորդ", Suffix: "-րդ", Masculine: "տասներեքերորդ", Feminine: "տասներեքերորդ", Neuter: "տասներեքերորդ"},
		{Number: 14, Word: "տասնչորսերորդ", Suffix: "-րդ", Masculine: "տասնչորսերորդ", Feminine: "տասնչորսերորդ", Neuter: "տասնչորսերորդ"},
		{Number: 15, Word: "տասնհինգերորդ", Suffix: "-րդ", Masculine: "տասնհինգերորդ", Feminine: "տասնհինգերորդ", Neuter: "տասնհինգերորդ"},
		{Number: 16, Word: "տասնվեցերորդ", Suffix: "-րդ", Masculine: "տասնվեցերորդ", Feminine: "տասնվեցերորդ", Neuter: "տասնվեցերորդ"},
		{Number: 17, Word: "տասնյոթերորդ", Suffix: "-րդ", Masculine: "տասնյոթերորդ", Feminine: "տասնյոթերորդ", Neuter: "տասնյոթերորդ"},
		{Number: 18, Word: "տասնութերորդ", Suffix: "-րդ", Masculine: "տասնութերորդ", Feminine: "տասնութերորդ", Neuter: "տասնութերորդ"},
		{Number: 19, Word: "տասնիններորդ", Suffix: "-րդ", Masculine: "տասնիններորդ", Feminine: "տասնիններորդ", Neuter: "տասնիններորդ"},
		{Number: 20, Word: "քսաներորդ", Suffix: "-րդ", Masculine: "քսաներորդ", Feminine: "քսաներորդ", Neuter: "քսաներորդ"},
		{Number: 21, Word: "քսանմեկերորդ", Suffix: "-րդ", Masculine: "քսանմեկերորդ", Feminine: "քսանմեկերորդ", Neuter: "քսանմեկերորդ"},
		{Number: 30, Word: "երեսուներորդ", Suffix: "-րդ", Masculine: "երեսուներորդ", Feminine: "երեսուներորդ", Neuter: "երեսուներորդ"},
		{Number: 40, Word: "քառասուներորդ", Suffix: "-րդ", Masculine: "քառասուներորդ", Feminine: "քառասուներորդ", Neuter: "քառասուներորդ"},
		{Number: 50, Word: "հիսուներորդ", Suffix: "-րդ", Masculine: "հիսուներորդ", Feminine: "հիսուներորդ", Neuter: "հիսուներորդ"},
		{Number: 60, Word: "վաթսուներորդ", Suffix: "-րդ", Masculine: "վաթսուներորդ", Feminine: "վաթսուներորդ", Neuter: "վաթսուներորդ"},
		{Number: 70, Word: "յոթանասուներորդ", Suffix: "-րդ", Masculine: "յոթանասուներորդ", Feminine: "յոթանասուներորդ", Neuter: "յոթանասուներորդ"},
		{Number: 80, Word: "ութսուներորդ", Suffix: "-րդ", Masculine: "ութսուներորդ", Feminine: "ութսուներորդ", Neuter: "ութսուներորդ"},
		{Number: 90, Word: "իննսուներորդ", Suffix: "-րդ", Masculine: "իննսուներորդ", Feminine: "իննսուներորդ", Neuter: "իննսուներորդ"},
		{Number: 100, Word: "հարյուրերորդ", Suffix: "-րդ", Masculine: "հարյուրերորդ", Feminine: "հարյուրերորդ", Neuter: "հարյուրերորդ"},
		{Number: 1000, Word: "հազարերորդ", Suffix: "-րդ", Masculine: "հազարերորդ", Feminine: "հազարերորդ", Neuter: "հազարերորդ"},
	},
	LocaleFormatter: &ArmenianFormatter{},
}

// ArmenianFormatter handles Armenian (hy-AM) formatting
type ArmenianFormatter struct{}

func (f *ArmenianFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *ArmenianFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *ArmenianFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *ArmenianFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *ArmenianFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *ArmenianFormatter) ChopDecimal(number decimal.Decimal, places int) decimal.Decimal {
	return number.Truncate(int32(places))
}

func (f *ArmenianFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *ArmenianFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
