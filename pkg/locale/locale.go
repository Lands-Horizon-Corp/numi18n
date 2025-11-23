package locale

import (
	"strings"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

type NumI18NLocale struct {
	Locale []pkg.NumI18NLocale
}

func NewNumI18Locale() NumI18NLocale {
	locales := []pkg.NumI18NLocale{
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
	return NumI18NLocale{
		Locale: locales,
	}
}

// FindByCountryName returns all locales for a specific country name
func (n NumI18NLocale) FindByCountryName(countryName string) []pkg.NumI18NLocale {
	var results []pkg.NumI18NLocale
	countryName = strings.TrimSpace(strings.ToLower(countryName))
	for _, locale := range n.Locale {
		if strings.ToLower(locale.NumI18Identifier.CountryName) == countryName {
			results = append(results, locale)
		}
	}
	return results
}

// FindByCurrency returns all locales that use a specific currency
func (n NumI18NLocale) FindByCurrency(currency string) []pkg.NumI18NLocale {
	var results []pkg.NumI18NLocale
	currency = strings.TrimSpace(strings.ToUpper(currency))
	for _, locale := range n.Locale {
		if strings.ToUpper(locale.NumI18Identifier.Currency) == currency {
			results = append(results, locale)
		}
	}
	return results
}

// FindByISO3166Alpha2 returns the locale for a specific ISO 3166-1 alpha-2 code
func (n NumI18NLocale) FindByISO3166Alpha2(iso2 string) []pkg.NumI18NLocale {
	var results []pkg.NumI18NLocale
	iso2 = strings.TrimSpace(strings.ToUpper(iso2))
	for _, locale := range n.Locale {
		if strings.ToUpper(locale.NumI18Identifier.ISO3166Alpha2) == iso2 {
			results = append(results, locale)
		}
	}
	return results
}

// FindByISO3166Alpha3 returns the locale for a specific ISO 3166-1 alpha-3 code
func (n NumI18NLocale) FindByISO3166Alpha3(iso3 string) []pkg.NumI18NLocale {
	var results []pkg.NumI18NLocale
	iso3 = strings.TrimSpace(strings.ToUpper(iso3))
	for _, locale := range n.Locale {
		if strings.ToUpper(locale.NumI18Identifier.ISO3166Alpha3) == iso3 {
			results = append(results, locale)
		}
	}
	return results
}

// FindByISO3166Numeric returns the locale for a specific ISO 3166 numeric code
func (n NumI18NLocale) FindByISO3166Numeric(numeric string) []pkg.NumI18NLocale {
	var results []pkg.NumI18NLocale
	numeric = strings.TrimSpace(numeric)
	for _, locale := range n.Locale {
		if locale.NumI18Identifier.ISO3166Numeric == numeric {
			results = append(results, locale)
		}
	}
	return results
}

// FindByLocale returns the locale for a specific locale string (e.g., "en-US", "fr-FR")
func (n NumI18NLocale) FindByLocale(localeStr string) []pkg.NumI18NLocale {
	var results []pkg.NumI18NLocale
	localeStr = strings.TrimSpace(strings.ToLower(localeStr))
	for _, locale := range n.Locale {
		if strings.ToLower(locale.NumI18Identifier.Locale) == localeStr {
			results = append(results, locale)
		}
	}
	return results
}

// FindByTimezone returns all locales that use a specific timezone
func (n NumI18NLocale) FindByTimezone(timezone string) []pkg.NumI18NLocale {
	var results []pkg.NumI18NLocale
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
func (n NumI18NLocale) FindByLanguage(language string) []pkg.NumI18NLocale {
	var results []pkg.NumI18NLocale
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
func (n NumI18NLocale) Find(query string) []pkg.NumI18NLocale {
	if query == "" {
		return []pkg.NumI18NLocale{}
	}

	query = strings.TrimSpace(query)
	queryLower := strings.ToLower(query)
	queryUpper := strings.ToUpper(query)

	// Use a map to avoid duplicates
	resultMap := make(map[string]pkg.NumI18NLocale)

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
	var results []pkg.NumI18NLocale
	for _, locale := range resultMap {
		results = append(results, locale)
	}
	return results
}

// AllLocales returns all available locales
func (n NumI18NLocale) AllLocales() []pkg.NumI18NLocale {
	return n.Locale
}
