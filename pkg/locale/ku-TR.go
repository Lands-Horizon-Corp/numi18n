package locale

import "github.com/shopspring/decimal"

// TRKULocale is a NumI18NLocale configured for Turkey (ku-TR)
var TRKULocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Lîre",
		Plural:   "Lîre",
		Singular: "Lîre",
		Symbol:   "₺",
		FractionUnit: FractionUnit{
			Name:     "Kuruş",
			Plural:   "Kuruş",
			Singular: "Kuruş",
			Symbol:   "kr",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Turkey",
		Currency:       "TRY",
		ISO3166Alpha2:  "TR",
		ISO3166Alpha3:  "TUR",
		ISO3166Numeric: "792",
		Locale:         "ku-TR",
		Timezone:       []string{"Europe/Istanbul"},
		Language:       "ku",
	},
	Texts: Texts{
		And:   "Û",
		Minus: "Kêm",
		Only:  "Tenê",
		Point: "Xal",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Yek kuadrîlyon"},
		{Number: 1000000000000, Value: "Yek trîlyon"},
		{Number: 1000000000, Value: "Yek mîlyar"},
		{Number: 1000000, Value: "Yek mîlyon"},
		{Number: 1000, Value: "Yek hezar"},
		{Number: 100, Value: "Yek sed"},
		{Number: 90, Value: "Nod"},
		{Number: 80, Value: "Heştê"},
		{Number: 70, Value: "Heftê"},
		{Number: 60, Value: "Şêst"},
		{Number: 50, Value: "Pêncî"},
		{Number: 40, Value: "Çil"},
		{Number: 30, Value: "Sî"},
		{Number: 20, Value: "Bîst"},
		{Number: 19, Value: "Nozdeh"},
		{Number: 18, Value: "Hijdeh"},
		{Number: 17, Value: "Hevdeh"},
		{Number: 16, Value: "Şazdeh"},
		{Number: 15, Value: "Pazdeh"},
		{Number: 14, Value: "Çardeh"},
		{Number: 13, Value: "Sêzdeh"},
		{Number: 12, Value: "Dwanzdeh"},
		{Number: 11, Value: "Yazdeh"},
		{Number: 10, Value: "Deh"},
		{Number: 9, Value: "Neh"},
		{Number: 8, Value: "Heşt"},
		{Number: 7, Value: "Heft"},
		{Number: 6, Value: "Şeş"},
		{Number: 5, Value: "Pênc"},
		{Number: 4, Value: "Çar"},
		{Number: 3, Value: "Sê"},
		{Number: 2, Value: "Du"},
		{Number: 1, Value: "Yek"},
		{Number: 0, Value: "Sifir"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Yek sed"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "yekem", Suffix: "-em", Masculine: "yekem", Feminine: "yekem", Neuter: "yekem"},
		{Number: 2, Word: "duyem", Suffix: "-em", Masculine: "duyem", Feminine: "duyem", Neuter: "duyem"},
		{Number: 3, Word: "sêyem", Suffix: "-em", Masculine: "sêyem", Feminine: "sêyem", Neuter: "sêyem"},
		{Number: 4, Word: "çarem", Suffix: "-em", Masculine: "çarem", Feminine: "çarem", Neuter: "çarem"},
		{Number: 5, Word: "pêncem", Suffix: "-em", Masculine: "pêncem", Feminine: "pêncem", Neuter: "pêncem"},
		{Number: 6, Word: "şeşem", Suffix: "-em", Masculine: "şeşem", Feminine: "şeşem", Neuter: "şeşem"},
		{Number: 7, Word: "heftem", Suffix: "-em", Masculine: "heftem", Feminine: "heftem", Neuter: "heftem"},
		{Number: 8, Word: "heştem", Suffix: "-em", Masculine: "heştem", Feminine: "heştem", Neuter: "heştem"},
		{Number: 9, Word: "nehem", Suffix: "-em", Masculine: "nehem", Feminine: "nehem", Neuter: "nehem"},
		{Number: 10, Word: "dehem", Suffix: "-em", Masculine: "dehem", Feminine: "dehem", Neuter: "dehem"},
		{Number: 11, Word: "yazdehem", Suffix: "-em", Masculine: "yazdehem", Feminine: "yazdehem", Neuter: "yazdehem"},
		{Number: 12, Word: "dwanzdehem", Suffix: "-em", Masculine: "dwanzdehem", Feminine: "dwanzdehem", Neuter: "dwanzdehem"},
		{Number: 13, Word: "sêzdehem", Suffix: "-em", Masculine: "sêzdehem", Feminine: "sêzdehem", Neuter: "sêzdehem"},
		{Number: 14, Word: "çardehem", Suffix: "-em", Masculine: "çardehem", Feminine: "çardehem", Neuter: "çardehem"},
		{Number: 15, Word: "pazdehem", Suffix: "-em", Masculine: "pazdehem", Feminine: "pazdehem", Neuter: "pazdehem"},
		{Number: 16, Word: "şazdehem", Suffix: "-em", Masculine: "şazdehem", Feminine: "şazdehem", Neuter: "şazdehem"},
		{Number: 17, Word: "hevdehem", Suffix: "-em", Masculine: "hevdehem", Feminine: "hevdehem", Neuter: "hevdehem"},
		{Number: 18, Word: "hijdehem", Suffix: "-em", Masculine: "hijdehem", Feminine: "hijdehem", Neuter: "hijdehem"},
		{Number: 19, Word: "nozdehem", Suffix: "-em", Masculine: "nozdehem", Feminine: "nozdehem", Neuter: "nozdehem"},
		{Number: 20, Word: "bîstem", Suffix: "-em", Masculine: "bîstem", Feminine: "bîstem", Neuter: "bîstem"},
		{Number: 21, Word: "bîst û yekem", Suffix: "-em", Masculine: "bîst û yekem", Feminine: "bîst û yekem", Neuter: "bîst û yekem"},
		{Number: 30, Word: "sîyem", Suffix: "-em", Masculine: "sîyem", Feminine: "sîyem", Neuter: "sîyem"},
		{Number: 40, Word: "çilem", Suffix: "-em", Masculine: "çilem", Feminine: "çilem", Neuter: "çilem"},
		{Number: 50, Word: "pêncîyem", Suffix: "-em", Masculine: "pêncîyem", Feminine: "pêncîyem", Neuter: "pêncîyem"},
		{Number: 60, Word: "şêstem", Suffix: "-em", Masculine: "şêstem", Feminine: "şêstem", Neuter: "şêstem"},
		{Number: 70, Word: "heftêyem", Suffix: "-em", Masculine: "heftêyem", Feminine: "heftêyem", Neuter: "heftêyem"},
		{Number: 80, Word: "heştêyem", Suffix: "-em", Masculine: "heştêyem", Feminine: "heştêyem", Neuter: "heştêyem"},
		{Number: 90, Word: "nodem", Suffix: "-em", Masculine: "nodem", Feminine: "nodem", Neuter: "nodem"},
		{Number: 100, Word: "sedem", Suffix: "-em", Masculine: "sedem", Feminine: "sedem", Neuter: "sedem"},
		{Number: 1000, Word: "hezarem", Suffix: "-em", Masculine: "hezarem", Feminine: "hezarem", Neuter: "hezarem"},
	},
	LocaleFormatter: &KurdishFormatter{},
}

// KurdishFormatter handles Kurdish (ku-TR) formatting
type KurdishFormatter struct{}

func (f *KurdishFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *KurdishFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	// Kurdish doesn't distinguish between singular and plural for currency in this context
	return result + " " + currencyName
}

func (f *KurdishFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *KurdishFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	// Kurdish doesn't distinguish between singular and plural for fraction units in this context
	return result + " " + fractionName
}

func (f *KurdishFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *KurdishFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 2
	}
	return amount.Truncate(int32(precision))
}
