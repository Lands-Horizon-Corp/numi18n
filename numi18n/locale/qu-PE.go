package locale

import (
	"github.com/shopspring/decimal"
)

// QUPELocale represents the Quechua (Peru) locale
var QUPELocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Sol",
		Plural:   "Soles",
		Singular: "Sol",
		Symbol:   "S/",
		FractionUnit: FractionUnit{
			Name:     "Céntimo",
			Plural:   "Céntimos",
			Singular: "Céntimo",
			Symbol:   "c",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Peru",
		Currency:       "PEN",
		ISO3166Alpha2:  "PE",
		ISO3166Alpha3:  "PER",
		ISO3166Numeric: "604",
		Locale:         "qu-PE",
		Timezone:       []string{"America/Lima"},
		Language:       "qu",
		Emoji:          "🇵🇪",
	},
	Texts: Texts{
		And:   "wan",
		Minus: "ayqiq",
		Only:  "sapalla",
		Point: "t'uqu",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "waranqa waranqa waranqa waranqa waranqa"},
		{Number: 1000000000000, Value: "waranqa waranqa waranqa waranqa"},
		{Number: 1000000000, Value: "waranqa waranqa waranqa"},
		{Number: 1000000, Value: "waranqa waranqa"},
		{Number: 100000, Value: "pachak waranqa"},
		{Number: 90000, Value: "isqun chunka waranqa"},
		{Number: 80000, Value: "pusaq chunka waranqa"},
		{Number: 70000, Value: "qanchis chunka waranqa"},
		{Number: 60000, Value: "suqta chunka waranqa"},
		{Number: 50000, Value: "pichqa chunka waranqa"},
		{Number: 40000, Value: "tawa chunka waranqa"},
		{Number: 30000, Value: "kimsa chunka waranqa"},
		{Number: 20000, Value: "iskay chunka waranqa"},
		{Number: 19000, Value: "chunka isqunniyuq waranqa"},
		{Number: 18000, Value: "chunka pusaqniyuq waranqa"},
		{Number: 17000, Value: "chunka qanchisniyuq waranqa"},
		{Number: 16000, Value: "chunka suqtaniyuq waranqa"},
		{Number: 15000, Value: "chunka pichqaniyuq waranqa"},
		{Number: 14000, Value: "chunka tawaniyuq waranqa"},
		{Number: 13000, Value: "chunka kimsaniyuq waranqa"},
		{Number: 12000, Value: "chunka iskayniyuq waranqa"},
		{Number: 11000, Value: "chunka hukniyuq waranqa"},
		{Number: 10000, Value: "chunka waranqa"},
		{Number: 9000, Value: "isqun waranqa"},
		{Number: 8000, Value: "pusaq waranqa"},
		{Number: 7000, Value: "qanchis waranqa"},
		{Number: 6000, Value: "suqta waranqa"},
		{Number: 5000, Value: "pichqa waranqa"},
		{Number: 4000, Value: "tawa waranqa"},
		{Number: 3000, Value: "kimsa waranqa"},
		{Number: 2000, Value: "iskay waranqa"},
		{Number: 1000, Value: "waranqa"},
		{Number: 900, Value: "isqun pachak"},
		{Number: 800, Value: "pusaq pachak"},
		{Number: 700, Value: "qanchis pachak"},
		{Number: 600, Value: "suqta pachak"},
		{Number: 500, Value: "pichqa pachak"},
		{Number: 400, Value: "tawa pachak"},
		{Number: 300, Value: "kimsa pachak"},
		{Number: 200, Value: "iskay pachak"},
		{Number: 100, Value: "pachak"},
		{Number: 99, Value: "isqun chunka isqunniyuq"},
		{Number: 98, Value: "isqun chunka pusaqniyuq"},
		{Number: 97, Value: "isqun chunka qanchisniyuq"},
		{Number: 96, Value: "isqun chunka suqtaniyuq"},
		{Number: 95, Value: "isqun chunka pichqaniyuq"},
		{Number: 94, Value: "isqun chunka tawaniyuq"},
		{Number: 93, Value: "isqun chunka kimsaniyuq"},
		{Number: 92, Value: "isqun chunka iskayniyuq"},
		{Number: 91, Value: "isqun chunka hukniyuq"},
		{Number: 90, Value: "isqun chunka"},
		{Number: 89, Value: "pusaq chunka isqunniyuq"},
		{Number: 88, Value: "pusaq chunka pusaqniyuq"},
		{Number: 87, Value: "pusaq chunka qanchisniyuq"},
		{Number: 86, Value: "pusaq chunka suqtaniyuq"},
		{Number: 85, Value: "pusaq chunka pichqaniyuq"},
		{Number: 84, Value: "pusaq chunka tawaniyuq"},
		{Number: 83, Value: "pusaq chunka kimsaniyuq"},
		{Number: 82, Value: "pusaq chunka iskayniyuq"},
		{Number: 81, Value: "pusaq chunka hukniyuq"},
		{Number: 80, Value: "pusaq chunka"},
		{Number: 79, Value: "qanchis chunka isqunniyuq"},
		{Number: 78, Value: "qanchis chunka pusaqniyuq"},
		{Number: 77, Value: "qanchis chunka qanchisniyuq"},
		{Number: 76, Value: "qanchis chunka suqtaniyuq"},
		{Number: 75, Value: "qanchis chunka pichqaniyuq"},
		{Number: 74, Value: "qanchis chunka tawaniyuq"},
		{Number: 73, Value: "qanchis chunka kimsaniyuq"},
		{Number: 72, Value: "qanchis chunka iskayniyuq"},
		{Number: 71, Value: "qanchis chunka hukniyuq"},
		{Number: 70, Value: "qanchis chunka"},
		{Number: 69, Value: "suqta chunka isqunniyuq"},
		{Number: 68, Value: "suqta chunka pusaqniyuq"},
		{Number: 67, Value: "suqta chunka qanchisniyuq"},
		{Number: 66, Value: "suqta chunka suqtaniyuq"},
		{Number: 65, Value: "suqta chunka pichqaniyuq"},
		{Number: 64, Value: "suqta chunka tawaniyuq"},
		{Number: 63, Value: "suqta chunka kimsaniyuq"},
		{Number: 62, Value: "suqta chunka iskayniyuq"},
		{Number: 61, Value: "suqta chunka hukniyuq"},
		{Number: 60, Value: "suqta chunka"},
		{Number: 59, Value: "pichqa chunka isqunniyuq"},
		{Number: 58, Value: "pichqa chunka pusaqniyuq"},
		{Number: 57, Value: "pichqa chunka qanchisniyuq"},
		{Number: 56, Value: "pichqa chunka suqtaniyuq"},
		{Number: 55, Value: "pichqa chunka pichqaniyuq"},
		{Number: 54, Value: "pichqa chunka tawaniyuq"},
		{Number: 53, Value: "pichqa chunka kimsaniyuq"},
		{Number: 52, Value: "pichqa chunka iskayniyuq"},
		{Number: 51, Value: "pichqa chunka hukniyuq"},
		{Number: 50, Value: "pichqa chunka"},
		{Number: 49, Value: "tawa chunka isqunniyuq"},
		{Number: 48, Value: "tawa chunka pusaqniyuq"},
		{Number: 47, Value: "tawa chunka qanchisniyuq"},
		{Number: 46, Value: "tawa chunka suqtaniyuq"},
		{Number: 45, Value: "tawa chunka pichqaniyuq"},
		{Number: 44, Value: "tawa chunka tawaniyuq"},
		{Number: 43, Value: "tawa chunka kimsaniyuq"},
		{Number: 42, Value: "tawa chunka iskayniyuq"},
		{Number: 41, Value: "tawa chunka hukniyuq"},
		{Number: 40, Value: "tawa chunka"},
		{Number: 39, Value: "kimsa chunka isqunniyuq"},
		{Number: 38, Value: "kimsa chunka pusaqniyuq"},
		{Number: 37, Value: "kimsa chunka qanchisniyuq"},
		{Number: 36, Value: "kimsa chunka suqtaniyuq"},
		{Number: 35, Value: "kimsa chunka pichqaniyuq"},
		{Number: 34, Value: "kimsa chunka tawaniyuq"},
		{Number: 33, Value: "kimsa chunka kimsaniyuq"},
		{Number: 32, Value: "kimsa chunka iskayniyuq"},
		{Number: 31, Value: "kimsa chunka hukniyuq"},
		{Number: 30, Value: "kimsa chunka"},
		{Number: 29, Value: "iskay chunka isqunniyuq"},
		{Number: 28, Value: "iskay chunka pusaqniyuq"},
		{Number: 27, Value: "iskay chunka qanchisniyuq"},
		{Number: 26, Value: "iskay chunka suqtaniyuq"},
		{Number: 25, Value: "iskay chunka pichqaniyuq"},
		{Number: 24, Value: "iskay chunka tawaniyuq"},
		{Number: 23, Value: "iskay chunka kimsaniyuq"},
		{Number: 22, Value: "iskay chunka iskayniyuq"},
		{Number: 21, Value: "iskay chunka hukniyuq"},
		{Number: 20, Value: "iskay chunka"},
		{Number: 19, Value: "chunka isqunniyuq"},
		{Number: 18, Value: "chunka pusaqniyuq"},
		{Number: 17, Value: "chunka qanchisniyuq"},
		{Number: 16, Value: "chunka suqtaniyuq"},
		{Number: 15, Value: "chunka pichqaniyuq"},
		{Number: 14, Value: "chunka tawaniyuq"},
		{Number: 13, Value: "chunka kimsaniyuq"},
		{Number: 12, Value: "chunka iskayniyuq"},
		{Number: 11, Value: "chunka hukniyuq"},
		{Number: 10, Value: "chunka"},
		{Number: 9, Value: "isqun"},
		{Number: 8, Value: "pusaq"},
		{Number: 7, Value: "qanchis"},
		{Number: 6, Value: "suqta"},
		{Number: 5, Value: "pichqa"},
		{Number: 4, Value: "tawa"},
		{Number: 3, Value: "kimsa"},
		{Number: 2, Value: "iskay"},
		{Number: 1, Value: "huk"},
		{Number: 0, Value: "ch'usaq"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Pachak"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "ñawpaq", Suffix: "ñiq", Masculine: "ñawpaq", Feminine: "ñawpaq", Neuter: "ñawpaq"},
		{Number: 2, Word: "iskayñiq", Suffix: "ñiq", Masculine: "iskayñiq", Feminine: "iskayñiq", Neuter: "iskayñiq"},
		{Number: 3, Word: "kimsañiq", Suffix: "ñiq", Masculine: "kimsañiq", Feminine: "kimsañiq", Neuter: "kimsañiq"},
		{Number: 4, Word: "tawañiq", Suffix: "ñiq", Masculine: "tawañiq", Feminine: "tawañiq", Neuter: "tawañiq"},
		{Number: 5, Word: "pichqañiq", Suffix: "ñiq", Masculine: "pichqañiq", Feminine: "pichqañiq", Neuter: "pichqañiq"},
		{Number: 6, Word: "suqtañiq", Suffix: "ñiq", Masculine: "suqtañiq", Feminine: "suqtañiq", Neuter: "suqtañiq"},
		{Number: 7, Word: "qanchisñiq", Suffix: "ñiq", Masculine: "qanchisñiq", Feminine: "qanchisñiq", Neuter: "qanchisñiq"},
		{Number: 8, Word: "pusaqñiq", Suffix: "ñiq", Masculine: "pusaqñiq", Feminine: "pusaqñiq", Neuter: "pusaqñiq"},
		{Number: 9, Word: "isqunñiq", Suffix: "ñiq", Masculine: "isqunñiq", Feminine: "isqunñiq", Neuter: "isqunñiq"},
		{Number: 10, Word: "chunkañiq", Suffix: "ñiq", Masculine: "chunkañiq", Feminine: "chunkañiq", Neuter: "chunkañiq"},
		{Number: 11, Word: "chunka hukniyuqñiq", Suffix: "ñiq", Masculine: "chunka hukniyuqñiq", Feminine: "chunka hukniyuqñiq", Neuter: "chunka hukniyuqñiq"},
		{Number: 12, Word: "chunka iskayniyuqñiq", Suffix: "ñiq", Masculine: "chunka iskayniyuqñiq", Feminine: "chunka iskayniyuqñiq", Neuter: "chunka iskayniyuqñiq"},
		{Number: 20, Word: "iskay chunkañiq", Suffix: "ñiq", Masculine: "iskay chunkañiq", Feminine: "iskay chunkañiq", Neuter: "iskay chunkañiq"},
		{Number: 21, Word: "iskay chunka hukniyuqñiq", Suffix: "ñiq", Masculine: "iskay chunka hukniyuqñiq", Feminine: "iskay chunka hukniyuqñiq", Neuter: "iskay chunka hukniyuqñiq"},
		{Number: 30, Word: "kimsa chunkañiq", Suffix: "ñiq", Masculine: "kimsa chunkañiq", Feminine: "kimsa chunkañiq", Neuter: "kimsa chunkañiq"},
		{Number: 50, Word: "pichqa chunkañiq", Suffix: "ñiq", Masculine: "pichqa chunkañiq", Feminine: "pichqa chunkañiq", Neuter: "pichqa chunkañiq"},
		{Number: 100, Word: "pachakñiq", Suffix: "ñiq", Masculine: "pachakñiq", Feminine: "pachakñiq", Neuter: "pachakñiq"},
		{Number: 1000, Word: "waranqañiq", Suffix: "ñiq", Masculine: "waranqañiq", Feminine: "waranqañiq", Neuter: "waranqañiq"},
	},
	LocaleFormatter: &QuechuaPeruFormatter{},
}

// QuechuaPeruFormatter handles Quechua (Peru) formatting
type QuechuaPeruFormatter struct{}

func (f *QuechuaPeruFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *QuechuaPeruFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *QuechuaPeruFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *QuechuaPeruFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *QuechuaPeruFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *QuechuaPeruFormatter) ChopDecimal(value decimal.Decimal, precision int) decimal.Decimal {
	return value.Truncate(int32(precision))
}

func (f *QuechuaPeruFormatter) FormatDecimalNumber(amount float64) string {
	return FormatPolishDecimal(amount)
}
func (f *QuechuaPeruFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Polish conventions (comma separators, period decimal, prefix symbol)
	return FormatPolishCurrency(amount, currencySymbol)
}
