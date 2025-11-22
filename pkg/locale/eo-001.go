package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// EOLocale is a NumI18NLocale configured for Esperanto (eo-001)
var EOLocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "Euro",
		Plural:   "Eŭroj",
		Singular: "Eŭro",
		Symbol:   "€",
		FractionUnit: pkg.FractionUnit{
			Name:     "Cento",
			Plural:   "Centoj",
			Singular: "Cento",
			Symbol:   "¢",
		},
	},
	Texts: pkg.Texts{
		And:   "Kaj",
		Minus: "Minus",
		Only:  "Nur",
		Point: "Punkto",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "Kvadriliono"},
		{Number: 1000000000000, Value: "Triliono"},
		{Number: 1000000000, Value: "Bilio"},
		{Number: 1000000, Value: "Miliono"},
		{Number: 1000, Value: "Mil"},
		{Number: 100, Value: "Cent"},
		{Number: 90, Value: "Kudek"},
		{Number: 80, Value: "Okdek"},
		{Number: 70, Value: "Sepdek"},
		{Number: 60, Value: "Sesdek"},
		{Number: 50, Value: "Kvindek"},
		{Number: 40, Value: "Kvardek"},
		{Number: 30, Value: "Tridek"},
		{Number: 20, Value: "Dudek"},
		{Number: 19, Value: "Dek na naŭ"},
		{Number: 18, Value: "Dek ok"},
		{Number: 17, Value: "Dek sep"},
		{Number: 16, Value: "Dek ses"},
		{Number: 15, Value: "Dek kvin"},
		{Number: 14, Value: "Dek kvar"},
		{Number: 13, Value: "Dek tri"},
		{Number: 12, Value: "Dek du"},
		{Number: 11, Value: "Dek unu"},
		{Number: 10, Value: "Dek"},
		{Number: 9, Value: "Nau"},
		{Number: 8, Value: "Ok"},
		{Number: 7, Value: "Sep"},
		{Number: 6, Value: "Ses"},
		{Number: 5, Value: "Kvin"},
		{Number: 4, Value: "Kvar"},
		{Number: 3, Value: "Tri"},
		{Number: 2, Value: "Du"},
		{Number: 1, Value: "Unu"},
		{Number: 0, Value: "Nulo"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Unu Cento"},
	},
}
