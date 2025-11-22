package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// DECHLocale is a NumI18NLocale configured for German (Switzerland) - de-CH
var DECHLocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "Schweizer Franken",
		Plural:   "Schweizer Franken",
		Singular: "Schweizer Franken",
		Symbol:   "CHF",
		FractionUnit: pkg.FractionUnit{
			Name:     "Rappen",
			Plural:   "Rappen",
			Singular: "Rappen",
			Symbol:   "Rp",
		},
	},
	Texts: pkg.Texts{
		And:   "und",
		Minus: "minus",
		Only:  "nur",
		Point: "Komma",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "Billiarde"},
		{Number: 1000000000000, Value: "Billion"},
		{Number: 1000000000, Value: "Milliarde"},
		{Number: 1000000, Value: "Million"},
		{Number: 1000, Value: "Tausend"},
		{Number: 100, Value: "Hundert"},
		{Number: 90, Value: "Neunzig"},
		{Number: 80, Value: "Achtzig"},
		{Number: 70, Value: "Siebzig"},
		{Number: 60, Value: "Sechzig"},
		{Number: 50, Value: "Fünfzig"},
		{Number: 40, Value: "Vierzig"},
		{Number: 30, Value: "Dreißig"},
		{Number: 20, Value: "Zwanzig"},
		{Number: 19, Value: "Neunzehn"},
		{Number: 18, Value: "Achtzehn"},
		{Number: 17, Value: "Siebzehn"},
		{Number: 16, Value: "Sechzehn"},
		{Number: 15, Value: "Fünfzehn"},
		{Number: 14, Value: "Vierzehn"},
		{Number: 13, Value: "Dreizehn"},
		{Number: 12, Value: "Zwölf"},
		{Number: 11, Value: "Elf"},
		{Number: 10, Value: "Zehn"},
		{Number: 9, Value: "Neun"},
		{Number: 8, Value: "Acht"},
		{Number: 7, Value: "Sieben"},
		{Number: 6, Value: "Sechs"},
		{Number: 5, Value: "Fünf"},
		{Number: 4, Value: "Vier"},
		{Number: 3, Value: "Drei"},
		{Number: 2, Value: "Zwei"},
		{Number: 1, Value: "Eins"},
		{Number: 0, Value: "Null"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Einhundert"},
	},
}
