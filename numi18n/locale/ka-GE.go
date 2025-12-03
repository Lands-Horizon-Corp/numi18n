package locale

import (
	"github.com/shopspring/decimal"
)

// GELocale is a NumI18NLocale configured for Georgia (ka-GE)
var GELocale = NumI18NLocale{
	Currency: Currency{
		Name:     "ლარი",
		Plural:   "ლარი",
		Singular: "ლარი",
		Symbol:   "₾",
		FractionUnit: FractionUnit{
			Name:     "თეთრი",
			Plural:   "თეთრი",
			Singular: "თეთრი",
			Symbol:   "თ",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Georgia",
		Currency:       "GEL",
		ISO3166Alpha2:  "GE",
		ISO3166Alpha3:  "GEO",
		ISO3166Numeric: "268",
		Locale:         "ka-GE",
		Timezone:       []string{"Asia/Tbilisi"},
		Language:       "ka",
		Emoji:          "🇬🇪",
	},
	Texts: Texts{
		And:   "და",
		Minus: "მინუს",
		Only:  "მხოლოდ",
		Point: "წერტილი",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Ერთი Კვადრილიონი"},
		{Number: 1000000000000, Value: "Ერთი Ტრილიონი"},
		{Number: 1000000000, Value: "Ერთი Მილიარდი"},
		{Number: 1000000, Value: "Ერთი Მილიონი"},
		{Number: 1000, Value: "Ერთი Ათასი"},
		{Number: 100, Value: "Ერთი Ასი"},
		{Number: 90, Value: "ოთხმოცდაათი"},
		{Number: 80, Value: "ოთხმოცი"},
		{Number: 70, Value: "სამოცდაათი"},
		{Number: 60, Value: "სამოცი"},
		{Number: 50, Value: "ორმოცდაათი"},
		{Number: 40, Value: "ორმოცი"},
		{Number: 30, Value: "ოცდაათი"},
		{Number: 20, Value: "ოცი"},
		{Number: 19, Value: "ცხრამეტი"},
		{Number: 18, Value: "თვრამეტი"},
		{Number: 17, Value: "ჩვიდმეტი"},
		{Number: 16, Value: "თექვსმეტი"},
		{Number: 15, Value: "თხუთმეტი"},
		{Number: 14, Value: "თოთხმეტი"},
		{Number: 13, Value: "ცამეტი"},
		{Number: 12, Value: "თორმეტი"},
		{Number: 11, Value: "თერთმეტი"},
		{Number: 10, Value: "ათი"},
		{Number: 9, Value: "ცხრა"},
		{Number: 8, Value: "რვა"},
		{Number: 7, Value: "შვიდი"},
		{Number: 6, Value: "ექვსი"},
		{Number: 5, Value: "ხუთი"},
		{Number: 4, Value: "ოთხი"},
		{Number: 3, Value: "სამი"},
		{Number: 2, Value: "ორი"},
		{Number: 1, Value: "ერთი"},
		{Number: 0, Value: "ნული"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Ერთი Ასი"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "პირველი", Suffix: "-ე", Masculine: "პირველი", Feminine: "პირველი", Neuter: "პირველი"},
		{Number: 2, Word: "მეორე", Suffix: "-ე", Masculine: "მეორე", Feminine: "მეორე", Neuter: "მეორე"},
		{Number: 3, Word: "მესამე", Suffix: "-ე", Masculine: "მესამე", Feminine: "მესამე", Neuter: "მესამე"},
		{Number: 4, Word: "მეოთხე", Suffix: "-ე", Masculine: "მეოთხე", Feminine: "მეოთხე", Neuter: "მეოთხე"},
		{Number: 5, Word: "მეხუთე", Suffix: "-ე", Masculine: "მეხუთე", Feminine: "მეხუთე", Neuter: "მეხუთე"},
		{Number: 6, Word: "მეექვსე", Suffix: "-ე", Masculine: "მეექვსე", Feminine: "მეექვსე", Neuter: "მეექვსე"},
		{Number: 7, Word: "მეშვიდე", Suffix: "-ე", Masculine: "მეშვიდე", Feminine: "მეშვიდე", Neuter: "მეშვიდე"},
		{Number: 8, Word: "მერვე", Suffix: "-ე", Masculine: "მერვე", Feminine: "მერვე", Neuter: "მერვე"},
		{Number: 9, Word: "მეცხრე", Suffix: "-ე", Masculine: "მეცხრე", Feminine: "მეცხრე", Neuter: "მეცხრე"},
		{Number: 10, Word: "მეათე", Suffix: "-ე", Masculine: "მეათე", Feminine: "მეათე", Neuter: "მეათე"},
		{Number: 11, Word: "მეთერთმეტე", Suffix: "-ე", Masculine: "მეთერთმეტე", Feminine: "მეთერთმეტე", Neuter: "მეთერთმეტე"},
		{Number: 12, Word: "მეథორმეტე", Suffix: "-ე", Masculine: "მეთორმეტე", Feminine: "მეთორმეტე", Neuter: "მეთორმეტე"},
		{Number: 13, Word: "მეცამეტე", Suffix: "-ე", Masculine: "მეცამეტე", Feminine: "მეცამეტე", Neuter: "მეცამეტე"},
		{Number: 14, Word: "მეთოთხმეტე", Suffix: "-ე", Masculine: "მეთოთხმეტე", Feminine: "მეთოთხმეტე", Neuter: "მეთოთხმეტე"},
		{Number: 15, Word: "მეთხუთმეტე", Suffix: "-ე", Masculine: "მეთხუთმეტე", Feminine: "მეთხუთმეტე", Neuter: "მეთხუთმეტე"},
		{Number: 16, Word: "მეთექვსმეტე", Suffix: "-ე", Masculine: "მეთექვსმეტე", Feminine: "მეთექვსმეტე", Neuter: "მეთექვსმეტე"},
		{Number: 17, Word: "მეჩვიდმეტე", Suffix: "-ე", Masculine: "მეჩვიდმეტე", Feminine: "მეჩვიდმეტე", Neuter: "მეჩვიდმეტე"},
		{Number: 18, Word: "მეთვრამეტე", Suffix: "-ე", Masculine: "მეთვრამეტე", Feminine: "მეთვრამეტე", Neuter: "მეთვრამეტე"},
		{Number: 19, Word: "მეცხრამეტე", Suffix: "-ე", Masculine: "მეცხრამეტე", Feminine: "მეცხრამეტე", Neuter: "მეცხრამეტე"},
		{Number: 20, Word: "მეოცე", Suffix: "-ე", Masculine: "მეოცე", Feminine: "მეოცე", Neuter: "მეოცე"},
		{Number: 21, Word: "ოცდაპირველი", Suffix: "-ე", Masculine: "ოცდაპირველი", Feminine: "ოცდაპირველი", Neuter: "ოცდაპირველი"},
		{Number: 30, Word: "მეოცდაათე", Suffix: "-ე", Masculine: "მეოცდაათე", Feminine: "მეოცდაათე", Neuter: "მეოცდაათე"},
		{Number: 40, Word: "მეორმოცე", Suffix: "-ე", Masculine: "მეორმოცე", Feminine: "მეორმოცე", Neuter: "მეორმოცე"},
		{Number: 50, Word: "მეორმოცდაათე", Suffix: "-ე", Masculine: "მეორმოცდაათე", Feminine: "მეორმოცდაათე", Neuter: "მეორმოცდაათე"},
		{Number: 60, Word: "მესამოცე", Suffix: "-ე", Masculine: "მესამოცე", Feminine: "მესამოცე", Neuter: "მესამოცე"},
		{Number: 70, Word: "მესამოცდაათე", Suffix: "-ე", Masculine: "მესამოცდაათე", Feminine: "მესამოცდაათე", Neuter: "მესამოცდაათე"},
		{Number: 80, Word: "მეოთხმოცე", Suffix: "-ე", Masculine: "მეოთხმოცე", Feminine: "მეოთხმოცე", Neuter: "მეოთხმოცე"},
		{Number: 90, Word: "მეოთხმოცდაათე", Suffix: "-ე", Masculine: "მეოთხმოცდაათე", Feminine: "მეოთხმოცდაათე", Neuter: "მეოთხმოცდაათე"},
		{Number: 100, Word: "მეასე", Suffix: "-ე", Masculine: "მეასე", Feminine: "მეასე", Neuter: "მეასე"},
		{Number: 1000, Word: "მეათასე", Suffix: "-ე", Masculine: "მეათასე", Feminine: "მეათასე", Neuter: "მეათასე"},
	},
	LocaleFormatter: &GeorgianFormatter{},
}

// GeorgianFormatter handles Georgian formatting
type GeorgianFormatter struct{}

func (f *GeorgianFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *GeorgianFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *GeorgianFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *GeorgianFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *GeorgianFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *GeorgianFormatter) ChopDecimal(value decimal.Decimal, places int) decimal.Decimal {
	multiplier := decimal.NewFromInt(1)
	for range places {
		multiplier = multiplier.Mul(decimal.NewFromInt(10))
	}
	return value.Mul(multiplier).Truncate(0).Div(multiplier)
}

func (f *GeorgianFormatter) FormatDecimalNumber(amount float64) string {
	return FormatPolishDecimal(amount)
}
func (f *GeorgianFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Polish conventions (comma separators, period decimal, prefix symbol)
	return FormatPolishCurrency(amount, currencySymbol)
}
