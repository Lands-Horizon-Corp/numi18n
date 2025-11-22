package locale

import "github.com/Lands-Horizon-Corp/numi18n/pkg"

// ELGRLocale is a NumI18NLocale configured for Greek (Greece) - el-GR
var ELGRLocale = pkg.NumI18NLocale{
	Currency: pkg.Currency{
		Name:     "Ευρώ",
		Plural:   "Ευρώ",
		Singular: "Ευρώ",
		Symbol:   "€",
		FractionUnit: pkg.FractionUnit{
			Name:     "Λεπτό",
			Plural:   "Λεπτά",
			Singular: "Λεπτό",
			Symbol:   "c",
		},
	},
	Texts: pkg.Texts{
		And:   "και",
		Minus: "μείον",
		Only:  "μόνο",
		Point: "τελεία",
	},
	NumberWordsMapping: []pkg.NumberWordMapping{
		{Number: 1000000000000000, Value: "Τετράκις εκατομμύριο"},
		{Number: 1000000000000, Value: "Τρισεκατομμύριο"},
		{Number: 1000000000, Value: "Δισεκατομμύριο"},
		{Number: 1000000, Value: "Εκατομμύριο"},
		{Number: 1000, Value: "Χίλια"},
		{Number: 100, Value: "Εκατό"},
		{Number: 90, Value: "Ενενήντα"},
		{Number: 80, Value: "Ογδόντα"},
		{Number: 70, Value: "Εβδομήντα"},
		{Number: 60, Value: "Εξήντα"},
		{Number: 50, Value: "Πενήντα"},
		{Number: 40, Value: "Σαράντα"},
		{Number: 30, Value: "Τριάντα"},
		{Number: 20, Value: "Είκοσι"},
		{Number: 19, Value: "Δεκαεννέα"},
		{Number: 18, Value: "Δεκαοκτώ"},
		{Number: 17, Value: "Δεκαεπτά"},
		{Number: 16, Value: "Δεκαέξι"},
		{Number: 15, Value: "Δεκαπέντε"},
		{Number: 14, Value: "Δεκατέσσερα"},
		{Number: 13, Value: "Δεκατρία"},
		{Number: 12, Value: "Δώδεκα"},
		{Number: 11, Value: "Έντεκα"},
		{Number: 10, Value: "Δέκα"},
		{Number: 9, Value: "Εννέα"},
		{Number: 8, Value: "Οκτώ"},
		{Number: 7, Value: "Επτά"},
		{Number: 6, Value: "Έξι"},
		{Number: 5, Value: "Πέντε"},
		{Number: 4, Value: "Τέσσερα"},
		{Number: 3, Value: "Τρία"},
		{Number: 2, Value: "Δύο"},
		{Number: 1, Value: "Ένα"},
		{Number: 0, Value: "Μηδέν"},
	},
	ExactWordsMapping: []pkg.ExactWordMapping{
		{Number: 100, Value: "Εκατό"},
	},
}
