package currency

type NumI18NCurrency struct {
	CountryName    string // Full country name, e.g., "Philippines"
	Currency       string // Currency code, e.g., "PHP", "USD"
	Symbol         string // Currency symbol, e.g., "₱", "$"
	ISO3166Alpha2  string // ISO 3166-1 alpha-2 code, e.g., "PH"
	ISO3166Alpha3  string // ISO 3166-1 alpha-3 code, e.g., "PHL"
	ISO3166Numeric string // ISO 3166 numeric code, e.g., "608"
	Locale         string // Locale string, e.g., "en-PH", "fr-FR"
	Timezone       string // Timezone, e.g., "Asia/Manila"
	Language       string // Language code, e.g., "en", "fr", "ja"
	FractionDigits int    // Number of fraction digits, e.g., 2 for most currencies
	SymbolNative   string // Native symbol, if different from standard, e.g., "₱" for PHP
}
