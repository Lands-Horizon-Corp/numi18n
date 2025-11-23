package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// HULocale is a NumI18NLocale configured for Hungary (hu-HU)
var HULocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "Forint",
		Plural:   "Forint",
		Singular: "Forint",
		Symbol:   "Ft",
		FractionUnit: pkg.FractionUnit{
			Name:     "Fillér",
			Plural:   "Fillér",
			Singular: "Fillér",
			Symbol:   "f",
		},
	},
	Texts: pkg.Texts{
		And:   "És",
		Minus: "Mínusz",
		Only:  "Csak",
		Point: "Pont",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "Egy billiárd"},
		{Number: 1000000000000, Value: "Egy billió"},
		{Number: 1000000000, Value: "Egy milliárd"},
		{Number: 1000000, Value: "Egy millió"},
		{Number: 1000, Value: "Egy ezer"},
		{Number: 100, Value: "Egy száz"},
		{Number: 90, Value: "Kilencven"},
		{Number: 80, Value: "Nyolcvan"},
		{Number: 70, Value: "Hetven"},
		{Number: 60, Value: "Hatvan"},
		{Number: 50, Value: "Ötven"},
		{Number: 40, Value: "Negyven"},
		{Number: 30, Value: "Harminc"},
		{Number: 20, Value: "Húsz"},
		{Number: 19, Value: "Tizenkilenc"},
		{Number: 18, Value: "Tizennyolc"},
		{Number: 17, Value: "Tizenhét"},
		{Number: 16, Value: "Tizenhat"},
		{Number: 15, Value: "Tizenöt"},
		{Number: 14, Value: "Tizennégy"},
		{Number: 13, Value: "Tizenhárom"},
		{Number: 12, Value: "Tizenkettő"},
		{Number: 11, Value: "Tizenegy"},
		{Number: 10, Value: "Tíz"},
		{Number: 9, Value: "Kilenc"},
		{Number: 8, Value: "Nyolc"},
		{Number: 7, Value: "Hét"},
		{Number: 6, Value: "Hat"},
		{Number: 5, Value: "Öt"},
		{Number: 4, Value: "Négy"},
		{Number: 3, Value: "Három"},
		{Number: 2, Value: "Kettő"},
		{Number: 1, Value: "Egy"},
		{Number: 0, Value: "Nulla"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Egy száz"},
	},
}
