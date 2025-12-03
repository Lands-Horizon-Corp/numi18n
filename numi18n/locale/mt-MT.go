package locale

import (
	"github.com/shopspring/decimal"
)

// MTMTLocale represents the Maltese (Malta) locale
var MTMTLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Euro",
		Plural:   "Ewro",
		Singular: "Ewro",
		Symbol:   "€",
		FractionUnit: FractionUnit{
			Name:     "Ċenteżmu",
			Plural:   "Ċenteżmi",
			Singular: "Ċenteżmu",
			Symbol:   "c",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Malta",
		Currency:       "EUR",
		ISO3166Alpha2:  "MT",
		ISO3166Alpha3:  "MLT",
		ISO3166Numeric: "470",
		Locale:         "mt-MT",
		Timezone:       []string{"Europe/Malta"},
		Language:       "mt",
		Emoji:          "🇲🇹",
		PhoneCode:      "+356",
		Domain:         ".mt",
	},
	Texts: Texts{
		And:   "u",
		Minus: "inqas",
		Only:  "biss",
		Point: "punt",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "biljun ta' biljuni"},
		{Number: 1000000000000, Value: "biljun"},
		{Number: 1000000000, Value: "biljun żgħir"},
		{Number: 1000000, Value: "miljun"},
		{Number: 100000, Value: "mija elf"},
		{Number: 90000, Value: "disgħin elf"},
		{Number: 80000, Value: "tmenin elf"},
		{Number: 70000, Value: "sebgħin elf"},
		{Number: 60000, Value: "sittin elf"},
		{Number: 50000, Value: "ħamsin elf"},
		{Number: 40000, Value: "erbgħin elf"},
		{Number: 30000, Value: "tletin elf"},
		{Number: 20000, Value: "għoxrin elf"},
		{Number: 19000, Value: "dsatax-il elf"},
		{Number: 18000, Value: "tmintax-il elf"},
		{Number: 17000, Value: "sbatax-il elf"},
		{Number: 16000, Value: "sittax-il elf"},
		{Number: 15000, Value: "ħmistax-il elf"},
		{Number: 14000, Value: "erbatax-il elf"},
		{Number: 13000, Value: "tlettax-il elf"},
		{Number: 12000, Value: "tnax-il elf"},
		{Number: 11000, Value: "ħdax-il elf"},
		{Number: 10000, Value: "għaxar eluf"},
		{Number: 9000, Value: "disgħa eluf"},
		{Number: 8000, Value: "tmien eluf"},
		{Number: 7000, Value: "seba' eluf"},
		{Number: 6000, Value: "sitt eluf"},
		{Number: 5000, Value: "ħamsa eluf"},
		{Number: 4000, Value: "erba' eluf"},
		{Number: 3000, Value: "tlett eluf"},
		{Number: 2000, Value: "elfejn"},
		{Number: 1000, Value: "elf"},
		{Number: 900, Value: "disgħa mija"},
		{Number: 800, Value: "tmien mija"},
		{Number: 700, Value: "seba' mija"},
		{Number: 600, Value: "sitt mija"},
		{Number: 500, Value: "ħames mija"},
		{Number: 400, Value: "erba' mija"},
		{Number: 300, Value: "tlett mija"},
		{Number: 200, Value: "mitejn"},
		{Number: 100, Value: "mija"},
		{Number: 99, Value: "disgħin u disgħa"},
		{Number: 98, Value: "disgħin u tmienja"},
		{Number: 97, Value: "disgħin u sebgħa"},
		{Number: 96, Value: "disgħin u sitta"},
		{Number: 95, Value: "disgħin u ħamsa"},
		{Number: 94, Value: "disgħin u erbgħa"},
		{Number: 93, Value: "disgħin u tlieta"},
		{Number: 92, Value: "disgħin u tnejn"},
		{Number: 91, Value: "disgħin u wieħed"},
		{Number: 90, Value: "disgħin"},
		{Number: 89, Value: "tmenin u disgħa"},
		{Number: 88, Value: "tmenin u tmienja"},
		{Number: 87, Value: "tmenin u sebgħa"},
		{Number: 86, Value: "tmenin u sitta"},
		{Number: 85, Value: "tmenin u ħamsa"},
		{Number: 84, Value: "tmenin u erbgħa"},
		{Number: 83, Value: "tmenin u tlieta"},
		{Number: 82, Value: "tmenin u tnejn"},
		{Number: 81, Value: "tmenin u wieħed"},
		{Number: 80, Value: "tmenin"},
		{Number: 79, Value: "sebgħin u disgħa"},
		{Number: 78, Value: "sebgħin u tmienja"},
		{Number: 77, Value: "sebgħin u sebgħa"},
		{Number: 76, Value: "sebgħin u sitta"},
		{Number: 75, Value: "sebgħin u ħamsa"},
		{Number: 74, Value: "sebgħin u erbgħa"},
		{Number: 73, Value: "sebgħin u tlieta"},
		{Number: 72, Value: "sebgħin u tnejn"},
		{Number: 71, Value: "sebgħin u wieħed"},
		{Number: 70, Value: "sebgħin"},
		{Number: 69, Value: "sittin u disgħa"},
		{Number: 68, Value: "sittin u tmienja"},
		{Number: 67, Value: "sittin u sebgħa"},
		{Number: 66, Value: "sittin u sitta"},
		{Number: 65, Value: "sittin u ħamsa"},
		{Number: 64, Value: "sittin u erbgħa"},
		{Number: 63, Value: "sittin u tlieta"},
		{Number: 62, Value: "sittin u tnejn"},
		{Number: 61, Value: "sittin u wieħed"},
		{Number: 60, Value: "sittin"},
		{Number: 59, Value: "ħamsin u disgħa"},
		{Number: 58, Value: "ħamsin u tmienja"},
		{Number: 57, Value: "ħamsin u sebgħa"},
		{Number: 56, Value: "ħamsin u sitta"},
		{Number: 55, Value: "ħamsin u ħamsa"},
		{Number: 54, Value: "ħamsin u erbgħa"},
		{Number: 53, Value: "ħamsin u tlieta"},
		{Number: 52, Value: "ħamsin u tnejn"},
		{Number: 51, Value: "ħamsin u wieħed"},
		{Number: 50, Value: "ħamsin"},
		{Number: 49, Value: "erbgħin u disgħa"},
		{Number: 48, Value: "erbgħin u tmienja"},
		{Number: 47, Value: "erbgħin u sebgħa"},
		{Number: 46, Value: "erbgħin u sitta"},
		{Number: 45, Value: "erbgħin u ħamsa"},
		{Number: 44, Value: "erbgħin u erbgħa"},
		{Number: 43, Value: "erbgħin u tlieta"},
		{Number: 42, Value: "erbgħin u tnejn"},
		{Number: 41, Value: "erbgħin u wieħed"},
		{Number: 40, Value: "erbgħin"},
		{Number: 39, Value: "tletin u disgħa"},
		{Number: 38, Value: "tletin u tmienja"},
		{Number: 37, Value: "tletin u sebgħa"},
		{Number: 36, Value: "tletin u sitta"},
		{Number: 35, Value: "tletin u ħamsa"},
		{Number: 34, Value: "tletin u erbgħa"},
		{Number: 33, Value: "tletin u tlieta"},
		{Number: 32, Value: "tletin u tnejn"},
		{Number: 31, Value: "tletin u wieħed"},
		{Number: 30, Value: "tletin"},
		{Number: 29, Value: "għoxrin u disgħa"},
		{Number: 28, Value: "għoxrin u tmienja"},
		{Number: 27, Value: "għoxrin u sebgħa"},
		{Number: 26, Value: "għoxrin u sitta"},
		{Number: 25, Value: "għoxrin u ħamsa"},
		{Number: 24, Value: "għoxrin u erbgħa"},
		{Number: 23, Value: "għoxrin u tlieta"},
		{Number: 22, Value: "għoxrin u tnejn"},
		{Number: 21, Value: "għoxrin u wieħed"},
		{Number: 20, Value: "għoxrin"},
		{Number: 19, Value: "dsatax"},
		{Number: 18, Value: "tmintax"},
		{Number: 17, Value: "sbatax"},
		{Number: 16, Value: "sittax"},
		{Number: 15, Value: "ħmistax"},
		{Number: 14, Value: "erbatax"},
		{Number: 13, Value: "tlettax"},
		{Number: 12, Value: "tnax"},
		{Number: 11, Value: "ħdax"},
		{Number: 10, Value: "għaxar"},
		{Number: 9, Value: "disgħa"},
		{Number: 8, Value: "tmienja"},
		{Number: 7, Value: "sebgħa"},
		{Number: 6, Value: "sitta"},
		{Number: 5, Value: "ħamsa"},
		{Number: 4, Value: "erbgħa"},
		{Number: 3, Value: "tlieta"},
		{Number: 2, Value: "tnejn"},
		{Number: 1, Value: "wieħed"},
		{Number: 0, Value: "żero"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Mija waħda"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "l-ewwel", Suffix: "-ewwel", Masculine: "l-ewwel", Feminine: "l-ewwel", Neuter: "l-ewwel"},
		{Number: 2, Word: "it-tieni", Suffix: "-tieni", Masculine: "it-tieni", Feminine: "it-tieni", Neuter: "it-tieni"},
		{Number: 3, Word: "it-tielet", Suffix: "-tielet", Masculine: "it-tielet", Feminine: "it-tielet", Neuter: "it-tielet"},
		{Number: 4, Word: "ir-raba'", Suffix: "-raba'", Masculine: "ir-raba'", Feminine: "ir-raba'", Neuter: "ir-raba'"},
		{Number: 5, Word: "il-ħames", Suffix: "-ħames", Masculine: "il-ħames", Feminine: "il-ħames", Neuter: "il-ħames"},
		{Number: 6, Word: "is-sitt", Suffix: "-sitt", Masculine: "is-sitt", Feminine: "is-sitt", Neuter: "is-sitt"},
		{Number: 7, Word: "is-seba'", Suffix: "-seba'", Masculine: "is-seba'", Feminine: "is-seba'", Neuter: "is-seba'"},
		{Number: 8, Word: "it-tmien", Suffix: "-tmien", Masculine: "it-tmien", Feminine: "it-tmien", Neuter: "it-tmien"},
		{Number: 9, Word: "id-disgħa", Suffix: "-disgħa", Masculine: "id-disgħa", Feminine: "id-disgħa", Neuter: "id-disgħa"},
		{Number: 10, Word: "l-għaxar", Suffix: "-għaxar", Masculine: "l-għaxar", Feminine: "l-għaxar", Neuter: "l-għaxar"},
		{Number: 11, Word: "il-ħdax", Suffix: "-ħdax", Masculine: "il-ħdax", Feminine: "il-ħdax", Neuter: "il-ħdax"},
		{Number: 12, Word: "it-tnax", Suffix: "-tnax", Masculine: "it-tnax", Feminine: "it-tnax", Neuter: "it-tnax"},
		{Number: 20, Word: "l-għoxrin", Suffix: "-għoxrin", Masculine: "l-għoxrin", Feminine: "l-għoxrin", Neuter: "l-għoxrin"},
		{Number: 21, Word: "l-wieħed u għoxrin", Suffix: "-wieħed u għoxrin", Masculine: "l-wieħed u għoxrin", Feminine: "l-wieħed u għoxrin", Neuter: "l-wieħed u għoxrin"},
		{Number: 30, Word: "it-tletin", Suffix: "-tletin", Masculine: "it-tletin", Feminine: "it-tletin", Neuter: "it-tletin"},
		{Number: 100, Word: "il-mija", Suffix: "-mija", Masculine: "il-mija", Feminine: "il-mija", Neuter: "il-mija"},
		{Number: 1000, Word: "l-elf", Suffix: "-elf", Masculine: "l-elf", Feminine: "l-elf", Neuter: "l-elf"},
	},
	LocaleFormatter: &MalteseFormatter{},
}

// MalteseFormatter handles Maltese-specific formatting
type MalteseFormatter struct{}

func (f *MalteseFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *MalteseFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *MalteseFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *MalteseFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *MalteseFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *MalteseFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 0
	}
	return amount.Truncate(int32(precision))
}

func (f *MalteseFormatter) FormatDecimalNumber(amount float64) string {
	return FormatEuropeanDecimal(amount)
}
func (f *MalteseFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with European conventions (period separators, comma decimal, prefix symbol)
	return FormatEuropeanCurrency(amount, currencySymbol)
}
