package locale

import (
	"strings"
	"github.com/shopspring/decimal"
)

// MSMYLocale represents the Malay (Malaysia) locale
var MSMYLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Ringgit",
		Plural:   "Ringgit",
		Singular: "Ringgit",
		Symbol:   "RM",
		FractionUnit: FractionUnit{
			Name:     "Sen",
			Plural:   "Sen",
			Singular: "Sen",
			Symbol:   "s",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Malaysia",
		Currency:       "MYR",
		ISO3166Alpha2:  "MY",
		ISO3166Alpha3:  "MYS",
		ISO3166Numeric: "458",
		Locale:         "ms-MY",
		Timezone:       []string{"Asia/Kuala_Lumpur"},
		Language:       "ms",
	},
	Texts: Texts{
		And:   "dan",
		Minus: "tolak",
		Only:  "sahaja",
		Point: "titik",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "satu kuadrilion"},
		{Number: 1000000000000, Value: "satu trilion"},
		{Number: 1000000000, Value: "satu bilion"},
		{Number: 1000000, Value: "satu juta"},
		{Number: 100000, Value: "seratus ribu"},
		{Number: 90000, Value: "sembilan puluh ribu"},
		{Number: 80000, Value: "lapan puluh ribu"},
		{Number: 70000, Value: "tujuh puluh ribu"},
		{Number: 60000, Value: "enam puluh ribu"},
		{Number: 50000, Value: "lima puluh ribu"},
		{Number: 40000, Value: "empat puluh ribu"},
		{Number: 30000, Value: "tiga puluh ribu"},
		{Number: 20000, Value: "dua puluh ribu"},
		{Number: 19000, Value: "sembilan belas ribu"},
		{Number: 18000, Value: "lapan belas ribu"},
		{Number: 17000, Value: "tujuh belas ribu"},
		{Number: 16000, Value: "enam belas ribu"},
		{Number: 15000, Value: "lima belas ribu"},
		{Number: 14000, Value: "empat belas ribu"},
		{Number: 13000, Value: "tiga belas ribu"},
		{Number: 12000, Value: "dua belas ribu"},
		{Number: 11000, Value: "sebelas ribu"},
		{Number: 10000, Value: "sepuluh ribu"},
		{Number: 9000, Value: "sembilan ribu"},
		{Number: 8000, Value: "lapan ribu"},
		{Number: 7000, Value: "tujuh ribu"},
		{Number: 6000, Value: "enam ribu"},
		{Number: 5000, Value: "lima ribu"},
		{Number: 4000, Value: "empat ribu"},
		{Number: 3000, Value: "tiga ribu"},
		{Number: 2000, Value: "dua ribu"},
		{Number: 1000, Value: "seribu"},
		{Number: 900, Value: "sembilan ratus"},
		{Number: 800, Value: "lapan ratus"},
		{Number: 700, Value: "tujuh ratus"},
		{Number: 600, Value: "enam ratus"},
		{Number: 500, Value: "lima ratus"},
		{Number: 400, Value: "empat ratus"},
		{Number: 300, Value: "tiga ratus"},
		{Number: 200, Value: "dua ratus"},
		{Number: 100, Value: "seratus"},
		{Number: 99, Value: "sembilan puluh sembilan"},
		{Number: 98, Value: "sembilan puluh lapan"},
		{Number: 97, Value: "sembilan puluh tujuh"},
		{Number: 96, Value: "sembilan puluh enam"},
		{Number: 95, Value: "sembilan puluh lima"},
		{Number: 94, Value: "sembilan puluh empat"},
		{Number: 93, Value: "sembilan puluh tiga"},
		{Number: 92, Value: "sembilan puluh dua"},
		{Number: 91, Value: "sembilan puluh satu"},
		{Number: 90, Value: "sembilan puluh"},
		{Number: 89, Value: "lapan puluh sembilan"},
		{Number: 88, Value: "lapan puluh lapan"},
		{Number: 87, Value: "lapan puluh tujuh"},
		{Number: 86, Value: "lapan puluh enam"},
		{Number: 85, Value: "lapan puluh lima"},
		{Number: 84, Value: "lapan puluh empat"},
		{Number: 83, Value: "lapan puluh tiga"},
		{Number: 82, Value: "lapan puluh dua"},
		{Number: 81, Value: "lapan puluh satu"},
		{Number: 80, Value: "lapan puluh"},
		{Number: 79, Value: "tujuh puluh sembilan"},
		{Number: 78, Value: "tujuh puluh lapan"},
		{Number: 77, Value: "tujuh puluh tujuh"},
		{Number: 76, Value: "tujuh puluh enam"},
		{Number: 75, Value: "tujuh puluh lima"},
		{Number: 74, Value: "tujuh puluh empat"},
		{Number: 73, Value: "tujuh puluh tiga"},
		{Number: 72, Value: "tujuh puluh dua"},
		{Number: 71, Value: "tujuh puluh satu"},
		{Number: 70, Value: "tujuh puluh"},
		{Number: 69, Value: "enam puluh sembilan"},
		{Number: 68, Value: "enam puluh lapan"},
		{Number: 67, Value: "enam puluh tujuh"},
		{Number: 66, Value: "enam puluh enam"},
		{Number: 65, Value: "enam puluh lima"},
		{Number: 64, Value: "enam puluh empat"},
		{Number: 63, Value: "enam puluh tiga"},
		{Number: 62, Value: "enam puluh dua"},
		{Number: 61, Value: "enam puluh satu"},
		{Number: 60, Value: "enam puluh"},
		{Number: 59, Value: "lima puluh sembilan"},
		{Number: 58, Value: "lima puluh lapan"},
		{Number: 57, Value: "lima puluh tujuh"},
		{Number: 56, Value: "lima puluh enam"},
		{Number: 55, Value: "lima puluh lima"},
		{Number: 54, Value: "lima puluh empat"},
		{Number: 53, Value: "lima puluh tiga"},
		{Number: 52, Value: "lima puluh dua"},
		{Number: 51, Value: "lima puluh satu"},
		{Number: 50, Value: "lima puluh"},
		{Number: 49, Value: "empat puluh sembilan"},
		{Number: 48, Value: "empat puluh lapan"},
		{Number: 47, Value: "empat puluh tujuh"},
		{Number: 46, Value: "empat puluh enam"},
		{Number: 45, Value: "empat puluh lima"},
		{Number: 44, Value: "empat puluh empat"},
		{Number: 43, Value: "empat puluh tiga"},
		{Number: 42, Value: "empat puluh dua"},
		{Number: 41, Value: "empat puluh satu"},
		{Number: 40, Value: "empat puluh"},
		{Number: 39, Value: "tiga puluh sembilan"},
		{Number: 38, Value: "tiga puluh lapan"},
		{Number: 37, Value: "tiga puluh tujuh"},
		{Number: 36, Value: "tiga puluh enam"},
		{Number: 35, Value: "tiga puluh lima"},
		{Number: 34, Value: "tiga puluh empat"},
		{Number: 33, Value: "tiga puluh tiga"},
		{Number: 32, Value: "tiga puluh dua"},
		{Number: 31, Value: "tiga puluh satu"},
		{Number: 30, Value: "tiga puluh"},
		{Number: 29, Value: "dua puluh sembilan"},
		{Number: 28, Value: "dua puluh lapan"},
		{Number: 27, Value: "dua puluh tujuh"},
		{Number: 26, Value: "dua puluh enam"},
		{Number: 25, Value: "dua puluh lima"},
		{Number: 24, Value: "dua puluh empat"},
		{Number: 23, Value: "dua puluh tiga"},
		{Number: 22, Value: "dua puluh dua"},
		{Number: 21, Value: "dua puluh satu"},
		{Number: 20, Value: "dua puluh"},
		{Number: 19, Value: "sembilan belas"},
		{Number: 18, Value: "lapan belas"},
		{Number: 17, Value: "tujuh belas"},
		{Number: 16, Value: "enam belas"},
		{Number: 15, Value: "lima belas"},
		{Number: 14, Value: "empat belas"},
		{Number: 13, Value: "tiga belas"},
		{Number: 12, Value: "dua belas"},
		{Number: 11, Value: "sebelas"},
		{Number: 10, Value: "sepuluh"},
		{Number: 9, Value: "sembilan"},
		{Number: 8, Value: "lapan"},
		{Number: 7, Value: "tujuh"},
		{Number: 6, Value: "enam"},
		{Number: 5, Value: "lima"},
		{Number: 4, Value: "empat"},
		{Number: 3, Value: "tiga"},
		{Number: 2, Value: "dua"},
		{Number: 1, Value: "satu"},
		{Number: 0, Value: "sifar"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Seratus"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "pertama", Suffix: "ke-", Masculine: "pertama", Feminine: "pertama", Neuter: "pertama"},
		{Number: 2, Word: "kedua", Suffix: "ke-", Masculine: "kedua", Feminine: "kedua", Neuter: "kedua"},
		{Number: 3, Word: "ketiga", Suffix: "ke-", Masculine: "ketiga", Feminine: "ketiga", Neuter: "ketiga"},
		{Number: 4, Word: "keempat", Suffix: "ke-", Masculine: "keempat", Feminine: "keempat", Neuter: "keempat"},
		{Number: 5, Word: "kelima", Suffix: "ke-", Masculine: "kelima", Feminine: "kelima", Neuter: "kelima"},
		{Number: 6, Word: "keenam", Suffix: "ke-", Masculine: "keenam", Feminine: "keenam", Neuter: "keenam"},
		{Number: 7, Word: "ketujuh", Suffix: "ke-", Masculine: "ketujuh", Feminine: "ketujuh", Neuter: "ketujuh"},
		{Number: 8, Word: "kelapan", Suffix: "ke-", Masculine: "kelapan", Feminine: "kelapan", Neuter: "kelapan"},
		{Number: 9, Word: "kesembilan", Suffix: "ke-", Masculine: "kesembilan", Feminine: "kesembilan", Neuter: "kesembilan"},
		{Number: 10, Word: "kesepuluh", Suffix: "ke-", Masculine: "kesepuluh", Feminine: "kesepuluh", Neuter: "kesepuluh"},
		{Number: 11, Word: "kesebelas", Suffix: "ke-", Masculine: "kesebelas", Feminine: "kesebelas", Neuter: "kesebelas"},
		{Number: 12, Word: "kedua belas", Suffix: "ke-", Masculine: "kedua belas", Feminine: "kedua belas", Neuter: "kedua belas"},
		{Number: 20, Word: "kedua puluh", Suffix: "ke-", Masculine: "kedua puluh", Feminine: "kedua puluh", Neuter: "kedua puluh"},
		{Number: 21, Word: "kedua puluh satu", Suffix: "ke-", Masculine: "kedua puluh satu", Feminine: "kedua puluh satu", Neuter: "kedua puluh satu"},
		{Number: 30, Word: "ketiga puluh", Suffix: "ke-", Masculine: "ketiga puluh", Feminine: "ketiga puluh", Neuter: "ketiga puluh"},
		{Number: 100, Word: "keseratus", Suffix: "ke-", Masculine: "keseratus", Feminine: "keseratus", Neuter: "keseratus"},
		{Number: 1000, Word: "keseribu", Suffix: "ke-", Masculine: "keseribu", Feminine: "keseribu", Neuter: "keseribu"},
	},
	LocaleFormatter: &MalaysianFormatter{},
}

// MalaysianFormatter handles Malaysian-specific formatting
type MalaysianFormatter struct{}

func (f *MalaysianFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *MalaysianFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *MalaysianFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *MalaysianFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *MalaysianFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *MalaysianFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 0
	}
	return amount.Truncate(int32(precision))
}


func (f *MalaysianFormatter) FormatDecimalNumber(amount float64) string {
	return DefaultFormatDecimalNumber(amount, ",", ".")
}
func (f *MalaysianFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
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
