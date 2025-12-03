package locale

import (
	"github.com/shopspring/decimal"
)

// KGLocale is a NumI18NLocale configured for Kyrgyzstan (ky-KG)
var KGLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Сом",
		Plural:   "Сом",
		Singular: "Сом",
		Symbol:   "с",
		FractionUnit: FractionUnit{
			Name:     "Тыйын",
			Plural:   "Тыйын",
			Singular: "Тыйын",
			Symbol:   "т",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Kyrgyzstan",
		Currency:       "KGS",
		ISO3166Alpha2:  "KG",
		ISO3166Alpha3:  "KGZ",
		ISO3166Numeric: "417",
		Locale:         "ky-KG",
		Timezone:       []string{"Asia/Bishkek"},
		Language:       "ky",
		Emoji:          "🇰🇬",
		PhoneCode:      "+996",
		Domain:         ".kg",
	},
	Texts: Texts{
		And:   "Жана",
		Minus: "Минус",
		Only:  "Гана",
		Point: "Чекит",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Бир квадриллион"},
		{Number: 1000000000000, Value: "Бир триллион"},
		{Number: 1000000000, Value: "Бир миллиард"},
		{Number: 1000000, Value: "Бир миллион"},
		{Number: 1000, Value: "Бир миң"},
		{Number: 100, Value: "Бир жүз"},
		{Number: 90, Value: "Токсон"},
		{Number: 80, Value: "Сексен"},
		{Number: 70, Value: "Жетимиш"},
		{Number: 60, Value: "Алтымыш"},
		{Number: 50, Value: "Элүү"},
		{Number: 40, Value: "Кырк"},
		{Number: 30, Value: "Отуз"},
		{Number: 20, Value: "Жыйырма"},
		{Number: 19, Value: "Он тогуз"},
		{Number: 18, Value: "Он сегиз"},
		{Number: 17, Value: "Он жети"},
		{Number: 16, Value: "Он алты"},
		{Number: 15, Value: "Он беш"},
		{Number: 14, Value: "Он төрт"},
		{Number: 13, Value: "Он үч"},
		{Number: 12, Value: "Он эки"},
		{Number: 11, Value: "Он бир"},
		{Number: 10, Value: "Он"},
		{Number: 9, Value: "Тогуз"},
		{Number: 8, Value: "Сегиз"},
		{Number: 7, Value: "Жети"},
		{Number: 6, Value: "Алты"},
		{Number: 5, Value: "Беш"},
		{Number: 4, Value: "Төрт"},
		{Number: 3, Value: "Үч"},
		{Number: 2, Value: "Эки"},
		{Number: 1, Value: "Бир"},
		{Number: 0, Value: "Нөл"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Бир жүз"},
	},
	LocaleFormatter: &KyrgyzFormatter{},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "биринчи", Suffix: "-нчи", Masculine: "биринчи", Feminine: "биринчи", Neuter: "биринчи"},
		{Number: 2, Word: "экинчи", Suffix: "-нчи", Masculine: "экинчи", Feminine: "экинчи", Neuter: "экинчи"},
		{Number: 3, Word: "үчүнчү", Suffix: "-нчү", Masculine: "үчүнчү", Feminine: "үчүнчү", Neuter: "үчүнчү"},
		{Number: 4, Word: "төртүнчү", Suffix: "-нчү", Masculine: "төртүнчү", Feminine: "төртүнчү", Neuter: "төртүнчү"},
		{Number: 5, Word: "бешинчи", Suffix: "-нчи", Masculine: "бешинчи", Feminine: "бешинчи", Neuter: "бешинчи"},
		{Number: 6, Word: "алтынчы", Suffix: "-нчы", Masculine: "алтынчы", Feminine: "алтынчы", Neuter: "алтынчы"},
		{Number: 7, Word: "жетинчи", Suffix: "-нчи", Masculine: "жетинчи", Feminine: "жетинчи", Neuter: "жетинчи"},
		{Number: 8, Word: "сегизинчи", Suffix: "-нчи", Masculine: "сегизинчи", Feminine: "сегизинчи", Neuter: "сегизинчи"},
		{Number: 9, Word: "тогузунчу", Suffix: "-нчу", Masculine: "тогузунчу", Feminine: "тогузунчу", Neuter: "тогузунчу"},
		{Number: 10, Word: "онунчу", Suffix: "-нчу", Masculine: "онунчу", Feminine: "онунчу", Neuter: "онунчу"},
		{Number: 11, Word: "он биринчи", Suffix: "-нчи", Masculine: "он биринчи", Feminine: "он биринчи", Neuter: "он биринчи"},
		{Number: 12, Word: "он экинчи", Suffix: "-нчи", Masculine: "он экинчи", Feminine: "он экинчи", Neuter: "он экинчи"},
		{Number: 13, Word: "он үчүнчү", Suffix: "-нчү", Masculine: "он үчүнчү", Feminine: "он үчүнчү", Neuter: "он үчүнчү"},
		{Number: 14, Word: "он төртүнчү", Suffix: "-нчү", Masculine: "он төртүнчү", Feminine: "он төртүнчү", Neuter: "он төртүнчү"},
		{Number: 15, Word: "он бешинчи", Suffix: "-нчи", Masculine: "он бешинчи", Feminine: "он бешинчи", Neuter: "он бешинчи"},
		{Number: 16, Word: "он алтынчы", Suffix: "-нчы", Masculine: "он алтынчы", Feminine: "он алтынчы", Neuter: "он алтынчы"},
		{Number: 17, Word: "он жетинчи", Suffix: "-нчи", Masculine: "он жетинчи", Feminine: "он жетинчи", Neuter: "он жетинчи"},
		{Number: 18, Word: "он сегизинчи", Suffix: "-нчи", Masculine: "он сегизинчи", Feminine: "он сегизинчи", Neuter: "он сегизинчи"},
		{Number: 19, Word: "он тогузунчу", Suffix: "-нчу", Masculine: "он тогузунчу", Feminine: "он тогузунчу", Neuter: "он тогузунчу"},
		{Number: 20, Word: "жыйырманчы", Suffix: "-нчы", Masculine: "жыйырманчы", Feminine: "жыйырманчы", Neuter: "жыйырманчы"},
		{Number: 21, Word: "жыйырма биринчи", Suffix: "-нчи", Masculine: "жыйырма биринчи", Feminine: "жыйырма биринчи", Neuter: "жыйырма биринчи"},
		{Number: 30, Word: "отузунчу", Suffix: "-нчу", Masculine: "отузунчу", Feminine: "отузунчу", Neuter: "отузунчу"},
		{Number: 40, Word: "кыркынчы", Suffix: "-нчы", Masculine: "кыркынчы", Feminine: "кыркынчы", Neuter: "кыркынчы"},
		{Number: 50, Word: "элүүнчү", Suffix: "-нчү", Masculine: "элүүнчү", Feminine: "элүүнчү", Neuter: "элүүнчү"},
		{Number: 60, Word: "алтымышынчы", Suffix: "-нчы", Masculine: "алтымышынчы", Feminine: "алтымышынчы", Neuter: "алтымышынчы"},
		{Number: 70, Word: "жетимишинчи", Suffix: "-нчи", Masculine: "жетимишинчи", Feminine: "жетимишинчи", Neuter: "жетимишинчи"},
		{Number: 80, Word: "сексенинчи", Suffix: "-нчи", Masculine: "сексенинчи", Feminine: "сексенинчи", Neuter: "сексенинчи"},
		{Number: 90, Word: "токсонунчу", Suffix: "-нчу", Masculine: "токсонунчу", Feminine: "токсонунчу", Neuter: "токсонунчу"},
		{Number: 100, Word: "жүзүнчү", Suffix: "-нчү", Masculine: "жүзүнчү", Feminine: "жүзүнчү", Neuter: "жүзүнчү"},
		{Number: 1000, Word: "миңинчи", Suffix: "-нчи", Masculine: "миңинчи", Feminine: "миңинчи", Neuter: "миңинчи"},
	},
}

// KyrgyzFormatter handles Kyrgyz formatting
type KyrgyzFormatter struct{}

func (f *KyrgyzFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *KyrgyzFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	// In Kyrgyz, Som doesn't change form for singular/plural
	return result + " " + currencyName
}

func (f *KyrgyzFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *KyrgyzFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	// In Kyrgyz, Tyiyn doesn't change form for singular/plural
	return result + " " + fractionName
}

func (f *KyrgyzFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *KyrgyzFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	return amount.Truncate(int32(precision))
}

func (f *KyrgyzFormatter) FormatDecimalNumber(amount float64) string {
	return FormatPolishDecimal(amount)
}
func (f *KyrgyzFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	return FormatPolishCurrency(amount, currencySymbol)
}
