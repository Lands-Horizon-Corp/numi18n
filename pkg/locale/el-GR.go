package locale

import (
	"github.com/shopspring/decimal"
)

// ELGRLocale is a NumI18NLocale configured for Greek (Greece) - el-GR
var ELGRLocale = NumI18NLocale{
	LocaleFormatter: &GreekFormatter{},
	Currency: Currency{
		Name:     "Ευρώ",
		Plural:   "Ευρώ",
		Singular: "Ευρώ",
		Symbol:   "€",
		FractionUnit: FractionUnit{
			Name:     "Λεπτό",
			Plural:   "Λεπτά",
			Singular: "Λεπτό",
			Symbol:   "c",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Greece",
		Currency:       "EUR",
		ISO3166Alpha2:  "GR",
		ISO3166Alpha3:  "GRC",
		ISO3166Numeric: "300",
		Locale:         "el-GR",
		Timezone:       []string{"Europe/Athens"},
		Language:       "el",
	},
	Texts: Texts{
		And:   "και",
		Minus: "μείον",
		Only:  "μόνο",
		Point: "τελεία",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Τετράκις εκατομμύριο"},
		{Number: 1000000000000, Value: "Τρισεκατομμύριο"},
		{Number: 1000000000, Value: "Δισεκατομμύριο"},
		{Number: 1000000, Value: "Εκατομμύριο"},
		{Number: 1000, Value: "Χίλια"},
		{Number: 100, Value: "Εκατό"},
		{Number: 90, Value: "Ενενήντα"},
		{Number: 80, Value: "Ογδόντα"},
		{Number: 70, Value: "Εβδομήντα"},
		{Number: 60, Value: "Εξήντα"},
		{Number: 50, Value: "Πενήντα"},
		{Number: 40, Value: "Σαράντα"},
		{Number: 30, Value: "Τριάντα"},
		{Number: 20, Value: "Είκοσι"},
		{Number: 19, Value: "Δεκαεννέα"},
		{Number: 18, Value: "Δεκαοκτώ"},
		{Number: 17, Value: "Δεκαεπτά"},
		{Number: 16, Value: "Δεκαέξι"},
		{Number: 15, Value: "Δεκαπέντε"},
		{Number: 14, Value: "Δεκατέσσερα"},
		{Number: 13, Value: "Δεκατρία"},
		{Number: 12, Value: "Δώδεκα"},
		{Number: 11, Value: "Έντεκα"},
		{Number: 10, Value: "Δέκα"},
		{Number: 9, Value: "Εννέα"},
		{Number: 8, Value: "Οκτώ"},
		{Number: 7, Value: "Επτά"},
		{Number: 6, Value: "Έξι"},
		{Number: 5, Value: "Πέντε"},
		{Number: 4, Value: "Τέσσερα"},
		{Number: 3, Value: "Τρία"},
		{Number: 2, Value: "Δύο"},
		{Number: 1, Value: "Ένα"},
		{Number: 0, Value: "Μηδέν"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Εκατό"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "πρώτη", Suffix: "ος/η/ο", Masculine: "πρώτος", Feminine: "πρώτη", Neuter: "πρώτο"},
		{Number: 2, Word: "δεύτερη", Suffix: "ος/η/ο", Masculine: "δεύτερος", Feminine: "δεύτερη", Neuter: "δεύτερο"},
		{Number: 3, Word: "τρίτη", Suffix: "ος/η/ο", Masculine: "τρίτος", Feminine: "τρίτη", Neuter: "τρίτο"},
		{Number: 4, Word: "τέταρτη", Suffix: "ος/η/ο", Masculine: "τέταρτος", Feminine: "τέταρτη", Neuter: "τέταρτο"},
		{Number: 5, Word: "πέμπτη", Suffix: "ος/η/ο", Masculine: "πέμπτος", Feminine: "πέμπτη", Neuter: "πέμπτο"},
		{Number: 6, Word: "έκτη", Suffix: "ος/η/ο", Masculine: "έκτος", Feminine: "έκτη", Neuter: "έκτο"},
		{Number: 7, Word: "έβδομη", Suffix: "ος/η/ο", Masculine: "έβδομος", Feminine: "έβδομη", Neuter: "έβδομο"},
		{Number: 8, Word: "όγδοη", Suffix: "ος/η/ο", Masculine: "όγδοος", Feminine: "όγδοη", Neuter: "όγδοο"},
		{Number: 9, Word: "ένατη", Suffix: "ος/η/ο", Masculine: "ένατος", Feminine: "ένατη", Neuter: "ένατο"},
		{Number: 10, Word: "δέκατη", Suffix: "ος/η/ο", Masculine: "δέκατος", Feminine: "δέκατη", Neuter: "δέκατο"},
		{Number: 11, Word: "ενδέκατη", Suffix: "ος/η/ο", Masculine: "ενδέκατος", Feminine: "ενδέκατη", Neuter: "ενδέκατο"},
		{Number: 12, Word: "δωδέκατη", Suffix: "ος/η/ο", Masculine: "δωδέκατος", Feminine: "δωδέκατη", Neuter: "δωδέκατο"},
		{Number: 13, Word: "δέκατη τρίτη", Suffix: "ος/η/ο", Masculine: "δέκατος τρίτος", Feminine: "δέκατη τρίτη", Neuter: "δέκατο τρίτο"},
		{Number: 14, Word: "δέκατη τέταρτη", Suffix: "ος/η/ο", Masculine: "δέκατος τέταρτος", Feminine: "δέκατη τέταρτη", Neuter: "δέκατο τέταρτο"},
		{Number: 15, Word: "δέκατη πέμπτη", Suffix: "ος/η/ο", Masculine: "δέκατος πέμπτος", Feminine: "δέκατη πέμπτη", Neuter: "δέκατο πέμπτο"},
		{Number: 16, Word: "δέκατη έκτη", Suffix: "ος/η/ο", Masculine: "δέκατος έκτος", Feminine: "δέκατη έκτη", Neuter: "δέκατο έκτο"},
		{Number: 17, Word: "δέκατη έβδομη", Suffix: "ος/η/ο", Masculine: "δέκατος έβδομος", Feminine: "δέκατη έβδομη", Neuter: "δέκατο έβδομο"},
		{Number: 18, Word: "δέκατη όγδοη", Suffix: "ος/η/ο", Masculine: "δέκατος όγδοος", Feminine: "δέκατη όγδοη", Neuter: "δέκατο όγδοο"},
		{Number: 19, Word: "δέκατη ένατη", Suffix: "ος/η/ο", Masculine: "δέκατος ένατος", Feminine: "δέκατη ένατη", Neuter: "δέκατο ένατο"},
		{Number: 20, Word: "εικοστή", Suffix: "ος/η/ο", Masculine: "εικοστός", Feminine: "εικοστή", Neuter: "εικοστό"},
		{Number: 21, Word: "εικοστή πρώτη", Suffix: "ος/η/ο", Masculine: "εικοστός πρώτος", Feminine: "εικοστή πρώτη", Neuter: "εικοστό πρώτο"},
		{Number: 30, Word: "τριακοστή", Suffix: "ος/η/ο", Masculine: "τριακοστός", Feminine: "τριακοστή", Neuter: "τριακοστό"},
		{Number: 40, Word: "τεσσαρακοστή", Suffix: "ος/η/ο", Masculine: "τεσσαρακοστός", Feminine: "τεσσαρακοστή", Neuter: "τεσσαρακοστό"},
		{Number: 50, Word: "πεντηκοστή", Suffix: "ος/η/ο", Masculine: "πεντηκοστός", Feminine: "πεντηκοστή", Neuter: "πεντηκοστό"},
		{Number: 100, Word: "εκατοστή", Suffix: "ος/η/ο", Masculine: "εκατοστός", Feminine: "εκατοστή", Neuter: "εκατοστό"},
		{Number: 1000, Word: "χιλιοστή", Suffix: "ος/η/ο", Masculine: "χιλιοστός", Feminine: "χιλιοστή", Neuter: "χιλιοστό"},
		{Number: 1000000, Word: "εκατομμυριοστή", Suffix: "ος/η/ο", Masculine: "εκατομμυριοστός", Feminine: "εκατομμυριοστή", Neuter: "εκατομμυριοστό"},
	},
}

// GreekFormatter handles Greek (Greece) formatting
type GreekFormatter struct{}

func (f *GreekFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *GreekFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *GreekFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *GreekFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *GreekFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *GreekFormatter) ChopDecimal(value decimal.Decimal, precision int) decimal.Decimal {
	return value.Truncate(int32(precision))
}

func (f *GreekFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *GreekFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
