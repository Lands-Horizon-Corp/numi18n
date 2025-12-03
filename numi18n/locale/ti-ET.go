package locale

import (
	"github.com/shopspring/decimal"
)

// TIETLocale represents the Tigrinya (Ethiopia) locale
var TIETLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Ethiopian Birr",
		Plural:   "ብር",
		Singular: "ብር",
		Symbol:   "ETB",
		FractionUnit: FractionUnit{
			Name:     "Santim",
			Plural:   "ሳንቲም",
			Singular: "ሳንቲም",
			Symbol:   "ሳ",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Ethiopia",
		Currency:       "ETB",
		ISO3166Alpha2:  "ET",
		ISO3166Alpha3:  "ETH",
		ISO3166Numeric: "231",
		Locale:         "ti-ET",
		Timezone:       []string{"Africa/Addis_Ababa"},
		Language:       "ti",
		Emoji:          "🇪🇹",
		PhoneCode:      "+251",
		Domain:         ".et",
	},
	Texts: Texts{
		And:   "ከምኡ'ውን",
		Minus: "ጉሓፍ",
		Only:  "ጥራይ",
		Point: "ነጥቢ",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "ሓደ ትሪልዮን"},
		{Number: 1000000000000, Value: "ሓደ ትሪልዮን"},
		{Number: 1000000000, Value: "ሓደ ቢልዮን"},
		{Number: 1000000, Value: "ሓደ ሚልዮን"},
		{Number: 100000, Value: "ሓደ መቶ ሽሕ"},
		{Number: 90000, Value: "ተስዓተ ሽሕ"},
		{Number: 80000, Value: "ሸሞንተ ሽሕ"},
		{Number: 70000, Value: "ሸውዓተ ሽሕ"},
		{Number: 60000, Value: "ሽድሽተ ሽሕ"},
		{Number: 50000, Value: "ሓምሽተ ሽሕ"},
		{Number: 40000, Value: "ኣርባዕተ ሽሕ"},
		{Number: 30000, Value: "ሰላሳ ሽሕ"},
		{Number: 20000, Value: "ዕስራ ሽሕ"},
		{Number: 19000, Value: "ተስዓተ ዓሰር ሽሕ"},
		{Number: 18000, Value: "ሸሞንተ ዓሰር ሽሕ"},
		{Number: 17000, Value: "ሸውዓተ ዓሰር ሽሕ"},
		{Number: 16000, Value: "ሽድሽተ ዓሰር ሽሕ"},
		{Number: 15000, Value: "ሓምሽተ ዓሰር ሽሕ"},
		{Number: 14000, Value: "ኣርባዕተ ዓሰር ሽሕ"},
		{Number: 13000, Value: "ሰለስተ ዓሰር ሽሕ"},
		{Number: 12000, Value: "ክልተ ዓሰር ሽሕ"},
		{Number: 11000, Value: "ሓዳሽ ዓሰር ሽሕ"},
		{Number: 10000, Value: "ዓሰር ሽሕ"},
		{Number: 9000, Value: "ተስዓተ ሽሕ"},
		{Number: 8000, Value: "ሸሞንተ ሽሕ"},
		{Number: 7000, Value: "ሸውዓተ ሽሕ"},
		{Number: 6000, Value: "ሽድሽተ ሽሕ"},
		{Number: 5000, Value: "ሓምሽተ ሽሕ"},
		{Number: 4000, Value: "ኣርባዕተ ሽሕ"},
		{Number: 3000, Value: "ሰለስተ ሽሕ"},
		{Number: 2000, Value: "ክልተ ሽሕ"},
		{Number: 1000, Value: "ሓደ ሽሕ"},
		{Number: 900, Value: "ተስዓተ መቶ"},
		{Number: 800, Value: "ሸሞንተ መቶ"},
		{Number: 700, Value: "ሸውዓተ መቶ"},
		{Number: 600, Value: "ሽድሽተ መቶ"},
		{Number: 500, Value: "ሓምሽተ መቶ"},
		{Number: 400, Value: "ኣርባዕተ መቶ"},
		{Number: 300, Value: "ሰለስተ መቶ"},
		{Number: 200, Value: "ክልተ መቶ"},
		{Number: 100, Value: "ሓደ መቶ"},
		{Number: 99, Value: "ተስዓተ ሰላሳ ተስዓተ"},
		{Number: 98, Value: "ተስዓተ ሰላሳ ሸሞንተ"},
		{Number: 97, Value: "ተስዓተ ሰላሳ ሸውዓተ"},
		{Number: 96, Value: "ተስዓተ ሰላሳ ሽድሽተ"},
		{Number: 95, Value: "ተስዓተ ሰላሳ ሓምሽተ"},
		{Number: 94, Value: "ተስዓተ ሰላሳ ኣርባዕተ"},
		{Number: 93, Value: "ተስዓተ ሰላሳ ሰለስተ"},
		{Number: 92, Value: "ተስዓተ ሰላሳ ክልተ"},
		{Number: 91, Value: "ተስዓተ ሰላሳ ሓደ"},
		{Number: 90, Value: "ተስዓተ ሰላሳ"},
		{Number: 89, Value: "ሸሞንተ ሰላሳ ተስዓተ"},
		{Number: 88, Value: "ሸሞንተ ሰላሳ ሸሞንተ"},
		{Number: 87, Value: "ሸሞንተ ሰላሳ ሸውዓተ"},
		{Number: 86, Value: "ሸሞንተ ሰላሳ ሽድሽተ"},
		{Number: 85, Value: "ሸሞንተ ሰላሳ ሓምሽተ"},
		{Number: 84, Value: "ሸሞንተ ሰላሳ ኣርባዕተ"},
		{Number: 83, Value: "ሸሞንተ ሰላሳ ሰለስተ"},
		{Number: 82, Value: "ሸሞንተ ሰላሳ ክልተ"},
		{Number: 81, Value: "ሸሞንተ ሰላሳ ሓደ"},
		{Number: 80, Value: "ሸሞንተ ሰላሳ"},
		{Number: 79, Value: "ሸውዓተ ሰላሳ ተስዓተ"},
		{Number: 78, Value: "ሸውዓተ ሰላሳ ሸሞንተ"},
		{Number: 77, Value: "ሸውዓተ ሰላሳ ሸውዓተ"},
		{Number: 76, Value: "ሸውዓተ ሰላሳ ሽድሽተ"},
		{Number: 75, Value: "ሸውዓተ ሰላሳ ሓምሽተ"},
		{Number: 74, Value: "ሸውዓተ ሰላሳ ኣርባዕተ"},
		{Number: 73, Value: "ሸውዓተ ሰላሳ ሰለስተ"},
		{Number: 72, Value: "ሸውዓተ ሰላሳ ክልተ"},
		{Number: 71, Value: "ሸውዓተ ሰላሳ ሓደ"},
		{Number: 70, Value: "ሸውዓተ ሰላሳ"},
		{Number: 69, Value: "ሽድሽተ ሰላሳ ተስዓተ"},
		{Number: 68, Value: "ሽድሽተ ሰላሳ ሸሞንተ"},
		{Number: 67, Value: "ሽድሽተ ሰላሳ ሸውዓተ"},
		{Number: 66, Value: "ሽድሽተ ሰላሳ ሽድሽተ"},
		{Number: 65, Value: "ሽድሽተ ሰላሳ ሓምሽተ"},
		{Number: 64, Value: "ሽድሽተ ሰላሳ ኣርባዕተ"},
		{Number: 63, Value: "ሽድሽተ ሰላሳ ሰለስተ"},
		{Number: 62, Value: "ሽድሽተ ሰላሳ ክልተ"},
		{Number: 61, Value: "ሽድሽተ ሰላሳ ሓደ"},
		{Number: 60, Value: "ሽድሽተ ሰላሳ"},
		{Number: 59, Value: "ሓምሽተ ሰላሳ ተስዓተ"},
		{Number: 58, Value: "ሓምሽተ ሰላሳ ሸሞንተ"},
		{Number: 57, Value: "ሓምሽተ ሰላሳ ሸውዓተ"},
		{Number: 56, Value: "ሓምሽተ ሰላሳ ሽድሽተ"},
		{Number: 55, Value: "ሓምሽተ ሰላሳ ሓምሽተ"},
		{Number: 54, Value: "ሓምሽተ ሰላሳ ኣርባዕተ"},
		{Number: 53, Value: "ሓምሽተ ሰላሳ ሰለስተ"},
		{Number: 52, Value: "ሓምሽተ ሰላሳ ክልተ"},
		{Number: 51, Value: "ሓምሽተ ሰላሳ ሓደ"},
		{Number: 50, Value: "ሓምሽተ ሰላሳ"},
		{Number: 49, Value: "ኣርባዕተ ሰላሳ ተስዓተ"},
		{Number: 48, Value: "ኣርባዕተ ሰላሳ ሸሞንተ"},
		{Number: 47, Value: "ኣርባዕተ ሰላሳ ሸውዓተ"},
		{Number: 46, Value: "ኣርባዕተ ሰላሳ ሽድሽተ"},
		{Number: 45, Value: "ኣርባዕተ ሰላሳ ሓምሽተ"},
		{Number: 44, Value: "ኣርባዕተ ሰላሳ ኣርባዕተ"},
		{Number: 43, Value: "ኣርባዕተ ሰላሳ ሰለስተ"},
		{Number: 42, Value: "ኣርባዕተ ሰላሳ ክልተ"},
		{Number: 41, Value: "ኣርባዕተ ሰላሳ ሓደ"},
		{Number: 40, Value: "ኣርባዕተ ሰላሳ"},
		{Number: 39, Value: "ሰለስተ ሰላሳ ተስዓተ"},
		{Number: 38, Value: "ሰለስተ ሰላሳ ሸሞንተ"},
		{Number: 37, Value: "ሰለስተ ሰላሳ ሸውዓተ"},
		{Number: 36, Value: "ሰለስተ ሰላሳ ሽድሽተ"},
		{Number: 35, Value: "ሰለስተ ሰላሳ ሓምሽተ"},
		{Number: 34, Value: "ሰለስተ ሰላሳ ኣርባዕተ"},
		{Number: 33, Value: "ሰለስተ ሰላሳ ሰለስተ"},
		{Number: 32, Value: "ሰለስተ ሰላሳ ክልተ"},
		{Number: 31, Value: "ሰለስተ ሰላሳ ሓደ"},
		{Number: 30, Value: "ሰለስተ ሰላሳ"},
		{Number: 29, Value: "ክልተ ሰላሳ ተስዓተ"},
		{Number: 28, Value: "ክልተ ሰላሳ ሸሞንተ"},
		{Number: 27, Value: "ክልተ ሰላሳ ሸውዓተ"},
		{Number: 26, Value: "ክልተ ሰላሳ ሽድሽተ"},
		{Number: 25, Value: "ክልተ ሰላሳ ሓምሽተ"},
		{Number: 24, Value: "ክልተ ሰላሳ ኣርባዕተ"},
		{Number: 23, Value: "ክልተ ሰላሳ ሰለስተ"},
		{Number: 22, Value: "ክልተ ሰላሳ ክልተ"},
		{Number: 21, Value: "ክልተ ሰላሳ ሓደ"},
		{Number: 20, Value: "ዕስራ"},
		{Number: 19, Value: "ተስዓተ ዓሰር"},
		{Number: 18, Value: "ሸሞንተ ዓሰር"},
		{Number: 17, Value: "ሸውዓተ ዓሰር"},
		{Number: 16, Value: "ሽድሽተ ዓሰር"},
		{Number: 15, Value: "ሓምሽተ ዓሰር"},
		{Number: 14, Value: "ኣርባዕተ ዓሰር"},
		{Number: 13, Value: "ሰለስተ ዓሰር"},
		{Number: 12, Value: "ክልተ ዓሰር"},
		{Number: 11, Value: "ሓዳሽ ዓሰር"},
		{Number: 10, Value: "ዓሰር"},
		{Number: 9, Value: "ተስዓተ"},
		{Number: 8, Value: "ሸሞንተ"},
		{Number: 7, Value: "ሸውዓተ"},
		{Number: 6, Value: "ሽድሽተ"},
		{Number: 5, Value: "ሓምሽተ"},
		{Number: 4, Value: "ኣርባዕተ"},
		{Number: 3, Value: "ሰለስተ"},
		{Number: 2, Value: "ክልተ"},
		{Number: 1, Value: "ሓደ"},
		{Number: 0, Value: "ዜሮ"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "ሓደ መቶ"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "ቀዳማይ", Suffix: ".", Masculine: "ቀዳማይ", Feminine: "ቀዳመይቲ", Neuter: "ቀዳማይ"},
		{Number: 2, Word: "ካልኣይ", Suffix: ".", Masculine: "ካልኣይ", Feminine: "ካልኣይቲ", Neuter: "ካልኣይ"},
		{Number: 3, Word: "ሳልሳይ", Suffix: ".", Masculine: "ሳልሳይ", Feminine: "ሳልሳይቲ", Neuter: "ሳልሳይ"},
		{Number: 4, Word: "ራብዓይ", Suffix: ".", Masculine: "ራብዓይ", Feminine: "ራብዓይቲ", Neuter: "ራብዓይ"},
		{Number: 5, Word: "ሓሙሻይ", Suffix: ".", Masculine: "ሓሙሻይ", Feminine: "ሓሙሻይቲ", Neuter: "ሓሙሻይ"},
		{Number: 6, Word: "ሻዱሻይ", Suffix: ".", Masculine: "ሻዱሻይ", Feminine: "ሻዱሻይቲ", Neuter: "ሻዱሻይ"},
		{Number: 7, Word: "ሸውዓተኛ", Suffix: ".", Masculine: "ሸውዓተኛ", Feminine: "ሸውዓተኛይቲ", Neuter: "ሸውዓተኛ"},
		{Number: 8, Word: "ሸሞንተኛ", Suffix: ".", Masculine: "ሸሞንተኛ", Feminine: "ሸሞንተኛይቲ", Neuter: "ሸሞንተኛ"},
		{Number: 9, Word: "ተስዓተኛ", Suffix: ".", Masculine: "ተስዓተኛ", Feminine: "ተስዓተኛይቲ", Neuter: "ተስዓተኛ"},
		{Number: 10, Word: "ዓስራይ", Suffix: ".", Masculine: "ዓስራይ", Feminine: "ዓስራይቲ", Neuter: "ዓስራይ"},
		{Number: 11, Word: "ዓስራ ሓደኛ", Suffix: ".", Masculine: "ዓስራ ሓደኛ", Feminine: "ዓስራ ሓደኛይቲ", Neuter: "ዓስራ ሓደኛ"},
		{Number: 12, Word: "ዓስራ ክልተኛ", Suffix: ".", Masculine: "ዓስራ ክልተኛ", Feminine: "ዓስራ ክልተኛይቲ", Neuter: "ዓስራ ክልተኛ"},
		{Number: 20, Word: "ዕስራይ", Suffix: ".", Masculine: "ዕስራይ", Feminine: "ዕስራይቲ", Neuter: "ዕስራይ"},
		{Number: 21, Word: "ዕስራ ሓደኛ", Suffix: ".", Masculine: "ዕስራ ሓደኛ", Feminine: "ዕስራ ሓደኛይቲ", Neuter: "ዕስራ ሓደኛ"},
		{Number: 30, Word: "ሰላሳይ", Suffix: ".", Masculine: "ሰላሳይ", Feminine: "ሰላሳይቲ", Neuter: "ሰላሳይ"},
		{Number: 50, Word: "ሓምሳይ", Suffix: ".", Masculine: "ሓምሳይ", Feminine: "ሓምሳይቲ", Neuter: "ሓምሳይ"},
		{Number: 100, Word: "መቶይ", Suffix: ".", Masculine: "መቶይ", Feminine: "መቶይቲ", Neuter: "መቶይ"},
		{Number: 1000, Word: "ሽሕይ", Suffix: ".", Masculine: "ሽሕይ", Feminine: "ሽሕይቲ", Neuter: "ሽሕይ"},
	},
	LocaleFormatter: &TigrayEthiopiaFormatter{},
}

// TigrayEthiopiaFormatter handles Tigrinya (Ethiopia) formatting
type TigrayEthiopiaFormatter struct{}

func (f *TigrayEthiopiaFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *TigrayEthiopiaFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	return result + " " + currencyName
}

func (f *TigrayEthiopiaFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *TigrayEthiopiaFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	return result + " " + fractionName
}

func (f *TigrayEthiopiaFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *TigrayEthiopiaFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 2
	}
	return amount.Truncate(int32(precision))
}

func (f *TigrayEthiopiaFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *TigrayEthiopiaFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
