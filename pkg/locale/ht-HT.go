package locale

import (
	"strings"
	"github.com/shopspring/decimal"
)

// HTLocale is a NumI18NLocale configured for Haiti (ht-HT)
var HTLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Goud",
		Plural:   "Goud",
		Singular: "Goud",
		Symbol:   "G",
		FractionUnit: FractionUnit{
			Name:     "Santim",
			Plural:   "Santim",
			Singular: "Santim",
			Symbol:   "¢",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Haiti",
		Currency:       "HTG",
		ISO3166Alpha2:  "HT",
		ISO3166Alpha3:  "HTI",
		ISO3166Numeric: "332",
		Locale:         "ht-HT",
		Timezone:       []string{"America/Port-au-Prince"},
		Language:       "ht",
	},
	Texts: Texts{
		And:   "ak",
		Minus: "Mwens",
		Only:  "sèlman",
		Point: "pwen",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "yon mil milya milya"},
		{Number: 1000000000000, Value: "yon mil milya"},
		{Number: 1000000000, Value: "yon milya"},
		{Number: 1000000, Value: "yon milyon"},
		{Number: 1000, Value: "yon mil"},
		{Number: 100, Value: "yon san"},
		{Number: 90, Value: "katre-ven-dis"},
		{Number: 80, Value: "katre-ven"},
		{Number: 70, Value: "swasann-dis"},
		{Number: 60, Value: "swasann"},
		{Number: 50, Value: "senkann"},
		{Number: 40, Value: "karant"},
		{Number: 30, Value: "trant"},
		{Number: 20, Value: "ven"},
		{Number: 19, Value: "dis-nèf"},
		{Number: 18, Value: "dis-uit"},
		{Number: 17, Value: "dis-sèt"},
		{Number: 16, Value: "sèz"},
		{Number: 15, Value: "kenz"},
		{Number: 14, Value: "katòz"},
		{Number: 13, Value: "trèz"},
		{Number: 12, Value: "douz"},
		{Number: 11, Value: "onz"},
		{Number: 10, Value: "dis"},
		{Number: 9, Value: "nèf"},
		{Number: 8, Value: "uit"},
		{Number: 7, Value: "sèt"},
		{Number: 6, Value: "sis"},
		{Number: 5, Value: "senk"},
		{Number: 4, Value: "kat"},
		{Number: 3, Value: "twa"},
		{Number: 2, Value: "de"},
		{Number: 1, Value: "yon"},
		{Number: 0, Value: "zewo"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "yon san"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "premye", Suffix: "ye", Masculine: "premye", Feminine: "premye", Neuter: "premye"},
		{Number: 2, Word: "dezyèm", Suffix: "èm", Masculine: "dezyèm", Feminine: "dezyèm", Neuter: "dezyèm"},
		{Number: 3, Word: "twazyèm", Suffix: "èm", Masculine: "twazyèm", Feminine: "twazyèm", Neuter: "twazyèm"},
		{Number: 4, Word: "katriyèm", Suffix: "èm", Masculine: "katriyèm", Feminine: "katriyèm", Neuter: "katriyèm"},
		{Number: 5, Word: "senkyèm", Suffix: "èm", Masculine: "senkyèm", Feminine: "senkyèm", Neuter: "senkyèm"},
		{Number: 6, Word: "sizyèm", Suffix: "èm", Masculine: "sizyèm", Feminine: "sizyèm", Neuter: "sizyèm"},
		{Number: 7, Word: "sètyèm", Suffix: "èm", Masculine: "sètyèm", Feminine: "sètyèm", Neuter: "sètyèm"},
		{Number: 8, Word: "ityèm", Suffix: "èm", Masculine: "ityèm", Feminine: "ityèm", Neuter: "ityèm"},
		{Number: 9, Word: "nèfyèm", Suffix: "èm", Masculine: "nèfyèm", Feminine: "nèfyèm", Neuter: "nèfyèm"},
		{Number: 10, Word: "dizyèm", Suffix: "èm", Masculine: "dizyèm", Feminine: "dizyèm", Neuter: "dizyèm"},
		{Number: 11, Word: "onzyèm", Suffix: "èm", Masculine: "onzyèm", Feminine: "onzyèm", Neuter: "onzyèm"},
		{Number: 12, Word: "douzyèm", Suffix: "èm", Masculine: "douzyèm", Feminine: "douzyèm", Neuter: "douzyèm"},
		{Number: 13, Word: "trèzyèm", Suffix: "èm", Masculine: "trèzyèm", Feminine: "trèzyèm", Neuter: "trèzyèm"},
		{Number: 14, Word: "katòzyèm", Suffix: "èm", Masculine: "katòzyèm", Feminine: "katòzyèm", Neuter: "katòzyèm"},
		{Number: 15, Word: "kenzyèm", Suffix: "èm", Masculine: "kenzyèm", Feminine: "kenzyèm", Neuter: "kenzyèm"},
		{Number: 16, Word: "sèzyèm", Suffix: "èm", Masculine: "sèzyèm", Feminine: "sèzyèm", Neuter: "sèzyèm"},
		{Number: 17, Word: "dis-sètyèm", Suffix: "èm", Masculine: "dis-sètyèm", Feminine: "dis-sètyèm", Neuter: "dis-sètyèm"},
		{Number: 18, Word: "dis-ityèm", Suffix: "èm", Masculine: "dis-ityèm", Feminine: "dis-ityèm", Neuter: "dis-ityèm"},
		{Number: 19, Word: "dis-nèfyèm", Suffix: "èm", Masculine: "dis-nèfyèm", Feminine: "dis-nèfyèm", Neuter: "dis-nèfyèm"},
		{Number: 20, Word: "ventyèm", Suffix: "èm", Masculine: "ventyèm", Feminine: "ventyèm", Neuter: "ventyèm"},
		{Number: 21, Word: "ven-etyèm", Suffix: "èm", Masculine: "ven-etyèm", Feminine: "ven-etyèm", Neuter: "ven-etyèm"},
		{Number: 30, Word: "trantyèm", Suffix: "èm", Masculine: "trantyèm", Feminine: "trantyèm", Neuter: "trantyèm"},
		{Number: 40, Word: "karantyèm", Suffix: "èm", Masculine: "karantyèm", Feminine: "karantyèm", Neuter: "karantyèm"},
		{Number: 50, Word: "senkannyèm", Suffix: "èm", Masculine: "senkannyèm", Feminine: "senkannyèm", Neuter: "senkannyèm"},
		{Number: 60, Word: "swasannyèm", Suffix: "èm", Masculine: "swasannyèm", Feminine: "swasannyèm", Neuter: "swasannyèm"},
		{Number: 70, Word: "swasann-dizyèm", Suffix: "èm", Masculine: "swasann-dizyèm", Feminine: "swasann-dizyèm", Neuter: "swasann-dizyèm"},
		{Number: 80, Word: "katre-ventyèm", Suffix: "èm", Masculine: "katre-ventyèm", Feminine: "katre-ventyèm", Neuter: "katre-ventyèm"},
		{Number: 90, Word: "katre-ven-dizyèm", Suffix: "èm", Masculine: "katre-ven-dizyèm", Feminine: "katre-ven-dizyèm", Neuter: "katre-ven-dizyèm"},
		{Number: 100, Word: "san-tyèm", Suffix: "èm", Masculine: "san-tyèm", Feminine: "san-tyèm", Neuter: "san-tyèm"},
		{Number: 1000, Word: "mil-tyèm", Suffix: "èm", Masculine: "mil-tyèm", Feminine: "mil-tyèm", Neuter: "mil-tyèm"},
	},
	LocaleFormatter: &HaitianCreoleFormatter{},
}

// HaitianCreoleFormatter handles Haitian Creole formatting
type HaitianCreoleFormatter struct{}

func (f *HaitianCreoleFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *HaitianCreoleFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *HaitianCreoleFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *HaitianCreoleFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *HaitianCreoleFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *HaitianCreoleFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	return amount.Truncate(int32(precision))
}


func (f *HaitianCreoleFormatter) FormatDecimalNumber(amount float64) string {
	return DefaultFormatDecimalNumber(amount, ",", ".")
}
func (f *HaitianCreoleFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
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
