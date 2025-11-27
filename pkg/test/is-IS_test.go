package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_Icelandic_Numbers(t *testing.T) {
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
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Núll",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Fimm",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Fimmtán",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Þrjátíu",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Fjörutíu Sjö",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Eitt Hundrað",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tvö Hundrað Fimmtíu Sex",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Eitt Þúsund",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ein Milljón",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ein Milljón Tvö Hundrað Þrjátíu Fjögur Þúsund Fimm Hundrað Sextíu Sjö",
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

func TestToWords_Icelandic_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One króna",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Eitt Króna",
		},
		{
			name:   "Multiple krónur",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Fimm Krónur",
		},
		{
			name:   "Zero krónur",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Núll Krónur",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Ein Milljón Krónur",
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

func TestToWords_Icelandic_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Króna and one eyrir",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Fimm Krónur Og Eitt Eyrir",
		},
		{
			name:   "Króna and multiple aurar",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Fimm Krónur Og Tuttugu Fimm Aurar",
		},
		{
			name:   "Only aurar",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Núll Krónur Og Níutíu Níu Aurar",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Eitt Þúsund Tvö Hundrað Þrjátíu Fjögur Krónur Og Fimmtíu Sex Aurar",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Eitt Hundrað Tuttugu Þrjú Og Fjörutíu Fimm",
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

func TestToWords_Icelandic_Negative(t *testing.T) {
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
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Mínus Fimmtíu",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Mínus Tuttugu Fimm Krónur Og Sjötíu Fimm Aurar",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "Neikvæð",
					},
				},
			},
			expected: "Neikvæð Eitt Hundrað Krónur",
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

func TestToWords_Icelandic_Formatting(t *testing.T) {
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
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "EITT HUNDRAÐ TUTTUGU ÞRJÚ KRÓNUR",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "eitt hundrað tuttugu þrjú krónur",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Níu Hundrað Níutíu Níu Aðeins",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Fimm Hundrað Krónur Aðeins",
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

func TestToWords_Icelandic_CustomCurrency(t *testing.T) {
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
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "evra",
						Plural: "evrur",
					},
				},
			},
			expected: "Eitt Hundrað evrur",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "dalur",
						Plural:           "dalir",
						FractionUnitName: "sent",
						FractionPlural:   "sent",
					},
				},
			},
			expected: "Fimmtíu dalir Og Tuttugu Fimm sent",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "pund",
						Plural: "pund",
					},
				},
			},
			expected: "Eitt pund",
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

func TestToWords_Icelandic_EdgeCases(t *testing.T) {
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
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Núll Krónur Og Eitt Eyrir",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ellefu",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tólf",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tuttugu Eitt",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Eitt Hundrað Eitt",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Eitt Þúsund Eitt",
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
