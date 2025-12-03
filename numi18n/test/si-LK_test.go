package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_SinhalaSriLanka_Numbers(t *testing.T) {
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
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ශුන්\u200dය",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "පන්ච",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "පන්චදශ",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ත්\u200dරිංශ",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "චත්වාරිංශ සහ සප්ත",
		},
		{
			name:   "One hundred",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "සියය",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ද්වි සියය පන්චාශ ෂට්",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "දහස",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "මිලියන",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ඒක මිලියන ද්වි සියය ත්\u200dරිංශ චතුර් දහස පන්ච සියය ෂෂ්ටිය සප්ත",
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
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One Sri Lankan rupee",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ඒක ශ්\u200dරී ලංකා රුපියල",
		},
		{
			name:   "Multiple Sri Lankan rupees",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "පන්ච ශ්\u200dරී ලංකා රුපියල්",
		},
		{
			name:   "Zero Sri Lankan rupees",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ශුන්\u200dය ශ්\u200dරී ලංකා රුපියල්",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "මිලියන ශ්\u200dරී ලංකා රුපියල්",
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
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Rupees and one cent",
			amount: 5.01,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "පන්ච ශ්\u200dරී ලංකා රුපියල් සහ ඒක සෙන්ට්",
		},
		{
			name:   "Rupees and multiple cents",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "පන්ච ශ්\u200dරී ලංකා රුපියල් සහ විංශ සහ පන්ච සෙන්ට්",
		},
		{
			name:   "Only cents",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ශුන්\u200dය ශ්\u200dරී ලංකා රුපියල් සහ අනූව සහ නව සෙන්ට්",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ඒක දහස ද්වි සියය ත්\u200dරිංශ චතුර් ශ්\u200dරී ලංකා රුපියල් සහ පන්චාශ සහ ෂට් සෙන්ට්",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ඒක සියය විංශ ත්\u200dරි සහ චත්වාරිංශ සහ පන්ච",
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
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Negative number basic",
			amount: -50,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "අඩු පන්චාශ",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "අඩු විංශ සහ පන්ච ශ්\u200dරී ලංකා රුපියල් සහ සප්තතිය සහ පන්ච සෙන්ට්",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &numi18n.OverrideOptions{
						NegativeWord: "ඍණ",
					},
				},
			},
			expected: "ඍණ සියය ශ්\u200dරී ලංකා රුපියල්",
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
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Uppercase",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "ඒක සියය විංශ ත්\u200dරි ශ්\u200dරී ලංකා රුපියල්",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "ඒක සියය විංශ ත්\u200dරි ශ්\u200dරී ලංකා රුපියල්",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "නව සියය අනූව නව පමණක්",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "පන්ච සියය ශ්\u200dරී ලංකා රුපියල් පමණක්",
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
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Custom currency name",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
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
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
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
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
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
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Very small decimal",
			amount: 0.01,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ශුන්\u200dය ශ්\u200dරී ලංකා රුපියල් සහ ඒක සෙන්ට්",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ඒකාදශ",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ද්වාදශ",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "විංශ සහ ඒක",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ඒක සියය ඒක",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &numi18n.NumI18NOptions{
				Locale: "si-LK",
				WordDetails: &numi18n.WordDetails{
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
