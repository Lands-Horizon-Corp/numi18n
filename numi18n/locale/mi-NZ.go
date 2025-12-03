package locale

import (
	"github.com/shopspring/decimal"
)

// MINZLocale represents the Māori (New Zealand) locale
var MINZLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Tāra",
		Plural:   "Tāra",
		Singular: "Tāra",
		Symbol:   "$",
		FractionUnit: FractionUnit{
			Name:     "Hēneti",
			Plural:   "Hēneti",
			Singular: "Hēneti",
			Symbol:   "c",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "New Zealand",
		Currency:       "NZD",
		ISO3166Alpha2:  "NZ",
		ISO3166Alpha3:  "NZL",
		ISO3166Numeric: "554",
		Locale:         "mi-NZ",
		Timezone:       []string{"Pacific/Auckland"},
		Language:       "mi",
		Emoji:          "🇳🇿",
	},
	Texts: Texts{
		And:   "me",
		Minus: "rēhia",
		Only:  "anake",
		Point: "ira",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "kotahi rau miriona miriona"},
		{Number: 1000000000000, Value: "kotahi miriona miriona"},
		{Number: 1000000000, Value: "kotahi piriona"},
		{Number: 1000000, Value: "kotahi miriona"},
		{Number: 100000, Value: "kotahi rau mano"},
		{Number: 90000, Value: "iwa tekau mano"},
		{Number: 80000, Value: "waru tekau mano"},
		{Number: 70000, Value: "whitu tekau mano"},
		{Number: 60000, Value: "ono tekau mano"},
		{Number: 50000, Value: "rima tekau mano"},
		{Number: 40000, Value: "whā tekau mano"},
		{Number: 30000, Value: "toru tekau mano"},
		{Number: 20000, Value: "rua tekau mano"},
		{Number: 19000, Value: "tekau mā iwa mano"},
		{Number: 18000, Value: "tekau mā waru mano"},
		{Number: 17000, Value: "tekau mā whitu mano"},
		{Number: 16000, Value: "tekau mā ono mano"},
		{Number: 15000, Value: "tekau mā rima mano"},
		{Number: 14000, Value: "tekau mā whā mano"},
		{Number: 13000, Value: "tekau mā toru mano"},
		{Number: 12000, Value: "tekau mā rua mano"},
		{Number: 11000, Value: "tekau mā kotahi mano"},
		{Number: 10000, Value: "tekau mano"},
		{Number: 9000, Value: "iwa mano"},
		{Number: 8000, Value: "waru mano"},
		{Number: 7000, Value: "whitu mano"},
		{Number: 6000, Value: "ono mano"},
		{Number: 5000, Value: "rima mano"},
		{Number: 4000, Value: "whā mano"},
		{Number: 3000, Value: "toru mano"},
		{Number: 2000, Value: "rua mano"},
		{Number: 1000, Value: "kotahi mano"},
		{Number: 900, Value: "iwa rau"},
		{Number: 800, Value: "waru rau"},
		{Number: 700, Value: "whitu rau"},
		{Number: 600, Value: "ono rau"},
		{Number: 500, Value: "rima rau"},
		{Number: 400, Value: "whā rau"},
		{Number: 300, Value: "toru rau"},
		{Number: 200, Value: "rua rau"},
		{Number: 100, Value: "kotahi rau"},
		{Number: 99, Value: "iwa tekau mā iwa"},
		{Number: 98, Value: "iwa tekau mā waru"},
		{Number: 97, Value: "iwa tekau mā whitu"},
		{Number: 96, Value: "iwa tekau mā ono"},
		{Number: 95, Value: "iwa tekau mā rima"},
		{Number: 94, Value: "iwa tekau mā whā"},
		{Number: 93, Value: "iwa tekau mā toru"},
		{Number: 92, Value: "iwa tekau mā rua"},
		{Number: 91, Value: "iwa tekau mā kotahi"},
		{Number: 90, Value: "iwa tekau"},
		{Number: 89, Value: "waru tekau mā iwa"},
		{Number: 88, Value: "waru tekau mā waru"},
		{Number: 87, Value: "waru tekau mā whitu"},
		{Number: 86, Value: "waru tekau mā ono"},
		{Number: 85, Value: "waru tekau mā rima"},
		{Number: 84, Value: "waru tekau mā whā"},
		{Number: 83, Value: "waru tekau mā toru"},
		{Number: 82, Value: "waru tekau mā rua"},
		{Number: 81, Value: "waru tekau mā kotahi"},
		{Number: 80, Value: "waru tekau"},
		{Number: 79, Value: "whitu tekau mā iwa"},
		{Number: 78, Value: "whitu tekau mā waru"},
		{Number: 77, Value: "whitu tekau mā whitu"},
		{Number: 76, Value: "whitu tekau mā ono"},
		{Number: 75, Value: "whitu tekau mā rima"},
		{Number: 74, Value: "whitu tekau mā whā"},
		{Number: 73, Value: "whitu tekau mā toru"},
		{Number: 72, Value: "whitu tekau mā rua"},
		{Number: 71, Value: "whitu tekau mā kotahi"},
		{Number: 70, Value: "whitu tekau"},
		{Number: 69, Value: "ono tekau mā iwa"},
		{Number: 68, Value: "ono tekau mā waru"},
		{Number: 67, Value: "ono tekau mā whitu"},
		{Number: 66, Value: "ono tekau mā ono"},
		{Number: 65, Value: "ono tekau mā rima"},
		{Number: 64, Value: "ono tekau mā whā"},
		{Number: 63, Value: "ono tekau mā toru"},
		{Number: 62, Value: "ono tekau mā rua"},
		{Number: 61, Value: "ono tekau mā kotahi"},
		{Number: 60, Value: "ono tekau"},
		{Number: 59, Value: "rima tekau mā iwa"},
		{Number: 58, Value: "rima tekau mā waru"},
		{Number: 57, Value: "rima tekau mā whitu"},
		{Number: 56, Value: "rima tekau mā ono"},
		{Number: 55, Value: "rima tekau mā rima"},
		{Number: 54, Value: "rima tekau mā whā"},
		{Number: 53, Value: "rima tekau mā toru"},
		{Number: 52, Value: "rima tekau mā rua"},
		{Number: 51, Value: "rima tekau mā kotahi"},
		{Number: 50, Value: "rima tekau"},
		{Number: 49, Value: "whā tekau mā iwa"},
		{Number: 48, Value: "whā tekau mā waru"},
		{Number: 47, Value: "whā tekau mā whitu"},
		{Number: 46, Value: "whā tekau mā ono"},
		{Number: 45, Value: "whā tekau mā rima"},
		{Number: 44, Value: "whā tekau mā whā"},
		{Number: 43, Value: "whā tekau mā toru"},
		{Number: 42, Value: "whā tekau mā rua"},
		{Number: 41, Value: "whā tekau mā kotahi"},
		{Number: 40, Value: "whā tekau"},
		{Number: 39, Value: "toru tekau mā iwa"},
		{Number: 38, Value: "toru tekau mā waru"},
		{Number: 37, Value: "toru tekau mā whitu"},
		{Number: 36, Value: "toru tekau mā ono"},
		{Number: 35, Value: "toru tekau mā rima"},
		{Number: 34, Value: "toru tekau mā whā"},
		{Number: 33, Value: "toru tekau mā toru"},
		{Number: 32, Value: "toru tekau mā rua"},
		{Number: 31, Value: "toru tekau mā kotahi"},
		{Number: 30, Value: "toru tekau"},
		{Number: 29, Value: "rua tekau mā iwa"},
		{Number: 28, Value: "rua tekau mā waru"},
		{Number: 27, Value: "rua tekau mā whitu"},
		{Number: 26, Value: "rua tekau mā ono"},
		{Number: 25, Value: "rua tekau mā rima"},
		{Number: 24, Value: "rua tekau mā whā"},
		{Number: 23, Value: "rua tekau mā toru"},
		{Number: 22, Value: "rua tekau mā rua"},
		{Number: 21, Value: "rua tekau mā kotahi"},
		{Number: 20, Value: "rua tekau"},
		{Number: 19, Value: "tekau mā iwa"},
		{Number: 18, Value: "tekau mā waru"},
		{Number: 17, Value: "tekau mā whitu"},
		{Number: 16, Value: "tekau mā ono"},
		{Number: 15, Value: "tekau mā rima"},
		{Number: 14, Value: "tekau mā whā"},
		{Number: 13, Value: "tekau mā toru"},
		{Number: 12, Value: "tekau mā rua"},
		{Number: 11, Value: "tekau mā kotahi"},
		{Number: 10, Value: "tekau"},
		{Number: 9, Value: "iwa"},
		{Number: 8, Value: "waru"},
		{Number: 7, Value: "whitu"},
		{Number: 6, Value: "ono"},
		{Number: 5, Value: "rima"},
		{Number: 4, Value: "whā"},
		{Number: 3, Value: "toru"},
		{Number: 2, Value: "rua"},
		{Number: 1, Value: "kotahi"},
		{Number: 0, Value: "kore"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Kotahi rau"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "tuatahi", Suffix: "o", Masculine: "tuatahi", Feminine: "tuatahi", Neuter: "tuatahi"},
		{Number: 2, Word: "tuarua", Suffix: "o", Masculine: "tuarua", Feminine: "tuarua", Neuter: "tuarua"},
		{Number: 3, Word: "tuatoru", Suffix: "o", Masculine: "tuatoru", Feminine: "tuatoru", Neuter: "tuatoru"},
		{Number: 4, Word: "tuawhā", Suffix: "o", Masculine: "tuawhā", Feminine: "tuawhā", Neuter: "tuawhā"},
		{Number: 5, Word: "tuarima", Suffix: "o", Masculine: "tuarima", Feminine: "tuarima", Neuter: "tuarima"},
		{Number: 6, Word: "tuaono", Suffix: "o", Masculine: "tuaono", Feminine: "tuaono", Neuter: "tuaono"},
		{Number: 7, Word: "tuawhitu", Suffix: "o", Masculine: "tuawhitu", Feminine: "tuawhitu", Neuter: "tuawhitu"},
		{Number: 8, Word: "tuawaru", Suffix: "o", Masculine: "tuawaru", Feminine: "tuawaru", Neuter: "tuawaru"},
		{Number: 9, Word: "tuaiwa", Suffix: "o", Masculine: "tuaiwa", Feminine: "tuaiwa", Neuter: "tuaiwa"},
		{Number: 10, Word: "tuatekau", Suffix: "o", Masculine: "tuatekau", Feminine: "tuatekau", Neuter: "tuatekau"},
		{Number: 11, Word: "tua tekau mā kotahi", Suffix: "o", Masculine: "tua tekau mā kotahi", Feminine: "tua tekau mā kotahi", Neuter: "tua tekau mā kotahi"},
		{Number: 12, Word: "tua tekau mā rua", Suffix: "o", Masculine: "tua tekau mā rua", Feminine: "tua tekau mā rua", Neuter: "tua tekau mā rua"},
		{Number: 20, Word: "tua rua tekau", Suffix: "o", Masculine: "tua rua tekau", Feminine: "tua rua tekau", Neuter: "tua rua tekau"},
		{Number: 21, Word: "tua rua tekau mā kotahi", Suffix: "o", Masculine: "tua rua tekau mā kotahi", Feminine: "tua rua tekau mā kotahi", Neuter: "tua rua tekau mā kotahi"},
		{Number: 30, Word: "tua toru tekau", Suffix: "o", Masculine: "tua toru tekau", Feminine: "tua toru tekau", Neuter: "tua toru tekau"},
		{Number: 100, Word: "tua kotahi rau", Suffix: "o", Masculine: "tua kotahi rau", Feminine: "tua kotahi rau", Neuter: "tua kotahi rau"},
		{Number: 1000, Word: "tua kotahi mano", Suffix: "o", Masculine: "tua kotahi mano", Feminine: "tua kotahi mano", Neuter: "tua kotahi mano"},
	},
	LocaleFormatter: &MāoriFormatter{},
}

// MāoriFormatter handles Māori-specific formatting
type MāoriFormatter struct{}

func (f *MāoriFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *MāoriFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *MāoriFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *MāoriFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *MāoriFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *MāoriFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 0
	}
	return amount.Truncate(int32(precision))
}

func (f *MāoriFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *MāoriFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
