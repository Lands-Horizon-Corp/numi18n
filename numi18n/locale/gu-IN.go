package locale

import (
	"github.com/shopspring/decimal"
)

// INGULocale is a NumI18NLocale configured for India (gu-IN)
var INGULocale = NumI18NLocale{
	LocaleFormatter: &GujaratiFormatter{},
	Currency: Currency{
		Name:     "રૂપિયો",
		Plural:   "રૂપિયા",
		Singular: "રૂપિયો",
		Symbol:   "₹",
		FractionUnit: FractionUnit{
			Name:     "પૈસો",
			Plural:   "પૈસા",
			Singular: "પૈસો",
			Symbol:   "પ",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "India",
		Currency:       "INR",
		ISO3166Alpha2:  "IN",
		ISO3166Alpha3:  "IND",
		ISO3166Numeric: "356",
		Locale:         "gu-IN",
		Timezone:       []string{"Asia/Kolkata"},
		Language:       "gu",
		Emoji:          "🇮🇳",
		PhoneCode:      "+91",
		Domain:         ".in",
	},
	Texts: Texts{
		And:   "અને",
		Minus: "ઋણ",
		Only:  "માત્ર",
		Point: "બિંદુ",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "એક પદમ"},
		{Number: 1000000000000, Value: "એક નીલ"},
		{Number: 1000000000, Value: "એક અબજ"},
		{Number: 1000000, Value: "એક લાખ"},
		{Number: 100000, Value: "એક લાખ"},
		{Number: 1000, Value: "એક હજાર"},
		{Number: 100, Value: "એક સો"},
		{Number: 90, Value: "નેવું"},
		{Number: 80, Value: "એંસી"},
		{Number: 70, Value: "સિત્તેર"},
		{Number: 60, Value: "સાઠ"},
		{Number: 50, Value: "પચાસ"},
		{Number: 40, Value: "ચાલીસ"},
		{Number: 30, Value: "ત્રીસ"},
		{Number: 20, Value: "વીસ"},
		{Number: 19, Value: "ઓગણીસ"},
		{Number: 18, Value: "અઢાર"},
		{Number: 17, Value: "સત્તર"},
		{Number: 16, Value: "સોળ"},
		{Number: 15, Value: "પંદર"},
		{Number: 14, Value: "ચૌદ"},
		{Number: 13, Value: "તેર"},
		{Number: 12, Value: "બાર"},
		{Number: 11, Value: "અગિયાર"},
		{Number: 10, Value: "દસ"},
		{Number: 9, Value: "નવ"},
		{Number: 8, Value: "આઠ"},
		{Number: 7, Value: "સાત"},
		{Number: 6, Value: "છ"},
		{Number: 5, Value: "પાંચ"},
		{Number: 4, Value: "ચાર"},
		{Number: 3, Value: "ત્રણ"},
		{Number: 2, Value: "બે"},
		{Number: 1, Value: "એક"},
		{Number: 0, Value: "શૂન્ય"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "એક સો"},
		{Number: 100000, Value: "એક લાખ"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "પહેલો", Suffix: "મો", Masculine: "પહેલો", Feminine: "પહેલી", Neuter: "પહેલું"},
		{Number: 2, Word: "બીજો", Suffix: "મો", Masculine: "બીજો", Feminine: "બીજી", Neuter: "બીજું"},
		{Number: 3, Word: "ત્રીજો", Suffix: "મો", Masculine: "ત્રીજો", Feminine: "ત્રીજી", Neuter: "ત્રીજું"},
		{Number: 4, Word: "ચોથો", Suffix: "મો", Masculine: "ચોથો", Feminine: "ચોથી", Neuter: "ચોથું"},
		{Number: 5, Word: "પાંચમો", Suffix: "મો", Masculine: "પાંચમો", Feminine: "પાંચમી", Neuter: "પાંચમું"},
		{Number: 6, Word: "છઠ્ઠો", Suffix: "મો", Masculine: "છઠ્ઠો", Feminine: "છઠ્ઠી", Neuter: "છઠ્ઠું"},
		{Number: 7, Word: "સાતમો", Suffix: "મો", Masculine: "સાતમો", Feminine: "સાતમી", Neuter: "સાતમું"},
		{Number: 8, Word: "આઠમો", Suffix: "મો", Masculine: "આઠમો", Feminine: "આઠમી", Neuter: "આઠમું"},
		{Number: 9, Word: "નવમો", Suffix: "મો", Masculine: "નવમો", Feminine: "નવમી", Neuter: "નવમું"},
		{Number: 10, Word: "દસમો", Suffix: "મો", Masculine: "દસમો", Feminine: "દસમી", Neuter: "દસમું"},
		{Number: 11, Word: "અગિયારમો", Suffix: "મો", Masculine: "અગિયારમો", Feminine: "અગિયારમી", Neuter: "અગિયારમું"},
		{Number: 12, Word: "બારમો", Suffix: "મો", Masculine: "બારમો", Feminine: "બારમી", Neuter: "બારમું"},
		{Number: 13, Word: "તેરમો", Suffix: "મો", Masculine: "તેરમો", Feminine: "તેરમી", Neuter: "તેરમું"},
		{Number: 14, Word: "ચૌદમો", Suffix: "મો", Masculine: "ચૌદમો", Feminine: "ચૌદમી", Neuter: "ચૌદમું"},
		{Number: 15, Word: "પંદરમો", Suffix: "મો", Masculine: "પંદરમો", Feminine: "પંદરમી", Neuter: "પંદરમું"},
		{Number: 16, Word: "સોળમો", Suffix: "મો", Masculine: "સોળમો", Feminine: "સોળમી", Neuter: "સોળમું"},
		{Number: 17, Word: "સત્તરમો", Suffix: "મો", Masculine: "સત્તરમો", Feminine: "સત્તરમી", Neuter: "સત્તરમું"},
		{Number: 18, Word: "અઢારમો", Suffix: "મો", Masculine: "અઢારમો", Feminine: "અઢારમી", Neuter: "અઢારમું"},
		{Number: 19, Word: "ઓગણીસમો", Suffix: "મો", Masculine: "ઓગણીસમો", Feminine: "ઓગણીસમી", Neuter: "ઓગણીસમું"},
		{Number: 20, Word: "વીસમો", Suffix: "મો", Masculine: "વીસમો", Feminine: "વીસમી", Neuter: "વીસમું"},
		{Number: 21, Word: "એકવીસમો", Suffix: "મો", Masculine: "એકવીસમો", Feminine: "એકવીસમી", Neuter: "એકવીસમું"},
		{Number: 30, Word: "ત્રીસમો", Suffix: "મો", Masculine: "ત્રીસમો", Feminine: "ત્રીસમી", Neuter: "ત્રીસમું"},
		{Number: 40, Word: "ચાલીસમો", Suffix: "મો", Masculine: "ચાલીસમો", Feminine: "ચાલીસમી", Neuter: "ચાલીસમું"},
		{Number: 50, Word: "પચાસમો", Suffix: "મો", Masculine: "પચાસમો", Feminine: "પચાસમી", Neuter: "પચાસમું"},
		{Number: 60, Word: "સાઠમો", Suffix: "મો", Masculine: "સાઠમો", Feminine: "સાઠમી", Neuter: "સાઠમું"},
		{Number: 70, Word: "સિત્તેરમો", Suffix: "મો", Masculine: "સિત્તેરમો", Feminine: "સિત્તેરમી", Neuter: "સિત્તેરમું"},
		{Number: 80, Word: "એંસીમો", Suffix: "મો", Masculine: "એંસીમો", Feminine: "એંસીમી", Neuter: "એંસીમું"},
		{Number: 90, Word: "નેવુંમો", Suffix: "મો", Masculine: "નેવુંમો", Feminine: "નેવુંમી", Neuter: "નેવુંમું"},
		{Number: 100, Word: "સોમો", Suffix: "મો", Masculine: "સોમો", Feminine: "સોમી", Neuter: "સોમું"},
		{Number: 1000, Word: "હજારમો", Suffix: "મો", Masculine: "હજારમો", Feminine: "હજારમી", Neuter: "હજારમું"},
	},
}

// GujaratiFormatter handles Gujarati (gu-IN) formatting
type GujaratiFormatter struct{}

func (f *GujaratiFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *GujaratiFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *GujaratiFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *GujaratiFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *GujaratiFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *GujaratiFormatter) ChopDecimal(d decimal.Decimal, precision int) decimal.Decimal {
	return d.Truncate(int32(precision))
}

func (f *GujaratiFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *GujaratiFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
