package locale

import "github.com/shopspring/decimal"

// SDPKLocale represents the Sindhi (Pakistan) locale
var SDPKLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "پاڪستاني رپيو",
		Plural:   "پاڪستاني رپيا",
		Singular: "پاڪستاني رپيو",
		Symbol:   "Rs",
		FractionUnit: FractionUnit{
			Name:     "پيسو",
			Plural:   "پيسا",
			Singular: "پيسو",
			Symbol:   "p",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Pakistan",
		Currency:       "PKR",
		ISO3166Alpha2:  "PK",
		ISO3166Alpha3:  "PAK",
		ISO3166Numeric: "586",
		Locale:         "sd-PK",
		Timezone:       []string{"Asia/Karachi"},
		Language:       "sd",
	},
	Texts: Texts{
		And:   "۽",
		Minus: "گهٽ",
		Only:  "فقط",
		Point: "نقطو",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "هڪ کرب"},
		{Number: 1000000000000, Value: "هڪ ارب"},
		{Number: 1000000000, Value: "هڪ کرور"},
		{Number: 1000000, Value: "هڪ لک"},
		{Number: 100000, Value: "لک"},
		{Number: 90000, Value: "نوي هزار"},
		{Number: 80000, Value: "اٺ هزار"},
		{Number: 70000, Value: "ست هزار"},
		{Number: 60000, Value: "ڇهه هزار"},
		{Number: 50000, Value: "پنجاهه هزار"},
		{Number: 40000, Value: "چاليهه هزار"},
		{Number: 30000, Value: "ٽيهه هزار"},
		{Number: 20000, Value: "ويهه هزار"},
		{Number: 19000, Value: "اڻويهه هزار"},
		{Number: 18000, Value: "اٺاريهه هزار"},
		{Number: 17000, Value: "ستاريهه هزار"},
		{Number: 16000, Value: "سوليهه هزار"},
		{Number: 15000, Value: "پندريهه هزار"},
		{Number: 14000, Value: "چوڏهه هزار"},
		{Number: 13000, Value: "تيريهه هزار"},
		{Number: 12000, Value: "ٻارهه هزار"},
		{Number: 11000, Value: "يارهن هزار"},
		{Number: 10000, Value: "ڏهه هزار"},
		{Number: 9000, Value: "نو هزار"},
		{Number: 8000, Value: "اٺ هزار"},
		{Number: 7000, Value: "ست هزار"},
		{Number: 6000, Value: "ڇهه هزار"},
		{Number: 5000, Value: "پنج هزار"},
		{Number: 4000, Value: "چار هزار"},
		{Number: 3000, Value: "ٽي هزار"},
		{Number: 2000, Value: "ٻه هزار"},
		{Number: 1000, Value: "هزار"},
		{Number: 900, Value: "نو سئو"},
		{Number: 800, Value: "اٺ سئو"},
		{Number: 700, Value: "ست سئو"},
		{Number: 600, Value: "ڇهه سئو"},
		{Number: 500, Value: "پنج سئو"},
		{Number: 400, Value: "چار سئو"},
		{Number: 300, Value: "ٽي سئو"},
		{Number: 200, Value: "ٻه سئو"},
		{Number: 100, Value: "سئو"},
		{Number: 99, Value: "نوي ۽ نو"},
		{Number: 98, Value: "نوي ۽ اٺ"},
		{Number: 97, Value: "نوي ۽ ست"},
		{Number: 96, Value: "نوي ۽ ڇهه"},
		{Number: 95, Value: "نوي ۽ پنج"},
		{Number: 94, Value: "نوي ۽ چار"},
		{Number: 93, Value: "نوي ۽ ٽي"},
		{Number: 92, Value: "نوي ۽ ٻه"},
		{Number: 91, Value: "نوي ۽ هڪ"},
		{Number: 90, Value: "نوي"},
		{Number: 89, Value: "اٺي ۽ نو"},
		{Number: 88, Value: "اٺي ۽ اٺ"},
		{Number: 87, Value: "اٺي ۽ ست"},
		{Number: 86, Value: "اٺي ۽ ڇهه"},
		{Number: 85, Value: "اٺي ۽ پنج"},
		{Number: 84, Value: "اٺي ۽ چار"},
		{Number: 83, Value: "اٺي ۽ ٽي"},
		{Number: 82, Value: "اٺي ۽ ٻه"},
		{Number: 81, Value: "اٺي ۽ هڪ"},
		{Number: 80, Value: "اٺي"},
		{Number: 79, Value: "ستر ۽ نو"},
		{Number: 78, Value: "ستر ۽ اٺ"},
		{Number: 77, Value: "ستر ۽ ست"},
		{Number: 76, Value: "ستر ۽ ڇهه"},
		{Number: 75, Value: "ستر ۽ پنج"},
		{Number: 74, Value: "ستر ۽ چار"},
		{Number: 73, Value: "ستر ۽ ٽي"},
		{Number: 72, Value: "ستر ۽ ٻه"},
		{Number: 71, Value: "ستر ۽ هڪ"},
		{Number: 70, Value: "ستر"},
		{Number: 69, Value: "سٺ ۽ نو"},
		{Number: 68, Value: "سٺ ۽ اٺ"},
		{Number: 67, Value: "سٺ ۽ ست"},
		{Number: 66, Value: "سٺ ۽ ڇهه"},
		{Number: 65, Value: "سٺ ۽ پنج"},
		{Number: 64, Value: "سٺ ۽ چار"},
		{Number: 63, Value: "سٺ ۽ ٽي"},
		{Number: 62, Value: "سٺ ۽ ٻه"},
		{Number: 61, Value: "سٺ ۽ هڪ"},
		{Number: 60, Value: "سٺ"},
		{Number: 59, Value: "پنجاهه ۽ نو"},
		{Number: 58, Value: "پنجاهه ۽ اٺ"},
		{Number: 57, Value: "پنجاهه ۽ ست"},
		{Number: 56, Value: "پنجاهه ۽ ڇهه"},
		{Number: 55, Value: "پنجاهه ۽ پنج"},
		{Number: 54, Value: "پنجاهه ۽ چار"},
		{Number: 53, Value: "پنجاهه ۽ ٽي"},
		{Number: 52, Value: "پنجاهه ۽ ٻه"},
		{Number: 51, Value: "پنجاهه ۽ هڪ"},
		{Number: 50, Value: "پنجاهه"},
		{Number: 49, Value: "چاليهه ۽ نو"},
		{Number: 48, Value: "چاليهه ۽ اٺ"},
		{Number: 47, Value: "چاليهه ۽ ست"},
		{Number: 46, Value: "چاليهه ۽ ڇهه"},
		{Number: 45, Value: "چاليهه ۽ پنج"},
		{Number: 44, Value: "چاليهه ۽ چار"},
		{Number: 43, Value: "چاليهه ۽ ٽي"},
		{Number: 42, Value: "چاليهه ۽ ٻه"},
		{Number: 41, Value: "چاليهه ۽ هڪ"},
		{Number: 40, Value: "چاليهه"},
		{Number: 39, Value: "ٽيهه ۽ نو"},
		{Number: 38, Value: "ٽيهه ۽ اٺ"},
		{Number: 37, Value: "ٽيهه ۽ ست"},
		{Number: 36, Value: "ٽيهه ۽ ڇهه"},
		{Number: 35, Value: "ٽيهه ۽ پنج"},
		{Number: 34, Value: "ٽيهه ۽ چار"},
		{Number: 33, Value: "ٽيهه ۽ ٽي"},
		{Number: 32, Value: "ٽيهه ۽ ٻه"},
		{Number: 31, Value: "ٽيهه ۽ هڪ"},
		{Number: 30, Value: "ٽيهه"},
		{Number: 29, Value: "ويهه ۽ نو"},
		{Number: 28, Value: "ويهه ۽ اٺ"},
		{Number: 27, Value: "ويهه ۽ ست"},
		{Number: 26, Value: "ويهه ۽ ڇهه"},
		{Number: 25, Value: "ويهه ۽ پنج"},
		{Number: 24, Value: "ويهه ۽ چار"},
		{Number: 23, Value: "ويهه ۽ ٽي"},
		{Number: 22, Value: "ويهه ۽ ٻه"},
		{Number: 21, Value: "ويهه ۽ هڪ"},
		{Number: 20, Value: "ويهه"},
		{Number: 19, Value: "اڻويهه"},
		{Number: 18, Value: "اٺاريهه"},
		{Number: 17, Value: "ستاريهه"},
		{Number: 16, Value: "سوليهه"},
		{Number: 15, Value: "پندريهه"},
		{Number: 14, Value: "چوڏهه"},
		{Number: 13, Value: "تيريهه"},
		{Number: 12, Value: "ٻارهه"},
		{Number: 11, Value: "يارهن"},
		{Number: 10, Value: "ڏهه"},
		{Number: 9, Value: "نو"},
		{Number: 8, Value: "اٺ"},
		{Number: 7, Value: "ست"},
		{Number: 6, Value: "ڇهه"},
		{Number: 5, Value: "پنج"},
		{Number: 4, Value: "چار"},
		{Number: 3, Value: "ٽي"},
		{Number: 2, Value: "ٻه"},
		{Number: 1, Value: "هڪ"},
		{Number: 0, Value: "صفر"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "سئو"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "پهريون", Suffix: "ون", Masculine: "پهريون", Feminine: "پهرياڻ", Neuter: "پهريون"},
		{Number: 2, Word: "ٻيون", Suffix: "ون", Masculine: "ٻيون", Feminine: "ٻياڻ", Neuter: "ٻيون"},
		{Number: 3, Word: "ٽيون", Suffix: "ون", Masculine: "ٽيون", Feminine: "ٽياڻ", Neuter: "ٽيون"},
		{Number: 4, Word: "چوٿون", Suffix: "ون", Masculine: "چوٿون", Feminine: "چوٿاڻ", Neuter: "چوٿون"},
		{Number: 5, Word: "پنجون", Suffix: "ون", Masculine: "پنجون", Feminine: "پنجاڻ", Neuter: "پنجون"},
		{Number: 6, Word: "ڇهون", Suffix: "ون", Masculine: "ڇهون", Feminine: "ڇهاڻ", Neuter: "ڇهون"},
		{Number: 7, Word: "ستون", Suffix: "ون", Masculine: "ستون", Feminine: "ستاڻ", Neuter: "ستون"},
		{Number: 8, Word: "اٺون", Suffix: "ون", Masculine: "اٺون", Feminine: "اٺاڻ", Neuter: "اٺون"},
		{Number: 9, Word: "نون", Suffix: "ون", Masculine: "نون", Feminine: "نواڻ", Neuter: "نون"},
		{Number: 10, Word: "ڏهون", Suffix: "ون", Masculine: "ڏهون", Feminine: "ڏهاڻ", Neuter: "ڏهون"},
		{Number: 11, Word: "يارهنون", Suffix: "ون", Masculine: "يارهنون", Feminine: "يارهناڻ", Neuter: "يارهنون"},
		{Number: 12, Word: "ٻارهنون", Suffix: "ون", Masculine: "ٻارهنون", Feminine: "ٻارهناڻ", Neuter: "ٻارهنون"},
		{Number: 20, Word: "ويهون", Suffix: "ون", Masculine: "ويهون", Feminine: "ويهاڣ", Neuter: "ويهون"},
		{Number: 21, Word: "ايڪويهون", Suffix: "ون", Masculine: "ايڪويهون", Feminine: "ايڪويهاڣ", Neuter: "ايڪويهون"},
		{Number: 30, Word: "ٽيهون", Suffix: "ون", Masculine: "ٽيهون", Feminine: "ٽيهاڣ", Neuter: "ٽيهون"},
		{Number: 50, Word: "پنجاهون", Suffix: "ون", Masculine: "پنجاهون", Feminine: "پنجاهاڣ", Neuter: "پنجاهون"},
		{Number: 100, Word: "سئون", Suffix: "ون", Masculine: "سئون", Feminine: "سئاڣ", Neuter: "سئون"},
		{Number: 1000, Word: "هزارون", Suffix: "ون", Masculine: "هزارون", Feminine: "هزاراڣ", Neuter: "هزارون"},
	},
	LocaleFormatter: &SindhiPakistanFormatter{},
}

// SindhiPakistanFormatter handles Sindhi (Pakistan) formatting
type SindhiPakistanFormatter struct{}

func (f *SindhiPakistanFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *SindhiPakistanFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *SindhiPakistanFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *SindhiPakistanFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *SindhiPakistanFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *SindhiPakistanFormatter) ChopDecimal(d decimal.Decimal, precision int) decimal.Decimal {
	return d.Truncate(int32(precision))
}
