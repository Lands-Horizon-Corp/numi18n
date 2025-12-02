package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_RomanianRomania_Numbers(t *testing.T) {
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
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Zero",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Cinci",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Cincisprezece",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Treizeci",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Patruzeci și șapte",
		},
		{
			name:   "One hundred",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "O sută",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "O mie",
		},
		{
			name:   "Large number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Unu un milion doi o sută treizeci patru o mie cinci o sută șaizeci șapte",
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

func TestToWords_RomanianRomania_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One leu",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Unu Leu românesc",
		},
		{
			name:   "Multiple lei",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Cinci Lei românești",
		},
		{
			name:   "Zero lei",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Zero Lei românești",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Un milion Lei românești",
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

func TestToWords_RomanianRomania_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Lei with bani",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Cinci Lei românești și douăzeci și cinci Bani",
		},
		{
			name:   "Only bani",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Zero Lei românești și nouăzeci și nouă Bani",
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

func TestToWords_RomanianRomania_Negative(t *testing.T) {
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
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Minus cincizeci",
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

func TestToWords_RomanianRomania_Formatting(t *testing.T) {
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
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "UNU O SUTĂ DOUĂZECI TREI LEI ROMÂNEȘTI",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "unu o sută douăzeci trei lei românești",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Nouă o sută nouăzeci nouă numai",
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

func TestToWords_RomanianRomania_CustomCurrency(t *testing.T) {
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
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "euro",
						Plural: "euro",
					},
				},
			},
			expected: "O sută euro",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "euro",
						Plural:           "euro",
						FractionUnitName: "cent",
						FractionPlural:   "cenți",
					},
				},
			},
			expected: "Cincizeci euro și douăzeci și cinci cenți",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "dolar",
						Plural: "dolari",
					},
				},
			},
			expected: "Unu dolar",
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

func TestToWords_RomanianRomania_EdgeCases(t *testing.T) {
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
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Zero Lei românești și unu Ban",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Unsprezece",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Doisprezece",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Douăzeci și unu",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Unu o sută unu",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Unu o mie unu",
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
