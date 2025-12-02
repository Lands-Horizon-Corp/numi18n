package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_PortugueseBrazil_Numbers(t *testing.T) {
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
				Locale: "pt-BR",
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
				Locale: "pt-BR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Cinco",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "pt-BR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Quinze",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "pt-BR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Trinta",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "pt-BR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Quarenta e sete",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "pt-BR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Cem",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "pt-BR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Dois cem cinquenta seis",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "pt-BR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Mil",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "pt-BR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Um milhão",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "pt-BR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Um um milhão dois cem trinta quatro mil cinco cem sessenta sete",
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

func TestToWords_PortugueseBrazil_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One real",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "pt-BR",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Um Real",
		},
		{
			name:   "Multiple reais",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "pt-BR",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Cinco Reais",
		},
		{
			name:   "Zero reais",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "pt-BR",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Zero Reais",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "pt-BR",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Um milhão Reais",
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

func TestToWords_PortugueseBrazil_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "pt-BR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Onze",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "pt-BR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Doze",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "pt-BR",
				WordDetails: &pkg.WordDetails{
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
