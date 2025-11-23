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

// FindByISO3166Alpha2 returns the locale for a specific ISO 3166-1 alpha-2 code
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

// FindByISO3166Alpha3 returns the locale for a specific ISO 3166-1 alpha-3 code
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
func (n NumI18NLocales) FindByISO3166Numeric(numeric string) []NumI18NLocale {
	var results []NumI18NLocale
	numeric = strings.TrimSpace(numeric)
	for _, locale := range n.Locale {
		if locale.NumI18Identifier.ISO3166Numeric == numeric {
			results = append(results, locale)
		}
	}
	return results
}

// FindByLocale returns the locale for a specific locale string (e.g., "en-US", "fr-FR")
func (n NumI18NLocales) FindByLocale(localeStr string) []NumI18NLocale {
	var results []NumI18NLocale
	localeStr = strings.TrimSpace(strings.ToLower(localeStr))
	for _, locale := range n.Locale {
		if strings.ToLower(locale.NumI18Identifier.Locale) == localeStr {
			results = append(results, locale)
		}
	}
	return results
}

// FindByTimezone returns all locales that use a specific timezone
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

// FindByLanguage returns all locales for a specific language code (e.g., "en", "fr")
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
// It searches in: CountryName, Currency, ISO codes, Locale, Timezone, and Language
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

// AllLocales returns all available locales
func (n NumI18NLocales) AllLocales() []NumI18NLocale {
	return n.Locale
}

func (n NumI18NLocales) FindMostMatchedLocale(identifier NumI18Identifier) *NumI18NLocale {
	if len(n.Locale) == 0 {
		return nil
	}

	type localeScore struct {
		locale       NumI18NLocale
		fieldMatches int
		priority     int
	}

	var candidates []localeScore

	for _, locale := range n.Locale {
		fieldMatches := 0
		priority := 0

		// Count field matches and calculate priority based on order
		if strings.EqualFold(
			strings.TrimSpace(locale.NumI18Identifier.Locale),
			strings.TrimSpace(identifier.Locale)) &&
			strings.TrimSpace(identifier.Locale) != "" {
			fieldMatches++
			priority += 7 // Highest priority
		}
		if strings.EqualFold(
			strings.TrimSpace(locale.NumI18Identifier.ISO3166Alpha2),
			strings.TrimSpace(identifier.ISO3166Alpha2)) &&
			strings.TrimSpace(identifier.ISO3166Alpha2) != "" {
			fieldMatches++
			priority += 6
		}
		if strings.EqualFold(
			strings.TrimSpace(locale.NumI18Identifier.ISO3166Alpha3),
			strings.TrimSpace(identifier.ISO3166Alpha3)) &&
			strings.TrimSpace(identifier.ISO3166Alpha3) != "" {
			fieldMatches++
			priority += 5
		}
		if strings.EqualFold(
			strings.TrimSpace(locale.NumI18Identifier.CountryName),
			strings.TrimSpace(identifier.CountryName)) &&
			strings.TrimSpace(identifier.CountryName) != "" {
			fieldMatches++
			priority += 4
		}
		if identifier.Timezone != nil {
			for _, searchTz := range identifier.Timezone {
				for _, localeTz := range locale.NumI18Identifier.Timezone {
					if strings.EqualFold(strings.TrimSpace(localeTz), strings.TrimSpace(searchTz)) {
						fieldMatches++
						priority += 3
						goto nextField1
					}
				}
			}
		}
	nextField1:
		if strings.EqualFold(strings.TrimSpace(locale.NumI18Identifier.Language), strings.TrimSpace(identifier.Language)) && strings.TrimSpace(identifier.Language) != "" {
			fieldMatches++
			priority += 2
		}
		if strings.TrimSpace(locale.NumI18Identifier.ISO3166Numeric) == strings.TrimSpace(identifier.ISO3166Numeric) && strings.TrimSpace(identifier.ISO3166Numeric) != "" {
			fieldMatches++
			priority += 1 // Lowest priority
		}

		// Currency bonus - if currencies match, add extra weight
		if strings.EqualFold(strings.TrimSpace(locale.NumI18Identifier.Currency), strings.TrimSpace(identifier.Currency)) && strings.TrimSpace(identifier.Currency) != "" {
			fieldMatches++
			priority += 8 // Currency gets highest bonus
		}

		if fieldMatches > 0 {
			candidates = append(candidates, localeScore{
				locale:       locale,
				fieldMatches: fieldMatches,
				priority:     priority,
			})
		}
	}

	if len(candidates) == 0 {
		return nil
	}

	// Sort by field matches first (descending), then by priority (descending)
	bestCandidate := candidates[0]
	for _, candidate := range candidates[1:] {
		if candidate.fieldMatches > bestCandidate.fieldMatches ||
			(candidate.fieldMatches == bestCandidate.fieldMatches && candidate.priority > bestCandidate.priority) {
			bestCandidate = candidate
		}
	}

	return &bestCandidate.locale
}

// FindMatch allows clients to find the best matching locale using any field value
// They can provide country name, locale string, currency, ISO codes, timezone, or language
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
