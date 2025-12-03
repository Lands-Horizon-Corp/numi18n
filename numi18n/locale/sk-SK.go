package locale

import (
	"github.com/shopspring/decimal"
)

// SKSKLocale represents the Slovak (Slovakia) locale
var SKSKLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Euro",
		Plural:   "Eurá",
		Singular: "Euro",
		Symbol:   "€",
		FractionUnit: FractionUnit{
			Name:     "Cent",
			Plural:   "Centy",
			Singular: "Cent",
			Symbol:   "c",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Slovakia",
		Currency:       "EUR",
		ISO3166Alpha2:  "SK",
		ISO3166Alpha3:  "SVK",
		ISO3166Numeric: "703",
		Locale:         "sk-SK",
		Timezone:       []string{"Europe/Bratislava"},
		Language:       "sk",
		Emoji:          "🇸🇰",
		PhoneCode:      "+421",
		Domain:         ".sk",
	},
	Texts: Texts{
		And:   "a",
		Minus: "mínus",
		Only:  "len",
		Point: "čiarka",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "biliarda"},
		{Number: 1000000000000, Value: "bilión"},
		{Number: 1000000000, Value: "miliarda"},
		{Number: 1000000, Value: "milión"},
		{Number: 100000, Value: "sto tisíc"},
		{Number: 90000, Value: "deväťdesiat tisíc"},
		{Number: 80000, Value: "osemdesiat tisíc"},
		{Number: 70000, Value: "sedemdesiat tisíc"},
		{Number: 60000, Value: "šesťdesiat tisíc"},
		{Number: 50000, Value: "päťdesiat tisíc"},
		{Number: 40000, Value: "štyridsať tisíc"},
		{Number: 30000, Value: "tridsať tisíc"},
		{Number: 20000, Value: "dvadsať tisíc"},
		{Number: 19000, Value: "devätnásť tisíc"},
		{Number: 18000, Value: "osemnásť tisíc"},
		{Number: 17000, Value: "sedemnásť tisíc"},
		{Number: 16000, Value: "šestnásť tisíc"},
		{Number: 15000, Value: "pätnásť tisíc"},
		{Number: 14000, Value: "štrnásť tisíc"},
		{Number: 13000, Value: "trinásť tisíc"},
		{Number: 12000, Value: "dvanásť tisíc"},
		{Number: 11000, Value: "jedenásť tisíc"},
		{Number: 10000, Value: "desať tisíc"},
		{Number: 9000, Value: "deväť tisíc"},
		{Number: 8000, Value: "osem tisíc"},
		{Number: 7000, Value: "sedem tisíc"},
		{Number: 6000, Value: "šesť tisíc"},
		{Number: 5000, Value: "päť tisíc"},
		{Number: 4000, Value: "štyri tisíc"},
		{Number: 3000, Value: "tri tisíc"},
		{Number: 2000, Value: "dva tisíc"},
		{Number: 1000, Value: "tisíc"},
		{Number: 900, Value: "deväťsto"},
		{Number: 800, Value: "osemsto"},
		{Number: 700, Value: "sedemsto"},
		{Number: 600, Value: "šesťsto"},
		{Number: 500, Value: "päťsto"},
		{Number: 400, Value: "štyristo"},
		{Number: 300, Value: "tristo"},
		{Number: 200, Value: "dvesto"},
		{Number: 100, Value: "sto"},
		{Number: 99, Value: "deväťdesiat deväť"},
		{Number: 98, Value: "deväťdesiat osem"},
		{Number: 97, Value: "deväťdesiat sedem"},
		{Number: 96, Value: "deväťdesiat šesť"},
		{Number: 95, Value: "deväťdesiat päť"},
		{Number: 94, Value: "deväťdesiat štyri"},
		{Number: 93, Value: "deväťdesiat tri"},
		{Number: 92, Value: "deväťdesiat dva"},
		{Number: 91, Value: "deväťdesiat jeden"},
		{Number: 90, Value: "deväťdesiat"},
		{Number: 89, Value: "osemdesiat deväť"},
		{Number: 88, Value: "osemdesiat osem"},
		{Number: 87, Value: "osemdesiat sedem"},
		{Number: 86, Value: "osemdesiat šesť"},
		{Number: 85, Value: "osemdesiat päť"},
		{Number: 84, Value: "osemdesiat štyri"},
		{Number: 83, Value: "osemdesiat tri"},
		{Number: 82, Value: "osemdesiat dva"},
		{Number: 81, Value: "osemdesiat jeden"},
		{Number: 80, Value: "osemdesiat"},
		{Number: 79, Value: "sedemdesiat deväť"},
		{Number: 78, Value: "sedemdesiat osem"},
		{Number: 77, Value: "sedemdesiat sedem"},
		{Number: 76, Value: "sedemdesiat šesť"},
		{Number: 75, Value: "sedemdesiat päť"},
		{Number: 74, Value: "sedemdesiat štyri"},
		{Number: 73, Value: "sedemdesiat tri"},
		{Number: 72, Value: "sedemdesiat dva"},
		{Number: 71, Value: "sedemdesiat jeden"},
		{Number: 70, Value: "sedemdesiat"},
		{Number: 69, Value: "šesťdesiat deväť"},
		{Number: 68, Value: "šesťdesiat osem"},
		{Number: 67, Value: "šesťdesiat sedem"},
		{Number: 66, Value: "šesťdesiat šesť"},
		{Number: 65, Value: "šesťdesiat päť"},
		{Number: 64, Value: "šesťdesiat štyri"},
		{Number: 63, Value: "šesťdesiat tri"},
		{Number: 62, Value: "šesťdesiat dva"},
		{Number: 61, Value: "šesťdesiat jeden"},
		{Number: 60, Value: "šesťdesiat"},
		{Number: 59, Value: "päťdesiat deväť"},
		{Number: 58, Value: "päťdesiat osem"},
		{Number: 57, Value: "päťdesiat sedem"},
		{Number: 56, Value: "päťdesiat šesť"},
		{Number: 55, Value: "päťdesiat päť"},
		{Number: 54, Value: "päťdesiat štyri"},
		{Number: 53, Value: "päťdesiat tri"},
		{Number: 52, Value: "päťdesiat dva"},
		{Number: 51, Value: "päťdesiat jeden"},
		{Number: 50, Value: "päťdesiat"},
		{Number: 49, Value: "štyridsať deväť"},
		{Number: 48, Value: "štyridsať osem"},
		{Number: 47, Value: "štyridsať sedem"},
		{Number: 46, Value: "štyridsať šesť"},
		{Number: 45, Value: "štyridsať päť"},
		{Number: 44, Value: "štyridsať štyri"},
		{Number: 43, Value: "štyridsať tri"},
		{Number: 42, Value: "štyridsať dva"},
		{Number: 41, Value: "štyridsať jeden"},
		{Number: 40, Value: "štyridsať"},
		{Number: 39, Value: "tridsať deväť"},
		{Number: 38, Value: "tridsať osem"},
		{Number: 37, Value: "tridsať sedem"},
		{Number: 36, Value: "tridsať šesť"},
		{Number: 35, Value: "tridsať päť"},
		{Number: 34, Value: "tridsať štyri"},
		{Number: 33, Value: "tridsať tri"},
		{Number: 32, Value: "tridsať dva"},
		{Number: 31, Value: "tridsať jeden"},
		{Number: 30, Value: "tridsať"},
		{Number: 29, Value: "dvadsať deväť"},
		{Number: 28, Value: "dvadsať osem"},
		{Number: 27, Value: "dvadsať sedem"},
		{Number: 26, Value: "dvadsať šesť"},
		{Number: 25, Value: "dvadsať päť"},
		{Number: 24, Value: "dvadsať štyri"},
		{Number: 23, Value: "dvadsať tri"},
		{Number: 22, Value: "dvadsať dva"},
		{Number: 21, Value: "dvadsať jeden"},
		{Number: 20, Value: "dvadsať"},
		{Number: 19, Value: "devätnásť"},
		{Number: 18, Value: "osemnásť"},
		{Number: 17, Value: "sedemnásť"},
		{Number: 16, Value: "šestnásť"},
		{Number: 15, Value: "pätnásť"},
		{Number: 14, Value: "štrnásť"},
		{Number: 13, Value: "trinásť"},
		{Number: 12, Value: "dvanásť"},
		{Number: 11, Value: "jedenásť"},
		{Number: 10, Value: "desať"},
		{Number: 9, Value: "deväť"},
		{Number: 8, Value: "osem"},
		{Number: 7, Value: "sedem"},
		{Number: 6, Value: "šesť"},
		{Number: 5, Value: "päť"},
		{Number: 4, Value: "štyri"},
		{Number: 3, Value: "tri"},
		{Number: 2, Value: "dva"},
		{Number: 1, Value: "jeden"},
		{Number: 0, Value: "nula"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Sto"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "prvý", Suffix: ".", Masculine: "prvý", Feminine: "prvá", Neuter: "prvé"},
		{Number: 2, Word: "druhý", Suffix: ".", Masculine: "druhý", Feminine: "druhá", Neuter: "druhé"},
		{Number: 3, Word: "tretí", Suffix: ".", Masculine: "tretí", Feminine: "tretia", Neuter: "tretie"},
		{Number: 4, Word: "štvrtý", Suffix: ".", Masculine: "štvrtý", Feminine: "štvrtá", Neuter: "štvrté"},
		{Number: 5, Word: "piaty", Suffix: ".", Masculine: "piaty", Feminine: "piata", Neuter: "piate"},
		{Number: 6, Word: "šiesty", Suffix: ".", Masculine: "šiesty", Feminine: "šiesta", Neuter: "šieste"},
		{Number: 7, Word: "siedmy", Suffix: ".", Masculine: "siedmy", Feminine: "siedma", Neuter: "siedme"},
		{Number: 8, Word: "ôsmy", Suffix: ".", Masculine: "ôsmy", Feminine: "ôsma", Neuter: "ôsme"},
		{Number: 9, Word: "deviaty", Suffix: ".", Masculine: "deviaty", Feminine: "deviata", Neuter: "deviate"},
		{Number: 10, Word: "desiaty", Suffix: ".", Masculine: "desiaty", Feminine: "desiata", Neuter: "desiate"},
		{Number: 11, Word: "jedenásty", Suffix: ".", Masculine: "jedenásty", Feminine: "jedenásta", Neuter: "jedenáste"},
		{Number: 12, Word: "dvanásty", Suffix: ".", Masculine: "dvanásty", Feminine: "dvanásta", Neuter: "dvanáste"},
		{Number: 20, Word: "dvadsiaty", Suffix: ".", Masculine: "dvadsiaty", Feminine: "dvadsiata", Neuter: "dvadsiate"},
		{Number: 21, Word: "dvadsiatyprvý", Suffix: ".", Masculine: "dvadsiatyprvý", Feminine: "dvadsiataprvá", Neuter: "dvadsiateprvé"},
		{Number: 30, Word: "tridsiaty", Suffix: ".", Masculine: "tridsiaty", Feminine: "tridsiata", Neuter: "tridsiate"},
		{Number: 50, Word: "päťdesiaty", Suffix: ".", Masculine: "päťdesiaty", Feminine: "päťdesiata", Neuter: "päťdesiate"},
		{Number: 100, Word: "stý", Suffix: ".", Masculine: "stý", Feminine: "stá", Neuter: "sté"},
		{Number: 1000, Word: "tisíci", Suffix: ".", Masculine: "tisíci", Feminine: "tisíca", Neuter: "tisíce"},
	},
	LocaleFormatter: &SlovakSlovakiaFormatter{},
}

// SlovakSlovakiaFormatter handles Slovak (Slovakia) formatting
type SlovakSlovakiaFormatter struct{}

func (f *SlovakSlovakiaFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *SlovakSlovakiaFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *SlovakSlovakiaFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *SlovakSlovakiaFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *SlovakSlovakiaFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *SlovakSlovakiaFormatter) ChopDecimal(d decimal.Decimal, precision int) decimal.Decimal {
	return d.Truncate(int32(precision))
}

func (f *SlovakSlovakiaFormatter) FormatDecimalNumber(amount float64) string {
	return FormatEuropeanDecimal(amount)
}
func (f *SlovakSlovakiaFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with European conventions (period separators, comma decimal, prefix symbol)
	return FormatEuropeanCurrency(amount, currencySymbol)
}
