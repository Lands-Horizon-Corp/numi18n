package locale

import (
	"github.com/shopspring/decimal"
)

// KZLocale is a NumI18NLocale configured for Kazakhstan (kk-KZ)
var KZLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Теңге",
		Plural:   "Теңге",
		Singular: "Теңге",
		Symbol:   "₸",
		FractionUnit: FractionUnit{
			Name:     "Тиын",
			Plural:   "Тиын",
			Singular: "Тиын",
			Symbol:   "т",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Kazakhstan",
		Currency:       "KZT",
		ISO3166Alpha2:  "KZ",
		ISO3166Alpha3:  "KAZ",
		ISO3166Numeric: "398",
		Locale:         "kk-KZ",
		Timezone:       []string{"Asia/Almaty"},
		Language:       "kk",
		Emoji:          "🇰🇿",
	},
	Texts: Texts{
		And:   "Және",
		Minus: "Минус",
		Only:  "Тек",
		Point: "Нүкте",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Бір квадрилион"},
		{Number: 1000000000000, Value: "Бір трилион"},
		{Number: 1000000000, Value: "Бір миллиард"},
		{Number: 1000000, Value: "Бір миллион"},
		{Number: 1000, Value: "Бір Мың"},
		{Number: 100, Value: "Бір Жүз"},
		{Number: 90, Value: "Тоқсан"},
		{Number: 80, Value: "Сексен"},
		{Number: 70, Value: "Жетпіс"},
		{Number: 60, Value: "Алпыс"},
		{Number: 50, Value: "Елу"},
		{Number: 40, Value: "Қырық"},
		{Number: 30, Value: "Отыз"},
		{Number: 20, Value: "Жиырма"},
		{Number: 19, Value: "Он Тоғыз"},
		{Number: 18, Value: "Он Сегіз"},
		{Number: 17, Value: "Он Жеті"},
		{Number: 16, Value: "Он Алты"},
		{Number: 15, Value: "Он Бес"},
		{Number: 14, Value: "Он Төрт"},
		{Number: 13, Value: "Он Үш"},
		{Number: 12, Value: "Он Екі"},
		{Number: 11, Value: "Он Бір"},
		{Number: 10, Value: "Он"},
		{Number: 9, Value: "Тоғыз"},
		{Number: 8, Value: "Сегіз"},
		{Number: 7, Value: "Жеті"},
		{Number: 6, Value: "Алты"},
		{Number: 5, Value: "Бес"},
		{Number: 4, Value: "Төрт"},
		{Number: 3, Value: "Үш"},
		{Number: 2, Value: "Екі"},
		{Number: 1, Value: "Бір"},
		{Number: 0, Value: "Нөл"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Бір Жүз"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "бірінші", Suffix: "-ші", Masculine: "бірінші", Feminine: "бірінші", Neuter: "бірінші"},
		{Number: 2, Word: "екінші", Suffix: "-ші", Masculine: "екінші", Feminine: "екінші", Neuter: "екінші"},
		{Number: 3, Word: "үшінші", Suffix: "-ші", Masculine: "үшінші", Feminine: "үшінші", Neuter: "үшінші"},
		{Number: 4, Word: "төртінші", Suffix: "-ші", Masculine: "төртінші", Feminine: "төртінші", Neuter: "төртінші"},
		{Number: 5, Word: "бесінші", Suffix: "-ші", Masculine: "бесінші", Feminine: "бесінші", Neuter: "бесінші"},
		{Number: 6, Word: "алтыншы", Suffix: "-шы", Masculine: "алтыншы", Feminine: "алтыншы", Neuter: "алтыншы"},
		{Number: 7, Word: "жетінші", Suffix: "-ші", Masculine: "жетінші", Feminine: "жетінші", Neuter: "жетінші"},
		{Number: 8, Word: "сегізінші", Suffix: "-ші", Masculine: "сегізінші", Feminine: "сегізінші", Neuter: "сегізінші"},
		{Number: 9, Word: "тоғызыншы", Suffix: "-шы", Masculine: "тоғызыншы", Feminine: "тоғызыншы", Neuter: "тоғызыншы"},
		{Number: 10, Word: "оныншы", Suffix: "-шы", Masculine: "оныншы", Feminine: "оныншы", Neuter: "оныншы"},
		{Number: 11, Word: "он бірінші", Suffix: "-ші", Masculine: "он бірінші", Feminine: "он бірінші", Neuter: "он бірінші"},
		{Number: 12, Word: "он екінші", Suffix: "-ші", Masculine: "он екінші", Feminine: "он екінші", Neuter: "он екінші"},
		{Number: 13, Word: "он үшінші", Suffix: "-ші", Masculine: "он үшінші", Feminine: "он үшінші", Neuter: "он үшінші"},
		{Number: 14, Word: "он төртінші", Suffix: "-ші", Masculine: "он төртінші", Feminine: "он төртінші", Neuter: "он төртінші"},
		{Number: 15, Word: "он бесінші", Suffix: "-ші", Masculine: "он бесінші", Feminine: "он бесінші", Neuter: "он бесінші"},
		{Number: 16, Word: "он алтыншы", Suffix: "-шы", Masculine: "он алтыншы", Feminine: "он алтыншы", Neuter: "он алтыншы"},
		{Number: 17, Word: "он жетінші", Suffix: "-ші", Masculine: "он жетінші", Feminine: "он жетінші", Neuter: "он жетінші"},
		{Number: 18, Word: "он сегізінші", Suffix: "-ші", Masculine: "он сегізінші", Feminine: "он сегізінші", Neuter: "он сегізінші"},
		{Number: 19, Word: "он тоғызыншы", Suffix: "-шы", Masculine: "он тоғызыншы", Feminine: "он тоғызыншы", Neuter: "он тоғызыншы"},
		{Number: 20, Word: "жиырмасыншы", Suffix: "-шы", Masculine: "жиырмасыншы", Feminine: "жиырмасыншы", Neuter: "жиырмасыншы"},
		{Number: 21, Word: "жиырма бірінші", Suffix: "-ші", Masculine: "жиырма бірінші", Feminine: "жиырма бірінші", Neuter: "жиырма бірінші"},
		{Number: 30, Word: "отызыншы", Suffix: "-шы", Masculine: "отызыншы", Feminine: "отызыншы", Neuter: "отызыншы"},
		{Number: 40, Word: "қырықыншы", Suffix: "-шы", Masculine: "қырықыншы", Feminine: "қырықыншы", Neuter: "қырықыншы"},
		{Number: 50, Word: "елуінші", Suffix: "-ші", Masculine: "елуінші", Feminine: "елуінші", Neuter: "елуінші"},
		{Number: 60, Word: "алпысыншы", Suffix: "-шы", Masculine: "алпысыншы", Feminine: "алпысыншы", Neuter: "алпысыншы"},
		{Number: 70, Word: "жетпісінші", Suffix: "-ші", Masculine: "жетпісінші", Feminine: "жетпісінші", Neuter: "жетпісінші"},
		{Number: 80, Word: "сексенінші", Suffix: "-ші", Masculine: "сексенінші", Feminine: "сексенінші", Neuter: "сексенінші"},
		{Number: 90, Word: "тоқсаныншы", Suffix: "-шы", Masculine: "тоқсаныншы", Feminine: "тоқсаныншы", Neuter: "тоқсаныншы"},
		{Number: 100, Word: "жүзінші", Suffix: "-ші", Masculine: "жүзінші", Feminine: "жүзінші", Neuter: "жүзінші"},
		{Number: 1000, Word: "мыңыншы", Suffix: "-шы", Masculine: "мыңыншы", Feminine: "мыңыншы", Neuter: "мыңыншы"},
	},
	LocaleFormatter: &KazakhFormatter{},
}

// KazakhFormatter handles Kazakh formatting
type KazakhFormatter struct{}

func (f *KazakhFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *KazakhFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *KazakhFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *KazakhFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *KazakhFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *KazakhFormatter) ChopDecimal(value decimal.Decimal, places int) decimal.Decimal {
	multiplier := decimal.NewFromInt(1)
	for range places {
		multiplier = multiplier.Mul(decimal.NewFromInt(10))
	}
	return value.Mul(multiplier).Truncate(0).Div(multiplier)
}

func (f *KazakhFormatter) FormatDecimalNumber(amount float64) string {
	return FormatPolishDecimal(amount)
}
func (f *KazakhFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	return FormatPolishCurrency(amount, currencySymbol)
}
