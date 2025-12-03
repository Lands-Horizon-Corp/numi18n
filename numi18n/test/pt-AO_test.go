package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_PortugueseAngola_Numbers(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Zero",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-AO",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Zero",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-AO",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Cinco",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-AO",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Quinze",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-AO",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Trinta",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-AO",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Quarenta e sete",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-AO",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Cem",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-AO",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Dois cem cinquenta seis",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-AO",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Mil",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-AO",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Milhão",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-AO",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Um milhão dois cem trinta quatro mil cinco cem sessenta sete",
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

func TestToWords_PortugueseAngola_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One kwanza",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-AO",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Um Kwanza",
		},
		{
			name:   "Multiple kwanzas",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-AO",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Cinco Kwanzas",
		},
		{
			name:   "Zero kwanzas",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-AO",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Zero Kwanzas",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-AO",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Milhão Kwanzas",
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

func TestToWords_PortugueseAngola_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-AO",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Onze",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-AO",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Doze",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-AO",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Vinte e um",
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
