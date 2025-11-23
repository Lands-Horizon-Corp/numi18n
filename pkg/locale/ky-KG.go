package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// KGLocale is a NumI18NLocale configured for Kyrgyzstan (ky-KG)
var KGLocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "Сом",
		Plural:   "Сом",
		Singular: "Сом",
		Symbol:   "с",
		FractionUnit: pkg.FractionUnit{
			Name:     "Тыйын",
			Plural:   "Тыйын",
			Singular: "Тыйын",
			Symbol:   "т",
		},
	},
	Texts: pkg.Texts{
		And:   "Жана",
		Minus: "Минус",
		Only:  "Гана",
		Point: "Чекит",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "Бир квадриллион"},
		{Number: 1000000000000, Value: "Бир триллион"},
		{Number: 1000000000, Value: "Бир миллиард"},
		{Number: 1000000, Value: "Бир миллион"},
		{Number: 1000, Value: "Бир миң"},
		{Number: 100, Value: "Бир жүз"},
		{Number: 90, Value: "Токсон"},
		{Number: 80, Value: "Сексен"},
		{Number: 70, Value: "Жетимиш"},
		{Number: 60, Value: "Алтымыш"},
		{Number: 50, Value: "Элүү"},
		{Number: 40, Value: "Кырк"},
		{Number: 30, Value: "Отуз"},
		{Number: 20, Value: "Жыйырма"},
		{Number: 19, Value: "Он тогуз"},
		{Number: 18, Value: "Он сегиз"},
		{Number: 17, Value: "Он жети"},
		{Number: 16, Value: "Он алты"},
		{Number: 15, Value: "Он беш"},
		{Number: 14, Value: "Он төрт"},
		{Number: 13, Value: "Он үч"},
		{Number: 12, Value: "Он эки"},
		{Number: 11, Value: "Он бир"},
		{Number: 10, Value: "Он"},
		{Number: 9, Value: "Тогуз"},
		{Number: 8, Value: "Сегиз"},
		{Number: 7, Value: "Жети"},
		{Number: 6, Value: "Алты"},
		{Number: 5, Value: "Беш"},
		{Number: 4, Value: "Төрт"},
		{Number: 3, Value: "Үч"},
		{Number: 2, Value: "Эки"},
		{Number: 1, Value: "Бир"},
		{Number: 0, Value: "Нөл"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Бир жүз"},
	},
}
