package locale

import (
	"github.com/shopspring/decimal"
)

// SQALLocale represents the Albanian (Albania) locale
var SQALLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Lek",
		Plural:   "Lekë",
		Singular: "Lek",
		Symbol:   "L",
		FractionUnit: FractionUnit{
			Name:     "Qindarkë",
			Plural:   "Qindarka",
			Singular: "Qindarkë",
			Symbol:   "q",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Albania",
		Currency:       "ALL",
		ISO3166Alpha2:  "AL",
		ISO3166Alpha3:  "ALB",
		ISO3166Numeric: "008",
		Locale:         "sq-AL",
		Timezone:       []string{"Europe/Tirane"},
		Language:       "sq",
		Emoji:          "🇦🇱",
	},
	Texts: Texts{
		And:   "dhe",
		Minus: "minus",
		Only:  "vetëm",
		Point: "pikë",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "një katrilion"},
		{Number: 1000000000000, Value: "një trilion"},
		{Number: 1000000000, Value: "një miliard"},
		{Number: 1000000, Value: "një milion"},
		{Number: 100000, Value: "njëqind mijë"},
		{Number: 90000, Value: "nëntëdhjetë mijë"},
		{Number: 80000, Value: "tetëdhjetë mijë"},
		{Number: 70000, Value: "shtatëdhjetë mijë"},
		{Number: 60000, Value: "gjashtëdhjetë mijë"},
		{Number: 50000, Value: "pesëdhjetë mijë"},
		{Number: 40000, Value: "dyzet mijë"},
		{Number: 30000, Value: "tridhjetë mijë"},
		{Number: 20000, Value: "njëzet mijë"},
		{Number: 19000, Value: "nëntëmbëdhjetë mijë"},
		{Number: 18000, Value: "tetëmbëdhjetë mijë"},
		{Number: 17000, Value: "shtatëmbëdhjetë mijë"},
		{Number: 16000, Value: "gjashtëmbëdhjetë mijë"},
		{Number: 15000, Value: "pesëmbëdhjetë mijë"},
		{Number: 14000, Value: "katërmbëdhjetë mijë"},
		{Number: 13000, Value: "trembëdhjetë mijë"},
		{Number: 12000, Value: "dymbëdhjetë mijë"},
		{Number: 11000, Value: "njëmbëdhjetë mijë"},
		{Number: 10000, Value: "dhjetë mijë"},
		{Number: 9000, Value: "nëntë mijë"},
		{Number: 8000, Value: "tetë mijë"},
		{Number: 7000, Value: "shtatë mijë"},
		{Number: 6000, Value: "gjashtë mijë"},
		{Number: 5000, Value: "pesë mijë"},
		{Number: 4000, Value: "katër mijë"},
		{Number: 3000, Value: "tre mijë"},
		{Number: 2000, Value: "dy mijë"},
		{Number: 1000, Value: "një mijë"},
		{Number: 900, Value: "nëntëqind"},
		{Number: 800, Value: "tetëqind"},
		{Number: 700, Value: "shtatëqind"},
		{Number: 600, Value: "gjashtëqind"},
		{Number: 500, Value: "pesëqind"},
		{Number: 400, Value: "katërqind"},
		{Number: 300, Value: "treqind"},
		{Number: 200, Value: "dyqind"},
		{Number: 100, Value: "njëqind"},
		{Number: 99, Value: "nëntëdhjetë e nëntë"},
		{Number: 98, Value: "nëntëdhjetë e tetë"},
		{Number: 97, Value: "nëntëdhjetë e shtatë"},
		{Number: 96, Value: "nëntëdhjetë e gjashtë"},
		{Number: 95, Value: "nëntëdhjetë e pesë"},
		{Number: 94, Value: "nëntëdhjetë e katër"},
		{Number: 93, Value: "nëntëdhjetë e tre"},
		{Number: 92, Value: "nëntëdhjetë e dy"},
		{Number: 91, Value: "nëntëdhjetë e një"},
		{Number: 90, Value: "nëntëdhjetë"},
		{Number: 89, Value: "tetëdhjetë e nëntë"},
		{Number: 88, Value: "tetëdhjetë e tetë"},
		{Number: 87, Value: "tetëdhjetë e shtatë"},
		{Number: 86, Value: "tetëdhjetë e gjashtë"},
		{Number: 85, Value: "tetëdhjetë e pesë"},
		{Number: 84, Value: "tetëdhjetë e katër"},
		{Number: 83, Value: "tetëdhjetë e tre"},
		{Number: 82, Value: "tetëdhjetë e dy"},
		{Number: 81, Value: "tetëdhjetë e një"},
		{Number: 80, Value: "tetëdhjetë"},
		{Number: 79, Value: "shtatëdhjetë e nëntë"},
		{Number: 78, Value: "shtatëdhjetë e tetë"},
		{Number: 77, Value: "shtatëdhjetë e shtatë"},
		{Number: 76, Value: "shtatëdhjetë e gjashtë"},
		{Number: 75, Value: "shtatëdhjetë e pesë"},
		{Number: 74, Value: "shtatëdhjetë e katër"},
		{Number: 73, Value: "shtatëdhjetë e tre"},
		{Number: 72, Value: "shtatëdhjetë e dy"},
		{Number: 71, Value: "shtatëdhjetë e një"},
		{Number: 70, Value: "shtatëdhjetë"},
		{Number: 69, Value: "gjashtëdhjetë e nëntë"},
		{Number: 68, Value: "gjashtëdhjetë e tetë"},
		{Number: 67, Value: "gjashtëdhjetë e shtatë"},
		{Number: 66, Value: "gjashtëdhjetë e gjashtë"},
		{Number: 65, Value: "gjashtëdhjetë e pesë"},
		{Number: 64, Value: "gjashtëdhjetë e katër"},
		{Number: 63, Value: "gjashtëdhjetë e tre"},
		{Number: 62, Value: "gjashtëdhjetë e dy"},
		{Number: 61, Value: "gjashtëdhjetë e një"},
		{Number: 60, Value: "gjashtëdhjetë"},
		{Number: 59, Value: "pesëdhjetë e nëntë"},
		{Number: 58, Value: "pesëdhjetë e tetë"},
		{Number: 57, Value: "pesëdhjetë e shtatë"},
		{Number: 56, Value: "pesëdhjetë e gjashtë"},
		{Number: 55, Value: "pesëdhjetë e pesë"},
		{Number: 54, Value: "pesëdhjetë e katër"},
		{Number: 53, Value: "pesëdhjetë e tre"},
		{Number: 52, Value: "pesëdhjetë e dy"},
		{Number: 51, Value: "pesëdhjetë e një"},
		{Number: 50, Value: "pesëdhjetë"},
		{Number: 49, Value: "dyzet e nëntë"},
		{Number: 48, Value: "dyzet e tetë"},
		{Number: 47, Value: "dyzet e shtatë"},
		{Number: 46, Value: "dyzet e gjashtë"},
		{Number: 45, Value: "dyzet e pesë"},
		{Number: 44, Value: "dyzet e katër"},
		{Number: 43, Value: "dyzet e tre"},
		{Number: 42, Value: "dyzet e dy"},
		{Number: 41, Value: "dyzet e një"},
		{Number: 40, Value: "dyzet"},
		{Number: 39, Value: "tridhjetë e nëntë"},
		{Number: 38, Value: "tridhjetë e tetë"},
		{Number: 37, Value: "tridhjetë e shtatë"},
		{Number: 36, Value: "tridhjetë e gjashtë"},
		{Number: 35, Value: "tridhjetë e pesë"},
		{Number: 34, Value: "tridhjetë e katër"},
		{Number: 33, Value: "tridhjetë e tre"},
		{Number: 32, Value: "tridhjetë e dy"},
		{Number: 31, Value: "tridhjetë e një"},
		{Number: 30, Value: "tridhjetë"},
		{Number: 29, Value: "njëzet e nëntë"},
		{Number: 28, Value: "njëzet e tetë"},
		{Number: 27, Value: "njëzet e shtatë"},
		{Number: 26, Value: "njëzet e gjashtë"},
		{Number: 25, Value: "njëzet e pesë"},
		{Number: 24, Value: "njëzet e katër"},
		{Number: 23, Value: "njëzet e tre"},
		{Number: 22, Value: "njëzet e dy"},
		{Number: 21, Value: "njëzet e një"},
		{Number: 20, Value: "njëzet"},
		{Number: 19, Value: "nëntëmbëdhjetë"},
		{Number: 18, Value: "tetëmbëdhjetë"},
		{Number: 17, Value: "shtatëmbëdhjetë"},
		{Number: 16, Value: "gjashtëmbëdhjetë"},
		{Number: 15, Value: "pesëmbëdhjetë"},
		{Number: 14, Value: "katërmbëdhjetë"},
		{Number: 13, Value: "trembëdhjetë"},
		{Number: 12, Value: "dymbëdhjetë"},
		{Number: 11, Value: "njëmbëdhjetë"},
		{Number: 10, Value: "dhjetë"},
		{Number: 9, Value: "nëntë"},
		{Number: 8, Value: "tetë"},
		{Number: 7, Value: "shtatë"},
		{Number: 6, Value: "gjashtë"},
		{Number: 5, Value: "pesë"},
		{Number: 4, Value: "katër"},
		{Number: 3, Value: "tre"},
		{Number: 2, Value: "dy"},
		{Number: 1, Value: "një"},
		{Number: 0, Value: "zero"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Njëqind"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "i parë", Suffix: "-të", Masculine: "i parë", Feminine: "e para", Neuter: "e parë"},
		{Number: 2, Word: "i dytë", Suffix: "-të", Masculine: "i dytë", Feminine: "e dyta", Neuter: "e dytë"},
		{Number: 3, Word: "i tretë", Suffix: "-të", Masculine: "i tretë", Feminine: "e treta", Neuter: "e tretë"},
		{Number: 4, Word: "i katërt", Suffix: "-të", Masculine: "i katërt", Feminine: "e katërta", Neuter: "e katërt"},
		{Number: 5, Word: "i pestë", Suffix: "-të", Masculine: "i pestë", Feminine: "e pesta", Neuter: "e pestë"},
		{Number: 6, Word: "i gjashtë", Suffix: "-të", Masculine: "i gjashtë", Feminine: "e gjashta", Neuter: "e gjashtë"},
		{Number: 7, Word: "i shtatë", Suffix: "-të", Masculine: "i shtatë", Feminine: "e shtata", Neuter: "e shtatë"},
		{Number: 8, Word: "i tetë", Suffix: "-të", Masculine: "i tetë", Feminine: "e teta", Neuter: "e tetë"},
		{Number: 9, Word: "i nëntë", Suffix: "-të", Masculine: "i nëntë", Feminine: "e nënta", Neuter: "e nëntë"},
		{Number: 10, Word: "i dhjetë", Suffix: "-të", Masculine: "i dhjetë", Feminine: "e dhjeta", Neuter: "e dhjetë"},
		{Number: 11, Word: "i njëmbëdhjetë", Suffix: "-të", Masculine: "i njëmbëdhjetë", Feminine: "e njëmbëdhjeta", Neuter: "e njëmbëdhjetë"},
		{Number: 12, Word: "i dymbëdhjetë", Suffix: "-të", Masculine: "i dymbëdhjetë", Feminine: "e dymbëdhjeta", Neuter: "e dymbëdhjetë"},
		{Number: 20, Word: "i njëzetë", Suffix: "-të", Masculine: "i njëzetë", Feminine: "e njëzeta", Neuter: "e njëzetë"},
		{Number: 21, Word: "i njëzet e parë", Suffix: "-të", Masculine: "i njëzet e parë", Feminine: "e njëzet e para", Neuter: "e njëzet e parë"},
		{Number: 30, Word: "i tridhjetë", Suffix: "-të", Masculine: "i tridhjetë", Feminine: "e tridhjeta", Neuter: "e tridhjetë"},
		{Number: 50, Word: "i pesëdhjetë", Suffix: "-të", Masculine: "i pesëdhjetë", Feminine: "e pesëdhjeta", Neuter: "e pesëdhjetë"},
		{Number: 100, Word: "i njëqindtë", Suffix: "-të", Masculine: "i njëqindtë", Feminine: "e njëqindta", Neuter: "e njëqindtë"},
		{Number: 1000, Word: "i mijëtë", Suffix: "-të", Masculine: "i mijëtë", Feminine: "e mijëta", Neuter: "e mijëtë"},
	},
	LocaleFormatter: &AlbanianAlbaniaFormatter{},
}

// AlbanianAlbaniaFormatter handles Albanian (Albania) formatting
type AlbanianAlbaniaFormatter struct{}

func (f *AlbanianAlbaniaFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *AlbanianAlbaniaFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *AlbanianAlbaniaFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *AlbanianAlbaniaFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *AlbanianAlbaniaFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *AlbanianAlbaniaFormatter) ChopDecimal(d decimal.Decimal, precision int) decimal.Decimal {
	return d.Truncate(int32(precision))
}

func (f *AlbanianAlbaniaFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *AlbanianAlbaniaFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
