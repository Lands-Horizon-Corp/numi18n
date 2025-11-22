package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// CSCZLocale is a NumI18NLocale configured for Czech (Czech Republic) - cs-CZ
var CSCZLocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "Koruna",
		Plural:   "Koruny",
		Singular: "Koruna",
		Symbol:   "Kč",
		FractionUnit: pkg.FractionUnit{
			Name:     "Haléř",
			Plural:   "Haléře",
			Singular: "Haléř",
			Symbol:   "h",
		},
	},
	Texts: pkg.Texts{
		And:   "a",
		Minus: "mínus",
		Only:  "pouze",
		Point: "tečka",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "Kvadrilion"},
		{Number: 1000000000000, Value: "Bilion"},
		{Number: 1000000000, Value: "Miliarda"},
		{Number: 1000000, Value: "Milion"},
		{Number: 1000, Value: "Tisíc"},
		{Number: 100, Value: "Sto"},
		{Number: 90, Value: "Devadesát"},
		{Number: 80, Value: "Osmdesát"},
		{Number: 70, Value: "Sedmdesát"},
		{Number: 60, Value: "Šedesát"},
		{Number: 50, Value: "Padesát"},
		{Number: 40, Value: "Čtyřicet"},
		{Number: 30, Value: "Třicet"},
		{Number: 20, Value: "Dvacet"},
		{Number: 19, Value: "Devatenáct"},
		{Number: 18, Value: "Osmnáct"},
		{Number: 17, Value: "Sedmnáct"},
		{Number: 16, Value: "Šestnáct"},
		{Number: 15, Value: "Patnáct"},
		{Number: 14, Value: "Čtrnáct"},
		{Number: 13, Value: "Třináct"},
		{Number: 12, Value: "Dvanáct"},
		{Number: 11, Value: "Jedenáct"},
		{Number: 10, Value: "Deset"},
		{Number: 9, Value: "Devět"},
		{Number: 8, Value: "Osm"},
		{Number: 7, Value: "Sedm"},
		{Number: 6, Value: "Šest"},
		{Number: 5, Value: "Pět"},
		{Number: 4, Value: "Čtyři"},
		{Number: 3, Value: "Tři"},
		{Number: 2, Value: "Dva"},
		{Number: 1, Value: "Jedna"},
		{Number: 0, Value: "Nula"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Sto"},
	},
}
