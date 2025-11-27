package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_Greenlandic_Numbers(t *testing.T) {
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
				Locale: "kl-GL",
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
				Locale: "kl-GL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tallimat",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "kl-GL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Pingasut-qulit",
		},
		{
			name:   "One hundred",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "kl-GL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ataaseq hundrede",
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

func TestToWords_Greenlandic_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One krone",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "kl-GL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Ataaseq Dansk krone",
		},
		{
			name:   "Multiple kroner",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "kl-GL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Tallimat Danske kroner",
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

func TestToWords_Greenlandic_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Kroner and øre",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "kl-GL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Tallimat Danske kroner Aamma Juullut Tallimat Øre",
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

func TestToWords_Greenlandic_Negative(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Negative number",
			amount: -50,
			options: &pkg.NumI18NOptions{
				Locale: "kl-GL",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Minus Tallimat-qulit",
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

func TestToWords_Greenlandic_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Twenty",
			amount: 20,
			options: &pkg.NumI18NOptions{
				Locale: "kl-GL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Juullut",
		},
		{
			name:   "Ten",
			amount: 10,
			options: &pkg.NumI18NOptions{
				Locale: "kl-GL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Qulit",
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
