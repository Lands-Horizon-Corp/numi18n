package locale

import (
	"strings"
)

type NumI18NLocales struct {
	Locale []NumI18NLocale
}

func NewNumI18Locales() NumI18NLocales {
	locales := []NumI18NLocale{
		AEELocale,
		AFLocale,
		AFZALocale,
		AMETLocale,
		AMLocale,
		AR001Locale,
		ARAELocale,
		ARBHLocale,
		ARDZLocale,
		AREGLocale,
		ARIQLocale,
		ARJOLocale,
		ARKWLocale,
		ARLBLocale,
		ARLYLocale,
		ARLocale,
		ARMALocale,
		AROMLocale,
		ARQALocale,
		ARSALocale,
		ARSYLocale,
		ARTNLocale,
		ARYELocale,
		ASINLocale,
		AUALocale,
		AZAZLocale,
		BEBYLocale,
		BELocale,
		BGBGLocale,
		BNBDLocale,
		BNINLocale,
		BOLocale,
		BSBALocale,
		CAESLocale,
		CAFRLocale,
		CALocale,
		CHFRLocale,
		CHITLocale,
		CLLocale,
		COLocale,
		CRLocale,
		CSCZLocale,
		CULocale,
		CYGBLocale,
		DADKLocale,
		DEATLocale,
		DECHLocale,
		DEDELocale,
		DOLocale,
		DVMVLocale,
		DZBTLocale,
		ECLocale,
		EELocale,
		ELGRLocale,
		ELocale,
		EN001Locale,
		EOLocale,
		ES001Locale,
		ESEULocale,
		ESGLLocale,
		FILocale,
		FOLocale,
		FR001Locale,
		FRLocale,
		GBGDLocale,
		GBLocale,
		GELocale,
		GHLocale,
		GLLocale,
		GTLocale,
		HKLocale,
		HNLocale,
		HRLocale,
		HTLocale,
		HULocale,
		IDJVLocale,
		IDLocale,
		IEGALocale,
		IELocale,
		ILLocale,
		INGULocale,
		INHILocale,
		INKNLocale,
		INLocale,
		IRLocale,
		ISLocale,
		ITLocale,
		JPLocale,
		KELocale,
		KGLocale,
		KHLocale,
		KRLocale,
		KZLocale,
		LBLULocale,
		LOLALocale,
		LTLTLocale,
		LUFRLocale,
		LVLVLocale,
		MAFRLocale,
		MGMGLocale,
		MINZLocale,
		MKMKLocale,
		MLINLocale,
		MNMNLocale,
		MRINLocale,
		MSMYLocale,
		MTLocale,
		MTMTLocale,
		MULocale,
		MYMMLocale,
		NENPLocale,
		NGHALocale,
		NGLocale,
		NILocale,
		NLBELocale,
		NLNLLocale,
		NNNOLocale,
		NONOLocale,
		NYMWLocale,
		NZLocale,
		ORINLocale,
		PAINLocale,
		PHPHLocale,
		PALocale,
		PELocale,
		PHLocale,
		PLPLLocale,
		PSAFLocale,
		PT001Locale,
		PTAOLocale,
		PTBRLocale,
		PTCVLocale,
		PTGWLocale,
		PTMOLocale,
		PTMZLocale,
		PTPTLocale,
		PYLocale,
		QUPELocale,
		RMCHLocale,
		ROROLocale,
		RU001Locale,
		RUBYLocale,
		RURULocale,
		RWRWLocale,
		SDPKLocale,
		SENOLocale,
		SGLocale,
		SILKLocale,
		SKSKLocale,
		SLSILocale,
		SNFRLocale,
		SOSOLocale,
		SQALLocale,
		SRRSLocale,
		SVFILocale,
		SVLocale,
		SVSELocale,
		SWKELocale,
		TAINLocale,
		THTHLocale,
		TIETLocale,
		TKTMLocale,
		TNBWLocale,
		TRKULocale,
		TRTRLocale,
		TZLocale,
		UGCNLocale,
		UKUALocale,
		URINLocale,
		URPKLocale,
		USLocale,
		UYLocale,
		UZUZLocale,
		VELocale,
		VIVNLocale,
		WOSNLocale,
		XHZALocale,
		YI001Locale,
		ZALocale,
		ZH001Locale,
		ZHCNLocale,
		ZHHKLocale,
		ZHMOLocale,
		ZHSGLocale,
		ZHTWLocale,
		ZMLocale,
		ZUZALocale,
		ZWLocale,
	}
	return NumI18NLocales{
		Locale: locales,
	}
}

