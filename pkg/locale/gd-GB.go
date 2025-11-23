package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// GBGDLocale is a NumI18NLocale configured for Great Britain (gd-GB)
var GBGDLocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "Punnd Sasannach",
		Plural:   "Puinnd Sasannach",
		Singular: "Punnd Sasannach",
		Symbol:   "£",
		FractionUnit: pkg.FractionUnit{
			Name:     "Sgillinn",
			Plural:   "Sgillinnean",
			Singular: "Sgillinn",
			Symbol:   "p",
		},
	},
	Texts: pkg.Texts{
		And:   "Agus",
		Minus: "As aonais",
		Only:  "A-mhàin",
		Point: "Puing",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "Mìle billean billean"},
		{Number: 1000000000000, Value: "Billean billean"},
		{Number: 1000000000, Value: "Billean"},
		{Number: 1000000, Value: "Millean"},
		{Number: 1000, Value: "Mìle"},
		{Number: 100, Value: "Ceud"},
		{Number: 90, Value: "Naochad"},
		{Number: 80, Value: "Ochdad"},
		{Number: 70, Value: "Seachdad"},
		{Number: 60, Value: "Seasgad"},
		{Number: 50, Value: "Leth-cheud"},
		{Number: 40, Value: "Dà fhichead"},
		{Number: 30, Value: "Trithead"},
		{Number: 20, Value: "Fichead"},
		{Number: 19, Value: "Naoi deug"},
		{Number: 18, Value: "Ochd deug"},
		{Number: 17, Value: "Seachd deug"},
		{Number: 16, Value: "Sia deug"},
		{Number: 15, Value: "Còig deug"},
		{Number: 14, Value: "Ceithir deug"},
		{Number: 13, Value: "Trì deug"},
		{Number: 12, Value: "Dà dheug"},
		{Number: 11, Value: "Aon deug"},
		{Number: 10, Value: "Deich"},
		{Number: 9, Value: "Naoi"},
		{Number: 8, Value: "Ochd"},
		{Number: 7, Value: "Seachd"},
		{Number: 6, Value: "Sia"},
		{Number: 5, Value: "Còig"},
		{Number: 4, Value: "Ceithir"},
		{Number: 3, Value: "Trì"},
		{Number: 2, Value: "Dà"},
		{Number: 1, Value: "Aon"},
		{Number: 0, Value: "Neoni"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Ceud"},
	},
}
