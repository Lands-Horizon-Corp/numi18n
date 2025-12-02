package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_WolofSN_Numbers(t *testing.T) {
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
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tus",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Juróom",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Fukk ak juróom",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ñett-fukk",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Juróom-benn ak juróom-ñaar",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Teemer",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ñaar benn teemer juróom-benn-fukk juróom-benn",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Benn junni",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Benn milliyon",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Benn benn milliyon ñaar benn teemer ñett-fukk ñeent benn junni juróom benn teemer juróom-ñaar juróom-ñaar",
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

func TestToWords_WolofSN_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One franc",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Benn franc CFA",
		},
		{
			name:   "Multiple francs",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Juróom franc CFA",
		},
		{
			name:   "Zero francs",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Tus franc CFA",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Benn milliyon franc CFA",
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

func TestToWords_WolofSN_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Francs and one centime",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Juróom franc CFA ak benn centime",
		},
		{
			name:   "Francs and multiple centimes",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Juróom franc CFA ak ñaar-fukk ak juróom centime",
		},
		{
			name:   "Only centimes",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Tus franc CFA ak juróom-ñeent-fukk ak juróom-ñeent centime",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Benn benn junni ñaar benn teemer ñett-fukk ñeent franc CFA ak juróom-benn-fukk ak juróom-benn centime",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Benn benn teemer ñaar-fukk ñett ak juróom-benn ak juróom",
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

func TestToWords_WolofSN_Negative(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Negative number basic",
			amount: -50,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Jiitu juróom-benn-fukk",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Jiitu ñaar-fukk ak juróom franc CFA ak juróom-ñaar-fukk ak juróom centime",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "fibar",
					},
				},
			},
			expected: "Fibar Teemer franc CFA",
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

func TestToWords_WolofSN_Formatting(t *testing.T) {
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
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "BENN BENN TEEMER ÑAAR-FUKK ÑETT FRANC CFA",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "benn benn teemer ñaar-fukk ñett franc cfa",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Juróom-ñeent benn teemer juróom-ñeent-fukk juróom-ñeent rekk",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Juróom teemer franc CFA rekk",
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

func TestToWords_WolofSN_CustomCurrency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Custom currency name",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "dalasi",
						Plural: "dalasi",
					},
				},
			},
			expected: "Teemer dalasi",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "naira",
						Plural:           "naira",
						FractionUnitName: "kobo",
						FractionPlural:   "kobo",
					},
				},
			},
			expected: "Juróom-benn-fukk naira ak ñaar-fukk ak juróom kobo",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "cedi",
						Plural: "cedis",
					},
				},
			},
			expected: "Benn cedi",
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

func TestToWords_WolofSN_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Very small decimal",
			amount: 0.01,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Tus franc CFA ak benn centime",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Fukk ak benn",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Fukk ak ñaar",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ñaar-fukk ak benn",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Benn benn teemer benn",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "wo-SN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Benn benn junni benn",
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