// FindByCountryName returns all locales for a specific country name
// Example: "United States" returns en-US, es-US, etc.
// Country names are matched case-insensitively
func (n NumI18NLocales) FindByCountryName(countryName string) []NumI18NLocale {
	var results []NumI18NLocale
	countryName = strings.TrimSpace(strings.ToLower(countryName))
	for _, locale := range n.Locale {
		if strings.ToLower(locale.NumI18Identifier.CountryName) == countryName {
			results = append(results, locale)
		}
	}
	return results
}

// FindByCurrency returns all locales that use a specific currency
// Example: "USD" returns en-US, en-CA (for some contexts), etc.
// Currency codes are matched case-insensitively using ISO 4217 codes
func (n NumI18NLocales) FindByCurrency(currency string) []NumI18NLocale {
	var results []NumI18NLocale
	currency = strings.TrimSpace(strings.ToUpper(currency))
	for _, locale := range n.Locale {
		if strings.ToUpper(locale.NumI18Identifier.Currency) == currency {
			results = append(results, locale)
		}
	}
	return results
}

// FindByISO3166Alpha2 returns all locales for a specific ISO 3166-1 alpha-2 code
// Example: "US" returns en-US, es-US, etc.
// ISO codes are matched case-insensitively (2-letter country codes)
func (n NumI18NLocales) FindByISO3166Alpha2(iso2 string) []NumI18NLocale {
	var results []NumI18NLocale
	iso2 = strings.TrimSpace(strings.ToUpper(iso2))
	for _, locale := range n.Locale {
		if strings.ToUpper(locale.NumI18Identifier.ISO3166Alpha2) == iso2 {
			results = append(results, locale)
		}
	}
	return results
}

// FindByISO3166Alpha3 returns all locales for a specific ISO 3166-1 alpha-3 code
// Example: "USA" returns en-US, es-US, etc.
// ISO codes are matched case-insensitively (3-letter country codes)
func (n NumI18NLocales) FindByISO3166Alpha3(iso3 string) []NumI18NLocale {
	var results []NumI18NLocale
	iso3 = strings.TrimSpace(strings.ToUpper(iso3))
	for _, locale := range n.Locale {
		if strings.ToUpper(locale.NumI18Identifier.ISO3166Alpha3) == iso3 {
			results = append(results, locale)
		}
	}
	return results
}

// FindByISO3166Numeric returns the locale for a specific ISO 3166 numeric code
// Example: "840" returns the US locale. Returns nil if not found.
// Numeric codes are unique identifiers (3-digit country codes)
func (n NumI18NLocales) FindByISO3166Numeric(numeric string) *NumI18NLocale {
	numeric = strings.TrimSpace(numeric)
	for _, locale := range n.Locale {
		if locale.NumI18Identifier.ISO3166Numeric == numeric {
			return &locale
		}
	}
	return nil
}

// FindByLocale returns the locale for a specific locale string (e.g., "en-US", "fr-FR")
// Returns nil if the locale is not found. Locale strings are unique identifiers.
func (n NumI18NLocales) FindByLocale(localeStr string) *NumI18NLocale {
	localeStr = strings.TrimSpace(strings.ToLower(localeStr))
	for _, locale := range n.Locale {
		if strings.ToLower(locale.NumI18Identifier.Locale) == localeStr {
			return &locale
		}
	}
	return nil
}

