package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// ISLocale is a NumI18NLocale configured for Iceland (is-IS)
var ISLocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "Króna",
		Plural:   "Krónur",
		Singular: "Króna",
		Symbol:   "kr",
		FractionUnit: pkg.FractionUnit{
			Name:     "Eyrir",
			Plural:   "Aurar",
			Singular: "Eyrir",
			Symbol:   "a",
		},
	},
	Texts: pkg.Texts{
		And:   "Og",
		Minus: "Mínus",
		Only:  "Aðeins",
		Point: "Komma",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "Ein trilljón trilljóna"},
		{Number: 1000000000000, Value: "Ein trilljón"},
		{Number: 1000000000, Value: "Einn milljarður"},
		{Number: 1000000, Value: "Ein milljón"},
		{Number: 1000, Value: "Eitt þúsund"},
		{Number: 100, Value: "Eitt hundrað"},
		{Number: 90, Value: "Níutíu"},
		{Number: 80, Value: "Áttatíu"},
		{Number: 70, Value: "Sjötíu"},
		{Number: 60, Value: "Sextíu"},
		{Number: 50, Value: "Fimmtíu"},
		{Number: 40, Value: "Fjörutíu"},
		{Number: 30, Value: "Þrjátíu"},
		{Number: 20, Value: "Tuttugu"},
		{Number: 19, Value: "Nítján"},
		{Number: 18, Value: "Átján"},
		{Number: 17, Value: "Sautján"},
		{Number: 16, Value: "Sextán"},
		{Number: 15, Value: "Fimmtán"},
		{Number: 14, Value: "Fjórtán"},
		{Number: 13, Value: "Þrettán"},
		{Number: 12, Value: "Tólf"},
		{Number: 11, Value: "Ellefu"},
		{Number: 10, Value: "Tíu"},
		{Number: 9, Value: "Níu"},
		{Number: 8, Value: "Átta"},
		{Number: 7, Value: "Sjö"},
		{Number: 6, Value: "Sex"},
		{Number: 5, Value: "Fimm"},
		{Number: 4, Value: "Fjögur"},
		{Number: 3, Value: "Þrjú"},
		{Number: 2, Value: "Tvö"},
		{Number: 1, Value: "Eitt"},
		{Number: 0, Value: "Núll"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Eitt hundrað"},
	},
}
