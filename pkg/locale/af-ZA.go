package locale

import (
	"strings"
	"github.com/shopspring/decimal"
)

// AFZALocale is a NumI18NLocale configured for Afrikaans (South Africa) - af-ZA
var AFZALocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Rand",
		Plural:   "Rande",
		Singular: "Rand",
		Symbol:   "R",
		FractionUnit: FractionUnit{
			Name:     "Sent",
			Plural:   "Sente",
			Singular: "Sent",
			Symbol:   "c",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "South Africa",
		Currency:       "ZAR",
		ISO3166Alpha2:  "ZA",
		ISO3166Alpha3:  "ZAF",
		ISO3166Numeric: "710",
		Locale:         "af-ZA",
		Timezone:       []string{"Africa/Johannesburg"},
		Language:       "af",
	},
	Texts: Texts{
		And:   "En",
		Minus: "Min",
		Only:  "Slegs",
		Point: "Komma",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Kwadrieljoen"},
		{Number: 1000000000000, Value: "Triljoen"},
		{Number: 1000000000, Value: "Miljard"},
		{Number: 1000000, Value: "Miljoen"},
		{Number: 1000, Value: "Duizend"},
		{Number: 100, Value: "Honderd"},
		{Number: 90, Value: "Negentig"},
		{Number: 80, Value: "Tagtig"},
		{Number: 70, Value: "Sewentig"},
		{Number: 60, Value: "Sestig"},
		{Number: 50, Value: "Vyftig"},
		{Number: 40, Value: "Veertig"},
		{Number: 30, Value: "Dertig"},
		{Number: 20, Value: "Twintig"},
		{Number: 19, Value: "Negentien"},
		{Number: 18, Value: "Agtien"},
		{Number: 17, Value: "Sewentien"},
		{Number: 16, Value: "Sestien"},
		{Number: 15, Value: "Vyftien"},
		{Number: 14, Value: "Veertien"},
		{Number: 13, Value: "Dertien"},
		{Number: 12, Value: "Twaalf"},
		{Number: 11, Value: "Elf"},
		{Number: 10, Value: "Tien"},
		{Number: 9, Value: "Nege"},
		{Number: 8, Value: "Agt"},
		{Number: 7, Value: "Sewe"},
		{Number: 6, Value: "Ses"},
		{Number: 5, Value: "Vyf"},
		{Number: 4, Value: "Vier"},
		{Number: 3, Value: "Drie"},
		{Number: 2, Value: "Twee"},
		{Number: 1, Value: "Een"},
		{Number: 0, Value: "Nul"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 1000000, Value: "Een Miljoen"},
		{Number: 1000, Value: "Een Duizend"},
		{Number: 100, Value: "Honderd"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "Eerste", Suffix: "de", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 2, Word: "Tweede", Suffix: "de", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 3, Word: "Derde", Suffix: "de", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 4, Word: "Vierde", Suffix: "de", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 5, Word: "Vyfde", Suffix: "de", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 6, Word: "Sesde", Suffix: "de", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 7, Word: "Sewende", Suffix: "de", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 8, Word: "Agtste", Suffix: "ste", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 9, Word: "Negende", Suffix: "de", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 10, Word: "Tiende", Suffix: "de", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 11, Word: "Elfde", Suffix: "de", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 12, Word: "Twaalfde", Suffix: "de", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 13, Word: "Dertiende", Suffix: "de", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 14, Word: "Veertiende", Suffix: "de", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 15, Word: "Vyftiende", Suffix: "de", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 16, Word: "Sestiende", Suffix: "de", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 17, Word: "Sewentiende", Suffix: "de", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 18, Word: "Agtiende", Suffix: "de", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 19, Word: "Negentiende", Suffix: "de", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 20, Word: "Twintigste", Suffix: "ste", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 21, Word: "Een-en-twintigste", Suffix: "ste", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 30, Word: "Dertigste", Suffix: "ste", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 40, Word: "Veertigste", Suffix: "ste", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 50, Word: "Vyftigste", Suffix: "ste", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 60, Word: "Sestigste", Suffix: "ste", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 70, Word: "Sewentigste", Suffix: "ste", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 80, Word: "Tagtigste", Suffix: "ste", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 90, Word: "Negentigste", Suffix: "ste", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 100, Word: "Honderdste", Suffix: "ste", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 1000, Word: "Duizendste", Suffix: "ste", Masculine: "", Feminine: "", Neuter: ""},
	},
	LocaleFormatter: &AfrikaansFormatter{},
}

