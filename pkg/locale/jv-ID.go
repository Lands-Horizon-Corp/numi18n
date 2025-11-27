package locale

import "github.com/shopspring/decimal"

// IDJVLocale is a NumI18NLocale configured for Indonesia (jv-ID)
var IDJVLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Rupiah",
		Plural:   "Rupiah",
		Singular: "Rupiah",
		Symbol:   "Rp",
		FractionUnit: FractionUnit{
			Name:     "Sen",
			Plural:   "Sen",
			Singular: "Sen",
			Symbol:   "s",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Indonesia",
		Currency:       "IDR",
		ISO3166Alpha2:  "ID",
		ISO3166Alpha3:  "IDN",
		ISO3166Numeric: "360",
		Locale:         "jv-ID",
		Timezone:       []string{"Asia/Jakarta"},
		Language:       "jv",
	},
	Texts: Texts{
		And:   "Lan",
		Minus: "Minus",
		Only:  "Mung",
		Point: "Titik",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Siji kuadriliyun"},
		{Number: 1000000000000, Value: "Siji triliyun"},
		{Number: 1000000000, Value: "Siji milyar"},
		{Number: 1000000, Value: "Siji Yuta"},
		{Number: 1000, Value: "Siji Ewu"},
		{Number: 100, Value: "Satus"},
		{Number: 90, Value: "Sangang Puluh"},
		{Number: 80, Value: "Wolung Puluh"},
		{Number: 70, Value: "Pitung Puluh"},
		{Number: 60, Value: "Sewidak"},
		{Number: 50, Value: "Seket"},
		{Number: 40, Value: "Patang Puluh"},
		{Number: 30, Value: "Telung Puluh"},
		{Number: 20, Value: "Rongpuluh"},
		{Number: 19, Value: "Sangalas"},
		{Number: 18, Value: "Wolulas"},
		{Number: 17, Value: "Pitulas"},
		{Number: 16, Value: "Nembelas"},
		{Number: 15, Value: "Limolas"},
		{Number: 14, Value: "Patbelas"},
		{Number: 13, Value: "Telulas"},
		{Number: 12, Value: "Rolas"},
		{Number: 11, Value: "Sewelas"},
		{Number: 10, Value: "Sepuluh"},
		{Number: 9, Value: "Sanga"},
		{Number: 8, Value: "Wolu"},
		{Number: 7, Value: "Pitu"},
		{Number: 6, Value: "Enem"},
		{Number: 5, Value: "Lima"},
		{Number: 4, Value: "Papat"},
		{Number: 3, Value: "Telu"},
		{Number: 2, Value: "Loro"},
		{Number: 1, Value: "Siji"},
		{Number: 0, Value: "Kosong"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Satus"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "kaping pisanan", Suffix: ".", Masculine: "kaping pisanan", Feminine: "kaping pisanan", Neuter: "kaping pisanan"},
		{Number: 2, Word: "kaping pindho", Suffix: ".", Masculine: "kaping pindho", Feminine: "kaping pindho", Neuter: "kaping pindho"},
		{Number: 3, Word: "kaping telu", Suffix: ".", Masculine: "kaping telu", Feminine: "kaping telu", Neuter: "kaping telu"},
		{Number: 4, Word: "kaping papat", Suffix: ".", Masculine: "kaping papat", Feminine: "kaping papat", Neuter: "kaping papat"},
		{Number: 5, Word: "kaping lima", Suffix: ".", Masculine: "kaping lima", Feminine: "kaping lima", Neuter: "kaping lima"},
		{Number: 6, Word: "kaping enem", Suffix: ".", Masculine: "kaping enem", Feminine: "kaping enem", Neuter: "kaping enem"},
		{Number: 7, Word: "kaping pitu", Suffix: ".", Masculine: "kaping pitu", Feminine: "kaping pitu", Neuter: "kaping pitu"},
		{Number: 8, Word: "kaping wolu", Suffix: ".", Masculine: "kaping wolu", Feminine: "kaping wolu", Neuter: "kaping wolu"},
		{Number: 9, Word: "kaping sanga", Suffix: ".", Masculine: "kaping sanga", Feminine: "kaping sanga", Neuter: "kaping sanga"},
		{Number: 10, Word: "kaping sepuluh", Suffix: ".", Masculine: "kaping sepuluh", Feminine: "kaping sepuluh", Neuter: "kaping sepuluh"},
		{Number: 11, Word: "kaping sewelas", Suffix: ".", Masculine: "kaping sewelas", Feminine: "kaping sewelas", Neuter: "kaping sewelas"},
		{Number: 12, Word: "kaping rolas", Suffix: ".", Masculine: "kaping rolas", Feminine: "kaping rolas", Neuter: "kaping rolas"},
		{Number: 13, Word: "kaping telulas", Suffix: ".", Masculine: "kaping telulas", Feminine: "kaping telulas", Neuter: "kaping telulas"},
		{Number: 14, Word: "kaping patbelas", Suffix: ".", Masculine: "kaping patbelas", Feminine: "kaping patbelas", Neuter: "kaping patbelas"},
		{Number: 15, Word: "kaping limolas", Suffix: ".", Masculine: "kaping limolas", Feminine: "kaping limolas", Neuter: "kaping limolas"},
		{Number: 16, Word: "kaping nembelas", Suffix: ".", Masculine: "kaping nembelas", Feminine: "kaping nembelas", Neuter: "kaping nembelas"},
		{Number: 17, Word: "kaping pitulas", Suffix: ".", Masculine: "kaping pitulas", Feminine: "kaping pitulas", Neuter: "kaping pitulas"},
		{Number: 18, Word: "kaping wolulas", Suffix: ".", Masculine: "kaping wolulas", Feminine: "kaping wolulas", Neuter: "kaping wolulas"},
		{Number: 19, Word: "kaping sangalas", Suffix: ".", Masculine: "kaping sangalas", Feminine: "kaping sangalas", Neuter: "kaping sangalas"},
		{Number: 20, Word: "kaping rongpuluh", Suffix: ".", Masculine: "kaping rongpuluh", Feminine: "kaping rongpuluh", Neuter: "kaping rongpuluh"},
		{Number: 21, Word: "kaping rongpuluh siji", Suffix: ".", Masculine: "kaping rongpuluh siji", Feminine: "kaping rongpuluh siji", Neuter: "kaping rongpuluh siji"},
		{Number: 30, Word: "kaping telung puluh", Suffix: ".", Masculine: "kaping telung puluh", Feminine: "kaping telung puluh", Neuter: "kaping telung puluh"},
		{Number: 40, Word: "kaping patang puluh", Suffix: ".", Masculine: "kaping patang puluh", Feminine: "kaping patang puluh", Neuter: "kaping patang puluh"},
		{Number: 50, Word: "kaping seket", Suffix: ".", Masculine: "kaping seket", Feminine: "kaping seket", Neuter: "kaping seket"},
		{Number: 60, Word: "kaping sewidak", Suffix: ".", Masculine: "kaping sewidak", Feminine: "kaping sewidak", Neuter: "kaping sewidak"},
		{Number: 70, Word: "kaping pitung puluh", Suffix: ".", Masculine: "kaping pitung puluh", Feminine: "kaping pitung puluh", Neuter: "kaping pitung puluh"},
		{Number: 80, Word: "kaping wolung puluh", Suffix: ".", Masculine: "kaping wolung puluh", Feminine: "kaping wolung puluh", Neuter: "kaping wolung puluh"},
		{Number: 90, Word: "kaping sangang puluh", Suffix: ".", Masculine: "kaping sangang puluh", Feminine: "kaping sangang puluh", Neuter: "kaping sangang puluh"},
		{Number: 100, Word: "kaping satus", Suffix: ".", Masculine: "kaping satus", Feminine: "kaping satus", Neuter: "kaping satus"},
		{Number: 1000, Word: "kaping sewu", Suffix: ".", Masculine: "kaping sewu", Feminine: "kaping sewu", Neuter: "kaping sewu"},
	},
	LocaleFormatter: &JavaneseFormatter{},
}

// JavaneseFormatter handles Javanese formatting
type JavaneseFormatter struct{}

func (f *JavaneseFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *JavaneseFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *JavaneseFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *JavaneseFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *JavaneseFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *JavaneseFormatter) ChopDecimal(value decimal.Decimal, places int) decimal.Decimal {
	multiplier := decimal.NewFromInt(1)
	for range places {
		multiplier = multiplier.Mul(decimal.NewFromInt(10))
	}
	return value.Mul(multiplier).Truncate(0).Div(multiplier)
}
