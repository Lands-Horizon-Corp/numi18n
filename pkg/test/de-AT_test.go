package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_AustrianGerman_Numbers(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Zero",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Null",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Fünf",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Fünfzehn",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Dreißig",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Siebenundvierzig",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Einhundert",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Zweihundertsechsundfünfzig",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Eintausend",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Million",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Eins million zwei hundert dreißig vier tausend fünf hundert sechzig sieben",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToWords(tt.amount)
			if result != tt.expected {
				t.Errorf("ToWords() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestToWords_AustrianGerman_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One Euro",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Eins Euro",
		},
		{
			name:   "Multiple Euros",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Fünf Euro",
		},
		{
			name:   "Zero Euros",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Null Euro",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Million Euro",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToWords(tt.amount)
			if result != tt.expected {
				t.Errorf("ToWords() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestToWords_AustrianGerman_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Euros and one Cent",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Fünf Euro und eins Cent",
		},
		{
			name:   "Euros and multiple Cents",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Fünf Euro und fünfundzwanzig Cent",
		},
		{
			name:   "Only Cents",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Null Euro und neunundneunzig Cent",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Eintausendzweihundertvierunddreißig Euro und sechsundfünfzig Cent",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Einhundertdreiundzwanzig und fünfundvierzig",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToWords(tt.amount)
			if result != tt.expected {
				t.Errorf("ToWords() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestToWords_AustrianGerman_Negative(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Negative number basic",
			amount: -50,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Minus fünfzig",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Minus fünfundzwanzig Euro und fünfundsiebzig Cent",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "Negativ",
					},
				},
			},
			expected: "Negativ einhundert Euro",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToWords(tt.amount)
			if result != tt.expected {
				t.Errorf("ToWords() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestToWords_AustrianGerman_Formatting(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Uppercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "EINHUNDERTDREIUNDZWANZIG EURO",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "einhundertdreiundzwanzig euro",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Neunhundertneunundneunzig nur",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Fünfhundert Euro nur",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToWords(tt.amount)
			if result != tt.expected {
				t.Errorf("ToWords() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestToWords_AustrianGerman_CustomCurrency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Custom currency name",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "Schilling",
						Plural: "Schilling",
					},
				},
			},
			expected: "Einhundert Schilling",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "Franken",
						Plural:           "Franken",
						FractionUnitName: "Rappen",
						FractionPlural:   "Rappen",
					},
				},
			},
			expected: "Fünfzig Franken und fünfundzwanzig Rappen",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "Mark",
						Plural: "Mark",
					},
				},
			},
			expected: "Eins Mark",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToWords(tt.amount)
			if result != tt.expected {
				t.Errorf("ToWords() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestToWords_AustrianGerman_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Very small decimal",
			amount: 0.01,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Null Euro und eins Cent",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Elf",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Zwölf",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Einsundzwanzig",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Einhunderteins",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Eintausendeins",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToWords(tt.amount)
			if result != tt.expected {
				t.Errorf("ToWords() = %v, want %v", result, tt.expected)
			}
		})
	}
}
