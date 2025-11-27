package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_FrenchSwitzerland_Numbers(t *testing.T) {
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
				Locale: "fr-CH",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Zéro",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "fr-CH",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Cinq",
		},
		{
			name:   "Seventy (Swiss French: Septante)",
			amount: 70,
			options: &pkg.NumI18NOptions{
				Locale: "fr-CH",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Septante",
		},
		{
			name:   "Ninety (Swiss French: Nonante)",
			amount: 90,
			options: &pkg.NumI18NOptions{
				Locale: "fr-CH",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Nonante",
		},
		{
			name:   "One hundred",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "fr-CH",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Cent",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "fr-CH",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Mille",
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