// AfrikaansFormatter handles Afrikaans formatting
type AfrikaansFormatter struct{}

func (f *AfrikaansFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	// Convert to decimal for consistent handling
	decNumber := decimal.NewFromInt(number)

	// Handle zero
	if decNumber.Equal(decimal.Zero) {
		return GetWordForNumber(decimal.Zero, targetLocale)
	}

	// Check for exact mappings first
	for _, mapping := range targetLocale.ExactWordsMapping {
		if decimal.NewFromInt(mapping.Number).Equal(decNumber) {
			return mapping.Value
		}
	}

	result := ""
	million := decimal.NewFromInt(1000000)
	thousand := decimal.NewFromInt(1000)
	hundred := decimal.NewFromInt(100)
	twentyOne := decimal.NewFromInt(21)
	ninety9 := decimal.NewFromInt(99)
	ten := decimal.NewFromInt(10)
	one := decimal.NewFromInt(1)

	// Handle millions and above
	if decNumber.GreaterThanOrEqual(million) {
		millions := decNumber.Div(million).Floor()
		if millions.Equal(one) {
			result += "Een " + GetWordForNumber(million, targetLocale)
		} else {
			result += f.FormatNumber(millions.IntPart(), targetLocale) + " " + GetWordForNumber(million, targetLocale)
		}
		remainder := decNumber.Mod(million)
		if remainder.GreaterThan(decimal.Zero) {
			result += " " + f.FormatNumber(remainder.IntPart(), targetLocale)
		}
		return result
	}

	// Handle thousands
	if decNumber.GreaterThanOrEqual(thousand) {
		thousands := decNumber.Div(thousand).Floor()
		if thousands.Equal(one) {
			result += "Een " + GetWordForNumber(thousand, targetLocale)
		} else {
			result += f.FormatNumber(thousands.IntPart(), targetLocale) + " " + GetWordForNumber(thousand, targetLocale)
		}
		remainder := decNumber.Mod(thousand)
		if remainder.GreaterThan(decimal.Zero) {
			result += " " + f.FormatNumber(remainder.IntPart(), targetLocale)
		}
		return result
	}

	// Handle hundreds
	if decNumber.GreaterThanOrEqual(hundred) {
		hundreds := decNumber.Div(hundred).Floor()
		if hundreds.Equal(one) {
			result += GetWordForNumber(hundred, targetLocale)
		} else {
			result += GetWordForNumber(hundreds, targetLocale) + " " + GetWordForNumber(hundred, targetLocale)
		}
		remainder := decNumber.Mod(hundred)
		if remainder.GreaterThan(decimal.Zero) {
			result += " " + f.FormatNumber(remainder.IntPart(), targetLocale)
		}
		return result
	}
	if decNumber.GreaterThanOrEqual(twentyOne) && decNumber.LessThanOrEqual(ninety9) {
		tens := decNumber.Div(ten).Floor().Mul(ten)
		ones := decNumber.Mod(ten)
		if ones.Equal(decimal.Zero) {
			return GetWordForNumber(tens, targetLocale)
		}
		return GetWordForNumber(ones, targetLocale) + " " + targetLocale.Texts.And + " " + GetWordForNumber(tens, targetLocale)
	}
	return GetWordForNumber(decNumber, targetLocale)
}

func (f *AfrikaansFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *AfrikaansFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *AfrikaansFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *AfrikaansFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *AfrikaansFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 2
	}
	return amount.Truncate(int32(precision))
}


func (f *AfrikaansFormatter) FormatDecimalNumber(amount float64) string {
	return DefaultFormatDecimalNumber(amount, ",", ".")
}
func (f *AfrikaansFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
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
