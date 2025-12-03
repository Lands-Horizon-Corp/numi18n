package locale

import (
	"github.com/shopspring/decimal"
)

// LBLULocale represents the Luxembourgish (Luxembourg) locale
var LBLULocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Euro",
		Plural:   "Euroen",
		Singular: "Euro",
		Symbol:   "€",
		FractionUnit: FractionUnit{
			Name:     "Cent",
			Plural:   "Centen",
			Singular: "Cent",
			Symbol:   "c",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Luxembourg",
		Currency:       "EUR",
		ISO3166Alpha2:  "LU",
		ISO3166Alpha3:  "LUX",
		ISO3166Numeric: "442",
		Locale:         "lb-LU",
		Timezone:       []string{"Europe/Luxembourg"},
		Language:       "lb",
		Emoji:          "🇱🇺",
	},
	Texts: Texts{
		And:   "an",
		Minus: "minus",
		Only:  "nëmmen",
		Point: "Komma",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Billiard"},
		{Number: 1000000000000, Value: "Billion"},
		{Number: 1000000000, Value: "Milliard"},
		{Number: 1000000, Value: "Millioun"},
		{Number: 100000, Value: "honnertdausend"},
		{Number: 10000, Value: "zéngdausend"},
		{Number: 9000, Value: "néngdausend"},
		{Number: 8000, Value: "aachtdausend"},
		{Number: 7000, Value: "siwwendausend"},
		{Number: 6000, Value: "sechsdausend"},
		{Number: 5000, Value: "fënneftausend"},
		{Number: 4000, Value: "véierdausend"},
		{Number: 3000, Value: "dräidausend"},
		{Number: 2000, Value: "zwoudausend"},
		{Number: 1000, Value: "dausend"},
		{Number: 900, Value: "nénghonnert"},
		{Number: 800, Value: "aachthonnert"},
		{Number: 700, Value: "siwwenhonnert"},
		{Number: 600, Value: "sechshonnert"},
		{Number: 500, Value: "fënnefhonnert"},
		{Number: 400, Value: "véierhonnert"},
		{Number: 300, Value: "dräihonnert"},
		{Number: 200, Value: "zweehonnert"},
		{Number: 100, Value: "honnert"},
		{Number: 99, Value: "nénganóngzeg"},
		{Number: 98, Value: "aachtanóngzeg"},
		{Number: 97, Value: "siwwenanóngzeg"},
		{Number: 96, Value: "sechsanóngzeg"},
		{Number: 95, Value: "fënnefanóngzeg"},
		{Number: 94, Value: "véieranóngzeg"},
		{Number: 93, Value: "dräianóngzeg"},
		{Number: 92, Value: "zwouanóngzeg"},
		{Number: 91, Value: "eenanóngzeg"},
		{Number: 90, Value: "nongzeg"},
		{Number: 89, Value: "nénganáachtzeg"},
		{Number: 88, Value: "aachtanáachtzeg"},
		{Number: 87, Value: "siwwenanáachtzeg"},
		{Number: 86, Value: "sechsanáachtzeg"},
		{Number: 85, Value: "fënnefanáachtzeg"},
		{Number: 84, Value: "véieranáachtzeg"},
		{Number: 83, Value: "dräianáachtzeg"},
		{Number: 82, Value: "zwouanáachtzeg"},
		{Number: 81, Value: "eenanáachtzeg"},
		{Number: 80, Value: "aachtzeg"},
		{Number: 79, Value: "nénganséwenzeg"},
		{Number: 78, Value: "aachtanséwenzeg"},
		{Number: 77, Value: "siwwenanséwenzeg"},
		{Number: 76, Value: "sechsanséwenzeg"},
		{Number: 75, Value: "fënnefanséwenzeg"},
		{Number: 74, Value: "véieranséwenzeg"},
		{Number: 73, Value: "dräianséwenzeg"},
		{Number: 72, Value: "zwouanséwenzeg"},
		{Number: 71, Value: "eenanséwenzeg"},
		{Number: 70, Value: "siwwenzeg"},
		{Number: 69, Value: "nénganséchzeg"},
		{Number: 68, Value: "aachtanséchzeg"},
		{Number: 67, Value: "siwwenanséchzeg"},
		{Number: 66, Value: "sechsanséchzeg"},
		{Number: 65, Value: "fënnefanséchzeg"},
		{Number: 64, Value: "véieranséchzeg"},
		{Number: 63, Value: "dräianséchzeg"},
		{Number: 62, Value: "zwouanséchzeg"},
		{Number: 61, Value: "eenanséchzeg"},
		{Number: 60, Value: "sechzeg"},
		{Number: 59, Value: "nénganfofzeg"},
		{Number: 58, Value: "aachtanfofzeg"},
		{Number: 57, Value: "siwwenanfofzeg"},
		{Number: 56, Value: "sechsanfofzeg"},
		{Number: 55, Value: "fënnefanfofzeg"},
		{Number: 54, Value: "véieranfofzeg"},
		{Number: 53, Value: "dräianfofzeg"},
		{Number: 52, Value: "zwouanfofzeg"},
		{Number: 51, Value: "eenanfofzeg"},
		{Number: 50, Value: "fofzeg"},
		{Number: 49, Value: "nénganvéierzeg"},
		{Number: 48, Value: "aachtanvéierzeg"},
		{Number: 47, Value: "siwwenanvéierzeg"},
		{Number: 46, Value: "sechsanvéierzeg"},
		{Number: 45, Value: "fënnefanvéierzeg"},
		{Number: 44, Value: "véieranvéierzeg"},
		{Number: 43, Value: "dräianvéierzeg"},
		{Number: 42, Value: "zwouanvéierzeg"},
		{Number: 41, Value: "eenanvéierzeg"},
		{Number: 40, Value: "véierzeg"},
		{Number: 39, Value: "néngandrësseg"},
		{Number: 38, Value: "aachtandrësseg"},
		{Number: 37, Value: "siwwenandrësseg"},
		{Number: 36, Value: "sechsandrësseg"},
		{Number: 35, Value: "fënnefandrësseg"},
		{Number: 34, Value: "véierandrësseg"},
		{Number: 33, Value: "dräiandrësseg"},
		{Number: 32, Value: "zwouandrësseg"},
		{Number: 31, Value: "eenandrësseg"},
		{Number: 30, Value: "drësseg"},
		{Number: 29, Value: "néngandzwanzeg"},
		{Number: 28, Value: "aachtandzwanzeg"},
		{Number: 27, Value: "siwwenandzwanzeg"},
		{Number: 26, Value: "sechsandzwanzeg"},
		{Number: 25, Value: "fënnefandzwanzeg"},
		{Number: 24, Value: "véierandzwanzeg"},
		{Number: 23, Value: "dräiandzwanzeg"},
		{Number: 22, Value: "zwouandzwanzeg"},
		{Number: 21, Value: "eenandzwanzeg"},
		{Number: 20, Value: "zwanzeg"},
		{Number: 19, Value: "nongzéng"},
		{Number: 18, Value: "aachtzéng"},
		{Number: 17, Value: "siwwenzéng"},
		{Number: 16, Value: "sechzéng"},
		{Number: 15, Value: "fënnefzéng"},
		{Number: 14, Value: "véierzéng"},
		{Number: 13, Value: "dräizéng"},
		{Number: 12, Value: "zwielef"},
		{Number: 11, Value: "eelef"},
		{Number: 10, Value: "zéng"},
		{Number: 9, Value: "néng"},
		{Number: 8, Value: "aacht"},
		{Number: 7, Value: "siwen"},
		{Number: 6, Value: "sechs"},
		{Number: 5, Value: "fënnef"},
		{Number: 4, Value: "véier"},
		{Number: 3, Value: "dräi"},
		{Number: 2, Value: "zwou"},
		{Number: 1, Value: "een"},
		{Number: 0, Value: "null"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Een Honnert"},
	},
	LocaleFormatter: &LuxembourgishFormatter{},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "éischten", Suffix: ".", Masculine: "éischten", Feminine: "éischt", Neuter: "éischt"},
		{Number: 2, Word: "zweeten", Suffix: ".", Masculine: "zweeten", Feminine: "zweet", Neuter: "zweet"},
		{Number: 3, Word: "drëtten", Suffix: ".", Masculine: "drëtten", Feminine: "drëtt", Neuter: "drëtt"},
		{Number: 4, Word: "véierten", Suffix: ".", Masculine: "véierten", Feminine: "véiert", Neuter: "véiert"},
		{Number: 5, Word: "fënneften", Suffix: ".", Masculine: "fënneften", Feminine: "fënneft", Neuter: "fënneft"},
		{Number: 6, Word: "sechsten", Suffix: ".", Masculine: "sechsten", Feminine: "sechst", Neuter: "sechst"},
		{Number: 7, Word: "siwenten", Suffix: ".", Masculine: "siwenten", Feminine: "siwent", Neuter: "siwent"},
		{Number: 8, Word: "aachten", Suffix: ".", Masculine: "aachten", Feminine: "aacht", Neuter: "aacht"},
		{Number: 9, Word: "néngten", Suffix: ".", Masculine: "néngten", Feminine: "néngt", Neuter: "néngt"},
		{Number: 10, Word: "zéngten", Suffix: ".", Masculine: "zéngten", Feminine: "zéngt", Neuter: "zéngt"},
		{Number: 11, Word: "eelften", Suffix: ".", Masculine: "eelften", Feminine: "eelft", Neuter: "eelft"},
		{Number: 12, Word: "zwieleften", Suffix: ".", Masculine: "zwieleften", Feminine: "zwieleft", Neuter: "zwieleft"},
		{Number: 13, Word: "dräizéngten", Suffix: ".", Masculine: "dräizéngten", Feminine: "dräizéngt", Neuter: "dräizéngt"},
		{Number: 14, Word: "véierzéngten", Suffix: ".", Masculine: "véierzéngten", Feminine: "véierzéngt", Neuter: "véierzéngt"},
		{Number: 15, Word: "fënnefzéngten", Suffix: ".", Masculine: "fënnefzéngten", Feminine: "fënnefzéngt", Neuter: "fënnefzéngt"},
		{Number: 16, Word: "sechzéngten", Suffix: ".", Masculine: "sechzéngten", Feminine: "sechzéngt", Neuter: "sechzéngt"},
		{Number: 17, Word: "siwwenzéngten", Suffix: ".", Masculine: "siwwenzéngten", Feminine: "siwwenzéngt", Neuter: "siwwenzéngt"},
		{Number: 18, Word: "aachtzéngten", Suffix: ".", Masculine: "aachtzéngten", Feminine: "aachtzéngt", Neuter: "aachtzéngt"},
		{Number: 19, Word: "nongzéngten", Suffix: ".", Masculine: "nongzéngten", Feminine: "nongzéngt", Neuter: "nongzéngt"},
		{Number: 20, Word: "zwanzegsten", Suffix: ".", Masculine: "zwanzegsten", Feminine: "zwanzegst", Neuter: "zwanzegst"},
		{Number: 21, Word: "eenandzwanzegsten", Suffix: ".", Masculine: "eenandzwanzegsten", Feminine: "eenandzwanzegst", Neuter: "eenandzwanzegst"},
		{Number: 30, Word: "drëssegsten", Suffix: ".", Masculine: "drëssegsten", Feminine: "drëssegst", Neuter: "drëssegst"},
		{Number: 40, Word: "véierzegsten", Suffix: ".", Masculine: "véierzegsten", Feminine: "véierzegst", Neuter: "véierzegst"},
		{Number: 50, Word: "fofzegsten", Suffix: ".", Masculine: "fofzegsten", Feminine: "fofzegst", Neuter: "fofzegst"},
		{Number: 60, Word: "sechzegsten", Suffix: ".", Masculine: "sechzegsten", Feminine: "sechzegst", Neuter: "sechzegst"},
		{Number: 70, Word: "siwwenzegsten", Suffix: ".", Masculine: "siwwenzegsten", Feminine: "siwwenzegst", Neuter: "siwwenzegst"},
		{Number: 80, Word: "aachtzegsten", Suffix: ".", Masculine: "aachtzegsten", Feminine: "aachtzegst", Neuter: "aachtzegst"},
		{Number: 90, Word: "nongzegsten", Suffix: ".", Masculine: "nongzegsten", Feminine: "nongzegst", Neuter: "nongzegst"},
		{Number: 100, Word: "honnertsten", Suffix: ".", Masculine: "honnertsten", Feminine: "honnertst", Neuter: "honnertst"},
		{Number: 1000, Word: "dausendsten", Suffix: ".", Masculine: "dausendsten", Feminine: "dausendst", Neuter: "dausendst"},
	},
}

// LuxembourgishFormatter handles Luxembourgish formatting
type LuxembourgishFormatter struct{}

func (f *LuxembourgishFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *LuxembourgishFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *LuxembourgishFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *LuxembourgishFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *LuxembourgishFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *LuxembourgishFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	return amount.Truncate(int32(precision))
}

func (f *LuxembourgishFormatter) FormatDecimalNumber(amount float64) string {
	return FormatEuropeanDecimal(amount)
}
func (f *LuxembourgishFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with European conventions (period separators, comma decimal, prefix symbol)
	return FormatEuropeanCurrency(amount, currencySymbol)
}
