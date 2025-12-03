package locale

import (
	"strings"

	"github.com/shopspring/decimal"
)

// NLBELocale represents the Dutch (Belgium) locale
var NLBELocale = NumI18NLocale{
	Currency: Currency{
		Name:     "euro",
		Plural:   "euro",
		Singular: "euro",
		Symbol:   "€",
		FractionUnit: FractionUnit{
			Name:     "cent",
			Plural:   "cent",
			Singular: "cent",
			Symbol:   "c",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Belgium",
		Currency:       "EUR",
		ISO3166Alpha2:  "BE",
		ISO3166Alpha3:  "BEL",
		ISO3166Numeric: "056",
		Locale:         "nl-BE",
		Timezone:       []string{"Europe/Brussels"},
		Language:       "nl",
		Emoji:          "🇧🇪",
	},
	Texts: Texts{
		And:   "en",
		Minus: "min",
		Only:  "alleen",
		Point: "komma",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "biljard"},
		{Number: 1000000000000, Value: "biljoen"},
		{Number: 1000000000, Value: "miljard"},
		{Number: 1000000, Value: "miljoen"},
		{Number: 100000, Value: "honderdduizend"},
		{Number: 90000, Value: "negentigduizend"},
		{Number: 80000, Value: "tachtigduizend"},
		{Number: 70000, Value: "zeventigduizend"},
		{Number: 60000, Value: "zestigduizend"},
		{Number: 50000, Value: "vijftigduizend"},
		{Number: 40000, Value: "veertigduizend"},
		{Number: 30000, Value: "dertigduizend"},
		{Number: 20000, Value: "twintigduizend"},
		{Number: 19000, Value: "negentienduizend"},
		{Number: 18000, Value: "achttienduizend"},
		{Number: 17000, Value: "zeventienduizend"},
		{Number: 16000, Value: "zestienduizend"},
		{Number: 15000, Value: "vijftienduizend"},
		{Number: 14000, Value: "veertienduizend"},
		{Number: 13000, Value: "dertienduizend"},
		{Number: 12000, Value: "twaalfduizend"},
		{Number: 11000, Value: "elfduizend"},
		{Number: 10000, Value: "tienduizend"},
		{Number: 9000, Value: "negenduizend"},
		{Number: 8000, Value: "achtduizend"},
		{Number: 7000, Value: "zevenduizend"},
		{Number: 6000, Value: "zesduizend"},
		{Number: 5000, Value: "vijfduizend"},
		{Number: 4000, Value: "vierduizend"},
		{Number: 3000, Value: "drieduizend"},
		{Number: 2000, Value: "tweeduizend"},
		{Number: 1000, Value: "duizend"},
		{Number: 900, Value: "negenhonderd"},
		{Number: 800, Value: "achthonderd"},
		{Number: 700, Value: "zevenhonderd"},
		{Number: 600, Value: "zeshonderd"},
		{Number: 500, Value: "vijfhonderd"},
		{Number: 400, Value: "vierhonderd"},
		{Number: 300, Value: "driehonderd"},
		{Number: 200, Value: "tweehonderd"},
		{Number: 100, Value: "honderd"},
		{Number: 99, Value: "negenennegentig"},
		{Number: 98, Value: "achtennegentig"},
		{Number: 97, Value: "zevenennegentig"},
		{Number: 96, Value: "zesennegentig"},
		{Number: 95, Value: "vijfennegentig"},
		{Number: 94, Value: "vierennegentig"},
		{Number: 93, Value: "drieënnegentig"},
		{Number: 92, Value: "tweeënnegentig"},
		{Number: 91, Value: "eenennegentig"},
		{Number: 90, Value: "negentig"},
		{Number: 89, Value: "negenentachtig"},
		{Number: 88, Value: "achtentachtig"},
		{Number: 87, Value: "zevenentachtig"},
		{Number: 86, Value: "zesentachtig"},
		{Number: 85, Value: "vijfentachtig"},
		{Number: 84, Value: "vierentachtig"},
		{Number: 83, Value: "drieëntachtig"},
		{Number: 82, Value: "tweeëntachtig"},
		{Number: 81, Value: "eenentachtig"},
		{Number: 80, Value: "tachtig"},
		{Number: 79, Value: "negenenzeventig"},
		{Number: 78, Value: "achtenenzeventig"},
		{Number: 77, Value: "zevenenzeventig"},
		{Number: 76, Value: "zesenzeventig"},
		{Number: 75, Value: "vijfenzeventig"},
		{Number: 74, Value: "vierenzeventig"},
		{Number: 73, Value: "drieënzeventig"},
		{Number: 72, Value: "tweeënzeventig"},
		{Number: 71, Value: "eenenzeventig"},
		{Number: 70, Value: "zeventig"},
		{Number: 69, Value: "negenenzestig"},
		{Number: 68, Value: "achtenzestig"},
		{Number: 67, Value: "zevenenzestig"},
		{Number: 66, Value: "zesenzestig"},
		{Number: 65, Value: "vijfenzestig"},
		{Number: 64, Value: "vierenzestig"},
		{Number: 63, Value: "drieënzestig"},
		{Number: 62, Value: "tweeënzestig"},
		{Number: 61, Value: "eenenzestig"},
		{Number: 60, Value: "zestig"},
		{Number: 59, Value: "negenenvijftig"},
		{Number: 58, Value: "achtenvijftig"},
		{Number: 57, Value: "zevenenvijftig"},
		{Number: 56, Value: "zesenvijftig"},
		{Number: 55, Value: "vijfenvijftig"},
		{Number: 54, Value: "vierenvijftig"},
		{Number: 53, Value: "drieënvijftig"},
		{Number: 52, Value: "tweeënvijftig"},
		{Number: 51, Value: "eenenvijftig"},
		{Number: 50, Value: "vijftig"},
		{Number: 49, Value: "negenenveertig"},
		{Number: 48, Value: "achtenveertig"},
		{Number: 47, Value: "zevenenveertig"},
		{Number: 46, Value: "zesenveertig"},
		{Number: 45, Value: "vijfenveertig"},
		{Number: 44, Value: "vierenveertig"},
		{Number: 43, Value: "drieënveertig"},
		{Number: 42, Value: "tweeënveertig"},
		{Number: 41, Value: "eenenveertig"},
		{Number: 40, Value: "veertig"},
		{Number: 39, Value: "negenendertig"},
		{Number: 38, Value: "achtendertig"},
		{Number: 37, Value: "zevenendertig"},
		{Number: 36, Value: "zesendertig"},
		{Number: 35, Value: "vijfendertig"},
		{Number: 34, Value: "vierendertig"},
		{Number: 33, Value: "drieëndertig"},
		{Number: 32, Value: "tweeëndertig"},
		{Number: 31, Value: "eenendertig"},
		{Number: 30, Value: "dertig"},
		{Number: 29, Value: "negenentwintig"},
		{Number: 28, Value: "achtentwintig"},
		{Number: 27, Value: "zevenentwintig"},
		{Number: 26, Value: "zesentwintig"},
		{Number: 25, Value: "vijfentwintig"},
		{Number: 24, Value: "vierentwintig"},
		{Number: 23, Value: "drieëntwintig"},
		{Number: 22, Value: "tweeëntwintig"},
		{Number: 21, Value: "eenentwintig"},
		{Number: 20, Value: "twintig"},
		{Number: 19, Value: "negentien"},
		{Number: 18, Value: "achttien"},
		{Number: 17, Value: "zeventien"},
		{Number: 16, Value: "zestien"},
		{Number: 15, Value: "vijftien"},
		{Number: 14, Value: "veertien"},
		{Number: 13, Value: "dertien"},
		{Number: 12, Value: "twaalf"},
		{Number: 11, Value: "elf"},
		{Number: 10, Value: "tien"},
		{Number: 9, Value: "negen"},
		{Number: 8, Value: "acht"},
		{Number: 7, Value: "zeven"},
		{Number: 6, Value: "zes"},
		{Number: 5, Value: "vijf"},
		{Number: 4, Value: "vier"},
		{Number: 3, Value: "drie"},
		{Number: 2, Value: "twee"},
		{Number: 1, Value: "een"},
		{Number: 0, Value: "nul"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Eenhonderd"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "eerste", Suffix: "-ste", Masculine: "eerste", Feminine: "eerste", Neuter: "eerste"},
		{Number: 2, Word: "tweede", Suffix: "-de", Masculine: "tweede", Feminine: "tweede", Neuter: "tweede"},
		{Number: 3, Word: "derde", Suffix: "-de", Masculine: "derde", Feminine: "derde", Neuter: "derde"},
		{Number: 4, Word: "vierde", Suffix: "-de", Masculine: "vierde", Feminine: "vierde", Neuter: "vierde"},
		{Number: 5, Word: "vijfde", Suffix: "-de", Masculine: "vijfde", Feminine: "vijfde", Neuter: "vijfde"},
		{Number: 6, Word: "zesde", Suffix: "-de", Masculine: "zesde", Feminine: "zesde", Neuter: "zesde"},
		{Number: 7, Word: "zevende", Suffix: "-de", Masculine: "zevende", Feminine: "zevende", Neuter: "zevende"},
		{Number: 8, Word: "achtste", Suffix: "-ste", Masculine: "achtste", Feminine: "achtste", Neuter: "achtste"},
		{Number: 9, Word: "negende", Suffix: "-de", Masculine: "negende", Feminine: "negende", Neuter: "negende"},
		{Number: 10, Word: "tiende", Suffix: "-de", Masculine: "tiende", Feminine: "tiende", Neuter: "tiende"},
		{Number: 11, Word: "elfde", Suffix: "-de", Masculine: "elfde", Feminine: "elfde", Neuter: "elfde"},
		{Number: 12, Word: "twaalfde", Suffix: "-de", Masculine: "twaalfde", Feminine: "twaalfde", Neuter: "twaalfde"},
		{Number: 20, Word: "twintigste", Suffix: "-ste", Masculine: "twintigste", Feminine: "twintigste", Neuter: "twintigste"},
		{Number: 21, Word: "eenentwintigste", Suffix: "-ste", Masculine: "eenentwintigste", Feminine: "eenentwintigste", Neuter: "eenentwintigste"},
		{Number: 30, Word: "dertigste", Suffix: "-ste", Masculine: "dertigste", Feminine: "dertigste", Neuter: "dertigste"},
		{Number: 100, Word: "honderdste", Suffix: "-ste", Masculine: "honderdste", Feminine: "honderdste", Neuter: "honderdste"},
		{Number: 1000, Word: "duizendste", Suffix: "-ste", Masculine: "duizendste", Feminine: "duizendste", Neuter: "duizendste"},
	},
	LocaleFormatter: &DutchBelgianFormatter{},
}

// DutchBelgianFormatter handles Dutch (Belgium)-specific formatting
type DutchBelgianFormatter struct{}

func (f *DutchBelgianFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return f.convertDutchNumber(number, targetLocale)
}

func (f *DutchBelgianFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *DutchBelgianFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *DutchBelgianFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *DutchBelgianFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *DutchBelgianFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 0
	}
	return amount.Truncate(int32(precision))
}

