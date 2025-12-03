package locale

import (
	"github.com/shopspring/decimal"
)

// BEBYLocale is a NumI18NLocale configured for Belarusian (Belarus) - be-BY
var BEBYLocale = NumI18NLocale{
	LocaleFormatter: &BelarusianFormatter{},
	Currency: Currency{
		Name:     "Беларускі рубель",
		Plural:   "Беларускія рублі",
		Singular: "Беларускі рубель",
		Symbol:   "BYN",
		FractionUnit: FractionUnit{
			Name:     "Капейка",
			Plural:   "Капейкі",
			Singular: "Капейка",
			Symbol:   "к.",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Belarus",
		Currency:       "BYN",
		ISO3166Alpha2:  "BY",
		ISO3166Alpha3:  "BLR",
		ISO3166Numeric: "112",
		Locale:         "be-BY",
		Timezone:       []string{"Europe/Minsk"},
		Language:       "be",
		Emoji:          "🇧🇾",
	},
	Texts: Texts{
		And:   "і",
		Minus: "мінус",
		Only:  "толькі",
		Point: "кропка",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Квадрыльён"},
		{Number: 1000000000000, Value: "Трыльён"},
		{Number: 1000000000, Value: "Мільярд"},
		{Number: 1000000, Value: "Мільён"},
		{Number: 1000, Value: "Тысяча"},
		{Number: 100, Value: "Сотня"},
		{Number: 90, Value: "Девяноста"},
		{Number: 80, Value: "Восемдесят"},
		{Number: 70, Value: "Семдесят"},
		{Number: 60, Value: "Шэсцьдзясят"},
		{Number: 50, Value: "Пяцьдзесят"},
		{Number: 40, Value: "Сорак"},
		{Number: 30, Value: "Трыццаць"},
		{Number: 20, Value: "Дваццаць"},
		{Number: 19, Value: "Дзесяць дзевяць"},
		{Number: 18, Value: "Восемнаццаць"},
		{Number: 17, Value: "Семнаццаць"},
		{Number: 16, Value: "Шаснаццаць"},
		{Number: 15, Value: "Пятнаццаць"},
		{Number: 14, Value: "Чатырнаццаць"},
		{Number: 13, Value: "Трынаццаць"},
		{Number: 12, Value: "Дванаццаць"},
		{Number: 11, Value: "Адзінаццаць"},
		{Number: 10, Value: "Дзесяць"},
		{Number: 9, Value: "Девяць"},
		{Number: 8, Value: "Восем"},
		{Number: 7, Value: "Сем"},
		{Number: 6, Value: "Шэсць"},
		{Number: 5, Value: "Пяць"},
		{Number: 4, Value: "Чатыры"},
		{Number: 3, Value: "Тры"},
		{Number: 2, Value: "Два"},
		{Number: 1, Value: "Адзін"},
		{Number: 0, Value: "Нуль"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Адзін Сотня"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "першы", Suffix: "", Masculine: "першы", Feminine: "першая", Neuter: "першае"},
		{Number: 2, Word: "другі", Suffix: "", Masculine: "другі", Feminine: "другая", Neuter: "другое"},
		{Number: 3, Word: "трэці", Suffix: "", Masculine: "трэці", Feminine: "трэцяя", Neuter: "трэцяе"},
		{Number: 4, Word: "чацвёрты", Suffix: "", Masculine: "чацвёрты", Feminine: "чацвёртая", Neuter: "чацвёртае"},
		{Number: 5, Word: "пяты", Suffix: "", Masculine: "пяты", Feminine: "пятая", Neuter: "пятае"},
		{Number: 6, Word: "шосты", Suffix: "", Masculine: "шосты", Feminine: "шостая", Neuter: "шостае"},
		{Number: 7, Word: "сёмы", Suffix: "", Masculine: "сёмы", Feminine: "сёмая", Neuter: "сёмае"},
		{Number: 8, Word: "восьмы", Suffix: "", Masculine: "восьмы", Feminine: "восьмая", Neuter: "восьмое"},
		{Number: 9, Word: "дзевяты", Suffix: "", Masculine: "дзевяты", Feminine: "дзевятая", Neuter: "дзевятае"},
		{Number: 10, Word: "дзесяты", Suffix: "", Masculine: "дзесяты", Feminine: "дзесятая", Neuter: "дзесятае"},
		{Number: 11, Word: "адзінаццаты", Suffix: "", Masculine: "адзінаццаты", Feminine: "адзінаццатая", Neuter: "адзінаццатае"},
		{Number: 12, Word: "дванаццаты", Suffix: "", Masculine: "дванаццаты", Feminine: "дванаццатая", Neuter: "дванаццатае"},
		{Number: 13, Word: "трынаццаты", Suffix: "", Masculine: "трынаццаты", Feminine: "трынаццатая", Neuter: "трынаццатае"},
		{Number: 14, Word: "чатырнаццаты", Suffix: "", Masculine: "чатырнаццаты", Feminine: "чатырнаццатая", Neuter: "чатырнаццатае"},
		{Number: 15, Word: "пятнаццаты", Suffix: "", Masculine: "пятнаццаты", Feminine: "пятнаццатая", Neuter: "пятнаццатае"},
		{Number: 16, Word: "шаснаццаты", Suffix: "", Masculine: "шаснаццаты", Feminine: "шаснаццатая", Neuter: "шаснаццатае"},
		{Number: 17, Word: "семнаццаты", Suffix: "", Masculine: "семнаццаты", Feminine: "семнаццатая", Neuter: "семнаццатае"},
		{Number: 18, Word: "восемнаццаты", Suffix: "", Masculine: "восемнаццаты", Feminine: "восемнаццатая", Neuter: "восемнаццатае"},
		{Number: 19, Word: "дзесяць дзевяты", Suffix: "", Masculine: "дзесяць дзевяты", Feminine: "дзесяць дзевятая", Neuter: "дзесяць дзевятае"},
		{Number: 20, Word: "дваццаты", Suffix: "", Masculine: "дваццаты", Feminine: "дваццатая", Neuter: "дваццатае"},
		{Number: 21, Word: "дваццаць першы", Suffix: "", Masculine: "дваццаць першы", Feminine: "дваццаць першая", Neuter: "дваццаць першае"},
		{Number: 30, Word: "трыццаты", Suffix: "", Masculine: "трыццаты", Feminine: "трыццатая", Neuter: "трыццатае"},
		{Number: 40, Word: "сорак", Suffix: "", Masculine: "сорак", Feminine: "сорак", Neuter: "сорак"},
		{Number: 50, Word: "пяцьдзесяты", Suffix: "", Masculine: "пяцьдзесяты", Feminine: "пяцьдзесятая", Neuter: "пяцьдзесятае"},
		{Number: 60, Word: "шэсцьдзесяты", Suffix: "", Masculine: "шэсцьдзесяты", Feminine: "шэсцьдзесятая", Neuter: "шэсцьдзесятае"},
		{Number: 70, Word: "семдзесяты", Suffix: "", Masculine: "семдзесяты", Feminine: "семдзесятая", Neuter: "семдзесятае"},
		{Number: 80, Word: "восьмдзесяты", Suffix: "", Masculine: "восьмдзесяты", Feminine: "восьмдзесятая", Neuter: "восьмдзесятае"},
		{Number: 90, Word: "дзевяносты", Suffix: "", Masculine: "дзевяносты", Feminine: "дзевяностая", Neuter: "дзевяностае"},
		{Number: 100, Word: "соты", Suffix: "", Masculine: "соты", Feminine: "сотая", Neuter: "сотае"},
		{Number: 1000, Word: "тысячны", Suffix: "", Masculine: "тысячны", Feminine: "тысячная", Neuter: "тысячнае"},
	},
}

// BelarusianFormatter handles Belarusian formatting
type BelarusianFormatter struct{}

func (f *BelarusianFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *BelarusianFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *BelarusianFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *BelarusianFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *BelarusianFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *BelarusianFormatter) ChopDecimal(d decimal.Decimal, precision int) decimal.Decimal {
	return d.Truncate(int32(precision))
}

func (f *BelarusianFormatter) FormatDecimalNumber(amount float64) string {
	return FormatPolishDecimal(amount)
}
func (f *BelarusianFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Polish conventions (comma separators, period decimal, prefix symbol)
	return FormatPolishCurrency(amount, currencySymbol)
}
