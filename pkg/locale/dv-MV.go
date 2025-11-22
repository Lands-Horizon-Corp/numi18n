package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// DVMVLocale is a NumI18NLocale configured for Dhivehi (Maldives) - dv-MV
var DVMVLocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "Rufiyaa",
		Plural:   "Rufiyaa",
		Singular: "Rufiyaa",
		Symbol:   "Rf",
		FractionUnit: pkg.FractionUnit{
			Name:     "Laari",
			Plural:   "Laari",
			Singular: "Laari",
			Symbol:   "l",
		},
	},
	Texts: pkg.Texts{
		And:   "އަންޑް",
		Minus: "މިނަސް",
		Only:  "ނުވަތަ",
		Point: "ޕޮއިންޓް",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "ކްވަޑްރިލިއަން"},
		{Number: 1000000000000, Value: "ޓްރިލިއަން"},
		{Number: 1000000000, Value: "ބިލިއަން"},
		{Number: 1000000, Value: "މިލިއަން"},
		{Number: 1000, Value: "ތިސަން"},
		{Number: 100, Value: "ސްތޯ"},
		{Number: 90, Value: "ނަވެސް"},
		{Number: 80, Value: "އޮވެސް"},
		{Number: 70, Value: "ސެބެތް"},
		{Number: 60, Value: "ސެކްސެތް"},
		{Number: 50, Value: "ފިވެސް"},
		{Number: 40, Value: "ފިސްރަ"},
		{Number: 30, Value: "ޓްރިސް"},
		{Number: 20, Value: "ޓްވިން"},
		{Number: 19, Value: "ނިންޓީން"},
		{Number: 18, Value: "އެއިޓީން"},
		{Number: 17, Value: "ސެވެންޓީން"},
		{Number: 16, Value: "ސިސްޓީން"},
		{Number: 15, Value: "ފިފްޓީން"},
		{Number: 14, Value: "ފޮޓްޝެން"},
		{Number: 13, Value: "ތްރީޝެން"},
		{Number: 12, Value: "ޓްވޮލްވް"},
		{Number: 11, Value: "އެލްވް"},
		{Number: 10, Value: "ޓެން"},
		{Number: 9, Value: "ނަވް"},
		{Number: 8, Value: "އޮވް"},
		{Number: 7, Value: "ސެވް"},
		{Number: 6, Value: "ސެސް"},
		{Number: 5, Value: "ފިފް"},
		{Number: 4, Value: "ފޯރ"},
		{Number: 3, Value: "ތްރީ"},
		{Number: 2, Value: "ޓްނޯ"},
		{Number: 1, Value: "އަން"},
		{Number: 0, Value: "ސިރޯ"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "ސްތޯ"},
	},
}
