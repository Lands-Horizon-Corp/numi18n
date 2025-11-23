package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// TRKULocale is a NumI18NLocale configured for Turkey (ku-TR)
var TRKULocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "Lîre",
		Plural:   "Lîre",
		Singular: "Lîre",
		Symbol:   "₺",
		FractionUnit: pkg.FractionUnit{
			Name:     "Kuruş",
			Plural:   "Kuruş",
			Singular: "Kuruş",
			Symbol:   "kr",
		},
	},
	Texts: pkg.Texts{
		And:   "Û",
		Minus: "Kêm",
		Only:  "Tenê",
		Point: "Xal",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "Yek kuadrîlyon"},
		{Number: 1000000000000, Value: "Yek trîlyon"},
		{Number: 1000000000, Value: "Yek mîlyar"},
		{Number: 1000000, Value: "Yek mîlyon"},
		{Number: 1000, Value: "Yek hezar"},
		{Number: 100, Value: "Yek sed"},
		{Number: 90, Value: "Nod"},
		{Number: 80, Value: "Heştê"},
		{Number: 70, Value: "Heftê"},
		{Number: 60, Value: "Şêst"},
		{Number: 50, Value: "Pêncî"},
		{Number: 40, Value: "Çil"},
		{Number: 30, Value: "Sî"},
		{Number: 20, Value: "Bîst"},
		{Number: 19, Value: "Nozdeh"},
		{Number: 18, Value: "Hijdeh"},
		{Number: 17, Value: "Hevdeh"},
		{Number: 16, Value: "Şazdeh"},
		{Number: 15, Value: "Pazdeh"},
		{Number: 14, Value: "Çardeh"},
		{Number: 13, Value: "Sêzdeh"},
		{Number: 12, Value: "Dwanzdeh"},
		{Number: 11, Value: "Yazdeh"},
		{Number: 10, Value: "Deh"},
		{Number: 9, Value: "Neh"},
		{Number: 8, Value: "Heşt"},
		{Number: 7, Value: "Heft"},
		{Number: 6, Value: "Şeş"},
		{Number: 5, Value: "Pênc"},
		{Number: 4, Value: "Çar"},
		{Number: 3, Value: "Sê"},
		{Number: 2, Value: "Du"},
		{Number: 1, Value: "Yek"},
		{Number: 0, Value: "Sifir"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Yek sed"},
	},
}
