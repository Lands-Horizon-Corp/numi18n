package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// PYLocale is a NumI18NLocale configured for Paraguay (es-PY)
var PYLocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "Guaraní",
		Plural:   "Guaraníes",
		Singular: "Guaraní",
		Symbol:   "₲",
		FractionUnit: pkg.FractionUnit{
			Name:     "Céntimo",
			Plural:   "Céntimos",
			Singular: "Céntimo",
			Symbol:   "¢",
		},
	},
	Texts: pkg.Texts{
		And:   "Y",
		Minus: "Menos",
		Only:  "Solo",
		Point: "Punto",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "Cuatrillón"},
		{Number: 1000000000000, Value: "Trillón"},
		{Number: 1000000000, Value: "Mil millones"},
		{Number: 1000000, Value: "Millón"},
		{Number: 1000, Value: "Mil"},
		{Number: 100, Value: "Cien"},
		{Number: 90, Value: "Noventa"},
		{Number: 80, Value: "Ochenta"},
		{Number: 70, Value: "Setenta"},
		{Number: 60, Value: "Sesenta"},
		{Number: 50, Value: "Cincuenta"},
		{Number: 40, Value: "Cuarenta"},
		{Number: 30, Value: "Treinta"},
		{Number: 20, Value: "Veinte"},
		{Number: 19, Value: "Diecinueve"},
		{Number: 18, Value: "Dieciocho"},
		{Number: 17, Value: "Diecisiete"},
		{Number: 16, Value: "Dieciséis"},
		{Number: 15, Value: "Quince"},
		{Number: 14, Value: "Catorce"},
		{Number: 13, Value: "Trece"},
		{Number: 12, Value: "Doce"},
		{Number: 11, Value: "Once"},
		{Number: 10, Value: "Diez"},
		{Number: 9, Value: "Nueve"},
		{Number: 8, Value: "Ocho"},
		{Number: 7, Value: "Siete"},
		{Number: 6, Value: "Seis"},
		{Number: 5, Value: "Cinco"},
		{Number: 4, Value: "Cuatro"},
		{Number: 3, Value: "Tres"},
		{Number: 2, Value: "Dos"},
		{Number: 1, Value: "Uno"},
		{Number: 0, Value: "Cero"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Cien"},
	},
}
