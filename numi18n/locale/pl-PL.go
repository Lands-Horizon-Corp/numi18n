package locale

import (
	"github.com/shopspring/decimal"
)

// PLPLLocale represents the Polish (Poland) locale
var PLPLLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Złoty",
		Plural:   "Złote",
		Singular: "Złoty",
		Symbol:   "zł",
		FractionUnit: FractionUnit{
			Name:     "Grosz",
			Plural:   "Grosze",
			Singular: "Grosz",
			Symbol:   "gr",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Poland",
		Currency:       "PLN",
		ISO3166Alpha2:  "PL",
		ISO3166Alpha3:  "POL",
		ISO3166Numeric: "616",
		Locale:         "pl-PL",
		Timezone:       []string{"Europe/Warsaw"},
		Language:       "pl",
		Emoji:          "🇵🇱",
		PhoneCode:      "+48",
		Domain:         ".pl",
	},
	Texts: Texts{
		And:   "i",
		Minus: "minus",
		Only:  "tylko",
		Point: "przecinek",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "biliard"},
		{Number: 1000000000000, Value: "bilion"},
		{Number: 1000000000, Value: "miliard"},
		{Number: 1000000, Value: "milion"},
		{Number: 100000, Value: "sto tysięcy"},
		{Number: 90000, Value: "dziewięćdziesiąt tysięcy"},
		{Number: 80000, Value: "osiemdziesiąt tysięcy"},
		{Number: 70000, Value: "siedemdziesiąt tysięcy"},
		{Number: 60000, Value: "sześćdziesiąt tysięcy"},
		{Number: 50000, Value: "pięćdziesiąt tysięcy"},
		{Number: 40000, Value: "czterdzieści tysięcy"},
		{Number: 30000, Value: "trzydzieści tysięcy"},
		{Number: 20000, Value: "dwadzieścia tysięcy"},
		{Number: 19000, Value: "dziewiętnaście tysięcy"},
		{Number: 18000, Value: "osiemnaście tysięcy"},
		{Number: 17000, Value: "siedemnaście tysięcy"},
		{Number: 16000, Value: "szesnaście tysięcy"},
		{Number: 15000, Value: "piętnaście tysięcy"},
		{Number: 14000, Value: "czternaście tysięcy"},
		{Number: 13000, Value: "trzynaście tysięcy"},
		{Number: 12000, Value: "dwanaście tysięcy"},
		{Number: 11000, Value: "jedenaście tysięcy"},
		{Number: 10000, Value: "dziesięć tysięcy"},
		{Number: 9000, Value: "dziewięć tysięcy"},
		{Number: 8000, Value: "osiem tysięcy"},
		{Number: 7000, Value: "siedem tysięcy"},
		{Number: 6000, Value: "sześć tysięcy"},
		{Number: 5000, Value: "pięć tysięcy"},
		{Number: 4000, Value: "cztery tysiące"},
		{Number: 3000, Value: "trzy tysiące"},
		{Number: 2000, Value: "dwa tysiące"},
		{Number: 1000, Value: "tysiąc"},
		{Number: 900, Value: "dziewięćset"},
		{Number: 800, Value: "osiemset"},
		{Number: 700, Value: "siedemset"},
		{Number: 600, Value: "sześćset"},
		{Number: 500, Value: "pięćset"},
		{Number: 400, Value: "czterysta"},
		{Number: 300, Value: "trzysta"},
		{Number: 200, Value: "dwieście"},
		{Number: 100, Value: "sto"},
		{Number: 99, Value: "dziewięćdziesiąt dziewięć"},
		{Number: 98, Value: "dziewięćdziesiąt osiem"},
		{Number: 97, Value: "dziewięćdziesiąt siedem"},
		{Number: 96, Value: "dziewięćdziesiąt sześć"},
		{Number: 95, Value: "dziewięćdziesiąt pięć"},
		{Number: 94, Value: "dziewięćdziesiąt cztery"},
		{Number: 93, Value: "dziewięćdziesiąt trzy"},
		{Number: 92, Value: "dziewięćdziesiąt dwa"},
		{Number: 91, Value: "dziewięćdziesiąt jeden"},
		{Number: 90, Value: "dziewięćdziesiąt"},
		{Number: 89, Value: "osiemdziesiąt dziewięć"},
		{Number: 88, Value: "osiemdziesiąt osiem"},
		{Number: 87, Value: "osiemdziesiąt siedem"},
		{Number: 86, Value: "osiemdziesiąt sześć"},
		{Number: 85, Value: "osiemdziesiąt pięć"},
		{Number: 84, Value: "osiemdziesiąt cztery"},
		{Number: 83, Value: "osiemdziesiąt trzy"},
		{Number: 82, Value: "osiemdziesiąt dwa"},
		{Number: 81, Value: "osiemdziesiąt jeden"},
		{Number: 80, Value: "osiemdziesiąt"},
		{Number: 79, Value: "siedemdziesiąt dziewięć"},
		{Number: 78, Value: "siedemdziesiąt osiem"},
		{Number: 77, Value: "siedemdziesiąt siedem"},
		{Number: 76, Value: "siedemdziesiąt sześć"},
		{Number: 75, Value: "siedemdziesiąt pięć"},
		{Number: 74, Value: "siedemdziesiąt cztery"},
		{Number: 73, Value: "siedemdziesiąt trzy"},
		{Number: 72, Value: "siedemdziesiąt dwa"},
		{Number: 71, Value: "siedemdziesiąt jeden"},
		{Number: 70, Value: "siedemdziesiąt"},
		{Number: 69, Value: "sześćdziesiąt dziewięć"},
		{Number: 68, Value: "sześćdziesiąt osiem"},
		{Number: 67, Value: "sześćdziesiąt siedem"},
		{Number: 66, Value: "sześćdziesiąt sześć"},
		{Number: 65, Value: "sześćdziesiąt pięć"},
		{Number: 64, Value: "sześćdziesiąt cztery"},
		{Number: 63, Value: "sześćdziesiąt trzy"},
		{Number: 62, Value: "sześćdziesiąt dwa"},
		{Number: 61, Value: "sześćdziesiąt jeden"},
		{Number: 60, Value: "sześćdziesiąt"},
		{Number: 59, Value: "pięćdziesiąt dziewięć"},
		{Number: 58, Value: "pięćdziesiąt osiem"},
		{Number: 57, Value: "pięćdziesiąt siedem"},
		{Number: 56, Value: "pięćdziesiąt sześć"},
		{Number: 55, Value: "pięćdziesiąt pięć"},
		{Number: 54, Value: "pięćdziesiąt cztery"},
		{Number: 53, Value: "pięćdziesiąt trzy"},
		{Number: 52, Value: "pięćdziesiąt dwa"},
		{Number: 51, Value: "pięćdziesiąt jeden"},
		{Number: 50, Value: "pięćdziesiąt"},
		{Number: 49, Value: "czterdzieści dziewięć"},
		{Number: 48, Value: "czterdzieści osiem"},
		{Number: 47, Value: "czterdzieści siedem"},
		{Number: 46, Value: "czterdzieści sześć"},
		{Number: 45, Value: "czterdzieści pięć"},
		{Number: 44, Value: "czterdzieści cztery"},
		{Number: 43, Value: "czterdzieści trzy"},
		{Number: 42, Value: "czterdzieści dwa"},
		{Number: 41, Value: "czterdzieści jeden"},
		{Number: 40, Value: "czterdzieści"},
		{Number: 39, Value: "trzydzieści dziewięć"},
		{Number: 38, Value: "trzydzieści osiem"},
		{Number: 37, Value: "trzydzieści siedem"},
		{Number: 36, Value: "trzydzieści sześć"},
		{Number: 35, Value: "trzydzieści pięć"},
		{Number: 34, Value: "trzydzieści cztery"},
		{Number: 33, Value: "trzydzieści trzy"},
		{Number: 32, Value: "trzydzieści dwa"},
		{Number: 31, Value: "trzydzieści jeden"},
		{Number: 30, Value: "trzydzieści"},
		{Number: 29, Value: "dwadzieścia dziewięć"},
		{Number: 28, Value: "dwadzieścia osiem"},
		{Number: 27, Value: "dwadzieścia siedem"},
		{Number: 26, Value: "dwadzieścia sześć"},
		{Number: 25, Value: "dwadzieścia pięć"},
		{Number: 24, Value: "dwadzieścia cztery"},
		{Number: 23, Value: "dwadzieścia trzy"},
		{Number: 22, Value: "dwadzieścia dwa"},
		{Number: 21, Value: "dwadzieścia jeden"},
		{Number: 20, Value: "dwadzieścia"},
		{Number: 19, Value: "dziewiętnaście"},
		{Number: 18, Value: "osiemnaście"},
		{Number: 17, Value: "siedemnaście"},
		{Number: 16, Value: "szesnaście"},
		{Number: 15, Value: "piętnaście"},
		{Number: 14, Value: "czternaście"},
		{Number: 13, Value: "trzynaście"},
		{Number: 12, Value: "dwanaście"},
		{Number: 11, Value: "jedenaście"},
		{Number: 10, Value: "dziesięć"},
		{Number: 9, Value: "dziewięć"},
		{Number: 8, Value: "osiem"},
		{Number: 7, Value: "siedem"},
		{Number: 6, Value: "sześć"},
		{Number: 5, Value: "pięć"},
		{Number: 4, Value: "cztery"},
		{Number: 3, Value: "trzy"},
		{Number: 2, Value: "dwa"},
		{Number: 1, Value: "jeden"},
		{Number: 0, Value: "zero"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Sto"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "pierwszy", Suffix: ".", Masculine: "pierwszy", Feminine: "pierwsza", Neuter: "pierwsze"},
		{Number: 2, Word: "drugi", Suffix: ".", Masculine: "drugi", Feminine: "druga", Neuter: "drugie"},
		{Number: 3, Word: "trzeci", Suffix: ".", Masculine: "trzeci", Feminine: "trzecia", Neuter: "trzecie"},
		{Number: 4, Word: "czwarty", Suffix: ".", Masculine: "czwarty", Feminine: "czwarta", Neuter: "czwarte"},
		{Number: 5, Word: "piąty", Suffix: ".", Masculine: "piąty", Feminine: "piąta", Neuter: "piąte"},
		{Number: 6, Word: "szósty", Suffix: ".", Masculine: "szósty", Feminine: "szósta", Neuter: "szóste"},
		{Number: 7, Word: "siódmy", Suffix: ".", Masculine: "siódmy", Feminine: "siódma", Neuter: "siódme"},
		{Number: 8, Word: "ósmy", Suffix: ".", Masculine: "ósmy", Feminine: "ósma", Neuter: "ósme"},
		{Number: 9, Word: "dziewiąty", Suffix: ".", Masculine: "dziewiąty", Feminine: "dziewiąta", Neuter: "dziewiąte"},
		{Number: 10, Word: "dziesiąty", Suffix: ".", Masculine: "dziesiąty", Feminine: "dziesiąta", Neuter: "dziesiąte"},
		{Number: 11, Word: "jedenasty", Suffix: ".", Masculine: "jedenasty", Feminine: "jedenasta", Neuter: "jedenaste"},
		{Number: 12, Word: "dwunasty", Suffix: ".", Masculine: "dwunasty", Feminine: "dwunasta", Neuter: "dwunaste"},
		{Number: 13, Word: "trzynasty", Suffix: ".", Masculine: "trzynasty", Feminine: "trzynasta", Neuter: "trzynaste"},
		{Number: 20, Word: "dwudziesty", Suffix: ".", Masculine: "dwudziesty", Feminine: "dwudziesta", Neuter: "dwudzieste"},
		{Number: 21, Word: "dwudziesty pierwszy", Suffix: ".", Masculine: "dwudziesty pierwszy", Feminine: "dwudziesta pierwsza", Neuter: "dwudzieste pierwsze"},
		{Number: 30, Word: "trzydziesty", Suffix: ".", Masculine: "trzydziesty", Feminine: "trzydziesta", Neuter: "trzydzieste"},
		{Number: 50, Word: "pięćdziesiąty", Suffix: ".", Masculine: "pięćdziesiąty", Feminine: "pięćdziesiąta", Neuter: "pięćdziesiąte"},
		{Number: 100, Word: "setny", Suffix: ".", Masculine: "setny", Feminine: "setna", Neuter: "setne"},
		{Number: 1000, Word: "tysięczny", Suffix: ".", Masculine: "tysięczny", Feminine: "tysięczna", Neuter: "tysięczne"},
	},
	LocaleFormatter: &PolishFormatter{},
}

// PolishFormatter handles Polish formatting
type PolishFormatter struct{}

func (f *PolishFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *PolishFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *PolishFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *PolishFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *PolishFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *PolishFormatter) ChopDecimal(value decimal.Decimal, precision int) decimal.Decimal {
	return value.Truncate(int32(precision))
}

func (f *PolishFormatter) FormatDecimalNumber(amount float64) string {
	return FormatPolishDecimal(amount)
}

func (f *PolishFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Polish conventions (comma separators, period decimal, prefix symbol)
	return FormatPolishCurrency(amount, currencySymbol)
}
