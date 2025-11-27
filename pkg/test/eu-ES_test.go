package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_BasqueSpain_Numbers(t *testing.T) {
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
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Zero",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Bost",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Hamabost",
		},
		{
			name:   "Twenty",
			amount: 20,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Hogei",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Berrogeita Zazpi",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ehun",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Bi Ehun Berrogeita Hamasei",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Mila",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Milioi Bat",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Bat Milioi Bi Ehun Hogeita Hamar Lau Mila Bost Ehun Hirurogei Zazpi",
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

func TestToWords_BasqueSpain_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One euro",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Bat Euro",
		},
		{
			name:   "Multiple euros",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Bost Euroak",
		},
		{
			name:   "Zero euros",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Zero Euroak",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Milioi Bat Euroak",
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

func TestToWords_BasqueSpain_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Euros and one zentimo",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Bost Euroak Eta Bat Zentimo",
		},
		{
			name:   "Euros and multiple zentimoak",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Bost Euroak Eta Hogeita Bost Zentimoak",
		},
		{
			name:   "Only zentimoak",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Zero Euroak Eta Laurogeita Hemeretzi Zentimoak",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Bat Mila Bi Ehun Hogeita Hamar Lau Euroak Eta Berrogeita Hamasei Zentimoak",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Ehun Hogei Hiru Eta Berrogei Bost",
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

func TestToWords_BasqueSpain_Negative(t *testing.T) {
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
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Minus Berrogeita Hamar",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Minus Hogeita Bost Euroak Eta Hirurogeita Hamabost Zentimoak",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "Negatibo",
					},
				},
			},
			expected: "Negatibo Ehun Euroak",
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

func TestToWords_BasqueSpain_Formatting(t *testing.T) {
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
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "EHUN HOGEI HIRU EUROAK",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "ehun hogei hiru euroak",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Bederatzi Ehun Laurogeita Hamar Bederatzi Soilik",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Bost Ehun Euroak Soilik",
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

func TestToWords_BasqueSpain_CustomCurrency(t *testing.T) {
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
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "peseta",
						Plural: "pesetak",
					},
				},
			},
			expected: "Ehun pesetak",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "dolar",
						Plural:           "dolarrak",
						FractionUnitName: "zentimo",
						FractionPlural:   "zentimoak",
					},
				},
			},
			expected: "Berrogeita Hamar dolarrak Eta Hogeita Bost zentimoak",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "libra",
						Plural: "librak",
					},
				},
			},
			expected: "Bat libra",
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

func TestToWords_BasqueSpain_EdgeCases(t *testing.T) {
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
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Zero Euroak Eta Bat Zentimo",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Hamaika",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Hamabi",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Hogeita Bat",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ehun Bat",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Mila Bat",
		},
		{
			name:   "Basque vigesimal - Eighty",
			amount: 80,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Laurogei",
		},
		{
			name:   "Basque vigesimal - Ninety",
			amount: 90,
			options: &pkg.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Laurogeita Hamar",
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
