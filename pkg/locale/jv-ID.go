package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// IDJVLocale is a NumI18NLocale configured for Indonesia (jv-ID)
var IDJVLocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "Rupiah",
		Plural:   "Rupiah",
		Singular: "Rupiah",
		Symbol:   "Rp",
		FractionUnit: pkg.FractionUnit{
			Name:     "Sen",
			Plural:   "Sen",
			Singular: "Sen",
			Symbol:   "s",
		},
	},
	Texts: pkg.Texts{
		And:   "Lan",
		Minus: "Minus",
		Only:  "Mung",
		Point: "Titik",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "Siji kuadriliyun"},
		{Number: 1000000000000, Value: "Siji triliyun"},
		{Number: 1000000000, Value: "Siji milyar"},
		{Number: 1000000, Value: "Siji yuta"},
		{Number: 1000, Value: "Siji ewu"},
		{Number: 100, Value: "Satus"},
		{Number: 90, Value: "Sangang puluh"},
		{Number: 80, Value: "Wolung puluh"},
		{Number: 70, Value: "Pitung puluh"},
		{Number: 60, Value: "Sewidak"},
		{Number: 50, Value: "Seket"},
		{Number: 40, Value: "Patang puluh"},
		{Number: 30, Value: "Telung puluh"},
		{Number: 20, Value: "Rongpuluh"},
		{Number: 19, Value: "Sangalas"},
		{Number: 18, Value: "Wolulas"},
		{Number: 17, Value: "Pitulas"},
		{Number: 16, Value: "Nembelas"},
		{Number: 15, Value: "Limolas"},
		{Number: 14, Value: "Patbelas"},
		{Number: 13, Value: "Telulas"},
		{Number: 12, Value: "Rolas"},
		{Number: 11, Value: "Sewelas"},
		{Number: 10, Value: "Sepuluh"},
		{Number: 9, Value: "Sanga"},
		{Number: 8, Value: "Wolu"},
		{Number: 7, Value: "Pitu"},
		{Number: 6, Value: "Enem"},
		{Number: 5, Value: "Lima"},
		{Number: 4, Value: "Papat"},
		{Number: 3, Value: "Telu"},
		{Number: 2, Value: "Loro"},
		{Number: 1, Value: "Siji"},
		{Number: 0, Value: "Kosong"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Satus"},
	},
}
