package locale

type Currency struct {
	Name         string // Currency name, e.g., "Dollar"
	Plural       string // Currency plural, e.g., "Dollars"
	Singular     string // Currency singular, e.g., "Dollar"
	Symbol       string // Currency symbol, e.g., "$"
	FractionUnit FractionUnit
}

type FractionUnit struct {
	Name     string // Fractional unit name, e.g., "Cent"
	Plural   string // Fractional plural, e.g., "Cents"
	Singular string // Fractional singular, e.g., "Cent"
	Symbol   string // Fraction symbol, e.g., "¢"
}

type Texts struct {
	And   string // Text for "and", e.g., "And"
	Minus string // Text for "minus", e.g., "Minus"
	Only  string // Text for "only", e.g., "Only"
	Point string // Text for "point", e.g., "Point"
}

type NumberWordMapping struct {
	Number int64  // Numeric value, e.g., 1000
	Value  string // Word representation, e.g., "Thousand"
}

type ExactWordMapping struct {
	Number int64  // Exact numeric value, e.g., 100
	Value  string // Exact word representation, e.g., "One Hundred"
}

type OrdinalMapping struct {
	Number    int64  // Ordinal numeric value, e.g., 1, 2, 3
	Word      string // Ordinal word, e.g., "first", "second", "third"
	Suffix    string // Ordinal suffix for numeric form, e.g., "st", "nd", "rd", "th"
	Masculine string // Masculine form, e.g., "primer" (Spanish), "premier" (French)
	Feminine  string // Feminine form, e.g., "primera" (Spanish), "première" (French)
	Neuter    string // Neuter form, e.g., "erstes" (German)
}

type NumI18Identifier struct {
	CountryName    string   // Full country name, e.g., "Philippines"
	Currency       string   // Currency code, e.g., "PHP", "USD"
	ISO3166Alpha2  string   // ISO 3166-1 alpha-2 code, e.g., "PH"
	ISO3166Alpha3  string   // ISO 3166-1 alpha-3 code, e.g., "PHL"
	ISO3166Numeric string   // ISO 3166 numeric code, e.g., "608"
	Locale         string   // Locale string, e.g., "en-PH", "fr-FR"
	Timezone       []string // Timezone, e.g., "Asia/Manila"
	Language       string   // Language code, e.g., "en", "fr", "ja"
}

type NumI18NLocale struct {
	Currency           Currency
	NumI18Identifier   NumI18Identifier
	Texts              Texts
	NumberWordsMapping []NumberWordMapping
	ExactWordsMapping  []ExactWordMapping
	OrdinalMapping     []OrdinalMapping
}
