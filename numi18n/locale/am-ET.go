package locale

import (
	"github.com/shopspring/decimal"
)

// AMETLocale is a NumI18NLocale configured for Amharic (Ethiopia) - am-ET
var AMETLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "ብር",
		Plural:   "ብር",
		Singular: "ብር",
		Symbol:   "Br",
		FractionUnit: FractionUnit{
			Name:     "ሚልስ",
			Plural:   "ሚልስ",
			Singular: "ሚልስ",
			Symbol:   "ms",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Ethiopia",
		Currency:       "ETB",
		ISO3166Alpha2:  "ET",
		ISO3166Alpha3:  "ETH",
		ISO3166Numeric: "231",
		Locale:         "am-ET",
		Timezone:       []string{"Africa/Addis_Ababa"},
		Language:       "am",
		Emoji:          "🇪🇹",
		PhoneCode:      "+251",
		Domain:         ".et",
	},
	Texts: Texts{
		And:   "እና",
		Minus: "አልፎ",
		Only:  "ብቻ",
		Point: "ነጥብ",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "ኩዋድሪሊዮን"},
		{Number: 1000000000000, Value: "ትሪሊዮን"},
		{Number: 1000000000, Value: "ቢሊዮን"},
		{Number: 1000000, Value: "ሚሊዮን"},
		{Number: 1000, Value: "ሺህ"},
		{Number: 100, Value: "መቶ"},
		{Number: 90, Value: "ዘጠና"},
		{Number: 80, Value: "ሰማንያ"},
		{Number: 70, Value: "ሰባ"},
		{Number: 60, Value: "ስድስት ዐሥር"},
		{Number: 50, Value: "ሃምሳ"},
		{Number: 40, Value: "አርባ"},
		{Number: 30, Value: "ሰላሳ"},
		{Number: 20, Value: "ሃያ"},
		{Number: 19, Value: "አስራ ዘጠኝ"},
		{Number: 18, Value: "አስራ ስምንት"},
		{Number: 17, Value: "አስራ ሰባት"},
		{Number: 16, Value: "አስራ ስድስት"},
		{Number: 15, Value: "አስራ አምስት"},
		{Number: 14, Value: "አስራ አራት"},
		{Number: 13, Value: "አስራ ሶስት"},
		{Number: 12, Value: "አስራ ሁለት"},
		{Number: 11, Value: "አስራ አንድ"},
		{Number: 10, Value: "አስር"},
		{Number: 9, Value: "ዘጠኝ"},
		{Number: 8, Value: "ስምንት"},
		{Number: 7, Value: "ሰባት"},
		{Number: 6, Value: "ስድስት"},
		{Number: 5, Value: "አምስት"},
		{Number: 4, Value: "አራት"},
		{Number: 3, Value: "ሶስት"},
		{Number: 2, Value: "ሁለት"},
		{Number: 1, Value: "አንድ"},
		{Number: 0, Value: "ዜሮ"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "መቶ"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "መጀመሪያ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 2, Word: "ሁለተኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 3, Word: "ሶስተኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 4, Word: "አራተኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 5, Word: "አምስተኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 6, Word: "ስድስተኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 7, Word: "ሰባተኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 8, Word: "ስምንተኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 9, Word: "ዘጠነኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 10, Word: "አስረኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 11, Word: "አስራ አንደኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 12, Word: "አስራ ሁለተኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 13, Word: "አስራ ሶስተኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 14, Word: "አስራ አራተኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 15, Word: "አስራ አምስተኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 16, Word: "አስራ ስድስተኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 17, Word: "አስራ ሰባተኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 18, Word: "አስራ ስምንተኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 19, Word: "አስራ ዘጠነኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 20, Word: "ሃያኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 21, Word: "ሃያ አንደኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 30, Word: "ሰላሳኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 40, Word: "አርባኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 50, Word: "ሃምሳኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 60, Word: "ስስት ዐሥርኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 70, Word: "ሰባኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 80, Word: "ሰማንያኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 90, Word: "ሰባ ዓሥርኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 100, Word: "መቶኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 1000, Word: "ሺህኛ", Suffix: "ኛ", Masculine: "", Feminine: "", Neuter: ""},
	},
	LocaleFormatter: &AmharicFormatter{},
}

// AmharicFormatter handles Amharic (am-ET) formatting
type AmharicFormatter struct{}

func (f *AmharicFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *AmharicFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	// In Amharic, currency name typically comes after the number
	// The currency name is the same for singular and plural (ብር)
	return result + " " + currencyName
}

func (f *AmharicFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	// In Amharic, "and" (እና) connects the whole and fractional parts
	return result + " " + andText + " " + fractionalWords
}

func (f *AmharicFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	// In Amharic, fraction name is the same for singular and plural (ሚልስ)
	return result + " " + fractionName
}

func (f *AmharicFormatter) FormatNegative(result, negativeWord string) string {
	// In Amharic, negative word (አልፎ) comes before the number
	return negativeWord + " " + result
}

func (f *AmharicFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 2
	}
	return amount.Truncate(int32(precision))
}

func (f *AmharicFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}

func (f *AmharicFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma thousands, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
