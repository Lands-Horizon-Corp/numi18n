package locale

import (
	"github.com/shopspring/decimal"
)

// UGCNLocale represents the Uyghur (China) locale
var UGCNLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Chinese Yuan",
		Plural:   "يۈەن",
		Singular: "يۈەن",
		Symbol:   "¥",
		FractionUnit: FractionUnit{
			Name:     "Jiao",
			Plural:   "جىياو",
			Singular: "جىياو",
			Symbol:   "角",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "China",
		Currency:       "CNY",
		ISO3166Alpha2:  "CN",
		ISO3166Alpha3:  "CHN",
		ISO3166Numeric: "156",
		Locale:         "ug-CN",
		Timezone:       []string{"Asia/Shanghai"},
		Language:       "ug",
		Emoji:          "🇨🇳",
	},
	Texts: Texts{
		And:   "ۋە",
		Minus: "مىنۇس",
		Only:  "پەقەت",
		Point: "چېكىت",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "بىر تىرىللىيون"},
		{Number: 1000000000000, Value: "بىر تىرىللىيون"},
		{Number: 1000000000, Value: "بىر مىليارد"},
		{Number: 1000000, Value: "بىر مىليون"},
		{Number: 100000, Value: "بىر يۈز مىڭ"},
		{Number: 90000, Value: "توقسان مىڭ"},
		{Number: 80000, Value: "سەكسەن مىڭ"},
		{Number: 70000, Value: "يەتمىش مىڭ"},
		{Number: 60000, Value: "ئاتمىش مىڭ"},
		{Number: 50000, Value: "ئەللىك مىڭ"},
		{Number: 40000, Value: "قىرىق مىڭ"},
		{Number: 30000, Value: "ئوتتۇز مىڭ"},
		{Number: 20000, Value: "يىگىرمە مىڭ"},
		{Number: 19000, Value: "ئون توققۇز مىڭ"},
		{Number: 18000, Value: "ئون سەككىز مىڭ"},
		{Number: 17000, Value: "ئون يەتتە مىڭ"},
		{Number: 16000, Value: "ئون ئالتە مىڭ"},
		{Number: 15000, Value: "ئون بەش مىڭ"},
		{Number: 14000, Value: "ئون تۆت مىڭ"},
		{Number: 13000, Value: "ئون ئۈچ مىڭ"},
		{Number: 12000, Value: "ئون ئىككى مىڭ"},
		{Number: 11000, Value: "ئون بىر مىڭ"},
		{Number: 10000, Value: "ئون مىڭ"},
		{Number: 9000, Value: "توققۇز مىڭ"},
		{Number: 8000, Value: "سەككىز مىڭ"},
		{Number: 7000, Value: "يەتتە مىڭ"},
		{Number: 6000, Value: "ئالتە مىڭ"},
		{Number: 5000, Value: "بەش مىڭ"},
		{Number: 4000, Value: "تۆت مىڭ"},
		{Number: 3000, Value: "ئۈچ مىڭ"},
		{Number: 2000, Value: "ئىككى مىڭ"},
		{Number: 1000, Value: "بىر مىڭ"},
		{Number: 900, Value: "توققۇز يۈز"},
		{Number: 800, Value: "سەككىز يۈز"},
		{Number: 700, Value: "يەتتە يۈز"},
		{Number: 600, Value: "ئالتە يۈز"},
		{Number: 500, Value: "بەش يۈز"},
		{Number: 400, Value: "تۆت يۈز"},
		{Number: 300, Value: "ئۈچ يۈز"},
		{Number: 200, Value: "ئىككى يۈز"},
		{Number: 100, Value: "بىر يۈز"},
		{Number: 99, Value: "توقسان توققۇز"},
		{Number: 98, Value: "توقسان سەككىز"},
		{Number: 97, Value: "توقسان يەتتە"},
		{Number: 96, Value: "توقسان ئالتە"},
		{Number: 95, Value: "توقسان بەش"},
		{Number: 94, Value: "توقسان تۆت"},
		{Number: 93, Value: "توقسان ئۈچ"},
		{Number: 92, Value: "توقسان ئىككى"},
		{Number: 91, Value: "توقسان بىر"},
		{Number: 90, Value: "توقسان"},
		{Number: 89, Value: "سەكسەن توققۇز"},
		{Number: 88, Value: "سەكسەن سەككىز"},
		{Number: 87, Value: "سەكسەن يەتتە"},
		{Number: 86, Value: "سەكسەن ئالتە"},
		{Number: 85, Value: "سەكسەن بەش"},
		{Number: 84, Value: "سەكسەن تۆت"},
		{Number: 83, Value: "سەكسەن ئۈچ"},
		{Number: 82, Value: "سەكسەن ئىككى"},
		{Number: 81, Value: "سەكسەن بىر"},
		{Number: 80, Value: "سەكسەن"},
		{Number: 79, Value: "يەتمىش توققۇز"},
		{Number: 78, Value: "يەتمىش سەككىز"},
		{Number: 77, Value: "يەتمىش يەتتە"},
		{Number: 76, Value: "يەتمىش ئالتە"},
		{Number: 75, Value: "يەتمىش بەش"},
		{Number: 74, Value: "يەتمىش تۆت"},
		{Number: 73, Value: "يەتمىش ئۈچ"},
		{Number: 72, Value: "يەتمىش ئىككى"},
		{Number: 71, Value: "يەتمىش بىر"},
		{Number: 70, Value: "يەتمىش"},
		{Number: 69, Value: "ئاتمىش توققۇز"},
		{Number: 68, Value: "ئاتمىش سەككىز"},
		{Number: 67, Value: "ئاتمىش يەتتە"},
		{Number: 66, Value: "ئاتمىش ئالتە"},
		{Number: 65, Value: "ئاتمىش بەش"},
		{Number: 64, Value: "ئاتمىش تۆت"},
		{Number: 63, Value: "ئاتمىش ئۈچ"},
		{Number: 62, Value: "ئاتمىش ئىككى"},
		{Number: 61, Value: "ئاتمىش بىر"},
		{Number: 60, Value: "ئاتمىش"},
		{Number: 59, Value: "ئەللىك توققۇز"},
		{Number: 58, Value: "ئەللىك سەككىز"},
		{Number: 57, Value: "ئەللىك يەتتە"},
		{Number: 56, Value: "ئەللىك ئالتە"},
		{Number: 55, Value: "ئەللىك بەش"},
		{Number: 54, Value: "ئەللىك تۆت"},
		{Number: 53, Value: "ئەللىك ئۈچ"},
		{Number: 52, Value: "ئەللىك ئىككى"},
		{Number: 51, Value: "ئەللىك بىر"},
		{Number: 50, Value: "ئەللىك"},
		{Number: 49, Value: "قىرىق توققۇز"},
		{Number: 48, Value: "قىرىق سەككىز"},
		{Number: 47, Value: "قىرىق يەتتە"},
		{Number: 46, Value: "قىرىق ئالتە"},
		{Number: 45, Value: "قىرىق بەش"},
		{Number: 44, Value: "قىرىق تۆت"},
		{Number: 43, Value: "قىرىق ئۈچ"},
		{Number: 42, Value: "قىرىق ئىككى"},
		{Number: 41, Value: "قىرىق بىر"},
		{Number: 40, Value: "قىرىق"},
		{Number: 39, Value: "ئوتتۇز توققۇز"},
		{Number: 38, Value: "ئوتتۇز سەككىز"},
		{Number: 37, Value: "ئوتتۇز يەتتە"},
		{Number: 36, Value: "ئوتتۇز ئالتە"},
		{Number: 35, Value: "ئوتتۇز بەش"},
		{Number: 34, Value: "ئوتتۇز تۆت"},
		{Number: 33, Value: "ئوتتۇز ئۈچ"},
		{Number: 32, Value: "ئوتتۇز ئىككى"},
		{Number: 31, Value: "ئوتتۇز بىر"},
		{Number: 30, Value: "ئوتتۇز"},
		{Number: 29, Value: "يىگىرمە توققۇز"},
		{Number: 28, Value: "يىگىرمە سەككىز"},
		{Number: 27, Value: "يىگىرمە يەتتە"},
		{Number: 26, Value: "يىگىرمە ئالتە"},
		{Number: 25, Value: "يىگىرمە بەش"},
		{Number: 24, Value: "يىگىرمە تۆت"},
		{Number: 23, Value: "يىگىرمە ئۈچ"},
		{Number: 22, Value: "يىگىرمە ئىككى"},
		{Number: 21, Value: "يىگىرمە بىر"},
		{Number: 20, Value: "يىگىرمە"},
		{Number: 19, Value: "ئون توققۇز"},
		{Number: 18, Value: "ئون سەككىز"},
		{Number: 17, Value: "ئون يەتتە"},
		{Number: 16, Value: "ئون ئالتە"},
		{Number: 15, Value: "ئون بەش"},
		{Number: 14, Value: "ئون تۆت"},
		{Number: 13, Value: "ئون ئۈچ"},
		{Number: 12, Value: "ئون ئىككى"},
		{Number: 11, Value: "ئون بىر"},
		{Number: 10, Value: "ئون"},
		{Number: 9, Value: "توققۇز"},
		{Number: 8, Value: "سەككىز"},
		{Number: 7, Value: "يەتتە"},
		{Number: 6, Value: "ئالتە"},
		{Number: 5, Value: "بەش"},
		{Number: 4, Value: "تۆت"},
		{Number: 3, Value: "ئۈچ"},
		{Number: 2, Value: "ئىككى"},
		{Number: 1, Value: "بىر"},
		{Number: 0, Value: "نۆل"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "بىر يۈز"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "بىرىنچى", Suffix: ".", Masculine: "بىرىنچى", Feminine: "بىرىنچى", Neuter: "بىرىنچى"},
		{Number: 2, Word: "ئىككىنچى", Suffix: ".", Masculine: "ئىككىنچى", Feminine: "ئىككىنچى", Neuter: "ئىككىنچى"},
		{Number: 3, Word: "ئۈچىنچى", Suffix: ".", Masculine: "ئۈچىنچى", Feminine: "ئۈچىنچى", Neuter: "ئۈچىنچى"},
		{Number: 4, Word: "تۆتىنچى", Suffix: ".", Masculine: "تۆتىنچى", Feminine: "تۆتىنچى", Neuter: "تۆتىنچى"},
		{Number: 5, Word: "بەشىنچى", Suffix: ".", Masculine: "بەشىنچى", Feminine: "بەشىنچى", Neuter: "بەشىنچى"},
		{Number: 6, Word: "ئالتىنچى", Suffix: ".", Masculine: "ئالتىنچى", Feminine: "ئالتىنچى", Neuter: "ئالتىنچى"},
		{Number: 7, Word: "يەتتىنچى", Suffix: ".", Masculine: "يەتتىنچى", Feminine: "يەتتىنچى", Neuter: "يەتتىنچى"},
		{Number: 8, Word: "سەككىزىنچى", Suffix: ".", Masculine: "سەككىزىنچى", Feminine: "سەككىزىنچى", Neuter: "سەككىزىنچى"},
		{Number: 9, Word: "توققۇزىنچى", Suffix: ".", Masculine: "توققۇزىنچى", Feminine: "توققۇزىنچى", Neuter: "توققۇزىنچى"},
		{Number: 10, Word: "ئونىنچى", Suffix: ".", Masculine: "ئونىنچى", Feminine: "ئونىنچى", Neuter: "ئونىنچى"},
		{Number: 11, Word: "ئون بىرىنچى", Suffix: ".", Masculine: "ئون بىرىنچى", Feminine: "ئون بىرىنچى", Neuter: "ئون بىرىنچى"},
		{Number: 12, Word: "ئون ئىككىنچى", Suffix: ".", Masculine: "ئون ئىككىنچى", Feminine: "ئون ئىككىنچى", Neuter: "ئون ئىككىنچى"},
		{Number: 20, Word: "يىگىرمىنچى", Suffix: ".", Masculine: "يىگىرمىنچى", Feminine: "يىگىرمىنچى", Neuter: "يىگىرمىنچى"},
		{Number: 21, Word: "يىگىرمە بىرىنچى", Suffix: ".", Masculine: "يىگىرمە بىرىنچى", Feminine: "يىگىرمە بىرىنچى", Neuter: "يىگىرمە بىرىنچى"},
		{Number: 30, Word: "ئوتتۇزىنچى", Suffix: ".", Masculine: "ئوتتۇزىنچى", Feminine: "ئوتتۇزىنچى", Neuter: "ئوتتۇزىنچى"},
		{Number: 50, Word: "ئەللىكىنچى", Suffix: ".", Masculine: "ئەللىكىنچى", Feminine: "ئەللىكىنچى", Neuter: "ئەللىكىنچى"},
		{Number: 100, Word: "يۈزىنچى", Suffix: ".", Masculine: "يۈزىنچى", Feminine: "يۈزىنچى", Neuter: "يۈزىنچى"},
		{Number: 1000, Word: "مىڭىنچى", Suffix: ".", Masculine: "مىڭىنچى", Feminine: "مىڭىنچى", Neuter: "مىڭىنچى"},
	},
	LocaleFormatter: &UyghurFormatter{},
}

// UyghurFormatter handles Uyghur (China) formatting
type UyghurFormatter struct{}

func (f *UyghurFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *UyghurFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *UyghurFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *UyghurFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *UyghurFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *UyghurFormatter) ChopDecimal(value decimal.Decimal, decimalPlaces int) decimal.Decimal {
	return value.Truncate(int32(decimalPlaces))
}

func (f *UyghurFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAsianDecimal(amount)
}
func (f *UyghurFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Asian conventions (no separators, period decimal, prefix symbol)
	return FormatAsianCurrency(amount, currencySymbol)
}
