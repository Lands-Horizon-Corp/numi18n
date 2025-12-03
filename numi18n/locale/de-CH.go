package locale

import (
	"strings"

	"github.com/shopspring/decimal"
)

// DECHLocale is a NumI18NLocale configured for German (Switzerland) - de-CH
var DECHLocale = NumI18NLocale{
	LocaleFormatter: &SwissGermanFormatter{},
	Currency: Currency{
		Name:     "Franken",
		Plural:   "Franken",
		Singular: "Franken",
		Symbol:   "CHF",
		FractionUnit: FractionUnit{
			Name:     "Rappen",
			Plural:   "Rappen",
			Singular: "Rappen",
			Symbol:   "Rp",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Switzerland",
		Currency:       "CHF",
		ISO3166Alpha2:  "CH",
		ISO3166Alpha3:  "CHE",
		ISO3166Numeric: "756",
		Locale:         "de-CH",
		Timezone:       []string{"Europe/Zurich"},
		Language:       "de",
	},
	Texts: Texts{
		And:   "und",
		Minus: "minus",
		Only:  "nur",
		Point: "komma",
	},

	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Billiarde"},
		{Number: 1000000000000, Value: "Billion"},
		{Number: 1000000000, Value: "Milliarde"},
		{Number: 1000000, Value: "Million"},
		{Number: 1000, Value: "Tausend"},
		{Number: 100, Value: "Hundert"},
		{Number: 90, Value: "Neunzig"},
		{Number: 80, Value: "Achtzig"},
		{Number: 70, Value: "Siebzig"},
		{Number: 60, Value: "Sechzig"},
		{Number: 50, Value: "Fünfzig"},
		{Number: 40, Value: "Vierzig"},
		{Number: 30, Value: "Dreißig"},
		{Number: 20, Value: "Zwanzig"},
		{Number: 19, Value: "Neunzehn"},
		{Number: 18, Value: "Achtzehn"},
		{Number: 17, Value: "Siebzehn"},
		{Number: 16, Value: "Sechzehn"},
		{Number: 15, Value: "Fünfzehn"},
		{Number: 14, Value: "Vierzehn"},
		{Number: 13, Value: "Dreizehn"},
		{Number: 12, Value: "Zwölf"},
		{Number: 11, Value: "Elf"},
		{Number: 10, Value: "Zehn"},
		{Number: 9, Value: "Neun"},
		{Number: 8, Value: "Acht"},
		{Number: 7, Value: "Sieben"},
		{Number: 6, Value: "Sechs"},
		{Number: 5, Value: "Fünf"},
		{Number: 4, Value: "Vier"},
		{Number: 3, Value: "Drei"},
		{Number: 2, Value: "Zwei"},
		{Number: 1, Value: "Eins"},
		{Number: 0, Value: "Null"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Einhundert"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "erste", Suffix: ".", Masculine: "erster", Feminine: "erste", Neuter: "erstes"},
		{Number: 2, Word: "zweite", Suffix: ".", Masculine: "zweiter", Feminine: "zweite", Neuter: "zweites"},
		{Number: 3, Word: "dritte", Suffix: ".", Masculine: "dritter", Feminine: "dritte", Neuter: "drittes"},
		{Number: 4, Word: "vierte", Suffix: ".", Masculine: "vierter", Feminine: "vierte", Neuter: "viertes"},
		{Number: 5, Word: "fünfte", Suffix: ".", Masculine: "fünfter", Feminine: "fünfte", Neuter: "fünftes"},
		{Number: 6, Word: "sechste", Suffix: ".", Masculine: "sechster", Feminine: "sechste", Neuter: "sechstes"},
		{Number: 7, Word: "siebte", Suffix: ".", Masculine: "siebter", Feminine: "siebte", Neuter: "siebtes"},
		{Number: 8, Word: "achte", Suffix: ".", Masculine: "achter", Feminine: "achte", Neuter: "achtes"},
		{Number: 9, Word: "neunte", Suffix: ".", Masculine: "neunter", Feminine: "neunte", Neuter: "neuntes"},
		{Number: 10, Word: "zehnte", Suffix: ".", Masculine: "zehnter", Feminine: "zehnte", Neuter: "zehntes"},
		{Number: 11, Word: "elfte", Suffix: ".", Masculine: "elfter", Feminine: "elfte", Neuter: "elftes"},
		{Number: 12, Word: "zwölfte", Suffix: ".", Masculine: "zwölfter", Feminine: "zwölfte", Neuter: "zwölftes"},
		{Number: 13, Word: "dreizehnte", Suffix: ".", Masculine: "dreizehnter", Feminine: "dreizehnte", Neuter: "dreizehntes"},
		{Number: 14, Word: "vierzehnte", Suffix: ".", Masculine: "vierzehnter", Feminine: "vierzehnte", Neuter: "vierzehntes"},
		{Number: 15, Word: "fünfzehnte", Suffix: ".", Masculine: "fünfzehnter", Feminine: "fünfzehnte", Neuter: "fünfzehntes"},
		{Number: 16, Word: "sechzehnte", Suffix: ".", Masculine: "sechzehnter", Feminine: "sechzehnte", Neuter: "sechzehntes"},
		{Number: 17, Word: "siebzehnte", Suffix: ".", Masculine: "siebzehnter", Feminine: "siebzehnte", Neuter: "siebzehntes"},
		{Number: 18, Word: "achtzehnte", Suffix: ".", Masculine: "achtzehnter", Feminine: "achtzehnte", Neuter: "achtzehntes"},
		{Number: 19, Word: "neunzehnte", Suffix: ".", Masculine: "neunzehnter", Feminine: "neunzehnte", Neuter: "neunzehntes"},
		{Number: 20, Word: "zwanzigste", Suffix: ".", Masculine: "zwanzigster", Feminine: "zwanzigste", Neuter: "zwanzigstes"},
		{Number: 21, Word: "einundzwanzigste", Suffix: ".", Masculine: "einundzwanzigster", Feminine: "einundzwanzigste", Neuter: "einundzwanzigstes"},
		{Number: 30, Word: "dreißigste", Suffix: ".", Masculine: "dreißigster", Feminine: "dreißigste", Neuter: "dreißigstes"},
		{Number: 40, Word: "vierzigste", Suffix: ".", Masculine: "vierzigster", Feminine: "vierzigste", Neuter: "vierzigstes"},
		{Number: 50, Word: "fünfzigste", Suffix: ".", Masculine: "fünfzigster", Feminine: "fünfzigste", Neuter: "fünfzigstes"},
		{Number: 100, Word: "hundertste", Suffix: ".", Masculine: "hundertster", Feminine: "hundertste", Neuter: "hundertstes"},
		{Number: 1000, Word: "tausendste", Suffix: ".", Masculine: "tausendster", Feminine: "tausendste", Neuter: "tausendstes"},
		{Number: 1000000, Word: "millionste", Suffix: ".", Masculine: "millionster", Feminine: "millionste", Neuter: "millionstes"},
	},
}

// SwissGermanFormatter handles German (Switzerland) formatting
type SwissGermanFormatter struct{}

func (f *SwissGermanFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return f.convertGermanNumber(number, targetLocale)
}

// convertGermanNumber converts numbers to German compound words
func (f *SwissGermanFormatter) convertGermanNumber(number int64, targetLocale NumI18NLocale) string {
	if number == 0 {
		return "null"
	}

	if number < 0 {
		return "minus " + f.convertGermanNumber(-number, targetLocale)
	}

	// Handle numbers 1-19 directly
	if number <= 19 {
		for _, mapping := range targetLocale.NumberWordsMapping {
			if mapping.Number == number {
				return strings.ToLower(mapping.Value)
			}
		}
	}

	// Handle 20-99 (compound tens)
	if number < 100 {
		tens := (number / 10) * 10
		ones := number % 10

		tensWord := ""
		for _, mapping := range targetLocale.NumberWordsMapping {
			if mapping.Number == tens {
				tensWord = strings.ToLower(mapping.Value)
				break
			}
		}

		if ones == 0 {
			return tensWord
		}

		onesWord := ""
		for _, mapping := range targetLocale.NumberWordsMapping {
			if mapping.Number == ones {
				onesWord = strings.ToLower(mapping.Value)
				break
			}
		}

		return onesWord + "und" + tensWord
	}

	// Handle 100-999 (hundreds)
	if number < 1000 {
		hundreds := number / 100
		remainder := number % 100

		result := ""
		if hundreds == 1 {
			result = "einhundert"
		} else {
			hundredsWord := f.convertGermanNumber(hundreds, targetLocale)
			result = hundredsWord + "hundert"
		}

		if remainder > 0 {
			result += f.convertGermanNumber(remainder, targetLocale)
		}

		return result
	}

	// Handle 1000+ (thousands, millions, etc.)
	if number < 1000000 {
		thousands := number / 1000
		remainder := number % 1000

		result := ""
		if thousands == 1 {
			result = "eintausend"
		} else {
			thousandsWord := f.convertGermanNumber(thousands, targetLocale)
			result = thousandsWord + "tausend"
		}

		if remainder > 0 {
			result += f.convertGermanNumber(remainder, targetLocale)
		}

		return result
	}

	// For larger numbers, fall back to generic conversion but lowercase
	return strings.ToLower(ConvertToWordsWithExactMappingInt64(number, targetLocale))
}

func (f *SwissGermanFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *SwissGermanFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *SwissGermanFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *SwissGermanFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *SwissGermanFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	return amount.Round(int32(precision))
}

func (f *SwissGermanFormatter) FormatDecimalNumber(amount float64) string {
	return FormatEuropeanDecimal(amount)
}

func (f *SwissGermanFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with European conventions (period separators, comma decimal, prefix symbol)
	return FormatEuropeanCurrency(amount, currencySymbol)
}
