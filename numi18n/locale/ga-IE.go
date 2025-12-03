package locale

import (
	"github.com/shopspring/decimal"
)

// IEGALocale is a NumI18NLocale configured for Ireland (ga-IE)
var IEGALocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Euro",
		Plural:   "Eorónna",
		Singular: "Euro",
		Symbol:   "€",
		FractionUnit: FractionUnit{
			Name:     "Cent",
			Plural:   "Ceinteannaí",
			Singular: "Cent",
			Symbol:   "c",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Ireland",
		Currency:       "EUR",
		ISO3166Alpha2:  "IE",
		ISO3166Alpha3:  "IRL",
		ISO3166Numeric: "372",
		Locale:         "ga-IE",
		Timezone:       []string{"Europe/Dublin"},
		Language:       "ga",
		Emoji:          "🇮🇪",
	},
	Texts: Texts{
		And:   "Agus",
		Minus: "Lúide",
		Only:  "Amháin",
		Point: "Pointe",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Mil billiún billiún"},
		{Number: 1000000000000, Value: "Billiún billiún"},
		{Number: 1000000000, Value: "Billiún"},
		{Number: 1000000, Value: "Milliún"},
		{Number: 1000, Value: "Míle"},
		{Number: 100, Value: "Céad"},
		{Number: 90, Value: "Nócha"},
		{Number: 80, Value: "Ochtó"},
		{Number: 70, Value: "Seachtó"},
		{Number: 60, Value: "Seasca"},
		{Number: 50, Value: "Caoga"},
		{Number: 40, Value: "Daichead"},
		{Number: 30, Value: "Tríocha"},
		{Number: 20, Value: "Fiche"},
		{Number: 19, Value: "Naoi déag"},
		{Number: 18, Value: "Ocht déag"},
		{Number: 17, Value: "Seacht déag"},
		{Number: 16, Value: "Sé déag"},
		{Number: 15, Value: "Cúig déag"},
		{Number: 14, Value: "Ceithre déag"},
		{Number: 13, Value: "Trí déag"},
		{Number: 12, Value: "Dó dhéag"},
		{Number: 11, Value: "Aon déag"},
		{Number: 10, Value: "Deich"},
		{Number: 9, Value: "Naoi"},
		{Number: 8, Value: "Ocht"},
		{Number: 7, Value: "Seacht"},
		{Number: 6, Value: "Sé"},
		{Number: 5, Value: "Cúig"},
		{Number: 4, Value: "Ceithre"},
		{Number: 3, Value: "Trí"},
		{Number: 2, Value: "Dó"},
		{Number: 1, Value: "Aon"},
		{Number: 0, Value: "Náid"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Céad"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "Céad", Suffix: "ú", Masculine: "Céad", Feminine: "Chéad", Neuter: "Céad"},
		{Number: 2, Word: "Dara", Suffix: "ú", Masculine: "Dara", Feminine: "Dara", Neuter: "Dara"},
		{Number: 3, Word: "Tríú", Suffix: "ú", Masculine: "Tríú", Feminine: "Tríú", Neuter: "Tríú"},
		{Number: 4, Word: "Ceathrú", Suffix: "ú", Masculine: "Ceathrú", Feminine: "Ceathrú", Neuter: "Ceathrú"},
		{Number: 5, Word: "Cúigiú", Suffix: "ú", Masculine: "Cúigiú", Feminine: "Cúigiú", Neuter: "Cúigiú"},
		{Number: 6, Word: "Séú", Suffix: "ú", Masculine: "Séú", Feminine: "Séú", Neuter: "Séú"},
		{Number: 7, Word: "Seachtú", Suffix: "ú", Masculine: "Seachtú", Feminine: "Seachtú", Neuter: "Seachtú"},
		{Number: 8, Word: "Ochtú", Suffix: "ú", Masculine: "Ochtú", Feminine: "Ochtú", Neuter: "Ochtú"},
		{Number: 9, Word: "Naoú", Suffix: "ú", Masculine: "Naoú", Feminine: "Naoú", Neuter: "Naoú"},
		{Number: 10, Word: "Deichiú", Suffix: "ú", Masculine: "Deichiú", Feminine: "Deichiú", Neuter: "Deichiú"},
		{Number: 11, Word: "Aon déagú", Suffix: "ú", Masculine: "Aon déagú", Feminine: "Aon déagú", Neuter: "Aon déagú"},
		{Number: 12, Word: "Dó dhéagú", Suffix: "ú", Masculine: "Dó dhéagú", Feminine: "Dó dhéagú", Neuter: "Dó dhéagú"},
		{Number: 13, Word: "Trí déagú", Suffix: "ú", Masculine: "Trí déagú", Feminine: "Trí déagú", Neuter: "Trí déagú"},
		{Number: 14, Word: "Ceithre déagú", Suffix: "ú", Masculine: "Ceithre déagú", Feminine: "Ceithre déagú", Neuter: "Ceithre déagú"},
		{Number: 15, Word: "Cúig déagú", Suffix: "ú", Masculine: "Cúig déagú", Feminine: "Cúig déagú", Neuter: "Cúig déagú"},
		{Number: 16, Word: "Sé déagú", Suffix: "ú", Masculine: "Sé déagú", Feminine: "Sé déagú", Neuter: "Sé déagú"},
		{Number: 17, Word: "Seacht déagú", Suffix: "ú", Masculine: "Seacht déagú", Feminine: "Seacht déagú", Neuter: "Seacht déagú"},
		{Number: 18, Word: "Ocht déagú", Suffix: "ú", Masculine: "Ocht déagú", Feminine: "Ocht déagú", Neuter: "Ocht déagú"},
		{Number: 19, Word: "Naoi déagú", Suffix: "ú", Masculine: "Naoi déagú", Feminine: "Naoi déagú", Neuter: "Naoi déagú"},
		{Number: 20, Word: "Fichiú", Suffix: "ú", Masculine: "Fichiú", Feminine: "Fichiú", Neuter: "Fichiú"},
		{Number: 21, Word: "Fiche a haon", Suffix: "ú", Masculine: "Fiche a haon", Feminine: "Fiche a haon", Neuter: "Fiche a haon"},
		{Number: 30, Word: "Tríochadú", Suffix: "ú", Masculine: "Tríochadú", Feminine: "Tríochadú", Neuter: "Tríochadú"},
		{Number: 40, Word: "Daicheadú", Suffix: "ú", Masculine: "Daicheadú", Feminine: "Daicheadú", Neuter: "Daicheadú"},
		{Number: 50, Word: "Caogadú", Suffix: "ú", Masculine: "Caogadú", Feminine: "Caogadú", Neuter: "Caogadú"},
		{Number: 60, Word: "Seascadú", Suffix: "ú", Masculine: "Seascadú", Feminine: "Seascadú", Neuter: "Seascadú"},
		{Number: 70, Word: "Seachtódú", Suffix: "ú", Masculine: "Seachtódú", Feminine: "Seachtódú", Neuter: "Seachtódú"},
		{Number: 80, Word: "Ochtódú", Suffix: "ú", Masculine: "Ochtódú", Feminine: "Ochtódú", Neuter: "Ochtódú"},
		{Number: 90, Word: "Nóchadú", Suffix: "ú", Masculine: "Nóchadú", Feminine: "Nóchadú", Neuter: "Nóchadú"},
		{Number: 100, Word: "Céadú", Suffix: "ú", Masculine: "Céadú", Feminine: "Chéadú", Neuter: "Céadú"},
		{Number: 1000, Word: "Míliú", Suffix: "ú", Masculine: "Míliú", Feminine: "Míliú", Neuter: "Míliú"},
	},
	LocaleFormatter: &IrishFormatter{},
}

// IrishFormatter handles Irish Gaelic (ga-IE) formatting
type IrishFormatter struct{}

func (f *IrishFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *IrishFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *IrishFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *IrishFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *IrishFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *IrishFormatter) ChopDecimal(num decimal.Decimal, precision int) decimal.Decimal {
	return num.Truncate(int32(precision))
}

func (f *IrishFormatter) FormatDecimalNumber(amount float64) string {
	return FormatEuropeanDecimal(amount)
}
func (f *IrishFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with European conventions (period separators, comma decimal, prefix symbol)
	return FormatEuropeanCurrency(amount, currencySymbol)
}
