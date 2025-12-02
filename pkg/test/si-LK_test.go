package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_SinhalaSriLanka_Numbers(t *testing.T) {
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
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ශුන්‍ය",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "පන්ච",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "පන්චදශ",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ත්‍රිංශ",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "චත්වාරිංශ සහ සප්ත",
		},
		{
			name:   "One hundred",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "සියය",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ද්වි සියය පන්චාශ ෂට්",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "දහස",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "මිලියන",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ඒක මිලියන ද්වි සියය ත්‍රිංශ චතුර් දහස පන්ච සියය ෂෂ්ටිය සප්ත",
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

func TestToWords_SinhalaSriLanka_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One Sri Lankan rupee",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ඒක ශ්‍රී ලංකා රුපියල",
		},
		{
			name:   "Multiple Sri Lankan rupees",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "පන්ච ශ්‍රී ලංකා රුපියල්",
		},
		{
			name:   "Zero Sri Lankan rupees",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ශුන්‍ය ශ්‍රී ලංකා රුපියල්",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "මිලියන ශ්‍රී ලංකා රුපියල්",
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

func TestToWords_SinhalaSriLanka_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Rupees and one cent",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "පන්ච ශ්‍රී ලංකා රුපියල් සහ ඒක සෙන්ට්",
		},
		{
			name:   "Rupees and multiple cents",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "පන්ච ශ්‍රී ලංකා රුපියල් සහ විංශ සහ පන්ච සෙන්ට්",
		},
		{
			name:   "Only cents",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ශුන්‍ය ශ්‍රී ලංකා රුපියල් සහ අනූව සහ නව සෙන්ට්",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ඒක දහස ද්වි සියය ත්‍රිංශ චතුර් ශ්‍රී ලංකා රුපියල් සහ පන්චාශ සහ ෂට් සෙන්ට්",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ඒක සියය විංශ ත්‍රි සහ චත්වාරිංශ සහ පන්ච",
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

func TestToWords_SinhalaSriLanka_Negative(t *testing.T) {
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
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "අඩු පන්චාශ",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "අඩු විංශ සහ පන්ච ශ්‍රී ලංකා රුපියල් සහ සප්තතිය සහ පන්ච සෙන්ට්",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "ඍණ",
					},
				},
			},
			expected: "ඍණ සියය ශ්‍රී ලංකා රුපියල්",
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

func TestToWords_SinhalaSriLanka_Formatting(t *testing.T) {
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
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "ඒක සියය විංශ ත්‍රි ශ්‍රී ලංකා රුපියල්",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "ඒක සියය විංශ ත්‍රි ශ්‍රී ලංකා රුපියල්",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "නව සියය අනූව නව පමණක්",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "පන්ච සියය ශ්‍රී ලංකා රුපියල් පමණක්",
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

func TestToWords_SinhalaSriLanka_CustomCurrency(t *testing.T) {
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
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "ඩොලර්",
						Plural: "ඩොලර්",
					},
				},
			},
			expected: "සියය ඩොලර්",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "යුරෝ",
						Plural:           "යුරෝ",
						FractionUnitName: "සෙන්ට්",
						FractionPlural:   "සෙන්ට්",
					},
				},
			},
			expected: "පන්චාශ යුරෝ සහ විංශ සහ පන්ච සෙන්ට්",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "පවුන්ඩ්",
						Plural: "පවුන්ඩ්",
					},
				},
			},
			expected: "ඒක පවුන්ඩ්",
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

func TestToWords_SinhalaSriLanka_EdgeCases(t *testing.T) {
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
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ශුන්‍ය ශ්‍රී ලංකා රුපියල් සහ ඒක සෙන්ට්",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ඒකාදශ",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ද්වාදශ",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "විංශ සහ ඒක",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ඒක සියය ඒක",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ඒක දහස ඒක",
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
