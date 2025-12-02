package locale

import (
	"strings"
	"github.com/shopspring/decimal"
)

// AZAZLocale is a NumI18NLocale configured for Azerbaijani (Azerbaijan) - az-AZ
var AZAZLocale = NumI18NLocale{
	LocaleFormatter: &AzerbaijaniFormatter{},
	Currency: Currency{
		Name:     "Manat",
		Plural:   "Manat",
		Singular: "Manat",
		Symbol:   "₼",
		FractionUnit: FractionUnit{
			Name:     "Qəpik",
			Plural:   "Qəpiklər",
			Singular: "Qəpik",
			Symbol:   "q",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Azerbaijan",
		Currency:       "AZN",
		ISO3166Alpha2:  "AZ",
		ISO3166Alpha3:  "AZE",
		ISO3166Numeric: "031",
		Locale:         "az-AZ",
		Timezone:       []string{"Asia/Baku"},
		Language:       "az",
	},
	Texts: Texts{
		And:   "Və",
		Minus: "Mənfi",
		Only:  "Yalnız",
		Point: "Nöqtə",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Kvadrilyon"},
		{Number: 1000000000000, Value: "Trilyon"},
		{Number: 1000000000, Value: "Milyard"},
		{Number: 1000000, Value: "Milyon"},
		{Number: 1000, Value: "Min"},
		{Number: 100, Value: "Yüz"},
		{Number: 90, Value: "Doxsan"},
		{Number: 80, Value: "Səksən"},
		{Number: 70, Value: "Yetmiş"},
		{Number: 60, Value: "Altmış"},
		{Number: 50, Value: "Əlli"},
		{Number: 40, Value: "Qırx"},
		{Number: 30, Value: "Otuz"},
		{Number: 20, Value: "İyirmi"},
		{Number: 19, Value: "On doqquz"},
		{Number: 18, Value: "On səkkiz"},
		{Number: 17, Value: "On yeddi"},
		{Number: 16, Value: "On altı"},
		{Number: 15, Value: "On beş"},
		{Number: 14, Value: "On dörd"},
		{Number: 13, Value: "On üç"},
		{Number: 12, Value: "On iki"},
		{Number: 11, Value: "On bir"},
		{Number: 10, Value: "On"},
		{Number: 9, Value: "Doqquz"},
		{Number: 8, Value: "Səkkiz"},
		{Number: 7, Value: "Yeddi"},
		{Number: 6, Value: "Altı"},
		{Number: 5, Value: "Beş"},
		{Number: 4, Value: "Dörd"},
		{Number: 3, Value: "Üç"},
		{Number: 2, Value: "İki"},
		{Number: 1, Value: "Bir"},
		{Number: 0, Value: "Sıfır"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Bir Yüz"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "Birinci", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 2, Word: "İkinci", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 3, Word: "Üçüncü", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 4, Word: "Dördüncü", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 5, Word: "Beşinci", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 6, Word: "Altıncı", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 7, Word: "Yeddinci", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 8, Word: "Səkkizinci", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 9, Word: "Doqquzuncu", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 10, Word: "Onuncu", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 11, Word: "On birinci", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 12, Word: "On ikinci", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 13, Word: "On üçüncü", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 14, Word: "On dördüncü", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 15, Word: "On beşinci", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 16, Word: "On altıncı", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 17, Word: "On yeddinci", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 18, Word: "On səkkizinci", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 19, Word: "On doqquzuncu", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 20, Word: "İyirminci", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 21, Word: "İyirmi birinci", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 30, Word: "Otuzuncu", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 40, Word: "Qırxıncı", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 50, Word: "Əllinci", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 60, Word: "Altmışıncı", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 70, Word: "Yetmişinci", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 80, Word: "Səksəninci", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 90, Word: "Doxsanıncı", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 100, Word: "Yüzüncü", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 1000, Word: "Mininci", Suffix: "", Masculine: "", Feminine: "", Neuter: ""},
	},
}

// AzerbaijaniFormatter handles Azerbaijani formatting
type AzerbaijaniFormatter struct{}

func (f *AzerbaijaniFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *AzerbaijaniFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *AzerbaijaniFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *AzerbaijaniFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *AzerbaijaniFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *AzerbaijaniFormatter) ChopDecimal(d decimal.Decimal, precision int) decimal.Decimal {
	return d.Truncate(int32(precision))
}


func (f *AzerbaijaniFormatter) FormatDecimalNumber(amount float64) string {
	return DefaultFormatDecimalNumber(amount, ",", ".")
}
func (f *AzerbaijaniFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
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
