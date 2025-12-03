package locale

import (
	"github.com/shopspring/decimal"
)

// SILKLocale represents the Sinhala (Sri Lanka) locale
var SILKLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "ශ්\u200dරී ලංකා රුපියල",
		Plural:   "ශ්\u200dරී ලංකා රුපියල්",
		Singular: "ශ්\u200dරී ලංකා රුපියල",
		Symbol:   "Rs",
		FractionUnit: FractionUnit{
			Name:     "සෙන්ට්",
			Plural:   "සෙන්ට්",
			Singular: "සෙන්ට්",
			Symbol:   "c",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Sri Lanka",
		Currency:       "LKR",
		ISO3166Alpha2:  "LK",
		ISO3166Alpha3:  "LKA",
		ISO3166Numeric: "144",
		Locale:         "si-LK",
		Timezone:       []string{"Asia/Colombo"},
		Language:       "si",
		Emoji:          "🇱🇰",
		PhoneCode:      "+94",
		Domain:         ".lk",
	},
	Texts: Texts{
		And:   "සහ",
		Minus: "අඩු",
		Only:  "පමණක්",
		Point: "ලක්ෂ්\u200dයය",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "බිලියන ලක්ෂ"},
		{Number: 1000000000000, Value: "බිලියන"},
		{Number: 1000000000, Value: "බිලියන"},
		{Number: 1000000, Value: "මිලියන"},
		{Number: 100000, Value: "ලක්ෂය"},
		{Number: 90000, Value: "අනූව දහස"},
		{Number: 80000, Value: "අෂ්ට දහස"},
		{Number: 70000, Value: "සප්ත දහස"},
		{Number: 60000, Value: "ෂෂ්ට දහස"},
		{Number: 50000, Value: "පන්චාස් දහස"},
		{Number: 40000, Value: "චත්වාරිස් දහස"},
		{Number: 30000, Value: "තිරිස් දහස"},
		{Number: 20000, Value: "විස් දහස"},
		{Number: 19000, Value: "ඒකාදශ දහස"},
		{Number: 18000, Value: "අෂ්ටාදශ දහස"},
		{Number: 17000, Value: "සප්තදශ දහස"},
		{Number: 16000, Value: "ෂෝඩශ දහස"},
		{Number: 15000, Value: "පන්චදශ දහස"},
		{Number: 14000, Value: "චතුර්දශ දහස"},
		{Number: 13000, Value: "ත්\u200dරයෝදශ දහස"},
		{Number: 12000, Value: "ද්වාදශ දහස"},
		{Number: 11000, Value: "ඒකාදශ දහස"},
		{Number: 10000, Value: "දශ දහස"},
		{Number: 9000, Value: "නව දහස"},
		{Number: 8000, Value: "අෂ්ට දහස"},
		{Number: 7000, Value: "සප්ත දහස"},
		{Number: 6000, Value: "ෂට් දහස"},
		{Number: 5000, Value: "පන්ච දහස"},
		{Number: 4000, Value: "චතුර් දහස"},
		{Number: 3000, Value: "ත්\u200dරි දහස"},
		{Number: 2000, Value: "ද්වි දහස"},
		{Number: 1000, Value: "දහස"},
		{Number: 900, Value: "නව සියය"},
		{Number: 800, Value: "අෂ්ට සියය"},
		{Number: 700, Value: "සප්ත සියය"},
		{Number: 600, Value: "ෂට් සියය"},
		{Number: 500, Value: "පන්ච සියය"},
		{Number: 400, Value: "චතුර් සියය"},
		{Number: 300, Value: "ත්\u200dරි සියය"},
		{Number: 200, Value: "ද්වි සියය"},
		{Number: 100, Value: "සියය"},
		{Number: 99, Value: "අනූව සහ නව"},
		{Number: 98, Value: "අනූව සහ අෂ්ට"},
		{Number: 97, Value: "අනූව සහ සප්ත"},
		{Number: 96, Value: "අනූව සහ ෂට්"},
		{Number: 95, Value: "අනූව සහ පන්ච"},
		{Number: 94, Value: "අනූව සහ චතුර්"},
		{Number: 93, Value: "අනූව සහ ත්\u200dරි"},
		{Number: 92, Value: "අනූව සහ ද්වි"},
		{Number: 91, Value: "අනූව සහ ඒක"},
		{Number: 90, Value: "අනූව"},
		{Number: 89, Value: "අෂීති සහ නව"},
		{Number: 88, Value: "අෂීති සහ අෂ්ට"},
		{Number: 87, Value: "අෂීති සහ සප්ත"},
		{Number: 86, Value: "අෂීති සහ ෂට්"},
		{Number: 85, Value: "අෂීති සහ පන්ච"},
		{Number: 84, Value: "අෂීති සහ චතුර්"},
		{Number: 83, Value: "අෂීති සහ ත්\u200dරි"},
		{Number: 82, Value: "අෂීති සහ ද්වි"},
		{Number: 81, Value: "අෂීති සහ ඒක"},
		{Number: 80, Value: "අෂීති"},
		{Number: 79, Value: "සප්තතිය සහ නව"},
		{Number: 78, Value: "සප්තතිය සහ අෂ්ට"},
		{Number: 77, Value: "සප්තතිය සහ සප්ත"},
		{Number: 76, Value: "සප්තතිය සහ ෂට්"},
		{Number: 75, Value: "සප්තතිය සහ පන්ච"},
		{Number: 74, Value: "සප්තතිය සහ චතුර්"},
		{Number: 73, Value: "සප්තතිය සහ ත්\u200dරි"},
		{Number: 72, Value: "සප්තතිය සහ ද්වි"},
		{Number: 71, Value: "සප්තතිය සහ ඒක"},
		{Number: 70, Value: "සප්තතිය"},
		{Number: 69, Value: "ෂෂ්ටිය සහ නව"},
		{Number: 68, Value: "ෂෂ්ටිය සහ අෂ්ට"},
		{Number: 67, Value: "ෂෂ්ටිය සහ සප්ත"},
		{Number: 66, Value: "ෂෂ්ටිය සහ ෂට්"},
		{Number: 65, Value: "ෂෂ්ටිය සහ පන්ච"},
		{Number: 64, Value: "ෂෂ්ටිය සහ චතුර්"},
		{Number: 63, Value: "ෂෂ්ටිය සහ ත්\u200dරි"},
		{Number: 62, Value: "ෂෂ්ටිය සහ ද්වි"},
		{Number: 61, Value: "ෂෂ්ටිය සහ ඒක"},
		{Number: 60, Value: "ෂෂ්ටිය"},
		{Number: 59, Value: "පන්චාශ සහ නව"},
		{Number: 58, Value: "පන්චාශ සහ අෂ්ට"},
		{Number: 57, Value: "පන්චාශ සහ සප්ත"},
		{Number: 56, Value: "පන්චාශ සහ ෂට්"},
		{Number: 55, Value: "පන්චාශ සහ පන්ච"},
		{Number: 54, Value: "පන්චාශ සහ චතුර්"},
		{Number: 53, Value: "පන්චාශ සහ ත්\u200dරි"},
		{Number: 52, Value: "පන්චාශ සහ ද්වි"},
		{Number: 51, Value: "පන්චාශ සහ ඒක"},
		{Number: 50, Value: "පන්චාශ"},
		{Number: 49, Value: "චත්වාරිංශ සහ නව"},
		{Number: 48, Value: "චත්වාරිංශ සහ අෂ්ට"},
		{Number: 47, Value: "චත්වාරිංශ සහ සප්ත"},
		{Number: 46, Value: "චත්වාරිංශ සහ ෂට්"},
		{Number: 45, Value: "චත්වාරිංශ සහ පන්ච"},
		{Number: 44, Value: "චත්වාරිංශ සහ චතුර්"},
		{Number: 43, Value: "චත්වාරිංශ සහ ත්\u200dරි"},
		{Number: 42, Value: "චත්වාරිංශ සහ ද්වි"},
		{Number: 41, Value: "චත්වාරිංශ සහ ඒක"},
		{Number: 40, Value: "චත්වාරිංශ"},
		{Number: 39, Value: "ත්\u200dරිංශ සහ නව"},
		{Number: 38, Value: "ත්\u200dරිංශ සහ අෂ්ට"},
		{Number: 37, Value: "ත්\u200dරිංශ සහ සප්ත"},
		{Number: 36, Value: "ත්\u200dරිංශ සහ ෂට්"},
		{Number: 35, Value: "ත්\u200dරිංශ සහ පන්ච"},
		{Number: 34, Value: "ත්\u200dරිංශ සහ චතුර්"},
		{Number: 33, Value: "ත්\u200dරිංශ සහ ත්\u200dරි"},
		{Number: 32, Value: "ත්\u200dරිංශ සහ ද්වි"},
		{Number: 31, Value: "ත්\u200dරිංශ සහ ඒක"},
		{Number: 30, Value: "ත්\u200dරිංශ"},
		{Number: 29, Value: "විංශ සහ නව"},
		{Number: 28, Value: "විංශ සහ අෂ්ට"},
		{Number: 27, Value: "විංශ සහ සප්ත"},
		{Number: 26, Value: "විංශ සහ ෂට්"},
		{Number: 25, Value: "විංශ සහ පන්ච"},
		{Number: 24, Value: "විංශ සහ චතුර්"},
		{Number: 23, Value: "විංශ සහ ත්\u200dරි"},
		{Number: 22, Value: "විංශ සහ ද්වි"},
		{Number: 21, Value: "විංශ සහ ඒක"},
		{Number: 20, Value: "විංශ"},
		{Number: 19, Value: "ඒකාදශ"},
		{Number: 18, Value: "අෂ්ටාදශ"},
		{Number: 17, Value: "සප්තදශ"},
		{Number: 16, Value: "ෂෝඩශ"},
		{Number: 15, Value: "පන්චදශ"},
		{Number: 14, Value: "චතුර්දශ"},
		{Number: 13, Value: "ත්\u200dරයෝදශ"},
		{Number: 12, Value: "ද්වාදශ"},
		{Number: 11, Value: "ඒකාදශ"},
		{Number: 10, Value: "දශ"},
		{Number: 9, Value: "නව"},
		{Number: 8, Value: "අෂ්ට"},
		{Number: 7, Value: "සප්ත"},
		{Number: 6, Value: "ෂට්"},
		{Number: 5, Value: "පන්ච"},
		{Number: 4, Value: "චතුර්"},
		{Number: 3, Value: "ත්\u200dරි"},
		{Number: 2, Value: "ද්වි"},
		{Number: 1, Value: "ඒක"},
		{Number: 0, Value: "ශුන්\u200dය"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "සියය"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "පළමු", Suffix: "වැනි", Masculine: "පළමු", Feminine: "පළමු", Neuter: "පළමු"},
		{Number: 2, Word: "දෙවැනි", Suffix: "වැනි", Masculine: "දෙවැනි", Feminine: "දෙවැනි", Neuter: "දෙවැනි"},
		{Number: 3, Word: "තෙවැනි", Suffix: "වැනි", Masculine: "තෙවැනි", Feminine: "තෙවැනි", Neuter: "තෙවැනි"},
		{Number: 4, Word: "සිව්වැනි", Suffix: "වැනි", Masculine: "සිව්වැනි", Feminine: "සිව්වැනි", Neuter: "සිව්වැනි"},
		{Number: 5, Word: "පස්වැනි", Suffix: "වැනි", Masculine: "පස්වැනි", Feminine: "පස්වැනි", Neuter: "පස්වැනි"},
		{Number: 6, Word: "හයවැනි", Suffix: "වැනි", Masculine: "හයවැනි", Feminine: "හයවැනි", Neuter: "හයවැනි"},
		{Number: 7, Word: "හත්වැනි", Suffix: "වැනි", Masculine: "හත්වැනි", Feminine: "හත්වැනි", Neuter: "හත්වැනි"},
		{Number: 8, Word: "අටවැනි", Suffix: "වැනි", Masculine: "අටවැනි", Feminine: "අටවැනි", Neuter: "අටවැනි"},
		{Number: 9, Word: "නවවැනි", Suffix: "වැනි", Masculine: "නවවැනි", Feminine: "නවවැනි", Neuter: "නවවැනි"},
		{Number: 10, Word: "දසවැනි", Suffix: "වැනි", Masculine: "දසවැනි", Feminine: "දසවැනි", Neuter: "දසවැනි"},
		{Number: 11, Word: "එකොළොස්වැනි", Suffix: "වැනි", Masculine: "එකොළොස්වැනි", Feminine: "එකොළොස්වැනි", Neuter: "එකොළොස්වැනි"},
		{Number: 12, Word: "දොළොස්වැනි", Suffix: "වැනි", Masculine: "දොළොස්වැනි", Feminine: "දොළොස්වැනි", Neuter: "දොළොස්වැනි"},
		{Number: 20, Word: "විසිවැනි", Suffix: "වැනි", Masculine: "විසිවැනි", Feminine: "විසිවැනි", Neuter: "විසිවැනි"},
		{Number: 21, Word: "විස්සේ පළමුවැනි", Suffix: "වැනි", Masculine: "විස්සේ පළමුවැනි", Feminine: "විස්සේ පළමුවැනි", Neuter: "විස්සේ පළමුවැනි"},
		{Number: 30, Word: "තිහිවැනි", Suffix: "වැනි", Masculine: "තිහිවැනි", Feminine: "තිහිවැනි", Neuter: "තිහිවැනි"},
		{Number: 50, Word: "පනහවැනි", Suffix: "වැනි", Masculine: "පනහවැනි", Feminine: "පනහවැනි", Neuter: "පනහවැනි"},
		{Number: 100, Word: "සියවැනි", Suffix: "වැනි", Masculine: "සියවැනි", Feminine: "සියවැනි", Neuter: "සියවැනි"},
		{Number: 1000, Word: "දහස්වැනි", Suffix: "වැනි", Masculine: "දහස්වැනි", Feminine: "දහස්වැනි", Neuter: "දහස්වැනි"},
	},
	LocaleFormatter: &SinhalaSriLankaFormatter{},
}

// SinhalaSriLankaFormatter handles Sinhala (Sri Lanka) formatting
type SinhalaSriLankaFormatter struct{}

func (f *SinhalaSriLankaFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *SinhalaSriLankaFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *SinhalaSriLankaFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *SinhalaSriLankaFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *SinhalaSriLankaFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *SinhalaSriLankaFormatter) ChopDecimal(d decimal.Decimal, precision int) decimal.Decimal {
	return d.Truncate(int32(precision))
}

func (f *SinhalaSriLankaFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *SinhalaSriLankaFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
