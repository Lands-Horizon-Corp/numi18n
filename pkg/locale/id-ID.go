package locale

import (
	"strings"

	"github.com/shopspring/decimal"
)

// IDLocale is a NumI18NLocale configured for Indonesia (id-ID)
var IDLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Rupiah",
		Plural:   "Rupiah",
		Singular: "Rupiah",
		Symbol:   "Rp",
		FractionUnit: FractionUnit{
			Name:     "Sen",
			Plural:   "Sen",
			Singular: "Sen",
			Symbol:   "s",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Indonesia",
		Currency:       "IDR",
		ISO3166Alpha2:  "ID",
		ISO3166Alpha3:  "IDN",
		ISO3166Numeric: "360",
		Locale:         "id-ID",
		Timezone:       []string{"Asia/Jakarta"},
		Language:       "id",
	},
	Texts: Texts{
		And:   "Dan",
		Minus: "Minus",
		Only:  "Hanya",
		Point: "Koma",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Kuadrilun"},
		{Number: 1000000000000, Value: "Trilun"},
		{Number: 1000000000, Value: "Miliar"},
		{Number: 1000000, Value: "Juta"},
		{Number: 1000, Value: "Ribu"},
		{Number: 100, Value: "Ratus"},
		{Number: 90, Value: "Sembilan Puluh"},
		{Number: 80, Value: "Delapan Puluh"},
		{Number: 70, Value: "Tujuh Puluh"},
		{Number: 60, Value: "Enam Puluh"},
		{Number: 50, Value: "Lima Puluh"},
		{Number: 40, Value: "Empat Puluh"},
		{Number: 30, Value: "Tiga Puluh"},
		{Number: 20, Value: "Dua Puluh"},
		{Number: 19, Value: "Sembilan Belas"},
		{Number: 18, Value: "Delapan Belas"},
		{Number: 17, Value: "Tujuh Belas"},
		{Number: 16, Value: "Enam Belas"},
		{Number: 15, Value: "Lima Belas"},
		{Number: 14, Value: "Empat Belas"},
		{Number: 13, Value: "Tiga Belas"},
		{Number: 12, Value: "Dua Belas"},
		{Number: 11, Value: "Sebelas"},
		{Number: 10, Value: "Sepuluh"},
		{Number: 9, Value: "Sembilan"},
		{Number: 8, Value: "Delapan"},
		{Number: 7, Value: "Tujuh"},
		{Number: 6, Value: "Enam"},
		{Number: 5, Value: "Lima"},
		{Number: 4, Value: "Empat"},
		{Number: 3, Value: "Tiga"},
		{Number: 2, Value: "Dua"},
		{Number: 1, Value: "Satu"},
		{Number: 0, Value: "Nol"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Seratus"},
		{Number: 1000, Value: "Seribu"},
		{Number: 1000000, Value: "Satu Juta"},
		{Number: 1000000000, Value: "Satu Miliar"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "pertama", Suffix: ".", Masculine: "pertama", Feminine: "pertama", Neuter: "pertama"},
		{Number: 2, Word: "kedua", Suffix: ".", Masculine: "kedua", Feminine: "kedua", Neuter: "kedua"},
		{Number: 3, Word: "ketiga", Suffix: ".", Masculine: "ketiga", Feminine: "ketiga", Neuter: "ketiga"},
		{Number: 4, Word: "keempat", Suffix: ".", Masculine: "keempat", Feminine: "keempat", Neuter: "keempat"},
		{Number: 5, Word: "kelima", Suffix: ".", Masculine: "kelima", Feminine: "kelima", Neuter: "kelima"},
		{Number: 6, Word: "keenam", Suffix: ".", Masculine: "keenam", Feminine: "keenam", Neuter: "keenam"},
		{Number: 7, Word: "ketujuh", Suffix: ".", Masculine: "ketujuh", Feminine: "ketujuh", Neuter: "ketujuh"},
		{Number: 8, Word: "kedelapan", Suffix: ".", Masculine: "kedelapan", Feminine: "kedelapan", Neuter: "kedelapan"},
		{Number: 9, Word: "kesembilan", Suffix: ".", Masculine: "kesembilan", Feminine: "kesembilan", Neuter: "kesembilan"},
		{Number: 10, Word: "kesepuluh", Suffix: ".", Masculine: "kesepuluh", Feminine: "kesepuluh", Neuter: "kesepuluh"},
		{Number: 11, Word: "kesebelas", Suffix: ".", Masculine: "kesebelas", Feminine: "kesebelas", Neuter: "kesebelas"},
		{Number: 12, Word: "kedua belas", Suffix: ".", Masculine: "kedua belas", Feminine: "kedua belas", Neuter: "kedua belas"},
		{Number: 13, Word: "ketiga belas", Suffix: ".", Masculine: "ketiga belas", Feminine: "ketiga belas", Neuter: "ketiga belas"},
		{Number: 14, Word: "keempat belas", Suffix: ".", Masculine: "keempat belas", Feminine: "keempat belas", Neuter: "keempat belas"},
		{Number: 15, Word: "kelima belas", Suffix: ".", Masculine: "kelima belas", Feminine: "kelima belas", Neuter: "kelima belas"},
		{Number: 16, Word: "keenam belas", Suffix: ".", Masculine: "keenam belas", Feminine: "keenam belas", Neuter: "keenam belas"},
		{Number: 17, Word: "ketujuh belas", Suffix: ".", Masculine: "ketujuh belas", Feminine: "ketujuh belas", Neuter: "ketujuh belas"},
		{Number: 18, Word: "kedelapan belas", Suffix: ".", Masculine: "kedelapan belas", Feminine: "kedelapan belas", Neuter: "kedelapan belas"},
		{Number: 19, Word: "kesembilan belas", Suffix: ".", Masculine: "kesembilan belas", Feminine: "kesembilan belas", Neuter: "kesembilan belas"},
		{Number: 20, Word: "kedua puluh", Suffix: ".", Masculine: "kedua puluh", Feminine: "kedua puluh", Neuter: "kedua puluh"},
		{Number: 21, Word: "kedua puluh satu", Suffix: ".", Masculine: "kedua puluh satu", Feminine: "kedua puluh satu", Neuter: "kedua puluh satu"},
		{Number: 30, Word: "ketiga puluh", Suffix: ".", Masculine: "ketiga puluh", Feminine: "ketiga puluh", Neuter: "ketiga puluh"},
		{Number: 40, Word: "keempat puluh", Suffix: ".", Masculine: "keempat puluh", Feminine: "keempat puluh", Neuter: "keempat puluh"},
		{Number: 50, Word: "kelima puluh", Suffix: ".", Masculine: "kelima puluh", Feminine: "kelima puluh", Neuter: "kelima puluh"},
		{Number: 60, Word: "keenam puluh", Suffix: ".", Masculine: "keenam puluh", Feminine: "keenam puluh", Neuter: "keenam puluh"},
		{Number: 70, Word: "ketujuh puluh", Suffix: ".", Masculine: "ketujuh puluh", Feminine: "ketujuh puluh", Neuter: "ketujuh puluh"},
		{Number: 80, Word: "kedelapan puluh", Suffix: ".", Masculine: "kedelapan puluh", Feminine: "kedelapan puluh", Neuter: "kedelapan puluh"},
		{Number: 90, Word: "kesembilan puluh", Suffix: ".", Masculine: "kesembilan puluh", Feminine: "kesembilan puluh", Neuter: "kesembilan puluh"},
		{Number: 100, Word: "keseratus", Suffix: ".", Masculine: "keseratus", Feminine: "keseratus", Neuter: "keseratus"},
		{Number: 1000, Word: "keseribu", Suffix: ".", Masculine: "keseribu", Feminine: "keseribu", Neuter: "keseribu"},
	},
	LocaleFormatter: &IndonesianFormatter{},
}

