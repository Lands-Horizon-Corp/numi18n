package locale

import (
	"strings"
	"github.com/shopspring/decimal"
)

// DVMVLocale is a NumI18NLocale configured for Dhivehi (Maldives) - dv-MV
var DVMVLocale = NumI18NLocale{
	LocaleFormatter: &DhivehiFormatter{},
	Currency: Currency{
		Name:     "Rufiyaa",
		Plural:   "Rufiyaa",
		Singular: "Rufiyaa",
		Symbol:   "Rf",
		FractionUnit: FractionUnit{
			Name:     "Laari",
			Plural:   "Laari",
			Singular: "Laari",
			Symbol:   "l",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Maldives",
		Currency:       "MVR",
		ISO3166Alpha2:  "MV",
		ISO3166Alpha3:  "MDV",
		ISO3166Numeric: "462",
		Locale:         "dv-MV",
		Timezone:       []string{"Indian/Maldives"},
		Language:       "dv",
	},
	Texts: Texts{
		And:   "އަންޑް",
		Minus: "މިނަސް",
		Only:  "ނުވަތަ",
		Point: "ޕޮއިންޓް",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "ކްވަޑްރިލިއަން"},
		{Number: 1000000000000, Value: "ޓްރިލިއަން"},
		{Number: 1000000000, Value: "ބިލިއަން"},
		{Number: 1000000, Value: "މިލިއަން"},
		{Number: 1000, Value: "ތިސަން"},
		{Number: 100, Value: "ސްތޯ"},
		{Number: 90, Value: "ނަވެސް"},
		{Number: 80, Value: "އޮވެސް"},
		{Number: 70, Value: "ސެބެތް"},
		{Number: 60, Value: "ސެކްސެތް"},
		{Number: 50, Value: "ފިވެސް"},
		{Number: 40, Value: "ފިސްރަ"},
		{Number: 30, Value: "ޓްރިސް"},
		{Number: 20, Value: "ޓްވިން"},
		{Number: 19, Value: "ނިންޓީން"},
		{Number: 18, Value: "އެއިޓީން"},
		{Number: 17, Value: "ސެވެންޓީން"},
		{Number: 16, Value: "ސިސްޓީން"},
		{Number: 15, Value: "ފިފްޓީން"},
		{Number: 14, Value: "ފޮޓްޝެން"},
		{Number: 13, Value: "ތްރީޝެން"},
		{Number: 12, Value: "ޓްވޮލްވް"},
		{Number: 11, Value: "އެލްވް"},
		{Number: 10, Value: "ޓެން"},
		{Number: 9, Value: "ނަވް"},
		{Number: 8, Value: "އޮވް"},
		{Number: 7, Value: "ސެވް"},
		{Number: 6, Value: "ސެސް"},
		{Number: 5, Value: "ފިފް"},
		{Number: 4, Value: "ފޯރ"},
		{Number: 3, Value: "ތްރީ"},
		{Number: 2, Value: "ޓްނޯ"},
		{Number: 1, Value: "އަން"},
		{Number: 0, Value: "ސިރޯ"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "ސްތޯ"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "ފުރަތަމަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 2, Word: "ދެވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 3, Word: "ތިންވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 4, Word: "ހަތަރުވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 5, Word: "ފަސްވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 6, Word: "ހަވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 7, Word: "ހަތްވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 8, Word: "އަށްވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 9, Word: "ނުވަވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 10, Word: "ދިހަވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 11, Word: "އެގާރަވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 12, Word: "ބާރަވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 13, Word: "ތޭރަވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 14, Word: "ސާދަވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 15, Word: "ފަނަރަވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 16, Word: "ސޯޅަވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 17, Word: "ސަތާރަވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 18, Word: "އަށާރަވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 19, Word: "ނަވާރަވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 20, Word: "ވަނަވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 21, Word: "އެކާވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 30, Word: "ތިރީހަވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 40, Word: "ސާޅަވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 50, Word: "ފަންސަވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 100, Word: "ސޭސާވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 1000, Word: "ހާސްވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 1000000, Word: "މިލިއަންވަނަ", Suffix: "ވަނަ", Masculine: "", Feminine: "", Neuter: ""},
	},
}

// DhivehiFormatter handles Dhivehi (Maldives) formatting
type DhivehiFormatter struct{}

func (f *DhivehiFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *DhivehiFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *DhivehiFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *DhivehiFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *DhivehiFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *DhivehiFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	return amount.Round(int32(precision))
}


func (f *DhivehiFormatter) FormatDecimalNumber(amount float64) string {
	return DefaultFormatDecimalNumber(amount, ",", ".")
}
func (f *DhivehiFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
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
