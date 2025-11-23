package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// KHLocale is a NumI18NLocale configured for Cambodia (km-KH)
var KHLocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "រៀល",
		Plural:   "រៀល",
		Singular: "រៀល",
		Symbol:   "៛",
		FractionUnit: pkg.FractionUnit{
			Name:     "សេន",
			Plural:   "សេន",
			Singular: "សេន",
			Symbol:   "សេន",
		},
	},
	Texts: pkg.Texts{
		And:   "និង",
		Minus: "ដក",
		Only:  "តែប៉ុណ្ណោះ",
		Point: "ចំណុច",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "មួយ​ខ្វាដ្រីលាន"},
		{Number: 1000000000000, Value: "មួយ​ទ្រីលាន"},
		{Number: 1000000000, Value: "មួយ​ពាន់​លាន"},
		{Number: 1000000, Value: "មួយ​លាន"},
		{Number: 1000, Value: "មួយ​ពាន់"},
		{Number: 100, Value: "មួយ​រយ"},
		{Number: 90, Value: "កៅ​ស៊ប"},
		{Number: 80, Value: "ប៉ែត​ស៊ប"},
		{Number: 70, Value: "ចិត​ស៊ប"},
		{Number: 60, Value: "ហុក​ស៊ប"},
		{Number: 50, Value: "ហា​ស៊ប"},
		{Number: 40, Value: "សែ​ស៊ប"},
		{Number: 30, Value: "សាម​ស៊ប"},
		{Number: 20, Value: "ម្ភៃ"},
		{Number: 19, Value: "ដប់​កៅ"},
		{Number: 18, Value: "ដប់​ប៉ែត"},
		{Number: 17, Value: "ដប់​ព្រាំ"},
		{Number: 16, Value: "ដប់​ប្រាំ​មួយ"},
		{Number: 15, Value: "ដប់​ប្រាំ"},
		{Number: 14, Value: "ដប់​បួន"},
		{Number: 13, Value: "ដប់​បី"},
		{Number: 12, Value: "ដប់​ពីរ"},
		{Number: 11, Value: "ដប់​មួយ"},
		{Number: 10, Value: "ដប់"},
		{Number: 9, Value: "កៅ"},
		{Number: 8, Value: "ប៉ែត"},
		{Number: 7, Value: "ព្រាំ"},
		{Number: 6, Value: "ប្រាំ​មួយ"},
		{Number: 5, Value: "ប្រាំ"},
		{Number: 4, Value: "បួន"},
		{Number: 3, Value: "បី"},
		{Number: 2, Value: "ពីរ"},
		{Number: 1, Value: "មួយ"},
		{Number: 0, Value: "សូន្យ"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "មួយ​រយ"},
	},
}
