package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_EnglishZW_Numbers(t *testing.T) {
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
				Locale: "en-ZW",
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
				Locale: "en-ZW",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Five",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "en-ZW",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "One Thousand",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "en-ZW",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "One Million",
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

func TestToWords_EnglishZW_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One Zimbabwean dollar",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "en-ZW",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "One Zimbabwean Dollar",
		},
		{
			name:   "Multiple Zimbabwean dollars",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "en-ZW",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Five Zimbabwean Dollars",
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

func TestToWords_EnglishZW_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Zimbabwean dollars and cents",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "en-ZW",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Five Zimbabwean Dollars And Twenty Five Cents",
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
