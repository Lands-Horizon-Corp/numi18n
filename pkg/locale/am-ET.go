package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// AMETLocale is a NumI18NLocale configured for Amharic (Ethiopia) - am-ET
var AMETLocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "ብር",
		Plural:   "ብር",
		Singular: "ብር",
		Symbol:   "Br",
		FractionUnit: pkg.FractionUnit{
			Name:     "ሚልስ",
			Plural:   "ሚልስ",
			Singular: "ሚልስ",
			Symbol:   "ms",
		},
	},
	Texts: pkg.Texts{
		And:   "እና",
		Minus: "አልፎ",
		Only:  "ብቻ",
		Point: "ነጥብ",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "ኩዋድሪሊዮን"},
		{Number: 1000000000000, Value: "ትሪሊዮን"},
		{Number: 1000000000, Value: "ቢሊዮን"},
		{Number: 1000000, Value: "ሚሊዮን"},
		{Number: 1000, Value: "ሺህ"},
		{Number: 100, Value: "መቶ"},
		{Number: 90, Value: "ሰባ ዓሥር"},
		{Number: 80, Value: "ሰማንያ"},
		{Number: 70, Value: "ሰባ"},
		{Number: 60, Value: "ስስት ዐሥር"},
		{Number: 50, Value: "ሃምሳ"},
		{Number: 40, Value: "አርባ"},
		{Number: 30, Value: "ሰላሳ"},
		{Number: 20, Value: "ሃያ"},
		{Number: 19, Value: "ታያ አምስት"},
		{Number: 18, Value: "አስራ ስምንት"},
		{Number: 17, Value: "አስራ ሰባ"},
		{Number: 16, Value: "አስራ ስስት"},
		{Number: 15, Value: "አስራ ሃምሳ"},
		{Number: 14, Value: "አስራ አርባ"},
		{Number: 13, Value: "አስራ ሶስት"},
		{Number: 12, Value: "አስራ ሁለት"},
		{Number: 11, Value: "አስራ አንድ"},
		{Number: 10, Value: "አስር"},
		{Number: 9, Value: "ዘጠኝ"},
		{Number: 8, Value: "ስምንት"},
		{Number: 7, Value: "ሰባ"},
		{Number: 6, Value: "ስድስት"},
		{Number: 5, Value: "አምስት"},
		{Number: 4, Value: "አራት"},
		{Number: 3, Value: "ሶስት"},
		{Number: 2, Value: "ሁለት"},
		{Number: 1, Value: "አንድ"},
		{Number: 0, Value: "ዜሮ"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "መቶ"},
	},
}
