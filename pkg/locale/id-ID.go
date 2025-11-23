package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// IDLocale is a NumI18NLocale configured for Indonesia (id-ID)
var IDLocale = pkg.NumI18NLocale{
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
		And:   "Dan",
		Minus: "Minus",
		Only:  "Hanya",
		Point: "Koma",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "Satu kuadrilun"},
		{Number: 1000000000000, Value: "Satu trilun"},
		{Number: 1000000000, Value: "Satu miliar"},
		{Number: 1000000, Value: "Satu juta"},
		{Number: 1000, Value: "Satu ribu"},
		{Number: 100, Value: "Seratus"},
		{Number: 90, Value: "Sembilan puluh"},
		{Number: 80, Value: "Delapan puluh"},
		{Number: 70, Value: "Tujuh puluh"},
		{Number: 60, Value: "Enam puluh"},
		{Number: 50, Value: "Lima puluh"},
		{Number: 40, Value: "Empat puluh"},
		{Number: 30, Value: "Tiga puluh"},
		{Number: 20, Value: "Dua puluh"},
		{Number: 19, Value: "Sembilan belas"},
		{Number: 18, Value: "Delapan belas"},
		{Number: 17, Value: "Tujuh belas"},
		{Number: 16, Value: "Enam belas"},
		{Number: 15, Value: "Lima belas"},
		{Number: 14, Value: "Empat belas"},
		{Number: 13, Value: "Tiga belas"},
		{Number: 12, Value: "Dua belas"},
		{Number: 11, Value: "Sebelas"},
		{Number: 10, Value: "Sepuluh"},
		{Number: 9, Value: "Sembilan"},
		{Number: 8, Value: "Delapan"},
		{Number: 7, Value: "Tujuh"},
		{Number: 6, Value: "Enam"},
		{Number: 5, Value: "Lima"},
		{Number: 4, Value: "Empat"},
		{Number: 3, Value: "Tiga"},
		{Number: 2, Value: "Dua"},
		{Number: 1, Value: "Satu"},
		{Number: 0, Value: "Nol"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Seratus"},
	},
}
