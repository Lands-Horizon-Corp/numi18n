package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_MāoriNewZealand_Numbers(t *testing.T) {
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
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Kore",
		},
		{
			name:   "One",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Kotahi",
		},
		{
			name:   "Five",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Rima",
		},
		{
			name:   "Ten",
			amount: 10,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tekau",
		},
		{
			name:   "Eleven",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tekau mā kotahi",
		},
		{
			name:   "Twenty",
			amount: 20,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Rua tekau",
		},
		{
			name:   "Twenty-one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Rua tekau mā kotahi",
		},
		{
			name:   "Fifty",
			amount: 50,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Rima tekau",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Kotahi rau",
		},
		{
			name:   "Five hundred",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Rima rau",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Kotahi mano",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Kotahi miriona",
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

func TestToWords_MāoriNewZealand_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Zero tāra",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Kore Tāra",
		},
		{
			name:   "One tāra",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Kotahi Tāra",
		},
		{
			name:   "Five tāra",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Rima Tāra",
		},
		{
			name:   "Twenty tāra",
			amount: 20,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Rua tekau Tāra",
		},
		{
			name:   "One hundred tāra",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Kotahi rau Tāra",
		},
		{
			name:   "One thousand tāra",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Kotahi mano Tāra",
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

func TestToWords_MāoriNewZealand_DecimalCurrency(t *testing.T) {
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
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Decimal:    true,
					Currency:   true,
				},
			},
			expected: "Kotahi Tāra me rua tekau Hēneti",
		},
		{
			name:   "Twenty point fifty",
			amount: 20.50,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Decimal:    true,
					Currency:   true,
				},
			},
			expected: "Rua tekau Tāra me rima tekau Hēneti",
		},
		{
			name:   "One hundred point ninety-nine",
			amount: 100.99,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Decimal:    true,
					Currency:   true,
				},
			},
			expected: "Kotahi rau Tāra me iwa tekau mā iwa Hēneti",
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

func TestToWords_MāoriNewZealand_NegativeNumbers(t *testing.T) {
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
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "Rēhia kotahi",
		},
		{
			name:   "Negative twenty",
			amount: -20,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "Rēhia rua tekau",
		},
		{
			name:   "Negative one hundred",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "Rēhia Kotahi rau",
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

func TestToWords_MāoriNewZealand_NegativeCurrency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Negative one tāra",
			amount: -1,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
					Currency:     true,
				},
			},
			expected: "Rēhia kotahi Tāra",
		},
		{
			name:   "Negative twenty point fifty",
			amount: -20.50,
			options: &numi18n.NumI18NOptions{
				Locale: "mi-NZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
					Decimal:      true,
					Currency:     true,
				},
			},
			expected: "Rēhia rua tekau Tāra me rima tekau Hēneti",
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
