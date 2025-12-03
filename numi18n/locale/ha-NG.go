package locale

import (
	"github.com/shopspring/decimal"
)

// NGHALocale is a NumI18NLocale configured for Nigeria (ha-NG)
var NGHALocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Naira",
		Plural:   "Naira",
		Singular: "Naira",
		Symbol:   "₦",
		FractionUnit: FractionUnit{
			Name:     "Kobo",
			Plural:   "Kobo",
			Singular: "Kobo",
			Symbol:   "k",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Nigeria",
		Currency:       "NGN",
		ISO3166Alpha2:  "NG",
		ISO3166Alpha3:  "NGA",
		ISO3166Numeric: "566",
		Locale:         "ha-NG",
		Timezone:       []string{"Africa/Lagos"},
		Language:       "ha",
		Emoji:          "🇳🇬",
		PhoneCode:      "+234",
		Domain:         ".ng",
	},
	Texts: Texts{
		And:   "Da",
		Minus: "Rashi",
		Only:  "Kawai",
		Point: "Aya",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Quadrillion daya"},
		{Number: 1000000000000, Value: "Trillion daya"},
		{Number: 1000000000, Value: "Billion daya"},
		{Number: 1000000, Value: "Miliyan daya"},
		{Number: 1000, Value: "Dubu daya"},
		{Number: 100, Value: "Dari daya"},
		{Number: 90, Value: "Casa'in"},
		{Number: 80, Value: "Tamanin"},
		{Number: 70, Value: "Saba'in"},
		{Number: 60, Value: "Sittin"},
		{Number: 50, Value: "Hamsin"},
		{Number: 40, Value: "Arba'in"},
		{Number: 30, Value: "Talatin"},
		{Number: 20, Value: "Ashirin"},
		{Number: 19, Value: "Sha tara"},
		{Number: 18, Value: "Sha takwas"},
		{Number: 17, Value: "Sha bakwai"},
		{Number: 16, Value: "Sha shida"},
		{Number: 15, Value: "Sha biyar"},
		{Number: 14, Value: "Sha hudu"},
		{Number: 13, Value: "Sha uku"},
		{Number: 12, Value: "Sha biyu"},
		{Number: 11, Value: "Sha daya"},
		{Number: 10, Value: "Goma"},
		{Number: 9, Value: "Tara"},
		{Number: 8, Value: "Takwas"},
		{Number: 7, Value: "Bakwai"},
		{Number: 6, Value: "Shida"},
		{Number: 5, Value: "Biyar"},
		{Number: 4, Value: "Hudu"},
		{Number: 3, Value: "Uku"},
		{Number: 2, Value: "Biyu"},
		{Number: 1, Value: "Daya"},
		{Number: 0, Value: "Sifiri"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Dari daya"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "Na farko", Suffix: ".", Masculine: "Na farko", Feminine: "Ta farko", Neuter: "Na farko"},
		{Number: 2, Word: "Na biyu", Suffix: ".", Masculine: "Na biyu", Feminine: "Ta biyu", Neuter: "Na biyu"},
		{Number: 3, Word: "Na uku", Suffix: ".", Masculine: "Na uku", Feminine: "Ta uku", Neuter: "Na uku"},
		{Number: 4, Word: "Na hudu", Suffix: ".", Masculine: "Na hudu", Feminine: "Ta hudu", Neuter: "Na hudu"},
		{Number: 5, Word: "Na biyar", Suffix: ".", Masculine: "Na biyar", Feminine: "Ta biyar", Neuter: "Na biyar"},
		{Number: 6, Word: "Na shida", Suffix: ".", Masculine: "Na shida", Feminine: "Ta shida", Neuter: "Na shida"},
		{Number: 7, Word: "Na bakwai", Suffix: ".", Masculine: "Na bakwai", Feminine: "Ta bakwai", Neuter: "Na bakwai"},
		{Number: 8, Word: "Na takwas", Suffix: ".", Masculine: "Na takwas", Feminine: "Ta takwas", Neuter: "Na takwas"},
		{Number: 9, Word: "Na tara", Suffix: ".", Masculine: "Na tara", Feminine: "Ta tara", Neuter: "Na tara"},
		{Number: 10, Word: "Na goma", Suffix: ".", Masculine: "Na goma", Feminine: "Ta goma", Neuter: "Na goma"},
		{Number: 11, Word: "Na sha daya", Suffix: ".", Masculine: "Na sha daya", Feminine: "Ta sha daya", Neuter: "Na sha daya"},
		{Number: 12, Word: "Na sha biyu", Suffix: ".", Masculine: "Na sha biyu", Feminine: "Ta sha biyu", Neuter: "Na sha biyu"},
		{Number: 13, Word: "Na sha uku", Suffix: ".", Masculine: "Na sha uku", Feminine: "Ta sha uku", Neuter: "Na sha uku"},
		{Number: 14, Word: "Na sha hudu", Suffix: ".", Masculine: "Na sha hudu", Feminine: "Ta sha hudu", Neuter: "Na sha hudu"},
		{Number: 15, Word: "Na sha biyar", Suffix: ".", Masculine: "Na sha biyar", Feminine: "Ta sha biyar", Neuter: "Na sha biyar"},
		{Number: 16, Word: "Na sha shida", Suffix: ".", Masculine: "Na sha shida", Feminine: "Ta sha shida", Neuter: "Na sha shida"},
		{Number: 17, Word: "Na sha bakwai", Suffix: ".", Masculine: "Na sha bakwai", Feminine: "Ta sha bakwai", Neuter: "Na sha bakwai"},
		{Number: 18, Word: "Na sha takwas", Suffix: ".", Masculine: "Na sha takwas", Feminine: "Ta sha takwas", Neuter: "Na sha takwas"},
		{Number: 19, Word: "Na sha tara", Suffix: ".", Masculine: "Na sha tara", Feminine: "Ta sha tara", Neuter: "Na sha tara"},
		{Number: 20, Word: "Na ashirin", Suffix: ".", Masculine: "Na ashirin", Feminine: "Ta ashirin", Neuter: "Na ashirin"},
		{Number: 21, Word: "Na ashirin da daya", Suffix: ".", Masculine: "Na ashirin da daya", Feminine: "Ta ashirin da daya", Neuter: "Na ashirin da daya"},
		{Number: 30, Word: "Na talatin", Suffix: ".", Masculine: "Na talatin", Feminine: "Ta talatin", Neuter: "Na talatin"},
		{Number: 40, Word: "Na arba'in", Suffix: ".", Masculine: "Na arba'in", Feminine: "Ta arba'in", Neuter: "Na arba'in"},
		{Number: 50, Word: "Na hamsin", Suffix: ".", Masculine: "Na hamsin", Feminine: "Ta hamsin", Neuter: "Na hamsin"},
		{Number: 60, Word: "Na sittin", Suffix: ".", Masculine: "Na sittin", Feminine: "Ta sittin", Neuter: "Na sittin"},
		{Number: 70, Word: "Na saba'in", Suffix: ".", Masculine: "Na saba'in", Feminine: "Ta saba'in", Neuter: "Na saba'in"},
		{Number: 80, Word: "Na tamanin", Suffix: ".", Masculine: "Na tamanin", Feminine: "Ta tamanin", Neuter: "Na tamanin"},
		{Number: 90, Word: "Na casa'in", Suffix: ".", Masculine: "Na casa'in", Feminine: "Ta casa'in", Neuter: "Na casa'in"},
		{Number: 100, Word: "Na dari", Suffix: ".", Masculine: "Na dari", Feminine: "Ta dari", Neuter: "Na dari"},
		{Number: 1000, Word: "Na dubu", Suffix: ".", Masculine: "Na dubu", Feminine: "Ta dubu", Neuter: "Na dubu"},
	},
	LocaleFormatter: &HausaFormatter{},
}

// HausaFormatter handles Hausa (ha-NG) formatting
type HausaFormatter struct{}

func (f *HausaFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *HausaFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	// Naira is the same for both singular and plural in Hausa
	return result + " " + currencyName
}

func (f *HausaFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *HausaFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	// Kobo is the same for both singular and plural in Hausa
	return result + " " + fractionName
}

func (f *HausaFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *HausaFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	// Standard decimal chopping - round to specified precision
	return amount.Round(int32(precision))
}

func (f *HausaFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *HausaFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
