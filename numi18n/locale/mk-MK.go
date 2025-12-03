package locale

import (
	"github.com/shopspring/decimal"
)

// MKMKLocale represents the Macedonian (North Macedonia) locale
var MKMKLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Денар",
		Plural:   "Денари",
		Singular: "Денар",
		Symbol:   "ден",
		FractionUnit: FractionUnit{
			Name:     "Дени",
			Plural:   "Дени",
			Singular: "Ден",
			Symbol:   "д",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "North Macedonia",
		Currency:       "MKD",
		ISO3166Alpha2:  "MK",
		ISO3166Alpha3:  "MKD",
		ISO3166Numeric: "807",
		Locale:         "mk-MK",
		Timezone:       []string{"Europe/Skopje"},
		Language:       "mk",
		Emoji:          "🇲🇰",
		PhoneCode:      "+389",
		Domain:         ".mk",
	},
	Texts: Texts{
		And:   "и",
		Minus: "минус",
		Only:  "само",
		Point: "точка",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "билијарда"},
		{Number: 1000000000000, Value: "билион"},
		{Number: 1000000000, Value: "милијарда"},
		{Number: 1000000, Value: "милион"},
		{Number: 100000, Value: "сто илјади"},
		{Number: 90000, Value: "деведесет илјади"},
		{Number: 80000, Value: "осумдесет илјади"},
		{Number: 70000, Value: "седумдесет илјади"},
		{Number: 60000, Value: "шеесет илјади"},
		{Number: 50000, Value: "педесет илјади"},
		{Number: 40000, Value: "четириесет илјади"},
		{Number: 30000, Value: "триесет илјади"},
		{Number: 20000, Value: "дваесет илјади"},
		{Number: 19000, Value: "деветнаесет илјади"},
		{Number: 18000, Value: "осумнаесет илјади"},
		{Number: 17000, Value: "седумнаесет илјади"},
		{Number: 16000, Value: "шеснаесет илјади"},
		{Number: 15000, Value: "петнаесет илјади"},
		{Number: 14000, Value: "четиринаесет илјади"},
		{Number: 13000, Value: "тринаесет илјади"},
		{Number: 12000, Value: "дванаесет илјади"},
		{Number: 11000, Value: "единаесет илјади"},
		{Number: 10000, Value: "десет илјади"},
		{Number: 9000, Value: "девет илјади"},
		{Number: 8000, Value: "осум илјади"},
		{Number: 7000, Value: "седум илјади"},
		{Number: 6000, Value: "шест илјади"},
		{Number: 5000, Value: "пет илјади"},
		{Number: 4000, Value: "четири илјади"},
		{Number: 3000, Value: "три илјади"},
		{Number: 2000, Value: "две илјади"},
		{Number: 1000, Value: "илјада"},
		{Number: 900, Value: "деветстотини"},
		{Number: 800, Value: "осумстотини"},
		{Number: 700, Value: "седумстотини"},
		{Number: 600, Value: "шестотини"},
		{Number: 500, Value: "петстотини"},
		{Number: 400, Value: "четиристотини"},
		{Number: 300, Value: "тристотини"},
		{Number: 200, Value: "двесте"},
		{Number: 100, Value: "сто"},
		{Number: 99, Value: "деведесет и девет"},
		{Number: 98, Value: "деведесет и осум"},
		{Number: 97, Value: "деведесет и седум"},
		{Number: 96, Value: "деведесет и шест"},
		{Number: 95, Value: "деведесет и пет"},
		{Number: 94, Value: "деведесет и четири"},
		{Number: 93, Value: "деведесет и три"},
		{Number: 92, Value: "деведесет и две"},
		{Number: 91, Value: "деведесет и еден"},
		{Number: 90, Value: "деведесет"},
		{Number: 89, Value: "осумдесет и девет"},
		{Number: 88, Value: "осумдесет и осум"},
		{Number: 87, Value: "осумдесет и седум"},
		{Number: 86, Value: "осумдесет и шест"},
		{Number: 85, Value: "осумдесет и пет"},
		{Number: 84, Value: "осумдесет и четири"},
		{Number: 83, Value: "осумдесет и три"},
		{Number: 82, Value: "осумдесет и две"},
		{Number: 81, Value: "осумдесет и еден"},
		{Number: 80, Value: "осумдесет"},
		{Number: 79, Value: "седумдесет и девет"},
		{Number: 78, Value: "седумдесет и осум"},
		{Number: 77, Value: "седумдесет и седум"},
		{Number: 76, Value: "седумдесет и шест"},
		{Number: 75, Value: "седумдесет и пет"},
		{Number: 74, Value: "седумдесет и четири"},
		{Number: 73, Value: "седумдесет и три"},
		{Number: 72, Value: "седумдесет и две"},
		{Number: 71, Value: "седумдесет и еден"},
		{Number: 70, Value: "седумдесет"},
		{Number: 69, Value: "шеесет и девет"},
		{Number: 68, Value: "шеесет и осум"},
		{Number: 67, Value: "шеесет и седум"},
		{Number: 66, Value: "шеесет и шест"},
		{Number: 65, Value: "шеесет и пет"},
		{Number: 64, Value: "шеесет и четири"},
		{Number: 63, Value: "шеесет и три"},
		{Number: 62, Value: "шеесет и две"},
		{Number: 61, Value: "шеесет и еден"},
		{Number: 60, Value: "шеесет"},
		{Number: 59, Value: "педесет и девет"},
		{Number: 58, Value: "педесет и осум"},
		{Number: 57, Value: "педесет и седум"},
		{Number: 56, Value: "педесет и шест"},
		{Number: 55, Value: "педесет и пет"},
		{Number: 54, Value: "педесет и четири"},
		{Number: 53, Value: "педесет и три"},
		{Number: 52, Value: "педесет и две"},
		{Number: 51, Value: "педесет и еден"},
		{Number: 50, Value: "педесет"},
		{Number: 49, Value: "четириесет и девет"},
		{Number: 48, Value: "четириесет и осум"},
		{Number: 47, Value: "четириесет и седум"},
		{Number: 46, Value: "четириесет и шест"},
		{Number: 45, Value: "четириесет и пет"},
		{Number: 44, Value: "четириесет и четири"},
		{Number: 43, Value: "четириесет и три"},
		{Number: 42, Value: "четириесет и две"},
		{Number: 41, Value: "четириесет и еден"},
		{Number: 40, Value: "четириесет"},
		{Number: 39, Value: "триесет и девет"},
		{Number: 38, Value: "триесет и осум"},
		{Number: 37, Value: "триесет и седум"},
		{Number: 36, Value: "триесет и шест"},
		{Number: 35, Value: "триесет и пет"},
		{Number: 34, Value: "триесет и четири"},
		{Number: 33, Value: "триесет и три"},
		{Number: 32, Value: "триесет и две"},
		{Number: 31, Value: "триесет и еден"},
		{Number: 30, Value: "триесет"},
		{Number: 29, Value: "дваесет и девет"},
		{Number: 28, Value: "дваесет и осум"},
		{Number: 27, Value: "дваесет и седум"},
		{Number: 26, Value: "дваесет и шест"},
		{Number: 25, Value: "дваесет и пет"},
		{Number: 24, Value: "дваесет и четири"},
		{Number: 23, Value: "дваесет и три"},
		{Number: 22, Value: "дваесет и две"},
		{Number: 21, Value: "дваесет и еден"},
		{Number: 20, Value: "дваесет"},
		{Number: 19, Value: "деветнаесет"},
		{Number: 18, Value: "осумнаесет"},
		{Number: 17, Value: "седумнаесет"},
		{Number: 16, Value: "шеснаесет"},
		{Number: 15, Value: "петнаесет"},
		{Number: 14, Value: "четиринаесет"},
		{Number: 13, Value: "тринаесет"},
		{Number: 12, Value: "дванаесет"},
		{Number: 11, Value: "единаесет"},
		{Number: 10, Value: "десет"},
		{Number: 9, Value: "девет"},
		{Number: 8, Value: "осум"},
		{Number: 7, Value: "седум"},
		{Number: 6, Value: "шест"},
		{Number: 5, Value: "пет"},
		{Number: 4, Value: "четири"},
		{Number: 3, Value: "три"},
		{Number: 2, Value: "две"},
		{Number: 1, Value: "еден"},
		{Number: 0, Value: "нула"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "први", Suffix: "-ви", Masculine: "први", Feminine: "прва", Neuter: "прво"},
		{Number: 2, Word: "втор", Suffix: "-ри", Masculine: "втор", Feminine: "втора", Neuter: "второ"},
		{Number: 3, Word: "трет", Suffix: "-ти", Masculine: "трет", Feminine: "трета", Neuter: "трето"},
		{Number: 4, Word: "четврт", Suffix: "-ти", Masculine: "четврт", Feminine: "четврта", Neuter: "четврто"},
		{Number: 5, Word: "петт", Suffix: "-ти", Masculine: "петт", Feminine: "петта", Neuter: "петто"},
		{Number: 6, Word: "шест", Suffix: "-ти", Masculine: "шест", Feminine: "шеста", Neuter: "шесто"},
		{Number: 7, Word: "седми", Suffix: "-ми", Masculine: "седми", Feminine: "седма", Neuter: "седмо"},
		{Number: 8, Word: "осми", Suffix: "-ми", Masculine: "осми", Feminine: "осма", Neuter: "осмо"},
		{Number: 9, Word: "деветти", Suffix: "-ти", Masculine: "деветти", Feminine: "деветта", Neuter: "деветто"},
		{Number: 10, Word: "десетти", Suffix: "-ти", Masculine: "десетти", Feminine: "десетта", Neuter: "десетто"},
		{Number: 11, Word: "единаесетти", Suffix: "-ти", Masculine: "единаесетти", Feminine: "единаесетта", Neuter: "единаесетто"},
		{Number: 12, Word: "дванаесетти", Suffix: "-ти", Masculine: "дванаесетти", Feminine: "дванаесетта", Neuter: "дванаесетто"},
		{Number: 20, Word: "дваесетти", Suffix: "-ти", Masculine: "дваесетти", Feminine: "дваесетта", Neuter: "дваесетто"},
		{Number: 21, Word: "дваесет и први", Suffix: "-ви", Masculine: "дваесет и први", Feminine: "дваесет и прва", Neuter: "дваесет и прво"},
		{Number: 30, Word: "триесетти", Suffix: "-ти", Masculine: "триесетти", Feminine: "триесетта", Neuter: "триесетто"},
		{Number: 100, Word: "стотти", Suffix: "-ти", Masculine: "стотти", Feminine: "стотта", Neuter: "стотто"},
		{Number: 1000, Word: "илјадитти", Suffix: "-ти", Masculine: "илјадитти", Feminine: "илјадитта", Neuter: "илјадитто"},
	},
	LocaleFormatter: &MacedonianFormatter{},
}

// MacedonianFormatter handles Macedonian formatting
type MacedonianFormatter struct{}

func (f *MacedonianFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *MacedonianFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *MacedonianFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *MacedonianFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *MacedonianFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *MacedonianFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 0
	}
	return amount.Truncate(int32(precision))
}

func (f *MacedonianFormatter) FormatDecimalNumber(amount float64) string {
	return FormatPolishDecimal(amount)
}
func (f *MacedonianFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Polish conventions (comma separators, period decimal, prefix symbol)
	return FormatPolishCurrency(amount, currencySymbol)
}
