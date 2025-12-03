package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_FrenchFrance_Numbers(t *testing.T) {
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
				Locale: "fr-FR",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Zéro",
		},
		{
			name:   "Seventy (Standard French)",
			amount: 70,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Soixante-dix",
		},
		{
			name:   "Ninety (Standard French)",
			amount: 90,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Quatre-vingt-dix",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Un million",
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
