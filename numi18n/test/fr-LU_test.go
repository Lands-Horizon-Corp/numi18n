package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_FrenchLuxembourg_Numbers(t *testing.T) {
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
				Locale: "fr-LU",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Zéro",
		},
		{
			name:   "Seventy",
			amount: 70,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-LU",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Soixante-dix",
		},
		{
			name:   "One hundred",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-LU",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Cent",
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
