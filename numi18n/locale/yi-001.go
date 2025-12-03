package locale

import (
	"github.com/shopspring/decimal"
)

// YI001Locale represents the Yiddish (World) locale
var YI001Locale = NumI18NLocale{
	LocaleFormatter: &YiddishFormatter{},
	Currency: Currency{
		Name:     "Dollar",
		Plural:   "דאָלערס",
		Singular: "דאָלער",
		Symbol:   "$",
		FractionUnit: FractionUnit{
			Name:     "Cent",
			Plural:   "צענט",
			Singular: "צענט",
			Symbol:   "¢",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "World",
		Currency:       "USD",
		ISO3166Alpha2:  "001",
		ISO3166Alpha3:  "WLD",
		ISO3166Numeric: "001",
		Locale:         "yi-001",
		Timezone:       []string{"UTC"},
		Language:       "yi",
		Emoji:          "🌍",
	},
	Texts: Texts{
		And:   "און",
		Minus: "מינוס",
		Only:  "בלויז",
		Point: "פּונקט",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "איין קוואַדריליאָן"},
		{Number: 1000000000000, Value: "איין טריליאָן"},
		{Number: 1000000000, Value: "איין ביליאָן"},
		{Number: 1000000, Value: "איין מיליאָן"},
		{Number: 100000, Value: "הונדערט טויזנט"},
		{Number: 90000, Value: "נײַנציק טויזנט"},
		{Number: 80000, Value: "אַכציק טויזנט"},
		{Number: 70000, Value: "זיבעציק טויזנט"},
		{Number: 60000, Value: "זעכציק טויזנט"},
		{Number: 50000, Value: "פֿופֿציק טויזנט"},
		{Number: 40000, Value: "פֿערציק טויזנט"},
		{Number: 30000, Value: "דרײַסיק טויזנט"},
		{Number: 20000, Value: "צוואַנציק טויזנט"},
		{Number: 19000, Value: "נײַנצן טויזנט"},
		{Number: 18000, Value: "אַכצן טויזנט"},
		{Number: 17000, Value: "זיבעצן טויזנט"},
		{Number: 16000, Value: "זעכצן טויזנט"},
		{Number: 15000, Value: "פֿופֿצן טויזנט"},
		{Number: 14000, Value: "פֿערצן טויזנט"},
		{Number: 13000, Value: "דרײַצן טויזנט"},
		{Number: 12000, Value: "צוועלף טויזנט"},
		{Number: 11000, Value: "עלף טויזנט"},
		{Number: 10000, Value: "צען טויזנט"},
		{Number: 9000, Value: "נײַן טויזנט"},
		{Number: 8000, Value: "אַכט טויזנט"},
		{Number: 7000, Value: "זיבן טויזנט"},
		{Number: 6000, Value: "זעקס טויזנט"},
		{Number: 5000, Value: "פֿינף טויזנט"},
		{Number: 4000, Value: "פֿיר טויזנט"},
		{Number: 3000, Value: "דרײַ טויזנט"},
		{Number: 2000, Value: "צוויי טויזנט"},
		{Number: 1000, Value: "איין טויזנט"},
		{Number: 900, Value: "נײַן הונדערט"},
		{Number: 800, Value: "אַכט הונדערט"},
		{Number: 700, Value: "זיבן הונדערט"},
		{Number: 600, Value: "זעקס הונדערט"},
		{Number: 500, Value: "פֿינף הונדערט"},
		{Number: 400, Value: "פֿיר הונדערט"},
		{Number: 300, Value: "דרײַ הונדערט"},
		{Number: 200, Value: "צוויי הונדערט"},
		{Number: 100, Value: "איין הונדערט"},
		{Number: 99, Value: "נײַנצן און אַכציק"},
		{Number: 98, Value: "אַכצן און אַכציק"},
		{Number: 97, Value: "זיבעצן און אַכציק"},
		{Number: 96, Value: "זעכצן און אַכציק"},
		{Number: 95, Value: "פֿופֿצן און אַכציק"},
		{Number: 94, Value: "פֿערצן און אַכציק"},
		{Number: 93, Value: "דרײַצן און אַכציק"},
		{Number: 92, Value: "צוועלף און אַכציק"},
		{Number: 91, Value: "עלף און אַכציק"},
		{Number: 90, Value: "נײַנציק"},
		{Number: 89, Value: "נײַן און אַכציק"},
		{Number: 88, Value: "אַכט און אַכציק"},
		{Number: 87, Value: "זיבן און אַכציק"},
		{Number: 86, Value: "זעקס און אַכציק"},
		{Number: 85, Value: "פֿינף און אַכציק"},
		{Number: 84, Value: "פֿיר און אַכציק"},
		{Number: 83, Value: "דרײַ און אַכציק"},
		{Number: 82, Value: "צוויי און אַכציק"},
		{Number: 81, Value: "איין און אַכציק"},
		{Number: 80, Value: "אַכציק"},
		{Number: 79, Value: "נײַן און זיבעציק"},
		{Number: 78, Value: "אַכט און זיבעציק"},
		{Number: 77, Value: "זיבן און זיבעציק"},
		{Number: 76, Value: "זעקס און זיבעציק"},
		{Number: 75, Value: "פֿינף און זיבעציק"},
		{Number: 74, Value: "פֿיר און זיבעציק"},
		{Number: 73, Value: "דרײַ און זיבעציק"},
		{Number: 72, Value: "צוויי און זיבעציק"},
		{Number: 71, Value: "איין און זיבעציק"},
		{Number: 70, Value: "זיבעציק"},
		{Number: 69, Value: "נײַן און זעכציק"},
		{Number: 68, Value: "אַכט און זעכציק"},
		{Number: 67, Value: "זיבן און זעכציק"},
		{Number: 66, Value: "זעקס און זעכציק"},
		{Number: 65, Value: "פֿינף און זעכציק"},
		{Number: 64, Value: "פֿיר און זעכציק"},
		{Number: 63, Value: "דרײַ און זעכציק"},
		{Number: 62, Value: "צוויי און זעכציק"},
		{Number: 61, Value: "איין און זעכציק"},
		{Number: 60, Value: "זעכציק"},
		{Number: 59, Value: "נײַן און פֿופֿציק"},
		{Number: 58, Value: "אַכט און פֿופֿציק"},
		{Number: 57, Value: "זיבן און פֿופֿציק"},
		{Number: 56, Value: "זעקס און פֿופֿציק"},
		{Number: 55, Value: "פֿינף און פֿופֿציק"},
		{Number: 54, Value: "פֿיר און פֿופֿציק"},
		{Number: 53, Value: "דרײַ און פֿופֿציק"},
		{Number: 52, Value: "צוויי און פֿופֿציק"},
		{Number: 51, Value: "איין און פֿופֿציק"},
		{Number: 50, Value: "פֿופֿציק"},
		{Number: 49, Value: "נײַן און פֿערציק"},
		{Number: 48, Value: "אַכט און פֿערציק"},
		{Number: 47, Value: "זיבן און פֿערציק"},
		{Number: 46, Value: "זעקס און פֿערציק"},
		{Number: 45, Value: "פֿינף און פֿערציק"},
		{Number: 44, Value: "פֿיר און פֿערציק"},
		{Number: 43, Value: "דרײַ און פֿערציק"},
		{Number: 42, Value: "צוויי און פֿערציק"},
		{Number: 41, Value: "איין און פֿערציק"},
		{Number: 40, Value: "פֿערציק"},
		{Number: 39, Value: "נײַן און דרײַסיק"},
		{Number: 38, Value: "אַכט און דרײַסיק"},
		{Number: 37, Value: "זיבן און דרײַסיק"},
		{Number: 36, Value: "זעקס און דרײַסיק"},
		{Number: 35, Value: "פֿינף און דרײַסיק"},
		{Number: 34, Value: "פֿיר און דרײַסיק"},
		{Number: 33, Value: "דרײַ און דרײַסיק"},
		{Number: 32, Value: "צוויי און דרײַסיק"},
		{Number: 31, Value: "איין און דרײַסיק"},
		{Number: 30, Value: "דרײַסיק"},
		{Number: 29, Value: "נײַן און צוואַנציק"},
		{Number: 28, Value: "אַכט און צוואַנציק"},
		{Number: 27, Value: "זיבן און צוואַנציק"},
		{Number: 26, Value: "זעקס און צוואַנציק"},
		{Number: 25, Value: "פֿינף און צוואַנציק"},
		{Number: 24, Value: "פֿיר און צוואַנציק"},
		{Number: 23, Value: "דרײַ און צוואַנציק"},
		{Number: 22, Value: "צוויי און צוואַנציק"},
		{Number: 21, Value: "איין און צוואַנציק"},
		{Number: 20, Value: "צוואַנציק"},
		{Number: 19, Value: "נײַנצן"},
		{Number: 18, Value: "אַכצן"},
		{Number: 17, Value: "זיבעצן"},
		{Number: 16, Value: "זעכצן"},
		{Number: 15, Value: "פֿופֿצן"},
		{Number: 14, Value: "פֿערצן"},
		{Number: 13, Value: "דרײַצן"},
		{Number: 12, Value: "צוועלף"},
		{Number: 11, Value: "עלף"},
		{Number: 10, Value: "צען"},
		{Number: 9, Value: "נײַן"},
		{Number: 8, Value: "אַכט"},
		{Number: 7, Value: "זיבן"},
		{Number: 6, Value: "זעקס"},
		{Number: 5, Value: "פֿינף"},
		{Number: 4, Value: "פֿיר"},
		{Number: 3, Value: "דרײַ"},
		{Number: 2, Value: "צוויי"},
		{Number: 1, Value: "איין"},
		{Number: 0, Value: "נול"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "הונדערט"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "ערשטער", Suffix: "-טער", Masculine: "ערשטער", Feminine: "ערשטע", Neuter: "ערשטע"},
		{Number: 2, Word: "צווייטער", Suffix: "-טער", Masculine: "צווייטער", Feminine: "צווייטע", Neuter: "צווייטע"},
		{Number: 3, Word: "דריטער", Suffix: "-טער", Masculine: "דריטער", Feminine: "דריטע", Neuter: "דריטע"},
		{Number: 4, Word: "פֿירטער", Suffix: "-טער", Masculine: "פֿירטער", Feminine: "פֿירטע", Neuter: "פֿירטע"},
		{Number: 5, Word: "פֿינפֿטער", Suffix: "-טער", Masculine: "פֿינפֿטער", Feminine: "פֿינפֿטע", Neuter: "פֿינפֿטע"},
		{Number: 6, Word: "זעקסטער", Suffix: "-טער", Masculine: "זעקסטער", Feminine: "זעקסטע", Neuter: "זעקסטע"},
		{Number: 7, Word: "זיבנטער", Suffix: "-טער", Masculine: "זיבנטער", Feminine: "זיבנטע", Neuter: "זיבנטע"},
		{Number: 8, Word: "אַכטער", Suffix: "-טער", Masculine: "אַכטער", Feminine: "אַכטע", Neuter: "אַכטע"},
		{Number: 9, Word: "נײַנטער", Suffix: "-טער", Masculine: "נײַנטער", Feminine: "נײַנטע", Neuter: "נײַנטע"},
		{Number: 10, Word: "צענטער", Suffix: "-טער", Masculine: "צענטער", Feminine: "צענטע", Neuter: "צענטע"},
		{Number: 11, Word: "עלפֿטער", Suffix: "-טער", Masculine: "עלפֿטער", Feminine: "עלפֿטע", Neuter: "עלפֿטע"},
		{Number: 12, Word: "צוועלפֿטער", Suffix: "-טער", Masculine: "צוועלפֿטער", Feminine: "צוועלפֿטע", Neuter: "צוועלפֿטע"},
		{Number: 13, Word: "דרײַצנטער", Suffix: "-טער", Masculine: "דרײַצנטער", Feminine: "דרײַצנטע", Neuter: "דרײַצנטע"},
		{Number: 14, Word: "פֿערצנטער", Suffix: "-טער", Masculine: "פֿערצנטער", Feminine: "פֿערצנטע", Neuter: "פֿערצנטע"},
		{Number: 15, Word: "פֿופֿצנטער", Suffix: "-טער", Masculine: "פֿופֿצנטער", Feminine: "פֿופֿצנטע", Neuter: "פֿופֿצנטע"},
		{Number: 16, Word: "זעכצנטער", Suffix: "-טער", Masculine: "זעכצנטער", Feminine: "זעכצנטע", Neuter: "זעכצנטע"},
		{Number: 17, Word: "זיבעצנטער", Suffix: "-טער", Masculine: "זיבעצנטער", Feminine: "זיבעצנטע", Neuter: "זיבעצנטע"},
		{Number: 18, Word: "אַכצנטער", Suffix: "-טער", Masculine: "אַכצנטער", Feminine: "אַכצנטע", Neuter: "אַכצנטע"},
		{Number: 19, Word: "נײַנצנטער", Suffix: "-טער", Masculine: "נײַנצנטער", Feminine: "נײַנצנטע", Neuter: "נײַנצנטע"},
		{Number: 20, Word: "צוואַנציקסטער", Suffix: "-טער", Masculine: "צוואַנציקסטער", Feminine: "צוואַנציקסטע", Neuter: "צוואַנציקסטע"},
		{Number: 21, Word: "איין און צוואַנציקסטער", Suffix: "-טער", Masculine: "איין און צוואַנציקסטער", Feminine: "איין און צוואַנציקסטע", Neuter: "איין און צוואַנציקסטע"},
		{Number: 30, Word: "דרײַסיקסטער", Suffix: "-טער", Masculine: "דרײַסיקסטער", Feminine: "דרײַסיקסטע", Neuter: "דרײַסיקסטע"},
		{Number: 40, Word: "פֿערציקסטער", Suffix: "-טער", Masculine: "פֿערציקסטער", Feminine: "פֿערציקסטע", Neuter: "פֿערציקסטע"},
		{Number: 50, Word: "פֿופֿציקסטער", Suffix: "-טער", Masculine: "פֿופֿציקסטער", Feminine: "פֿופֿציקסטע", Neuter: "פֿופֿציקסטע"},
		{Number: 60, Word: "זעכציקסטער", Suffix: "-טער", Masculine: "זעכציקסטער", Feminine: "זעכציקסטע", Neuter: "זעכציקסטע"},
		{Number: 70, Word: "זיבעציקסטער", Suffix: "-טער", Masculine: "זיבעציקסטער", Feminine: "זיבעציקסטע", Neuter: "זיבעציקסטע"},
		{Number: 80, Word: "אַכציקסטער", Suffix: "-טער", Masculine: "אַכציקסטער", Feminine: "אַכציקסטע", Neuter: "אַכציקסטע"},
		{Number: 90, Word: "נײַנציקסטער", Suffix: "-טער", Masculine: "נײַנציקסטער", Feminine: "נײַנציקסטע", Neuter: "נײַנציקסטע"},
		{Number: 100, Word: "הונדערטסטער", Suffix: "-טער", Masculine: "הונדערטסטער", Feminine: "הונדערטסטע", Neuter: "הונדערטסטע"},
		{Number: 1000, Word: "טויזנטסטער", Suffix: "-טער", Masculine: "טויזנטסטער", Feminine: "טויזנטסטע", Neuter: "טויזנטסטע"},
		{Number: 1000000, Word: "מיליאָנסטער", Suffix: "-טער", Masculine: "מיליאָנסטער", Feminine: "מיליאָנסטע", Neuter: "מיליאָנסטע"},
		{Number: 1000000000, Word: "ביליאָנסטער", Suffix: "-טער", Masculine: "ביליאָנסטער", Feminine: "ביליאָנסטע", Neuter: "ביליאָנסטע"},
	},
}

// YiddishFormatter handles Yiddish (World) formatting
type YiddishFormatter struct{}

func (f *YiddishFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *YiddishFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *YiddishFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *YiddishFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *YiddishFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *YiddishFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	multiplier := decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(precision)))
	return amount.Mul(multiplier).Truncate(0).Div(multiplier)
}

func (f *YiddishFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *YiddishFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
