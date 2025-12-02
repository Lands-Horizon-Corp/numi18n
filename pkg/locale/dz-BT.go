package locale

import (
	"strings"
	"github.com/shopspring/decimal"
)

// DZBTLocale is a NumI18NLocale configured for Dzongkha (Bhutan) - dz-BT
var DZBTLocale = NumI18NLocale{
	LocaleFormatter: &DzongkhaFormatter{},
	Currency: Currency{
		Name:     "Ngultrum",
		Plural:   "Ngultrum",
		Singular: "Ngultrum",
		Symbol:   "Nu",
		FractionUnit: FractionUnit{
			Name:     "Chhertum",
			Plural:   "Chhertum",
			Singular: "Chhertum",
			Symbol:   "Ch",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Bhutan",
		Currency:       "BTN",
		ISO3166Alpha2:  "BT",
		ISO3166Alpha3:  "BTN",
		ISO3166Numeric: "064",
		Locale:         "dz-BT",
		Timezone:       []string{"Asia/Thimphu"},
		Language:       "dz",
	},
	Texts: Texts{
		And:   "དང་",
		Minus: "མིང་མེད་",
		Only:  "ཡང་ན་",
		Point: "ཐུང་",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "ཀྱུར་དྲི་ལི་ཨན་"},
		{Number: 1000000000000, Value: "ཊྲི་ལི་ཨན་"},
		{Number: 1000000000, Value: "བི་ལི་ཨན་"},
		{Number: 1000000, Value: "མི་ལི་ཨན་"},
		{Number: 1000, Value: "མི་ལ་"},
		{Number: 100, Value: "སུང་"},
		{Number: 90, Value: "དང་ལྔ་དུས་དེབ་"},
		{Number: 80, Value: "བརྒྱ་བརྒྱ་"},
		{Number: 70, Value: "བདུན་བརྒྱ་"},
		{Number: 60, Value: "དགུ་བརྒྱ་"},
		{Number: 50, Value: "ལྔ་བརྒྱ་"},
		{Number: 40, Value: "བཞི་བརྒྱ་"},
		{Number: 30, Value: "སུམ་བརྒྱ་"},
		{Number: 20, Value: "ཉེ་བརྒྱ་"},
		{Number: 19, Value: "དགུ་གཅིག་"},
		{Number: 18, Value: "བཞི་གཉིས་"},
		{Number: 17, Value: "བདུན་གཉིས་"},
		{Number: 16, Value: "དགུ་གཉིས་"},
		{Number: 15, Value: "ལྔ་གཉིས་"},
		{Number: 14, Value: "བཞི་བཞི་"},
		{Number: 13, Value: "སུམ་གསུམ་"},
		{Number: 12, Value: "ཉེ་གཉིས་"},
		{Number: 11, Value: "གཅིག་གཅིག་"},
		{Number: 10, Value: "བཅུ་"},
		{Number: 9, Value: "དགུ་"},
		{Number: 8, Value: "བརྒྱད་"},
		{Number: 7, Value: "བདུན་"},
		{Number: 6, Value: "དྲུག་"},
		{Number: 5, Value: "ལྔ་"},
		{Number: 4, Value: "བཞི་"},
		{Number: 3, Value: "སུམ་"},
		{Number: 2, Value: "གཉིས་"},
		{Number: 1, Value: "གཅིག་"},
		{Number: 0, Value: "སོ་རོ་"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "སུང་"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "དང་པོ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 2, Word: "གཉིས་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 3, Word: "གསུམ་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 4, Word: "བཞི་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 5, Word: "ལྔ་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 6, Word: "དྲུག་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 7, Word: "བདུན་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 8, Word: "བརྒྱད་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 9, Word: "དགུ་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 10, Word: "བཅུ་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 11, Word: "བཅུ་གཅིག་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 12, Word: "བཅུ་གཉིས་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 13, Word: "བཅུ་གསུམ་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 14, Word: "བཅུ་བཞི་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 15, Word: "བཅུ་ལྔ་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 16, Word: "བཅུ་དྲུག་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 17, Word: "བཅུ་བདུན་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 18, Word: "བཅུ་བརྒྱད་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 19, Word: "བཅུ་དགུ་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 20, Word: "ཉི་ཤུ་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 21, Word: "ཉི་ཤུ་རྩ་གཅིག་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 30, Word: "སུམ་བརྒྱ་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 40, Word: "བཞི་བརྒྱ་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 50, Word: "ལྔ་བརྒྱ་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 100, Word: "སུང་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 1000, Word: "མི་ལ་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 1000000, Word: "མི་ལི་ཨན་པ་", Suffix: "པ་", Masculine: "", Feminine: "", Neuter: ""},
	},
}

// DzongkhaFormatter handles Dzongkha (Bhutan) formatting
type DzongkhaFormatter struct{}

func (f *DzongkhaFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *DzongkhaFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *DzongkhaFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *DzongkhaFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *DzongkhaFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *DzongkhaFormatter) ChopDecimal(value decimal.Decimal, precision int) decimal.Decimal {
	return value.Truncate(int32(precision))
}


func (f *DzongkhaFormatter) FormatDecimalNumber(amount float64) string {
	return DefaultFormatDecimalNumber(amount, ",", ".")
}
func (f *DzongkhaFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
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
