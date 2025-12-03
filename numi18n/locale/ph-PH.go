package locale

import (
	"strings"

	"github.com/shopspring/decimal"
)

// PHPHLocale is a NumI18NLocale configured for Filipino (Philippines) - ph-PH
var PHPHLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Peso",
		Plural:   "Piso",
		Singular: "Peso",
		Symbol:   "₱",
		FractionUnit: FractionUnit{
			Name:     "Sentimo",
			Plural:   "Sentimo",
			Singular: "Sentimo",
			Symbol:   "¢",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Pilipinas",
		Currency:       "PHP",
		ISO3166Alpha2:  "PH",
		ISO3166Alpha3:  "PHL",
		ISO3166Numeric: "608",
		Locale:         "ph-PH",
		Timezone:       []string{"Asia/Manila"},
		Language:       "ph",
		Emoji:          "🇵🇭",
	},
	Texts: Texts{
		And:   "at",
		Minus: "minus",
		Only:  "lamang",
		Point: "punto",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Katrilyun"},
		{Number: 1000000000000, Value: "Trilyun"},
		{Number: 1000000000, Value: "Bilyun"},
		{Number: 1000000, Value: "Milyun"},
		{Number: 1000, Value: "Libo"},
		{Number: 100, Value: "Daan"},
		{Number: 90, Value: "Siyamnapu"},
		{Number: 80, Value: "Walumpu"},
		{Number: 70, Value: "Pitumpu"},
		{Number: 60, Value: "Animnapu"},
		{Number: 50, Value: "Limampu"},
		{Number: 40, Value: "Apatnapu"},
		{Number: 30, Value: "Tatlong pu"},
		{Number: 20, Value: "Dalawampu"},
		{Number: 19, Value: "Labinsiyam"},
		{Number: 18, Value: "Labingwalo"},
		{Number: 17, Value: "Labimpito"},
		{Number: 16, Value: "Labing-anim"},
		{Number: 15, Value: "Labinlima"},
		{Number: 14, Value: "Labing-apat"},
		{Number: 13, Value: "Labintatlo"},
		{Number: 12, Value: "Labindalawa"},
		{Number: 11, Value: "Labing-isa"},
		{Number: 10, Value: "Sampu"},
		{Number: 9, Value: "Siyam"},
		{Number: 8, Value: "Walo"},
		{Number: 7, Value: "Pito"},
		{Number: 6, Value: "Anim"},
		{Number: 5, Value: "Lima"},
		{Number: 4, Value: "Apat"},
		{Number: 3, Value: "Tatlo"},
		{Number: 2, Value: "Dalawa"},
		{Number: 1, Value: "Isa"},
		{Number: 0, Value: "Zero"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Isang daan"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "Una", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 2, Word: "Ikalawa", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 3, Word: "Ikatlo", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 4, Word: "Ikaapat", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 5, Word: "Ikalima", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 6, Word: "Ikaanim", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 7, Word: "Ikapito", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 8, Word: "Ikawalo", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 9, Word: "Ikasiyam", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 10, Word: "Ikasampu", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 11, Word: "Ikalabing-isa", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 12, Word: "Ikalabindalawa", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 13, Word: "Ikalabintatlo", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 14, Word: "Ikalabing-apat", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 15, Word: "Ikalabinlima", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 16, Word: "Ikalabing-anim", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 17, Word: "Ikalabimpito", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 18, Word: "Ikalabingwalo", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 19, Word: "Ikalabinsiyam", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 20, Word: "Ikadalawampu", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 21, Word: "Ikadalawampu't isa", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 30, Word: "Ikatatlumpu", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 40, Word: "Ikaapatnapu", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 50, Word: "Ikalimampu", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 100, Word: "Ikadaan", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 1000, Word: "Ikalibo", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
		{Number: 1000000, Word: "Ikamilyun", Suffix: "-ng", Masculine: "", Feminine: "", Neuter: ""},
	},
	LocaleFormatter: &FilipinoFormatter{},
}

// FilipinoFormatter handles Filipino-specific formatting
type FilipinoFormatter struct{}

func (f *FilipinoFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	decimalNumber := decimal.NewFromInt(number)
	if decimalNumber.Equal(decimal.Zero) {
		return GetWordForNumber(decimal.Zero, targetLocale)
	}

	result := ""
	billion := decimal.NewFromInt(1000000000)
	million := decimal.NewFromInt(1000000)
	thousand := decimal.NewFromInt(1000)
	hundred := decimal.NewFromInt(100)

	// Handle large numbers
	result, decimalNumber = f.handleBillions(result, decimalNumber, billion, targetLocale)
	result, decimalNumber = f.handleMillions(result, decimalNumber, million, targetLocale)
	result, decimalNumber = f.handleThousands(result, decimalNumber, thousand, targetLocale)
	result, decimalNumber = f.handleHundreds(result, decimalNumber, hundred, targetLocale)

	// Handle remaining numbers (tens and ones)
	result = f.handleTensAndOnes(result, decimalNumber, targetLocale)

	return strings.TrimSpace(result)
}

