package locale

import (
	"github.com/shopspring/decimal"
)

// GLLocale is a NumI18NLocale configured for Greenland (kl-GL)
var GLLocale = NumI18NLocale{
	Currency: Currency{
		Name:     "Dansk krone",
		Plural:   "Danske kroner",
		Singular: "Dansk krone",
		Symbol:   "kr",
		FractionUnit: FractionUnit{
			Name:     "Øre",
			Plural:   "Øre",
			Singular: "Øre",
			Symbol:   "ø",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "Greenland",
		Currency:       "DKK",
		ISO3166Alpha2:  "GL",
		ISO3166Alpha3:  "GRL",
		ISO3166Numeric: "304",
		Locale:         "kl-GL",
		Timezone:       []string{"America/Godthab"},
		Language:       "kl",
	},
	Texts: Texts{
		And:   "Aamma",
		Minus: "Minus",
		Only:  "Kissaat",
		Point: "Tikkuartoq",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "Ataaseq quadrillion"},
		{Number: 1000000000000, Value: "Ataaseq trillion"},
		{Number: 1000000000, Value: "Ataaseq milliardi"},
		{Number: 1000000, Value: "Ataaseq million"},
		{Number: 1000, Value: "Ataaseq tusind"},
		{Number: 100, Value: "Ataaseq hundrede"},
		{Number: 90, Value: "Qulingiluat"},
		{Number: 80, Value: "Arfersanerit"},
		{Number: 70, Value: "Arfersanik-qulit"},
		{Number: 60, Value: "Arfersanik"},
		{Number: 50, Value: "Tallimat-qulit"},
		{Number: 40, Value: "Tallimat"},
		{Number: 30, Value: "Pingasut-qulit"},
		{Number: 20, Value: "Juullut"},
		{Number: 19, Value: "Arfineq-sisamat"},
		{Number: 18, Value: "Arfineq-pingasut"},
		{Number: 17, Value: "Arfineq-marluk"},
		{Number: 16, Value: "Arfineq-ataaseq"},
		{Number: 15, Value: "Aqqaneq-tallimat"},
		{Number: 14, Value: "Aqqaneq-sisamat"},
		{Number: 13, Value: "Aqqaneq-pingasut"},
		{Number: 12, Value: "Aqqaneq-marluk"},
		{Number: 11, Value: "Aqqaneq-ataaseq"},
		{Number: 10, Value: "Qulit"},
		{Number: 9, Value: "Qulingiluat"},
		{Number: 8, Value: "Arfineq-pingasut"},
		{Number: 7, Value: "Arfineq-marluk"},
		{Number: 6, Value: "Arfineq-ataaseq"},
		{Number: 5, Value: "Tallimat"},
		{Number: 4, Value: "Sisamat"},
		{Number: 3, Value: "Pingasut"},
		{Number: 2, Value: "Marluk"},
		{Number: 1, Value: "Ataaseq"},
		{Number: 0, Value: "Nul"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "Ataaseq hundrede"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "sivulliit", Suffix: ".", Masculine: "sivulliit", Feminine: "sivulliit", Neuter: "sivulliit"},
		{Number: 2, Word: "aappaat", Suffix: ".", Masculine: "aappaat", Feminine: "aappaat", Neuter: "aappaat"},
		{Number: 3, Word: "pingasoq", Suffix: ".", Masculine: "pingasoq", Feminine: "pingasoq", Neuter: "pingasoq"},
		{Number: 4, Word: "sisamasoq", Suffix: ".", Masculine: "sisamasoq", Feminine: "sisamasoq", Neuter: "sisamasoq"},
		{Number: 5, Word: "tallimasoq", Suffix: ".", Masculine: "tallimasoq", Feminine: "tallimasoq", Neuter: "tallimasoq"},
		{Number: 6, Word: "arfinilluunniit", Suffix: ".", Masculine: "arfinilluunniit", Feminine: "arfinilluunniit", Neuter: "arfinilluunniit"},
		{Number: 7, Word: "arfineq marlunniit", Suffix: ".", Masculine: "arfineq marlunniit", Feminine: "arfineq marlunniit", Neuter: "arfineq marlunniit"},
		{Number: 8, Word: "arfineq pingasunniit", Suffix: ".", Masculine: "arfineq pingasunniit", Feminine: "arfineq pingasunniit", Neuter: "arfineq pingasunniit"},
		{Number: 9, Word: "qulingiluanniit", Suffix: ".", Masculine: "qulingiluanniit", Feminine: "qulingiluanniit", Neuter: "qulingiluanniit"},
		{Number: 10, Word: "qulilluunniit", Suffix: ".", Masculine: "qulilluunniit", Feminine: "qulilluunniit", Neuter: "qulilluunniit"},
		{Number: 11, Word: "aqqaneq ataasiunniit", Suffix: ".", Masculine: "aqqaneq ataasiunniit", Feminine: "aqqaneq ataasiunniit", Neuter: "aqqaneq ataasiunniit"},
		{Number: 12, Word: "aqqaneq marluunniit", Suffix: ".", Masculine: "aqqaneq marluunniit", Feminine: "aqqaneq marluunniit", Neuter: "aqqaneq marluunniit"},
		{Number: 13, Word: "aqqaneq pingasunniit", Suffix: ".", Masculine: "aqqaneq pingasunniit", Feminine: "aqqaneq pingasunniit", Neuter: "aqqaneq pingasunniit"},
		{Number: 14, Word: "aqqaneq sisamanniit", Suffix: ".", Masculine: "aqqaneq sisamanniit", Feminine: "aqqaneq sisamanniit", Neuter: "aqqaneq sisamanniit"},
		{Number: 15, Word: "aqqaneq tallimanniit", Suffix: ".", Masculine: "aqqaneq tallimanniit", Feminine: "aqqaneq tallimanniit", Neuter: "aqqaneq tallimanniit"},
		{Number: 16, Word: "arfineq ataasiunniit", Suffix: ".", Masculine: "arfineq ataasiunniit", Feminine: "arfineq ataasiunniit", Neuter: "arfineq ataasiunniit"},
		{Number: 17, Word: "arfineq marluunniit", Suffix: ".", Masculine: "arfineq marluunniit", Feminine: "arfineq marluunniit", Neuter: "arfineq marluunniit"},
		{Number: 18, Word: "arfineq pingasunniit", Suffix: ".", Masculine: "arfineq pingasunniit", Feminine: "arfineq pingasunniit", Neuter: "arfineq pingasunniit"},
		{Number: 19, Word: "arfineq sisamanniit", Suffix: ".", Masculine: "arfineq sisamanniit", Feminine: "arfineq sisamanniit", Neuter: "arfineq sisamanniit"},
		{Number: 20, Word: "juulluunniit", Suffix: ".", Masculine: "juulluunniit", Feminine: "juulluunniit", Neuter: "juulluunniit"},
		{Number: 21, Word: "juullut ataasiunniit", Suffix: ".", Masculine: "juullut ataasiunniit", Feminine: "juullut ataasiunniit", Neuter: "juullut ataasiunniit"},
		{Number: 30, Word: "pingasut qulilluunniit", Suffix: ".", Masculine: "pingasut qulilluunniit", Feminine: "pingasut qulilluunniit", Neuter: "pingasut qulilluunniit"},
		{Number: 40, Word: "tallimanniit", Suffix: ".", Masculine: "tallimanniit", Feminine: "tallimanniit", Neuter: "tallimanniit"},
		{Number: 50, Word: "tallimat qulilluunniit", Suffix: ".", Masculine: "tallimat qulilluunniit", Feminine: "tallimat qulilluunniit", Neuter: "tallimat qulilluunniit"},
		{Number: 60, Word: "arfersanilluunniit", Suffix: ".", Masculine: "arfersanilluunniit", Feminine: "arfersanilluunniit", Neuter: "arfersanilluunniit"},
		{Number: 70, Word: "arfersanik qulilluunniit", Suffix: ".", Masculine: "arfersanik qulilluunniit", Feminine: "arfersanik qulilluunniit", Neuter: "arfersanik qulilluunniit"},
		{Number: 80, Word: "arfersanerilluunniit", Suffix: ".", Masculine: "arfersanerilluunniit", Feminine: "arfersanerilluunniit", Neuter: "arfersanerilluunniit"},
		{Number: 90, Word: "qulingiluanniit", Suffix: ".", Masculine: "qulingiluanniit", Feminine: "qulingiluanniit", Neuter: "qulingiluanniit"},
		{Number: 100, Word: "hundredeunniit", Suffix: ".", Masculine: "hundredeunniit", Feminine: "hundredeunniit", Neuter: "hundredeunniit"},
		{Number: 1000, Word: "tusindiunniit", Suffix: ".", Masculine: "tusindiunniit", Feminine: "tusindiunniit", Neuter: "tusindiunniit"},
	},
	LocaleFormatter: &GreenlandicFormatter{},
}

// GreenlandicFormatter handles Greenlandic formatting
type GreenlandicFormatter struct{}

func (f *GreenlandicFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *GreenlandicFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *GreenlandicFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *GreenlandicFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *GreenlandicFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *GreenlandicFormatter) ChopDecimal(value decimal.Decimal, places int) decimal.Decimal {
	multiplier := decimal.NewFromInt(1)
	for i := 0; i < places; i++ {
		multiplier = multiplier.Mul(decimal.NewFromInt(10))
	}
	return value.Mul(multiplier).Truncate(0).Div(multiplier)
}

func (f *GreenlandicFormatter) FormatDecimalNumber(amount float64) string {
	return FormatEuropeanDecimal(amount)
}
func (f *GreenlandicFormatter) FormatDecimalNumberWithCurrency(amount float64, targetLocale NumI18NLocale, overrideOptions *OverrideOptions) string {
	// Get currency symbol
	currencySymbol := targetLocale.Currency.Symbol
	if overrideOptions != nil && overrideOptions.Symbol != "" {
		currencySymbol = overrideOptions.Symbol
	}

	// Format with European conventions (period separators, comma decimal, prefix symbol)
	return FormatEuropeanCurrency(amount, currencySymbol)
}
