package locale

import (
	"strings"
	"github.com/shopspring/decimal"
)

// ITLocale is a NumI18NLocale configured for Italy (it-IT)
var ITLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Euro",
		Plural:   "Euro",
		Singular: "Euro",
		Symbol:   "€",
		FractionUnit: FractionUnit{
			Name:     "Centesimo",
			Plural:   "Centesimi",
			Singular: "Centesimo",
			Symbol:   "¢",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Italy",
		Currency:       "EUR",
		ISO3166Alpha2:  "IT",
		ISO3166Alpha3:  "ITA",
		ISO3166Numeric: "380",
		Locale:         "it-IT",
		Timezone:       []string{"Europe/Rome"},
		Language:       "it",
	},
	Texts: Texts{
		And:   "E",
		Minus: "Meno",
		Only:  "Solo",
		Point: "Virgola",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Un biliardo"},
		{Number: 1000000000000, Value: "Un bilione"},
		{Number: 1000000000, Value: "Un miliardo"},
		{Number: 1000000, Value: "Un milione"},
		{Number: 1000, Value: "Mille"},
		{Number: 100, Value: "Cento"},
		{Number: 90, Value: "Novanta"},
		{Number: 80, Value: "Ottanta"},
		{Number: 70, Value: "Settanta"},
		{Number: 60, Value: "Sessanta"},
		{Number: 50, Value: "Cinquanta"},
		{Number: 40, Value: "Quaranta"},
		{Number: 30, Value: "Trenta"},
		{Number: 20, Value: "Venti"},
		{Number: 19, Value: "Diciannove"},
		{Number: 18, Value: "Diciotto"},
		{Number: 17, Value: "Diciassette"},
		{Number: 16, Value: "Sedici"},
		{Number: 15, Value: "Quindici"},
		{Number: 14, Value: "Quattordici"},
		{Number: 13, Value: "Tredici"},
		{Number: 12, Value: "Dodici"},
		{Number: 11, Value: "Undici"},
		{Number: 10, Value: "Dieci"},
		{Number: 9, Value: "Nove"},
		{Number: 8, Value: "Otto"},
		{Number: 7, Value: "Sette"},
		{Number: 6, Value: "Sei"},
		{Number: 5, Value: "Cinque"},
		{Number: 4, Value: "Quattro"},
		{Number: 3, Value: "Tre"},
		{Number: 2, Value: "Due"},
		{Number: 1, Value: "Uno"},
		{Number: 0, Value: "Zero"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Cento"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "primo", Suffix: "°", Masculine: "primo", Feminine: "prima", Neuter: "primo"},
		{Number: 2, Word: "secondo", Suffix: "°", Masculine: "secondo", Feminine: "seconda", Neuter: "secondo"},
		{Number: 3, Word: "terzo", Suffix: "°", Masculine: "terzo", Feminine: "terza", Neuter: "terzo"},
		{Number: 4, Word: "quarto", Suffix: "°", Masculine: "quarto", Feminine: "quarta", Neuter: "quarto"},
		{Number: 5, Word: "quinto", Suffix: "°", Masculine: "quinto", Feminine: "quinta", Neuter: "quinto"},
		{Number: 6, Word: "sesto", Suffix: "°", Masculine: "sesto", Feminine: "sesta", Neuter: "sesto"},
		{Number: 7, Word: "settimo", Suffix: "°", Masculine: "settimo", Feminine: "settima", Neuter: "settimo"},
		{Number: 8, Word: "ottavo", Suffix: "°", Masculine: "ottavo", Feminine: "ottava", Neuter: "ottavo"},
		{Number: 9, Word: "nono", Suffix: "°", Masculine: "nono", Feminine: "nona", Neuter: "nono"},
		{Number: 10, Word: "decimo", Suffix: "°", Masculine: "decimo", Feminine: "decima", Neuter: "decimo"},
		{Number: 11, Word: "undicesimo", Suffix: "°", Masculine: "undicesimo", Feminine: "undicesima", Neuter: "undicesimo"},
		{Number: 12, Word: "dodicesimo", Suffix: "°", Masculine: "dodicesimo", Feminine: "dodicesima", Neuter: "dodicesimo"},
		{Number: 13, Word: "tredicesimo", Suffix: "°", Masculine: "tredicesimo", Feminine: "tredicesima", Neuter: "tredicesimo"},
		{Number: 14, Word: "quattordicesimo", Suffix: "°", Masculine: "quattordicesimo", Feminine: "quattordicesima", Neuter: "quattordicesimo"},
		{Number: 15, Word: "quindicesimo", Suffix: "°", Masculine: "quindicesimo", Feminine: "quindicesima", Neuter: "quindicesimo"},
		{Number: 16, Word: "sedicesimo", Suffix: "°", Masculine: "sedicesimo", Feminine: "sedicesima", Neuter: "sedicesimo"},
		{Number: 17, Word: "diciassettesimo", Suffix: "°", Masculine: "diciassettesimo", Feminine: "diciassettesima", Neuter: "diciassettesimo"},
		{Number: 18, Word: "diciottesimo", Suffix: "°", Masculine: "diciottesimo", Feminine: "diciottesima", Neuter: "diciottesimo"},
		{Number: 19, Word: "diciannovesimo", Suffix: "°", Masculine: "diciannovesimo", Feminine: "diciannovesima", Neuter: "diciannovesimo"},
		{Number: 20, Word: "ventesimo", Suffix: "°", Masculine: "ventesimo", Feminine: "ventesima", Neuter: "ventesimo"},
		{Number: 21, Word: "ventunesimo", Suffix: "°", Masculine: "ventunesimo", Feminine: "ventunesima", Neuter: "ventunesimo"},
		{Number: 30, Word: "trentesimo", Suffix: "°", Masculine: "trentesimo", Feminine: "trentesima", Neuter: "trentesimo"},
		{Number: 40, Word: "quarantesimo", Suffix: "°", Masculine: "quarantesimo", Feminine: "quarantesima", Neuter: "quarantesimo"},
		{Number: 50, Word: "cinquantesimo", Suffix: "°", Masculine: "cinquantesimo", Feminine: "cinquantesima", Neuter: "cinquantesimo"},
		{Number: 60, Word: "sessantesimo", Suffix: "°", Masculine: "sessantesimo", Feminine: "sessantesima", Neuter: "sessantesimo"},
		{Number: 70, Word: "settantesimo", Suffix: "°", Masculine: "settantesimo", Feminine: "settantesima", Neuter: "settantesimo"},
		{Number: 80, Word: "ottantesimo", Suffix: "°", Masculine: "ottantesimo", Feminine: "ottantesima", Neuter: "ottantesimo"},
		{Number: 90, Word: "novantesimo", Suffix: "°", Masculine: "novantesimo", Feminine: "novantesima", Neuter: "novantesimo"},
		{Number: 100, Word: "centesimo", Suffix: "°", Masculine: "centesimo", Feminine: "centesima", Neuter: "centesimo"},
		{Number: 1000, Word: "millesimo", Suffix: "°", Masculine: "millesimo", Feminine: "millesima", Neuter: "millesimo"},
	},
	LocaleFormatter: &ItalyFormatter{},
}

// ItalyFormatter handles Italian (Italy) formatting
type ItalyFormatter struct{}

func (f *ItalyFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *ItalyFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *ItalyFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *ItalyFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *ItalyFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *ItalyFormatter) ChopDecimal(value decimal.Decimal, places int) decimal.Decimal {
	multiplier := decimal.NewFromInt(1)
	for range places {
		multiplier = multiplier.Mul(decimal.NewFromInt(10))
	}
	return value.Mul(multiplier).Truncate(0).Div(multiplier)
}


func (f *ItalyFormatter) FormatDecimalNumber(amount float64) string {
	return DefaultFormatDecimalNumber(amount, ",", ".")
}
func (f *ItalyFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	formattedNumber := f.FormatDecimalNumber(amount)
	
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}
	
	// Default currency placement for this locale (prefix with symbol)
	if strings.HasPrefix(formattedNumber, "-") {
		formattedNumber = strings.TrimPrefix(formattedNumber, "-")
		return "-" + currencySymbol + formattedNumber
	}
	
	return currencySymbol + formattedNumber
}