// FindByTimezone returns all locales that use a specific timezone
// Example: "America/New_York" returns en-US, es-US, etc.
// Timezone matching is case-insensitive
func (n NumI18NLocales) FindByTimezone(timezone string) []NumI18NLocale {
	var results []NumI18NLocale
	timezone = strings.TrimSpace(timezone)
	for _, locale := range n.Locale {
		for _, tz := range locale.NumI18Identifier.Timezone {
			if strings.EqualFold(tz, timezone) {
				results = append(results, locale)
				break // Avoid duplicates if multiple timezones match
			}
		}
	}
	return results
}

// FindByLanguage returns all locales for a specific language code
// Example: "en" returns en-US, en-GB, en-CA, etc.
// Language codes are matched case-insensitively (2-letter ISO 639-1 codes)
func (n NumI18NLocales) FindByLanguage(language string) []NumI18NLocale {
	var results []NumI18NLocale
	language = strings.TrimSpace(strings.ToLower(language))
	for _, locale := range n.Locale {
		if strings.ToLower(locale.NumI18Identifier.Language) == language {
			results = append(results, locale)
		}
	}
	return results
}

// Find performs a flexible search across multiple fields and returns matching locales
// Searches in: CountryName, Currency, ISO codes, Locale, Timezone, and Language
// Example: "US" matches ISO codes, "USD" matches currency, "United" matches country name
// Uses partial matching and is case-insensitive. Returns deduplicated results.
func (n NumI18NLocales) Find(query string) []NumI18NLocale {
	if query == "" {
		return []NumI18NLocale{}
	}

	query = strings.TrimSpace(query)
	queryLower := strings.ToLower(query)
	queryUpper := strings.ToUpper(query)

	// Use a map to avoid duplicates
	resultMap := make(map[string]NumI18NLocale)

	for _, locale := range n.Locale {
		identifier := locale.NumI18Identifier
		matched := false
		// Check country name (case-insensitive)
		if strings.Contains(strings.ToLower(identifier.CountryName), queryLower) {
			matched = true
		}

		// Check currency (case-insensitive)
		if strings.Contains(strings.ToUpper(identifier.Currency), queryUpper) {
			matched = true
		}
		// Check ISO codes (case-insensitive)
		if strings.Contains(strings.ToUpper(identifier.ISO3166Alpha2), queryUpper) ||
			strings.Contains(strings.ToUpper(identifier.ISO3166Alpha3), queryUpper) ||
			strings.Contains(identifier.ISO3166Numeric, query) {
			matched = true
		}
		// Check locale string (case-insensitive)
		if strings.Contains(strings.ToLower(identifier.Locale), queryLower) {
			matched = true
		}
		// Check timezones (case-insensitive)
		for _, tz := range identifier.Timezone {
			if strings.Contains(strings.ToLower(tz), queryLower) {
				matched = true
				break
			}
		}
		// Check language (case-insensitive)
		if strings.Contains(strings.ToLower(identifier.Language), queryLower) {
			matched = true
		}

		if matched {
			resultMap[identifier.Locale] = locale
		}
	}

	// Convert map to slice
	var results []NumI18NLocale
	for _, locale := range resultMap {
		results = append(results, locale)
	}
	return results
}

// AllLocales returns all available locales in the system
// Useful for enumeration or when you need the complete list of supported locales
func (n NumI18NLocales) AllLocales() []NumI18NLocale {
	return n.Locale
}

// PerCountryLocales returns one preferred locale for each unique country/region
// This helps avoid duplicates and provides clean globalized locale selection
// Returns the most preferred locale per country based on language priority
func (n NumI18NLocales) PerCountryLocales() []NumI18NLocale {
	countryGroups := make(map[string][]NumI18NLocale)

	for _, locale := range n.Locale {
		key := locale.NumI18Identifier.ISO3166Alpha2
		if key == "" {
			key = locale.NumI18Identifier.CountryName
		}
		countryGroups[key] = append(countryGroups[key], locale)
	}
	var perCountryLocales []NumI18NLocale
	for _, locales := range countryGroups {
		if preferredLocale := n.prioritizeLocale(locales); preferredLocale != nil {
			perCountryLocales = append(perCountryLocales, *preferredLocale)
		}
	}
	return perCountryLocales
}

