package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_DutchNetherlands_Numbers(t *testing.T) {
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
				Locale: "nl-NL",
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
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Five",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Fifteen",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Thirty",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Forty Seven",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "One Hundred",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Two Hundred Fifty Six",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "One Thousand",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "One Million",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven",
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

func TestToWords_DutchNetherlands_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One euro",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "One Euro",
		},
		{
			name:   "Multiple euros",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Five Euro",
		},
		{
			name:   "Zero euros",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Zero Euro",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "One Million Euro",
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

func TestToWords_DutchNetherlands_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Euros and one cent",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Five Euro en One Cent",
		},
		{
			name:   "Euros and multiple cents",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Five Euro en Twenty Five Cent",
		},
		{
			name:   "Only cents",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Zero Euro en Ninety Nine Cent",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "One Thousand Two Hundred Thirty Four Euro en Fifty Six Cent",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "One Hundred Twenty Three en Forty Five",
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

func TestToWords_DutchNetherlands_Negative(t *testing.T) {
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
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Min Fifty",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Min Twenty Five Euro en Seventy Five Cent",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "Negatief",
					},
				},
			},
			expected: "Negatief One Hundred Euro",
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

func TestToWords_DutchNetherlands_Formatting(t *testing.T) {
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
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "ONE HUNDRED TWENTY THREE EURO",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "one hundred twenty three euro",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Nine Hundred Ninety Nine alleen",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Five Hundred Euro alleen",
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

func TestToWords_DutchNetherlands_CustomCurrency(t *testing.T) {
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
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "gulden",
						Plural: "gulden",
					},
				},
			},
			expected: "One Hundred gulden",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "dollar",
						Plural:           "dollars",
						FractionUnitName: "cent",
						FractionPlural:   "cent",
					},
				},
			},
			expected: "Fifty dollars en Twenty Five cent",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "pond",
						Plural: "pond",
					},
				},
			},
			expected: "One pond",
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

func TestToWords_DutchNetherlands_EdgeCases(t *testing.T) {
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
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Zero Euro en One Cent",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Eleven",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Twelve",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Twenty One",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "One Hundred One",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "One Thousand One",
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
