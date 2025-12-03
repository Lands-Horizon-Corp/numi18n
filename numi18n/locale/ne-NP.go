package locale

import (
	"github.com/shopspring/decimal"
)

// NENPLocale represents the Nepali (Nepal) locale
var NENPLocale = NumI18NLocale{

	Currency: Currency{
		Name:     "रुपैया",
		Plural:   "रुपैया",
		Singular: "रुपैया",
		Symbol:   "रु",
		FractionUnit: FractionUnit{
			Name:     "पैसा",
			Plural:   "पैसा",
			Singular: "पैसा",
			Symbol:   "प",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Nepal",
		Currency:       "NPR",
		ISO3166Alpha2:  "NP",
		ISO3166Alpha3:  "NPL",
		ISO3166Numeric: "524",
		Locale:         "ne-NP",
		Timezone:       []string{"Asia/Kathmandu"},
		Language:       "ne",
		Emoji:          "🇳🇵",
		PhoneCode:      "+977",
		Domain:         ".np",
	},
	Texts: Texts{
		And:   "र",
		Minus: "माइनस",
		Only:  "मात्र",
		Point: "बिन्दु",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 10000000000000000, Value: "दश खरब"},
		{Number: 1000000000000000, Value: "एक खरब"},
		{Number: 100000000000000, Value: "दश नील"},
		{Number: 10000000000000, Value: "एक नील"},
		{Number: 1000000000000, Value: "दश अर्ब"},
		{Number: 100000000000, Value: "एक अर्ब"},
		{Number: 10000000000, Value: "दश अरब"},
		{Number: 1000000000, Value: "एक अरब"},
		{Number: 100000000, Value: "दश करोड"},
		{Number: 10000000, Value: "एक करोड"},
		{Number: 1000000, Value: "दश लाख"},
		{Number: 100000, Value: "एक लाख"},
		{Number: 90000, Value: "नब्बे हजार"},
		{Number: 80000, Value: "अस्सी हजार"},
		{Number: 70000, Value: "सत्तरी हजार"},
		{Number: 60000, Value: "साठी हजार"},
		{Number: 50000, Value: "पचास हजार"},
		{Number: 40000, Value: "चालिस हजार"},
		{Number: 30000, Value: "तिस हजार"},
		{Number: 20000, Value: "बिस हजार"},
		{Number: 19000, Value: "उन्नाइस हजार"},
		{Number: 18000, Value: "अठार हजार"},
		{Number: 17000, Value: "सत्र हजार"},
		{Number: 16000, Value: "सोह्र हजार"},
		{Number: 15000, Value: "पन्ध्र हजार"},
		{Number: 14000, Value: "चौध हजार"},
		{Number: 13000, Value: "तेह्र हजार"},
		{Number: 12000, Value: "बाह्र हजार"},
		{Number: 11000, Value: "एघार हजार"},
		{Number: 10000, Value: "दस हजार"},
		{Number: 9000, Value: "नौ हजार"},
		{Number: 8000, Value: "आठ हजार"},
		{Number: 7000, Value: "सात हजार"},
		{Number: 6000, Value: "छ हजार"},
		{Number: 5000, Value: "पाँच हजार"},
		{Number: 4000, Value: "चार हजार"},
		{Number: 3000, Value: "तीन हजार"},
		{Number: 2000, Value: "दुई हजार"},
		{Number: 1000, Value: "एक हजार"},
		{Number: 900, Value: "नौ सय"},
		{Number: 800, Value: "आठ सय"},
		{Number: 700, Value: "सात सय"},
		{Number: 600, Value: "छ सय"},
		{Number: 500, Value: "पाँच सय"},
		{Number: 400, Value: "चार सय"},
		{Number: 300, Value: "तीन सय"},
		{Number: 200, Value: "दुई सय"},
		{Number: 100, Value: "एक सय"},
		{Number: 99, Value: "उनान्सय"},
		{Number: 98, Value: "अन्ठानबे"},
		{Number: 97, Value: "सन्तानबे"},
		{Number: 96, Value: "छयानबे"},
		{Number: 95, Value: "पन्चानबे"},
		{Number: 94, Value: "चौरानबे"},
		{Number: 93, Value: "त्रिरानबे"},
		{Number: 92, Value: "बयानबे"},
		{Number: 91, Value: "एकानबे"},
		{Number: 90, Value: "नब्बे"},
		{Number: 89, Value: "उनान्सी"},
		{Number: 88, Value: "अठ्ठासी"},
		{Number: 87, Value: "सतासी"},
		{Number: 86, Value: "छयासी"},
		{Number: 85, Value: "पचासी"},
		{Number: 84, Value: "चौरासी"},
		{Number: 83, Value: "त्रिरासी"},
		{Number: 82, Value: "बयासी"},
		{Number: 81, Value: "एकासी"},
		{Number: 80, Value: "अस्सी"},
		{Number: 79, Value: "उनान्सत्तरी"},
		{Number: 78, Value: "अठ्ठहत्तर"},
		{Number: 77, Value: "सतहत्तर"},
		{Number: 76, Value: "छहत्तर"},
		{Number: 75, Value: "पचहत्तर"},
		{Number: 74, Value: "चौहत्तर"},
		{Number: 73, Value: "त्रिहत्तर"},
		{Number: 72, Value: "बहत्तर"},
		{Number: 71, Value: "एकहत्तर"},
		{Number: 70, Value: "सत्तरी"},
		{Number: 69, Value: "उनान्सत्तरी"},
		{Number: 68, Value: "अठ्ठसट्ठी"},
		{Number: 67, Value: "सतसट्ठी"},
		{Number: 66, Value: "छैसट्ठी"},
		{Number: 65, Value: "पैंसट्ठी"},
		{Number: 64, Value: "चौंसट्ठी"},
		{Number: 63, Value: "त्रिसट्ठी"},
		{Number: 62, Value: "बैसट्ठी"},
		{Number: 61, Value: "एकसट्ठी"},
		{Number: 60, Value: "साठी"},
		{Number: 59, Value: "उनान्साठी"},
		{Number: 58, Value: "अन्ठाउन्न"},
		{Number: 57, Value: "सन्ताउन्न"},
		{Number: 56, Value: "छप्पन्न"},
		{Number: 55, Value: "पचपन्न"},
		{Number: 54, Value: "चौवन्न"},
		{Number: 53, Value: "त्रिपन्न"},
		{Number: 52, Value: "बाउन्न"},
		{Number: 51, Value: "एकाउन्न"},
		{Number: 50, Value: "पचास"},
		{Number: 49, Value: "उनान्चास"},
		{Number: 48, Value: "अठचालिस"},
		{Number: 47, Value: "सत्चालिस"},
		{Number: 46, Value: "छयालिस"},
		{Number: 45, Value: "पैंतालिस"},
		{Number: 44, Value: "चौवालिस"},
		{Number: 43, Value: "त्रिचालिस"},
		{Number: 42, Value: "बयालिस"},
		{Number: 41, Value: "एकचालिस"},
		{Number: 40, Value: "चालिस"},
		{Number: 39, Value: "उनान्चालिस"},
		{Number: 38, Value: "अठतिस"},
		{Number: 37, Value: "सैंतिस"},
		{Number: 36, Value: "छत्तिस"},
		{Number: 35, Value: "पैंतिस"},
		{Number: 34, Value: "चौंतिस"},
		{Number: 33, Value: "तेत्तिस"},
		{Number: 32, Value: "बत्तिस"},
		{Number: 31, Value: "एकतिस"},
		{Number: 30, Value: "तिस"},
		{Number: 29, Value: "उनान्तिस"},
		{Number: 28, Value: "अठ्ठाइस"},
		{Number: 27, Value: "सत्ताइस"},
		{Number: 26, Value: "छब्बिस"},
		{Number: 25, Value: "पच्चिस"},
		{Number: 24, Value: "चौबिस"},
		{Number: 23, Value: "तेइस"},
		{Number: 22, Value: "बाइस"},
		{Number: 21, Value: "एकाइस"},
		{Number: 20, Value: "बिस"},
		{Number: 19, Value: "उन्नाइस"},
		{Number: 18, Value: "अठार"},
		{Number: 17, Value: "सत्र"},
		{Number: 16, Value: "सोह्र"},
		{Number: 15, Value: "पन्ध्र"},
		{Number: 14, Value: "चौध"},
		{Number: 13, Value: "तेह्र"},
		{Number: 12, Value: "बाह्र"},
		{Number: 11, Value: "एघार"},
		{Number: 10, Value: "दस"},
		{Number: 9, Value: "नौ"},
		{Number: 8, Value: "आठ"},
		{Number: 7, Value: "सात"},
		{Number: 6, Value: "छ"},
		{Number: 5, Value: "पाँच"},
		{Number: 4, Value: "चार"},
		{Number: 3, Value: "तीन"},
		{Number: 2, Value: "दुई"},
		{Number: 1, Value: "एक"},
		{Number: 0, Value: "शून्य"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "एक सय"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "पहिलो", Suffix: "-औं", Masculine: "पहिलो", Feminine: "पहिली", Neuter: "पहिलो"},
		{Number: 2, Word: "दोस्रो", Suffix: "-औं", Masculine: "दोस्रो", Feminine: "दोस्री", Neuter: "दोस्रो"},
		{Number: 3, Word: "तेस्रो", Suffix: "-औं", Masculine: "तेस्रो", Feminine: "तेस्री", Neuter: "तेस्रो"},
		{Number: 4, Word: "चौथो", Suffix: "-औं", Masculine: "चौथो", Feminine: "चौथी", Neuter: "चौथो"},
		{Number: 5, Word: "पाँचौं", Suffix: "-औं", Masculine: "पाँचौं", Feminine: "पाँचौं", Neuter: "पाँचौं"},
		{Number: 6, Word: "छैठौं", Suffix: "-औं", Masculine: "छैठौं", Feminine: "छैठौं", Neuter: "छैठौं"},
		{Number: 7, Word: "सातौं", Suffix: "-औं", Masculine: "सातौं", Feminine: "सातौं", Neuter: "सातौं"},
		{Number: 8, Word: "आठौं", Suffix: "-औं", Masculine: "आठौं", Feminine: "आठौं", Neuter: "आठौं"},
		{Number: 9, Word: "नवौं", Suffix: "-औं", Masculine: "नवौं", Feminine: "नवौं", Neuter: "नवौं"},
		{Number: 10, Word: "दशौं", Suffix: "-औं", Masculine: "दशौं", Feminine: "दशौं", Neuter: "दशौं"},
		{Number: 11, Word: "एघारौं", Suffix: "-औं", Masculine: "एघारौं", Feminine: "एघारौं", Neuter: "एघारौं"},
		{Number: 12, Word: "बाह्रौं", Suffix: "-औं", Masculine: "बाह्रौं", Feminine: "बाह्रौं", Neuter: "बाह्रौं"},
		{Number: 20, Word: "बिसौं", Suffix: "-औं", Masculine: "बिसौं", Feminine: "बिसौं", Neuter: "बिसौं"},
		{Number: 21, Word: "एकाइसौं", Suffix: "-औं", Masculine: "एकाइसौं", Feminine: "एकाइसौं", Neuter: "एकाइसौं"},
		{Number: 30, Word: "तिसौं", Suffix: "-औं", Masculine: "तिसौं", Feminine: "तिसौं", Neuter: "तिसौं"},
		{Number: 100, Word: "सयौं", Suffix: "-औं", Masculine: "सयौं", Feminine: "सयौं", Neuter: "सयौं"},
		{Number: 1000, Word: "हजारौं", Suffix: "-औं", Masculine: "हजारौं", Feminine: "हजारौं", Neuter: "हजारौं"},
	},
	LocaleFormatter: &NepaliFormatter{},
}

// NepaliFormatter handles Nepali-specific formatting
type NepaliFormatter struct{}

func (f *NepaliFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *NepaliFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *NepaliFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *NepaliFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *NepaliFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *NepaliFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 0
	}
	return amount.Truncate(int32(precision))
}

func (f *NepaliFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *NepaliFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
