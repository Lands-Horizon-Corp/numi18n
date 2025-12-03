package locale

import (
	"github.com/shopspring/decimal"
)

// ISLocale is a NumI18NLocale configured for Iceland (is-IS)
var ISLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Króna",
		Plural:   "Krónur",
		Singular: "Króna",
		Symbol:   "kr",
		FractionUnit: FractionUnit{
			Name:     "Eyrir",
			Plural:   "Aurar",
			Singular: "Eyrir",
			Symbol:   "a",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Iceland",
		Currency:       "ISK",
		ISO3166Alpha2:  "IS",
		ISO3166Alpha3:  "ISL",
		ISO3166Numeric: "352",
		Locale:         "is-IS",
		Timezone:       []string{"Atlantic/Reykjavik"},
		Language:       "is",
		Emoji:          "🇮🇸",
		PhoneCode:      "+354",
		Domain:         ".is",
	},
	Texts: Texts{
		And:   "Og",
		Minus: "Mínus",
		Only:  "Aðeins",
		Point: "Komma",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Trilljón"},
		{Number: 1000000000000, Value: "Billjón"},
		{Number: 1000000000, Value: "Milljarður"},
		{Number: 1000000, Value: "Milljón"},
		{Number: 1000, Value: "Þúsund"},
		{Number: 100, Value: "Hundrað"},
		{Number: 90, Value: "Níutíu"},
		{Number: 80, Value: "Áttatíu"},
		{Number: 70, Value: "Sjötíu"},
		{Number: 60, Value: "Sextíu"},
		{Number: 50, Value: "Fimmtíu"},
		{Number: 40, Value: "Fjörutíu"},
		{Number: 30, Value: "Þrjátíu"},
		{Number: 20, Value: "Tuttugu"},
		{Number: 19, Value: "Nítján"},
		{Number: 18, Value: "Átján"},
		{Number: 17, Value: "Sautján"},
		{Number: 16, Value: "Sextán"},
		{Number: 15, Value: "Fimmtán"},
		{Number: 14, Value: "Fjórtán"},
		{Number: 13, Value: "Þrettán"},
		{Number: 12, Value: "Tólf"},
		{Number: 11, Value: "Ellefu"},
		{Number: 10, Value: "Tíu"},
		{Number: 9, Value: "Níu"},
		{Number: 8, Value: "Átta"},
		{Number: 7, Value: "Sjö"},
		{Number: 6, Value: "Sex"},
		{Number: 5, Value: "Fimm"},
		{Number: 4, Value: "Fjögur"},
		{Number: 3, Value: "Þrjú"},
		{Number: 2, Value: "Tvö"},
		{Number: 1, Value: "Eitt"},
		{Number: 0, Value: "Núll"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 1000000000000000, Value: "Ein Trilljón"},
		{Number: 1000000000000, Value: "Ein Billjón"},
		{Number: 100000000000, Value: "Eitt Hundrað Billjarðar"},
		{Number: 10000000000, Value: "Tíu Billjarðar"},
		{Number: 1000000000, Value: "Einn Milljarður"},
		{Number: 100000000, Value: "Eitt Hundrað Milljónir"},
		{Number: 90000000, Value: "Níutíu Milljónir"},
		{Number: 80000000, Value: "Áttatíu Milljónir"},
		{Number: 70000000, Value: "Sjötíu Milljónir"},
		{Number: 60000000, Value: "Sextíu Milljónir"},
		{Number: 50000000, Value: "Fimmtíu Milljónir"},
		{Number: 40000000, Value: "Fjörutíu Milljónir"},
		{Number: 30000000, Value: "Þrjátíu Milljónir"},
		{Number: 20000000, Value: "Tuttugu Milljónir"},
		{Number: 10000000, Value: "Tíu Milljónir"},
		{Number: 9000000, Value: "Níu Milljónir"},
		{Number: 8000000, Value: "Átta Milljónir"},
		{Number: 7000000, Value: "Sjö Milljónir"},
		{Number: 6000000, Value: "Sex Milljónir"},
		{Number: 5000000, Value: "Fimm Milljónir"},
		{Number: 4000000, Value: "Fjórar Milljónir"},
		{Number: 3000000, Value: "Þrjár Milljónir"},
		{Number: 2000000, Value: "Tvær Milljónir"},
		{Number: 1000000, Value: "Ein Milljón"},
		{Number: 900000, Value: "Níu Hundrað Þúsund"},
		{Number: 800000, Value: "Átta Hundrað Þúsund"},
		{Number: 700000, Value: "Sjö Hundrað Þúsund"},
		{Number: 600000, Value: "Sex Hundrað Þúsund"},
		{Number: 500000, Value: "Fimm Hundrað Þúsund"},
		{Number: 400000, Value: "Fjögur Hundrað Þúsund"},
		{Number: 300000, Value: "Þrjú Hundrað Þúsund"},
		{Number: 200000, Value: "Tvö Hundrað Þúsund"},
		{Number: 100000, Value: "Eitt Hundrað Þúsund"},
		{Number: 90000, Value: "Níutíu Þúsund"},
		{Number: 80000, Value: "Áttatíu Þúsund"},
		{Number: 70000, Value: "Sjötíu Þúsund"},
		{Number: 60000, Value: "Sextíu Þúsund"},
		{Number: 50000, Value: "Fimmtíu Þúsund"},
		{Number: 40000, Value: "Fjörutíu Þúsund"},
		{Number: 30000, Value: "Þrjátíu Þúsund"},
		{Number: 20000, Value: "Tuttugu Þúsund"},
		{Number: 10000, Value: "Tíu Þúsund"},
		{Number: 9000, Value: "Níu Þúsund"},
		{Number: 8000, Value: "Átta Þúsund"},
		{Number: 7000, Value: "Sjö Þúsund"},
		{Number: 6000, Value: "Sex Þúsund"},
		{Number: 5000, Value: "Fimm Þúsund"},
		{Number: 4000, Value: "Fjögur Þúsund"},
		{Number: 3000, Value: "Þrjú Þúsund"},
		{Number: 2000, Value: "Tvö Þúsund"},
		{Number: 1000, Value: "Eitt Þúsund"},
		{Number: 900, Value: "Níu Hundrað"},
		{Number: 800, Value: "Átta Hundrað"},
		{Number: 700, Value: "Sjö Hundrað"},
		{Number: 600, Value: "Sex Hundrað"},
		{Number: 500, Value: "Fimm Hundrað"},
		{Number: 400, Value: "Fjögur Hundrað"},
		{Number: 300, Value: "Þrjú Hundrað"},
		{Number: 200, Value: "Tvö Hundrað"},
		{Number: 100, Value: "Eitt Hundrað"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "fyrsti", Suffix: ".", Masculine: "fyrsti", Feminine: "fyrsta", Neuter: "fyrsta"},
		{Number: 2, Word: "annar", Suffix: ".", Masculine: "annar", Feminine: "önnur", Neuter: "annað"},
		{Number: 3, Word: "þriðji", Suffix: ".", Masculine: "þriðji", Feminine: "þriðja", Neuter: "þriðja"},
		{Number: 4, Word: "fjórði", Suffix: ".", Masculine: "fjórði", Feminine: "fjórða", Neuter: "fjórða"},
		{Number: 5, Word: "fimmti", Suffix: ".", Masculine: "fimmti", Feminine: "fimmta", Neuter: "fimmta"},
		{Number: 6, Word: "sjötti", Suffix: ".", Masculine: "sjötti", Feminine: "sjötta", Neuter: "sjötta"},
		{Number: 7, Word: "sjöundi", Suffix: ".", Masculine: "sjöundi", Feminine: "sjöunda", Neuter: "sjöunda"},
		{Number: 8, Word: "áttundi", Suffix: ".", Masculine: "áttundi", Feminine: "áttunda", Neuter: "áttunda"},
		{Number: 9, Word: "níundi", Suffix: ".", Masculine: "níundi", Feminine: "níunda", Neuter: "níunda"},
		{Number: 10, Word: "tíundi", Suffix: ".", Masculine: "tíundi", Feminine: "tíunda", Neuter: "tíunda"},
		{Number: 11, Word: "ellefti", Suffix: ".", Masculine: "ellefti", Feminine: "ellefta", Neuter: "ellefta"},
		{Number: 12, Word: "tólfti", Suffix: ".", Masculine: "tólfti", Feminine: "tólfta", Neuter: "tólfta"},
		{Number: 13, Word: "þrettándi", Suffix: ".", Masculine: "þrettándi", Feminine: "þrettánda", Neuter: "þrettánda"},
		{Number: 14, Word: "fjórtándi", Suffix: ".", Masculine: "fjórtándi", Feminine: "fjórtánda", Neuter: "fjórtánda"},
		{Number: 15, Word: "fimmtándi", Suffix: ".", Masculine: "fimmtándi", Feminine: "fimmtánda", Neuter: "fimmtánda"},
		{Number: 16, Word: "sextándi", Suffix: ".", Masculine: "sextándi", Feminine: "sextánda", Neuter: "sextánda"},
		{Number: 17, Word: "sautjándi", Suffix: ".", Masculine: "sautjándi", Feminine: "sautjánda", Neuter: "sautjánda"},
		{Number: 18, Word: "átjándi", Suffix: ".", Masculine: "átjándi", Feminine: "átjánda", Neuter: "átjánda"},
		{Number: 19, Word: "nítjándi", Suffix: ".", Masculine: "nítjándi", Feminine: "nítjánda", Neuter: "nítjánda"},
		{Number: 20, Word: "tuttugasti", Suffix: ".", Masculine: "tuttugasti", Feminine: "tuttugasta", Neuter: "tuttugasta"},
		{Number: 21, Word: "tuttugasti og fyrsti", Suffix: ".", Masculine: "tuttugasti og fyrsti", Feminine: "tuttugasta og fyrsta", Neuter: "tuttugasta og fyrsta"},
		{Number: 30, Word: "þrítugasti", Suffix: ".", Masculine: "þrítugasti", Feminine: "þrítugasta", Neuter: "þrítugasta"},
		{Number: 40, Word: "fertugasti", Suffix: ".", Masculine: "fertugasti", Feminine: "fertugasta", Neuter: "fertugasta"},
		{Number: 50, Word: "fimmtugasti", Suffix: ".", Masculine: "fimmtugasti", Feminine: "fimmtugasta", Neuter: "fimmtugasta"},
		{Number: 60, Word: "sextugasti", Suffix: ".", Masculine: "sextugasti", Feminine: "sextugasta", Neuter: "sextugasta"},
		{Number: 70, Word: "sjötugasti", Suffix: ".", Masculine: "sjötugasti", Feminine: "sjötugasta", Neuter: "sjötugasta"},
		{Number: 80, Word: "áttatugasti", Suffix: ".", Masculine: "áttatugasti", Feminine: "áttatugasta", Neuter: "áttatugasta"},
		{Number: 90, Word: "nítugasti", Suffix: ".", Masculine: "nítugasti", Feminine: "nítugasta", Neuter: "nítugasta"},
		{Number: 100, Word: "hundraðasti", Suffix: ".", Masculine: "hundraðasti", Feminine: "hundraðasta", Neuter: "hundraðasta"},
		{Number: 1000, Word: "þúsundasti", Suffix: ".", Masculine: "þúsundasti", Feminine: "þúsundasta", Neuter: "þúsundasta"},
	},
	LocaleFormatter: &IcelandicFormatter{},
}

