package locale

import (
	"github.com/shopspring/decimal"
)

// BGBGLocale is a NumI18NLocale configured for Bulgarian (Bulgaria) - bg-BG
var BGBGLocale = NumI18NLocale{
	LocaleFormatter: &BulgarianFormatter{},
	Currency: Currency{
		Name:     "лев",
		Plural:   "лева",
		Singular: "лев",
		Symbol:   "лв",
		FractionUnit: FractionUnit{
			Name:     "стотинка",
			Plural:   "стотинки",
			Singular: "стотинка",
			Symbol:   "ст.",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Bulgaria",
		Currency:       "BGN",
		ISO3166Alpha2:  "BG",
		ISO3166Alpha3:  "BGR",
		ISO3166Numeric: "100",
		Locale:         "bg-BG",
		Timezone:       []string{"Europe/Sofia"},
		Language:       "bg",
	},
	Texts: Texts{
		And:   "и",
		Minus: "минус",
		Only:  "само",
		Point: "точка",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Квадриллион"},
		{Number: 1000000000000, Value: "Трилион"},
		{Number: 1000000000, Value: "Милиард"},
		{Number: 1000000, Value: "Милион"},
		{Number: 1000, Value: "Хиляда"},
		{Number: 100, Value: "Сто"},
		{Number: 90, Value: "Сто деветдесет"},
		{Number: 80, Value: "Осемдесет"},
		{Number: 70, Value: "Седемдесет"},
		{Number: 60, Value: "Шестдесет"},
		{Number: 50, Value: "Петдесет"},
		{Number: 40, Value: "Четиридесет"},
		{Number: 30, Value: "Тридесет"},
		{Number: 20, Value: "Двадесет"},
		{Number: 19, Value: "Деветнадесет"},
		{Number: 18, Value: "Осемнадесет"},
		{Number: 17, Value: "Седемнадесет"},
		{Number: 16, Value: "Шестнадесет"},
		{Number: 15, Value: "Петнадесет"},
		{Number: 14, Value: "Четирнадесет"},
		{Number: 13, Value: "Тринадесет"},
		{Number: 12, Value: "Дванадесет"},
		{Number: 11, Value: "Единадесет"},
		{Number: 10, Value: "Десет"},
		{Number: 9, Value: "Девет"},
		{Number: 8, Value: "Осем"},
		{Number: 7, Value: "Седем"},
		{Number: 6, Value: "Шест"},
		{Number: 5, Value: "Пет"},
		{Number: 4, Value: "Четири"},
		{Number: 3, Value: "Три"},
		{Number: 2, Value: "Две"},
		{Number: 1, Value: "Едно"},
		{Number: 0, Value: "Нула"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Сто"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "първи", Suffix: "", Masculine: "първи", Feminine: "първа", Neuter: "първо"},
		{Number: 2, Word: "втори", Suffix: "", Masculine: "втори", Feminine: "втора", Neuter: "второ"},
		{Number: 3, Word: "трети", Suffix: "", Masculine: "трети", Feminine: "трета", Neuter: "трето"},
		{Number: 4, Word: "четвърти", Suffix: "", Masculine: "четвърти", Feminine: "четвърта", Neuter: "четвърто"},
		{Number: 5, Word: "пети", Suffix: "", Masculine: "пети", Feminine: "пета", Neuter: "пето"},
		{Number: 6, Word: "шести", Suffix: "", Masculine: "шести", Feminine: "шеста", Neuter: "шесто"},
		{Number: 7, Word: "седми", Suffix: "", Masculine: "седми", Feminine: "седма", Neuter: "седмо"},
		{Number: 8, Word: "осми", Suffix: "", Masculine: "осми", Feminine: "осма", Neuter: "осмо"},
		{Number: 9, Word: "девети", Suffix: "", Masculine: "девети", Feminine: "девета", Neuter: "девето"},
		{Number: 10, Word: "десети", Suffix: "", Masculine: "десети", Feminine: "десета", Neuter: "десето"},
		{Number: 11, Word: "единадесети", Suffix: "", Masculine: "единадесети", Feminine: "единадесета", Neuter: "единадесето"},
		{Number: 12, Word: "дванадесети", Suffix: "", Masculine: "дванадесети", Feminine: "дванадесета", Neuter: "дванадесето"},
		{Number: 13, Word: "тринадесети", Suffix: "", Masculine: "тринадесети", Feminine: "тринадесета", Neuter: "тринадесето"},
		{Number: 14, Word: "четирнадесети", Suffix: "", Masculine: "четирнадесети", Feminine: "четирнадесета", Neuter: "четирнадесето"},
		{Number: 15, Word: "петнадесети", Suffix: "", Masculine: "петнадесети", Feminine: "петнадесета", Neuter: "петнадесето"},
		{Number: 16, Word: "шестнадесети", Suffix: "", Masculine: "шестнадесети", Feminine: "шестнадесета", Neuter: "шестнадесето"},
		{Number: 17, Word: "седемнадесети", Suffix: "", Masculine: "седемнадесети", Feminine: "седемнадесета", Neuter: "седемнадесето"},
		{Number: 18, Word: "осемнадесети", Suffix: "", Masculine: "осемнадесети", Feminine: "осемнадесета", Neuter: "осемнадесето"},
		{Number: 19, Word: "деветнадесети", Suffix: "", Masculine: "деветнадесети", Feminine: "деветнадесета", Neuter: "деветнадесето"},
		{Number: 20, Word: "двадесети", Suffix: "", Masculine: "двадесети", Feminine: "двадесета", Neuter: "двадесето"},
		{Number: 21, Word: "двадесет и първи", Suffix: "", Masculine: "двадесет и първи", Feminine: "двадесет и първа", Neuter: "двадесет и първо"},
		{Number: 30, Word: "тридесети", Suffix: "", Masculine: "тридесети", Feminine: "тридесета", Neuter: "тридесето"},
		{Number: 40, Word: "четиридесети", Suffix: "", Masculine: "четиридесети", Feminine: "четиридесета", Neuter: "четиридесето"},
		{Number: 50, Word: "петдесети", Suffix: "", Masculine: "петдесети", Feminine: "петдесета", Neuter: "петдесето"},
		{Number: 60, Word: "шестдесети", Suffix: "", Masculine: "шестдесети", Feminine: "шестдесета", Neuter: "шестдесето"},
		{Number: 70, Word: "седемдесети", Suffix: "", Masculine: "седемдесети", Feminine: "седемдесета", Neuter: "седемдесето"},
		{Number: 80, Word: "осемдесети", Suffix: "", Masculine: "осемдесети", Feminine: "осемдесета", Neuter: "осемдесето"},
		{Number: 90, Word: "деветдесети", Suffix: "", Masculine: "деветдесети", Feminine: "деветдесета", Neuter: "деветдесето"},
		{Number: 100, Word: "стогодишен", Suffix: "", Masculine: "стогодишен", Feminine: "стогодишна", Neuter: "стогодишно"},
		{Number: 1000, Word: "хилядни", Suffix: "", Masculine: "хилядни", Feminine: "хилядна", Neuter: "хилядно"},
	},
}

// BulgarianFormatter handles Bulgarian formatting
type BulgarianFormatter struct{}

func (f *BulgarianFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *BulgarianFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *BulgarianFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *BulgarianFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *BulgarianFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *BulgarianFormatter) ChopDecimal(d decimal.Decimal, precision int) decimal.Decimal {
	return d.Truncate(int32(precision))
}

func (f *BulgarianFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *BulgarianFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
