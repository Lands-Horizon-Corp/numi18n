package locale

import (
	"github.com/shopspring/decimal"
)

// TAINLocale represents the Tamil (India) locale
var TAINLocale = NumI18NLocale{
	LocaleFormatter: &TamilIndiaFormatter{},
	Currency: Currency{
		Name:     "Indian Rupee",
		Plural:   "ரூபாய்கள்",
		Singular: "ரூபாய்",
		Symbol:   "₹",
		FractionUnit: FractionUnit{
			Name:     "Paisa",
			Plural:   "பைசாக்கள்",
			Singular: "பைசா",
			Symbol:   "p",
		},
	},
	NumI18Identifier: NumI18Identifier{
		CountryName:    "India",
		Currency:       "INR",
		ISO3166Alpha2:  "IN",
		ISO3166Alpha3:  "IND",
		ISO3166Numeric: "356",
		Locale:         "ta-IN",
		Timezone:       []string{"Asia/Kolkata"},
		Language:       "ta",
	},
	Texts: Texts{
		And:   "மற்றும்",
		Minus: "கழித்தல்",
		Only:  "மட்டும்",
		Point: "புள்ளி",
	},
	NumberWordsMapping: []NumberWordMapping{
		{Number: 1000000000000000, Value: "ஒரு பத்து லட்சம் கோடி"},
		{Number: 1000000000000, Value: "ஒரு லட்சம் கோடி"},
		{Number: 1000000000, Value: "ஒரு ஆயிரம் கோடி"},
		{Number: 1000000, Value: "ஒரு பத்து லட்சம்"},
		{Number: 100000, Value: "ஒரு லட்சம்"},
		{Number: 90000, Value: "தொண்ணூறு ஆயிரம்"},
		{Number: 80000, Value: "எண்பது ஆயிரம்"},
		{Number: 70000, Value: "எழுபது ஆயிரம்"},
		{Number: 60000, Value: "அறுபது ஆயிரம்"},
		{Number: 50000, Value: "ஐம்பது ஆயிரம்"},
		{Number: 40000, Value: "நாற்பது ஆயிரம்"},
		{Number: 30000, Value: "முப்பது ஆயிரம்"},
		{Number: 20000, Value: "இருபது ஆயிரம்"},
		{Number: 19000, Value: "பத்தொன்பது ஆயிரம்"},
		{Number: 18000, Value: "பதினெட்டு ஆயிரம்"},
		{Number: 17000, Value: "பதினேழு ஆயிரம்"},
		{Number: 16000, Value: "பதினாறு ஆயிரம்"},
		{Number: 15000, Value: "பதினைந்து ஆயிரம்"},
		{Number: 14000, Value: "பதினான்கு ஆயிரம்"},
		{Number: 13000, Value: "பதின்மூன்று ஆயிரம்"},
		{Number: 12000, Value: "பன்னிரண்டு ஆயிரம்"},
		{Number: 11000, Value: "பதினொன்று ஆயிரம்"},
		{Number: 10000, Value: "பத்து ஆயிரம்"},
		{Number: 9000, Value: "ஒன்பது ஆயிரம்"},
		{Number: 8000, Value: "எட்டு ஆயிரம்"},
		{Number: 7000, Value: "ஏழு ஆயிரம்"},
		{Number: 6000, Value: "ஆறு ஆயிரம்"},
		{Number: 5000, Value: "ஐந்து ஆயிரம்"},
		{Number: 4000, Value: "நான்கு ஆயிரம்"},
		{Number: 3000, Value: "மூன்று ஆயிரம்"},
		{Number: 2000, Value: "இரண்டு ஆயிரம்"},
		{Number: 1000, Value: "ஒரு ஆயிரம்"},
		{Number: 900, Value: "தொள்ளாயிரம்"},
		{Number: 800, Value: "எண்ணூறு"},
		{Number: 700, Value: "எழுநூறு"},
		{Number: 600, Value: "அறுநூறு"},
		{Number: 500, Value: "ஐந்நூறு"},
		{Number: 400, Value: "நானூறு"},
		{Number: 300, Value: "முன்னூறு"},
		{Number: 200, Value: "இருநூறு"},
		{Number: 100, Value: "நூறு"},
		{Number: 99, Value: "தொண்ணூற்று ஒன்பது"},
		{Number: 98, Value: "தொண்ணூற்று எட்டு"},
		{Number: 97, Value: "தொண்ணூற்று ஏழு"},
		{Number: 96, Value: "தொண்ணூற்று ஆறு"},
		{Number: 95, Value: "தொண்ணூற்று ஐந்து"},
		{Number: 94, Value: "தொண்ணூற்று நான்கு"},
		{Number: 93, Value: "தொண்ணூற்று மூன்று"},
		{Number: 92, Value: "தொண்ணூற்று இரண்டு"},
		{Number: 91, Value: "தொண்ணூற்று ஒன்று"},
		{Number: 90, Value: "தொண்ணூறு"},
		{Number: 89, Value: "எண்பத்து ஒன்பது"},
		{Number: 88, Value: "எண்பத்து எட்டு"},
		{Number: 87, Value: "எண்பத்து ஏழு"},
		{Number: 86, Value: "எண்பத்து ஆறு"},
		{Number: 85, Value: "எண்பத்து ஐந்து"},
		{Number: 84, Value: "எண்பத்து நான்கு"},
		{Number: 83, Value: "எண்பத்து மூன்று"},
		{Number: 82, Value: "எண்பத்து இரண்டு"},
		{Number: 81, Value: "எண்பத்து ஒன்று"},
		{Number: 80, Value: "எண்பது"},
		{Number: 79, Value: "எழுபத்து ஒன்பது"},
		{Number: 78, Value: "எழுபத்து எட்டு"},
		{Number: 77, Value: "எழுபத்து ஏழு"},
		{Number: 76, Value: "எழுபத்து ஆறு"},
		{Number: 75, Value: "எழுபத்து ஐந்து"},
		{Number: 74, Value: "எழுபத்து நான்கு"},
		{Number: 73, Value: "எழுபத்து மூன்று"},
		{Number: 72, Value: "எழுபத்து இரண்டு"},
		{Number: 71, Value: "எழுபத்து ஒன்று"},
		{Number: 70, Value: "எழுபது"},
		{Number: 69, Value: "அறுபத்து ஒன்பது"},
		{Number: 68, Value: "அறுபத்து எட்டு"},
		{Number: 67, Value: "அறுபத்து ஏழு"},
		{Number: 66, Value: "அறுபத்து ஆறு"},
		{Number: 65, Value: "அறுபத்து ஐந்து"},
		{Number: 64, Value: "அறுபத்து நான்கு"},
		{Number: 63, Value: "அறுபத்து மூன்று"},
		{Number: 62, Value: "அறுபத்து இரண்டு"},
		{Number: 61, Value: "அறுபத்து ஒன்று"},
		{Number: 60, Value: "அறுபது"},
		{Number: 59, Value: "ஐம்பத்து ஒன்பது"},
		{Number: 58, Value: "ஐம்பத்து எட்டு"},
		{Number: 57, Value: "ஐம்பத்து ஏழு"},
		{Number: 56, Value: "ஐம்பத்து ஆறு"},
		{Number: 55, Value: "ஐம்பத்து ஐந்து"},
		{Number: 54, Value: "ஐம்பத்து நான்கு"},
		{Number: 53, Value: "ஐம்பத்து மூன்று"},
		{Number: 52, Value: "ஐம்பத்து இரண்டு"},
		{Number: 51, Value: "ஐம்பத்து ஒன்று"},
		{Number: 50, Value: "ஐம்பது"},
		{Number: 49, Value: "நாற்பத்து ஒன்பது"},
		{Number: 48, Value: "நாற்பத்து எட்டு"},
		{Number: 47, Value: "நாற்பத்து ஏழு"},
		{Number: 46, Value: "நாற்பத்து ஆறு"},
		{Number: 45, Value: "நாற்பத்து ஐந்து"},
		{Number: 44, Value: "நாற்பத்து நான்கு"},
		{Number: 43, Value: "நாற்பத்து மூன்று"},
		{Number: 42, Value: "நாற்பத்து இரண்டு"},
		{Number: 41, Value: "நாற்பத்து ஒன்று"},
		{Number: 40, Value: "நாற்பது"},
		{Number: 39, Value: "முப்பத்து ஒன்பது"},
		{Number: 38, Value: "முப்பத்து எட்டு"},
		{Number: 37, Value: "முப்பத்து ஏழு"},
		{Number: 36, Value: "முப்பத்து ஆறு"},
		{Number: 35, Value: "முப்பத்து ஐந்து"},
		{Number: 34, Value: "முப்பத்து நான்கு"},
		{Number: 33, Value: "முப்பத்து மூன்று"},
		{Number: 32, Value: "முப்பத்து இரண்டு"},
		{Number: 31, Value: "முப்பத்து ஒன்று"},
		{Number: 30, Value: "முப்பது"},
		{Number: 29, Value: "இருபத்து ஒன்பது"},
		{Number: 28, Value: "இருபத்து எட்டு"},
		{Number: 27, Value: "இருபத்து ஏழு"},
		{Number: 26, Value: "இருபத்து ஆறு"},
		{Number: 25, Value: "இருபத்து ஐந்து"},
		{Number: 24, Value: "இருபத்து நான்கு"},
		{Number: 23, Value: "இருபத்து மூன்று"},
		{Number: 22, Value: "இருபத்து இரண்டு"},
		{Number: 21, Value: "இருபத்து ஒன்று"},
		{Number: 20, Value: "இருபது"},
		{Number: 19, Value: "பத்தொன்பது"},
		{Number: 18, Value: "பதினெட்டு"},
		{Number: 17, Value: "பதினேழு"},
		{Number: 16, Value: "பதினாறு"},
		{Number: 15, Value: "பதினைந்து"},
		{Number: 14, Value: "பதினான்கு"},
		{Number: 13, Value: "பதின்மூன்று"},
		{Number: 12, Value: "பன்னிரண்டு"},
		{Number: 11, Value: "பதினொன்று"},
		{Number: 10, Value: "பத்து"},
		{Number: 9, Value: "ஒன்பது"},
		{Number: 8, Value: "எட்டு"},
		{Number: 7, Value: "ஏழு"},
		{Number: 6, Value: "ஆறு"},
		{Number: 5, Value: "ஐந்து"},
		{Number: 4, Value: "நான்கு"},
		{Number: 3, Value: "மூன்று"},
		{Number: 2, Value: "இரண்டு"},
		{Number: 1, Value: "ஒன்று"},
		{Number: 0, Value: "பூஜ்யம்"},
	},
	ExactWordsMapping: []ExactWordMapping{
		{Number: 100, Value: "நூறு"},
	},
	OrdinalMapping: []OrdinalMapping{
		{Number: 1, Word: "முதலாவது", Suffix: "-ஆம்", Masculine: "முதலாவது", Feminine: "முதலாவது", Neuter: "முதலாவது"},
		{Number: 2, Word: "இரண்டாவது", Suffix: "-ஆம்", Masculine: "இரண்டாவது", Feminine: "இரண்டாவது", Neuter: "இரண்டாவது"},
		{Number: 3, Word: "மூன்றாவது", Suffix: "-ஆம்", Masculine: "மூன்றாவது", Feminine: "மூன்றாவது", Neuter: "மூன்றாவது"},
		{Number: 4, Word: "நான்காவது", Suffix: "-ஆம்", Masculine: "நான்காவது", Feminine: "நான்காவது", Neuter: "நான்காவது"},
		{Number: 5, Word: "ஐந்தாவது", Suffix: "-ஆம்", Masculine: "ஐந்தாவது", Feminine: "ஐந்தாவது", Neuter: "ஐந்தாவது"},
		{Number: 6, Word: "ஆறாவது", Suffix: "-ஆம்", Masculine: "ஆறாவது", Feminine: "ஆறாவது", Neuter: "ஆறாவது"},
		{Number: 7, Word: "ஏழாவது", Suffix: "-ஆம்", Masculine: "ஏழாவது", Feminine: "ஏழாவது", Neuter: "ஏழாவது"},
		{Number: 8, Word: "எட்டாவது", Suffix: "-ஆம்", Masculine: "எட்டாவது", Feminine: "எட்டாவது", Neuter: "எட்டாவது"},
		{Number: 9, Word: "ஒன்பதாவது", Suffix: "-ஆம்", Masculine: "ஒன்பதாவது", Feminine: "ஒன்பதாவது", Neuter: "ஒன்பதாவது"},
		{Number: 10, Word: "பத்தாவது", Suffix: "-ஆம்", Masculine: "பத்தாவது", Feminine: "பத்தாவது", Neuter: "பத்தாவது"},
		{Number: 11, Word: "பதினொன்றாவது", Suffix: "-ஆம்", Masculine: "பதினொன்றாவது", Feminine: "பதினொன்றாவது", Neuter: "பதினொன்றாவது"},
		{Number: 12, Word: "பன்னிரண்டாவது", Suffix: "-ஆம்", Masculine: "பன்னிரண்டாவது", Feminine: "பன்னிரண்டாவது", Neuter: "பன்னிரண்டாவது"},
		{Number: 20, Word: "இருபதாவது", Suffix: "-ஆம்", Masculine: "இருபதாவது", Feminine: "இருபதாவது", Neuter: "இருபதாவது"},
		{Number: 21, Word: "இருபத்து ஒன்றாவது", Suffix: "-ஆம்", Masculine: "இருபத்து ஒன்றாவது", Feminine: "இருபத்து ஒன்றாவது", Neuter: "இருபத்து ஒன்றாவது"},
		{Number: 30, Word: "முப்பதாவது", Suffix: "-ஆம்", Masculine: "முப்பதாவது", Feminine: "முப்பதாவது", Neuter: "முப்பதாவது"},
		{Number: 50, Word: "ஐம்பதாவது", Suffix: "-ஆம்", Masculine: "ஐம்பதாவது", Feminine: "ஐம்பதாவது", Neuter: "ஐம்பதாவது"},
		{Number: 100, Word: "நூறாவது", Suffix: "-ஆம்", Masculine: "நூறாவது", Feminine: "நூறாவது", Neuter: "நூறாவது"},
		{Number: 1000, Word: "ஆயிரமாவது", Suffix: "-ஆம்", Masculine: "ஆயிரமாவது", Feminine: "ஆயிரமாவது", Neuter: "ஆயிரமாவது"},
	},
}

// TamilIndiaFormatter handles Tamil (India) formatting
type TamilIndiaFormatter struct{}

func (f *TamilIndiaFormatter) FormatNumber(number int64, targetLocale NumI18NLocale) string {
	return ConvertToWordsWithExactMappingInt64(number, targetLocale)
}

func (f *TamilIndiaFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + " " + currencyName
	}
	return result + " " + currencyPlural
}

func (f *TamilIndiaFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + " " + andText + " " + fractionalWords
}

func (f *TamilIndiaFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + " " + fractionName
	}
	return result + " " + fractionPlural
}

func (f *TamilIndiaFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + " " + result
}

func (f *TamilIndiaFormatter) ChopDecimal(amount decimal.Decimal, precision int) decimal.Decimal {
	return amount.Truncate(int32(precision))
}
