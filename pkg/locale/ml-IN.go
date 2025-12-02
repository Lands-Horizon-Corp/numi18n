package locale

import (
	"strings"
	"github.com/shopspring/decimal"
)

// MLINLocale represents the Malayalam (India) locale
var MLINLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "രൂപ",
		Plural:   "രൂപകൾ",
		Singular: "രൂപ",
		Symbol:   "₹",
		FractionUnit: FractionUnit{
			Name:     "പൈസ",
			Plural:   "പൈസകൾ",
			Singular: "പൈസ",
			Symbol:   "p",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "India",
		Currency:       "INR",
		ISO3166Alpha2:  "IN",
		ISO3166Alpha3:  "IND",
		ISO3166Numeric: "356",
		Locale:         "ml-IN",
		Timezone:       []string{"Asia/Kolkata"},
		Language:       "ml",
	},
	Texts: Texts{
		And:   "ഉം",
		Minus: "കുറച്ച്",
		Only:  "മാത്രം",
		Point: "ബിന്ദു",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 10000000000000000, Value: "പത്ത് ശതകോടി കോടി"},
		{Number: 1000000000000000, Value: "ശതകോടി കോടി"},
		{Number: 100000000000000, Value: "പത്ത് ലക്ഷം കോടി"},
		{Number: 10000000000000, Value: "ലക്ഷം കോടി"},
		{Number: 1000000000000, Value: "പത്ത് ആയിരം കോടി"},
		{Number: 100000000000, Value: "ആയിരം കോടി"},
		{Number: 10000000000, Value: "നൂറ് കോടി"},
		{Number: 1000000000, Value: "പത്ത് കോടി"},
		{Number: 100000000, Value: "കോടി"},
		{Number: 10000000, Value: "പത്ത് ലക്ഷം"},
		{Number: 1000000, Value: "ലക്ഷം"},
		{Number: 100000, Value: "ലക്ഷം"},
		{Number: 90000, Value: "തൊണ്ണൂറ് ആയിരം"},
		{Number: 80000, Value: "എൺപത് ആയിരം"},
		{Number: 70000, Value: "എഴുപത് ആയിരം"},
		{Number: 60000, Value: "അറുപത് ആയിരം"},
		{Number: 50000, Value: "അൻപത് ആയിരം"},
		{Number: 40000, Value: "നാൽപ്പത് ആയിരം"},
		{Number: 30000, Value: "മുപ്പത് ആയിരം"},
		{Number: 20000, Value: "ഇരുപത് ആയിരം"},
		{Number: 19000, Value: "പത്തൊമ്പത് ആയിരം"},
		{Number: 18000, Value: "പതിനെട്ട് ആയിരം"},
		{Number: 17000, Value: "പതിനേഴ് ആയിരം"},
		{Number: 16000, Value: "പതിനാറ് ആയിരം"},
		{Number: 15000, Value: "പതിനഞ്ച് ആയിരം"},
		{Number: 14000, Value: "പതിനാല് ആയിരം"},
		{Number: 13000, Value: "പതിമൂന്ന് ആയിരം"},
		{Number: 12000, Value: "പന്ത്രണ്ട് ആയിരം"},
		{Number: 11000, Value: "പതിനൊന്ന് ആയിരം"},
		{Number: 10000, Value: "പത്ത് ആയിരം"},
		{Number: 9000, Value: "ഒമ്പത് ആയിരം"},
		{Number: 8000, Value: "എട്ട് ആയിരം"},
		{Number: 7000, Value: "ഏഴ് ആയിരം"},
		{Number: 6000, Value: "ആറ് ആയിരം"},
		{Number: 5000, Value: "അഞ്ച് ആയിരം"},
		{Number: 4000, Value: "നാല് ആയിരം"},
		{Number: 3000, Value: "മൂന്ന് ആയിരം"},
		{Number: 2000, Value: "രണ്ട് ആയിരം"},
		{Number: 1000, Value: "ആയിരം"},
		{Number: 900, Value: "തൊള്ളായിരം"},
		{Number: 800, Value: "എണ്ണൂറ്"},
		{Number: 700, Value: "എഴുനൂറ്"},
		{Number: 600, Value: "അറുനൂറ്"},
		{Number: 500, Value: "അഞ്ഞൂറ്"},
		{Number: 400, Value: "നാനൂറ്"},
		{Number: 300, Value: "മുന്നൂറ്"},
		{Number: 200, Value: "ഇരുനൂറ്"},
		{Number: 100, Value: "നൂറ്"},
		{Number: 99, Value: "തൊണ്ണൂറ്റൊമ്പത്"},
		{Number: 98, Value: "തൊണ്ണൂറ്റെട്ട്"},
		{Number: 97, Value: "തൊണ്ണൂറ്റേഴ്"},
		{Number: 96, Value: "തൊണ്ണൂറ്റാറ്"},
		{Number: 95, Value: "തൊണ്ണൂറ്റഞ്ച്"},
		{Number: 94, Value: "തൊണ്ണൂറ്റാല്"},
		{Number: 93, Value: "തൊണ്ണൂറ്റമൂന്ന്"},
		{Number: 92, Value: "തൊണ്ണൂറ്റരണ്ട്"},
		{Number: 91, Value: "തൊണ്ണൂറ്റൊന്ന്"},
		{Number: 90, Value: "തൊണ്ണൂറ്"},
		{Number: 89, Value: "എൺപത്തൊമ്പത്"},
		{Number: 88, Value: "എൺപത്തെട്ട്"},
		{Number: 87, Value: "എൺപത്തേഴ്"},
		{Number: 86, Value: "എൺപത്താറ്"},
		{Number: 85, Value: "എൺപത്തഞ്ച്"},
		{Number: 84, Value: "എൺപത്തനാല്"},
		{Number: 83, Value: "എൺപത്തമൂന്ന്"},
		{Number: 82, Value: "എൺപത്തരണ്ട്"},
		{Number: 81, Value: "എൺപത്തൊന്ന്"},
		{Number: 80, Value: "എൺപത്"},
		{Number: 79, Value: "എഴുപത്തൊമ്പത്"},
		{Number: 78, Value: "എഴുപത്തെട്ട്"},
		{Number: 77, Value: "എഴുപത്തേഴ്"},
		{Number: 76, Value: "എഴുപത്താറ്"},
		{Number: 75, Value: "എഴുപത്തഞ്ച്"},
		{Number: 74, Value: "എഴുപത്തനാല്"},
		{Number: 73, Value: "എഴുപത്തമൂന്ന്"},
		{Number: 72, Value: "എഴുപത്തരണ്ട്"},
		{Number: 71, Value: "എഴുപത്തൊന്ന്"},
		{Number: 70, Value: "എഴുപത്"},
		{Number: 69, Value: "അറുപത്തൊമ്പത്"},
		{Number: 68, Value: "അറുപത്തെട്ട്"},
		{Number: 67, Value: "അറുപത്തേഴ്"},
		{Number: 66, Value: "അറുപത്താറ്"},
		{Number: 65, Value: "അറുപത്തഞ്ച്"},
		{Number: 64, Value: "അറുപത്തനാല്"},
		{Number: 63, Value: "അറുപത്തമൂന്ന്"},
		{Number: 62, Value: "അറുപത്തരണ്ട്"},
		{Number: 61, Value: "അറുപത്തൊന്ന്"},
		{Number: 60, Value: "അറുപത്"},
		{Number: 59, Value: "അൻപത്തൊമ്പത്"},
		{Number: 58, Value: "അൻപത്തെട്ട്"},
		{Number: 57, Value: "അൻപത്തേഴ്"},
		{Number: 56, Value: "അൻപത്താറ്"},
		{Number: 55, Value: "അൻപത്തഞ്ച്"},
		{Number: 54, Value: "അൻപത്തനാല്"},
		{Number: 53, Value: "അൻപത്തമൂന്ന്"},
		{Number: 52, Value: "അൻപത്തരണ്ട്"},
		{Number: 51, Value: "അൻപത്തൊന്ന്"},
		{Number: 50, Value: "അൻപത്"},
		{Number: 49, Value: "നാൽപ്പത്തൊമ്പത്"},
		{Number: 48, Value: "നാൽപ്പത്തെട്ട്"},
		{Number: 47, Value: "നാൽപ്പത്തേഴ്"},
		{Number: 46, Value: "നാൽപ്പത്താറ്"},
		{Number: 45, Value: "നാൽപ്പത്തഞ്ച്"},
		{Number: 44, Value: "നാൽപ്പത്തനാല്"},
		{Number: 43, Value: "നാൽപ്പത്തമൂന്ന്"},
		{Number: 42, Value: "നാൽപ്പത്തരണ്ട്"},
		{Number: 41, Value: "നാൽപ്പത്തൊന്ന്"},
		{Number: 40, Value: "നാൽപ്പത്"},
		{Number: 39, Value: "മുപ്പത്തൊമ്പത്"},
		{Number: 38, Value: "മുപ്പത്തെട്ട്"},
		{Number: 37, Value: "മുപ്പത്തേഴ്"},
		{Number: 36, Value: "മുപ്പത്താറ്"},
		{Number: 35, Value: "മുപ്പത്തഞ്ച്"},
		{Number: 34, Value: "മുപ്പത്തനാല്"},
		{Number: 33, Value: "മുപ്പത്തമൂന്ന്"},
		{Number: 32, Value: "മുപ്പത്തരണ്ട്"},
		{Number: 31, Value: "മുപ്പത്തൊന്ന്"},
		{Number: 30, Value: "മുപ്പത്"},
		{Number: 29, Value: "ഇരുപത്തൊമ്പത്"},
		{Number: 28, Value: "ഇരുപത്തെട്ട്"},
		{Number: 27, Value: "ഇരുപത്തേഴ്"},
		{Number: 26, Value: "ഇരുപത്താറ്"},
		{Number: 25, Value: "ഇരുപത്തഞ്ച്"},
		{Number: 24, Value: "ഇരുപത്തനാല്"},
		{Number: 23, Value: "ഇരുപത്തമൂന്ന്"},
		{Number: 22, Value: "ഇരുപത്തരണ്ട്"},
		{Number: 21, Value: "ഇരുപത്തൊന്ന്"},
		{Number: 20, Value: "ഇരുപത്"},
		{Number: 19, Value: "പത്തൊമ്പത്"},
		{Number: 18, Value: "പതിനെട്ട്"},
		{Number: 17, Value: "പതിനേഴ്"},
		{Number: 16, Value: "പതിനാറ്"},
		{Number: 15, Value: "പതിനഞ്ച്"},
		{Number: 14, Value: "പതിനാല്"},
		{Number: 13, Value: "പതിമൂന്ന്"},
		{Number: 12, Value: "പന്ത്രണ്ട്"},
		{Number: 11, Value: "പതിനൊന്ന്"},
		{Number: 10, Value: "പത്ത്"},
		{Number: 9, Value: "ഒമ്പത്"},
		{Number: 8, Value: "എട്ട്"},
		{Number: 7, Value: "ഏഴ്"},
		{Number: 6, Value: "ആറ്"},
		{Number: 5, Value: "അഞ്ച്"},
		{Number: 4, Value: "നാല്"},
		{Number: 3, Value: "മൂന്ന്"},
		{Number: 2, Value: "രണ്ട്"},
		{Number: 1, Value: "ഒന്ന്"},
		{Number: 0, Value: "പൂജ്യം"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "ഒരു നൂറ്"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "ഒന്നാം", Suffix: "-ാം", Masculine: "ഒന്നാം", Feminine: "ഒന്നാം", Neuter: "ഒന്നാം"},
		{Number: 2, Word: "രണ്ടാം", Suffix: "-ാം", Masculine: "രണ്ടാം", Feminine: "രണ്ടാം", Neuter: "രണ്ടാം"},
		{Number: 3, Word: "മൂന്നാം", Suffix: "-ാം", Masculine: "മൂന്നാം", Feminine: "മൂന്നാം", Neuter: "മൂന്നാം"},
		{Number: 4, Word: "നാലാം", Suffix: "-ാം", Masculine: "നാലാം", Feminine: "നാലാം", Neuter: "നാലാം"},
		{Number: 5, Word: "അഞ്ചാം", Suffix: "-ാം", Masculine: "അഞ്ചാം", Feminine: "അഞ്ചാം", Neuter: "അഞ്ചാം"},
		{Number: 6, Word: "ആറാം", Suffix: "-ാം", Masculine: "ആറാം", Feminine: "ആറാം", Neuter: "ആറാം"},
		{Number: 7, Word: "ഏഴാം", Suffix: "-ാം", Masculine: "ഏഴാം", Feminine: "ഏഴാം", Neuter: "ഏഴാം"},
		{Number: 8, Word: "എട്ടാം", Suffix: "-ാം", Masculine: "എട്ടാം", Feminine: "എട്ടാം", Neuter: "എട്ടാം"},
		{Number: 9, Word: "ഒമ്പതാം", Suffix: "-ാം", Masculine: "ഒമ്പതാം", Feminine: "ഒമ്പതാം", Neuter: "ഒമ്പതാം"},
		{Number: 10, Word: "പത്താം", Suffix: "-ാം", Masculine: "പത്താം", Feminine: "പത്താം", Neuter: "പത്താം"},
		{Number: 11, Word: "പതിനൊന്നാം", Suffix: "-ാം", Masculine: "പതിനൊന്നാം", Feminine: "പതിനൊന്നാം", Neuter: "പതിനൊന്നാം"},
		{Number: 12, Word: "പന്ത്രണ്ടാം", Suffix: "-ാം", Masculine: "പന്ത്രണ്ടാം", Feminine: "പന്ത്രണ്ടാം", Neuter: "പന്ത്രണ്ടാം"},
		{Number: 20, Word: "ഇരുപതാം", Suffix: "-ാം", Masculine: "ഇരുപതാം", Feminine: "ഇരുപതാം", Neuter: "ഇരുപതാം"},
		{Number: 21, Word: "ഇരുപത്തൊന്നാം", Suffix: "-ാം", Masculine: "ഇരുപത്തൊന്നാം", Feminine: "ഇരുപത്തൊന്നാം", Neuter: "ഇരുപത്തൊന്നാം"},
		{Number: 30, Word: "മുപ്പതാം", Suffix: "-ാം", Masculine: "മുപ്പതാം", Feminine: "മുപ്പതാം", Neuter: "മുപ്പതാം"},
		{Number: 100, Word: "നൂറാം", Suffix: "-ാം", Masculine: "നൂറാം", Feminine: "നൂറാം", Neuter: "നൂറാം"},
		{Number: 1000, Word: "ആയിരാം", Suffix: "-ാം", Masculine: "ആയിരാം", Feminine: "ആയിരാം", Neuter: "ആയിരാം"},
	},
	LocaleFormatter: &MalayalamFormatter{},
}

// MalayalamFormatter handles Malayalam formatting
type MalayalamFormatter struct{}

func (f *MalayalamFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *MalayalamFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *MalayalamFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *MalayalamFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *MalayalamFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *MalayalamFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 0
	}
	return amount.Truncate(int32(precision))
}


func (f *MalayalamFormatter) FormatDecimalNumber(amount float64) string {
	return DefaultFormatDecimalNumber(amount, ",", ".")
}
func (f *MalayalamFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
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
