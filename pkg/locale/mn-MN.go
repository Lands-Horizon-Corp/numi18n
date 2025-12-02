package locale

import (
	"strings"
	"github.com/shopspring/decimal"
)

// MNMNLocale represents the Mongolian (Mongolia) locale
var MNMNLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Төгрөг",
		Plural:   "Төгрөг",
		Singular: "Төгрөг",
		Symbol:   "₮",
		FractionUnit: FractionUnit{
			Name:     "Мөнгө",
			Plural:   "Мөнгө",
			Singular: "Мөнгө",
			Symbol:   "м",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Mongolia",
		Currency:       "MNT",
		ISO3166Alpha2:  "MN",
		ISO3166Alpha3:  "MNG",
		ISO3166Numeric: "496",
		Locale:         "mn-MN",
		Timezone:       []string{"Asia/Ulaanbaatar"},
		Language:       "mn",
	},
	Texts: Texts{
		And:   "ба",
		Minus: "хасах",
		Only:  "зөвхөн",
		Point: "цэг",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "их наяд"},
		{Number: 1000000000000, Value: "их тэрбум"},
		{Number: 1000000000, Value: "тэрбум"},
		{Number: 1000000, Value: "сая"},
		{Number: 100000, Value: "зуун мянга"},
		{Number: 90000, Value: "ерэн мянга"},
		{Number: 80000, Value: "наян мянга"},
		{Number: 70000, Value: "далан мянга"},
		{Number: 60000, Value: "жаран мянга"},
		{Number: 50000, Value: "тавин мянга"},
		{Number: 40000, Value: "дөчин мянга"},
		{Number: 30000, Value: "гучин мянга"},
		{Number: 20000, Value: "хорин мянга"},
		{Number: 19000, Value: "арван есөн мянга"},
		{Number: 18000, Value: "арван найман мянга"},
		{Number: 17000, Value: "арван долоон мянга"},
		{Number: 16000, Value: "арван зургаан мянга"},
		{Number: 15000, Value: "арван таван мянга"},
		{Number: 14000, Value: "арван дөрвөн мянга"},
		{Number: 13000, Value: "арван гурван мянга"},
		{Number: 12000, Value: "арван хоёр мянга"},
		{Number: 11000, Value: "арван нэгэн мянга"},
		{Number: 10000, Value: "арван мянга"},
		{Number: 9000, Value: "есөн мянга"},
		{Number: 8000, Value: "найман мянга"},
		{Number: 7000, Value: "долоон мянга"},
		{Number: 6000, Value: "зургаан мянга"},
		{Number: 5000, Value: "таван мянга"},
		{Number: 4000, Value: "дөрвөн мянга"},
		{Number: 3000, Value: "гурван мянга"},
		{Number: 2000, Value: "хоёр мянга"},
		{Number: 1000, Value: "нэг мянга"},
		{Number: 900, Value: "есөн зуу"},
		{Number: 800, Value: "найман зуу"},
		{Number: 700, Value: "долоон зуу"},
		{Number: 600, Value: "зургаан зуу"},
		{Number: 500, Value: "таван зуу"},
		{Number: 400, Value: "дөрвөн зуу"},
		{Number: 300, Value: "гурван зуу"},
		{Number: 200, Value: "хоёр зуу"},
		{Number: 100, Value: "нэг зуу"},
		{Number: 99, Value: "ерэн есөн"},
		{Number: 98, Value: "ерэн найман"},
		{Number: 97, Value: "ерэн долоон"},
		{Number: 96, Value: "ерэн зургаан"},
		{Number: 95, Value: "ерэн таван"},
		{Number: 94, Value: "ерэн дөрвөн"},
		{Number: 93, Value: "ерэн гурван"},
		{Number: 92, Value: "ерэн хоёр"},
		{Number: 91, Value: "ерэн нэгэн"},
		{Number: 90, Value: "ерэн"},
		{Number: 89, Value: "наян есөн"},
		{Number: 88, Value: "наян найман"},
		{Number: 87, Value: "наян долоон"},
		{Number: 86, Value: "наян зургаан"},
		{Number: 85, Value: "наян таван"},
		{Number: 84, Value: "наян дөрвөн"},
		{Number: 83, Value: "наян гурван"},
		{Number: 82, Value: "наян хоёр"},
		{Number: 81, Value: "наян нэгэн"},
		{Number: 80, Value: "наян"},
		{Number: 79, Value: "далан есөн"},
		{Number: 78, Value: "далан найман"},
		{Number: 77, Value: "далан долоон"},
		{Number: 76, Value: "далан зургаан"},
		{Number: 75, Value: "далан таван"},
		{Number: 74, Value: "далан дөрвөн"},
		{Number: 73, Value: "далан гурван"},
		{Number: 72, Value: "далан хоёр"},
		{Number: 71, Value: "далан нэгэн"},
		{Number: 70, Value: "далан"},
		{Number: 69, Value: "жаран есөн"},
		{Number: 68, Value: "жаран найман"},
		{Number: 67, Value: "жаран долоон"},
		{Number: 66, Value: "жаран зургаан"},
		{Number: 65, Value: "жаран таван"},
		{Number: 64, Value: "жаран дөрвөн"},
		{Number: 63, Value: "жаран гурван"},
		{Number: 62, Value: "жаран хоёр"},
		{Number: 61, Value: "жаран нэгэн"},
		{Number: 60, Value: "жаран"},
		{Number: 59, Value: "тавин есөн"},
		{Number: 58, Value: "тавин найман"},
		{Number: 57, Value: "тавин долоон"},
		{Number: 56, Value: "тавин зургаан"},
		{Number: 55, Value: "тавин таван"},
		{Number: 54, Value: "тавин дөрвөн"},
		{Number: 53, Value: "тавин гурван"},
		{Number: 52, Value: "тавин хоёр"},
		{Number: 51, Value: "тавин нэгэн"},
		{Number: 50, Value: "тавин"},
		{Number: 49, Value: "дөчин есөн"},
		{Number: 48, Value: "дөчин найман"},
		{Number: 47, Value: "дөчин долоон"},
		{Number: 46, Value: "дөчин зургаан"},
		{Number: 45, Value: "дөчин таван"},
		{Number: 44, Value: "дөчин дөрвөн"},
		{Number: 43, Value: "дөчин гурван"},
		{Number: 42, Value: "дөчин хоёр"},
		{Number: 41, Value: "дөчин нэгэн"},
		{Number: 40, Value: "дөчин"},
		{Number: 39, Value: "гучин есөн"},
		{Number: 38, Value: "гучин найман"},
		{Number: 37, Value: "гучин долоон"},
		{Number: 36, Value: "гучин зургаан"},
		{Number: 35, Value: "гучин таван"},
		{Number: 34, Value: "гучин дөрвөн"},
		{Number: 33, Value: "гучин гурван"},
		{Number: 32, Value: "гучин хоёр"},
		{Number: 31, Value: "гучин нэгэн"},
		{Number: 30, Value: "гучин"},
		{Number: 29, Value: "хорин есөн"},
		{Number: 28, Value: "хорин найман"},
		{Number: 27, Value: "хорин долоон"},
		{Number: 26, Value: "хорин зургаан"},
		{Number: 25, Value: "хорин таван"},
		{Number: 24, Value: "хорин дөрвөн"},
		{Number: 23, Value: "хорин гурван"},
		{Number: 22, Value: "хорин хоёр"},
		{Number: 21, Value: "хорин нэгэн"},
		{Number: 20, Value: "хорин"},
		{Number: 19, Value: "арван есөн"},
		{Number: 18, Value: "арван найман"},
		{Number: 17, Value: "арван долоон"},
		{Number: 16, Value: "арван зургаан"},
		{Number: 15, Value: "арван таван"},
		{Number: 14, Value: "арван дөрвөн"},
		{Number: 13, Value: "арван гурван"},
		{Number: 12, Value: "арван хоёр"},
		{Number: 11, Value: "арван нэгэн"},
		{Number: 10, Value: "арван"},
		{Number: 9, Value: "есөн"},
		{Number: 8, Value: "найман"},
		{Number: 7, Value: "долоон"},
		{Number: 6, Value: "зургаан"},
		{Number: 5, Value: "таван"},
		{Number: 4, Value: "дөрвөн"},
		{Number: 3, Value: "гурван"},
		{Number: 2, Value: "хоёр"},
		{Number: 1, Value: "нэгэн"},
		{Number: 0, Value: "тэг"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Нэг зуу"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "нэгдүгээр", Suffix: "-дүгээр", Masculine: "нэгдүгээр", Feminine: "нэгдүгээр", Neuter: "нэгдүгээр"},
		{Number: 2, Word: "хоёрдугаар", Suffix: "-дугаар", Masculine: "хоёрдугаар", Feminine: "хоёрдугаар", Neuter: "хоёрдугаар"},
		{Number: 3, Word: "гуравдугаар", Suffix: "-дугаар", Masculine: "гуравдугаар", Feminine: "гуравдугаар", Neuter: "гуравдугаар"},
		{Number: 4, Word: "дөрөвдүгээр", Suffix: "-дүгээр", Masculine: "дөрөвдүгээр", Feminine: "дөрөвдүгээр", Neuter: "дөрөвдүгээр"},
		{Number: 5, Word: "тавдугаар", Suffix: "-дугаар", Masculine: "тавдугаар", Feminine: "тавдугаар", Neuter: "тавдугаар"},
		{Number: 6, Word: "зургаадугаар", Suffix: "-дугаар", Masculine: "зургаадугаар", Feminine: "зургаадугаар", Neuter: "зургаадугаар"},
		{Number: 7, Word: "долоодугаар", Suffix: "-дугаар", Masculine: "долоодугаар", Feminine: "долоодугаар", Neuter: "долоодугаар"},
		{Number: 8, Word: "найдугаар", Suffix: "-дугаар", Masculine: "найдугаар", Feminine: "найдугаар", Neuter: "найдугаар"},
		{Number: 9, Word: "есдүгээр", Suffix: "-дүгээр", Masculine: "есдүгээр", Feminine: "есдүгээр", Neuter: "есдүгээр"},
		{Number: 10, Word: "аравдугаар", Suffix: "-дугаар", Masculine: "аравдугаар", Feminine: "аравдугаар", Neuter: "аравдугаар"},
		{Number: 11, Word: "арван нэгдүгээр", Suffix: "-дүгээр", Masculine: "арван нэгдүгээр", Feminine: "арван нэгдүгээр", Neuter: "арван нэгдүгээр"},
		{Number: 12, Word: "арван хоёрдугаар", Suffix: "-дугаар", Masculine: "арван хоёрдугаар", Feminine: "арван хоёрдугаар", Neuter: "арван хоёрдугаар"},
		{Number: 20, Word: "хоридугаар", Suffix: "-дугаар", Masculine: "хоридугаар", Feminine: "хоридугаар", Neuter: "хоридугаар"},
		{Number: 21, Word: "хорин нэгдүгээр", Suffix: "-дүгээр", Masculine: "хорин нэгдүгээр", Feminine: "хорин нэгдүгээр", Neuter: "хорин нэгдүгээр"},
		{Number: 30, Word: "гучдугаар", Suffix: "-дугаар", Masculine: "гучдугаар", Feminine: "гучдугаар", Neuter: "гучдугаар"},
		{Number: 100, Word: "зуудугаар", Suffix: "-дугаар", Masculine: "зуудугаар", Feminine: "зуудугаар", Neuter: "зуудугаар"},
		{Number: 1000, Word: "мянгадугаар", Suffix: "-дугаар", Masculine: "мянгадугаар", Feminine: "мянгадугаар", Neuter: "мянгадугаар"},
	},
	LocaleFormatter: &MongolianFormatter{},
}

// MongolianFormatter handles Mongolian-specific formatting
type MongolianFormatter struct{}

func (f *MongolianFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *MongolianFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *MongolianFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *MongolianFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *MongolianFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *MongolianFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	if precision < 0 {
		precision = 0
	}
	return amount.Truncate(int32(precision))
}


func (f *MongolianFormatter) FormatDecimalNumber(amount float64) string {
	return DefaultFormatDecimalNumber(amount, ",", ".")
}
func (f *MongolianFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	formattedNumber := f.FormatDecimalNumber(amount)
	
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}
	
	// Default currency placement for this locale (prefix with symbol)
	if strings.HasPrefix(formattedNumber, "-") {
		formattedNumber = strings.TrimPrefix(formattedNumber, "-")
		return "-" + currencySymbol + formattedNumber
	}
	
	return currencySymbol + formattedNumber
}
