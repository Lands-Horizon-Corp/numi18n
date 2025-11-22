package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// ARIQLocale is a NumI18NLocale configured for Arabic (Iraq) - ar-IQ
var ARIQLocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "دينار",
		Plural:   "دنانير",
		Singular: "دينار",
		Symbol:   "ع.د",
		FractionUnit: pkg.FractionUnit{
			Name:     "فلس",
			Plural:   "فلوس",
			Singular: "فلس",
			Symbol:   "ف",
		},
	},
	Texts: pkg.Texts{
		And:   "و",
		Minus: "ناقص",
		Only:  "فقط",
		Point: "فاصل",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "كوادريليون"},
		{Number: 1000000000000, Value: "تريليون"},
		{Number: 1000000000, Value: "مليار"},
		{Number: 1000000, Value: "مليون"},
		{Number: 1000, Value: "ألف"},
		{Number: 100, Value: "مئة"},
		{Number: 90, Value: "تسعون"},
		{Number: 80, Value: "ثمانون"},
		{Number: 70, Value: "سبعون"},
		{Number: 60, Value: "ستون"},
		{Number: 50, Value: "خمسون"},
		{Number: 40, Value: "أربعون"},
		{Number: 30, Value: "ثلاثون"},
		{Number: 20, Value: "عشرون"},
		{Number: 19, Value: "تسعة عشر"},
		{Number: 18, Value: "ثمانية عشر"},
		{Number: 17, Value: "سبعة عشر"},
		{Number: 16, Value: "ستة عشر"},
		{Number: 15, Value: "خمسة عشر"},
		{Number: 14, Value: "أربعة عشر"},
		{Number: 13, Value: "ثلاثة عشر"},
		{Number: 12, Value: "اثنا عشر"},
		{Number: 11, Value: "أحد عشر"},
		{Number: 10, Value: "عشرة"},
		{Number: 9, Value: "تسعة"},
		{Number: 8, Value: "ثمانية"},
		{Number: 7, Value: "سبعة"},
		{Number: 6, Value: "ستة"},
		{Number: 5, Value: "خمسة"},
		{Number: 4, Value: "أربعة"},
		{Number: 3, Value: "ثلاثة"},
		{Number: 2, Value: "اثنان"},
		{Number: 1, Value: "واحد"},
		{Number: 0, Value: "صفر"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "مئة"},
	},
}
