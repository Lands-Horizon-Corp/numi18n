package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// FOLocale is a NumI18NLocale configured for Faroe Islands (fo-FO)
var FOLocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "Króna",
		Plural:   "Krónur",
		Singular: "Króna",
		Symbol:   "kr",
		FractionUnit: pkg.FractionUnit{
			Name:     "Oyra",
			Plural:   "Oyrur",
			Singular: "Oyra",
			Symbol:   "ø",
		},
	},
	Texts: pkg.Texts{
		And:   "Og",
		Minus: "Mínus",
		Only:  "Bara",
		Point: "Komma",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "Eitt quadrillion"},
		{Number: 1000000000000, Value: "Eitt trillion"},
		{Number: 1000000000, Value: "Ein milliard"},
		{Number: 1000000, Value: "Ein million"},
		{Number: 1000, Value: "Eitt túsund"},
		{Number: 100, Value: "Eitt hundrað"},
		{Number: 90, Value: "Níti"},
		{Number: 80, Value: "Áttati"},
		{Number: 70, Value: "Sjeyti"},
		{Number: 60, Value: "Seksti"},
		{Number: 50, Value: "Hálvtríss"},
		{Number: 40, Value: "Fyrti"},
		{Number: 30, Value: "Tríati"},
		{Number: 20, Value: "Tuttugu"},
		{Number: 19, Value: "Nítjan"},
		{Number: 18, Value: "Átjan"},
		{Number: 17, Value: "Seytan"},
		{Number: 16, Value: "Sekstan"},
		{Number: 15, Value: "Fimtan"},
		{Number: 14, Value: "Fjórtan"},
		{Number: 13, Value: "Trettan"},
		{Number: 12, Value: "Tólv"},
		{Number: 11, Value: "Ellivu"},
		{Number: 10, Value: "Tíggju"},
		{Number: 9, Value: "Níggju"},
		{Number: 8, Value: "Átta"},
		{Number: 7, Value: "Sjey"},
		{Number: 6, Value: "Seks"},
		{Number: 5, Value: "Fimm"},
		{Number: 4, Value: "Fýra"},
		{Number: 3, Value: "Trý"},
		{Number: 2, Value: "Tvey"},
		{Number: 1, Value: "Eitt"},
		{Number: 0, Value: "Null"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Eitt hundrað"},
	},
}
