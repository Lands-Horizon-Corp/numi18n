package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// AMLocale is a NumI18NLocale configured for Armenia (hy-AM)
var AMLocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "Դրամ",
		Plural:   "Դրամ",
		Singular: "Դրամ",
		Symbol:   "֏",
		FractionUnit: pkg.FractionUnit{
			Name:     "Լումա",
			Plural:   "Լումա",
			Singular: "Լումա",
			Symbol:   "լ",
		},
	},
	Texts: pkg.Texts{
		And:   "Եվ",
		Minus: "Մինուս",
		Only:  "Միայն",
		Point: "Կետ",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "Մեկ քվադրիլիոն"},
		{Number: 1000000000000, Value: "Մեկ տրիլիոն"},
		{Number: 1000000000, Value: "Մեկ միլիարդ"},
		{Number: 1000000, Value: "Մեկ միլիոն"},
		{Number: 1000, Value: "Մեկ հազար"},
		{Number: 100, Value: "Մեկ հարյուր"},
		{Number: 90, Value: "Իննսուն"},
		{Number: 80, Value: "Ութսուն"},
		{Number: 70, Value: "Յոթանասուն"},
		{Number: 60, Value: "Վաթսուն"},
		{Number: 50, Value: "Հիսուն"},
		{Number: 40, Value: "Քառասուն"},
		{Number: 30, Value: "Երեսուն"},
		{Number: 20, Value: "Քսան"},
		{Number: 19, Value: "Տասնինը"},
		{Number: 18, Value: "Տասնութ"},
		{Number: 17, Value: "Տասնյոթ"},
		{Number: 16, Value: "Տասնվեց"},
		{Number: 15, Value: "Տասնհինգ"},
		{Number: 14, Value: "Տասնչորս"},
		{Number: 13, Value: "Տասներեք"},
		{Number: 12, Value: "Տասներկու"},
		{Number: 11, Value: "Տասնմեկ"},
		{Number: 10, Value: "Տասը"},
		{Number: 9, Value: "Ինը"},
		{Number: 8, Value: "Ութ"},
		{Number: 7, Value: "Յոթ"},
		{Number: 6, Value: "Վեց"},
		{Number: 5, Value: "Հինգ"},
		{Number: 4, Value: "Չորս"},
		{Number: 3, Value: "Երեք"},
		{Number: 2, Value: "Երկու"},
		{Number: 1, Value: "Մեկ"},
		{Number: 0, Value: "Զրո"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Մեկ հարյուր"},
	},
}