// IndonesianFormatter handles Indonesian (id-ID) formatting
type IndonesianFormatter struct{}

func (f *IndonesianFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	// First check exact mappings
	exactResult := ConvertToWordsWithExactMappingInt64(number, targetLocale)
	if exactResult != ConvertToWordsGenericInt64(number, targetLocale) {
		return exactResult // Found exact mapping
	}

	// Handle Indonesian special forms by replacing the generic forms
	genericResult := ConvertToWordsGenericInt64(number, targetLocale)

	// Replace "Satu Ratus" with "Seratus" (one hundred)
	if genericResult == "Satu Ratus" {
		return "Seratus"
	}

	// Replace "Satu Ribu" with "Seribu" (one thousand)
	if genericResult == "Satu Ribu" {
		return "Seribu"
	}

	// For compound numbers, replace the patterns
	result := genericResult
	result = strings.ReplaceAll(result, "Satu Ratus", "Seratus")
	result = strings.ReplaceAll(result, "Satu Ribu", "Seribu")

	return result
}

func (f *IndonesianFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *IndonesianFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *IndonesianFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *IndonesianFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *IndonesianFormatter) ChopDecimal(number decimal.Decimal, places int) decimal.Decimal {
	return number.Truncate(int32(places))
}


func (f *IndonesianFormatter) FormatDecimalNumber(amount float64) string {
	return DefaultFormatDecimalNumber(amount, ",", ".")
}
func (f *IndonesianFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
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
