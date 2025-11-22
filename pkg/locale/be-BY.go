package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// BEBYLocale is a NumI18NLocale configured for Belarusian (Belarus) - be-BY
var BEBYLocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "Беларускі рубель",
		Plural:   "Беларускія рублі",
		Singular: "Беларускі рубель",
		Symbol:   "BYN",
		FractionUnit: pkg.FractionUnit{
			Name:     "Капейка",
			Plural:   "Капейкі",
			Singular: "Капейка",
			Symbol:   "к.",
		},
	},
	Texts: pkg.Texts{
		And:   "і",
		Minus: "мінус",
		Only:  "толькі",
		Point: "кропка",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "Квадрыльён"},
		{Number: 1000000000000, Value: "Трыльён"},
		{Number: 1000000000, Value: "Мільярд"},
		{Number: 1000000, Value: "Мільён"},
		{Number: 1000, Value: "Тысяча"},
		{Number: 100, Value: "Сотня"},
		{Number: 90, Value: "Девяноста"},
		{Number: 80, Value: "Восемдесят"},
		{Number: 70, Value: "Семдесят"},
		{Number: 60, Value: "Шэсцьдзясят"},
		{Number: 50, Value: "Пяцьдзесят"},
		{Number: 40, Value: "Сорак"},
		{Number: 30, Value: "Трыццаць"},
		{Number: 20, Value: "Дваццаць"},
		{Number: 19, Value: "Дзесяць дзевяць"},
		{Number: 18, Value: "Восемнаццаць"},
		{Number: 17, Value: "Семнаццаць"},
		{Number: 16, Value: "Шаснаццаць"},
		{Number: 15, Value: "Пятнаццаць"},
		{Number: 14, Value: "Чатырнаццаць"},
		{Number: 13, Value: "Трынаццаць"},
		{Number: 12, Value: "Дванаццаць"},
		{Number: 11, Value: "Адзінаццаць"},
		{Number: 10, Value: "Дзесяць"},
		{Number: 9, Value: "Девяць"},
		{Number: 8, Value: "Восем"},
		{Number: 7, Value: "Сем"},
		{Number: 6, Value: "Шэсць"},
		{Number: 5, Value: "Пяць"},
		{Number: 4, Value: "Чатыры"},
		{Number: 3, Value: "Тры"},
		{Number: 2, Value: "Два"},
		{Number: 1, Value: "Адзін"},
		{Number: 0, Value: "Нуль"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Адзін Сотня"},
	},
}
