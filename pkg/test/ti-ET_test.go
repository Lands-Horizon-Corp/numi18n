package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_TigrayEthiopia_Numbers(t *testing.T) {
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
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ዜሮ",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ሓምሽተ",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ሓምሽተ ዓሰር",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ሰለስተ ሰላሳ",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ኣርባዕተ ሰላሳ ሸውዓተ",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ሓደ መቶ",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ክልተ ሓደ መቶ ሓምሽተ ሰላሳ ሽድሽተ",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ሓደ ሽሕ",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ሓደ ሚልዮን",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ሓደ ሓደ ሚልዮን ክልተ ሓደ መቶ ሰለስተ ሰላሳ ኣርባዕተ ሓደ ሽሕ ሓምሽተ ሓደ መቶ ሽድሽተ ሰላሳ ሸውዓተ",
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

func TestToWords_TigrayEthiopia_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One birr",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ሓደ ብር",
		},
		{
			name:   "Multiple birrs",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ሓምሽተ ብር",
		},
		{
			name:   "Zero birrs",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ዜሮ ብር",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ሓደ ሚልዮን ብር",
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

func TestToWords_TigrayEthiopia_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Birrs and one santim",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ሓምሽተ ብር ከምኡ'ውን ሓደ ሳንቲም",
		},
		{
			name:   "Birrs and multiple santims",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ሓምሽተ ብር ከምኡ'ውን ክልተ ሰላሳ ሓምሽተ ሳንቲም",
		},
		{
			name:   "Only santims",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ዜሮ ብር ከምኡ'ውን ተስዓተ ሰላሳ ተስዓተ ሳንቲም",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ሓደ ሓደ ሽሕ ክልተ ሓደ መቶ ሰለስተ ሰላሳ ኣርባዕተ ብር ከምኡ'ውን ሓምሽተ ሰላሳ ሽድሽተ ሳንቲም",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ሓደ ሓደ መቶ ዕስራ ሰለስተ ከምኡ'ውን ኣርባዕተ ሰላሳ ሓምሽተ",
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

func TestToWords_TigrayEthiopia_Negative(t *testing.T) {
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
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ጉሓፍ ሓምሽተ ሰላሳ",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ጉሓፍ ክልተ ሰላሳ ሓምሽተ ብር ከምኡ'ውን ሸውዓተ ሰላሳ ሓምሽተ ሳንቲም",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "ዋላሓሓሳብ",
					},
				},
			},
			expected: "ዋላሓሓሳብ ሓደ መቶ ብር",
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

func TestToWords_TigrayEthiopia_Formatting(t *testing.T) {
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
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "ሓደ ሓደ መቶ ዕስራ ሰለስተ ብር",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "ሓደ ሓደ መቶ ዕስራ ሰለስተ ብር",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "ተስዓተ ሓደ መቶ ተስዓተ ሰላሳ ተስዓተ ጥራይ",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "ሓምሽተ መቶ ብር ጥራይ",
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

func TestToWords_TigrayEthiopia_CustomCurrency(t *testing.T) {
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
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "ዶላር",
						Plural: "ዶላር",
					},
				},
			},
			expected: "ሓደ መቶ ዶላር",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "ዩሮ",
						Plural:           "ዩሮ",
						FractionUnitName: "ሳንት",
						FractionPlural:   "ሳንት",
					},
				},
			},
			expected: "ሓምሽተ ሰላሳ ዩሮ ከምኡ'ውን ክልተ ሰላሳ ሓምሽተ ሳንት",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "ፓውንድ",
						Plural: "ፓውንድ",
					},
				},
			},
			expected: "ሓደ ፓውንድ",
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

func TestToWords_TigrayEthiopia_EdgeCases(t *testing.T) {
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
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ዜሮ ብር ከምኡ'ውን ሓደ ሳንቲም",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ሓዳሽ ዓሰር",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ክልተ ዓሰር",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ክልተ ሰላሳ ሓደ",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ሓደ ሓደ መቶ ሓደ",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "ti-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ሓደ ሓደ ሽሕ ሓደ",
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
