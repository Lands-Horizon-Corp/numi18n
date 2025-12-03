package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_ItalianSwiss_Numbers(t *testing.T) {
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
				Locale: "it-CH",
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
				Locale: "it-CH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Cinque",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "it-CH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Trenta",
		},
		{
			name:   "One hundred",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "it-CH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Cento",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "it-CH",
				WordDetails: &numi18n.WordDetails{
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

func TestToWords_ItalianSwiss_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One franc",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "it-CH",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Uno Franco svizzero",
		},
		{
			name:   "Multiple francs",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "it-CH",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Cinque Franchi svizzeri",
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
