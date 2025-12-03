package locale

import (
	"github.com/shopspring/decimal"
)

// PAINLocale represents the Punjabi (India) locale
var PAINLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "ਭਾਰਤੀ ਰੁਪਇਆ",
		Plural:   "ਭਾਰਤੀ ਰੁਪਏ",
		Singular: "ਭਾਰਤੀ ਰੁਪਇਆ",
		Symbol:   "₹",
		FractionUnit: FractionUnit{
			Name:     "ਪੈਸਾ",
			Plural:   "ਪੈਸੇ",
			Singular: "ਪੈਸਾ",
			Symbol:   "p",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "India",
		Currency:       "INR",
		ISO3166Alpha2:  "IN",
		ISO3166Alpha3:  "IND",
		ISO3166Numeric: "356",
		Locale:         "pa-IN",
		Timezone:       []string{"Asia/Kolkata"},
		Language:       "pa",
		Emoji:          "🇮🇳",
		PhoneCode:      "+91",
		Domain:         ".in",
	},
	Texts: Texts{
		And:   "ਅਤੇ",
		Minus: "ਮਾਇਨਸ",
		Only:  "ਸਿਰਫ਼",
		Point: "ਦਸ਼ਮਲਵ",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 10000000000000, Value: "ਦਸ ਖਰਬ"},
		{Number: 1000000000000, Value: "ਇੱਕ ਖਰਬ"},
		{Number: 100000000000, Value: "ਦਸ ਨੀਲ"},
		{Number: 10000000000, Value: "ਇੱਕ ਨੀਲ"},
		{Number: 1000000000, Value: "ਇੱਕ ਅਰਬ"},
		{Number: 100000000, Value: "ਦਸ ਕਰੋੜ"},
		{Number: 10000000, Value: "ਇੱਕ ਕਰੋੜ"},
		{Number: 1000000, Value: "ਦਸ ਲੱਖ"},
		{Number: 100000, Value: "ਇੱਕ ਲੱਖ"},
		{Number: 90000, Value: "ਨੱਬੇ ਹਜ਼ਾਰ"},
		{Number: 80000, Value: "ਅੱਸੀ ਹਜ਼ਾਰ"},
		{Number: 70000, Value: "ਸੱਤਰ ਹਜ਼ਾਰ"},
		{Number: 60000, Value: "ਸੱਠ ਹਜ਼ਾਰ"},
		{Number: 50000, Value: "ਪੰਜਾਹ ਹਜ਼ਾਰ"},
		{Number: 40000, Value: "ਚਾਲੀਸ ਹਜ਼ਾਰ"},
		{Number: 30000, Value: "ਤੀਹ ਹਜ਼ਾਰ"},
		{Number: 20000, Value: "ਵੀਹ ਹਜ਼ਾਰ"},
		{Number: 19000, Value: "ਉਨੀ ਹਜ਼ਾਰ"},
		{Number: 18000, Value: "ਅਠਾਰਾਂ ਹਜ਼ਾਰ"},
		{Number: 17000, Value: "ਸਤਾਰਾਂ ਹਜ਼ਾਰ"},
		{Number: 16000, Value: "ਸੋਲਾਂ ਹਜ਼ਾਰ"},
		{Number: 15000, Value: "ਪੰਦਰਾਂ ਹਜ਼ਾਰ"},
		{Number: 14000, Value: "ਚੌਦਾਂ ਹਜ਼ਾਰ"},
		{Number: 13000, Value: "ਤੇਰਾਂ ਹਜ਼ਾਰ"},
		{Number: 12000, Value: "ਬਾਰਾਂ ਹਜ਼ਾਰ"},
		{Number: 11000, Value: "ਗਿਆਰਾਂ ਹਜ਼ਾਰ"},
		{Number: 10000, Value: "ਦਸ ਹਜ਼ਾਰ"},
		{Number: 9000, Value: "ਨੌ ਹਜ਼ਾਰ"},
		{Number: 8000, Value: "ਅੱਠ ਹਜ਼ਾਰ"},
		{Number: 7000, Value: "ਸੱਤ ਹਜ਼ਾਰ"},
		{Number: 6000, Value: "ਛੇ ਹਜ਼ਾਰ"},
		{Number: 5000, Value: "ਪੰਜ ਹਜ਼ਾਰ"},
		{Number: 4000, Value: "ਚਾਰ ਹਜ਼ਾਰ"},
		{Number: 3000, Value: "ਤਿੰਨ ਹਜ਼ਾਰ"},
		{Number: 2000, Value: "ਦੋ ਹਜ਼ਾਰ"},
		{Number: 1000, Value: "ਇੱਕ ਹਜ਼ਾਰ"},
		{Number: 900, Value: "ਨੌ ਸੌ"},
		{Number: 800, Value: "ਅੱਠ ਸੌ"},
		{Number: 700, Value: "ਸੱਤ ਸੌ"},
		{Number: 600, Value: "ਛੇ ਸੌ"},
		{Number: 500, Value: "ਪੰਜ ਸੌ"},
		{Number: 400, Value: "ਚਾਰ ਸੌ"},
		{Number: 300, Value: "ਤਿੰਨ ਸੌ"},
		{Number: 200, Value: "ਦੋ ਸੌ"},
		{Number: 100, Value: "ਇੱਕ ਸੌ"},
		{Number: 99, Value: "ਨਿੰਨਾਣਵੇਂ"},
		{Number: 98, Value: "ਅਠਿਆਣਵੇਂ"},
		{Number: 97, Value: "ਸਤਿਆਣਵੇਂ"},
		{Number: 96, Value: "ਛਿਆਣਵੇਂ"},
		{Number: 95, Value: "ਪਚਿਆਣਵੇਂ"},
		{Number: 94, Value: "ਚੌਰਾਣਵੇਂ"},
		{Number: 93, Value: "ਤਿਰਾਣਵੇਂ"},
		{Number: 92, Value: "ਬਾਣਵੇਂ"},
		{Number: 91, Value: "ਇਕਾਣਵੇਂ"},
		{Number: 90, Value: "ਨੱਬੇ"},
		{Number: 89, Value: "ਨਵਾਸੀ"},
		{Number: 88, Value: "ਅਠਿਆਸੀ"},
		{Number: 87, Value: "ਸਤਿਆਸੀ"},
		{Number: 86, Value: "ਛਿਆਸੀ"},
		{Number: 85, Value: "ਪਚਿਆਸੀ"},
		{Number: 84, Value: "ਚੌਰਾਸੀ"},
		{Number: 83, Value: "ਤਿਰਾਸੀ"},
		{Number: 82, Value: "ਬਿਆਸੀ"},
		{Number: 81, Value: "ਇਕਿਆਸੀ"},
		{Number: 80, Value: "ਅੱਸੀ"},
		{Number: 79, Value: "ਉਨਾਸੀ"},
		{Number: 78, Value: "ਅਠਿਆਤਰ"},
		{Number: 77, Value: "ਸਤਿਆਤਰ"},
		{Number: 76, Value: "ਛਿਹੱਤਰ"},
		{Number: 75, Value: "ਪਚਿਹੱਤਰ"},
		{Number: 74, Value: "ਚੌਹੱਤਰ"},
		{Number: 73, Value: "ਤਿਹੱਤਰ"},
		{Number: 72, Value: "ਬਹੱਤਰ"},
		{Number: 71, Value: "ਇਕਹੱਤਰ"},
		{Number: 70, Value: "ਸੱਤਰ"},
		{Number: 69, Value: "ਉਨਹੱਤਰ"},
		{Number: 68, Value: "ਅਠਠਿਸਠ"},
		{Number: 67, Value: "ਸਤਿਸਠ"},
		{Number: 66, Value: "ਛਿਆਸਠ"},
		{Number: 65, Value: "ਪਤਿਸਠ"},
		{Number: 64, Value: "ਚੌਂਸਠ"},
		{Number: 63, Value: "ਤਿਰਸਠ"},
		{Number: 62, Value: "ਬਾਹਸਠ"},
		{Number: 61, Value: "ਇਕਸਠ"},
		{Number: 60, Value: "ਸੱਠ"},
		{Number: 59, Value: "ਉਨਸਠ"},
		{Number: 58, Value: "ਅਠਵੰਜਾ"},
		{Number: 57, Value: "ਸਤਵੰਜਾ"},
		{Number: 56, Value: "ਛਪਨ"},
		{Number: 55, Value: "ਪਚਵੰਜਾ"},
		{Number: 54, Value: "ਚੁਰੰਜਾ"},
		{Number: 53, Value: "ਤਰਵੰਜਾ"},
		{Number: 52, Value: "ਬਵੰਜਾ"},
		{Number: 51, Value: "ਇਕਵੰਜਾ"},
		{Number: 50, Value: "ਪੰਜਾਹ"},
		{Number: 49, Value: "ਉਨੰਜਾ"},
		{Number: 48, Value: "ਅਠਤਾਲੀ"},
		{Number: 47, Value: "ਸਤਤਾਲੀ"},
		{Number: 46, Value: "ਛਿਤਾਲੀ"},
		{Number: 45, Value: "ਪੰਤਾਲੀ"},
		{Number: 44, Value: "ਚੁੰਤਾਲੀ"},
		{Number: 43, Value: "ਤਰਤਾਲੀ"},
		{Number: 42, Value: "ਬਿਤਾਲੀ"},
		{Number: 41, Value: "ਇਕਤਾਲੀ"},
		{Number: 40, Value: "ਚਾਲੀਸ"},
		{Number: 39, Value: "ਉਨਤਾਲੀ"},
		{Number: 38, Value: "ਅਠੱਤੀਸ"},
		{Number: 37, Value: "ਸੈਂਤੀਸ"},
		{Number: 36, Value: "ਛੱਤੀਸ"},
		{Number: 35, Value: "ਪੈਂਤੀਸ"},
		{Number: 34, Value: "ਚੌਂਤੀਸ"},
		{Number: 33, Value: "ਤੇਂਤੀਸ"},
		{Number: 32, Value: "ਬੱਤੀਸ"},
		{Number: 31, Value: "ਇਕੱਤੀਸ"},
		{Number: 30, Value: "ਤੀਹ"},
		{Number: 29, Value: "ਉਨੱਤੀਸ"},
		{Number: 28, Value: "ਅੱਠਾਈ"},
		{Number: 27, Value: "ਸਤਾਈ"},
		{Number: 26, Value: "ਛੱਬੀਸ"},
		{Number: 25, Value: "ਪੱਚੀਸ"},
		{Number: 24, Value: "ਚੌਂਵੀਸ"},
		{Number: 23, Value: "ਤੇਈਸ"},
		{Number: 22, Value: "ਬਾਈਸ"},
		{Number: 21, Value: "ਇੱਕੀਸ"},
		{Number: 20, Value: "ਵੀਹ"},
		{Number: 19, Value: "ਉਨੀ"},
		{Number: 18, Value: "ਅਠਾਰਾਂ"},
		{Number: 17, Value: "ਸਤਾਰਾਂ"},
		{Number: 16, Value: "ਸੋਲਾਂ"},
		{Number: 15, Value: "ਪੰਦਰਾਂ"},
		{Number: 14, Value: "ਚੌਦਾਂ"},
		{Number: 13, Value: "ਤੇਰਾਂ"},
		{Number: 12, Value: "ਬਾਰਾਂ"},
		{Number: 11, Value: "ਗਿਆਰਾਂ"},
		{Number: 10, Value: "ਦਸ"},
		{Number: 9, Value: "ਨੌ"},
		{Number: 8, Value: "ਅੱਠ"},
		{Number: 7, Value: "ਸੱਤ"},
		{Number: 6, Value: "ਛੇ"},
		{Number: 5, Value: "ਪੰਜ"},
		{Number: 4, Value: "ਚਾਰ"},
		{Number: 3, Value: "ਤਿੰਨ"},
		{Number: 2, Value: "ਦੋ"},
		{Number: 1, Value: "ਇੱਕ"},
		{Number: 0, Value: "ਸਿਫ਼ਰ"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "ਇੱਕ ਸੌ"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "ਪਹਿਲਾ", Suffix: "-ਵਾਂ", Masculine: "ਪਹਿਲਾ", Feminine: "ਪਹਿਲੀ", Neuter: "ਪਹਿਲਾ"},
		{Number: 2, Word: "ਦੂਜਾ", Suffix: "-ਵਾਂ", Masculine: "ਦੂਜਾ", Feminine: "ਦੂਜੀ", Neuter: "ਦੂਜਾ"},
		{Number: 3, Word: "ਤੀਜਾ", Suffix: "-ਵਾਂ", Masculine: "ਤੀਜਾ", Feminine: "ਤੀਜੀ", Neuter: "ਤੀਜਾ"},
		{Number: 4, Word: "ਚੌਥਾ", Suffix: "-ਵਾਂ", Masculine: "ਚੌਥਾ", Feminine: "ਚੌਥੀ", Neuter: "ਚੌਥਾ"},
		{Number: 5, Word: "ਪੰਜਵਾਂ", Suffix: "-ਵਾਂ", Masculine: "ਪੰਜਵਾਂ", Feminine: "ਪੰਜਵੀਂ", Neuter: "ਪੰਜਵਾਂ"},
		{Number: 6, Word: "ਛੇਵਾਂ", Suffix: "-ਵਾਂ", Masculine: "ਛੇਵਾਂ", Feminine: "ਛੇਵੀਂ", Neuter: "ਛੇਵਾਂ"},
		{Number: 7, Word: "ਸੱਤਵਾਂ", Suffix: "-ਵਾਂ", Masculine: "ਸੱਤਵਾਂ", Feminine: "ਸੱਤਵੀਂ", Neuter: "ਸੱਤਵਾਂ"},
		{Number: 8, Word: "ਅੱਠਵਾਂ", Suffix: "-ਵਾਂ", Masculine: "ਅੱਠਵਾਂ", Feminine: "ਅੱਠਵੀਂ", Neuter: "ਅੱਠਵਾਂ"},
		{Number: 9, Word: "ਨੌਵਾਂ", Suffix: "-ਵਾਂ", Masculine: "ਨੌਵਾਂ", Feminine: "ਨੌਵੀਂ", Neuter: "ਨੌਵਾਂ"},
		{Number: 10, Word: "ਦਸਵਾਂ", Suffix: "-ਵਾਂ", Masculine: "ਦਸਵਾਂ", Feminine: "ਦਸਵੀਂ", Neuter: "ਦਸਵਾਂ"},
		{Number: 11, Word: "ਗਿਆਰਵਾਂ", Suffix: "-ਵਾਂ", Masculine: "ਗਿਆਰਵਾਂ", Feminine: "ਗਿਆਰਵੀਂ", Neuter: "ਗਿਆਰਵਾਂ"},
		{Number: 12, Word: "ਬਾਰਵਾਂ", Suffix: "-ਵਾਂ", Masculine: "ਬਾਰਵਾਂ", Feminine: "ਬਾਰਵੀਂ", Neuter: "ਬਾਰਵਾਂ"},
		{Number: 20, Word: "ਵੀਹਵਾਂ", Suffix: "-ਵਾਂ", Masculine: "ਵੀਹਵਾਂ", Feminine: "ਵੀਹਵੀਂ", Neuter: "ਵੀਹਵਾਂ"},
		{Number: 21, Word: "ਇੱਕੀਵਾਂ", Suffix: "-ਵਾਂ", Masculine: "ਇੱਕੀਵਾਂ", Feminine: "ਇੱਕੀਵੀਂ", Neuter: "ਇੱਕੀਵਾਂ"},
		{Number: 30, Word: "ਤੀਹਵਾਂ", Suffix: "-ਵਾਂ", Masculine: "ਤੀਹਵਾਂ", Feminine: "ਤੀਹਵੀਂ", Neuter: "ਤੀਹਵਾਂ"},
		{Number: 50, Word: "ਪੰਜਾਹਵਾਂ", Suffix: "-ਵਾਂ", Masculine: "ਪੰਜਾਹਵਾਂ", Feminine: "ਪੰਜਾਹਵੀਂ", Neuter: "ਪੰਜਾਹਵਾਂ"},
		{Number: 100, Word: "ਸੌਵਾਂ", Suffix: "-ਵਾਂ", Masculine: "ਸੌਵਾਂ", Feminine: "ਸੌਵੀਂ", Neuter: "ਸੌਵਾਂ"},
		{Number: 1000, Word: "ਹਜ਼ਾਰਵਾਂ", Suffix: "-ਵਾਂ", Masculine: "ਹਜ਼ਾਰਵਾਂ", Feminine: "ਹਜ਼ਾਰਵੀਂ", Neuter: "ਹਜ਼ਾਰਵਾਂ"},
	},
	LocaleFormatter: &PunjabiFormatter{},
}

// PunjabiFormatter handles Punjabi (India) formatting
type PunjabiFormatter struct{}

func (f *PunjabiFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *PunjabiFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *PunjabiFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *PunjabiFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *PunjabiFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *PunjabiFormatter) ChopDecimal(decimal decimal.Decimal, precision int) decimal.Decimal {
	return decimal.Truncate(int32(precision))
}

func (f *PunjabiFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *PunjabiFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
