package locale

import (
	"github.com/shopspring/decimal"
)

// RUBYLocale represents the Russian (Belarus) locale
var RUBYLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Белорусский рубль",
		Plural:   "Белорусских рублей",
		Singular: "Белорусский рубль",
		Symbol:   "BYN",
		FractionUnit: FractionUnit{
			Name:     "Копейка",
			Plural:   "Копеек",
			Singular: "Копейка",
			Symbol:   "к",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Belarus",
		Currency:       "BYN",
		ISO3166Alpha2:  "BY",
		ISO3166Alpha3:  "BLR",
		ISO3166Numeric: "112",
		Locale:         "ru-BY",
		Timezone:       []string{"Europe/Minsk"},
		Language:       "ru",
	},
	Texts: Texts{
		And:   "и",
		Minus: "минус",
		Only:  "только",
		Point: "запятая",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "квадриллион"},
		{Number: 1000000000000, Value: "триллион"},
		{Number: 1000000000, Value: "миллиард"},
		{Number: 1000000, Value: "миллион"},
		{Number: 100000, Value: "сто тысяч"},
		{Number: 90000, Value: "девяносто тысяч"},
		{Number: 80000, Value: "восемьдесят тысяч"},
		{Number: 70000, Value: "семьдесят тысяч"},
		{Number: 60000, Value: "шестьдесят тысяч"},
		{Number: 50000, Value: "пятьдесят тысяч"},
		{Number: 40000, Value: "сорок тысяч"},
		{Number: 30000, Value: "тридцать тысяч"},
		{Number: 20000, Value: "двадцать тысяч"},
		{Number: 19000, Value: "девятнадцать тысяч"},
		{Number: 18000, Value: "восемнадцать тысяч"},
		{Number: 17000, Value: "семнадцать тысяч"},
		{Number: 16000, Value: "шестнадцать тысяч"},
		{Number: 15000, Value: "пятнадцать тысяч"},
		{Number: 14000, Value: "четырнадцать тысяч"},
		{Number: 13000, Value: "тринадцать тысяч"},
		{Number: 12000, Value: "двенадцать тысяч"},
		{Number: 11000, Value: "одиннадцать тысяч"},
		{Number: 10000, Value: "десять тысяч"},
		{Number: 9000, Value: "девять тысяч"},
		{Number: 8000, Value: "восемь тысяч"},
		{Number: 7000, Value: "семь тысяч"},
		{Number: 6000, Value: "шесть тысяч"},
		{Number: 5000, Value: "пять тысяч"},
		{Number: 4000, Value: "четыре тысячи"},
		{Number: 3000, Value: "три тысячи"},
		{Number: 2000, Value: "две тысячи"},
		{Number: 1000, Value: "тысяча"},
		{Number: 900, Value: "девятьсот"},
		{Number: 800, Value: "восемьсот"},
		{Number: 700, Value: "семьсот"},
		{Number: 600, Value: "шестьсот"},
		{Number: 500, Value: "пятьсот"},
		{Number: 400, Value: "четыреста"},
		{Number: 300, Value: "триста"},
		{Number: 200, Value: "двести"},
		{Number: 100, Value: "сто"},
		{Number: 99, Value: "девяносто девять"},
		{Number: 98, Value: "девяносто восемь"},
		{Number: 97, Value: "девяносто семь"},
		{Number: 96, Value: "девяносто шесть"},
		{Number: 95, Value: "девяносто пять"},
		{Number: 94, Value: "девяносто четыре"},
		{Number: 93, Value: "девяносто три"},
		{Number: 92, Value: "девяносто два"},
		{Number: 91, Value: "девяносто один"},
		{Number: 90, Value: "девяносто"},
		{Number: 89, Value: "восемьдесят девять"},
		{Number: 88, Value: "восемьдесят восемь"},
		{Number: 87, Value: "восемьдесят семь"},
		{Number: 86, Value: "восемьдесят шесть"},
		{Number: 85, Value: "восемьдесят пять"},
		{Number: 84, Value: "восемьдесят четыре"},
		{Number: 83, Value: "восемьдесят три"},
		{Number: 82, Value: "восемьдесят два"},
		{Number: 81, Value: "восемьдесят один"},
		{Number: 80, Value: "восемьдесят"},
		{Number: 79, Value: "семьдесят девять"},
		{Number: 78, Value: "семьдесят восемь"},
		{Number: 77, Value: "семьдесят семь"},
		{Number: 76, Value: "семьдесят шесть"},
		{Number: 75, Value: "семьдесят пять"},
		{Number: 74, Value: "семьдесят четыре"},
		{Number: 73, Value: "семьдесят три"},
		{Number: 72, Value: "семьдесят два"},
		{Number: 71, Value: "семьдесят один"},
		{Number: 70, Value: "семьдесят"},
		{Number: 69, Value: "шестьдесят девять"},
		{Number: 68, Value: "шестьдесят восемь"},
		{Number: 67, Value: "шестьдесят семь"},
		{Number: 66, Value: "шестьдесят шесть"},
		{Number: 65, Value: "шестьдесят пять"},
		{Number: 64, Value: "шестьдесят четыре"},
		{Number: 63, Value: "шестьдесят три"},
		{Number: 62, Value: "шестьдесят два"},
		{Number: 61, Value: "шестьдесят один"},
		{Number: 60, Value: "шестьдесят"},
		{Number: 59, Value: "пятьдесят девять"},
		{Number: 58, Value: "пятьдесят восемь"},
		{Number: 57, Value: "пятьдесят семь"},
		{Number: 56, Value: "пятьдесят шесть"},
		{Number: 55, Value: "пятьдесят пять"},
		{Number: 54, Value: "пятьдесят четыре"},
		{Number: 53, Value: "пятьдесят три"},
		{Number: 52, Value: "пятьдесят два"},
		{Number: 51, Value: "пятьдесят один"},
		{Number: 50, Value: "пятьдесят"},
		{Number: 49, Value: "сорок девять"},
		{Number: 48, Value: "сорок восемь"},
		{Number: 47, Value: "сорок семь"},
		{Number: 46, Value: "сорок шесть"},
		{Number: 45, Value: "сорок пять"},
		{Number: 44, Value: "сорок четыре"},
		{Number: 43, Value: "сорок три"},
		{Number: 42, Value: "сорок два"},
		{Number: 41, Value: "сорок один"},
		{Number: 40, Value: "сорок"},
		{Number: 39, Value: "тридцать девять"},
		{Number: 38, Value: "тридцать восемь"},
		{Number: 37, Value: "тридцать семь"},
		{Number: 36, Value: "тридцать шесть"},
		{Number: 35, Value: "тридцать пять"},
		{Number: 34, Value: "тридцать четыре"},
		{Number: 33, Value: "тридцать три"},
		{Number: 32, Value: "тридцать два"},
		{Number: 31, Value: "тридцать один"},
		{Number: 30, Value: "тридцать"},
		{Number: 29, Value: "двадцать девять"},
		{Number: 28, Value: "двадцать восемь"},
		{Number: 27, Value: "двадцать семь"},
		{Number: 26, Value: "двадцать шесть"},
		{Number: 25, Value: "двадцать пять"},
		{Number: 24, Value: "двадцать четыре"},
		{Number: 23, Value: "двадцать три"},
		{Number: 22, Value: "двадцать два"},
		{Number: 21, Value: "двадцать один"},
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
		{Number: 100, Value: "Сто"},
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
	LocaleFormatter: &RussianBelarusFormatter{},
}

// RussianBelarusFormatter implements LocaleFormatter for Russian (Belarus)
type RussianBelarusFormatter struct{}

func (f *RussianBelarusFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *RussianBelarusFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *RussianBelarusFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *RussianBelarusFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return " " + fractionName
	}
	return " " + fractionPlural
}

func (f *RussianBelarusFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *RussianBelarusFormatter) ChopDecimal(value decimal.Decimal, precision int) decimal.Decimal {
	return value.Truncate(int32(precision))
}

func (f *RussianBelarusFormatter) FormatDecimalNumber(amount float64) string {
	return FormatPolishDecimal(amount)
}
func (f *RussianBelarusFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Polish conventions (comma separators, period decimal, prefix symbol)
	return FormatPolishCurrency(amount, currencySymbol)
}
