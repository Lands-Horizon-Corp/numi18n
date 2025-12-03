package locale

import (
	"github.com/shopspring/decimal"
)

// ORINLocale represents the Odia (India) locale
var ORINLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "ଭାରତୀୟ ଟଙ୍କା",
		Plural:   "ଭାରତୀୟ ଟଙ୍କା",
		Singular: "ଭାରତୀୟ ଟଙ୍କା",
		Symbol:   "₹",
		FractionUnit: FractionUnit{
			Name:     "ପଇସା",
			Plural:   "ପଇସା",
			Singular: "ପଇସା",
			Symbol:   "p",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "India",
		Currency:       "INR",
		ISO3166Alpha2:  "IN",
		ISO3166Alpha3:  "IND",
		ISO3166Numeric: "356",
		Locale:         "or-IN",
		Timezone:       []string{"Asia/Kolkata"},
		Language:       "or",
	},
	Texts: Texts{
		And:   "ଏବଂ",
		Minus: "ମାଇନସ",
		Only:  "କେବଳ",
		Point: "ଦଶମିକ",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 10000000000000, Value: "ଦଶ ଖରବ"},
		{Number: 1000000000000, Value: "ଏକ ଖରବ"},
		{Number: 100000000000, Value: "ଦଶ ନୀଳ"},
		{Number: 10000000000, Value: "ଏକ ନୀଳ"},
		{Number: 1000000000, Value: "ଏକ ଅରବ"},
		{Number: 100000000, Value: "ଦଶ କୋଟି"},
		{Number: 10000000, Value: "ଏକ କୋଟି"},
		{Number: 1000000, Value: "ଦଶ ଲକ୍ଷ"},
		{Number: 100000, Value: "ଏକ ଲକ୍ଷ"},
		{Number: 90000, Value: "ନବେ ହଜାର"},
		{Number: 80000, Value: "ଅଶୀ ହଜାର"},
		{Number: 70000, Value: "ସତୁରି ହଜାର"},
		{Number: 60000, Value: "ଷାଠିଏ ହଜାର"},
		{Number: 50000, Value: "ପଚାଶ ହଜାର"},
		{Number: 40000, Value: "ଚାଳିଶ ହଜାର"},
		{Number: 30000, Value: "ତିରିଶ ହଜାର"},
		{Number: 20000, Value: "କୋଡ଼ିଏ ହଜାର"},
		{Number: 19000, Value: "ଉଣେଇଶ ହଜାର"},
		{Number: 18000, Value: "ଅଠର ହଜାର"},
		{Number: 17000, Value: "ସତର ହଜାର"},
		{Number: 16000, Value: "ଷୋଡ଼ଶ ହଜାର"},
		{Number: 15000, Value: "ପନ୍ଦର ହଜାର"},
		{Number: 14000, Value: "ଚଉଦ ହଜାର"},
		{Number: 13000, Value: "ତେର ହଜାର"},
		{Number: 12000, Value: "ବାର ହଜାର"},
		{Number: 11000, Value: "ଏଗାର ହଜାର"},
		{Number: 10000, Value: "ଦଶ ହଜାର"},
		{Number: 9000, Value: "ନଅ ହଜାର"},
		{Number: 8000, Value: "ଆଠ ହଜାର"},
		{Number: 7000, Value: "ସାତ ହଜାର"},
		{Number: 6000, Value: "ଛଅ ହଜାର"},
		{Number: 5000, Value: "ପାଞ୍ଚ ହଜାର"},
		{Number: 4000, Value: "ଚାରି ହଜାର"},
		{Number: 3000, Value: "ତିନି ହଜାର"},
		{Number: 2000, Value: "ଦୁଇ ହଜାର"},
		{Number: 1000, Value: "ଏକ ହଜାର"},
		{Number: 900, Value: "ନଅ ଶହ"},
		{Number: 800, Value: "ଆଠ ଶହ"},
		{Number: 700, Value: "ସାତ ଶହ"},
		{Number: 600, Value: "ଛଅ ଶହ"},
		{Number: 500, Value: "ପାଞ୍ଚ ଶହ"},
		{Number: 400, Value: "ଚାରି ଶହ"},
		{Number: 300, Value: "ତିନି ଶହ"},
		{Number: 200, Value: "ଦୁଇ ଶହ"},
		{Number: 100, Value: "ଏକ ଶହ"},
		{Number: 99, Value: "ନବେ ନଅ"},
		{Number: 98, Value: "ନବେ ଆଠ"},
		{Number: 97, Value: "ନବେ ସାତ"},
		{Number: 96, Value: "ନବେ ଛଅ"},
		{Number: 95, Value: "ନବେ ପାଞ୍ଚ"},
		{Number: 94, Value: "ନବେ ଚାରି"},
		{Number: 93, Value: "ନବେ ତିନି"},
		{Number: 92, Value: "ନବେ ଦୁଇ"},
		{Number: 91, Value: "ନବେ ଏକ"},
		{Number: 90, Value: "ନବେ"},
		{Number: 89, Value: "ଅଶୀ ନଅ"},
		{Number: 88, Value: "ଅଶୀ ଆଠ"},
		{Number: 87, Value: "ଅଶୀ ସାତ"},
		{Number: 86, Value: "ଅଶୀ ଛଅ"},
		{Number: 85, Value: "ଅଶୀ ପାଞ୍ଚ"},
		{Number: 84, Value: "ଅଶୀ ଚାରି"},
		{Number: 83, Value: "ଅଶୀ ତିନି"},
		{Number: 82, Value: "ଅଶୀ ଦୁଇ"},
		{Number: 81, Value: "ଅଶୀ ଏକ"},
		{Number: 80, Value: "ଅଶୀ"},
		{Number: 79, Value: "ସତୁରି ନଅ"},
		{Number: 78, Value: "ସତୁରି ଆଠ"},
		{Number: 77, Value: "ସତୁରି ସାତ"},
		{Number: 76, Value: "ସତୁରି ଛଅ"},
		{Number: 75, Value: "ସତୁରି ପାଞ୍ଚ"},
		{Number: 74, Value: "ସତୁରି ଚାରି"},
		{Number: 73, Value: "ସତୁରି ତିନି"},
		{Number: 72, Value: "ସତୁରି ଦୁଇ"},
		{Number: 71, Value: "ସତୁରି ଏକ"},
		{Number: 70, Value: "ସତୁରି"},
		{Number: 69, Value: "ଷାଠିଏ ନଅ"},
		{Number: 68, Value: "ଷାଠିଏ ଆଠ"},
		{Number: 67, Value: "ଷାଠିଏ ସାତ"},
		{Number: 66, Value: "ଷାଠିଏ ଛଅ"},
		{Number: 65, Value: "ଷାଠିଏ ପାଞ୍ଚ"},
		{Number: 64, Value: "ଷାଠିଏ ଚାରି"},
		{Number: 63, Value: "ଷାଠିଏ ତିନି"},
		{Number: 62, Value: "ଷାଠିଏ ଦୁଇ"},
		{Number: 61, Value: "ଷାଠିଏ ଏକ"},
		{Number: 60, Value: "ଷାଠିଏ"},
		{Number: 59, Value: "ପଚାଶ ନଅ"},
		{Number: 58, Value: "ପଚାଶ ଆଠ"},
		{Number: 57, Value: "ପଚାଶ ସାତ"},
		{Number: 56, Value: "ପଚାଶ ଛଅ"},
		{Number: 55, Value: "ପଚାଶ ପାଞ୍ଚ"},
		{Number: 54, Value: "ପଚାଶ ଚାରି"},
		{Number: 53, Value: "ପଚାଶ ତିନି"},
		{Number: 52, Value: "ପଚାଶ ଦୁଇ"},
		{Number: 51, Value: "ପଚାଶ ଏକ"},
		{Number: 50, Value: "ପଚାଶ"},
		{Number: 49, Value: "ଚାଳିଶ ନଅ"},
		{Number: 48, Value: "ଚାଳିଶ ଆଠ"},
		{Number: 47, Value: "ଚାଳିଶ ସାତ"},
		{Number: 46, Value: "ଚାଳିଶ ଛଅ"},
		{Number: 45, Value: "ଚାଳିଶ ପାଞ୍ଚ"},
		{Number: 44, Value: "ଚାଳିଶ ଚାରି"},
		{Number: 43, Value: "ଚାଳିଶ ତିନି"},
		{Number: 42, Value: "ଚାଳିଶ ଦୁଇ"},
		{Number: 41, Value: "ଚାଳିଶ ଏକ"},
		{Number: 40, Value: "ଚାଳିଶ"},
		{Number: 39, Value: "ତିରିଶ ନଅ"},
		{Number: 38, Value: "ତିରିଶ ଆଠ"},
		{Number: 37, Value: "ତିରିଶ ସାତ"},
		{Number: 36, Value: "ତିରିଶ ଛଅ"},
		{Number: 35, Value: "ତିରିଶ ପାଞ୍ଚ"},
		{Number: 34, Value: "ତିରିଶ ଚାରି"},
		{Number: 33, Value: "ତିରିଶ ତିନି"},
		{Number: 32, Value: "ତିରିଶ ଦୁଇ"},
		{Number: 31, Value: "ତିରିଶ ଏକ"},
		{Number: 30, Value: "ତିରିଶ"},
		{Number: 29, Value: "କୋଡ଼ିଏ ନଅ"},
		{Number: 28, Value: "କୋଡ଼ିଏ ଆଠ"},
		{Number: 27, Value: "କୋଡ଼ିଏ ସାତ"},
		{Number: 26, Value: "କୋଡ଼ିଏ ଛଅ"},
		{Number: 25, Value: "କୋଡ଼ିଏ ପାଞ୍ଚ"},
		{Number: 24, Value: "କୋଡ଼ିଏ ଚାରି"},
		{Number: 23, Value: "କୋଡ଼ିଏ ତିନି"},
		{Number: 22, Value: "କୋଡ଼ିଏ ଦୁଇ"},
		{Number: 21, Value: "କୋଡ଼ିଏ ଏକ"},
		{Number: 20, Value: "କୋଡ଼ିଏ"},
		{Number: 19, Value: "ଉଣେଇଶ"},
		{Number: 18, Value: "ଅଠର"},
		{Number: 17, Value: "ସତର"},
		{Number: 16, Value: "ଷୋଡ଼ଶ"},
		{Number: 15, Value: "ପନ୍ଦର"},
		{Number: 14, Value: "ଚଉଦ"},
		{Number: 13, Value: "ତେର"},
		{Number: 12, Value: "ବାର"},
		{Number: 11, Value: "ଏଗାର"},
		{Number: 10, Value: "ଦଶ"},
		{Number: 9, Value: "ନଅ"},
		{Number: 8, Value: "ଆଠ"},
		{Number: 7, Value: "ସାତ"},
		{Number: 6, Value: "ଛଅ"},
		{Number: 5, Value: "ପାଞ୍ଚ"},
		{Number: 4, Value: "ଚାରି"},
		{Number: 3, Value: "ତିନି"},
		{Number: 2, Value: "ଦୁଇ"},
		{Number: 1, Value: "ଏକ"},
		{Number: 0, Value: "ଶୂନ୍ୟ"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "ଏକ ଶହ"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "ପ୍ରଥମ", Suffix: "-ମ", Masculine: "ପ୍ରଥମ", Feminine: "ପ୍ରଥମ", Neuter: "ପ୍ରଥମ"},
		{Number: 2, Word: "ଦ୍ୱିତୀୟ", Suffix: "-ୟ", Masculine: "ଦ୍ୱିତୀୟ", Feminine: "ଦ୍ୱିତୀୟ", Neuter: "ଦ୍ୱିତୀୟ"},
		{Number: 3, Word: "ତୃତୀୟ", Suffix: "-ୟ", Masculine: "ତୃତୀୟ", Feminine: "ତୃତୀୟ", Neuter: "ତୃତୀୟ"},
		{Number: 4, Word: "ଚତୁର୍ଥ", Suffix: "-ଥ", Masculine: "ଚତୁର୍ଥ", Feminine: "ଚତୁର୍ଥ", Neuter: "ଚତୁର୍ଥ"},
		{Number: 5, Word: "ପଞ୍ଚମ", Suffix: "-ମ", Masculine: "ପଞ୍ଚମ", Feminine: "ପଞ୍ଚମ", Neuter: "ପଞ୍ଚମ"},
		{Number: 6, Word: "ଷଷ୍ଠ", Suffix: "-ଷ୍ଠ", Masculine: "ଷଷ୍ଠ", Feminine: "ଷଷ୍ଠ", Neuter: "ଷଷ୍ଠ"},
		{Number: 7, Word: "ସପ୍ତମ", Suffix: "-ମ", Masculine: "ସପ୍ତମ", Feminine: "ସପ୍ତମ", Neuter: "ସପ୍ତମ"},
		{Number: 8, Word: "ଅଷ୍ଟମ", Suffix: "-ମ", Masculine: "ଅଷ୍ଟମ", Feminine: "ଅଷ୍ଟମ", Neuter: "ଅଷ୍ଟମ"},
		{Number: 9, Word: "ନବମ", Suffix: "-ମ", Masculine: "ନବମ", Feminine: "ନବମ", Neuter: "ନବମ"},
		{Number: 10, Word: "ଦଶମ", Suffix: "-ମ", Masculine: "ଦଶମ", Feminine: "ଦଶମ", Neuter: "ଦଶମ"},
		{Number: 11, Word: "ଏକାଦଶ", Suffix: "-ଶ", Masculine: "ଏକାଦଶ", Feminine: "ଏକାଦଶ", Neuter: "ଏକାଦଶ"},
		{Number: 12, Word: "ଦ୍ୱାଦଶ", Suffix: "-ଶ", Masculine: "ଦ୍ୱାଦଶ", Feminine: "ଦ୍ୱାଦଶ", Neuter: "ଦ୍ୱାଦଶ"},
		{Number: 20, Word: "ବିଂଶତିତମ", Suffix: "-ତମ", Masculine: "ବିଂଶତିତମ", Feminine: "ବିଂଶତିତମ", Neuter: "ବିଂଶତିତମ"},
		{Number: 21, Word: "ଏକବିଂଶତିତମ", Suffix: "-ତମ", Masculine: "ଏକବିଂଶତିତମ", Feminine: "ଏକବିଂଶତିତମ", Neuter: "ଏକବିଂଶତିତମ"},
		{Number: 30, Word: "ତ୍ରିଂଶତିତମ", Suffix: "-ତମ", Masculine: "ତ୍ରିଂଶତିତମ", Feminine: "ତ୍ରିଂଶତିତମ", Neuter: "ତ୍ରିଂଶତିତମ"},
		{Number: 100, Word: "ଶତତମ", Suffix: "-ତମ", Masculine: "ଶତତମ", Feminine: "ଶତତମ", Neuter: "ଶତତମ"},
		{Number: 1000, Word: "ସହସ୍ରତମ", Suffix: "-ତମ", Masculine: "ସହସ୍ରତମ", Feminine: "ସହସ୍ରତମ", Neuter: "ସହସ୍ରତମ"},
	},
	LocaleFormatter: &OdiaFormatter{},
}

// OdiaFormatter handles Odia (India) formatting
type OdiaFormatter struct{}

func (f *OdiaFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *OdiaFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *OdiaFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *OdiaFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *OdiaFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *OdiaFormatter) ChopDecimal(decimal decimal.Decimal, precision int) decimal.Decimal {
	return decimal.Truncate(int32(precision))
}

func (f *OdiaFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *OdiaFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
