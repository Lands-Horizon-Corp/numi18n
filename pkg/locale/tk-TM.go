package locale

import "github.com/shopspring/decimal"

// TKTMLocale represents the Turkmen (Turkmenistan) locale
var TKTMLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Turkmen Manat",
		Plural:   "Manat",
		Singular: "Manat",
		Symbol:   "TMT",
		FractionUnit: FractionUnit{
			Name:     "Tenge",
			Plural:   "Tenge",
			Singular: "Tenge",
			Symbol:   "t",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Turkmenistan",
		Currency:       "TMT",
		ISO3166Alpha2:  "TM",
		ISO3166Alpha3:  "TKM",
		ISO3166Numeric: "795",
		Locale:         "tk-TM",
		Timezone:       []string{"Asia/Ashgabat"},
		Language:       "tk",
	},
	Texts: Texts{
		And:   "we",
		Minus: "minus",
		Only:  "diňe",
		Point: "nokat",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "bir trilliard"},
		{Number: 1000000000000, Value: "bir trillion"},
		{Number: 1000000000, Value: "bir milliard"},
		{Number: 1000000, Value: "bir million"},
		{Number: 100000, Value: "ýüz müň"},
		{Number: 90000, Value: "togsan müň"},
		{Number: 80000, Value: "segsen müň"},
		{Number: 70000, Value: "ýetmiş müň"},
		{Number: 60000, Value: "altmyş müň"},
		{Number: 50000, Value: "elli müň"},
		{Number: 40000, Value: "kyrk müň"},
		{Number: 30000, Value: "otuz müň"},
		{Number: 20000, Value: "ýigrimi müň"},
		{Number: 19000, Value: "on dokuz müň"},
		{Number: 18000, Value: "on sekiz müň"},
		{Number: 17000, Value: "on ýedi müň"},
		{Number: 16000, Value: "on alty müň"},
		{Number: 15000, Value: "on bäş müň"},
		{Number: 14000, Value: "on dört müň"},
		{Number: 13000, Value: "on üç müň"},
		{Number: 12000, Value: "on iki müň"},
		{Number: 11000, Value: "on bir müň"},
		{Number: 10000, Value: "on müň"},
		{Number: 9000, Value: "dokuz müň"},
		{Number: 8000, Value: "sekiz müň"},
		{Number: 7000, Value: "ýedi müň"},
		{Number: 6000, Value: "alty müň"},
		{Number: 5000, Value: "bäş müň"},
		{Number: 4000, Value: "dört müň"},
		{Number: 3000, Value: "üç müň"},
		{Number: 2000, Value: "iki müň"},
		{Number: 1000, Value: "bir müň"},
		{Number: 900, Value: "dokuz ýüz"},
		{Number: 800, Value: "sekiz ýüz"},
		{Number: 700, Value: "ýedi ýüz"},
		{Number: 600, Value: "alty ýüz"},
		{Number: 500, Value: "bäş ýüz"},
		{Number: 400, Value: "dört ýüz"},
		{Number: 300, Value: "üç ýüz"},
		{Number: 200, Value: "iki ýüz"},
		{Number: 100, Value: "bir ýüz"},
		{Number: 99, Value: "togsan dokuz"},
		{Number: 98, Value: "togsan sekiz"},
		{Number: 97, Value: "togsan ýedi"},
		{Number: 96, Value: "togsan alty"},
		{Number: 95, Value: "togsan bäş"},
		{Number: 94, Value: "togsan dört"},
		{Number: 93, Value: "togsan üç"},
		{Number: 92, Value: "togsan iki"},
		{Number: 91, Value: "togsan bir"},
		{Number: 90, Value: "togsan"},
		{Number: 89, Value: "segsen dokuz"},
		{Number: 88, Value: "segsen sekiz"},
		{Number: 87, Value: "segsen ýedi"},
		{Number: 86, Value: "segsen alty"},
		{Number: 85, Value: "segsen bäş"},
		{Number: 84, Value: "segsen dört"},
		{Number: 83, Value: "segsen üç"},
		{Number: 82, Value: "segsen iki"},
		{Number: 81, Value: "segsen bir"},
		{Number: 80, Value: "segsen"},
		{Number: 79, Value: "ýetmiş dokuz"},
		{Number: 78, Value: "ýetmiş sekiz"},
		{Number: 77, Value: "ýetmiş ýedi"},
		{Number: 76, Value: "ýetmiş alty"},
		{Number: 75, Value: "ýetmiş bäş"},
		{Number: 74, Value: "ýetmiş dört"},
		{Number: 73, Value: "ýetmiş üç"},
		{Number: 72, Value: "ýetmiş iki"},
		{Number: 71, Value: "ýetmiş bir"},
		{Number: 70, Value: "ýetmiş"},
		{Number: 69, Value: "altmyş dokuz"},
		{Number: 68, Value: "altmyş sekiz"},
		{Number: 67, Value: "altmyş ýedi"},
		{Number: 66, Value: "altmyş alty"},
		{Number: 65, Value: "altmyş bäş"},
		{Number: 64, Value: "altmyş dört"},
		{Number: 63, Value: "altmyş üç"},
		{Number: 62, Value: "altmyş iki"},
		{Number: 61, Value: "altmyş bir"},
		{Number: 60, Value: "altmyş"},
		{Number: 59, Value: "elli dokuz"},
		{Number: 58, Value: "elli sekiz"},
		{Number: 57, Value: "elli ýedi"},
		{Number: 56, Value: "elli alty"},
		{Number: 55, Value: "elli bäş"},
		{Number: 54, Value: "elli dört"},
		{Number: 53, Value: "elli üç"},
		{Number: 52, Value: "elli iki"},
		{Number: 51, Value: "elli bir"},
		{Number: 50, Value: "elli"},
		{Number: 49, Value: "kyrk dokuz"},
		{Number: 48, Value: "kyrk sekiz"},
		{Number: 47, Value: "kyrk ýedi"},
		{Number: 46, Value: "kyrk alty"},
		{Number: 45, Value: "kyrk bäş"},
		{Number: 44, Value: "kyrk dört"},
		{Number: 43, Value: "kyrk üç"},
		{Number: 42, Value: "kyrk iki"},
		{Number: 41, Value: "kyrk bir"},
		{Number: 40, Value: "kyrk"},
		{Number: 39, Value: "otuz dokuz"},
		{Number: 38, Value: "otuz sekiz"},
		{Number: 37, Value: "otuz ýedi"},
		{Number: 36, Value: "otuz alty"},
		{Number: 35, Value: "otuz bäş"},
		{Number: 34, Value: "otuz dört"},
		{Number: 33, Value: "otuz üç"},
		{Number: 32, Value: "otuz iki"},
		{Number: 31, Value: "otuz bir"},
		{Number: 30, Value: "otuz"},
		{Number: 29, Value: "ýigrimi dokuz"},
		{Number: 28, Value: "ýigrimi sekiz"},
		{Number: 27, Value: "ýigrimi ýedi"},
		{Number: 26, Value: "ýigrimi alty"},
		{Number: 25, Value: "ýigrimi bäş"},
		{Number: 24, Value: "ýigrimi dört"},
		{Number: 23, Value: "ýigrimi üç"},
		{Number: 22, Value: "ýigrimi iki"},
		{Number: 21, Value: "ýigrimi bir"},
		{Number: 20, Value: "ýigrimi"},
		{Number: 19, Value: "on dokuz"},
		{Number: 18, Value: "on sekiz"},
		{Number: 17, Value: "on ýedi"},
		{Number: 16, Value: "on alty"},
		{Number: 15, Value: "on bäş"},
		{Number: 14, Value: "on dört"},
		{Number: 13, Value: "on üç"},
		{Number: 12, Value: "on iki"},
		{Number: 11, Value: "on bir"},
		{Number: 10, Value: "on"},
		{Number: 9, Value: "dokuz"},
		{Number: 8, Value: "sekiz"},
		{Number: 7, Value: "ýedi"},
		{Number: 6, Value: "alty"},
		{Number: 5, Value: "bäş"},
		{Number: 4, Value: "dört"},
		{Number: 3, Value: "üç"},
		{Number: 2, Value: "iki"},
		{Number: 1, Value: "bir"},
		{Number: 0, Value: "nol"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Bir ýüz"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "birinji", Suffix: "-nji", Masculine: "birinji", Feminine: "birinji", Neuter: "birinji"},
		{Number: 2, Word: "ikinji", Suffix: "-nji", Masculine: "ikinji", Feminine: "ikinji", Neuter: "ikinji"},
		{Number: 3, Word: "üçünji", Suffix: "-nji", Masculine: "üçünji", Feminine: "üçünji", Neuter: "üçünji"},
		{Number: 4, Word: "dördünji", Suffix: "-nji", Masculine: "dördünji", Feminine: "dördünji", Neuter: "dördünji"},
		{Number: 5, Word: "bäşinji", Suffix: "-nji", Masculine: "bäşinji", Feminine: "bäşinji", Neuter: "bäşinji"},
		{Number: 6, Word: "altynjy", Suffix: "-nji", Masculine: "altynjy", Feminine: "altynjy", Neuter: "altynjy"},
		{Number: 7, Word: "ýediji", Suffix: "-nji", Masculine: "ýediji", Feminine: "ýediji", Neuter: "ýediji"},
		{Number: 8, Word: "sekizinji", Suffix: "-nji", Masculine: "sekizinji", Feminine: "sekizinji", Neuter: "sekizinji"},
		{Number: 9, Word: "dokuzynjy", Suffix: "-nji", Masculine: "dokuzynjy", Feminine: "dokuzynjy", Neuter: "dokuzynjy"},
		{Number: 10, Word: "onunji", Suffix: "-nji", Masculine: "onunji", Feminine: "onunji", Neuter: "onunji"},
		{Number: 11, Word: "on birinji", Suffix: "-nji", Masculine: "on birinji", Feminine: "on birinji", Neuter: "on birinji"},
		{Number: 12, Word: "on ikinji", Suffix: "-nji", Masculine: "on ikinji", Feminine: "on ikinji", Neuter: "on ikinji"},
		{Number: 20, Word: "ýigriminji", Suffix: "-nji", Masculine: "ýigriminji", Feminine: "ýigriminji", Neuter: "ýigriminji"},
		{Number: 21, Word: "ýigrimi birinji", Suffix: "-nji", Masculine: "ýigrimi birinji", Feminine: "ýigrimi birinji", Neuter: "ýigrimi birinji"},
		{Number: 30, Word: "otuzunji", Suffix: "-nji", Masculine: "otuzunji", Feminine: "otuzunji", Neuter: "otuzunji"},
		{Number: 50, Word: "elliji", Suffix: "-nji", Masculine: "elliji", Feminine: "elliji", Neuter: "elliji"},
		{Number: 100, Word: "ýüzünji", Suffix: "-nji", Masculine: "ýüzünji", Feminine: "ýüzünji", Neuter: "ýüzünji"},
		{Number: 1000, Word: "müňünji", Suffix: "-nji", Masculine: "müňünji", Feminine: "müňünji", Neuter: "müňünji"},
	},
	LocaleFormatter: &TurkmenFormatter{},
}

// TurkmenFormatter handles Turkmen (Turkmenistan) formatting
type TurkmenFormatter struct{}

func (f *TurkmenFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *TurkmenFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *TurkmenFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *TurkmenFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *TurkmenFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *TurkmenFormatter) ChopDecimal(value decimal.Decimal, decimalPlaces int) decimal.Decimal {
	return value.Truncate(int32(decimalPlaces))
}
