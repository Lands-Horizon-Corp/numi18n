package locale

import (
	"github.com/shopspring/decimal"
)

// INHILocale is a NumI18NLocale configured for India (hi-IN)
var INHILocale = NumI18NLocale{
	Currency: Currency{
		Name:     "रुपया",
		Plural:   "रुपये",
		Singular: "रुपया",
		Symbol:   "₹",
		FractionUnit: FractionUnit{
			Name:     "पैसा",
			Plural:   "पैसे",
			Singular: "पैसा",
			Symbol:   "प",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "India",
		Currency:       "INR",
		ISO3166Alpha2:  "IN",
		ISO3166Alpha3:  "IND",
		ISO3166Numeric: "356",
		Locale:         "hi-IN",
		Timezone:       []string{"Asia/Kolkata"},
		Language:       "hi",
		Emoji:          "🇮🇳",
		PhoneCode:      "+91",
		Domain:         ".in",
	},
	Texts: Texts{
		And:   "और",
		Minus: "ऋण",
		Only:  "केवल",
		Point: "बिंदु",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "एक पद्म"},
		{Number: 1000000000000, Value: "एक नील"},
		{Number: 1000000000, Value: "एक अरब"},
		{Number: 10000000, Value: "एक करोड़"},
		{Number: 1000000, Value: "दस लाख"},
		{Number: 100000, Value: "एक लाख"},
		{Number: 1000, Value: "एक हजार"},
		{Number: 100, Value: "एक सौ"},
		{Number: 90, Value: "नब्बे"},
		{Number: 80, Value: "अस्सी"},
		{Number: 70, Value: "सत्तर"},
		{Number: 60, Value: "साठ"},
		{Number: 50, Value: "पचास"},
		{Number: 40, Value: "चालीस"},
		{Number: 30, Value: "तीस"},
		{Number: 20, Value: "बीस"},
		{Number: 19, Value: "उन्नीस"},
		{Number: 18, Value: "अठारह"},
		{Number: 17, Value: "सत्रह"},
		{Number: 16, Value: "सोलह"},
		{Number: 15, Value: "पंद्रह"},
		{Number: 14, Value: "चौदह"},
		{Number: 13, Value: "तेरह"},
		{Number: 12, Value: "बारह"},
		{Number: 11, Value: "ग्यारह"},
		{Number: 10, Value: "दस"},
		{Number: 9, Value: "नौ"},
		{Number: 8, Value: "आठ"},
		{Number: 7, Value: "सात"},
		{Number: 6, Value: "छह"},
		{Number: 5, Value: "पांच"},
		{Number: 4, Value: "चार"},
		{Number: 3, Value: "तीन"},
		{Number: 2, Value: "दो"},
		{Number: 1, Value: "एक"},
		{Number: 0, Value: "शून्य"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "एक सौ"},
		{Number: 100000, Value: "एक लाख"},
		{Number: 10000000, Value: "एक करोड़"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "पहला", Suffix: "वाँ", Masculine: "पहला", Feminine: "पहली", Neuter: "पहला"},
		{Number: 2, Word: "दूसरा", Suffix: "वाँ", Masculine: "दूसरा", Feminine: "दूसरी", Neuter: "दूसरा"},
		{Number: 3, Word: "तीसरा", Suffix: "वाँ", Masculine: "तीसरा", Feminine: "तीसरी", Neuter: "तीसरा"},
		{Number: 4, Word: "चौथा", Suffix: "वाँ", Masculine: "चौथा", Feminine: "चौथी", Neuter: "चौथा"},
		{Number: 5, Word: "पाँचवाँ", Suffix: "वाँ", Masculine: "पाँचवाँ", Feminine: "पाँचवीं", Neuter: "पाँचवाँ"},
		{Number: 6, Word: "छठा", Suffix: "वाँ", Masculine: "छठा", Feminine: "छठी", Neuter: "छठा"},
		{Number: 7, Word: "सातवाँ", Suffix: "वाँ", Masculine: "सातवाँ", Feminine: "सातवीं", Neuter: "सातवाँ"},
		{Number: 8, Word: "आठवाँ", Suffix: "वाँ", Masculine: "आठवाँ", Feminine: "आठवीं", Neuter: "आठवाँ"},
		{Number: 9, Word: "नौवाँ", Suffix: "वाँ", Masculine: "नौवाँ", Feminine: "नौवीं", Neuter: "नौवाँ"},
		{Number: 10, Word: "दसवाँ", Suffix: "वाँ", Masculine: "दसवाँ", Feminine: "दसवीं", Neuter: "दसवाँ"},
		{Number: 11, Word: "ग्यारहवाँ", Suffix: "वाँ", Masculine: "ग्यारहवाँ", Feminine: "ग्यारहवीं", Neuter: "ग्यारहवाँ"},
		{Number: 12, Word: "बारहवाँ", Suffix: "वाँ", Masculine: "बारहवाँ", Feminine: "बारहवीं", Neuter: "बारहवाँ"},
		{Number: 13, Word: "तेरहवाँ", Suffix: "वाँ", Masculine: "तेरहवाँ", Feminine: "तेरहवीं", Neuter: "तेरहवाँ"},
		{Number: 14, Word: "चौदहवाँ", Suffix: "वाँ", Masculine: "चौदहवाँ", Feminine: "चौदहवीं", Neuter: "चौदहवाँ"},
		{Number: 15, Word: "पंद्रहवाँ", Suffix: "वाँ", Masculine: "पंद्रहवाँ", Feminine: "पंद्रहवीं", Neuter: "पंद्रहवाँ"},
		{Number: 16, Word: "सोलहवाँ", Suffix: "वाँ", Masculine: "सोलहवाँ", Feminine: "सोलहवीं", Neuter: "सोलहवाँ"},
		{Number: 17, Word: "सत्रहवाँ", Suffix: "वाँ", Masculine: "सत्रहवाँ", Feminine: "सत्रहवीं", Neuter: "सत्रहवाँ"},
		{Number: 18, Word: "अठारहवाँ", Suffix: "वाँ", Masculine: "अठारहवाँ", Feminine: "अठारहवीं", Neuter: "अठारहवाँ"},
		{Number: 19, Word: "उन्नीसवाँ", Suffix: "वाँ", Masculine: "उन्नीसवाँ", Feminine: "उन्नीसवीं", Neuter: "उन्नीसवाँ"},
		{Number: 20, Word: "बीसवाँ", Suffix: "वाँ", Masculine: "बीसवाँ", Feminine: "बीसवीं", Neuter: "बीसवाँ"},
		{Number: 21, Word: "इक्कीसवाँ", Suffix: "वाँ", Masculine: "इक्कीसवाँ", Feminine: "इक्कीसवीं", Neuter: "इक्कीसवाँ"},
		{Number: 30, Word: "तीसवाँ", Suffix: "वाँ", Masculine: "तीसवाँ", Feminine: "तीसवीं", Neuter: "तीसवाँ"},
		{Number: 40, Word: "चालीसवाँ", Suffix: "वाँ", Masculine: "चालीसवाँ", Feminine: "चालीसवीं", Neuter: "चालीसवाँ"},
		{Number: 50, Word: "पचासवाँ", Suffix: "वाँ", Masculine: "पचासवाँ", Feminine: "पचासवीं", Neuter: "पचासवाँ"},
		{Number: 60, Word: "साठवाँ", Suffix: "वाँ", Masculine: "साठवाँ", Feminine: "साठवीं", Neuter: "साठवाँ"},
		{Number: 70, Word: "सत्तरवाँ", Suffix: "वाँ", Masculine: "सत्तरवाँ", Feminine: "सत्तरवीं", Neuter: "सत्तरवाँ"},
		{Number: 80, Word: "अस्सीवाँ", Suffix: "वाँ", Masculine: "अस्सीवाँ", Feminine: "अस्सीवीं", Neuter: "अस्सीवाँ"},
		{Number: 90, Word: "नब्बेवाँ", Suffix: "वाँ", Masculine: "नब्बेवाँ", Feminine: "नब्बेवीं", Neuter: "नब्बेवाँ"},
		{Number: 100, Word: "सौवाँ", Suffix: "वाँ", Masculine: "सौवाँ", Feminine: "सौवीं", Neuter: "सौवाँ"},
		{Number: 1000, Word: "हजारवाँ", Suffix: "वाँ", Masculine: "हजारवाँ", Feminine: "हजारवीं", Neuter: "हजारवाँ"},
	},
	LocaleFormatter: &HindiFormatter{},
}

// HindiFormatter handles Hindi (hi-IN) formatting
type HindiFormatter struct{}

func (f *HindiFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *HindiFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *HindiFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *HindiFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *HindiFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *HindiFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	// Standard decimal chopping - round to specified precision
	return amount.Round(int32(precision))
}

func (f *HindiFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}

func (f *HindiFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	return FormatAngloCurrency(amount, currencySymbol)
}
