package locale

import (
	"github.com/shopspring/decimal"
)

// UKUALocale represents the Ukrainian (Ukraine) locale
var UKUALocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Ukrainian Hryvnia",
		Plural:   "Гривні",
		Singular: "Гривня",
		Symbol:   "₴",
		FractionUnit: FractionUnit{
			Name:     "Kopiyka",
			Plural:   "Копійки",
			Singular: "Копійка",
			Symbol:   "к",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Ukraine",
		Currency:       "UAH",
		ISO3166Alpha2:  "UA",
		ISO3166Alpha3:  "UKR",
		ISO3166Numeric: "804",
		Locale:         "uk-UA",
		Timezone:       []string{"Europe/Kiev"},
		Language:       "uk",
		Emoji:          "🇺🇦",
		PhoneCode:      "+380",
		Domain:         ".ua",
	},
	Texts: Texts{
		And:   "і",
		Minus: "мінус",
		Only:  "лише",
		Point: "кома",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "один квадрільйон"},
		{Number: 1000000000000, Value: "один трільйон"},
		{Number: 1000000000, Value: "один мільярд"},
		{Number: 1000000, Value: "один мільйон"},
		{Number: 100000, Value: "сто тисяч"},
		{Number: 90000, Value: "дев'яносто тисяч"},
		{Number: 80000, Value: "вісімдесят тисяч"},
		{Number: 70000, Value: "сімдесят тисяч"},
		{Number: 60000, Value: "шістдесят тисяч"},
		{Number: 50000, Value: "п'ятдесят тисяч"},
		{Number: 40000, Value: "сорок тисяч"},
		{Number: 30000, Value: "тридцять тисяч"},
		{Number: 20000, Value: "двадцять тисяч"},
		{Number: 19000, Value: "дев'ятнадцять тисяч"},
		{Number: 18000, Value: "вісімнадцять тисяч"},
		{Number: 17000, Value: "сімнадцять тисяч"},
		{Number: 16000, Value: "шістнадцять тисяч"},
		{Number: 15000, Value: "п'ятнадцять тисяч"},
		{Number: 14000, Value: "чотирнадцять тисяч"},
		{Number: 13000, Value: "тринадцять тисяч"},
		{Number: 12000, Value: "дванадцять тисяч"},
		{Number: 11000, Value: "одинадцять тисяч"},
		{Number: 10000, Value: "десять тисяч"},
		{Number: 9000, Value: "дев'ять тисяч"},
		{Number: 8000, Value: "вісім тисяч"},
		{Number: 7000, Value: "сім тисяч"},
		{Number: 6000, Value: "шість тисяч"},
		{Number: 5000, Value: "п'ять тисяч"},
		{Number: 4000, Value: "чотири тисячі"},
		{Number: 3000, Value: "три тисячі"},
		{Number: 2000, Value: "дві тисячі"},
		{Number: 1000, Value: "одна тисяча"},
		{Number: 900, Value: "дев'ятсот"},
		{Number: 800, Value: "вісімсот"},
		{Number: 700, Value: "сімсот"},
		{Number: 600, Value: "шістсот"},
		{Number: 500, Value: "п'ятсот"},
		{Number: 400, Value: "чотириста"},
		{Number: 300, Value: "триста"},
		{Number: 200, Value: "двісті"},
		{Number: 100, Value: "сто"},
		{Number: 99, Value: "дев'яносто дев'ять"},
		{Number: 98, Value: "дев'яносто вісім"},
		{Number: 97, Value: "дев'яносто сім"},
		{Number: 96, Value: "дев'яносто шість"},
		{Number: 95, Value: "дев'яносто п'ять"},
		{Number: 94, Value: "дев'яносто чотири"},
		{Number: 93, Value: "дев'яносто три"},
		{Number: 92, Value: "дев'яносто два"},
		{Number: 91, Value: "дев'яносто один"},
		{Number: 90, Value: "дев'яносто"},
		{Number: 89, Value: "вісімдесят дев'ять"},
		{Number: 88, Value: "вісімдесят вісім"},
		{Number: 87, Value: "вісімдесят сім"},
		{Number: 86, Value: "вісімдесят шість"},
		{Number: 85, Value: "вісімдесят п'ять"},
		{Number: 84, Value: "вісімдесят чотири"},
		{Number: 83, Value: "вісімдесят три"},
		{Number: 82, Value: "вісімдесят два"},
		{Number: 81, Value: "вісімдесят один"},
		{Number: 80, Value: "вісімдесят"},
		{Number: 79, Value: "сімдесят дев'ять"},
		{Number: 78, Value: "сімдесят вісім"},
		{Number: 77, Value: "сімдесят сім"},
		{Number: 76, Value: "сімдесят шість"},
		{Number: 75, Value: "сімдесят п'ять"},
		{Number: 74, Value: "сімдесят чотири"},
		{Number: 73, Value: "сімдесят три"},
		{Number: 72, Value: "сімдесят два"},
		{Number: 71, Value: "сімдесят один"},
		{Number: 70, Value: "сімдесят"},
		{Number: 69, Value: "шістдесят дев'ять"},
		{Number: 68, Value: "шістдесят вісім"},
		{Number: 67, Value: "шістдесят сім"},
		{Number: 66, Value: "шістдесят шість"},
		{Number: 65, Value: "шістдесят п'ять"},
		{Number: 64, Value: "шістдесят чотири"},
		{Number: 63, Value: "шістдесят три"},
		{Number: 62, Value: "шістдесят два"},
		{Number: 61, Value: "шістдесят один"},
		{Number: 60, Value: "шістдесят"},
		{Number: 59, Value: "п'ятдесят дев'ять"},
		{Number: 58, Value: "п'ятдесят вісім"},
		{Number: 57, Value: "п'ятдесят сім"},
		{Number: 56, Value: "п'ятдесят шість"},
		{Number: 55, Value: "п'ятдесят п'ять"},
		{Number: 54, Value: "п'ятдесят чотири"},
		{Number: 53, Value: "п'ятдесят три"},
		{Number: 52, Value: "п'ятдесят два"},
		{Number: 51, Value: "п'ятдесят один"},
		{Number: 50, Value: "п'ятдесят"},
		{Number: 49, Value: "сорок дев'ять"},
		{Number: 48, Value: "сорок вісім"},
		{Number: 47, Value: "сорок сім"},
		{Number: 46, Value: "сорок шість"},
		{Number: 45, Value: "сорок п'ять"},
		{Number: 44, Value: "сорок чотири"},
		{Number: 43, Value: "сорок три"},
		{Number: 42, Value: "сорок два"},
		{Number: 41, Value: "сорок один"},
		{Number: 40, Value: "сорок"},
		{Number: 39, Value: "тридцять дев'ять"},
		{Number: 38, Value: "тридцять вісім"},
		{Number: 37, Value: "тридцять сім"},
		{Number: 36, Value: "тридцять шість"},
		{Number: 35, Value: "тридцять п'ять"},
		{Number: 34, Value: "тридцять чотири"},
		{Number: 33, Value: "тридцять три"},
		{Number: 32, Value: "тридцять два"},
		{Number: 31, Value: "тридцять один"},
		{Number: 30, Value: "тридцять"},
		{Number: 29, Value: "двадцять дев'ять"},
		{Number: 28, Value: "двадцять вісім"},
		{Number: 27, Value: "двадцять сім"},
		{Number: 26, Value: "двадцять шість"},
		{Number: 25, Value: "двадцять п'ять"},
		{Number: 24, Value: "двадцять чотири"},
		{Number: 23, Value: "двадцять три"},
		{Number: 22, Value: "двадцять два"},
		{Number: 21, Value: "двадцять один"},
		{Number: 20, Value: "двадцять"},
		{Number: 19, Value: "дев'ятнадцять"},
		{Number: 18, Value: "вісімнадцять"},
		{Number: 17, Value: "сімнадцять"},
		{Number: 16, Value: "шістнадцять"},
		{Number: 15, Value: "п'ятнадцять"},
		{Number: 14, Value: "чотирнадцять"},
		{Number: 13, Value: "тринадцять"},
		{Number: 12, Value: "дванадцять"},
		{Number: 11, Value: "одинадцять"},
		{Number: 10, Value: "десять"},
		{Number: 9, Value: "дев'ять"},
		{Number: 8, Value: "вісім"},
		{Number: 7, Value: "сім"},
		{Number: 6, Value: "шість"},
		{Number: 5, Value: "п'ять"},
		{Number: 4, Value: "чотири"},
		{Number: 3, Value: "три"},
		{Number: 2, Value: "два"},
		{Number: 1, Value: "один"},
		{Number: 0, Value: "нуль"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Сто"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "перший", Suffix: "-й", Masculine: "перший", Feminine: "перша", Neuter: "перше"},
		{Number: 2, Word: "другий", Suffix: "-й", Masculine: "другий", Feminine: "друга", Neuter: "друге"},
		{Number: 3, Word: "третій", Suffix: "-й", Masculine: "третій", Feminine: "третя", Neuter: "третє"},
		{Number: 4, Word: "четвертий", Suffix: "-й", Masculine: "четвертий", Feminine: "четверта", Neuter: "четверте"},
		{Number: 5, Word: "п'ятий", Suffix: "-й", Masculine: "п'ятий", Feminine: "п'ята", Neuter: "п'яте"},
		{Number: 6, Word: "шостий", Suffix: "-й", Masculine: "шостий", Feminine: "шоста", Neuter: "шосте"},
		{Number: 7, Word: "сьомий", Suffix: "-й", Masculine: "сьомий", Feminine: "сьома", Neuter: "сьоме"},
		{Number: 8, Word: "восьмий", Suffix: "-й", Masculine: "восьмий", Feminine: "восьма", Neuter: "восьме"},
		{Number: 9, Word: "дев'ятий", Suffix: "-й", Masculine: "дев'ятий", Feminine: "дев'ята", Neuter: "дев'яте"},
		{Number: 10, Word: "десятий", Suffix: "-й", Masculine: "десятий", Feminine: "десята", Neuter: "десяте"},
		{Number: 11, Word: "одинадцятий", Suffix: "-й", Masculine: "одинадцятий", Feminine: "одинадцята", Neuter: "одинадцяте"},
		{Number: 12, Word: "дванадцятий", Suffix: "-й", Masculine: "дванадцятий", Feminine: "дванадцята", Neuter: "дванадцяте"},
		{Number: 13, Word: "тринадцятий", Suffix: "-й", Masculine: "тринадцятий", Feminine: "тринадцята", Neuter: "тринадцяте"},
		{Number: 14, Word: "чотирнадцятий", Suffix: "-й", Masculine: "чотирнадцятий", Feminine: "чотирнадцята", Neuter: "чотирнадцяте"},
		{Number: 15, Word: "п'ятнадцятий", Suffix: "-й", Masculine: "п'ятнадцятий", Feminine: "п'ятнадцята", Neuter: "п'ятнадцяте"},
		{Number: 16, Word: "шістнадцятий", Suffix: "-й", Masculine: "шістнадцятий", Feminine: "шістнадцята", Neuter: "шістнадцяте"},
		{Number: 17, Word: "сімнадцятий", Suffix: "-й", Masculine: "сімнадцятий", Feminine: "сімнадцята", Neuter: "сімнадцяте"},
		{Number: 18, Word: "вісімнадцятий", Suffix: "-й", Masculine: "вісімнадцятий", Feminine: "вісімнадцята", Neuter: "вісімнадцяте"},
		{Number: 19, Word: "дев'ятнадцятий", Suffix: "-й", Masculine: "дев'ятнадцятий", Feminine: "дев'ятнадцята", Neuter: "дев'ятнадцяте"},
		{Number: 20, Word: "двадцятий", Suffix: "-й", Masculine: "двадцятий", Feminine: "двадцята", Neuter: "двадцяте"},
		{Number: 21, Word: "двадцять перший", Suffix: "-й", Masculine: "двадцять перший", Feminine: "двадцять перша", Neuter: "двадцять перше"},
		{Number: 30, Word: "тридцятий", Suffix: "-й", Masculine: "тридцятий", Feminine: "тридцята", Neuter: "тридцяте"},
		{Number: 40, Word: "сороковий", Suffix: "-й", Masculine: "сороковий", Feminine: "сорокова", Neuter: "сорокове"},
		{Number: 50, Word: "п'ятдесятий", Suffix: "-й", Masculine: "п'ятдесятий", Feminine: "п'ятдесята", Neuter: "п'ятдесяте"},
		{Number: 60, Word: "шістдесятий", Suffix: "-й", Masculine: "шістдесятий", Feminine: "шістдесята", Neuter: "шістдесяте"},
		{Number: 70, Word: "сімдесятий", Suffix: "-й", Masculine: "сімдесятий", Feminine: "сімдесята", Neuter: "сімдесяте"},
		{Number: 80, Word: "вісімдесятий", Suffix: "-й", Masculine: "вісімдесятий", Feminine: "вісімдесята", Neuter: "вісімдесяте"},
		{Number: 90, Word: "дев'яностий", Suffix: "-й", Masculine: "дев'яностий", Feminine: "дев'яноста", Neuter: "дев'яносте"},
		{Number: 100, Word: "сотий", Suffix: "-й", Masculine: "сотий", Feminine: "сота", Neuter: "соте"},
		{Number: 1000, Word: "тисячний", Suffix: "-й", Masculine: "тисячний", Feminine: "тисячна", Neuter: "тисячне"},
		{Number: 1000000, Word: "мільйонний", Suffix: "-й", Masculine: "мільйонний", Feminine: "мільйонна", Neuter: "мільйонне"},
	},
	LocaleFormatter: &UkrainianFormatter{},
}

// UkrainianFormatter handles Ukrainian (Ukraine) formatting
type UkrainianFormatter struct{}

func (f *UkrainianFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *UkrainianFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *UkrainianFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *UkrainianFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *UkrainianFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *UkrainianFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	return amount.Truncate(int32(precision))
}

func (f *UkrainianFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *UkrainianFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
