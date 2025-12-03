package locale

import (
	"github.com/shopspring/decimal"
)

// MYMMLocale represents the Myanmar (Myanmar) locale
var MYMMLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "ကျပ်",
		Plural:   "ကျပ်",
		Singular: "ကျပ်",
		Symbol:   "K",
		FractionUnit: FractionUnit{
			Name:     "ပြား",
			Plural:   "ပြား",
			Singular: "ပြား",
			Symbol:   "p",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Myanmar",
		Currency:       "MMK",
		ISO3166Alpha2:  "MM",
		ISO3166Alpha3:  "MMR",
		ISO3166Numeric: "104",
		Locale:         "my-MM",
		Timezone:       []string{"Asia/Yangon"},
		Language:       "my",
	},
	Texts: Texts{
		And:   "နှင့်",
		Minus: "နုတ်",
		Only:  "သာ",
		Point: "အမှတ်",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "ကုဋိကုဋိ"},
		{Number: 1000000000000, Value: "တထောင်ကုဋိ"},
		{Number: 1000000000, Value: "တဆယ်ကုဋိ"},
		{Number: 100000000, Value: "ကုဋိ"},
		{Number: 10000000, Value: "တသောင်းသိန်း"},
		{Number: 1000000, Value: "သိန်း"},
		{Number: 100000, Value: "သိန်း"},
		{Number: 90000, Value: "ကိုးသောင်း"},
		{Number: 80000, Value: "ရှစ်သောင်း"},
		{Number: 70000, Value: "ခုနစ်သောင်း"},
		{Number: 60000, Value: "ခြောက်သောင်း"},
		{Number: 50000, Value: "ငါးသောင်း"},
		{Number: 40000, Value: "လေးသောင်း"},
		{Number: 30000, Value: "သုံးသောင်း"},
		{Number: 20000, Value: "နှစ်သောင်း"},
		{Number: 19000, Value: "ဆယ့်ကိုးထောင်"},
		{Number: 18000, Value: "ဆယ့်ရှစ်ထောင်"},
		{Number: 17000, Value: "ဆယ့်ခုနစ်ထောင်"},
		{Number: 16000, Value: "ဆယ့်ခြောက်ထောင်"},
		{Number: 15000, Value: "ဆယ့်ငါးထောင်"},
		{Number: 14000, Value: "ဆယ့်လေးထောင်"},
		{Number: 13000, Value: "ဆယ့်သုံးထောင်"},
		{Number: 12000, Value: "ဆယ့်နှစ်ထောင်"},
		{Number: 11000, Value: "ဆယ့်တထောင်"},
		{Number: 10000, Value: "တသောင်း"},
		{Number: 9000, Value: "ကိုးထောင်"},
		{Number: 8000, Value: "ရှစ်ထောင်"},
		{Number: 7000, Value: "ခုနစ်ထောင်"},
		{Number: 6000, Value: "ခြောက်ထောင်"},
		{Number: 5000, Value: "ငါးထောင်"},
		{Number: 4000, Value: "လေးထောင်"},
		{Number: 3000, Value: "သုံးထောင်"},
		{Number: 2000, Value: "နှစ်ထောင်"},
		{Number: 1000, Value: "တထောင်"},
		{Number: 900, Value: "ကိုးရာ"},
		{Number: 800, Value: "ရှစ်ရာ"},
		{Number: 700, Value: "ခုနစ်ရာ"},
		{Number: 600, Value: "ခြောက်ရာ"},
		{Number: 500, Value: "ငါးရာ"},
		{Number: 400, Value: "လေးရာ"},
		{Number: 300, Value: "သုံးရာ"},
		{Number: 200, Value: "နှစ်ရာ"},
		{Number: 100, Value: "တရာ"},
		{Number: 99, Value: "ကိုးဆယ့်ကိုး"},
		{Number: 98, Value: "ကိုးဆယ့်ရှစ်"},
		{Number: 97, Value: "ကိုးဆယ့်ခုနစ်"},
		{Number: 96, Value: "ကိုးဆယ့်ခြောက်"},
		{Number: 95, Value: "ကိုးဆယ့်ငါး"},
		{Number: 94, Value: "ကိုးဆယ့်လေး"},
		{Number: 93, Value: "ကိုးဆယ့်သုံး"},
		{Number: 92, Value: "ကိုးဆယ့်နှစ်"},
		{Number: 91, Value: "ကိုးဆယ့်တ"},
		{Number: 90, Value: "ကိုးဆယ်"},
		{Number: 89, Value: "ရှစ်ဆယ့်ကိုး"},
		{Number: 88, Value: "ရှစ်ဆယ့်ရှစ်"},
		{Number: 87, Value: "ရှစ်ဆယ့်ခုနစ်"},
		{Number: 86, Value: "ရှစ်ဆယ့်ခြောက်"},
		{Number: 85, Value: "ရှစ်ဆယ့်ငါး"},
		{Number: 84, Value: "ရှစ်ဆယ့်လေး"},
		{Number: 83, Value: "ရှစ်ဆယ့်သုံး"},
		{Number: 82, Value: "ရှစ်ဆယ့်နှစ်"},
		{Number: 81, Value: "ရှစ်ဆယ့်တ"},
		{Number: 80, Value: "ရှစ်ဆယ်"},
		{Number: 79, Value: "ခုနစ်ဆယ့်ကိုး"},
		{Number: 78, Value: "ခုနစ်ဆယ့်ရှစ်"},
		{Number: 77, Value: "ခုနစ်ဆယ့်ခုနစ်"},
		{Number: 76, Value: "ခုနစ်ဆယ့်ခြောက်"},
		{Number: 75, Value: "ခုနစ်ဆယ့်ငါး"},
		{Number: 74, Value: "ခုနစ်ဆယ့်လေး"},
		{Number: 73, Value: "ခုနစ်ဆယ့်သုံး"},
		{Number: 72, Value: "ခုနစ်ဆယ့်နှစ်"},
		{Number: 71, Value: "ခုနစ်ဆယ့်တ"},
		{Number: 70, Value: "ခုနစ်ဆယ်"},
		{Number: 69, Value: "ခြောက်ဆယ့်ကိုး"},
		{Number: 68, Value: "ခြောက်ဆယ့်ရှစ်"},
		{Number: 67, Value: "ခြောက်ဆယ့်ခုနစ်"},
		{Number: 66, Value: "ခြောက်ဆယ့်ခြောက်"},
		{Number: 65, Value: "ခြောက်ဆယ့်ငါး"},
		{Number: 64, Value: "ခြောက်ဆယ့်လေး"},
		{Number: 63, Value: "ခြောက်ဆယ့်သုံး"},
		{Number: 62, Value: "ခြောက်ဆယ့်နှစ်"},
		{Number: 61, Value: "ခြောက်ဆယ့်တ"},
		{Number: 60, Value: "ခြောက်ဆယ်"},
		{Number: 59, Value: "ငါးဆယ့်ကိုး"},
		{Number: 58, Value: "ငါးဆယ့်ရှစ်"},
		{Number: 57, Value: "ငါးဆယ့်ခုနစ်"},
		{Number: 56, Value: "ငါးဆယ့်ခြောက်"},
		{Number: 55, Value: "ငါးဆယ့်ငါး"},
		{Number: 54, Value: "ငါးဆယ့်လေး"},
		{Number: 53, Value: "ငါးဆယ့်သုံး"},
		{Number: 52, Value: "ငါးဆယ့်နှစ်"},
		{Number: 51, Value: "ငါးဆယ့်တ"},
		{Number: 50, Value: "ငါးဆယ်"},
		{Number: 49, Value: "လေးဆယ့်ကိုး"},
		{Number: 48, Value: "လေးဆယ့်ရှစ်"},
		{Number: 47, Value: "လေးဆယ့်ခုနစ်"},
		{Number: 46, Value: "လေးဆယ့်ခြောက်"},
		{Number: 45, Value: "လေးဆယ့်ငါး"},
		{Number: 44, Value: "လေးဆယ့်လေး"},
		{Number: 43, Value: "လေးဆယ့်သုံး"},
		{Number: 42, Value: "လေးဆယ့်နှစ်"},
		{Number: 41, Value: "လေးဆယ့်တ"},
		{Number: 40, Value: "လေးဆယ်"},
		{Number: 39, Value: "သုံးဆယ့်ကိုး"},
		{Number: 38, Value: "သုံးဆယ့်ရှစ်"},
		{Number: 37, Value: "သုံးဆယ့်ခုနစ်"},
		{Number: 36, Value: "သုံးဆယ့်ခြောက်"},
		{Number: 35, Value: "သုံးဆယ့်ငါး"},
		{Number: 34, Value: "သုံးဆယ့်လေး"},
		{Number: 33, Value: "သုံးဆယ့်သုံး"},
		{Number: 32, Value: "သုံးဆယ့်နှစ်"},
		{Number: 31, Value: "သုံးဆယ့်တ"},
		{Number: 30, Value: "သုံးဆယ်"},
		{Number: 29, Value: "နှစ်ဆယ့်ကိုး"},
		{Number: 28, Value: "နှစ်ဆယ့်ရှစ်"},
		{Number: 27, Value: "နှစ်ဆယ့်ခုနစ်"},
		{Number: 26, Value: "နှစ်ဆယ့်ခြောက်"},
		{Number: 25, Value: "နှစ်ဆယ့်ငါး"},
		{Number: 24, Value: "နှစ်ဆယ့်လေး"},
		{Number: 23, Value: "နှစ်ဆယ့်သုံး"},
		{Number: 22, Value: "နှစ်ဆယ့်နှစ်"},
		{Number: 21, Value: "နှစ်ဆယ့်တ"},
		{Number: 20, Value: "နှစ်ဆယ်"},
		{Number: 19, Value: "ဆယ့်ကိုး"},
		{Number: 18, Value: "ဆယ့်ရှစ်"},
		{Number: 17, Value: "ဆယ့်ခုနစ်"},
		{Number: 16, Value: "ဆယ့်ခြောက်"},
		{Number: 15, Value: "ဆယ့်ငါး"},
		{Number: 14, Value: "ဆယ့်လေး"},
		{Number: 13, Value: "ဆယ့်သုံး"},
		{Number: 12, Value: "ဆယ့်နှစ်"},
		{Number: 11, Value: "ဆယ့်တ"},
		{Number: 10, Value: "ဆယ်"},
		{Number: 9, Value: "ကိုး"},
		{Number: 8, Value: "ရှစ်"},
		{Number: 7, Value: "ခုနစ်"},
		{Number: 6, Value: "ခြောက်"},
		{Number: 5, Value: "ငါး"},
		{Number: 4, Value: "လေး"},
		{Number: 3, Value: "သုံး"},
		{Number: 2, Value: "နှစ်"},
		{Number: 1, Value: "တ"},
		{Number: 0, Value: "သုည"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "တရာ"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "ပထမ", Suffix: "ပထမ", Masculine: "ပထမ", Feminine: "ပထမ", Neuter: "ပထမ"},
		{Number: 2, Word: "ဒုတိယ", Suffix: "ဒုတိယ", Masculine: "ဒုတိယ", Feminine: "ဒုတိယ", Neuter: "ဒုတိယ"},
		{Number: 3, Word: "တတိယ", Suffix: "တတိယ", Masculine: "တတိယ", Feminine: "တတိယ", Neuter: "တတိယ"},
		{Number: 4, Word: "စတုတ္ထ", Suffix: "စတုတ္ထ", Masculine: "စတုတ္ထ", Feminine: "စတုတ္ထ", Neuter: "စတုတ္ထ"},
		{Number: 5, Word: "ပဉ္စမ", Suffix: "ပဉ္စမ", Masculine: "ပဉ္စမ", Feminine: "ပဉ္စမ", Neuter: "ပဉ္စမ"},
		{Number: 6, Word: "ဆဋ္ဌမ", Suffix: "ဆဋ္ဌမ", Masculine: "ဆဋ္ဌမ", Feminine: "ဆဋ္ဌမ", Neuter: "ဆဋ္ဌမ"},
		{Number: 7, Word: "သတ္တမ", Suffix: "သတ္တမ", Masculine: "သတ္တမ", Feminine: "သတ္တမ", Neuter: "သတ္တမ"},
		{Number: 8, Word: "အဋ္ဌမ", Suffix: "အဋ္ဌမ", Masculine: "အဋ္ဌမ", Feminine: "အဋ္ဌမ", Neuter: "အဋ္ဌမ"},
		{Number: 9, Word: "နဝမ", Suffix: "နဝမ", Masculine: "နဝမ", Feminine: "နဝမ", Neuter: "နဝမ"},
		{Number: 10, Word: "ဒသမ", Suffix: "ဒသမ", Masculine: "ဒသမ", Feminine: "ဒသမ", Neuter: "ဒသမ"},
		{Number: 11, Word: "ဧကာဒသမ", Suffix: "ဧကာဒသမ", Masculine: "ဧကာဒသမ", Feminine: "ဧကာဒသမ", Neuter: "ဧကာဒသမ"},
		{Number: 12, Word: "ဒွါဒသမ", Suffix: "ဒွါဒသမ", Masculine: "ဒွါဒသမ", Feminine: "ဒွါဒသမ", Neuter: "ဒွါဒသမ"},
		{Number: 20, Word: "ဝိံသတိမ", Suffix: "ဝိံသတိမ", Masculine: "ဝိံသတိမ", Feminine: "ဝိံသတိမ", Neuter: "ဝိံသတိမ"},
		{Number: 21, Word: "ဧကာဝိံသတိမ", Suffix: "ဧကာဝိံသတိမ", Masculine: "ဧကာဝိံသတိမ", Feminine: "ဧကာဝိံသတိမ", Neuter: "ဧကာဝိံသတိမ"},
		{Number: 30, Word: "တိံသတိမ", Suffix: "တိံသတိမ", Masculine: "တိံသတိမ", Feminine: "တိံသတိမ", Neuter: "တိံသတိမ"},
		{Number: 100, Word: "သတမ", Suffix: "သတမ", Masculine: "သတမ", Feminine: "သတမ", Neuter: "သတမ"},
		{Number: 1000, Word: "သဟဿမ", Suffix: "သဟဿမ", Masculine: "သဟဿမ", Feminine: "သဟဿမ", Neuter: "သဟဿမ"},
	},
	LocaleFormatter: &MyanmarFormatter{},
}

// MyanmarFormatter handles Myanmar-specific formatting
type MyanmarFormatter struct{}

func (f *MyanmarFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *MyanmarFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *MyanmarFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *MyanmarFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *MyanmarFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *MyanmarFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 0
	}
	return amount.Truncate(int32(precision))
}

func (f *MyanmarFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *MyanmarFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
