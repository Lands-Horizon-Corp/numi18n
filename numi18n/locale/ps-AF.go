package locale

import (
	"github.com/shopspring/decimal"
)

// PSAFLocale represents the Pashto (Afghanistan) locale
var PSAFLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "افغانۍ",
		Plural:   "افغانۍ",
		Singular: "افغانۍ",
		Symbol:   "؋",
		FractionUnit: FractionUnit{
			Name:     "پول",
			Plural:   "پول",
			Singular: "پول",
			Symbol:   "پ",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Afghanistan",
		Currency:       "AFN",
		ISO3166Alpha2:  "AF",
		ISO3166Alpha3:  "AFG",
		ISO3166Numeric: "004",
		Locale:         "ps-AF",
		Timezone:       []string{"Asia/Kabul"},
		Language:       "ps",
		Emoji:          "🇦🇫",
		PhoneCode:      "+93",
		Domain:         ".af",
	},
	Texts: Texts{
		And:   "او",
		Minus: "منفي",
		Only:  "یوازې",
		Point: "ټکي",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "یو ترلیون"},
		{Number: 1000000000000, Value: "یو بلیون"},
		{Number: 1000000000, Value: "یو ملیون"},
		{Number: 1000000, Value: "یو میلیون"},
		{Number: 100000, Value: "یو سل زره"},
		{Number: 90000, Value: "نوی زره"},
		{Number: 80000, Value: "اتیا زره"},
		{Number: 70000, Value: "اویا زره"},
		{Number: 60000, Value: "شپیته زره"},
		{Number: 50000, Value: "پنځوس زره"},
		{Number: 40000, Value: "څلویښت زره"},
		{Number: 30000, Value: "دیرش زره"},
		{Number: 20000, Value: "شل زره"},
		{Number: 19000, Value: "نولس زره"},
		{Number: 18000, Value: "اتلس زره"},
		{Number: 17000, Value: "اولس زره"},
		{Number: 16000, Value: "شپاړس زره"},
		{Number: 15000, Value: "پنځلس زره"},
		{Number: 14000, Value: "څوارلس زره"},
		{Number: 13000, Value: "دیارلس زره"},
		{Number: 12000, Value: "دولس زره"},
		{Number: 11000, Value: "یولس زره"},
		{Number: 10000, Value: "لس زره"},
		{Number: 9000, Value: "نهه زره"},
		{Number: 8000, Value: "اته زره"},
		{Number: 7000, Value: "اووه زره"},
		{Number: 6000, Value: "شپږ زره"},
		{Number: 5000, Value: "پنځه زره"},
		{Number: 4000, Value: "څلور زره"},
		{Number: 3000, Value: "درې زره"},
		{Number: 2000, Value: "دوه زره"},
		{Number: 1000, Value: "یو زره"},
		{Number: 900, Value: "نهه سوه"},
		{Number: 800, Value: "اته سوه"},
		{Number: 700, Value: "اووه سوه"},
		{Number: 600, Value: "شپږ سوه"},
		{Number: 500, Value: "پنځه سوه"},
		{Number: 400, Value: "څلور سوه"},
		{Number: 300, Value: "درې سوه"},
		{Number: 200, Value: "دوه سوه"},
		{Number: 100, Value: "یو سل"},
		{Number: 99, Value: "نوی نهه"},
		{Number: 98, Value: "نوی اته"},
		{Number: 97, Value: "نوی اووه"},
		{Number: 96, Value: "نوی شپږ"},
		{Number: 95, Value: "نوی پنځه"},
		{Number: 94, Value: "نوی څلور"},
		{Number: 93, Value: "نوی درې"},
		{Number: 92, Value: "نوی دوه"},
		{Number: 91, Value: "نوی یو"},
		{Number: 90, Value: "نوی"},
		{Number: 89, Value: "اتیا نهه"},
		{Number: 88, Value: "اتیا اته"},
		{Number: 87, Value: "اتیا اووه"},
		{Number: 86, Value: "اتیا شپږ"},
		{Number: 85, Value: "اتیا پنځه"},
		{Number: 84, Value: "اتیا څلور"},
		{Number: 83, Value: "اتیا درې"},
		{Number: 82, Value: "اتیا دوه"},
		{Number: 81, Value: "اتیا یو"},
		{Number: 80, Value: "اتیا"},
		{Number: 79, Value: "اویا نهه"},
		{Number: 78, Value: "اویا اته"},
		{Number: 77, Value: "اویا اووه"},
		{Number: 76, Value: "اویا شپږ"},
		{Number: 75, Value: "اویا پنځه"},
		{Number: 74, Value: "اویا څلور"},
		{Number: 73, Value: "اویا درې"},
		{Number: 72, Value: "اویا دوه"},
		{Number: 71, Value: "اویا یو"},
		{Number: 70, Value: "اویا"},
		{Number: 69, Value: "شپیته نهه"},
		{Number: 68, Value: "شپیته اته"},
		{Number: 67, Value: "شپیته اووه"},
		{Number: 66, Value: "شپیته شپږ"},
		{Number: 65, Value: "شپیته پنځه"},
		{Number: 64, Value: "شپیته څلور"},
		{Number: 63, Value: "شپیته درې"},
		{Number: 62, Value: "شپیته دوه"},
		{Number: 61, Value: "شپیته یو"},
		{Number: 60, Value: "شپیته"},
		{Number: 59, Value: "پنځوس نهه"},
		{Number: 58, Value: "پنځوس اته"},
		{Number: 57, Value: "پنځوس اووه"},
		{Number: 56, Value: "پنځوس شپږ"},
		{Number: 55, Value: "پنځوس پنځه"},
		{Number: 54, Value: "پنځوس څلور"},
		{Number: 53, Value: "پنځوس درې"},
		{Number: 52, Value: "پنځوس دوه"},
		{Number: 51, Value: "پنځوس یو"},
		{Number: 50, Value: "پنځوس"},
		{Number: 49, Value: "څلویښت نهه"},
		{Number: 48, Value: "څلویښت اته"},
		{Number: 47, Value: "څلویښت اووه"},
		{Number: 46, Value: "څلویښت شپږ"},
		{Number: 45, Value: "څلویښت پنځه"},
		{Number: 44, Value: "څلویښت څلور"},
		{Number: 43, Value: "څلویښت درې"},
		{Number: 42, Value: "څلویښت دوه"},
		{Number: 41, Value: "څلویښت یو"},
		{Number: 40, Value: "څلویښت"},
		{Number: 39, Value: "دیرش نهه"},
		{Number: 38, Value: "دیرش اته"},
		{Number: 37, Value: "دیرش اووه"},
		{Number: 36, Value: "دیرش شپږ"},
		{Number: 35, Value: "دیرش پنځه"},
		{Number: 34, Value: "دیرش څلور"},
		{Number: 33, Value: "دیرش درې"},
		{Number: 32, Value: "دیرش دوه"},
		{Number: 31, Value: "دیرش یو"},
		{Number: 30, Value: "دیرش"},
		{Number: 29, Value: "شل نهه"},
		{Number: 28, Value: "شل اته"},
		{Number: 27, Value: "شل اووه"},
		{Number: 26, Value: "شل شپږ"},
		{Number: 25, Value: "شل پنځه"},
		{Number: 24, Value: "شل څلور"},
		{Number: 23, Value: "شل درې"},
		{Number: 22, Value: "شل دوه"},
		{Number: 21, Value: "شل یو"},
		{Number: 20, Value: "شل"},
		{Number: 19, Value: "نولس"},
		{Number: 18, Value: "اتلس"},
		{Number: 17, Value: "اولس"},
		{Number: 16, Value: "شپاړس"},
		{Number: 15, Value: "پنځلس"},
		{Number: 14, Value: "څوارلس"},
		{Number: 13, Value: "دیارلس"},
		{Number: 12, Value: "دولس"},
		{Number: 11, Value: "یولس"},
		{Number: 10, Value: "لس"},
		{Number: 9, Value: "نهه"},
		{Number: 8, Value: "اته"},
		{Number: 7, Value: "اووه"},
		{Number: 6, Value: "شپږ"},
		{Number: 5, Value: "پنځه"},
		{Number: 4, Value: "څلور"},
		{Number: 3, Value: "درې"},
		{Number: 2, Value: "دوه"},
		{Number: 1, Value: "یو"},
		{Number: 0, Value: "صفر"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "یو سل"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "لومړی", Suffix: "م", Masculine: "لومړی", Feminine: "لومړۍ", Neuter: "لومړی"},
		{Number: 2, Word: "دویم", Suffix: "م", Masculine: "دویم", Feminine: "دویمه", Neuter: "دویم"},
		{Number: 3, Word: "دریم", Suffix: "م", Masculine: "دریم", Feminine: "دریمه", Neuter: "دریم"},
		{Number: 4, Word: "څلورم", Suffix: "م", Masculine: "څلورم", Feminine: "څلورمه", Neuter: "څلورم"},
		{Number: 5, Word: "پنځم", Suffix: "م", Masculine: "پنځم", Feminine: "پنځمه", Neuter: "پنځم"},
		{Number: 6, Word: "شپږم", Suffix: "م", Masculine: "شپږم", Feminine: "شپږمه", Neuter: "شپږم"},
		{Number: 7, Word: "اووم", Suffix: "م", Masculine: "اووم", Feminine: "اوومه", Neuter: "اووم"},
		{Number: 8, Word: "اتم", Suffix: "م", Masculine: "اتم", Feminine: "اتمه", Neuter: "اتم"},
		{Number: 9, Word: "نهم", Suffix: "م", Masculine: "نهم", Feminine: "نهمه", Neuter: "نهم"},
		{Number: 10, Word: "لسم", Suffix: "م", Masculine: "لسم", Feminine: "لسمه", Neuter: "لسم"},
		{Number: 11, Word: "یولسم", Suffix: "م", Masculine: "یولسم", Feminine: "یولسمه", Neuter: "یولسم"},
		{Number: 12, Word: "دولسم", Suffix: "م", Masculine: "دولسم", Feminine: "دولسمه", Neuter: "دولسم"},
		{Number: 20, Word: "شلم", Suffix: "م", Masculine: "شلم", Feminine: "شلمه", Neuter: "شلم"},
		{Number: 21, Word: "یو ویشتم", Suffix: "م", Masculine: "یو ویشتم", Feminine: "یو ویشتمه", Neuter: "یو ویشتم"},
		{Number: 30, Word: "دیرشم", Suffix: "م", Masculine: "دیرشم", Feminine: "دیرشمه", Neuter: "دیرشم"},
		{Number: 50, Word: "پنځوسم", Suffix: "م", Masculine: "پنځوسم", Feminine: "پنځوسمه", Neuter: "پنځوسم"},
		{Number: 100, Word: "سلم", Suffix: "م", Masculine: "سلم", Feminine: "سلمه", Neuter: "سلم"},
		{Number: 1000, Word: "زرم", Suffix: "م", Masculine: "زرم", Feminine: "زرمه", Neuter: "زرم"},
	},
	LocaleFormatter: &PashtoFormatter{},
}

// PashtoFormatter handles Pashto formatting
type PashtoFormatter struct{}

func (f *PashtoFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *PashtoFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *PashtoFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *PashtoFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *PashtoFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *PashtoFormatter) ChopDecimal(value decimal.Decimal, precision int) decimal.Decimal {
	return value.Truncate(int32(precision))
}

func (f *PashtoFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *PashtoFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
