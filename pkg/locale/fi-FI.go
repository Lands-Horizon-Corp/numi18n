package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// FILocale is a NumI18NLocale configured for Finland (fi-FI)
var FILocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "Euro",
		Plural:   "Euroa",
		Singular: "Euro",
		Symbol:   "€",
		FractionUnit: pkg.FractionUnit{
			Name:     "Sentti",
			Plural:   "Senttiä",
			Singular: "Sentti",
			Symbol:   "¢",
		},
	},
	Texts: pkg.Texts{
		And:   "Ja",
		Minus: "Miinus",
		Only:  "Vain",
		Point: "Pilkku",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "Kvadriljoona"},
		{Number: 1000000000000, Value: "Triljoona"},
		{Number: 1000000000, Value: "Miljardi"},
		{Number: 1000000, Value: "Miljoona"},
		{Number: 1000, Value: "Tuhat"},
		{Number: 100, Value: "Sata"},
		{Number: 90, Value: "Yhdeksänkymmentä"},
		{Number: 80, Value: "Kahdeksankymmentä"},
		{Number: 70, Value: "Seitsemänkymmentä"},
		{Number: 60, Value: "Kuusikymmentä"},
		{Number: 50, Value: "Viisikymmentä"},
		{Number: 40, Value: "Neljäkymmentä"},
		{Number: 30, Value: "Kolmekymmentä"},
		{Number: 20, Value: "Kaksikymmentä"},
		{Number: 19, Value: "Yhdeksäntoista"},
		{Number: 18, Value: "Kahdeksantoista"},
		{Number: 17, Value: "Seitsemäntoista"},
		{Number: 16, Value: "Kuusitoista"},
		{Number: 15, Value: "Viisitoista"},
		{Number: 14, Value: "Neljätoista"},
		{Number: 13, Value: "Kolmetoista"},
		{Number: 12, Value: "Kaksitoista"},
		{Number: 11, Value: "Yksitoista"},
		{Number: 10, Value: "Kymmenen"},
		{Number: 9, Value: "Yhdeksän"},
		{Number: 8, Value: "Kahdeksan"},
		{Number: 7, Value: "Seitsemän"},
		{Number: 6, Value: "Kuusi"},
		{Number: 5, Value: "Viisi"},
		{Number: 4, Value: "Neljä"},
		{Number: 3, Value: "Kolme"},
		{Number: 2, Value: "Kaksi"},
		{Number: 1, Value: "Yksi"},
		{Number: 0, Value: "Nolla"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Sata"},
	},
}
