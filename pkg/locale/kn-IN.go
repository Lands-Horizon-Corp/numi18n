package locale

import (
	"github.com/shopspring/decimal"
)

// INKNLocale is a NumI18NLocale configured for India (kn-IN)
var INKNLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "ರೂಪಾಯಿ",
		Plural:   "ರೂಪಾಯಿಗಳು",
		Singular: "ರೂಪಾಯಿ",
		Symbol:   "₹",
		FractionUnit: FractionUnit{
			Name:     "ಪೈಸೆ",
			Plural:   "ಪೈಸೆಗಳು",
			Singular: "ಪೈಸೆ",
			Symbol:   "ಪ",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "India",
		Currency:       "INR",
		ISO3166Alpha2:  "IN",
		ISO3166Alpha3:  "IND",
		ISO3166Numeric: "356",
		Locale:         "kn-IN",
		Timezone:       []string{"Asia/Kolkata"},
		Language:       "kn",
	},
	Texts: Texts{
		And:   "ಮತ್ತು",
		Minus: "ಮೈನಸ್",
		Only:  "ಮಾತ್ರ",
		Point: "ಬಿಂದು",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "ಒಂದು ಪದ್ಮ"},
		{Number: 1000000000000, Value: "ಒಂದು ನೀಲ"},
		{Number: 1000000000, Value: "ಒಂದು ಅಬ್ಜ"},
		{Number: 10000000, Value: "ಒಂದು ಕೋಟಿ"},
		{Number: 1000000, Value: "ಹತ್ತು ಲಕ್ಷ"},
		{Number: 100000, Value: "ಒಂದು ಲಕ್ಷ"},
		{Number: 1000, Value: "ಒಂದು ಸಾವಿರ"},
		{Number: 100, Value: "ಒಂದು ನೂರು"},
		{Number: 90, Value: "ತೊಂಬತ್ತು"},
		{Number: 80, Value: "ಎಂಬತ್ತು"},
		{Number: 70, Value: "ಎಪ್ಪತ್ತು"},
		{Number: 60, Value: "ಅರುವತ್ತು"},
		{Number: 50, Value: "ಐವತ್ತು"},
		{Number: 40, Value: "ನಲವತ್ತು"},
		{Number: 30, Value: "ಮೂವತ್ತು"},
		{Number: 20, Value: "ಇಪ್ಪತ್ತು"},
		{Number: 19, Value: "ಹತ್ತೊಂಬತ್ತು"},
		{Number: 18, Value: "ಹದಿನೆಂಟು"},
		{Number: 17, Value: "ಹದಿನೇಳು"},
		{Number: 16, Value: "ಹದಿನಾರು"},
		{Number: 15, Value: "ಹದಿನೈದು"},
		{Number: 14, Value: "ಹದಿನಾಲ್ಕು"},
		{Number: 13, Value: "ಹದಿಮೂರು"},
		{Number: 12, Value: "ಹನ್ನೆರಡು"},
		{Number: 11, Value: "ಹನ್ನೊಂದು"},
		{Number: 10, Value: "ಹತ್ತು"},
		{Number: 9, Value: "ಒಂಬತ್ತು"},
		{Number: 8, Value: "ಎಂಟು"},
		{Number: 7, Value: "ಏಳು"},
		{Number: 6, Value: "ಆರು"},
		{Number: 5, Value: "ಐದು"},
		{Number: 4, Value: "ನಾಲ್ಕು"},
		{Number: 3, Value: "ಮೂರು"},
		{Number: 2, Value: "ಎರಡು"},
		{Number: 1, Value: "ಒಂದು"},
		{Number: 0, Value: "ಶೂನ್ಯ"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "ಒಂದು ನೂರು"},
		{Number: 100000, Value: "ಒಂದು ಲಕ್ಷ"},
		{Number: 10000000, Value: "ಒಂದು ಕೋಟಿ"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "ಮೊದಲನೇ", Suffix: "-ನೇ", Masculine: "ಮೊದಲನೇ", Feminine: "ಮೊದಲನೇ", Neuter: "ಮೊದಲನೇ"},
		{Number: 2, Word: "ಎರಡನೇ", Suffix: "-ನೇ", Masculine: "ಎರಡನೇ", Feminine: "ಎರಡನೇ", Neuter: "ಎರಡನೇ"},
		{Number: 3, Word: "ಮೂರನೇ", Suffix: "-ನೇ", Masculine: "ಮೂರನೇ", Feminine: "ಮೂರನೇ", Neuter: "ಮೂರನೇ"},
		{Number: 4, Word: "ನಾಲ್ಕನೇ", Suffix: "-ನೇ", Masculine: "ನಾಲ್ಕನೇ", Feminine: "ನಾಲ್ಕನೇ", Neuter: "ನಾಲ್ಕನೇ"},
		{Number: 5, Word: "ಐದನೇ", Suffix: "-ನೇ", Masculine: "ಐದನೇ", Feminine: "ಐದನೇ", Neuter: "ಐದನೇ"},
		{Number: 6, Word: "ಆರನೇ", Suffix: "-ನೇ", Masculine: "ಆರನೇ", Feminine: "ಆರನೇ", Neuter: "ಆರನೇ"},
		{Number: 7, Word: "ಏಳನೇ", Suffix: "-ನೇ", Masculine: "ಏಳನೇ", Feminine: "ಏಳನೇ", Neuter: "ಏಳನೇ"},
		{Number: 8, Word: "ಎಂಟನೇ", Suffix: "-ನೇ", Masculine: "ಎಂಟನೇ", Feminine: "ಎಂಟನೇ", Neuter: "ಎಂಟನೇ"},
		{Number: 9, Word: "ಒಂಬತ್ತನೇ", Suffix: "-ನೇ", Masculine: "ಒಂಬತ್ತನೇ", Feminine: "ಒಂಬತ್ತನೇ", Neuter: "ಒಂಬತ್ತನೇ"},
		{Number: 10, Word: "ಹತ್ತನೇ", Suffix: "-ನೇ", Masculine: "ಹತ್ತನೇ", Feminine: "ಹತ್ತನೇ", Neuter: "ಹತ್ತನೇ"},
		{Number: 11, Word: "ಹನ್ನೊಂದನೇ", Suffix: "-ನೇ", Masculine: "ಹನ್ನೊಂದನೇ", Feminine: "ಹನ್ನೊಂದನೇ", Neuter: "ಹನ್ನೊಂದನೇ"},
		{Number: 12, Word: "ಹನ್ನೆರಡನೇ", Suffix: "-ನೇ", Masculine: "ಹನ್ನೆರಡನೇ", Feminine: "ಹನ್ನೆರಡನೇ", Neuter: "ಹನ್ನೆರಡನೇ"},
		{Number: 13, Word: "ಹದಿಮೂರನೇ", Suffix: "-ನೇ", Masculine: "ಹದಿಮೂರನೇ", Feminine: "ಹದಿಮೂರನೇ", Neuter: "ಹದಿಮೂರನೇ"},
		{Number: 14, Word: "ಹದಿನಾಲ್ಕನೇ", Suffix: "-ನೇ", Masculine: "ಹದಿನಾಲ್ಕನೇ", Feminine: "ಹದಿನಾಲ್ಕನೇ", Neuter: "ಹದಿನಾಲ್ಕನೇ"},
		{Number: 15, Word: "ಹದಿನೈದನೇ", Suffix: "-ನೇ", Masculine: "ಹದಿನೈದನೇ", Feminine: "ಹದಿನೈದನೇ", Neuter: "ಹದಿನೈದನೇ"},
		{Number: 16, Word: "ಹದಿನಾರನೇ", Suffix: "-ನೇ", Masculine: "ಹದಿನಾರನೇ", Feminine: "ಹದಿನಾರನೇ", Neuter: "ಹದಿನಾರನೇ"},
		{Number: 17, Word: "ಹದಿನೇಳನೇ", Suffix: "-ನೇ", Masculine: "ಹದಿನೇಳನೇ", Feminine: "ಹದಿನೇಳನೇ", Neuter: "ಹದಿನೇಳನೇ"},
		{Number: 18, Word: "ಹದಿನೆಂಟನೇ", Suffix: "-ನೇ", Masculine: "ಹದಿನೆಂಟನೇ", Feminine: "ಹದಿನೆಂಟನೇ", Neuter: "ಹದಿನೆಂಟನೇ"},
		{Number: 19, Word: "ಹತ್ತೊಂಬತ್ತನೇ", Suffix: "-ನೇ", Masculine: "ಹತ್ತೊಂಬತ್ತನೇ", Feminine: "ಹತ್ತೊಂಬತ್ತನೇ", Neuter: "ಹತ್ತೊಂಬತ್ತನೇ"},
		{Number: 20, Word: "ಇಪ್ಪತ್ತನೇ", Suffix: "-ನೇ", Masculine: "ಇಪ್ಪತ್ತನೇ", Feminine: "ಇಪ್ಪತ್ತನೇ", Neuter: "ಇಪ್ಪತ್ತನೇ"},
		{Number: 21, Word: "ಇಪ್ಪತ್ತೊಂದನೇ", Suffix: "-ನೇ", Masculine: "ಇಪ್ಪತ್ತೊಂದನೇ", Feminine: "ಇಪ್ಪತ್ತೊಂದನೇ", Neuter: "ಇಪ್ಪತ್ತೊಂದನೇ"},
		{Number: 30, Word: "ಮೂವತ್ತನೇ", Suffix: "-ನೇ", Masculine: "ಮೂವತ್ತನೇ", Feminine: "ಮೂವತ್ತನೇ", Neuter: "ಮೂವತ್ತನೇ"},
		{Number: 40, Word: "ನಲವತ್ತನೇ", Suffix: "-ನೇ", Masculine: "ನಲವತ್ತನೇ", Feminine: "ನಲವತ್ತನೇ", Neuter: "ನಲವತ್ತನೇ"},
		{Number: 50, Word: "ಐವತ್ತನೇ", Suffix: "-ನೇ", Masculine: "ಐವತ್ತನೇ", Feminine: "ಐವತ್ತನೇ", Neuter: "ಐವತ್ತನೇ"},
		{Number: 60, Word: "ಅರುವತ್ತನೇ", Suffix: "-ನೇ", Masculine: "ಅರುವತ್ತನೇ", Feminine: "ಅರುವತ್ತನೇ", Neuter: "ಅರುವತ್ತನೇ"},
		{Number: 70, Word: "ಎಪ್ಪತ್ತನೇ", Suffix: "-ನೇ", Masculine: "ಎಪ್ಪತ್ತನೇ", Feminine: "ಎಪ್ಪತ್ತನೇ", Neuter: "ಎಪ್ಪತ್ತನೇ"},
		{Number: 80, Word: "ಎಂಬತ್ತನೇ", Suffix: "-ನೇ", Masculine: "ಎಂಬತ್ತನೇ", Feminine: "ಎಂಬತ್ತನೇ", Neuter: "ಎಂಬತ್ತನೇ"},
		{Number: 90, Word: "ತೊಂಬತ್ತನೇ", Suffix: "-ನೇ", Masculine: "ತೊಂಬತ್ತನೇ", Feminine: "ತೊಂಬತ್ತನೇ", Neuter: "ತೊಂಬತ್ತನೇ"},
		{Number: 100, Word: "ನೂರನೇ", Suffix: "-ನೇ", Masculine: "ನೂರನೇ", Feminine: "ನೂರನೇ", Neuter: "ನೂರನೇ"},
		{Number: 1000, Word: "ಸಾವಿರನೇ", Suffix: "-ನೇ", Masculine: "ಸಾವಿರನೇ", Feminine: "ಸಾವಿರನೇ", Neuter: "ಸಾವಿರನೇ"},
	},
	LocaleFormatter: &KannadaFormatter{},
}

// KannadaFormatter handles Kannada (kn-IN) formatting
type KannadaFormatter struct{}

func (f *KannadaFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *KannadaFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *KannadaFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *KannadaFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *KannadaFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *KannadaFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 2
	}
	return amount.Truncate(int32(precision))
}

func (f *KannadaFormatter) FormatDecimalNumber(amount float64) string {
	return FormatAngloDecimal(amount)
}
func (f *KannadaFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with Anglo conventions (comma separators, period decimal, prefix symbol)
	return FormatAngloCurrency(amount, currencySymbol)
}
