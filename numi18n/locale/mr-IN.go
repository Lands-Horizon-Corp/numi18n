package locale

import (
	"github.com/shopspring/decimal"
)

// MRINLocale represents the Marathi (India) locale
var MRINLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "रुपया",
		Plural:   "रुपये",
		Singular: "रुपया",
		Symbol:   "₹",
		FractionUnit: FractionUnit{
			Name:     "पैसा",
			Plural:   "पैसे",
			Singular: "पैसा",
			Symbol:   "p",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "India",
		Currency:       "INR",
		ISO3166Alpha2:  "IN",
		ISO3166Alpha3:  "IND",
		ISO3166Numeric: "356",
		Locale:         "mr-IN",
		Timezone:       []string{"Asia/Kolkata"},
		Language:       "mr",
		Emoji:          "🇮🇳",
		PhoneCode:      "+91",
		Domain:         ".in",
	},
	Texts: Texts{
		And:   "आणि",
		Minus: "वजा",
		Only:  "फक्त",
		Point: "बिंदू",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 10000000000000000, Value: "दहा पद्म"},
		{Number: 1000000000000000, Value: "एक पद्म"},
		{Number: 100000000000000, Value: "दहा नील"},
		{Number: 10000000000000, Value: "एक नील"},
		{Number: 1000000000000, Value: "दहा खर्व"},
		{Number: 100000000000, Value: "एक खर्व"},
		{Number: 10000000000, Value: "दहा अब्ज"},
		{Number: 1000000000, Value: "एक अब्ज"},
		{Number: 100000000, Value: "दहा कोटी"},
		{Number: 10000000, Value: "एक कोटी"},
		{Number: 1000000, Value: "दहा लाख"},
		{Number: 100000, Value: "एक लाख"},
		{Number: 90000, Value: "नव्वद हजार"},
		{Number: 80000, Value: "ऐंशी हजार"},
		{Number: 70000, Value: "सत्तर हजार"},
		{Number: 60000, Value: "साठ हजार"},
		{Number: 50000, Value: "पन्नास हजार"},
		{Number: 40000, Value: "चाळीस हजार"},
		{Number: 30000, Value: "तीस हजार"},
		{Number: 20000, Value: "वीस हजार"},
		{Number: 19000, Value: "एकोणीस हजार"},
		{Number: 18000, Value: "अठरा हजार"},
		{Number: 17000, Value: "सतरा हजार"},
		{Number: 16000, Value: "सोळा हजार"},
		{Number: 15000, Value: "पंधरा हजार"},
		{Number: 14000, Value: "चौदा हजार"},
		{Number: 13000, Value: "तेरा हजार"},
		{Number: 12000, Value: "बारा हजार"},
		{Number: 11000, Value: "अकरा हजार"},
		{Number: 10000, Value: "दहा हजार"},
		{Number: 9000, Value: "नऊ हजार"},
		{Number: 8000, Value: "आठ हजार"},
		{Number: 7000, Value: "सात हजार"},
		{Number: 6000, Value: "सहा हजार"},
		{Number: 5000, Value: "पाच हजार"},
		{Number: 4000, Value: "चार हजार"},
		{Number: 3000, Value: "तीन हजार"},
		{Number: 2000, Value: "दोन हजार"},
		{Number: 1000, Value: "एक हजार"},
		{Number: 900, Value: "नऊशे"},
		{Number: 800, Value: "आठशे"},
		{Number: 700, Value: "सातशे"},
		{Number: 600, Value: "सहाशे"},
		{Number: 500, Value: "पाचशे"},
		{Number: 400, Value: "चारशे"},
		{Number: 300, Value: "तीनशे"},
		{Number: 200, Value: "दोनशे"},
		{Number: 100, Value: "एकशे"},
		{Number: 99, Value: "नव्याण्णव"},
		{Number: 98, Value: "अठ्ठ्याण्णव"},
		{Number: 97, Value: "सव्याण्णव"},
		{Number: 96, Value: "शहाण्णव"},
		{Number: 95, Value: "पंच्याण्णव"},
		{Number: 94, Value: "चौऱ्याण्णव"},
		{Number: 93, Value: "त्र्याण्णव"},
		{Number: 92, Value: "ब्याण्णव"},
		{Number: 91, Value: "एक्याण्णव"},
		{Number: 90, Value: "नव्वद"},
		{Number: 89, Value: "एकोणनव्वद"},
		{Number: 88, Value: "अठ्ठयाशी"},
		{Number: 87, Value: "सत्त्याशी"},
		{Number: 86, Value: "शहयाशी"},
		{Number: 85, Value: "पंच्याशी"},
		{Number: 84, Value: "चौऱ्याशी"},
		{Number: 83, Value: "त्र्याशी"},
		{Number: 82, Value: "ब्याशी"},
		{Number: 81, Value: "एक्याशी"},
		{Number: 80, Value: "ऐंशी"},
		{Number: 79, Value: "एकोणऐंशी"},
		{Number: 78, Value: "अठ्याहत्तर"},
		{Number: 77, Value: "सत्याहत्तर"},
		{Number: 76, Value: "शहात्तर"},
		{Number: 75, Value: "पंचाहत्तर"},
		{Number: 74, Value: "चौऱ्याहत्तर"},
		{Number: 73, Value: "त्र्याहत्तर"},
		{Number: 72, Value: "बहात्तर"},
		{Number: 71, Value: "एकाहत्तर"},
		{Number: 70, Value: "सत्तर"},
		{Number: 69, Value: "एकोणसत्तर"},
		{Number: 68, Value: "अडुसष्ट"},
		{Number: 67, Value: "सदुसष्ट"},
		{Number: 66, Value: "सहासष्ट"},
		{Number: 65, Value: "पासष्ट"},
		{Number: 64, Value: "चौसष्ट"},
		{Number: 63, Value: "त्रेसष्ट"},
		{Number: 62, Value: "बासष्ट"},
		{Number: 61, Value: "एकसष्ट"},
		{Number: 60, Value: "साठ"},
		{Number: 59, Value: "एकोणसाठ"},
		{Number: 58, Value: "अठ्ठावन्न"},
		{Number: 57, Value: "सत्तावन्न"},
		{Number: 56, Value: "छप्पन्न"},
		{Number: 55, Value: "पंचावन्न"},
		{Number: 54, Value: "चोपन्न"},
		{Number: 53, Value: "त्रेपन्न"},
		{Number: 52, Value: "बावन्न"},
		{Number: 51, Value: "एकावन्न"},
		{Number: 50, Value: "पन्नास"},
		{Number: 49, Value: "एकोणपन्नास"},
		{Number: 48, Value: "अठ्ठेचाळीस"},
		{Number: 47, Value: "सत्तेचाळीस"},
		{Number: 46, Value: "सेहेचाळीस"},
		{Number: 45, Value: "पंचेचाळीस"},
		{Number: 44, Value: "चव्वेचाळीस"},
		{Number: 43, Value: "त्रेचाळीस"},
		{Number: 42, Value: "बेचाळीस"},
		{Number: 41, Value: "एकेचाळीस"},
		{Number: 40, Value: "चाळीस"},
		{Number: 39, Value: "एकोणचाळीस"},
		{Number: 38, Value: "अडतीस"},
		{Number: 37, Value: "सदतीस"},
		{Number: 36, Value: "छत्तीस"},
		{Number: 35, Value: "पस्तीस"},
		{Number: 34, Value: "चौतीस"},
		{Number: 33, Value: "तेहतीस"},
		{Number: 32, Value: "बत्तीस"},
		{Number: 31, Value: "एकतीस"},
		{Number: 30, Value: "तीस"},
		{Number: 29, Value: "एकोणतीस"},
		{Number: 28, Value: "अठ्ठावीस"},
		{Number: 27, Value: "सत्तावीस"},
		{Number: 26, Value: "सव्वीस"},
		{Number: 25, Value: "पंचवीस"},
		{Number: 24, Value: "चोवीस"},
		{Number: 23, Value: "तेवीस"},
		{Number: 22, Value: "बावीस"},
		{Number: 21, Value: "एकवीस"},
		{Number: 20, Value: "वीस"},
		{Number: 19, Value: "एकोणीस"},
		{Number: 18, Value: "अठरा"},
		{Number: 17, Value: "सतरा"},
		{Number: 16, Value: "सोळा"},
		{Number: 15, Value: "पंधरा"},
		{Number: 14, Value: "चौदा"},
		{Number: 13, Value: "तेरा"},
		{Number: 12, Value: "बारा"},
		{Number: 11, Value: "अकरा"},
		{Number: 10, Value: "दहा"},
		{Number: 9, Value: "नऊ"},
		{Number: 8, Value: "आठ"},
		{Number: 7, Value: "सात"},
		{Number: 6, Value: "सहा"},
		{Number: 5, Value: "पाच"},
		{Number: 4, Value: "चार"},
		{Number: 3, Value: "तीन"},
		{Number: 2, Value: "दोन"},
		{Number: 1, Value: "एक"},
		{Number: 0, Value: "शून्य"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "एक शंभर"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "पहिला", Suffix: "-वा", Masculine: "पहिला", Feminine: "पहिली", Neuter: "पहिले"},
		{Number: 2, Word: "दुसरा", Suffix: "-रा", Masculine: "दुसरा", Feminine: "दुसरी", Neuter: "दुसरे"},
		{Number: 3, Word: "तिसरा", Suffix: "-रा", Masculine: "तिसरा", Feminine: "तिसरी", Neuter: "तिसरे"},
		{Number: 4, Word: "चौथा", Suffix: "-वा", Masculine: "चौथा", Feminine: "चौथी", Neuter: "चौथे"},
		{Number: 5, Word: "पाचवा", Suffix: "-वा", Masculine: "पाचवा", Feminine: "पाचवी", Neuter: "पाचवे"},
		{Number: 6, Word: "सहावा", Suffix: "-वा", Masculine: "सहावा", Feminine: "सहावी", Neuter: "सहावे"},
		{Number: 7, Word: "सातवा", Suffix: "-वा", Masculine: "सातवा", Feminine: "सातवी", Neuter: "सातवे"},
		{Number: 8, Word: "आठवा", Suffix: "-वा", Masculine: "आठवा", Feminine: "आठवी", Neuter: "आठवे"},
		{Number: 9, Word: "नववा", Suffix: "-वा", Masculine: "नववा", Feminine: "नववी", Neuter: "नववे"},
		{Number: 10, Word: "दहावा", Suffix: "-वा", Masculine: "दहावा", Feminine: "दहावी", Neuter: "दहावे"},
		{Number: 11, Word: "अकरावा", Suffix: "-वा", Masculine: "अकरावा", Feminine: "अकरावी", Neuter: "अकरावे"},
		{Number: 12, Word: "बारावा", Suffix: "-वा", Masculine: "बारावा", Feminine: "बारावी", Neuter: "बारावे"},
		{Number: 20, Word: "वीसावा", Suffix: "-वा", Masculine: "वीसावा", Feminine: "वीसावी", Neuter: "वीसावे"},
		{Number: 21, Word: "एकविसावा", Suffix: "-वा", Masculine: "एकविसावा", Feminine: "एकविसावी", Neuter: "एकविसावे"},
		{Number: 30, Word: "तीसावा", Suffix: "-वा", Masculine: "तीसावा", Feminine: "तीसावी", Neuter: "तीसावे"},
		{Number: 100, Word: "शंभरावा", Suffix: "-वा", Masculine: "शंभरावा", Feminine: "शंभरावी", Neuter: "शंभरावे"},
		{Number: 1000, Word: "हजारावा", Suffix: "-वा", Masculine: "हजारावा", Feminine: "हजारावी", Neuter: "हजारावे"},
	},
	LocaleFormatter: &MarathiFormatter{},
}

// MarathiFormatter handles Marathi-specific formatting
type MarathiFormatter struct{}

func (f *MarathiFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *MarathiFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *MarathiFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *MarathiFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *MarathiFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *MarathiFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 0
	}
	return amount.Truncate(int32(precision))
}

func (f *MarathiFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *MarathiFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
