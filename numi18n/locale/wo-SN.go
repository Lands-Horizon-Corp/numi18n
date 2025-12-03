package locale

import (
	"github.com/shopspring/decimal"
)

// WOSNLocale represents the Wolof (Senegal) locale
var WOSNLocale = NumI18NLocale{
	LocaleFormatter: &WolofFormatter{},
	Currency: Currency{
		Name:     "West African CFA Franc",
		Plural:   "franc CFA",
		Singular: "franc CFA",
		Symbol:   "F CFA",
		FractionUnit: FractionUnit{
			Name:     "Centime",
			Plural:   "centime",
			Singular: "centime",
			Symbol:   "c",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Senegal",
		Currency:       "XOF",
		ISO3166Alpha2:  "SN",
		ISO3166Alpha3:  "SEN",
		ISO3166Numeric: "686",
		Locale:         "wo-SN",
		Timezone:       []string{"Africa/Dakar"},
		Language:       "wo",
		Emoji:          "🇸🇳",
		PhoneCode:      "+221",
		Domain:         ".sn",
	},
	Texts: Texts{
		And:   "ak",
		Minus: "jiitu",
		Only:  "rekk",
		Point: "wax",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "benn kadralliyon"},
		{Number: 1000000000000, Value: "benn trilliyon"},
		{Number: 1000000000, Value: "benn milliyard"},
		{Number: 1000000, Value: "benn milliyon"},
		{Number: 100000, Value: "teemer junni"},
		{Number: 90000, Value: "juróom-ñeent-fukk junni"},
		{Number: 80000, Value: "juróom-ñeent junni"},
		{Number: 70000, Value: "juróom-ñaar-fukk junni"},
		{Number: 60000, Value: "juróom-ñaar junni"},
		{Number: 50000, Value: "juróom-benn-fukk junni"},
		{Number: 40000, Value: "juróom-benn junni"},
		{Number: 30000, Value: "ñett-fukk junni"},
		{Number: 20000, Value: "ñaar-fukk junni"},
		{Number: 19000, Value: "fukk ak juróom-ñeent junni"},
		{Number: 18000, Value: "fukk ak juróom-ñeent junni"},
		{Number: 17000, Value: "fukk ak juróom-ñaar junni"},
		{Number: 16000, Value: "fukk ak juróom-benn junni"},
		{Number: 15000, Value: "fukk ak juróom junni"},
		{Number: 14000, Value: "fukk ak ñeent junni"},
		{Number: 13000, Value: "fukk ak ñett junni"},
		{Number: 12000, Value: "fukk ak ñaar junni"},
		{Number: 11000, Value: "fukk ak benn junni"},
		{Number: 10000, Value: "fukk junni"},
		{Number: 9000, Value: "juróom-ñeent junni"},
		{Number: 8000, Value: "juróom-ñett junni"},
		{Number: 7000, Value: "juróom-ñaar junni"},
		{Number: 6000, Value: "juróom-benn junni"},
		{Number: 5000, Value: "juróom junni"},
		{Number: 4000, Value: "ñeent junni"},
		{Number: 3000, Value: "ñett junni"},
		{Number: 2000, Value: "ñaar junni"},
		{Number: 1000, Value: "benn junni"},
		{Number: 900, Value: "juróom-ñeent teemer"},
		{Number: 800, Value: "juróom-ñett teemer"},
		{Number: 700, Value: "juróom-ñaar teemer"},
		{Number: 600, Value: "juróom-benn teemer"},
		{Number: 500, Value: "juróom teemer"},
		{Number: 400, Value: "ñeent teemer"},
		{Number: 300, Value: "ñett teemer"},
		{Number: 200, Value: "ñaar teemer"},
		{Number: 100, Value: "benn teemer"},
		{Number: 99, Value: "juróom-ñeent-fukk ak juróom-ñeent"},
		{Number: 98, Value: "juróom-ñeent-fukk ak juróom-ñett"},
		{Number: 97, Value: "juróom-ñeent-fukk ak juróom-ñaar"},
		{Number: 96, Value: "juróom-ñeent-fukk ak juróom-benn"},
		{Number: 95, Value: "juróom-ñeent-fukk ak juróom"},
		{Number: 94, Value: "juróom-ñeent-fukk ak ñeent"},
		{Number: 93, Value: "juróom-ñeent-fukk ak ñett"},
		{Number: 92, Value: "juróom-ñeent-fukk ak ñaar"},
		{Number: 91, Value: "juróom-ñeent-fukk ak benn"},
		{Number: 90, Value: "juróom-ñeent-fukk"},
		{Number: 89, Value: "juróom-ñeent-fukk ak juróom-ñeent"},
		{Number: 88, Value: "juróom-ñeent-fukk ak juróom-ñett"},
		{Number: 87, Value: "juróom-ñeent-fukk ak juróom-ñaar"},
		{Number: 86, Value: "juróom-ñeent-fukk ak juróom-benn"},
		{Number: 85, Value: "juróom-ñeent-fukk ak juróom"},
		{Number: 84, Value: "juróom-ñeent-fukk ak ñeent"},
		{Number: 83, Value: "juróom-ñeent-fukk ak ñett"},
		{Number: 82, Value: "juróom-ñeent-fukk ak ñaar"},
		{Number: 81, Value: "juróom-ñeent-fukk ak benn"},
		{Number: 80, Value: "juróom-ñeent"},
		{Number: 79, Value: "juróom-ñaar-fukk ak juróom-ñeent"},
		{Number: 78, Value: "juróom-ñaar-fukk ak juróom-ñett"},
		{Number: 77, Value: "juróom-ñaar-fukk ak juróom-ñaar"},
		{Number: 76, Value: "juróom-ñaar-fukk ak juróom-benn"},
		{Number: 75, Value: "juróom-ñaar-fukk ak juróom"},
		{Number: 74, Value: "juróom-ñaar-fukk ak ñeent"},
		{Number: 73, Value: "juróom-ñaar-fukk ak ñett"},
		{Number: 72, Value: "juróom-ñaar-fukk ak ñaar"},
		{Number: 71, Value: "juróom-ñaar-fukk ak benn"},
		{Number: 70, Value: "juróom-ñaar-fukk"},
		{Number: 69, Value: "juróom-ñaar ak juróom-ñeent"},
		{Number: 68, Value: "juróom-ñaar ak juróom-ñett"},
		{Number: 67, Value: "juróom-ñaar ak juróom-ñaar"},
		{Number: 66, Value: "juróom-ñaar ak juróom-benn"},
		{Number: 65, Value: "juróom-ñaar ak juróom"},
		{Number: 64, Value: "juróom-ñaar ak ñeent"},
		{Number: 63, Value: "juróom-ñaar ak ñett"},
		{Number: 62, Value: "juróom-ñaar ak ñaar"},
		{Number: 61, Value: "juróom-ñaar ak benn"},
		{Number: 60, Value: "juróom-ñaar"},
		{Number: 59, Value: "juróom-benn-fukk ak juróom-ñeent"},
		{Number: 58, Value: "juróom-benn-fukk ak juróom-ñett"},
		{Number: 57, Value: "juróom-benn-fukk ak juróom-ñaar"},
		{Number: 56, Value: "juróom-benn-fukk ak juróom-benn"},
		{Number: 55, Value: "juróom-benn-fukk ak juróom"},
		{Number: 54, Value: "juróom-benn-fukk ak ñeent"},
		{Number: 53, Value: "juróom-benn-fukk ak ñett"},
		{Number: 52, Value: "juróom-benn-fukk ak ñaar"},
		{Number: 51, Value: "juróom-benn-fukk ak benn"},
		{Number: 50, Value: "juróom-benn-fukk"},
		{Number: 49, Value: "juróom-benn ak juróom-ñeent"},
		{Number: 48, Value: "juróom-benn ak juróom-ñett"},
		{Number: 47, Value: "juróom-benn ak juróom-ñaar"},
		{Number: 46, Value: "juróom-benn ak juróom-benn"},
		{Number: 45, Value: "juróom-benn ak juróom"},
		{Number: 44, Value: "juróom-benn ak ñeent"},
		{Number: 43, Value: "juróom-benn ak ñett"},
		{Number: 42, Value: "juróom-benn ak ñaar"},
		{Number: 41, Value: "juróom-benn ak benn"},
		{Number: 40, Value: "juróom-benn"},
		{Number: 39, Value: "ñett-fukk ak juróom-ñeent"},
		{Number: 38, Value: "ñett-fukk ak juróom-ñett"},
		{Number: 37, Value: "ñett-fukk ak juróom-ñaar"},
		{Number: 36, Value: "ñett-fukk ak juróom-benn"},
		{Number: 35, Value: "ñett-fukk ak juróom"},
		{Number: 34, Value: "ñett-fukk ak ñeent"},
		{Number: 33, Value: "ñett-fukk ak ñett"},
		{Number: 32, Value: "ñett-fukk ak ñaar"},
		{Number: 31, Value: "ñett-fukk ak benn"},
		{Number: 30, Value: "ñett-fukk"},
		{Number: 29, Value: "ñaar-fukk ak juróom-ñeent"},
		{Number: 28, Value: "ñaar-fukk ak juróom-ñett"},
		{Number: 27, Value: "ñaar-fukk ak juróom-ñaar"},
		{Number: 26, Value: "ñaar-fukk ak juróom-benn"},
		{Number: 25, Value: "ñaar-fukk ak juróom"},
		{Number: 24, Value: "ñaar-fukk ak ñeent"},
		{Number: 23, Value: "ñaar-fukk ak ñett"},
		{Number: 22, Value: "ñaar-fukk ak ñaar"},
		{Number: 21, Value: "ñaar-fukk ak benn"},
		{Number: 20, Value: "ñaar-fukk"},
		{Number: 19, Value: "fukk ak juróom-ñeent"},
		{Number: 18, Value: "fukk ak juróom-ñett"},
		{Number: 17, Value: "fukk ak juróom-ñaar"},
		{Number: 16, Value: "fukk ak juróom-benn"},
		{Number: 15, Value: "fukk ak juróom"},
		{Number: 14, Value: "fukk ak ñeent"},
		{Number: 13, Value: "fukk ak ñett"},
		{Number: 12, Value: "fukk ak ñaar"},
		{Number: 11, Value: "fukk ak benn"},
		{Number: 10, Value: "fukk"},
		{Number: 9, Value: "juróom-ñeent"},
		{Number: 8, Value: "juróom-ñett"},
		{Number: 7, Value: "juróom-ñaar"},
		{Number: 6, Value: "juróom-benn"},
		{Number: 5, Value: "juróom"},
		{Number: 4, Value: "ñeent"},
		{Number: 3, Value: "ñett"},
		{Number: 2, Value: "ñaar"},
		{Number: 1, Value: "benn"},
		{Number: 0, Value: "tus"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Teemer"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "bu njëkk", Suffix: "-ëem", Masculine: "bu njëkk", Feminine: "bu njëkk", Neuter: ""},
		{Number: 2, Word: "bu ñaareel", Suffix: "-ëem", Masculine: "bu ñaareel", Feminine: "bu ñaareel", Neuter: ""},
		{Number: 3, Word: "bu ñetteel", Suffix: "-ëem", Masculine: "bu ñetteel", Feminine: "bu ñetteel", Neuter: ""},
		{Number: 4, Word: "bu ñeenteel", Suffix: "-ëem", Masculine: "bu ñeenteel", Feminine: "bu ñeenteel", Neuter: ""},
		{Number: 5, Word: "bu juroomeel", Suffix: "-ëem", Masculine: "bu juroomeel", Feminine: "bu juroomeel", Neuter: ""},
		{Number: 6, Word: "bu juróom-benneel", Suffix: "-ëem", Masculine: "bu juróom-benneel", Feminine: "bu juróom-benneel", Neuter: ""},
		{Number: 7, Word: "bu juróom-ñaareel", Suffix: "-ëem", Masculine: "bu juróom-ñaareel", Feminine: "bu juróom-ñaareel", Neuter: ""},
		{Number: 8, Word: "bu juróom-ñetteel", Suffix: "-ëem", Masculine: "bu juróom-ñetteel", Feminine: "bu juróom-ñetteel", Neuter: ""},
		{Number: 9, Word: "bu juróom-ñeenteel", Suffix: "-ëem", Masculine: "bu juróom-ñeenteel", Feminine: "bu juróom-ñeenteel", Neuter: ""},
		{Number: 10, Word: "bu fukkeel", Suffix: "-ëem", Masculine: "bu fukkeel", Feminine: "bu fukkeel", Neuter: ""},
		{Number: 11, Word: "bu fukk ak benneel", Suffix: "-ëem", Masculine: "bu fukk ak benneel", Feminine: "bu fukk ak benneel", Neuter: ""},
		{Number: 12, Word: "bu fukk ak ñaareel", Suffix: "-ëem", Masculine: "bu fukk ak ñaareel", Feminine: "bu fukk ak ñaareel", Neuter: ""},
		{Number: 13, Word: "bu fukk ak ñetteel", Suffix: "-ëem", Masculine: "bu fukk ak ñetteel", Feminine: "bu fukk ak ñetteel", Neuter: ""},
		{Number: 14, Word: "bu fukk ak ñeenteel", Suffix: "-ëem", Masculine: "bu fukk ak ñeenteel", Feminine: "bu fukk ak ñeenteel", Neuter: ""},
		{Number: 15, Word: "bu fukk ak juroomeel", Suffix: "-ëem", Masculine: "bu fukk ak juroomeel", Feminine: "bu fukk ak juroomeel", Neuter: ""},
		{Number: 16, Word: "bu fukk ak juróom-benneel", Suffix: "-ëem", Masculine: "bu fukk ak juróom-benneel", Feminine: "bu fukk ak juróom-benneel", Neuter: ""},
		{Number: 17, Word: "bu fukk ak juróom-ñaareel", Suffix: "-ëem", Masculine: "bu fukk ak juróom-ñaareel", Feminine: "bu fukk ak juróom-ñaareel", Neuter: ""},
		{Number: 18, Word: "bu fukk ak juróom-ñetteel", Suffix: "-ëem", Masculine: "bu fukk ak juróom-ñetteel", Feminine: "bu fukk ak juróom-ñetteel", Neuter: ""},
		{Number: 19, Word: "bu fukk ak juróom-ñeenteel", Suffix: "-ëem", Masculine: "bu fukk ak juróom-ñeenteel", Feminine: "bu fukk ak juróom-ñeenteel", Neuter: ""},
		{Number: 20, Word: "bu ñaar-fukkeel", Suffix: "-ëem", Masculine: "bu ñaar-fukkeel", Feminine: "bu ñaar-fukkeel", Neuter: ""},
		{Number: 21, Word: "bu ñaar-fukk ak benneel", Suffix: "-ëem", Masculine: "bu ñaar-fukk ak benneel", Feminine: "bu ñaar-fukk ak benneel", Neuter: ""},
		{Number: 30, Word: "bu ñett-fukkeel", Suffix: "-ëem", Masculine: "bu ñett-fukkeel", Feminine: "bu ñett-fukkeel", Neuter: ""},
		{Number: 40, Word: "bu juróom-benneel", Suffix: "-ëem", Masculine: "bu juróom-benneel", Feminine: "bu juróom-benneel", Neuter: ""},
		{Number: 50, Word: "bu juróom-benn-fukkeel", Suffix: "-ëem", Masculine: "bu juróom-benn-fukkeel", Feminine: "bu juróom-benn-fukkeel", Neuter: ""},
		{Number: 60, Word: "bu juróom-ñaareel", Suffix: "-ëem", Masculine: "bu juróom-ñaareel", Feminine: "bu juróom-ñaareel", Neuter: ""},
		{Number: 70, Word: "bu juróom-ñaar-fukkeel", Suffix: "-ëem", Masculine: "bu juróom-ñaar-fukkeel", Feminine: "bu juróom-ñaar-fukkeel", Neuter: ""},
		{Number: 80, Word: "bu juróom-ñeenteel", Suffix: "-ëem", Masculine: "bu juróom-ñeenteel", Feminine: "bu juróom-ñeenteel", Neuter: ""},
		{Number: 90, Word: "bu juróom-ñeent-fukkeel", Suffix: "-ëem", Masculine: "bu juróom-ñeent-fukkeel", Feminine: "bu juróom-ñeent-fukkeel", Neuter: ""},
		{Number: 100, Word: "bu teemeereel", Suffix: "-ëem", Masculine: "bu teemeereel", Feminine: "bu teemeereel", Neuter: ""},
		{Number: 1000, Word: "bu junneel", Suffix: "-ëem", Masculine: "bu junneel", Feminine: "bu junneel", Neuter: ""},
		{Number: 1000000, Word: "bu milliyoneel", Suffix: "-ëem", Masculine: "bu milliyoneel", Feminine: "bu milliyoneel", Neuter: ""},
		{Number: 1000000000, Word: "bu milliyard", Suffix: "-ëem", Masculine: "bu milliyardeel", Feminine: "bu milliyardeel", Neuter: ""},
	},
}

// WolofFormatter handles Wolof (Senegal) formatting
type WolofFormatter struct{}

func (f *WolofFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *WolofFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *WolofFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *WolofFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *WolofFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *WolofFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	multiplier := decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(precision)))
	return amount.Mul(multiplier).Truncate(0).Div(multiplier)
}

func (f *WolofFormatter) FormatDecimalNumber(amount float64) string {
	return FormatFrenchDecimal(amount)
}
func (f *WolofFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with French conventions (space separators, comma decimal, prefix symbol)
	return FormatFrenchCurrency(amount, currencySymbol)
}
