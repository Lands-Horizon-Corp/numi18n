package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// BGBGLocale is a NumI18NLocale configured for Bulgarian (Bulgaria) - bg-BG
var BGBGLocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "лев",
		Plural:   "лева",
		Singular: "лев",
		Symbol:   "лв",
		FractionUnit: pkg.FractionUnit{
			Name:     "стотинка",
			Plural:   "стотинки",
			Singular: "стотинка",
			Symbol:   "ст.",
		},
	},
	Texts: pkg.Texts{
		And:   "и",
		Minus: "минус",
		Only:  "само",
		Point: "точка",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "Квадриллион"},
		{Number: 1000000000000, Value: "Трилион"},
		{Number: 1000000000, Value: "Милиард"},
		{Number: 1000000, Value: "Милион"},
		{Number: 1000, Value: "Хиляда"},
		{Number: 100, Value: "Сто"},
		{Number: 90, Value: "Сто деветдесет"},
		{Number: 80, Value: "Осемдесет"},
		{Number: 70, Value: "Седемдесет"},
		{Number: 60, Value: "Шестдесет"},
		{Number: 50, Value: "Петдесет"},
		{Number: 40, Value: "Четиридесет"},
		{Number: 30, Value: "Тридесет"},
		{Number: 20, Value: "Двадесет"},
		{Number: 19, Value: "Деветнадесет"},
		{Number: 18, Value: "Осемнадесет"},
		{Number: 17, Value: "Седемнадесет"},
		{Number: 16, Value: "Шестнадесет"},
		{Number: 15, Value: "Петнадесет"},
		{Number: 14, Value: "Четирнадесет"},
		{Number: 13, Value: "Тринадесет"},
		{Number: 12, Value: "Дванадесет"},
		{Number: 11, Value: "Единадесет"},
		{Number: 10, Value: "Десет"},
		{Number: 9, Value: "Девет"},
		{Number: 8, Value: "Осем"},
		{Number: 7, Value: "Седем"},
		{Number: 6, Value: "Шест"},
		{Number: 5, Value: "Пет"},
		{Number: 4, Value: "Четири"},
		{Number: 3, Value: "Три"},
		{Number: 2, Value: "Две"},
		{Number: 1, Value: "Едно"},
		{Number: 0, Value: "Нула"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Сто"},
	},
}