func (f *DutchBelgianFormatter) FormatDecimalNumber(amount float64) string {
	return FormatEuropeanDecimal(amount)
}

func (f *DutchBelgianFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with European conventions (period separators, comma decimal, prefix symbol)
	return FormatEuropeanCurrency(amount, currencySymbol)
}

// convertDutchNumber converts numbers to Dutch compound words
func (f *DutchBelgianFormatter) convertDutchNumber(number int64, targetLocale NumI18NLocale) string {
	if number == 0 {
		return "nul"
	}

	if number < 0 {
		return "min " + f.convertDutchNumber(-number, targetLocale)
	}

	// Handle numbers 1-19 directly
	if number <= 19 {
		for _, mapping := range targetLocale.NumberWordsMapping {
			if mapping.Number == number {
				return strings.ToLower(mapping.Value)
			}
		}
	}

	// Handle 20-99 (compound tens - Dutch puts ones before tens with "en")
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

		return onesWord + "en" + tensWord
	}

	// Handle 100-999 (hundreds)
	if number < 1000 {
		hundreds := number / 100
		remainder := number % 100

		result := ""
		if hundreds == 1 {
			result = "honderd"
		} else {
			hundredsWord := f.convertDutchNumber(hundreds, targetLocale)
			result = hundredsWord + "honderd"
		}

		if remainder > 0 {
			result += f.convertDutchNumber(remainder, targetLocale)
		}

		return result
	}

	// Handle 1000+ (thousands)
	if number < 1000000 {
		thousands := number / 1000
		remainder := number % 1000

		result := ""
		if thousands == 1 {
			result = "duizend"
		} else {
			thousandsWord := f.convertDutchNumber(thousands, targetLocale)
			result = thousandsWord + "duizend"
		}

		if remainder > 0 {
			result += f.convertDutchNumber(remainder, targetLocale)
		}

		return result
	}

	// For larger numbers, fall back to generic conversion but lowercase
	return strings.ToLower(ConvertToWordsWithExactMappingInt64(number, targetLocale))
}
