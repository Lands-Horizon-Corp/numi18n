package pkg

type OverrideOptions struct {
	Name             string // Currency name override, e.g., "peso"
	Plural           string // Currency plural, e.g., "pesos"
	Symbol           string // Currency symbol, e.g., "₱"
	FractionUnitName string // Fractional unit name, e.g., "centavo"
	FractionPlural   string // Fractional plural, e.g., "centavos"
	FractionSymbol   string // Fraction symbol, e.g., "¢"
	NegativeWord     string // Custom negative word, e.g., "minus"
}

type WordDetails struct {
	Currency        bool // If true, include currency words, e.g., "one hundred pesos"
	Decimal         bool // If true, include decimals/fractions in words, e.g., "fifty-six centavos"
	Only            bool // If true, output only words without numbers or symbols
	Capitalize      bool // Capitalize the first letter of the resulting words
	Uppercase       bool // Convert entire output to uppercase
	Lowercase       bool // Convert entire output to lowercase
	NegativeWord    bool // If true, use the negative word instead of a negative sign, e.g., "minus one"
	OverrideOptions *OverrideOptions
}

type OrdinalDetails struct {
	SuffixOverride   string // Override suffix for numeric ordinals, e.g., "nd", "rd", "th", "º"
	PrefixOverride   string // Prefix for ordinal words, e.g., Japanese "第", Tagalog "ika-"
	UseWords         bool   // If true, outputs ordinal as words instead of numeric, e.g., "first"
	Gender           string // "masculine", "feminine", "neuter" for languages with gendered ordinals
	Plural           bool   // If true, pluralizes ordinal words, e.g., "firsts" instead of "first"
	Case             string // Grammatical case, e.g., "nominative", "accusative" (for Slavic languages)
	Style            string // Style of ordinal word, e.g., "short" | "long"
	Capitalize       bool   // Capitalize the first letter of the ordinal word
	FallbackToNumber bool   // If true, fallback to numeric form if word not available
	Separator        string // Separator between prefix/suffix and number/word, e.g., "", "-", " "
}
type NumI18NOptions struct {
	CountryName    string // Full country name, e.g., "Philippines"
	Currency       string // Currency code, e.g., "PHP", "USD"
	ISO3166Alpha2  string // ISO 3166-1 alpha-2 code, e.g., "PH"
	ISO3166Alpha3  string // ISO 3166-1 alpha-3 code, e.g., "PHL"
	ISO3166Numeric string // ISO 3166 numeric code, e.g., "608"
	Locale         string // Locale string, e.g., "en-PH", "fr-FR"
	Timezone       string // Timezone, e.g., "Asia/Manila"
	Language       string // Language code, e.g., "en", "fr", "ja"
	WordDetails    *WordDetails
	OrdinalDetails *OrdinalDetails
}
