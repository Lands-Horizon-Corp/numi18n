package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// ESEULocale is a NumI18NLocale configured for Basque - Spain (eu-ES)
var ESEULocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "Euro",
		Plural:   "Euroak",
		Singular: "Euro",
		Symbol:   "€",
		FractionUnit: pkg.FractionUnit{
			Name:     "Zentimo",
			Plural:   "Zentimoak",
			Singular: "Zentimo",
			Symbol:   "¢",
		},
	},
	Texts: pkg.Texts{
		And:   "Eta",
		Minus: "Minus",
		Only:  "Soilik",
		Point: "Puntua",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "Kualdrilioi"},
		{Number: 1000000000000, Value: "Trilioi"},
		{Number: 1000000000, Value: "Mila milioi"},
		{Number: 1000000, Value: "Milioi"},
		{Number: 1000, Value: "Mila"},
		{Number: 100, Value: "Ehun"},
		{Number: 90, Value: "Laurogeita hamar"},
		{Number: 80, Value: "Laurogei"},
		{Number: 70, Value: "Hirurogeita hamar"},
		{Number: 60, Value: "Hirurogei"},
		{Number: 50, Value: "Berrogeita hamar"},
		{Number: 40, Value: "Berrogei"},
		{Number: 30, Value: "Hogeita hamar"},
		{Number: 20, Value: "Hogei"},
		{Number: 19, Value: "Hemeretzi"},
		{Number: 18, Value: "Hemezortzi"},
		{Number: 17, Value: "Hamazazpi"},
		{Number: 16, Value: "Hamasei"},
		{Number: 15, Value: "Hamabost"},
		{Number: 14, Value: "Hamalau"},
		{Number: 13, Value: "Hamahiru"},
		{Number: 12, Value: "Hamabi"},
		{Number: 11, Value: "Hamaika"},
		{Number: 10, Value: "Hamar"},
		{Number: 9, Value: "Bederatzi"},
		{Number: 8, Value: "Zortzi"},
		{Number: 7, Value: "Zazpi"},
		{Number: 6, Value: "Sei"},
		{Number: 5, Value: "Bost"},
		{Number: 4, Value: "Lau"},
		{Number: 3, Value: "Hiru"},
		{Number: 2, Value: "Bi"},
		{Number: 1, Value: "Bat"},
		{Number: 0, Value: "Zero"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Ehun"},
	},
}