// FindMostMatchedLocale finds the best matching locale for the given identifier
// Prioritizes exact locale string matches, then falls back to other criteria with smart prioritization
func (n NumI18NLocales) FindMostMatchedLocale(identifier NumI18Identifier) *NumI18NLocale {
	// Try exact locale match first (most specific)
	if result := n.findByExactLocale(identifier); result != nil {
		return result
	}

	// Try language + country combination
	if result := n.findByLanguageAndCountry(identifier); result != nil {
		return result
	}

	// Try currency-based search
	if result := n.findByCurrencyFallback(identifier); result != nil {
		return result
	}

	// Try ISO code-based search
	if result := n.findByISOCodes(identifier); result != nil {
		return result
	}

	// No match found
	return nil
}

func (n NumI18NLocales) findByExactLocale(identifier NumI18Identifier) *NumI18NLocale {
	if identifier.Locale != "" {
		return n.FindByLocale(identifier.Locale)
	}
	return nil
}

func (n NumI18NLocales) findByLanguageAndCountry(identifier NumI18Identifier) *NumI18NLocale {
	if identifier.Language == "" {
		return nil
	}

	languageLocales := n.FindByLanguage(identifier.Language)

	if identifier.ISO3166Alpha2 != "" {
		countryLocales := n.FindByISO3166Alpha2(identifier.ISO3166Alpha2)
		// Find intersection
		for _, langLocale := range languageLocales {
			for _, countryLocale := range countryLocales {
				if langLocale.NumI18Identifier.Locale == countryLocale.NumI18Identifier.Locale {
					return &langLocale
				}
			}
		}
	}

	// Return prioritized language match if no country match
	if len(languageLocales) > 0 {
		return n.prioritizeLocale(languageLocales)
	}

	return nil
}

func (n NumI18NLocales) findByCurrencyFallback(identifier NumI18Identifier) *NumI18NLocale {
	if identifier.Currency != "" {
		currencyLocales := n.FindByCurrency(identifier.Currency)
		if len(currencyLocales) > 0 {
			return n.prioritizeLocale(currencyLocales)
		}
	}
	return nil
}

func (n NumI18NLocales) findByISOCodes(identifier NumI18Identifier) *NumI18NLocale {
	// Try ISO3166Alpha2
	if identifier.ISO3166Alpha2 != "" {
		alpha2Locales := n.FindByISO3166Alpha2(identifier.ISO3166Alpha2)
		if len(alpha2Locales) > 0 {
			return n.prioritizeLocale(alpha2Locales)
		}
	}

	// Try ISO3166Alpha3
	if identifier.ISO3166Alpha3 != "" {
		alpha3Locales := n.FindByISO3166Alpha3(identifier.ISO3166Alpha3)
		if len(alpha3Locales) > 0 {
			return n.prioritizeLocale(alpha3Locales)
		}
	}

	// Try ISO3166Numeric
	if identifier.ISO3166Numeric != "" {
		return n.FindByISO3166Numeric(identifier.ISO3166Numeric)
	}

	return nil
}

// FindMatch finds the best matching locale using any field value
// Accepts: country name, locale string, currency, ISO codes, timezone, or language
// Example: "US", "en-US", "USD", "America/New_York" all return en-US locale
// Returns nil if no match is found. Uses intelligent matching priority.
func (n NumI18NLocales) FindMatch(query string) *NumI18NLocale {
	if strings.TrimSpace(query) == "" {
		return nil
	}
	query = strings.TrimSpace(query)
	searchIdentifier := NumI18Identifier{
		CountryName:    query,
		Currency:       query,
		ISO3166Alpha2:  query,
		ISO3166Alpha3:  query,
		ISO3166Numeric: query,
		Locale:         query,
		Language:       query,
		Timezone:       []string{query},
	}
	return n.FindMostMatchedLocale(searchIdentifier)
}
