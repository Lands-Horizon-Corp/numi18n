package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// ZH001Locale is a NumI18NLocale configured for Chinese (World) - zh-001
var ZH001Locale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "元",
		Plural:   "元",
		Singular: "元",
		Symbol:   "¥",
		FractionUnit: pkg.FractionUnit{
			Name:     "分",
			Plural:   "分",
			Singular: "分",
			Symbol:   "¢",
		},
	},
	Texts: pkg.Texts{
		And:   "和",
		Minus: "负",
		Only:  "仅",
		Point: "点",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "千万亿"}, // Quadrillion
		{Number: 1000000000000, Value: "万亿"},     // Trillion
		{Number: 1000000000, Value: "十亿"},        // Billion
		{Number: 1000000, Value: "百万"},           // Million
		{Number: 1000, Value: "千"},               // Thousand
		{Number: 100, Value: "百"},                // Hundred
		{Number: 90, Value: "九十"},
		{Number: 80, Value: "八十"},
		{Number: 70, Value: "七十"},
		{Number: 60, Value: "六十"},
		{Number: 50, Value: "五十"},
		{Number: 40, Value: "四十"},
		{Number: 30, Value: "三十"},
		{Number: 20, Value: "二十"},
		{Number: 19, Value: "十九"},
		{Number: 18, Value: "十八"},
		{Number: 17, Value: "十七"},
		{Number: 16, Value: "十六"},
		{Number: 15, Value: "十五"},
		{Number: 14, Value: "十四"},
		{Number: 13, Value: "十三"},
		{Number: 12, Value: "十二"},
		{Number: 11, Value: "十一"},
		{Number: 10, Value: "十"},
		{Number: 9, Value: "九"},
		{Number: 8, Value: "八"},
		{Number: 7, Value: "七"},
		{Number: 6, Value: "六"},
		{Number: 5, Value: "五"},
		{Number: 4, Value: "四"},
		{Number: 3, Value: "三"},
		{Number: 2, Value: "二"},
		{Number: 1, Value: "一"},
		{Number: 0, Value: "零"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "一百"},
	},
}
