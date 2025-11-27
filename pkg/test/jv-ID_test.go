package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_Javanese_Numbers(t *testing.T) {
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
				Locale: "jv-ID",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Kosong",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "jv-ID",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Lima",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "jv-ID",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Limolas",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "jv-ID",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Telung Puluh",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "jv-ID",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Patang Puluh Pitu",
		},
		{
			name:   "One hundred",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "jv-ID",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Satus",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "jv-ID",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Siji Ewu",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "jv-ID",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Siji Yuta",
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

func TestToWords_Javanese_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One rupiah",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "jv-ID",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Siji Rupiah",
		},
		{
			name:   "Multiple rupiah",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "jv-ID",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Lima Rupiah",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "jv-ID",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Siji Yuta Rupiah",
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

func TestToWords_Javanese_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Rupiah and sen",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "jv-ID",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Lima Rupiah Lan Rongpuluh Lima Sen",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "jv-ID",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Siji Satus Rongpuluh Telu Lan Patang Puluh Lima",
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

func TestToWords_Javanese_Negative(t *testing.T) {
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
				Locale: "jv-ID",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Minus Seket",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "jv-ID",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Minus Rongpuluh Lima Rupiah Lan Pitung Puluh Lima Sen",
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

func TestToWords_Javanese_Formatting(t *testing.T) {
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
				Locale: "jv-ID",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "SIJI SATUS RONGPULUH TELU RUPIAH",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "jv-ID",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "siji satus rongpuluh telu rupiah",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "jv-ID",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Sanga Satus Sangang Puluh Sanga Mung",
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

func TestToWords_Javanese_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Eleven",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "jv-ID",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Sewelas",
		},
		{
			name:   "Twelve",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "jv-ID",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Rolas",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "jv-ID",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Rongpuluh Siji",
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
