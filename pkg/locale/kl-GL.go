package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// GLLocale is a NumI18NLocale configured for Greenland (kl-GL)
var GLLocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "Dansk krone",
		Plural:   "Danske kroner",
		Singular: "Dansk krone",
		Symbol:   "kr",
		FractionUnit: pkg.FractionUnit{
			Name:     "Øre",
			Plural:   "Øre",
			Singular: "Øre",
			Symbol:   "ø",
		},
	},
	Texts: pkg.Texts{
		And:   "Aamma",
		Minus: "Minus",
		Only:  "Kissaat",
		Point: "Tikkuartoq",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "Ataaseq quadrillion"},
		{Number: 1000000000000, Value: "Ataaseq trillion"},
		{Number: 1000000000, Value: "Ataaseq milliardi"},
		{Number: 1000000, Value: "Ataaseq million"},
		{Number: 1000, Value: "Ataaseq tusind"},
		{Number: 100, Value: "Ataaseq hundrede"},
		{Number: 90, Value: "Qulingiluat"},
		{Number: 80, Value: "Arfersanerit"},
		{Number: 70, Value: "Arfersanik-qulit"},
		{Number: 60, Value: "Arfersanik"},
		{Number: 50, Value: "Tallimat-qulit"},
		{Number: 40, Value: "Tallimat"},
		{Number: 30, Value: "Pingasut-qulit"},
		{Number: 20, Value: "Juullut"},
		{Number: 19, Value: "Arfineq-sisamat"},
		{Number: 18, Value: "Arfineq-pingasut"},
		{Number: 17, Value: "Arfineq-marluk"},
		{Number: 16, Value: "Arfineq-ataaseq"},
		{Number: 15, Value: "Aqqaneq-tallimat"},
		{Number: 14, Value: "Aqqaneq-sisamat"},
		{Number: 13, Value: "Aqqaneq-pingasut"},
		{Number: 12, Value: "Aqqaneq-marluk"},
		{Number: 11, Value: "Aqqaneq-ataaseq"},
		{Number: 10, Value: "Qulit"},
		{Number: 9, Value: "Qulingiluat"},
		{Number: 8, Value: "Arfineq-pingasut"},
		{Number: 7, Value: "Arfineq-marluk"},
		{Number: 6, Value: "Arfineq-ataaseq"},
		{Number: 5, Value: "Tallimat"},
		{Number: 4, Value: "Sisamat"},
		{Number: 3, Value: "Pingasut"},
		{Number: 2, Value: "Marluk"},
		{Number: 1, Value: "Ataaseq"},
		{Number: 0, Value: "Nul"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Ataaseq hundrede"},
	},
}
