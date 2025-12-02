package locale

import (
	"strings"
	"github.com/shopspring/decimal"
)

// HRLocale is a NumI18NLocale configured for Croatia (hr-HR)
var HRLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Euro",
		Plural:   "Eura",
		Singular: "Euro",
		Symbol:   "€",
		FractionUnit: FractionUnit{
			Name:     "Cent",
			Plural:   "Centa",
			Singular: "Cent",
			Symbol:   "¢",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Croatia",
		Currency:       "HRK",
		ISO3166Alpha2:  "HR",
		ISO3166Alpha3:  "HRV",
		ISO3166Numeric: "191",
		Locale:         "hr-HR",
		Timezone:       []string{"Europe/Zagreb"},
		Language:       "hr",
	},
	Texts: Texts{
		And:   "I",
		Minus: "Minus",
		Only:  "Samo",
		Point: "Zarez",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Jedan kvadrilijun"},
		{Number: 1000000000000, Value: "Jedan trilijun"},
		{Number: 1000000000, Value: "Jedan milijarda"},
		{Number: 1000000, Value: "Jedan milijun"},
		{Number: 1000, Value: "Jedna tisuća"},
		{Number: 100, Value: "Sto"},
		{Number: 90, Value: "Devedeset"},
		{Number: 80, Value: "Osamdeset"},
		{Number: 70, Value: "Sedamdeset"},
		{Number: 60, Value: "Šezdeset"},
		{Number: 50, Value: "Pedeset"},
		{Number: 40, Value: "Četrdeset"},
		{Number: 30, Value: "Trideset"},
		{Number: 20, Value: "Dvadeset"},
		{Number: 19, Value: "Devetnaest"},
		{Number: 18, Value: "Osamnaest"},
		{Number: 17, Value: "Sedamnaest"},
		{Number: 16, Value: "Šesnaest"},
		{Number: 15, Value: "Petnaest"},
		{Number: 14, Value: "Četrnaest"},
		{Number: 13, Value: "Trinaest"},
		{Number: 12, Value: "Dvanaest"},
		{Number: 11, Value: "Jedanaest"},
		{Number: 10, Value: "Deset"},
		{Number: 9, Value: "Devet"},
		{Number: 8, Value: "Osam"},
		{Number: 7, Value: "Sedam"},
		{Number: 6, Value: "Šest"},
		{Number: 5, Value: "Pet"},
		{Number: 4, Value: "Četiri"},
		{Number: 3, Value: "Tri"},
		{Number: 2, Value: "Dva"},
		{Number: 1, Value: "Jedan"},
		{Number: 0, Value: "Nula"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Sto"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "Prvi", Suffix: ".", Masculine: "Prvi", Feminine: "Prva", Neuter: "Prvo"},
		{Number: 2, Word: "Drugi", Suffix: ".", Masculine: "Drugi", Feminine: "Druga", Neuter: "Drugo"},
		{Number: 3, Word: "Treći", Suffix: ".", Masculine: "Treći", Feminine: "Treća", Neuter: "Treće"},
		{Number: 4, Word: "Četvrti", Suffix: ".", Masculine: "Četvrti", Feminine: "Četvrta", Neuter: "Četvrto"},
		{Number: 5, Word: "Peti", Suffix: ".", Masculine: "Peti", Feminine: "Peta", Neuter: "Peto"},
		{Number: 6, Word: "Šesti", Suffix: ".", Masculine: "Šesti", Feminine: "Šesta", Neuter: "Šesto"},
		{Number: 7, Word: "Sedmi", Suffix: ".", Masculine: "Sedmi", Feminine: "Sedma", Neuter: "Sedmo"},
		{Number: 8, Word: "Osmi", Suffix: ".", Masculine: "Osmi", Feminine: "Osma", Neuter: "Osmo"},
		{Number: 9, Word: "Deveti", Suffix: ".", Masculine: "Deveti", Feminine: "Deveta", Neuter: "Deveto"},
		{Number: 10, Word: "Deseti", Suffix: ".", Masculine: "Deseti", Feminine: "Deseta", Neuter: "Deseto"},
		{Number: 11, Word: "Jedanaesti", Suffix: ".", Masculine: "Jedanaesti", Feminine: "Jedanaesta", Neuter: "Jedanaesto"},
		{Number: 12, Word: "Dvanaesti", Suffix: ".", Masculine: "Dvanaesti", Feminine: "Dvanaesta", Neuter: "Dvanaesto"},
		{Number: 13, Word: "Trinaesti", Suffix: ".", Masculine: "Trinaesti", Feminine: "Trinaesta", Neuter: "Trinaesto"},
		{Number: 14, Word: "Četrnaesti", Suffix: ".", Masculine: "Četrnaesti", Feminine: "Četrnaesta", Neuter: "Četrnaesto"},
		{Number: 15, Word: "Petnaesti", Suffix: ".", Masculine: "Petnaesti", Feminine: "Petnaesta", Neuter: "Petnaesto"},
		{Number: 16, Word: "Šesnaesti", Suffix: ".", Masculine: "Šesnaesti", Feminine: "Šesnaesta", Neuter: "Šesnaesto"},
		{Number: 17, Word: "Sedamnaesti", Suffix: ".", Masculine: "Sedamnaesti", Feminine: "Sedamnaesta", Neuter: "Sedamnaesto"},
		{Number: 18, Word: "Osamnaesti", Suffix: ".", Masculine: "Osamnaesti", Feminine: "Osamnaesta", Neuter: "Osamnaesto"},
		{Number: 19, Word: "Devetnaesti", Suffix: ".", Masculine: "Devetnaesti", Feminine: "Devetnaesta", Neuter: "Devetnaesto"},
		{Number: 20, Word: "Dvadeseti", Suffix: ".", Masculine: "Dvadeseti", Feminine: "Dvadeseta", Neuter: "Dvadeseto"},
		{Number: 21, Word: "Dvadeset prvi", Suffix: ".", Masculine: "Dvadeset prvi", Feminine: "Dvadeset prva", Neuter: "Dvadeset prvo"},
		{Number: 30, Word: "Trideseti", Suffix: ".", Masculine: "Trideseti", Feminine: "Trideseta", Neuter: "Trideseto"},
		{Number: 40, Word: "Četrdeseti", Suffix: ".", Masculine: "Četrdeseti", Feminine: "Četrdeseta", Neuter: "Četrdeseto"},
		{Number: 50, Word: "Pedeseti", Suffix: ".", Masculine: "Pedeseti", Feminine: "Pedeseta", Neuter: "Pedeseto"},
		{Number: 60, Word: "Šezdeseti", Suffix: ".", Masculine: "Šezdeseti", Feminine: "Šezdeseta", Neuter: "Šezdeseto"},
		{Number: 70, Word: "Sedamdeseti", Suffix: ".", Masculine: "Sedamdeseti", Feminine: "Sedamdeseta", Neuter: "Sedamdeseto"},
		{Number: 80, Word: "Osamdeseti", Suffix: ".", Masculine: "Osamdeseti", Feminine: "Osamdeseta", Neuter: "Osamdeseto"},
		{Number: 90, Word: "Devedeseti", Suffix: ".", Masculine: "Devedeseti", Feminine: "Devedeseta", Neuter: "Devedeseto"},
		{Number: 100, Word: "Stoti", Suffix: ".", Masculine: "Stoti", Feminine: "Stota", Neuter: "Stoto"},
		{Number: 1000, Word: "Tisućiti", Suffix: ".", Masculine: "Tisućiti", Feminine: "Tisućita", Neuter: "Tisućito"},
	},
	LocaleFormatter: &CroatianFormatter{},
}

// CroatianFormatter handles Croatian formatting
type CroatianFormatter struct{}

func (f *CroatianFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *CroatianFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *CroatianFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *CroatianFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *CroatianFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *CroatianFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	return amount.Truncate(int32(precision))
}


func (f *CroatianFormatter) FormatDecimalNumber(amount float64) string {
	return DefaultFormatDecimalNumber(amount, ",", ".")
}
func (f *CroatianFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	formattedNumber := f.FormatDecimalNumber(amount)
	
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}
	
	// Default currency placement for this locale (prefix with symbol)
	if strings.HasPrefix(formattedNumber, "-") {
		formattedNumber = strings.TrimPrefix(formattedNumber, "-")
		return "-" + currencySymbol + formattedNumber
	}
	
	return currencySymbol + formattedNumber
}
