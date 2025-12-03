package locale

import (
	"github.com/shopspring/decimal"
)

// SRRSLocale represents the Serbian (Serbia) locale
var SRRSLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Serbian Dinar",
		Plural:   "Динара",
		Singular: "Динар",
		Symbol:   "РСД",
		FractionUnit: FractionUnit{
			Name:     "Para",
			Plural:   "Паре",
			Singular: "Пара",
			Symbol:   "п",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Serbia",
		Currency:       "RSD",
		ISO3166Alpha2:  "RS",
		ISO3166Alpha3:  "SRB",
		ISO3166Numeric: "688",
		Locale:         "sr-RS",
		Timezone:       []string{"Europe/Belgrade"},
		Language:       "sr",
		Emoji:          "🇷🇸",
	},
	Texts: Texts{
		And:   "и",
		Minus: "минус",
		Only:  "само",
		Point: "тачка",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "хиљада билијарди"},
		{Number: 1000000000000, Value: "билијарда"},
		{Number: 1000000000, Value: "милијарда"},
		{Number: 1000000, Value: "милион"},
		{Number: 100000, Value: "сто хиљада"},
		{Number: 90000, Value: "деведесет хиљада"},
		{Number: 80000, Value: "осамдесет хиљада"},
		{Number: 70000, Value: "седамдесет хиљада"},
		{Number: 60000, Value: "шездесет хиљада"},
		{Number: 50000, Value: "педесет хиљада"},
		{Number: 40000, Value: "четрдесет хиљада"},
		{Number: 30000, Value: "тридесет хиљада"},
		{Number: 20000, Value: "двадесет хиљада"},
		{Number: 19000, Value: "деветнаест хиљада"},
		{Number: 18000, Value: "осамнаест хиљада"},
		{Number: 17000, Value: "седамнаест хиљада"},
		{Number: 16000, Value: "шеснаест хиљада"},
		{Number: 15000, Value: "петнаест хиљада"},
		{Number: 14000, Value: "четрнаест хиљада"},
		{Number: 13000, Value: "тринаест хиљада"},
		{Number: 12000, Value: "дванаест хиљада"},
		{Number: 11000, Value: "једанаест хиљада"},
		{Number: 10000, Value: "десет хиљада"},
		{Number: 9000, Value: "девет хиљада"},
		{Number: 8000, Value: "осам хиљада"},
		{Number: 7000, Value: "седам хиљада"},
		{Number: 6000, Value: "шест хиљада"},
		{Number: 5000, Value: "пет хиљада"},
		{Number: 4000, Value: "четири хиљада"},
		{Number: 3000, Value: "три хиљада"},
		{Number: 2000, Value: "две хиљаде"},
		{Number: 1000, Value: "хиљада"},
		{Number: 900, Value: "деветсто"},
		{Number: 800, Value: "осамсто"},
		{Number: 700, Value: "седамсто"},
		{Number: 600, Value: "шестсто"},
		{Number: 500, Value: "петсто"},
		{Number: 400, Value: "четиристо"},
		{Number: 300, Value: "тристо"},
		{Number: 200, Value: "двеста"},
		{Number: 100, Value: "сто"},
		{Number: 99, Value: "деведесет девет"},
		{Number: 98, Value: "деведесет осам"},
		{Number: 97, Value: "деведесет седам"},
		{Number: 96, Value: "деведесет шест"},
		{Number: 95, Value: "деведесет пет"},
		{Number: 94, Value: "деведесет четири"},
		{Number: 93, Value: "деведесет три"},
		{Number: 92, Value: "деведесет два"},
		{Number: 91, Value: "деведесет један"},
		{Number: 90, Value: "деведесет"},
		{Number: 89, Value: "осамдесет девет"},
		{Number: 88, Value: "осамдесет осам"},
		{Number: 87, Value: "осамдесет седам"},
		{Number: 86, Value: "осамдесет шест"},
		{Number: 85, Value: "осамдесет пет"},
		{Number: 84, Value: "осамдесет четири"},
		{Number: 83, Value: "осамдесет три"},
		{Number: 82, Value: "осамдесет два"},
		{Number: 81, Value: "осамдесет један"},
		{Number: 80, Value: "осамдесет"},
		{Number: 79, Value: "седамдесет девет"},
		{Number: 78, Value: "седамдесет осам"},
		{Number: 77, Value: "седамдесет седам"},
		{Number: 76, Value: "седамдесет шест"},
		{Number: 75, Value: "седамдесет пет"},
		{Number: 74, Value: "седамдесет четири"},
		{Number: 73, Value: "седамдесет три"},
		{Number: 72, Value: "седамдесет два"},
		{Number: 71, Value: "седамдесет један"},
		{Number: 70, Value: "седамдесет"},
		{Number: 69, Value: "шездесет девет"},
		{Number: 68, Value: "шездесет осам"},
		{Number: 67, Value: "шездесет седам"},
		{Number: 66, Value: "шездесет шест"},
		{Number: 65, Value: "шездесет пет"},
		{Number: 64, Value: "шездесет четири"},
		{Number: 63, Value: "шездесет три"},
		{Number: 62, Value: "шездесет два"},
		{Number: 61, Value: "шездесет један"},
		{Number: 60, Value: "шездесет"},
		{Number: 59, Value: "педесет девет"},
		{Number: 58, Value: "педесет осам"},
		{Number: 57, Value: "педесет седам"},
		{Number: 56, Value: "педесет шест"},
		{Number: 55, Value: "педесет пет"},
		{Number: 54, Value: "педесет четири"},
		{Number: 53, Value: "педесет три"},
		{Number: 52, Value: "педесет два"},
		{Number: 51, Value: "педесет један"},
		{Number: 50, Value: "педесет"},
		{Number: 49, Value: "четрдесет девет"},
		{Number: 48, Value: "четрдесет осам"},
		{Number: 47, Value: "четрдесет седам"},
		{Number: 46, Value: "четрдесет шест"},
		{Number: 45, Value: "четрдесет пет"},
		{Number: 44, Value: "четрдесет четири"},
		{Number: 43, Value: "четрдесет три"},
		{Number: 42, Value: "четрдесет два"},
		{Number: 41, Value: "четрдесет један"},
		{Number: 40, Value: "четрдесет"},
		{Number: 39, Value: "тридесет девет"},
		{Number: 38, Value: "тридесет осам"},
		{Number: 37, Value: "тридесет седам"},
		{Number: 36, Value: "тридесет шест"},
		{Number: 35, Value: "тридесет пет"},
		{Number: 34, Value: "тридесет четири"},
		{Number: 33, Value: "тридесет три"},
		{Number: 32, Value: "тридесет два"},
		{Number: 31, Value: "тридесет један"},
		{Number: 30, Value: "тридесет"},
		{Number: 29, Value: "двадесет девет"},
		{Number: 28, Value: "двадесет осам"},
		{Number: 27, Value: "двадесет седам"},
		{Number: 26, Value: "двадесет шест"},
		{Number: 25, Value: "двадесет пет"},
		{Number: 24, Value: "двадесет четири"},
		{Number: 23, Value: "двадесет три"},
		{Number: 22, Value: "двадесет два"},
		{Number: 21, Value: "двадесет један"},
		{Number: 20, Value: "двадесет"},
		{Number: 19, Value: "деветнаест"},
		{Number: 18, Value: "осамнаест"},
		{Number: 17, Value: "седамнаест"},
		{Number: 16, Value: "шеснаест"},
		{Number: 15, Value: "петнаест"},
		{Number: 14, Value: "четрнаест"},
		{Number: 13, Value: "тринаест"},
		{Number: 12, Value: "дванаест"},
		{Number: 11, Value: "једанаест"},
		{Number: 10, Value: "десет"},
		{Number: 9, Value: "девет"},
		{Number: 8, Value: "осам"},
		{Number: 7, Value: "седам"},
		{Number: 6, Value: "шест"},
		{Number: 5, Value: "пет"},
		{Number: 4, Value: "четири"},
		{Number: 3, Value: "три"},
		{Number: 2, Value: "два"},
		{Number: 1, Value: "један"},
		{Number: 0, Value: "нула"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Сто"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "први", Suffix: ".", Masculine: "први", Feminine: "прва", Neuter: "прво"},
		{Number: 2, Word: "други", Suffix: ".", Masculine: "други", Feminine: "друга", Neuter: "друго"},
		{Number: 3, Word: "трећи", Suffix: ".", Masculine: "трећи", Feminine: "трећа", Neuter: "треће"},
		{Number: 4, Word: "четврти", Suffix: ".", Masculine: "четврти", Feminine: "четврта", Neuter: "четврто"},
		{Number: 5, Word: "пети", Suffix: ".", Masculine: "пети", Feminine: "пета", Neuter: "пето"},
		{Number: 6, Word: "шести", Suffix: ".", Masculine: "шести", Feminine: "шеста", Neuter: "шесто"},
		{Number: 7, Word: "седми", Suffix: ".", Masculine: "седми", Feminine: "седма", Neuter: "седмо"},
		{Number: 8, Word: "осми", Suffix: ".", Masculine: "осми", Feminine: "осма", Neuter: "осмо"},
		{Number: 9, Word: "девети", Suffix: ".", Masculine: "девети", Feminine: "девета", Neuter: "девето"},
		{Number: 10, Word: "десети", Suffix: ".", Masculine: "десети", Feminine: "десета", Neuter: "десето"},
		{Number: 11, Word: "једанаести", Suffix: ".", Masculine: "једанаести", Feminine: "једанаеста", Neuter: "једанаесто"},
		{Number: 12, Word: "дванаести", Suffix: ".", Masculine: "дванаести", Feminine: "дванаеста", Neuter: "дванаесто"},
		{Number: 20, Word: "двадесети", Suffix: ".", Masculine: "двадесети", Feminine: "двадесета", Neuter: "двадесето"},
		{Number: 21, Word: "двадесет први", Suffix: ".", Masculine: "двадесет први", Feminine: "двадесет прва", Neuter: "двадесет прво"},
		{Number: 30, Word: "тридесети", Suffix: ".", Masculine: "тридесети", Feminine: "тридесета", Neuter: "тридесето"},
		{Number: 50, Word: "педесети", Suffix: ".", Masculine: "педесети", Feminine: "педесета", Neuter: "педесето"},
		{Number: 100, Word: "стоти", Suffix: ".", Masculine: "стоти", Feminine: "стота", Neuter: "сто"},
		{Number: 1000, Word: "хиљадити", Suffix: ".", Masculine: "хиљадити", Feminine: "хиљадита", Neuter: "хиљадито"},
	},
	LocaleFormatter: &SerbianSerbiaFormatter{},
}

// SerbianSerbiaFormatter handles Serbian (Serbia) formatting
type SerbianSerbiaFormatter struct{}

func (f *SerbianSerbiaFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *SerbianSerbiaFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *SerbianSerbiaFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *SerbianSerbiaFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *SerbianSerbiaFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *SerbianSerbiaFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	return amount.Truncate(int32(precision))
}

func (f *SerbianSerbiaFormatter) FormatDecimalNumber(amount float64) string {
	return FormatEuropeanDecimal(amount)
}
func (f *SerbianSerbiaFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with European conventions (period separators, comma decimal, prefix symbol)
	return FormatEuropeanCurrency(amount, currencySymbol)
}
