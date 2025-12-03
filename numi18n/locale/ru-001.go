package locale

import (
	"github.com/shopspring/decimal"
)

// RU001Locale is a NumI18NLocale configured for Russian (World) - ru-001
var RU001Locale = NumI18NLocale{
	Currency: Currency{
		Name:     "рубль",
		Plural:   "рублей",
		Singular: "рубль",
		Symbol:   "₽",
		FractionUnit: FractionUnit{
			Name:     "копейка",
			Plural:   "копеек",
			Singular: "копейка",
			Symbol:   "к",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "World",
		Currency:       "USD",
		ISO3166Alpha2:  "001",
		ISO3166Alpha3:  "WLD",
		ISO3166Numeric: "001",
		Locale:         "ru-001",
		Timezone:       []string{"UTC"},
		Language:       "ru",
		Emoji:          "🌍",
	},
	Texts: Texts{
		And:   "и",
		Minus: "минус",
		Only:  "только",
		Point: "точка",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "квадриллион"},
		{Number: 1000000000000, Value: "триллион"},
		{Number: 1000000000, Value: "миллиард"},
		{Number: 1000000, Value: "миллион"},
		{Number: 1000, Value: "тысяча"},
		{Number: 100, Value: "сто"},
		{Number: 90, Value: "девяносто"},
		{Number: 80, Value: "восемьдесят"},
		{Number: 70, Value: "семьдесят"},
		{Number: 60, Value: "шестьдесят"},
		{Number: 50, Value: "пятьдесят"},
		{Number: 40, Value: "сорок"},
		{Number: 30, Value: "тридцать"},
		{Number: 20, Value: "двадцать"},
		{Number: 19, Value: "девятнадцать"},
		{Number: 18, Value: "восемнадцать"},
		{Number: 17, Value: "семнадцать"},
		{Number: 16, Value: "шестнадцать"},
		{Number: 15, Value: "пятнадцать"},
		{Number: 14, Value: "четырнадцать"},
		{Number: 13, Value: "тринадцать"},
		{Number: 12, Value: "двенадцать"},
		{Number: 11, Value: "одиннадцать"},
		{Number: 10, Value: "десять"},
		{Number: 9, Value: "девять"},
		{Number: 8, Value: "восемь"},
		{Number: 7, Value: "семь"},
		{Number: 6, Value: "шесть"},
		{Number: 5, Value: "пять"},
		{Number: 4, Value: "четыре"},
		{Number: 3, Value: "три"},
		{Number: 2, Value: "два"},
		{Number: 1, Value: "один"},
		{Number: 0, Value: "ноль"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "сто"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "первый", Suffix: "-й", Masculine: "первый", Feminine: "первая", Neuter: "первое"},
		{Number: 2, Word: "второй", Suffix: "-й", Masculine: "второй", Feminine: "вторая", Neuter: "второе"},
		{Number: 3, Word: "третий", Suffix: "-й", Masculine: "третий", Feminine: "третья", Neuter: "третье"},
		{Number: 4, Word: "четвёртый", Suffix: "-й", Masculine: "четвёртый", Feminine: "четвёртая", Neuter: "четвёртое"},
		{Number: 5, Word: "пятый", Suffix: "-й", Masculine: "пятый", Feminine: "пятая", Neuter: "пятое"},
		{Number: 6, Word: "шестой", Suffix: "-й", Masculine: "шестой", Feminine: "шестая", Neuter: "шестое"},
		{Number: 7, Word: "седьмой", Suffix: "-й", Masculine: "седьмой", Feminine: "седьмая", Neuter: "седьмое"},
		{Number: 8, Word: "восьмой", Suffix: "-й", Masculine: "восьмой", Feminine: "восьмая", Neuter: "восьмое"},
		{Number: 9, Word: "девятый", Suffix: "-й", Masculine: "девятый", Feminine: "девятая", Neuter: "девятое"},
		{Number: 10, Word: "десятый", Suffix: "-й", Masculine: "десятый", Feminine: "десятая", Neuter: "десятое"},
		{Number: 11, Word: "одиннадцатый", Suffix: "-й", Masculine: "одиннадцатый", Feminine: "одиннадцатая", Neuter: "одиннадцатое"},
		{Number: 12, Word: "двенадцатый", Suffix: "-й", Masculine: "двенадцатый", Feminine: "двенадцатая", Neuter: "двенадцатое"},
		{Number: 20, Word: "двадцатый", Suffix: "-й", Masculine: "двадцатый", Feminine: "двадцатая", Neuter: "двадцатое"},
		{Number: 21, Word: "двадцать первый", Suffix: "-й", Masculine: "двадцать первый", Feminine: "двадцать первая", Neuter: "двадцать первое"},
		{Number: 30, Word: "тридцатый", Suffix: "-й", Masculine: "тридцатый", Feminine: "тридцатая", Neuter: "тридцатое"},
		{Number: 50, Word: "пятидесятый", Suffix: "-й", Masculine: "пятидесятый", Feminine: "пятидесятая", Neuter: "пятидесятое"},
		{Number: 100, Word: "сотый", Suffix: "-й", Masculine: "сотый", Feminine: "сотая", Neuter: "сотое"},
		{Number: 1000, Word: "тысячный", Suffix: "-й", Masculine: "тысячный", Feminine: "тысячная", Neuter: "тысячное"},
	},
	LocaleFormatter: &RussianFormatter{},
}

// RussianFormatter handles Russian formatting
type RussianFormatter struct{}

func (f *RussianFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *RussianFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *RussianFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *RussianFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *RussianFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *RussianFormatter) ChopDecimal(value decimal.Decimal, precision int) decimal.Decimal {
	return value.Truncate(int32(precision))
}

func (f *RussianFormatter) FormatDecimalNumber(amount float64) string {
	return FormatPolishDecimal(amount)
}
func (f *RussianFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Polish conventions (comma separators, period decimal, prefix symbol)
	return FormatPolishCurrency(amount, currencySymbol)
}
