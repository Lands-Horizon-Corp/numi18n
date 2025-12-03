package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_MalagasyMadagascar_Numbers(t *testing.T) {
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
				Locale: "mg-MG",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Aotra",
		},
		{
			name:   "One",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "mg-MG",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Iraika",
		},
		{
			name:   "Five",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "mg-MG",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Dimy",
		},
		{
			name:   "Ten",
			amount: 10,
			options: &numi18n.NumI18NOptions{
				Locale: "mg-MG",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Folo",
		},
		{
			name:   "Eleven",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "mg-MG",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Iraika ambin'ny folo",
		},
		{
			name:   "Twenty",
			amount: 20,
			options: &numi18n.NumI18NOptions{
				Locale: "mg-MG",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Roapolo",
		},
		{
			name:   "Twenty-one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "mg-MG",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Iraika ambin'ny roapolo",
		},
		{
			name:   "Fifty",
			amount: 50,
			options: &numi18n.NumI18NOptions{
				Locale: "mg-MG",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Dimampolo",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "mg-MG",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Zato iray",
		},
		{
			name:   "Five hundred",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "mg-MG",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Dimy zato",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "mg-MG",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Arivo",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "mg-MG",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tapitrisa kely",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToWords(tt.amount)
			if result != tt.expected {
				t.Errorf("ToWords(%f) = %q, want %q", tt.amount, result, tt.expected)
			}
		})
	}
}

func TestToWords_MalagasyMadagascar_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Zero ariary",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "mg-MG",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Aotra Ariary",
		},
		{
			name:   "One ariary",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "mg-MG",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Iraika Ariary",
		},
		{
			name:   "Five ariary",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "mg-MG",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Dimy Ariary",
		},
		{
			name:   "One hundred ariary",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "mg-MG",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Zato iray Ariary",
		},
		{
			name:   "One thousand ariary",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "mg-MG",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Arivo Ariary",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToWords(tt.amount)
			if result != tt.expected {
				t.Errorf("ToWords(%f) = %q, want %q", tt.amount, result, tt.expected)
			}
		})
	}
}

func TestToWords_MalagasyMadagascar_DecimalCurrency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One point twenty",
			amount: 1.20,
			options: &numi18n.NumI18NOptions{
				Locale: "mg-MG",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Decimal:    true,
					Currency:   true,
				},
			},
			expected: "Iraika Ariary ary roapolo Iraimbilanja",
		},
		{
			name:   "Twenty point fifty",
			amount: 20.50,
			options: &numi18n.NumI18NOptions{
				Locale: "mg-MG",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Decimal:    true,
					Currency:   true,
				},
			},
			expected: "Roapolo Ariary ary dimampolo Iraimbilanja",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToWords(tt.amount)
			if result != tt.expected {
				t.Errorf("ToWords(%f) = %q, want %q", tt.amount, result, tt.expected)
			}
		})
	}
}

func TestToWords_MalagasyMadagascar_NegativeNumbers(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Negative one",
			amount: -1,
			options: &numi18n.NumI18NOptions{
				Locale: "mg-MG",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "Miiba iraika",
		},
		{
			name:   "Negative twenty",
			amount: -20,
			options: &numi18n.NumI18NOptions{
				Locale: "mg-MG",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "Miiba roapolo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToWords(tt.amount)
			if result != tt.expected {
				t.Errorf("ToWords(%f) = %q, want %q", tt.amount, result, tt.expected)
			}
		})
	}
}