// IcelandicFormatter handles Icelandic (is-IS) formatting
type IcelandicFormatter struct{}

func (f *IcelandicFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	// First check exact mappings
	exactResult := ConvertToWordsWithExactMappingInt64(number, targetLocale)
	if exactResult != ConvertToWordsGenericInt64(number, targetLocale) {
		return exactResult // Found exact mapping
	}

	// Handle Icelandic gender agreement for millions
	decimalNumber := decimal.NewFromInt(number)
	if decimalNumber.Equal(decimal.Zero) {
		return GetWordForNumber(decimal.Zero, targetLocale)
	}

	result := ""
	million := decimal.NewFromInt(1000000)

	// Handle millions with gender agreement
	if decimalNumber.GreaterThanOrEqual(million) {
		millionPart := decimalNumber.Div(million).Floor()
		if millionPart.GreaterThan(decimal.Zero) {
			if millionPart.Equal(decimal.NewFromInt(1)) {
				result += "Ein Milljón" // Feminine form for "one million"
			} else {
				result += ConvertToWordsGenericInt64(millionPart.IntPart(), targetLocale) + " Milljónir"
			}
		}
		decimalNumber = decimalNumber.Mod(million)
		if decimalNumber.GreaterThan(decimal.Zero) {
			result += " " + ConvertToWordsGenericInt64(decimalNumber.IntPart(), targetLocale)
		}
		return result
	}

	// For numbers less than 1 million, use the standard generic conversion
	return ConvertToWordsGenericInt64(number, targetLocale)
}

func (f *IcelandicFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *IcelandicFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *IcelandicFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *IcelandicFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *IcelandicFormatter) ChopDecimal(number decimal.Decimal, places int) decimal.Decimal {
	return number.Truncate(int32(places))
}

func (f *IcelandicFormatter) FormatDecimalNumber(amount float64) string {
	return FormatEuropeanDecimal(amount)
}
func (f *IcelandicFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with European conventions (period separators, comma decimal, prefix symbol)
	return FormatEuropeanCurrency(amount, currencySymbol)
}
