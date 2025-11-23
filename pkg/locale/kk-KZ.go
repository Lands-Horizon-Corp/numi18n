package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// KZLocale is a NumI18NLocale configured for Kazakhstan (kk-KZ)
var KZLocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "Теңге",
		Plural:   "Теңге",
		Singular: "Теңге",
		Symbol:   "₸",
		FractionUnit: pkg.FractionUnit{
			Name:     "Тиын",
			Plural:   "Тиын",
			Singular: "Тиын",
			Symbol:   "т",
		},
	},
	Texts: pkg.Texts{
		And:   "Және",
		Minus: "Минус",
		Only:  "Тек",
		Point: "Нүкте",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "Бір квадрилион"},
		{Number: 1000000000000, Value: "Бір трилион"},
		{Number: 1000000000, Value: "Бір миллиард"},
		{Number: 1000000, Value: "Бір миллион"},
		{Number: 1000, Value: "Бір мың"},
		{Number: 100, Value: "Бір жүз"},
		{Number: 90, Value: "Тоқсан"},
		{Number: 80, Value: "Сексен"},
		{Number: 70, Value: "Жетпіс"},
		{Number: 60, Value: "Алпыс"},
		{Number: 50, Value: "Елу"},
		{Number: 40, Value: "Қырық"},
		{Number: 30, Value: "Отыз"},
		{Number: 20, Value: "Жиырма"},
		{Number: 19, Value: "Он тоғыз"},
		{Number: 18, Value: "Он сегіз"},
		{Number: 17, Value: "Он жеті"},
		{Number: 16, Value: "Он алты"},
		{Number: 15, Value: "Он бес"},
		{Number: 14, Value: "Он төрт"},
		{Number: 13, Value: "Он үш"},
		{Number: 12, Value: "Он екі"},
		{Number: 11, Value: "Он бір"},
		{Number: 10, Value: "Он"},
		{Number: 9, Value: "Тоғыз"},
		{Number: 8, Value: "Сегіз"},
		{Number: 7, Value: "Жеті"},
		{Number: 6, Value: "Алты"},
		{Number: 5, Value: "Бес"},
		{Number: 4, Value: "Төрт"},
		{Number: 3, Value: "Үш"},
		{Number: 2, Value: "Екі"},
		{Number: 1, Value: "Бір"},
		{Number: 0, Value: "Нөл"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Бір жүз"},
	},
}
