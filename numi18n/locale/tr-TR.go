package locale

import (
	"github.com/shopspring/decimal"
)

// TRTRLocale represents the Turkish (Turkey) locale
var TRTRLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "lira",
		Plural:   "lira",
		Singular: "lira",
		Symbol:   "₺",
		FractionUnit: FractionUnit{
			Name:     "kuruş",
			Plural:   "kuruş",
			Singular: "kuruş",
			Symbol:   "kr",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Turkey",
		Currency:       "TRY",
		ISO3166Alpha2:  "TR",
		ISO3166Alpha3:  "TUR",
		ISO3166Numeric: "792",
		Locale:         "tr-TR",
		Timezone:       []string{"Europe/Istanbul"},
		Language:       "tr",
		Emoji:          "🇹🇷",
	},
	Texts: Texts{
		And:   "ve",
		Minus: "eksi",
		Only:  "sadece",
		Point: "virgül",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "bir katrilyon"},
		{Number: 1000000000000, Value: "bir trilyon"},
		{Number: 1000000000, Value: "bir milyar"},
		{Number: 1000000, Value: "bir milyon"},
		{Number: 100000, Value: "yüz bin"},
		{Number: 90000, Value: "doksan bin"},
		{Number: 80000, Value: "seksen bin"},
		{Number: 70000, Value: "yetmiş bin"},
		{Number: 60000, Value: "altmış bin"},
		{Number: 50000, Value: "elli bin"},
		{Number: 40000, Value: "kırk bin"},
		{Number: 30000, Value: "otuz bin"},
		{Number: 20000, Value: "yirmi bin"},
		{Number: 19000, Value: "on dokuz bin"},
		{Number: 18000, Value: "on sekiz bin"},
		{Number: 17000, Value: "on yedi bin"},
		{Number: 16000, Value: "on altı bin"},
		{Number: 15000, Value: "on beş bin"},
		{Number: 14000, Value: "on dört bin"},
		{Number: 13000, Value: "on üç bin"},
		{Number: 12000, Value: "on iki bin"},
		{Number: 11000, Value: "on bir bin"},
		{Number: 10000, Value: "on bin"},
		{Number: 9000, Value: "dokuz bin"},
		{Number: 8000, Value: "sekiz bin"},
		{Number: 7000, Value: "yedi bin"},
		{Number: 6000, Value: "altı bin"},
		{Number: 5000, Value: "beş bin"},
		{Number: 4000, Value: "dört bin"},
		{Number: 3000, Value: "üç bin"},
		{Number: 2000, Value: "iki bin"},
		{Number: 1000, Value: "bin"},
		{Number: 900, Value: "dokuz yüz"},
		{Number: 800, Value: "sekiz yüz"},
		{Number: 700, Value: "yedi yüz"},
		{Number: 600, Value: "altı yüz"},
		{Number: 500, Value: "beş yüz"},
		{Number: 400, Value: "dört yüz"},
		{Number: 300, Value: "üç yüz"},
		{Number: 200, Value: "iki yüz"},
		{Number: 100, Value: "yüz"},
		{Number: 99, Value: "doksan dokuz"},
		{Number: 98, Value: "doksan sekiz"},
		{Number: 97, Value: "doksan yedi"},
		{Number: 96, Value: "doksan altı"},
		{Number: 95, Value: "doksan beş"},
		{Number: 94, Value: "doksan dört"},
		{Number: 93, Value: "doksan üç"},
		{Number: 92, Value: "doksan iki"},
		{Number: 91, Value: "doksan bir"},
		{Number: 90, Value: "doksan"},
		{Number: 89, Value: "seksen dokuz"},
		{Number: 88, Value: "seksen sekiz"},
		{Number: 87, Value: "seksen yedi"},
		{Number: 86, Value: "seksen altı"},
		{Number: 85, Value: "seksen beş"},
		{Number: 84, Value: "seksen dört"},
		{Number: 83, Value: "seksen üç"},
		{Number: 82, Value: "seksen iki"},
		{Number: 81, Value: "seksen bir"},
		{Number: 80, Value: "seksen"},
		{Number: 79, Value: "yetmiş dokuz"},
		{Number: 78, Value: "yetmiş sekiz"},
		{Number: 77, Value: "yetmiş yedi"},
		{Number: 76, Value: "yetmiş altı"},
		{Number: 75, Value: "yetmiş beş"},
		{Number: 74, Value: "yetmiş dört"},
		{Number: 73, Value: "yetmiş üç"},
		{Number: 72, Value: "yetmiş iki"},
		{Number: 71, Value: "yetmiş bir"},
		{Number: 70, Value: "yetmiş"},
		{Number: 69, Value: "altmış dokuz"},
		{Number: 68, Value: "altmış sekiz"},
		{Number: 67, Value: "altmış yedi"},
		{Number: 66, Value: "altmış altı"},
		{Number: 65, Value: "altmış beş"},
		{Number: 64, Value: "altmış dört"},
		{Number: 63, Value: "altmış üç"},
		{Number: 62, Value: "altmış iki"},
		{Number: 61, Value: "altmış bir"},
		{Number: 60, Value: "altmış"},
		{Number: 59, Value: "elli dokuz"},
		{Number: 58, Value: "elli sekiz"},
		{Number: 57, Value: "elli yedi"},
		{Number: 56, Value: "elli altı"},
		{Number: 55, Value: "elli beş"},
		{Number: 54, Value: "elli dört"},
		{Number: 53, Value: "elli üç"},
		{Number: 52, Value: "elli iki"},
		{Number: 51, Value: "elli bir"},
		{Number: 50, Value: "elli"},
		{Number: 49, Value: "kırk dokuz"},
		{Number: 48, Value: "kırk sekiz"},
		{Number: 47, Value: "kırk yedi"},
		{Number: 46, Value: "kırk altı"},
		{Number: 45, Value: "kırk beş"},
		{Number: 44, Value: "kırk dört"},
		{Number: 43, Value: "kırk üç"},
		{Number: 42, Value: "kırk iki"},
		{Number: 41, Value: "kırk bir"},
		{Number: 40, Value: "kırk"},
		{Number: 39, Value: "otuz dokuz"},
		{Number: 38, Value: "otuz sekiz"},
		{Number: 37, Value: "otuz yedi"},
		{Number: 36, Value: "otuz altı"},
		{Number: 35, Value: "otuz beş"},
		{Number: 34, Value: "otuz dört"},
		{Number: 33, Value: "otuz üç"},
		{Number: 32, Value: "otuz iki"},
		{Number: 31, Value: "otuz bir"},
		{Number: 30, Value: "otuz"},
		{Number: 29, Value: "yirmi dokuz"},
		{Number: 28, Value: "yirmi sekiz"},
		{Number: 27, Value: "yirmi yedi"},
		{Number: 26, Value: "yirmi altı"},
		{Number: 25, Value: "yirmi beş"},
		{Number: 24, Value: "yirmi dört"},
		{Number: 23, Value: "yirmi üç"},
		{Number: 22, Value: "yirmi iki"},
		{Number: 21, Value: "yirmi bir"},
		{Number: 20, Value: "yirmi"},
		{Number: 19, Value: "on dokuz"},
		{Number: 18, Value: "on sekiz"},
		{Number: 17, Value: "on yedi"},
		{Number: 16, Value: "on altı"},
		{Number: 15, Value: "on beş"},
		{Number: 14, Value: "on dört"},
		{Number: 13, Value: "on üç"},
		{Number: 12, Value: "on iki"},
		{Number: 11, Value: "on bir"},
		{Number: 10, Value: "on"},
		{Number: 9, Value: "dokuz"},
		{Number: 8, Value: "sekiz"},
		{Number: 7, Value: "yedi"},
		{Number: 6, Value: "altı"},
		{Number: 5, Value: "beş"},
		{Number: 4, Value: "dört"},
		{Number: 3, Value: "üç"},
		{Number: 2, Value: "iki"},
		{Number: 1, Value: "bir"},
		{Number: 0, Value: "sıfır"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Yüz"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "birinci", Suffix: ".", Masculine: "birinci", Feminine: "birinci", Neuter: "birinci"},
		{Number: 2, Word: "ikinci", Suffix: ".", Masculine: "ikinci", Feminine: "ikinci", Neuter: "ikinci"},
		{Number: 3, Word: "üçüncü", Suffix: ".", Masculine: "üçüncü", Feminine: "üçüncü", Neuter: "üçüncü"},
		{Number: 4, Word: "dördüncü", Suffix: ".", Masculine: "dördüncü", Feminine: "dördüncü", Neuter: "dördüncü"},
		{Number: 5, Word: "beşinci", Suffix: ".", Masculine: "beşinci", Feminine: "beşinci", Neuter: "beşinci"},
		{Number: 6, Word: "altıncı", Suffix: ".", Masculine: "altıncı", Feminine: "altıncı", Neuter: "altıncı"},
		{Number: 7, Word: "yedinci", Suffix: ".", Masculine: "yedinci", Feminine: "yedinci", Neuter: "yedinci"},
		{Number: 8, Word: "sekizinci", Suffix: ".", Masculine: "sekizinci", Feminine: "sekizinci", Neuter: "sekizinci"},
		{Number: 9, Word: "dokuzuncu", Suffix: ".", Masculine: "dokuzuncu", Feminine: "dokuzuncu", Neuter: "dokuzuncu"},
		{Number: 10, Word: "onuncu", Suffix: ".", Masculine: "onuncu", Feminine: "onuncu", Neuter: "onuncu"},
		{Number: 11, Word: "on birinci", Suffix: ".", Masculine: "on birinci", Feminine: "on birinci", Neuter: "on birinci"},
		{Number: 12, Word: "on ikinci", Suffix: ".", Masculine: "on ikinci", Feminine: "on ikinci", Neuter: "on ikinci"},
		{Number: 20, Word: "yirminci", Suffix: ".", Masculine: "yirminci", Feminine: "yirminci", Neuter: "yirminci"},
		{Number: 21, Word: "yirmi birinci", Suffix: ".", Masculine: "yirmi birinci", Feminine: "yirmi birinci", Neuter: "yirmi birinci"},
		{Number: 30, Word: "otuzuncu", Suffix: ".", Masculine: "otuzuncu", Feminine: "otuzuncu", Neuter: "otuzuncu"},
		{Number: 50, Word: "ellinci", Suffix: ".", Masculine: "ellinci", Feminine: "ellinci", Neuter: "ellinci"},
		{Number: 100, Word: "yüzüncü", Suffix: ".", Masculine: "yüzüncü", Feminine: "yüzüncü", Neuter: "yüzüncü"},
		{Number: 1000, Word: "bininci", Suffix: ".", Masculine: "bininci", Feminine: "bininci", Neuter: "bininci"},
	},
	LocaleFormatter: &TurkishFormatter{},
}

// TurkishFormatter handles Turkish (Turkey) formatting
type TurkishFormatter struct{}

func (f *TurkishFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *TurkishFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *TurkishFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *TurkishFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *TurkishFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *TurkishFormatter) ChopDecimal(value decimal.Decimal, decimalPlaces int) decimal.Decimal {
	return value.Truncate(int32(decimalPlaces))
}

func (f *TurkishFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *TurkishFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
