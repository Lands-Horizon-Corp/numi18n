package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// AFZALocale is a NumI18NLocale configured for Afrikaans (South Africa) - af-ZA
var AFZALocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "Rand",
		Plural:   "Rande",
		Singular: "Rand",
		Symbol:   "R",
		FractionUnit: pkg.FractionUnit{
			Name:     "Sent",
			Plural:   "Sente",
			Singular: "Sent",
			Symbol:   "c",
		},
	},
	Texts: pkg.Texts{
		And:   "En",
		Minus: "Min",
		Only:  "Slegs",
		Point: "Komma",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "Kwadrieljoen"},
		{Number: 1000000000000, Value: "Triljoen"},
		{Number: 1000000000, Value: "Miljard"},
		{Number: 1000000, Value: "Miljoen"},
		{Number: 1000, Value: "Duizend"},
		{Number: 100, Value: "Honderd"},
		{Number: 90, Value: "Negentig"},
		{Number: 80, Value: "Tagtig"},
		{Number: 70, Value: "Sewentig"},
		{Number: 60, Value: "Sestig"},
		{Number: 50, Value: "Vyftig"},
		{Number: 40, Value: "Veertig"},
		{Number: 30, Value: "Dertig"},
		{Number: 20, Value: "Twintig"},
		{Number: 19, Value: "Negentien"},
		{Number: 18, Value: "Agtien"},
		{Number: 17, Value: "Sewentien"},
		{Number: 16, Value: "Sestien"},
		{Number: 15, Value: "Vyftien"},
		{Number: 14, Value: "Veertien"},
		{Number: 13, Value: "Dertien"},
		{Number: 12, Value: "Twaalf"},
		{Number: 11, Value: "Elf"},
		{Number: 10, Value: "Tien"},
		{Number: 9, Value: "Nege"},
		{Number: 8, Value: "Agt"},
		{Number: 7, Value: "Sewe"},
		{Number: 6, Value: "Ses"},
		{Number: 5, Value: "Vyf"},
		{Number: 4, Value: "Vier"},
		{Number: 3, Value: "Drie"},
		{Number: 2, Value: "Twee"},
		{Number: 1, Value: "Een"},
		{Number: 0, Value: "Nul"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Honderd"},
	},
}
