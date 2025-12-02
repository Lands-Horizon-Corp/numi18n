package locale

import (
	"strings"
	"github.com/shopspring/decimal"
)

// ILLocale is a NumI18NLocale configured for Israel (he-IL)
var ILLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "שקל",
		Plural:   "שקלים",
		Singular: "שקל",
		Symbol:   "₪",
		FractionUnit: FractionUnit{
			Name:     "אגורה",
			Plural:   "אגורות",
			Singular: "אגורה",
			Symbol:   "ע׳",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Israel",
		Currency:       "ILS",
		ISO3166Alpha2:  "IL",
		ISO3166Alpha3:  "ISR",
		ISO3166Numeric: "376",
		Locale:         "he-IL",
		Timezone:       []string{"Asia/Jerusalem"},
		Language:       "he",
	},
	Texts: Texts{
		And:   "ו",
		Minus: "מינוס",
		Only:  "בלבד",
		Point: "נקודה",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "אחד קוודריליון"},
		{Number: 1000000000000, Value: "אחד טריליון"},
		{Number: 1000000000, Value: "אחד מיליארד"},
		{Number: 1000000, Value: "אחד מיליון"},
		{Number: 1000, Value: "אלף"},
		{Number: 100, Value: "מאה"},
		{Number: 90, Value: "תשעים"},
		{Number: 80, Value: "שמונים"},
		{Number: 70, Value: "שבעים"},
		{Number: 60, Value: "שישים"},
		{Number: 50, Value: "חמישים"},
		{Number: 40, Value: "ארבעים"},
		{Number: 30, Value: "שלושים"},
		{Number: 20, Value: "עשרים"},
		{Number: 19, Value: "תשע עשרה"},
		{Number: 18, Value: "שמונה עשרה"},
		{Number: 17, Value: "שבע עשרה"},
		{Number: 16, Value: "שש עשרה"},
		{Number: 15, Value: "חמש עשרה"},
		{Number: 14, Value: "ארבע עשרה"},
		{Number: 13, Value: "שלוש עשרה"},
		{Number: 12, Value: "שתים עשרה"},
		{Number: 11, Value: "אחת עשרה"},
		{Number: 10, Value: "עשר"},
		{Number: 9, Value: "תשע"},
		{Number: 8, Value: "שמונה"},
		{Number: 7, Value: "שבע"},
		{Number: 6, Value: "שש"},
		{Number: 5, Value: "חמש"},
		{Number: 4, Value: "ארבע"},
		{Number: 3, Value: "שלוש"},
		{Number: 2, Value: "שתים"},
		{Number: 1, Value: "אחד"},
		{Number: 0, Value: "אפס"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "מאה"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "ראשון", Suffix: "׳", Masculine: "ראשון", Feminine: "ראשונה", Neuter: "ראשון"},
		{Number: 2, Word: "שני", Suffix: "׳", Masculine: "שני", Feminine: "שנייה", Neuter: "שני"},
		{Number: 3, Word: "שלישי", Suffix: "׳", Masculine: "שלישי", Feminine: "שלישית", Neuter: "שלישי"},
		{Number: 4, Word: "רביעי", Suffix: "׳", Masculine: "רביעי", Feminine: "רביעית", Neuter: "רביעי"},
		{Number: 5, Word: "חמישי", Suffix: "׳", Masculine: "חמישי", Feminine: "חמישית", Neuter: "חמישי"},
		{Number: 6, Word: "שישי", Suffix: "׳", Masculine: "שישי", Feminine: "שישית", Neuter: "שישי"},
		{Number: 7, Word: "שביעי", Suffix: "׳", Masculine: "שביעי", Feminine: "שביעית", Neuter: "שביעי"},
		{Number: 8, Word: "שמיני", Suffix: "׳", Masculine: "שמיני", Feminine: "שמינית", Neuter: "שמיני"},
		{Number: 9, Word: "תשיעי", Suffix: "׳", Masculine: "תשיעי", Feminine: "תשיעית", Neuter: "תשיעי"},
		{Number: 10, Word: "עשירי", Suffix: "׳", Masculine: "עשירי", Feminine: "עשירית", Neuter: "עשירי"},
		{Number: 11, Word: "אחד עשר", Suffix: "׳", Masculine: "אחד עשר", Feminine: "אחת עשרה", Neuter: "אחד עשר"},
		{Number: 12, Word: "שנים עשר", Suffix: "׳", Masculine: "שנים עשר", Feminine: "שתים עשרה", Neuter: "שנים עשר"},
		{Number: 13, Word: "שלושה עשר", Suffix: "׳", Masculine: "שלושה עשר", Feminine: "שלוש עשרה", Neuter: "שלושה עשר"},
		{Number: 14, Word: "ארבעה עשר", Suffix: "׳", Masculine: "ארבעה עשר", Feminine: "ארבע עשרה", Neuter: "ארבעה עשר"},
		{Number: 15, Word: "חמישה עשר", Suffix: "׳", Masculine: "חמישה עשר", Feminine: "חמש עשרה", Neuter: "חמישה עשר"},
		{Number: 16, Word: "שישה עשר", Suffix: "׳", Masculine: "שישה עשר", Feminine: "שש עשרה", Neuter: "שישה עשר"},
		{Number: 17, Word: "שבעה עשר", Suffix: "׳", Masculine: "שבעה עשר", Feminine: "שבע עשרה", Neuter: "שבעה עשר"},
		{Number: 18, Word: "שמונה עשר", Suffix: "׳", Masculine: "שמונה עשר", Feminine: "שמונה עשרה", Neuter: "שמונה עשר"},
		{Number: 19, Word: "תשעה עשר", Suffix: "׳", Masculine: "תשעה עשר", Feminine: "תשע עשרה", Neuter: "תשעה עשר"},
		{Number: 20, Word: "עשרים", Suffix: "׳", Masculine: "עשרים", Feminine: "עשרים", Neuter: "עשרים"},
		{Number: 21, Word: "עשרים ואחד", Suffix: "׳", Masculine: "עשרים ואחד", Feminine: "עשרים ואחת", Neuter: "עשרים ואחד"},
		{Number: 30, Word: "שלושים", Suffix: "׳", Masculine: "שלושים", Feminine: "שלושים", Neuter: "שלושים"},
		{Number: 40, Word: "ארבעים", Suffix: "׳", Masculine: "ארבעים", Feminine: "ארבעים", Neuter: "ארבעים"},
		{Number: 50, Word: "חמישים", Suffix: "׳", Masculine: "חמישים", Feminine: "חמישים", Neuter: "חמישים"},
		{Number: 60, Word: "שישים", Suffix: "׳", Masculine: "שישים", Feminine: "שישים", Neuter: "שישים"},
		{Number: 70, Word: "שבעים", Suffix: "׳", Masculine: "שבעים", Feminine: "שבעים", Neuter: "שבעים"},
		{Number: 80, Word: "שמונים", Suffix: "׳", Masculine: "שמונים", Feminine: "שמונים", Neuter: "שמונים"},
		{Number: 90, Word: "תשעים", Suffix: "׳", Masculine: "תשעים", Feminine: "תשעים", Neuter: "תשעים"},
		{Number: 100, Word: "מאה", Suffix: "׳", Masculine: "מאה", Feminine: "מאה", Neuter: "מאה"},
		{Number: 1000, Word: "אלף", Suffix: "׳", Masculine: "אלף", Feminine: "אלף", Neuter: "אלף"},
	},
	LocaleFormatter: &HebrewFormatter{},
}

// HebrewFormatter handles Hebrew (he-IL) formatting
type HebrewFormatter struct{}

func (f *HebrewFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *HebrewFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *HebrewFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *HebrewFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *HebrewFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *HebrewFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	// Standard decimal chopping - round to specified precision
	return amount.Round(int32(precision))
}


func (f *HebrewFormatter) FormatDecimalNumber(amount float64) string {
	return DefaultFormatDecimalNumber(amount, ",", ".")
}
func (f *HebrewFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
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