func (f *FilipinoFormatter) handleBillions(result string, decimalNumber, billion decimal.Decimal, targetLocale NumI18NLocale) (string, decimal.Decimal) {
	if decimalNumber.GreaterThanOrEqual(billion) {
		billions := decimalNumber.Div(billion).Floor()
		if billions.Equal(decimal.NewFromInt(1)) {
			result += "Isang " + GetWordForNumber(billion, targetLocale)
		} else {
			result += f.FormatNumber(billions.IntPart(), targetLocale) + " " + GetWordForNumber(billion, targetLocale)
		}
		decimalNumber = decimalNumber.Mod(billion)
		if decimalNumber.GreaterThan(decimal.Zero) {
			result += " "
		}
	}
	return result, decimalNumber
}

func (f *FilipinoFormatter) handleMillions(result string, decimalNumber, million decimal.Decimal, targetLocale NumI18NLocale) (string, decimal.Decimal) {
	if decimalNumber.GreaterThanOrEqual(million) {
		millions := decimalNumber.Div(million).Floor()
		if millions.Equal(decimal.NewFromInt(1)) {
			result += "Isang " + GetWordForNumber(million, targetLocale)
		} else {
			result += f.FormatNumber(millions.IntPart(), targetLocale) + " " + GetWordForNumber(million, targetLocale)
		}
		decimalNumber = decimalNumber.Mod(million)
		if decimalNumber.GreaterThan(decimal.Zero) {
			result += " "
		}
	}
	return result, decimalNumber
}

func (f *FilipinoFormatter) handleThousands(result string, decimalNumber, thousand decimal.Decimal, targetLocale NumI18NLocale) (string, decimal.Decimal) {
	if decimalNumber.GreaterThanOrEqual(thousand) {
		thousands := decimalNumber.Div(thousand).Floor()
		if thousands.Equal(decimal.NewFromInt(1)) {
			result += "Isang " + GetWordForNumber(thousand, targetLocale)
		} else {
			result += f.FormatNumber(thousands.IntPart(), targetLocale) + " " + GetWordForNumber(thousand, targetLocale)
		}
		decimalNumber = decimalNumber.Mod(thousand)
		if decimalNumber.GreaterThan(decimal.Zero) {
			result += " "
		}
	}
	return result, decimalNumber
}

func (f *FilipinoFormatter) handleHundreds(result string, decimalNumber, hundred decimal.Decimal, targetLocale NumI18NLocale) (string, decimal.Decimal) {
	if decimalNumber.GreaterThanOrEqual(hundred) {
		hundreds := decimalNumber.Div(hundred).Floor()
		if hundreds.Equal(decimal.NewFromInt(1)) {
			// Use exact mapping for "Isang daan"
			for _, mapping := range targetLocale.ExactWordsMapping {
				if mapping.Number == 100 {
					result += mapping.Value
					break
				}
			}
		} else {
			result += f.FormatNumber(hundreds.IntPart(), targetLocale) + " " + GetWordForNumber(hundred, targetLocale)
		}
		decimalNumber = decimalNumber.Mod(hundred)
		if decimalNumber.GreaterThan(decimal.Zero) {
			result += " "
		}
	}
	return result, decimalNumber
}

func (f *FilipinoFormatter) handleTensAndOnes(result string, decimalNumber decimal.Decimal, targetLocale NumI18NLocale) string {
	twenty := decimal.NewFromInt(20)
	ten := decimal.NewFromInt(10)

	if decimalNumber.GreaterThanOrEqual(twenty) {
		tens := decimalNumber.Div(ten).Floor()
		tensNumber := tens.Mul(ten)
		result += GetWordForNumber(tensNumber, targetLocale)
		decimalNumber = decimalNumber.Mod(ten)

		result = f.handleOnesAfterTens(result, decimalNumber, targetLocale)
	} else if decimalNumber.GreaterThan(decimal.Zero) {
		result = f.handleSingleDigits(result, decimalNumber, targetLocale)
	}

	return result
}

func (f *FilipinoFormatter) handleOnesAfterTens(result string, decimalNumber decimal.Decimal, targetLocale NumI18NLocale) string {
	one := decimal.NewFromInt(1)

	if decimalNumber.GreaterThan(decimal.Zero) {
		if decimalNumber.Equal(one) {
			result += " Isa"
		} else {
			found := f.findExactMapping(decimalNumber, targetLocale)
			if found != "" {
				result += " " + found
			} else {
				result += " " + GetWordForNumber(decimalNumber, targetLocale)
			}
		}
	}
	return result
}

func (f *FilipinoFormatter) handleSingleDigits(result string, decimalNumber decimal.Decimal, targetLocale NumI18NLocale) string {
	found := f.findExactMapping(decimalNumber, targetLocale)
	if found != "" {
		result += found
	} else {
		result += GetWordForNumber(decimalNumber, targetLocale)
	}
	return result
}

func (f *FilipinoFormatter) findExactMapping(decimalNumber decimal.Decimal, targetLocale NumI18NLocale) string {
	for _, mapping := range targetLocale.ExactWordsMapping {
		if decimal.NewFromInt(mapping.Number).Equal(decimalNumber) {
			return mapping.Value
		}
	}
	return ""
}

func (f *FilipinoFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *FilipinoFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *FilipinoFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *FilipinoFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *FilipinoFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 2
	}
	return amount.Truncate(int32(precision))
}

func (f *FilipinoFormatter) FormatDecimalNumber(amount float64) string {
	return FormatDecimalWithSeparators(amount, ",")
}

func (f *FilipinoFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with thousand separators, currency symbol, and symbol prefix (true for Philippines) - always shows 2 decimal places
	return FormatUSCurrencyWithSeparators(amount, ",", currencySymbol, true)
}
