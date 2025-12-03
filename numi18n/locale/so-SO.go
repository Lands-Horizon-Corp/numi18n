package locale

import (
	"github.com/shopspring/decimal"
)

// SOSOLocale represents the Somali (Somalia) locale
var SOSOLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Somali Shilling",
		Plural:   "Shilin Soomaaliyeed",
		Singular: "Shilin Soomaaliyeed",
		Symbol:   "SOS",
		FractionUnit: FractionUnit{
			Name:     "Cent",
			Plural:   "Cent",
			Singular: "Cent",
			Symbol:   "c",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Somalia",
		Currency:       "SOS",
		ISO3166Alpha2:  "SO",
		ISO3166Alpha3:  "SOM",
		ISO3166Numeric: "706",
		Locale:         "so-SO",
		Timezone:       []string{"Africa/Mogadishu"},
		Language:       "so",
		Emoji:          "🇸🇴",
	},
	Texts: Texts{
		And:   "iyo",
		Minus: "naqsi",
		Only:  "kaliya",
		Point: "dhibic",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "kun bilyan"},
		{Number: 1000000000000, Value: "bilyan"},
		{Number: 1000000000, Value: "bilyan"},
		{Number: 1000000, Value: "milyan"},
		{Number: 100000, Value: "boqol kun"},
		{Number: 90000, Value: "sagaashan kun"},
		{Number: 80000, Value: "siddeetan kun"},
		{Number: 70000, Value: "todobaatan kun"},
		{Number: 60000, Value: "lixdatan kun"},
		{Number: 50000, Value: "kontan kun"},
		{Number: 40000, Value: "afaratan kun"},
		{Number: 30000, Value: "soddonka kun"},
		{Number: 20000, Value: "labaatan kun"},
		{Number: 19000, Value: "sagaal iyo toban kun"},
		{Number: 18000, Value: "siddeed iyo toban kun"},
		{Number: 17000, Value: "toddoba iyo toban kun"},
		{Number: 16000, Value: "lix iyo toban kun"},
		{Number: 15000, Value: "shan iyo toban kun"},
		{Number: 14000, Value: "afar iyo toban kun"},
		{Number: 13000, Value: "saddex iyo toban kun"},
		{Number: 12000, Value: "laba iyo toban kun"},
		{Number: 11000, Value: "kow iyo toban kun"},
		{Number: 10000, Value: "toban kun"},
		{Number: 9000, Value: "sagaal kun"},
		{Number: 8000, Value: "siddeed kun"},
		{Number: 7000, Value: "toddoba kun"},
		{Number: 6000, Value: "lix kun"},
		{Number: 5000, Value: "shan kun"},
		{Number: 4000, Value: "afar kun"},
		{Number: 3000, Value: "saddex kun"},
		{Number: 2000, Value: "laba kun"},
		{Number: 1000, Value: "kun"},
		{Number: 900, Value: "sagaal boqol"},
		{Number: 800, Value: "siddeed boqol"},
		{Number: 700, Value: "toddoba boqol"},
		{Number: 600, Value: "lix boqol"},
		{Number: 500, Value: "shan boqol"},
		{Number: 400, Value: "afar boqol"},
		{Number: 300, Value: "saddex boqol"},
		{Number: 200, Value: "laba boqol"},
		{Number: 100, Value: "boqol"},
		{Number: 99, Value: "sagaashan iyo sagaal"},
		{Number: 98, Value: "sagaashan iyo siddeed"},
		{Number: 97, Value: "sagaashan iyo toddoba"},
		{Number: 96, Value: "sagaashan iyo lix"},
		{Number: 95, Value: "sagaashan iyo shan"},
		{Number: 94, Value: "sagaashan iyo afar"},
		{Number: 93, Value: "sagaashan iyo saddex"},
		{Number: 92, Value: "sagaashan iyo laba"},
		{Number: 91, Value: "sagaashan iyo kow"},
		{Number: 90, Value: "sagaashan"},
		{Number: 89, Value: "siddeetan iyo sagaal"},
		{Number: 88, Value: "siddeetan iyo siddeed"},
		{Number: 87, Value: "siddeetan iyo toddoba"},
		{Number: 86, Value: "siddeetan iyo lix"},
		{Number: 85, Value: "siddeetan iyo shan"},
		{Number: 84, Value: "siddeetan iyo afar"},
		{Number: 83, Value: "siddeetan iyo saddex"},
		{Number: 82, Value: "siddeetan iyo laba"},
		{Number: 81, Value: "siddeetan iyo kow"},
		{Number: 80, Value: "siddeetan"},
		{Number: 79, Value: "todobaatan iyo sagaal"},
		{Number: 78, Value: "todobaatan iyo siddeed"},
		{Number: 77, Value: "todobaatan iyo toddoba"},
		{Number: 76, Value: "todobaatan iyo lix"},
		{Number: 75, Value: "todobaatan iyo shan"},
		{Number: 74, Value: "todobaatan iyo afar"},
		{Number: 73, Value: "todobaatan iyo saddex"},
		{Number: 72, Value: "todobaatan iyo laba"},
		{Number: 71, Value: "todobaatan iyo kow"},
		{Number: 70, Value: "todobaatan"},
		{Number: 69, Value: "lixdatan iyo sagaal"},
		{Number: 68, Value: "lixdatan iyo siddeed"},
		{Number: 67, Value: "lixdatan iyo toddoba"},
		{Number: 66, Value: "lixdatan iyo lix"},
		{Number: 65, Value: "lixdatan iyo shan"},
		{Number: 64, Value: "lixdatan iyo afar"},
		{Number: 63, Value: "lixdatan iyo saddex"},
		{Number: 62, Value: "lixdatan iyo laba"},
		{Number: 61, Value: "lixdatan iyo kow"},
		{Number: 60, Value: "lixdatan"},
		{Number: 59, Value: "kontan iyo sagaal"},
		{Number: 58, Value: "kontan iyo siddeed"},
		{Number: 57, Value: "kontan iyo toddoba"},
		{Number: 56, Value: "kontan iyo lix"},
		{Number: 55, Value: "kontan iyo shan"},
		{Number: 54, Value: "kontan iyo afar"},
		{Number: 53, Value: "kontan iyo saddex"},
		{Number: 52, Value: "kontan iyo laba"},
		{Number: 51, Value: "kontan iyo kow"},
		{Number: 50, Value: "kontan"},
		{Number: 49, Value: "afaratan iyo sagaal"},
		{Number: 48, Value: "afaratan iyo siddeed"},
		{Number: 47, Value: "afaratan iyo toddoba"},
		{Number: 46, Value: "afaratan iyo lix"},
		{Number: 45, Value: "afaratan iyo shan"},
		{Number: 44, Value: "afaratan iyo afar"},
		{Number: 43, Value: "afaratan iyo saddex"},
		{Number: 42, Value: "afaratan iyo laba"},
		{Number: 41, Value: "afaratan iyo kow"},
		{Number: 40, Value: "afaratan"},
		{Number: 39, Value: "soddonka iyo sagaal"},
		{Number: 38, Value: "soddonka iyo siddeed"},
		{Number: 37, Value: "soddonka iyo toddoba"},
		{Number: 36, Value: "soddonka iyo lix"},
		{Number: 35, Value: "soddonka iyo shan"},
		{Number: 34, Value: "soddonka iyo afar"},
		{Number: 33, Value: "soddonka iyo saddex"},
		{Number: 32, Value: "soddonka iyo laba"},
		{Number: 31, Value: "soddonka iyo kow"},
		{Number: 30, Value: "soddonka"},
		{Number: 29, Value: "labaatan iyo sagaal"},
		{Number: 28, Value: "labaatan iyo siddeed"},
		{Number: 27, Value: "labaatan iyo toddoba"},
		{Number: 26, Value: "labaatan iyo lix"},
		{Number: 25, Value: "labaatan iyo shan"},
		{Number: 24, Value: "labaatan iyo afar"},
		{Number: 23, Value: "labaatan iyo saddex"},
		{Number: 22, Value: "labaatan iyo laba"},
		{Number: 21, Value: "labaatan iyo kow"},
		{Number: 20, Value: "labaatan"},
		{Number: 19, Value: "sagaal iyo toban"},
		{Number: 18, Value: "siddeed iyo toban"},
		{Number: 17, Value: "toddoba iyo toban"},
		{Number: 16, Value: "lix iyo toban"},
		{Number: 15, Value: "shan iyo toban"},
		{Number: 14, Value: "afar iyo toban"},
		{Number: 13, Value: "saddex iyo toban"},
		{Number: 12, Value: "laba iyo toban"},
		{Number: 11, Value: "kow iyo toban"},
		{Number: 10, Value: "toban"},
		{Number: 9, Value: "sagaal"},
		{Number: 8, Value: "siddeed"},
		{Number: 7, Value: "toddoba"},
		{Number: 6, Value: "lix"},
		{Number: 5, Value: "shan"},
		{Number: 4, Value: "afar"},
		{Number: 3, Value: "saddex"},
		{Number: 2, Value: "laba"},
		{Number: 1, Value: "kow"},
		{Number: 0, Value: "eber"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Boqol"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "kowaad", Suffix: "aad", Masculine: "kowaad", Feminine: "kowaad", Neuter: "kowaad"},
		{Number: 2, Word: "labaad", Suffix: "aad", Masculine: "labaad", Feminine: "labaad", Neuter: "labaad"},
		{Number: 3, Word: "saddexaad", Suffix: "aad", Masculine: "saddexaad", Feminine: "saddexaad", Neuter: "saddexaad"},
		{Number: 4, Word: "afaraad", Suffix: "aad", Masculine: "afaraad", Feminine: "afaraad", Neuter: "afaraad"},
		{Number: 5, Word: "shanaad", Suffix: "aad", Masculine: "shanaad", Feminine: "shanaad", Neuter: "shanaad"},
		{Number: 6, Word: "lixaad", Suffix: "aad", Masculine: "lixaad", Feminine: "lixaad", Neuter: "lixaad"},
		{Number: 7, Word: "todobaad", Suffix: "aad", Masculine: "todobaad", Feminine: "todobaad", Neuter: "todobaad"},
		{Number: 8, Word: "siddeedaad", Suffix: "aad", Masculine: "siddeedaad", Feminine: "siddeedaad", Neuter: "siddeedaad"},
		{Number: 9, Word: "sagaalaad", Suffix: "aad", Masculine: "sagaalaad", Feminine: "sagaalaad", Neuter: "sagaalaad"},
		{Number: 10, Word: "tobanaad", Suffix: "aad", Masculine: "tobanaad", Feminine: "tobanaad", Neuter: "tobanaad"},
		{Number: 11, Word: "kow iyo tobanaad", Suffix: "aad", Masculine: "kow iyo tobanaad", Feminine: "kow iyo tobanaad", Neuter: "kow iyo tobanaad"},
		{Number: 12, Word: "laba iyo tobanaad", Suffix: "aad", Masculine: "laba iyo tobanaad", Feminine: "laba iyo tobanaad", Neuter: "laba iyo tobanaad"},
		{Number: 20, Word: "labaatanaad", Suffix: "aad", Masculine: "labaatanaad", Feminine: "labaatanaad", Neuter: "labaatanaad"},
		{Number: 21, Word: "labaatan iyo kowaad", Suffix: "aad", Masculine: "labaatan iyo kowaad", Feminine: "labaatan iyo kowaad", Neuter: "labaatan iyo kowaad"},
		{Number: 30, Word: "soddnakaad", Suffix: "aad", Masculine: "soddnakaad", Feminine: "soddnakaad", Neuter: "soddnakaad"},
		{Number: 50, Word: "kontanaad", Suffix: "aad", Masculine: "kontanaad", Feminine: "kontanaad", Neuter: "kontanaad"},
		{Number: 100, Word: "boqolaad", Suffix: "aad", Masculine: "boqolaad", Feminine: "boqolaad", Neuter: "boqolaad"},
		{Number: 1000, Word: "kunaad", Suffix: "aad", Masculine: "kunaad", Feminine: "kunaad", Neuter: "kunaad"},
	},
	LocaleFormatter: &SomaliSomaliaFormatter{},
}

// SomaliSomaliaFormatter handles Somali (Somalia) formatting
type SomaliSomaliaFormatter struct{}

func (f *SomaliSomaliaFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *SomaliSomaliaFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *SomaliSomaliaFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *SomaliSomaliaFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *SomaliSomaliaFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *SomaliSomaliaFormatter) ChopDecimal(d decimal.Decimal, precision int) decimal.Decimal {
	return d.Truncate(int32(precision))
}

func (f *SomaliSomaliaFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *SomaliSomaliaFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
