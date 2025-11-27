package locale

import "github.com/shopspring/decimal"

// EELocale is a NumI18NLocale configured for Estonia (et-EE)
var EELocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Euro",
		Plural:   "Eurot",
		Singular: "Euro",
		Symbol:   "€",
		FractionUnit: FractionUnit{
			Name:     "Sent",
			Plural:   "Senti",
			Singular: "Sent",
			Symbol:   "¢",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Estonia",
		Currency:       "EUR",
		ISO3166Alpha2:  "EE",
		ISO3166Alpha3:  "EST",
		ISO3166Numeric: "233",
		Locale:         "et-EE",
		Timezone:       []string{"Europe/Tallinn"},
		Language:       "et",
	},
	Texts: Texts{
		And:   "Ja",
		Minus: "Miinus",
		Only:  "Ainult",
		Point: "Koma",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Biljard"},
		{Number: 1000000000000, Value: "Biljon"},
		{Number: 1000000000, Value: "Miljard"},
		{Number: 1000000, Value: "Miljon"},
		{Number: 1000, Value: "Tuhat"},
		{Number: 100, Value: "Sada"},
		{Number: 90, Value: "Üheksakümmend"},
		{Number: 80, Value: "Kaheksakümmend"},
		{Number: 70, Value: "Seitsekümmend"},
		{Number: 60, Value: "Kuuskümmend"},
		{Number: 50, Value: "Viiskümmend"},
		{Number: 40, Value: "Nelikümmend"},
		{Number: 30, Value: "Kolmkümmend"},
		{Number: 20, Value: "Kakskümmend"},
		{Number: 19, Value: "Üheksateist"},
		{Number: 18, Value: "Kaheksateist"},
		{Number: 17, Value: "Seitseteist"},
		{Number: 16, Value: "Kuusteist"},
		{Number: 15, Value: "Viisteist"},
		{Number: 14, Value: "Neliteist"},
		{Number: 13, Value: "Kolmteist"},
		{Number: 12, Value: "Kaksteist"},
		{Number: 11, Value: "Üksteist"},
		{Number: 10, Value: "Kümme"},
		{Number: 9, Value: "Üheksa"},
		{Number: 8, Value: "Kaheksa"},
		{Number: 7, Value: "Seitse"},
		{Number: 6, Value: "Kuus"},
		{Number: 5, Value: "Viis"},
		{Number: 4, Value: "Neli"},
		{Number: 3, Value: "Kolm"},
		{Number: 2, Value: "Kaks"},
		{Number: 1, Value: "Üks"},
		{Number: 0, Value: "Null"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{100, "Üks Sada"},
		{1000, "Üks Tuhat"},
		{1000000, "Üks Miljon"},
		{1000000000, "Üks Miljard"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "Esimene", Suffix: ".", Masculine: "Esimene", Feminine: "Esimene", Neuter: "Esimene"},
		{Number: 2, Word: "Teine", Suffix: ".", Masculine: "Teine", Feminine: "Teine", Neuter: "Teine"},
		{Number: 3, Word: "Kolmas", Suffix: ".", Masculine: "Kolmas", Feminine: "Kolmas", Neuter: "Kolmas"},
		{Number: 4, Word: "Neljas", Suffix: ".", Masculine: "Neljas", Feminine: "Neljas", Neuter: "Neljas"},
		{Number: 5, Word: "Viies", Suffix: ".", Masculine: "Viies", Feminine: "Viies", Neuter: "Viies"},
		{Number: 6, Word: "Kuues", Suffix: ".", Masculine: "Kuues", Feminine: "Kuues", Neuter: "Kuues"},
		{Number: 7, Word: "Seitsmes", Suffix: ".", Masculine: "Seitsmes", Feminine: "Seitsmes", Neuter: "Seitsmes"},
		{Number: 8, Word: "Kaheksas", Suffix: ".", Masculine: "Kaheksas", Feminine: "Kaheksas", Neuter: "Kaheksas"},
		{Number: 9, Word: "Üheksas", Suffix: ".", Masculine: "Üheksas", Feminine: "Üheksas", Neuter: "Üheksas"},
		{Number: 10, Word: "Kümnes", Suffix: ".", Masculine: "Kümnes", Feminine: "Kümnes", Neuter: "Kümnes"},
		{Number: 11, Word: "Üheteistkümnes", Suffix: ".", Masculine: "Üheteistkümnes", Feminine: "Üheteistkümnes", Neuter: "Üheteistkümnes"},
		{Number: 12, Word: "Kaheteistkümnes", Suffix: ".", Masculine: "Kaheteistkümnes", Feminine: "Kaheteistkümnes", Neuter: "Kaheteistkümnes"},
		{Number: 13, Word: "Kolmeteistkümnes", Suffix: ".", Masculine: "Kolmeteistkümnes", Feminine: "Kolmeteistkümnes", Neuter: "Kolmeteistkümnes"},
		{Number: 14, Word: "Neliteistkümnes", Suffix: ".", Masculine: "Neliteistkümnes", Feminine: "Neliteistkümnes", Neuter: "Neliteistkümnes"},
		{Number: 15, Word: "Viieteistkümnes", Suffix: ".", Masculine: "Viieteistkümnes", Feminine: "Viieteistkümnes", Neuter: "Viieteistkümnes"},
		{Number: 16, Word: "Kuueteistkümnes", Suffix: ".", Masculine: "Kuueteistkümnes", Feminine: "Kuueteistkümnes", Neuter: "Kuueteistkümnes"},
		{Number: 17, Word: "Seitsmeteistkümnes", Suffix: ".", Masculine: "Seitsmeteistkümnes", Feminine: "Seitsmeteistkümnes", Neuter: "Seitsmeteistkümnes"},
		{Number: 18, Word: "Kaheksateistkümnes", Suffix: ".", Masculine: "Kaheksateistkümnes", Feminine: "Kaheksateistkümnes", Neuter: "Kaheksateistkümnes"},
		{Number: 19, Word: "Üheksateistkümnes", Suffix: ".", Masculine: "Üheksateistkümnes", Feminine: "Üheksateistkümnes", Neuter: "Üheksateistkümnes"},
		{Number: 20, Word: "Kahekümnes", Suffix: ".", Masculine: "Kahekümnes", Feminine: "Kahekümnes", Neuter: "Kahekümnes"},
		{Number: 21, Word: "Kahekümne esimene", Suffix: ".", Masculine: "Kahekümne esimene", Feminine: "Kahekümne esimene", Neuter: "Kahekümne esimene"},
		{Number: 30, Word: "Kolmekümnes", Suffix: ".", Masculine: "Kolmekümnes", Feminine: "Kolmekümnes", Neuter: "Kolmekümnes"},
		{Number: 40, Word: "Nelikümnes", Suffix: ".", Masculine: "Nelikümnes", Feminine: "Nelikümnes", Neuter: "Nelikümnes"},
		{Number: 50, Word: "Viiekümnes", Suffix: ".", Masculine: "Viiekümnes", Feminine: "Viiekümnes", Neuter: "Viiekümnes"},
		{Number: 60, Word: "Kuuekümnes", Suffix: ".", Masculine: "Kuuekümnes", Feminine: "Kuuekümnes", Neuter: "Kuuekümnes"},
		{Number: 70, Word: "Seitsmekümnes", Suffix: ".", Masculine: "Seitsmekümnes", Feminine: "Seitsmekümnes", Neuter: "Seitsmekümnes"},
		{Number: 80, Word: "Kaheksakümnes", Suffix: ".", Masculine: "Kaheksakümnes", Feminine: "Kaheksakümnes", Neuter: "Kaheksakümnes"},
		{Number: 90, Word: "Üheksakümnes", Suffix: ".", Masculine: "Üheksakümnes", Feminine: "Üheksakümnes", Neuter: "Üheksakümnes"},
		{Number: 100, Word: "Sajas", Suffix: ".", Masculine: "Sajas", Feminine: "Sajas", Neuter: "Sajas"},
		{Number: 1000, Word: "Tuhandes", Suffix: ".", Masculine: "Tuhandes", Feminine: "Tuhandes", Neuter: "Tuhandes"},
	},
	LocaleFormatter: &EstonianFormatter{},
}

// EstonianFormatter handles Estonian formatting
type EstonianFormatter struct{}

func (f *EstonianFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *EstonianFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *EstonianFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *EstonianFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *EstonianFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *EstonianFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 0
	}
	return amount.Truncate(int32(precision))
}
