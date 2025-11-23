package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// CHITLocale is a NumI18NLocale configured for Switzerland (it-CH)
var CHITLocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "Franco svizzero",
		Plural:   "Franchi svizzeri",
		Singular: "Franco svizzero",
		Symbol:   "CHF",
		FractionUnit: pkg.FractionUnit{
			Name:     "Centesimo",
			Plural:   "Centesimi",
			Singular: "Centesimo",
			Symbol:   "¢",
		},
	},
	Texts: pkg.Texts{
		And:   "E",
		Minus: "Meno",
		Only:  "Solo",
		Point: "Virgola",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "Un biliardo"},
		{Number: 1000000000000, Value: "Un bilione"},
		{Number: 1000000000, Value: "Un miliardo"},
		{Number: 1000000, Value: "Un milione"},
		{Number: 1000, Value: "Mille"},
		{Number: 100, Value: "Cento"},
		{Number: 90, Value: "Novanta"},
		{Number: 80, Value: "Ottanta"},
		{Number: 70, Value: "Settanta"},
		{Number: 60, Value: "Sessanta"},
		{Number: 50, Value: "Cinquanta"},
		{Number: 40, Value: "Quaranta"},
		{Number: 30, Value: "Trenta"},
		{Number: 20, Value: "Venti"},
		{Number: 19, Value: "Diciannove"},
		{Number: 18, Value: "Diciotto"},
		{Number: 17, Value: "Diciassette"},
		{Number: 16, Value: "Sedici"},
		{Number: 15, Value: "Quindici"},
		{Number: 14, Value: "Quattordici"},
		{Number: 13, Value: "Tredici"},
		{Number: 12, Value: "Dodici"},
		{Number: 11, Value: "Undici"},
		{Number: 10, Value: "Dieci"},
		{Number: 9, Value: "Nove"},
		{Number: 8, Value: "Otto"},
		{Number: 7, Value: "Sette"},
		{Number: 6, Value: "Sei"},
		{Number: 5, Value: "Cinque"},
		{Number: 4, Value: "Quattro"},
		{Number: 3, Value: "Tre"},
		{Number: 2, Value: "Due"},
		{Number: 1, Value: "Uno"},
		{Number: 0, Value: "Zero"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Cento"},
	},
}
