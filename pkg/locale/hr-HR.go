package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// HRLocale is a NumI18NLocale configured for Croatia (hr-HR)
var HRLocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "Euro",
		Plural:   "Eura",
		Singular: "Euro",
		Symbol:   "€",
		FractionUnit: pkg.FractionUnit{
			Name:     "Cent",
			Plural:   "Centa",
			Singular: "Cent",
			Symbol:   "¢",
		},
	},
	Texts: pkg.Texts{
		And:   "I",
		Minus: "Minus",
		Only:  "Samo",
		Point: "Zarez",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
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
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Sto"},
	},
}
