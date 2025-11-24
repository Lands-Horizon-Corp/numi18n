package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_AfrikaansZA_Numbers(t *testing.T) {
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
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Nul",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Vyf",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Vyftien",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Dertig",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Sewe En Veertig",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Honderd",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Twee Honderd Ses En Vyftig",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Een Duizend",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Een Miljoen",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Een Miljoen Twee Honderd Vier En Dertig Duizend Vyf Honderd Sewe En Sestig",
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

func TestToWords_AfrikaansZA_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One rand",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Een Rand",
		},
		{
			name:   "Multiple rande",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Vyf Rande",
		},
		{
			name:   "Zero rande",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Nul Rande",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Een Miljoen Rande",
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

func TestToWords_AfrikaansZA_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Rande and one sent",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Vyf Rande En Een Sent",
		},
		{
			name:   "Rande and multiple sente",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Vyf Rande En Vyf En Twintig Sente",
		},
		{
			name:   "Only sente",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Nul Rande En Nege En Negentig Sente",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Een Duizend Twee Honderd Vier En Dertig Rande En Ses En Vyftig Sente",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Honderd Drie En Twintig En Vyf En Veertig",
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

func TestToWords_AfrikaansZA_Negative(t *testing.T) {
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
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Min Vyftig",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Min Vyf En Twintig Rande En Vyf En Sewentig Sente",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "Negatief",
					},
				},
			},
			expected: "Negatief Honderd Rande",
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

func TestToWords_AfrikaansZA_Formatting(t *testing.T) {
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
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "HONDERD DRIE EN TWINTIG RANDE",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "honderd drie en twintig rande",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Nege Honderd Nege En Negentig Slegs",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Vyf Honderd Rande Slegs",
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

func TestToWords_AfrikaansZA_EdgeCases(t *testing.T) {
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
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Nul Rande En Een Sent",
		},
		{
			name:   "Eleven (elf)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Elf",
		},
		{
			name:   "Twelve (twaalf)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Twaalf",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Een En Twintig",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Honderd Een",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Een Duizend Een",
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
