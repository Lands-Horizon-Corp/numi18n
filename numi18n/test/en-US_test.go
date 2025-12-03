package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_EnglishUS_Numbers(t *testing.T) {
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
				Locale: "en-US",
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
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Five",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Fifteen",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Thirty",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Forty Seven",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "One Hundred",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Two Hundred Fifty Six",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "One Thousand",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "One Million",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
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

func TestToWords_EnglishUS_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One dollar",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "One Dollar",
		},
		{
			name:   "Multiple dollars",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Five Dollars",
		},
		{
			name:   "Zero dollars",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Zero Dollars",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "One Million Dollars",
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

func TestToWords_EnglishUS_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Dollars and one cent",
			amount: 5.01,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Five Dollars And One Cent",
		},
		{
			name:   "Dollars and multiple cents",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Five Dollars And Twenty Five Cents",
		},
		{
			name:   "Only cents",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Zero Dollars And Ninety Nine Cents",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "One Thousand Two Hundred Thirty Four Dollars And Fifty Six Cents",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "One Hundred Twenty Three And Forty Five",
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

func TestToWords_EnglishUS_Negative(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Negative number basic",
			amount: -50,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Minus Fifty",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Minus Twenty Five Dollars And Seventy Five Cents",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &numi18n.OverrideOptions{
						NegativeWord: "Negative",
					},
				},
			},
			expected: "Negative One Hundred Dollars",
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

func TestToWords_EnglishUS_Formatting(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Uppercase",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "ONE HUNDRED TWENTY THREE DOLLARS",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "one hundred twenty three dollars",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Nine Hundred Ninety Nine Only",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Five Hundred Dollars Only",
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

func TestToWords_EnglishUS_CustomCurrency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Custom currency name",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "peso",
						Plural: "pesos",
					},
				},
			},
			expected: "One Hundred pesos",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "euro",
						Plural:           "euros",
						FractionUnitName: "cent",
						FractionPlural:   "cents",
					},
				},
			},
			expected: "Fifty euros And Twenty Five cents",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "pound",
						Plural: "pounds",
					},
				},
			},
			expected: "One pound",
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

func TestToWords_EnglishUS_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Very small decimal",
			amount: 0.01,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Zero Dollars And One Cent",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Eleven",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Twelve",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Twenty One",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "One Hundred One",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
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

// ========== ToFormat Tests ==========

func TestToFormat_EnglishUS_Numbers(t *testing.T) {
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
				Locale: "en-US",
			},
			expected: "0",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "5",
		},
		{
			name:   "Two digits",
			amount: 25,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "25",
		},
		{
			name:   "Three digits",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "100",
		},
		{
			name:   "Four digits with comma",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "1,000",
		},
		{
			name:   "Five digits with comma",
			amount: 12345,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "12,345",
		},
		{
			name:   "Six digits with commas",
			amount: 123456,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "123,456",
		},
		{
			name:   "Seven digits with commas",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "1,234,567",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "1,000,000",
		},
		{
			name:   "Large number with commas",
			amount: 1234567890,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "1,234,567,890",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expected {
				t.Errorf("ToFormat(%f) = %q, expected %q", tt.amount, result, tt.expected)
			}
		})
	}
}

func TestToFormat_EnglishUS_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Decimal with one place",
			amount: 5.5,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "5.5",
		},
		{
			name:   "Decimal with two places",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "5.25",
		},
		{
			name:   "Decimal with multiple places",
			amount: 123.456789,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "123.456789",
		},
		{
			name:   "Large number with decimals",
			amount: 1234567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "1,234,567.89",
		},
		{
			name:   "Small decimal",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "0.99",
		},
		{
			name:   "Complex decimal with commas",
			amount: 12345678.123456,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "12,345,678.123456",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expected {
				t.Errorf("ToFormat(%f) = %q, expected %q", tt.amount, result, tt.expected)
			}
		})
	}
}

func TestToFormat_EnglishUS_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Simple currency",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "$1,234.56",
		},
		{
			name:   "Large currency amount",
			amount: 1234567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "$1,234,567.89",
		},
		{
			name:   "Negative currency",
			amount: -987.65,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "-$987.65",
		},
		{
			name:   "Override currency symbol",
			amount: 1000.50,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Symbol: "€",
					},
				},
			},
			expected: "€1,000.50",
		},
		{
			name:   "Zero with currency",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "$0.00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expected {
				t.Errorf("ToFormat(%f) = %q, expected %q", tt.amount, result, tt.expected)
			}
		})
	}
}

func TestToFormat_EnglishUS_Negative(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Negative single digit",
			amount: -5,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "-5",
		},
		{
			name:   "Negative with commas",
			amount: -1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "-1,234,567",
		},
		{
			name:   "Negative with decimals",
			amount: -1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "-1,234.56",
		},
		{
			name:   "Negative decimal only",
			amount: -0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "-0.99",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expected {
				t.Errorf("ToFormat(%f) = %q, expected %q", tt.amount, result, tt.expected)
			}
		})
	}
}

func TestToFormat_EnglishUS_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Very small positive number",
			amount: 0.001,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "0.001",
		},
		{
			name:   "Very small negative number",
			amount: -0.001,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "-0.001",
		},
		{
			name:   "Very large number",
			amount: 999999999999.99,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "999,999,999,999.99",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expected {
				t.Errorf("ToFormat(%f) = %q, expected %q", tt.amount, result, tt.expected)
			}
		})
	}
}
