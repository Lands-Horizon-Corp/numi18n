package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_EnglishUS(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Basic whole number",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "One Hundred Twenty Three Dollars",
		},
		{
			name:   "Number with decimals",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "One Hundred Twenty Three Dollars And Forty Five Cents",
		},
		{
			name:   "Negative number",
			amount: -50.25,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Minus Fifty Dollars And Twenty Five Cents",
		},
		{
			name:   "Zero amount",
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
			name:   "With override options",
			amount: 100.50,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "peso",
						Plural:           "pesos",
						FractionUnitName: "centavo",
						FractionPlural:   "centavos",
					},
				},
			},
			expected: "One Hundred pesos And Fifty centavos",
		},
		{
			name:   "Only words without currency",
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

func TestToWords_Japanese(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Basic whole number",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "百二十三円",
		},
		{
			name:   "Number with decimals",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "百二十三円と四十五銭",
		},
		{
			name:   "Large number",
			amount: 10000,
			options: &numi18n.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "一万円",
		},
		{
			name:   "Negative number",
			amount: -50,
			options: &numi18n.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
				},
			},
			expected: "マイナス五十円",
		},
		{
			name:   "Zero amount",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "零円",
		},
		{
			name:   "With override options",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "ドル",
						Plural: "ドル",
					},
				},
			},
			expected: "千ドル",
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

func TestToWords_Filipino(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Basic whole number",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Isang daan Dalawampu Tatlo Piso",
		},
		{
			name:   "Number with decimals",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "Isang daan Dalawampu Tatlo Piso at Apatnapu Lima Sentimo",
		},
		{
			name:   "Single peso",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Isa Peso",
		},
		{
			name:   "Negative number",
			amount: -50.25,
			options: &numi18n.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
				},
			},
			expected: "minus Limampu Piso at Dalawampu Lima Sentimo",
		},
		{
			name:   "Zero amount",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Zero Piso",
		},
		{
			name:   "With only flag",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Only:     true,
				},
			},
			expected: "Lima Daan Piso lamang",
		},
		{
			name:   "With override options",
			amount: 100.50,
			options: &numi18n.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "dolyar",
						Plural:           "mga dolyar",
						FractionUnitName: "sentimo",
						FractionPlural:   "mga sentimo",
					},
				},
			},
			expected: "Isang daan mga dolyar at Limampu mga sentimo",
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

func TestToWords_EdgeCases(t *testing.T) {
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
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "Zero Dollars And One Cent",
		},
		{
			name:   "Large number",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "One Million Dollars",
		},
		{
			name:   "No locale fallback to US",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "One Hundred Dollars",
		},
		{
			name:   "Currency matching fallback",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Currency: "USD",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "One Hundred Dollars",
		},
		{
			name:   "Uppercase formatting",
			amount: 50,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "FIFTY DOLLARS",
		},
		{
			name:   "Lowercase formatting",
			amount: 50,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "fifty dollars",
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

func TestToWords_OverrideOptions(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Custom negative word",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
					OverrideOptions: &numi18n.OverrideOptions{
						NegativeWord: "negative",
					},
				},
			},
			expected: "negative One Hundred Dollars",
		},
		{
			name:   "All overrides combined",
			amount: -123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "euro",
						Plural:           "euros",
						FractionUnitName: "cent",
						FractionPlural:   "cents",
						NegativeWord:     "negative",
					},
				},
			},
			expected: "Negative One Hundred Twenty Three euros And Forty Five cents",
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

func TestToFormat_EnglishUS(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Basic whole number",
			amount: 1234,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "1,234",
		},
		{
			name:   "Number with decimals",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "1,234.56",
		},
		{
			name:   "Basic currency",
			amount: 1234,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "$1,234.00",
		},
		{
			name:   "Currency with decimals",
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
			name:   "Negative currency",
			amount: -1000,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "-$1,000.00",
		},
		{
			name:   "Override currency symbol",
			amount: 5000,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Symbol: "€",
					},
				},
			},
			expected: "€5,000.00",
		},
		{
			name:   "Zero amount",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "0",
		},
		{
			name:   "Zero currency",
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

func TestToFormat_Japanese(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Basic whole number (no separators)",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "1234567",
		},
		{
			name:   "Number with decimals (no separators)",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "1234.56",
		},
		{
			name:   "Basic currency",
			amount: 1234,
			options: &numi18n.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "¥1234",
		},
		{
			name:   "Currency with decimals (rounded)",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "¥1235",
		},
		{
			name:   "Negative currency",
			amount: -1000,
			options: &numi18n.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "-¥1000",
		},
		{
			name:   "Override currency symbol",
			amount: 5000,
			options: &numi18n.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Symbol: "$",
					},
				},
			},
			expected: "$5000",
		},
		{
			name:   "Large number (no separators)",
			amount: 123456789,
			options: &numi18n.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "123456789",
		},
		{
			name:   "Zero currency",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "¥0",
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

func TestToFormat_Filipino(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Basic whole number with separators",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "ph-PH",
			},
			expected: "1,234,567",
		},
		{
			name:   "Number with decimals and separators",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "ph-PH",
			},
			expected: "1,234.56",
		},
		{
			name:   "Basic currency",
			amount: 1234,
			options: &numi18n.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "₱1,234.00",
		},
		{
			name:   "Currency with decimals",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "₱1,234.56",
		},
		{
			name:   "Negative currency",
			amount: -1000,
			options: &numi18n.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "-₱1,000.00",
		},
		{
			name:   "Override currency symbol",
			amount: 5000,
			options: &numi18n.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Symbol: "$",
					},
				},
			},
			expected: "$5,000.00",
		},
		{
			name:   "Small number without separators",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "ph-PH",
			},
			expected: "123",
		},
		{
			name:   "Zero currency",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "₱0.00",
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

func TestToFormat_OverrideOptions(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Override symbol only",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Symbol: "€",
					},
				},
			},
			expected: "€1,000.00",
		},
		{
			name:   "Override symbol with Japanese locale",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Symbol: "USD",
					},
				},
			},
			expected: "USD1000",
		},
		{
			name:   "Override symbol with Philippines locale",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Symbol: "¥",
					},
				},
			},
			expected: "¥1,000.00",
		},
		{
			name:   "Complex currency override with custom symbol",
			amount: 12345.67,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Symbol: "BTC",
					},
				},
			},
			expected: "BTC12,345.67",
		},
		{
			name:   "Negative with override symbol",
			amount: -500,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Symbol: "£",
					},
				},
			},
			expected: "-£500.00",
		},
		{
			name:   "Zero with override symbol",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Symbol: "₹",
					},
				},
			},
			expected: "₹0.00",
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

func TestToFormat_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Very small decimal",
			amount: 0.001,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "0.001",
		},
		{
			name:   "Very large number US",
			amount: 999999999999.99,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "999,999,999,999.99",
		},
		{
			name:   "Very large number Japan (no separators)",
			amount: 999999999999.99,
			options: &numi18n.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "999999999999.99",
		},
		{
			name:   "Very large number Philippines",
			amount: 999999999999.99,
			options: &numi18n.NumI18NOptions{
				Locale: "ph-PH",
			},
			expected: "999,999,999,999.99",
		},
		{
			name:   "Many decimal places",
			amount: 123.456789,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expected: "123.456789",
		},
		{
			name:   "Currency with many decimals (US rounds to 2)",
			amount: 123.456789,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "$123.46",
		},
		{
			name:   "Currency with many decimals (Japan rounds to integer)",
			amount: 123.456789,
			options: &numi18n.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "¥123",
		},
		{
			name:   "Currency with many decimals (Philippines rounds to 2)",
			amount: 123.456789,
			options: &numi18n.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "₱123.46",
		},
		{
			name:     "No locale fallback to US",
			amount:   1234,
			options:  &numi18n.NumI18NOptions{},
			expected: "1,234",
		},
		{
			name:   "Currency with no locale fallback",
			amount: 1234,
			options: &numi18n.NumI18NOptions{
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "$1,234.00",
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

// ========== French Locale ToFormat Tests ==========

func TestToFormat_French_Numbers(t *testing.T) {
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
				Locale: "fr-FR",
			},
			expected: "0",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
			},
			expected: "5",
		},
		{
			name:   "Two digits",
			amount: 42,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
			},
			expected: "42",
		},
		{
			name:   "Three digits",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
			},
			expected: "123",
		},
		{
			name:   "Four digits with space separator",
			amount: 1234,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
			},
			expected: "1 234",
		},
		{
			name:   "Five digits with space separator",
			amount: 12345,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
			},
			expected: "12 345",
		},
		{
			name:   "Six digits with space separators",
			amount: 123456,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
			},
			expected: "123 456",
		},
		{
			name:   "Seven digits with space separators",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
			},
			expected: "1 234 567",
		},
		{
			name:   "One million with space separators",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
			},
			expected: "1 000 000",
		},
		{
			name:   "Large number with space separators",
			amount: 12345678,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
			},
			expected: "12 345 678",
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

func TestToFormat_French_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Decimal with one place",
			amount: 123.4,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
			},
			expected: "123,4",
		},
		{
			name:   "Decimal with two places",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
			},
			expected: "123,45",
		},
		{
			name:   "Decimal with multiple places",
			amount: 123.456789,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
			},
			expected: "123,456789",
		},
		{
			name:   "Large number with decimals and space separators",
			amount: 1234567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
			},
			expected: "1 234 567,89",
		},
		{
			name:   "Small decimal",
			amount: 0.123,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
			},
			expected: "0,123",
		},
		{
			name:   "Complex decimal with space separators",
			amount: 123456.123456,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
			},
			expected: "123 456,123456",
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

func TestToFormat_French_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Simple euro amount",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "€123,45",
		},
		{
			name:   "Large euro amount with space separators",
			amount: 1234567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "€1 234 567,89",
		},
		{
			name:   "Negative euro",
			amount: -123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "-€123,45",
		},
		{
			name:   "Override currency symbol",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Symbol: "$",
					},
				},
			},
			expected: "$123,45",
		},
		{
			name:   "Zero euro",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "€0",
		},
		{
			name:   "Euro with decimals",
			amount: 1000.50,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "€1 000,5",
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

func TestToFormat_French_Negative(t *testing.T) {
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
				Locale: "fr-FR",
			},
			expected: "-5",
		},
		{
			name:   "Negative large number with space separators",
			amount: -1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
			},
			expected: "-1 234 567",
		},
		{
			name:   "Negative with decimals",
			amount: -123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
			},
			expected: "-123,45",
		},
		{
			name:   "Negative decimal only",
			amount: -0.123,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
			},
			expected: "-0,123",
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

func TestToFormat_French_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Very small positive number",
			amount: 0.0001,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
			},
			expected: "0,0001",
		},
		{
			name:   "Very small negative number",
			amount: -0.0001,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
			},
			expected: "-0,0001",
		},
		{
			name:   "Very large number with space separators",
			amount: 999999999.99,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
			},
			expected: "999 999 999,99",
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

func TestToFormat_FrenchCanada_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Canadian dollar amount",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-CA",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "C$123,45",
		},
		{
			name:   "Large Canadian dollar with space separators",
			amount: 1234567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-CA",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "C$1 234 567,89",
		},
		{
			name:   "Negative Canadian dollar",
			amount: -123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-CA",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "-C$123,45",
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

// ========== European Locale ToFormat Tests ==========

func TestToFormat_European_Numbers(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "German number with period separators",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "de-DE",
			},
			expected: "1.234,56",
		},
		{
			name:   "Dutch number with period separators",
			amount: 1234567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "nl-NL",
			},
			expected: "1.234.567,89",
		},
		{
			name:   "Danish number with period separators",
			amount: 5000.25,
			options: &numi18n.NumI18NOptions{
				Locale: "da-DK",
			},
			expected: "5.000,25",
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

func TestToFormat_Asian_Numbers(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		// Korean formatting (no thousands separator, period decimal)
		{
			name:   "Korean number basic",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "₩1234.56",
		},
		{
			name:   "Korean number large",
			amount: 1234567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "₩1234567.89",
		},
		// Chinese formatting (no thousands separator, period decimal)
		{
			name:   "Chinese number basic",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "zh-CN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "¥1234.56",
		},
		{
			name:   "Chinese number large",
			amount: 1234567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "zh-CN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "¥1234567.89",
		},
		// Thai formatting (no thousands separator, period decimal)
		{
			name:   "Thai number basic",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "฿1234.56",
		},
		{
			name:   "Thai number large",
			amount: 1234567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "฿1234567.89",
		},
		// Japanese formatting (no thousands separator, no decimal for currency)
		{
			name:   "Japanese number basic",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "¥1235", // Japanese Yen has no decimals
		},
		{
			name:   "Japanese number large",
			amount: 1234567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "¥1234568", // Japanese Yen has no decimals
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

func TestToFormat_African_Numbers(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		// Afrikaans South Africa (af-ZA)
		{
			name:   "Afrikaans basic number",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "af-ZA",
			},
			expected: "1,234.56",
		},
		{
			name:   "Afrikaans large number",
			amount: 1234567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "af-ZA",
			},
			expected: "1,234,567.89",
		},
		{
			name:   "Afrikaans currency",
			amount: 1500.75,
			options: &numi18n.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "R1,500.75",
		},
		// Amharic Ethiopia (am-ET)
		{
			name:   "Amharic basic number",
			amount: 2500.25,
			options: &numi18n.NumI18NOptions{
				Locale: "am-ET",
			},
			expected: "2,500.25",
		},
		{
			name:   "Amharic currency",
			amount: 750.50,
			options: &numi18n.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Br750.5",
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

func TestToFormat_Arabic_Numbers(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		// Arabic World (ar-001)
		{
			name:   "Arabic world basic number",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
			},
			expected: "1,234.56",
		},
		{
			name:   "Arabic world currency",
			amount: 1500.75,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "$1,500.75",
		},
		// Arabic UAE (ar-AE)
		{
			name:   "Arabic UAE basic number",
			amount: 2500.25,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-AE",
			},
			expected: "2,500.25",
		},
		{
			name:   "Arabic UAE currency",
			amount: 750.50,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-AE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "د.إ750.5",
		},
		// Arabic Bahrain (ar-BH)
		{
			name:   "Arabic Bahrain basic number",
			amount: 999.99,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-BH",
			},
			expected: "999.99",
		},
		{
			name:   "Arabic Bahrain currency",
			amount: 1200.00,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-BH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "د.ب1,200",
		},
		// Arabic Algeria (ar-DZ)
		{
			name:   "Arabic Algeria basic number",
			amount: 3456.78,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-DZ",
			},
			expected: "3,456.78",
		},
		{
			name:   "Arabic Algeria currency",
			amount: 5000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-DZ",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "دج5,000",
		},
		// Arabic Egypt (ar-EG)
		{
			name:   "Arabic Egypt basic number",
			amount: 7890.12,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-EG",
			},
			expected: "7,890.12",
		},
		{
			name:   "Arabic Egypt currency",
			amount: 2500.50,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-EG",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "ج.م2,500.5",
		},
		// Arabic Iraq (ar-IQ)
		{
			name:   "Arabic Iraq basic number",
			amount: 4567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-IQ",
			},
			expected: "4,567.89",
		},
		{
			name:   "Arabic Iraq currency",
			amount: 10000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "ع.د10,000",
		},
		// Arabic Jordan (ar-JO)
		{
			name:   "Arabic Jordan basic number",
			amount: 6543.21,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-JO",
			},
			expected: "6,543.21",
		},
		{
			name:   "Arabic Jordan currency",
			amount: 800.75,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-JO",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "د.أ800.75",
		},
		// Arabic Kuwait (ar-KW)
		{
			name:   "Arabic Kuwait basic number",
			amount: 3210.98,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-KW",
			},
			expected: "3,210.98",
		},
		{
			name:   "Arabic Kuwait currency",
			amount: 450.25,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-KW",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "د.ك450.25",
		},
		// Arabic Lebanon (ar-LB)
		{
			name:   "Arabic Lebanon basic number",
			amount: 8765.43,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-LB",
			},
			expected: "8,765.43",
		},
		{
			name:   "Arabic Lebanon currency",
			amount: 1500000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-LB",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "ل.ل1,500,000",
		},
		// Arabic Libya (ar-LY)
		{
			name:   "Arabic Libya basic number",
			amount: 2468.13,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-LY",
			},
			expected: "2,468.13",
		},
		{
			name:   "Arabic Libya currency",
			amount: 650.00,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-LY",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "ل.د650",
		},
		// Arabic Morocco (ar-MA)
		{
			name:   "Arabic Morocco basic number",
			amount: 9876.54,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-MA",
			},
			expected: "9,876.54",
		},
		{
			name:   "Arabic Morocco currency",
			amount: 1200.75,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-MA",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "د.م1,200.75",
		},
		// Arabic Oman (ar-OM)
		{
			name:   "Arabic Oman basic number",
			amount: 1357.92,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-OM",
			},
			expected: "1,357.92",
		},
		{
			name:   "Arabic Oman currency",
			amount: 300.50,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-OM",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "ر.ع.300.5",
		},
		// Arabic Qatar (ar-QA)
		{
			name:   "Arabic Qatar basic number",
			amount: 8642.97,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-QA",
			},
			expected: "8,642.97",
		},
		{
			name:   "Arabic Qatar currency",
			amount: 2500.25,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-QA",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "ر.ق2,500.25",
		},
		// Arabic Saudi Arabia (ar-SA)
		{
			name:   "Arabic Saudi basic number",
			amount: 7531.86,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-SA",
			},
			expected: "7,531.86",
		},
		{
			name:   "Arabic Saudi currency",
			amount: 1800.00,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-SA",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "ر.س1,800",
		},
		// Arabic Syria (ar-SY)
		{
			name:   "Arabic Syria basic number",
			amount: 4269.15,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-SY",
			},
			expected: "4,269.15",
		},
		{
			name:   "Arabic Syria currency",
			amount: 50000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-SY",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "ل.س50,000",
		},
		// Arabic Tunisia (ar-TN)
		{
			name:   "Arabic Tunisia basic number",
			amount: 6174.82,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-TN",
			},
			expected: "6,174.82",
		},
		{
			name:   "Arabic Tunisia currency",
			amount: 950.75,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-TN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "د.ت950.75",
		},
		// Arabic Yemen (ar-YE)
		{
			name:   "Arabic Yemen basic number",
			amount: 8395.67,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-YE",
			},
			expected: "8,395.67",
		},
		{
			name:   "Arabic Yemen currency",
			amount: 25000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-YE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "ر.ي25,000",
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

func TestToFormat_Indian_Caucasus_Numbers(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		// Assamese India (as-IN)
		{
			name:   "Assamese basic number",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "as-IN",
			},
			expected: "1,234.56",
		},
		{
			name:   "Assamese large number",
			amount: 9876543.21,
			options: &numi18n.NumI18NOptions{
				Locale: "as-IN",
			},
			expected: "9,876,543.21",
		},
		{
			name:   "Assamese currency",
			amount: 2500.75,
			options: &numi18n.NumI18NOptions{
				Locale: "as-IN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "₹2,500.75",
		},
		// Azerbaijani Azerbaijan (az-AZ)
		{
			name:   "Azerbaijani basic number",
			amount: 3456.78,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
			},
			expected: "3,456.78",
		},
		{
			name:   "Azerbaijani large number",
			amount: 1234567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
			},
			expected: "1,234,567.89",
		},
		{
			name:   "Azerbaijani currency",
			amount: 850.25,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "₼850.25",
		},
		{
			name:   "Azerbaijani small decimal",
			amount: 0.123,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
			},
			expected: "0.123",
		},
		{
			name:   "Azerbaijani zero",
			amount: 0.0,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
			},
			expected: "0",
		},
		{
			name:   "Azerbaijani negative number",
			amount: -1500.50,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
			},
			expected: "-1,500.5",
		},
		{
			name:   "Azerbaijani negative currency",
			amount: -750.25,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "-₼750.25",
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

func TestToFormat_EdgeCases_SpecifiedLocales(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		// Test edge cases with various locales
		{
			name:   "Very large number Arabic",
			amount: 999999999999.99,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-SA",
			},
			expected: "999,999,999,999.99",
		},
		{
			name:   "Very small decimal African",
			amount: 0.001,
			options: &numi18n.NumI18NOptions{
				Locale: "af-ZA",
			},
			expected: "0.001",
		},
		{
			name:   "Zero with currency Arabic",
			amount: 0.0,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-EG",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "ج.م0",
		},
		{
			name:   "Large currency Amharic",
			amount: 1000000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Br1,000,000",
		},
		{
			name:   "Complex decimal Assamese",
			amount: 123456.123456,
			options: &numi18n.NumI18NOptions{
				Locale: "as-IN",
			},
			expected: "123,456.123456",
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

// Test comprehensive English locales worldwide
func TestToFormat_English_Commonwealth(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// United Kingdom - Pound Sterling
		{
			name:   "en-GB Pound",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-GB",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "£1,234.56",
		},

		// Australia - Australian Dollar
		{
			name:   "en-AU AUD",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-AU",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "A$1,234.56",
		},

		// Canada - Canadian Dollar
		{
			name:   "en-CA CAD",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-CA",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "C$1,234.56",
		},

		// New Zealand
		{
			name:   "en-NZ NZD",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-NZ",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "NZ$1,234.56",
		},

		// Ireland - Euro
		{
			name:   "en-IE Euro",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-IE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1,234.56",
		},

		// South Africa - Rand
		{
			name:   "en-ZA Rand",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-ZA",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "R1 234,56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test English locales in Asia and Africa with different currencies
func TestToFormat_English_AsiaAfrica(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Hong Kong - Hong Kong Dollar
		{
			name:   "en-HK HKD",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-HK",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "HK$1,234.56",
		},

		// Singapore - Singapore Dollar
		{
			name:   "en-SG SGD",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-SG",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "S$1,234.56",
		},

		// India - Indian Rupee
		{
			name:   "en-IN INR",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-IN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "₹1,234.56",
		},

		// Kenya - Kenyan Shilling
		{
			name:   "en-KE KES",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-KE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "KSh1,234.56",
		},

		// Ghana - Ghanaian Cedi
		{
			name:   "en-GH GHS",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-GH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "GH₵1,234.56",
		},

		// Nigeria - Nigerian Naira
		{
			name:   "en-NG NGN",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-NG",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "₦1,234.56",
		},

		// Philippines - Philippine Peso
		{
			name:   "en-PH PHP",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-PH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "₱1,234.56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test Greek and other European locales
func TestToFormat_Other_European(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Greek - Euro
		{
			name:   "el-GR Euro",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "el-GR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1,234.56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test edge cases with various locales
func TestToFormat_Global_EdgeCases(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Zero amounts
		{
			name:   "en-GB Zero Pound",
			amount: 0.00,
			options: &numi18n.NumI18NOptions{
				Locale: "en-GB",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "£0.00",
		},
		{
			name:   "en-AU Zero AUD",
			amount: 0.00,
			options: &numi18n.NumI18NOptions{
				Locale: "en-AU",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "A$0.00",
		},

		// Large amounts
		{
			name:   "en-GB Million Pound",
			amount: 1000000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "en-GB",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "£1,000,000.00",
		},
		{
			name:   "en-CA Million CAD",
			amount: 1000000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "en-CA",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "C$1,000,000.00",
		},

		// Negative amounts
		{
			name:   "en-GB Negative Pound",
			amount: -1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-GB",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "-£1,234.56",
		},
		{
			name:   "en-IN Negative Rupee",
			amount: -1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-IN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "-₹1,234.56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test decimal formatting without currency
func TestToFormat_Decimal_Only(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// European decimal formatting (period thousands, comma decimal)
		{
			name:   "de-DE Decimal",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "de-DE",
			},
			expectedFormat: "1.234,56",
		},
		{
			name:   "da-DK Decimal",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "da-DK",
			},
			expectedFormat: "1.234,56",
		},

		// Anglo decimal formatting (comma thousands, period decimal)
		{
			name:   "en-US Decimal",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
			},
			expectedFormat: "1,234.56",
		},
		{
			name:   "en-GB Decimal",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-GB",
			},
			expectedFormat: "1,234.56",
		},
		{
			name:   "en-AU Decimal",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-AU",
			},
			expectedFormat: "1,234.56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

func TestToFormat_European_Slavic_Numbers(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		// Belarusian (Belarus) - be-BY
		{
			name:   "Belarusian_basic_number",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "be-BY",
			},
			expected: "1,234.56",
		},
		{
			name:   "Belarusian_large_number",
			amount: 10000000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "be-BY",
			},
			expected: "10,000,000",
		},
		{
			name:   "Belarusian_currency",
			amount: 1500.75,
			options: &numi18n.NumI18NOptions{
				Locale: "be-BY",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "BYN1,500.75",
		},

		// Bulgarian (Bulgaria) - bg-BG
		{
			name:   "Bulgarian_basic_number",
			amount: 2345.67,
			options: &numi18n.NumI18NOptions{
				Locale: "bg-BG",
			},
			expected: "2,345.67",
		},
		{
			name:   "Bulgarian_large_number",
			amount: 5000000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "bg-BG",
			},
			expected: "5,000,000",
		},
		{
			name:   "Bulgarian_currency",
			amount: 750.25,
			options: &numi18n.NumI18NOptions{
				Locale: "bg-BG",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "лв750.25",
		},

		// Bosnian (Bosnia and Herzegovina) - bs-BA
		{
			name:   "Bosnian_basic_number",
			amount: 3456.78,
			options: &numi18n.NumI18NOptions{
				Locale: "bs-BA",
			},
			expected: "3,456.78",
		},
		{
			name:   "Bosnian_currency",
			amount: 999.99,
			options: &numi18n.NumI18NOptions{
				Locale: "bs-BA",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "KM999.99",
		},

		// Czech (Czech Republic) - cs-CZ
		{
			name:   "Czech_basic_number",
			amount: 4567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "cs-CZ",
			},
			expected: "4,567.89",
		},
		{
			name:   "Czech_currency",
			amount: 1200.50,
			options: &numi18n.NumI18NOptions{
				Locale: "cs-CZ",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Kč1,200.5",
		},

		// Welsh (United Kingdom) - cy-GB
		{
			name:   "Welsh_basic_number",
			amount: 5678.90,
			options: &numi18n.NumI18NOptions{
				Locale: "cy-GB",
			},
			expected: "5,678.9",
		},
		{
			name:   "Welsh_currency",
			amount: 2500.00,
			options: &numi18n.NumI18NOptions{
				Locale: "cy-GB",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "£2,500",
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

func TestToFormat_Romance_Catalan_Numbers(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		// Catalan (Spain) - ca-ES
		{
			name:   "Catalan_basic_number",
			amount: 6789.01,
			options: &numi18n.NumI18NOptions{
				Locale: "ca-ES",
			},
			expected: "6,789.01",
		},
		{
			name:   "Catalan_large_number",
			amount: 8000000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "ca-ES",
			},
			expected: "8,000,000",
		},
		{
			name:   "Catalan_currency",
			amount: 1800.75,
			options: &numi18n.NumI18NOptions{
				Locale: "ca-ES",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "€1,800.75",
		},
		{
			name:   "Catalan_small_decimal",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "ca-ES",
			},
			expected: "0.99",
		},
		{
			name:   "Catalan_zero",
			amount: 0.00,
			options: &numi18n.NumI18NOptions{
				Locale: "ca-ES",
			},
			expected: "0",
		},
		{
			name:   "Catalan_negative_number",
			amount: -1500.50,
			options: &numi18n.NumI18NOptions{
				Locale: "ca-ES",
			},
			expected: "-1,500.5",
		},
		{
			name:   "Catalan_negative_currency",
			amount: -750.25,
			options: &numi18n.NumI18NOptions{
				Locale: "ca-ES",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "-€750.25",
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

func TestToFormat_Bengali_Numbers(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		// Bengali (Bangladesh) - bn-BD
		{
			name:   "Bengali_BD_basic_number",
			amount: 7890.12,
			options: &numi18n.NumI18NOptions{
				Locale: "bn-BD",
			},
			expected: "7,890.12",
		},
		{
			name:   "Bengali_BD_large_number",
			amount: 3000000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "bn-BD",
			},
			expected: "3,000,000",
		},
		{
			name:   "Bengali_BD_currency",
			amount: 2200.50,
			options: &numi18n.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "৳2,200.5",
		},

		// Bengali (India) - bn-IN
		{
			name:   "Bengali_IN_basic_number",
			amount: 8901.23,
			options: &numi18n.NumI18NOptions{
				Locale: "bn-IN",
			},
			expected: "8,901.23",
		},
		{
			name:   "Bengali_IN_large_number",
			amount: 4500000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "bn-IN",
			},
			expected: "4,500,000",
		},
		{
			name:   "Bengali_IN_currency",
			amount: 3300.75,
			options: &numi18n.NumI18NOptions{
				Locale: "bn-IN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "₹3,300.75",
		},
		{
			name:   "Bengali_IN_small_decimal",
			amount: 0.05,
			options: &numi18n.NumI18NOptions{
				Locale: "bn-IN",
			},
			expected: "0.05",
		},
		{
			name:   "Bengali_IN_zero",
			amount: 0.00,
			options: &numi18n.NumI18NOptions{
				Locale: "bn-IN",
			},
			expected: "0",
		},
		{
			name:   "Bengali_IN_negative_number",
			amount: -2500.25,
			options: &numi18n.NumI18NOptions{
				Locale: "bn-IN",
			},
			expected: "-2,500.25",
		},
		{
			name:   "Bengali_IN_negative_currency",
			amount: -1100.80,
			options: &numi18n.NumI18NOptions{
				Locale: "bn-IN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "-₹1,100.8",
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

func TestToFormat_EdgeCases_EuropeanBengali(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Very_large_number_Belarusian",
			amount: 999999999.99,
			options: &numi18n.NumI18NOptions{
				Locale: "be-BY",
			},
			expected: "999,999,999.99",
		},
		{
			name:   "Very_small_decimal_Bulgarian",
			amount: 0.001,
			options: &numi18n.NumI18NOptions{
				Locale: "bg-BG",
			},
			expected: "0.001",
		},
		{
			name:   "Zero_with_currency_Czech",
			amount: 0.00,
			options: &numi18n.NumI18NOptions{
				Locale: "cs-CZ",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Kč0",
		},
		{
			name:   "Large_currency_Welsh",
			amount: 1000000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "cy-GB",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "£1,000,000",
		},
		{
			name:   "Complex_decimal_Bengali_BD",
			amount: 12345.6789,
			options: &numi18n.NumI18NOptions{
				Locale: "bn-BD",
			},
			expected: "12,345.6789",
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

// Test Polish and Eastern European locales
func TestToFormat_Polish_EasternEuropean(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Poland - Zloty
		{
			name:   "pl-PL Zloty",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "pl-PL",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "zł1,234.56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test Portuguese locales worldwide
func TestToFormat_Portuguese_Locales(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Portugal - Euro
		{
			name:   "pt-PT Euro",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-PT",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1,234.56",
		},
		// Brazil - Real
		{
			name:   "pt-BR Real",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-BR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "R$1,234.56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test Russian and Slavic locales
func TestToFormat_Russian_Slavic_Locales(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Russia - Ruble
		{
			name:   "ru-RU Ruble",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "ru-RU",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "₽1,234.56",
		},
		// Ukraine - Hryvnia
		{
			name:   "uk-UA Hryvnia",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "uk-UA",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "₴1,234.56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test Chinese locales comprehensive
func TestToFormat_Chinese_Comprehensive(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// China - Yuan
		{
			name:   "zh-CN Yuan",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "zh-CN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "¥1234.56",
		},
		// Hong Kong - HK Dollar
		{
			name:   "zh-HK HKD",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "zh-HK",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "HK$1234.56",
		},
		// Taiwan - New Taiwan Dollar
		{
			name:   "zh-TW TWD",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "zh-TW",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "NT$1234.56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test Asian and Middle Eastern locales
func TestToFormat_Asian_MiddleEastern(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Turkish - Lira
		{
			name:   "tr-TR Lira",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "tr-TR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "₺1,234.56",
		},
		// Thai - Baht
		{
			name:   "th-TH Baht",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "฿1234.56",
		},
		// Vietnamese - Dong
		{
			name:   "vi-VN Dong",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "vi-VN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "₫1,234.56",
		},
		// Urdu Pakistan - Rupee
		{
			name:   "ur-PK Rupee",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "ur-PK",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "Rs1,234.56",
		},
		// Tamil India - Rupee
		{
			name:   "ta-IN Rupee",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "₹1,234.56",
		},
		// Uzbek - Som
		{
			name:   "uz-UZ Som",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "uz-UZ",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "so'm1,234.56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test Additional locale groups
func TestToFormat_Additional_Locales(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Romania - Leu
		{
			name:   "ro-RO Leu",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "ro-RO",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "RON1,234.56",
		},
		// Slovakia - Euro
		{
			name:   "sk-SK Euro",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1.234,56",
		},
		// Serbia - Dinar
		{
			name:   "sr-RS Dinar",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "sr-RS",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "РСД1.234,56",
		},
		// Albanian - Lek
		{
			name:   "sq-AL Lek",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "sq-AL",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "L1,234.56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

func TestToFormat_Nordic_Scandinavian(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Sweden - Krona
		{
			name:   "sv-SE Krona",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "sv-SE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "kr1.234,56",
		},
		// Norway - Krone
		{
			name:   "nb-NO Krone",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "nb-NO",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "$1,234.56",
		},
		// Denmark - Krone
		{
			name:   "da-DK Krone",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "da-DK",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "kr1.234,56",
		},
		// Iceland - Krona
		{
			name:   "is-IS Krona",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "kr1.234,56",
		},
		// Finland - Euro
		{
			name:   "fi-FI Euro",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1.234,56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

func TestToFormat_Additional_African_Locales(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// South Africa - Rand
		{
			name:   "en-ZA Rand",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-ZA",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "R1 234,56",
		},
		// Nigeria - Naira
		{
			name:   "en-NG Naira",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-NG",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "₦1,234.56",
		},
		// Ghana - Cedi
		{
			name:   "en-GH Cedi",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-GH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "GH₵1,234.56",
		},
		// Kenya - Shilling
		{
			name:   "en-KE Shilling",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-KE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "KSh1,234.56",
		},
		// Tanzania - Shilling
		{
			name:   "en-TZ Shilling",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-TZ",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "TSh1,234.56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

func TestToFormat_Latin_American_Locales(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Mexico - Peso
		{
			name:   "es-MX Peso",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "es-MX",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "$1,234.56",
		},
		// Argentina - Peso
		{
			name:   "es-AR Peso",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "es-AR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "$1,234.56",
		},
		// Brazil - Real
		{
			name:   "pt-BR Real",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-BR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "R$1,234.56",
		},
		// Colombia - Peso
		{
			name:   "es-CO Peso",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "es-CO",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "$1,234.56",
		},
		// Chile - Peso
		{
			name:   "es-CL Peso",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "es-CL",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "$1,234.56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

func TestToFormat_Additional_European_Locales(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Switzerland - Franc
		{
			name:   "de-CH Franc",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "de-CH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "CHF1.234,56",
		},
		// Austria - Euro
		{
			name:   "de-AT Euro",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1.234,56",
		},
		// Belgium - Euro
		{
			name:   "nl-BE Euro",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1.234,56",
		},
		// Norway - Krone (using no-NO)
		{
			name:   "no-NO Krone",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "no-NO",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "kr1.234,56",
		},
		// Luxembourg - Euro
		{
			name:   "fr-LU Euro",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-LU",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1 234,56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test French locales worldwide
func TestToFormat_French_Locales(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// France - Euro
		{
			name:   "fr-FR Euro",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1 234,56",
		},

		// Canada French - Canadian Dollar
		{
			name:   "fr-CA CAD",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-CA",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "C$1 234,56",
		},

		// Belgium French - Euro
		{
			name:   "fr-BE Euro",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-BE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1 234,56",
		},

		// Switzerland French - Swiss Franc
		{
			name:   "fr-CH CHF",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-CH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "CHF1 234,56",
		},

		// Luxembourg French - Euro
		{
			name:   "fr-LU Euro",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-LU",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1 234,56",
		},

		// World French - Euro
		{
			name:   "fr-001 Euro",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-001",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1 234,56",
		},

		// Morocco French - Dirham
		{
			name:   "fr-MA MAD",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-MA",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "MAD1 234,56",
		},

		// Senegal French - CFA Franc
		{
			name:   "fr-SN XOF",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-SN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "CFA1 234,56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test Italian locales
func TestToFormat_Italian_Locales(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Italy - Euro
		{
			name:   "it-IT Euro",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "it-IT",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1.234,56",
		},

		// Switzerland Italian - Swiss Franc
		{
			name:   "it-CH CHF",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "it-CH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "CHF1.234,56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test Nordic/Northern European locales
func TestToFormat_Nordic_Locales(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Finland - Euro
		{
			name:   "fi-FI Euro",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1.234,56",
		},

		// Iceland - Icelandic Krona
		{
			name:   "is-IS ISK",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "kr1.234,56",
		},

		// Faroe Islands - Danish Krone
		{
			name:   "fo-FO DKK",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "kr1,234.56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test Celtic locales
func TestToFormat_Celtic_Locales(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Irish Gaelic - Euro
		{
			name:   "ga-IE Euro",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "ga-IE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1.234,56",
		},

		// Scottish Gaelic - Pound
		{
			name:   "gd-GB Pound",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "gd-GB",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "£1,234.56",
		},

		// Galician - Euro
		{
			name:   "gl-ES Euro",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "gl-ES",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1.234,56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test Middle Eastern and Persian locales
func TestToFormat_MiddleEast_Locales(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Iran Persian - Rial
		{
			name:   "fa-IR IRR",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "﷼1,234.56",
		},

		// Afghanistan Persian - Afghani
		{
			name:   "fa-AF AFN",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "؋1,234.56",
		},

		// Hebrew Israel - Shekel
		{
			name:   "he-IL ILS",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "he-IL",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "₪1,234.56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test Indian subcontinent locales
func TestToFormat_Indian_Locales(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Hindi India - Rupee
		{
			name:   "hi-IN INR",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "₹1,234.56",
		},

		// Gujarati India - Rupee
		{
			name:   "gu-IN INR",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "₹1,234.56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test Eastern European locales
func TestToFormat_EasternEuropean_Locales(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Hungarian - Forint
		{
			name:   "hu-HU HUF",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "hu-HU",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "Ft1.234,56",
		},

		// Croatian - Kuna
		{
			name:   "hr-HR HRK",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "hr-HR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1.234,56",
		},

		// Armenian - Dram
		{
			name:   "hy-AM AMD",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "֏1,234.56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test African locales
func TestToFormat_African_Locales(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Hausa Nigeria - Naira
		{
			name:   "ha-NG NGN",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "ha-NG",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "₦1,234.56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test Asian locales
func TestToFormat_Asian_Locales(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Japanese - Yen (no decimals)
		{
			name:   "ja-JP JPY",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "¥1235",
		},

		// Indonesian - Rupiah
		{
			name:   "id-ID IDR",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "id-ID",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "Rp1,234.56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test Caribbean locales
func TestToFormat_Caribbean_Locales(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Haitian Creole - Gourde
		{
			name:   "ht-HT HTG",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "ht-HT",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "G1 234,56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test large amounts and edge cases
func TestToFormat_Additional_EdgeCases(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// French large amount
		{
			name:   "fr-FR Large Euro",
			amount: 1234567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1 234 567,89",
		},

		// Japanese large amount (no decimals)
		{
			name:   "ja-JP Large Yen",
			amount: 1234567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "¥1234568",
		},

		// Italian negative amount
		{
			name:   "it-IT Negative Euro",
			amount: -1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "it-IT",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "-€1.234,56",
		},

		// French zero amount
		{
			name:   "fr-FR Zero Euro",
			amount: 0.00,
			options: &numi18n.NumI18NOptions{
				Locale: "fr-FR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test European locales that use European formatting (period thousands separator, comma decimal separator, prefix currency)
func TestToFormat_European_EuroZone(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// German locales
		{
			name:   "de-DE Euro",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "de-DE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1.234,56",
		},
		{
			name:   "de-AT Euro",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1.234,56",
		},

		// Basque (Spain) - Euro
		{
			name:   "eu-ES Euro",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1.234,56",
		},

		// Test larger amounts
		{
			name:   "de-DE Large Euro",
			amount: 1234567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "de-DE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1.234.567,89",
		},
		{
			name:   "de-AT Large Euro",
			amount: 1234567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1.234.567,89",
		},
		{
			name:   "eu-ES Large Euro",
			amount: 1234567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1.234.567,89",
		},

		// Test small amounts
		{
			name:   "de-DE Small Euro",
			amount: 12.34,
			options: &numi18n.NumI18NOptions{
				Locale: "de-DE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€12,34",
		},
		{
			name:   "de-AT Small Euro",
			amount: 12.34,
			options: &numi18n.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€12,34",
		},
		{
			name:   "eu-ES Small Euro",
			amount: 12.34,
			options: &numi18n.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€12,34",
		},

		// Test whole numbers
		{
			name:   "de-DE Whole Euro",
			amount: 1000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "de-DE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1.000",
		},
		{
			name:   "de-AT Whole Euro",
			amount: 1000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1.000",
		},
		{
			name:   "eu-ES Whole Euro",
			amount: 1000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1.000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

func TestToFormat_European_Swiss(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Swiss Franc (CHF)
		{
			name:   "de-CH CHF",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "de-CH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "CHF1.234,56",
		},
		{
			name:   "de-CH Large CHF",
			amount: 1234567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "de-CH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "CHF1.234.567,89",
		},
		{
			name:   "de-CH Small CHF",
			amount: 12.34,
			options: &numi18n.NumI18NOptions{
				Locale: "de-CH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "CHF12,34",
		},
		{
			name:   "de-CH Whole CHF",
			amount: 1000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "de-CH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "CHF1.000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

func TestToFormat_European_Danish(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Danish Krone (kr)
		{
			name:   "da-DK Krone",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "da-DK",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "kr1.234,56",
		},
		{
			name:   "da-DK Large Krone",
			amount: 1234567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "da-DK",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "kr1.234.567,89",
		},
		{
			name:   "da-DK Small Krone",
			amount: 12.34,
			options: &numi18n.NumI18NOptions{
				Locale: "da-DK",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "kr12,34",
		},
		{
			name:   "da-DK Whole Krone",
			amount: 1000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "da-DK",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "kr1.000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test English locales that use Anglo formatting (comma thousands separator, period decimal separator, prefix currency)
func TestToFormat_English_Major(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// US Dollar
		{
			name:   "en-US USD",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "$1,234.56",
		},
		{
			name:   "en-US Large USD",
			amount: 1234567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "$1,234,567.89",
		},
		{
			name:   "en-US Small USD",
			amount: 12.34,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "$12.34",
		},
		{
			name:   "en-US Whole USD",
			amount: 1000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "$1,000.00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test negative numbers across different locale groups
func TestToFormat_European_Negative(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// European formatting - negative amounts
		{
			name:   "de-DE Negative Euro",
			amount: -1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "de-DE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "-€1.234,56",
		},
		{
			name:   "de-AT Negative Euro",
			amount: -1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "-€1.234,56",
		},
		{
			name:   "de-CH Negative CHF",
			amount: -1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "de-CH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "-CHF1.234,56",
		},
		{
			name:   "da-DK Negative Krone",
			amount: -1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "da-DK",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "-kr1.234,56",
		},
		{
			name:   "eu-ES Negative Euro",
			amount: -1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "-€1.234,56",
		},

		// English formatting - negative amounts
		{
			name:   "en-US Negative USD",
			amount: -1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "-$1,234.56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

// Test edge cases
func TestToFormat_European_EdgeCases(t *testing.T) {
	tests := []struct {
		name           string
		amount         float64
		options        *numi18n.NumI18NOptions
		expectedFormat string
	}{
		// Zero values
		{
			name:   "de-DE Zero Euro",
			amount: 0.00,
			options: &numi18n.NumI18NOptions{
				Locale: "de-DE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€0",
		},
		{
			name:   "en-US Zero USD",
			amount: 0.00,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "$0.00",
		},

		// Very small amounts
		{
			name:   "de-DE Cents Euro",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "de-DE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€0,99",
		},
		{
			name:   "en-US Cents USD",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "$0.99",
		},

		// Very large amounts
		{
			name:   "de-DE Million Euro",
			amount: 1000000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "de-DE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "€1.000.000",
		},
		{
			name:   "en-US Million USD",
			amount: 1000000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expectedFormat: "$1,000,000.00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToFormat(tt.amount)
			if result != tt.expectedFormat {
				t.Errorf("ToFormat() = %v, want %v", result, tt.expectedFormat)
			}
		})
	}
}

func TestToFormat_Spanish_Europe_World(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		// Spanish (World) - es-001
		{
			name:   "Spanish_World_basic_number",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "es-001",
			},
			expected: "1,234.56",
		},
		{
			name:   "Spanish_World_currency",
			amount: 1500.75,
			options: &numi18n.NumI18NOptions{
				Locale: "es-001",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "€1,500.75",
		},

		// Spanish (Spain) - es-ES
		{
			name:   "Spanish_Spain_basic_number",
			amount: 2345.67,
			options: &numi18n.NumI18NOptions{
				Locale: "es-ES",
			},
			expected: "2,345.67",
		},
		{
			name:   "Spanish_Spain_currency",
			amount: 999.99,
			options: &numi18n.NumI18NOptions{
				Locale: "es-ES",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "€999.99",
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

func TestToFormat_Spanish_SouthAmerica_North(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		// Spanish (Argentina) - es-AR
		{
			name:   "Spanish_Argentina_basic_number",
			amount: 3456.78,
			options: &numi18n.NumI18NOptions{
				Locale: "es-AR",
			},
			expected: "3,456.78",
		},
		{
			name:   "Spanish_Argentina_currency",
			amount: 1200.50,
			options: &numi18n.NumI18NOptions{
				Locale: "es-AR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "$1,200.5",
		},

		// Spanish (Bolivia) - es-BO
		{
			name:   "Spanish_Bolivia_basic_number",
			amount: 4567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "es-BO",
			},
			expected: "4,567.89",
		},
		{
			name:   "Spanish_Bolivia_currency",
			amount: 800.25,
			options: &numi18n.NumI18NOptions{
				Locale: "es-BO",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Bs800.25",
		},

		// Spanish (Chile) - es-CL
		{
			name:   "Spanish_Chile_basic_number",
			amount: 5678.90,
			options: &numi18n.NumI18NOptions{
				Locale: "es-CL",
			},
			expected: "5,678.9",
		},
		{
			name:   "Spanish_Chile_currency",
			amount: 2500.00,
			options: &numi18n.NumI18NOptions{
				Locale: "es-CL",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "$2,500",
		},

		// Spanish (Colombia) - es-CO
		{
			name:   "Spanish_Colombia_basic_number",
			amount: 6789.01,
			options: &numi18n.NumI18NOptions{
				Locale: "es-CO",
			},
			expected: "6,789.01",
		},
		{
			name:   "Spanish_Colombia_currency",
			amount: 1800.75,
			options: &numi18n.NumI18NOptions{
				Locale: "es-CO",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "$1,800.75",
		},

		// Spanish (Ecuador) - es-EC
		{
			name:   "Spanish_Ecuador_basic_number",
			amount: 7890.12,
			options: &numi18n.NumI18NOptions{
				Locale: "es-EC",
			},
			expected: "7,890.12",
		},
		{
			name:   "Spanish_Ecuador_currency",
			amount: 750.50,
			options: &numi18n.NumI18NOptions{
				Locale: "es-EC",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "$750.5",
		},

		// Spanish (Paraguay) - es-PY
		{
			name:   "Spanish_Paraguay_basic_number",
			amount: 8901.23,
			options: &numi18n.NumI18NOptions{
				Locale: "es-PY",
			},
			expected: "8,901.23",
		},
		{
			name:   "Spanish_Paraguay_currency",
			amount: 3300.75,
			options: &numi18n.NumI18NOptions{
				Locale: "es-PY",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "₲3,300.75",
		},

		// Spanish (Peru) - es-PE
		{
			name:   "Spanish_Peru_basic_number",
			amount: 9012.34,
			options: &numi18n.NumI18NOptions{
				Locale: "es-PE",
			},
			expected: "9,012.34",
		},
		{
			name:   "Spanish_Peru_currency",
			amount: 1100.80,
			options: &numi18n.NumI18NOptions{
				Locale: "es-PE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "S/.1,100.8",
		},

		// Spanish (Uruguay) - es-UY
		{
			name:   "Spanish_Uruguay_basic_number",
			amount: 1123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "es-UY",
			},
			expected: "1,123.45",
		},
		{
			name:   "Spanish_Uruguay_currency",
			amount: 2200.60,
			options: &numi18n.NumI18NOptions{
				Locale: "es-UY",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "$U2,200.6",
		},

		// Spanish (Venezuela) - es-VE
		{
			name:   "Spanish_Venezuela_basic_number",
			amount: 2234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "es-VE",
			},
			expected: "2,234.56",
		},
		{
			name:   "Spanish_Venezuela_currency",
			amount: 1500.25,
			options: &numi18n.NumI18NOptions{
				Locale: "es-VE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Bs.1,500.25",
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

func TestToFormat_Spanish_NorthAmerica_Caribbean(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		// Spanish (Mexico) - es-MX
		{
			name:   "Spanish_Mexico_basic_number",
			amount: 3345.67,
			options: &numi18n.NumI18NOptions{
				Locale: "es-MX",
			},
			expected: "3,345.67",
		},
		{
			name:   "Spanish_Mexico_currency",
			amount: 1250.75,
			options: &numi18n.NumI18NOptions{
				Locale: "es-MX",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "$1,250.75",
		},

		// Spanish (Cuba) - es-CU
		{
			name:   "Spanish_Cuba_basic_number",
			amount: 4456.78,
			options: &numi18n.NumI18NOptions{
				Locale: "es-CU",
			},
			expected: "4,456.78",
		},
		{
			name:   "Spanish_Cuba_currency",
			amount: 950.50,
			options: &numi18n.NumI18NOptions{
				Locale: "es-CU",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "$950.5",
		},

		// Spanish (Dominican Republic) - es-DO
		{
			name:   "Spanish_DominicanRepublic_basic_number",
			amount: 5567.89,
			options: &numi18n.NumI18NOptions{
				Locale: "es-DO",
			},
			expected: "5,567.89",
		},
		{
			name:   "Spanish_DominicanRepublic_currency",
			amount: 1750.25,
			options: &numi18n.NumI18NOptions{
				Locale: "es-DO",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "RD$1,750.25",
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

func TestToFormat_Spanish_CentralAmerica(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		// Spanish (Costa Rica) - es-CR
		{
			name:   "Spanish_CostaRica_basic_number",
			amount: 6678.90,
			options: &numi18n.NumI18NOptions{
				Locale: "es-CR",
			},
			expected: "6,678.9",
		},
		{
			name:   "Spanish_CostaRica_currency",
			amount: 2100.50,
			options: &numi18n.NumI18NOptions{
				Locale: "es-CR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "₡2,100.5",
		},

		// Spanish (El Salvador) - es-SV
		{
			name:   "Spanish_ElSalvador_basic_number",
			amount: 7789.01,
			options: &numi18n.NumI18NOptions{
				Locale: "es-SV",
			},
			expected: "7,789.01",
		},
		{
			name:   "Spanish_ElSalvador_currency",
			amount: 1350.75,
			options: &numi18n.NumI18NOptions{
				Locale: "es-SV",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "$1,350.75",
		},

		// Spanish (Guatemala) - es-GT
		{
			name:   "Spanish_Guatemala_basic_number",
			amount: 8890.12,
			options: &numi18n.NumI18NOptions{
				Locale: "es-GT",
			},
			expected: "8,890.12",
		},
		{
			name:   "Spanish_Guatemala_currency",
			amount: 2800.25,
			options: &numi18n.NumI18NOptions{
				Locale: "es-GT",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Q2,800.25",
		},

		// Spanish (Honduras) - es-HN
		{
			name:   "Spanish_Honduras_basic_number",
			amount: 9901.23,
			options: &numi18n.NumI18NOptions{
				Locale: "es-HN",
			},
			expected: "9,901.23",
		},
		{
			name:   "Spanish_Honduras_currency",
			amount: 1650.80,
			options: &numi18n.NumI18NOptions{
				Locale: "es-HN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "L1,650.8",
		},

		// Spanish (Nicaragua) - es-NI
		{
			name:   "Spanish_Nicaragua_basic_number",
			amount: 1012.34,
			options: &numi18n.NumI18NOptions{
				Locale: "es-NI",
			},
			expected: "1,012.34",
		},
		{
			name:   "Spanish_Nicaragua_currency",
			amount: 950.60,
			options: &numi18n.NumI18NOptions{
				Locale: "es-NI",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "C$950.6",
		},

		// Spanish (Panama) - es-PA
		{
			name:   "Spanish_Panama_basic_number",
			amount: 1123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "es-PA",
			},
			expected: "1,123.45",
		},
		{
			name:   "Spanish_Panama_currency",
			amount: 1400.25,
			options: &numi18n.NumI18NOptions{
				Locale: "es-PA",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "B/.1,400.25",
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

func TestToFormat_Spanish_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Very_large_number_Spain",
			amount: 999999999.99,
			options: &numi18n.NumI18NOptions{
				Locale: "es-ES",
			},
			expected: "999,999,999.99",
		},
		{
			name:   "Very_small_decimal_Mexico",
			amount: 0.001,
			options: &numi18n.NumI18NOptions{
				Locale: "es-MX",
			},
			expected: "0.001",
		},
		{
			name:   "Zero_with_currency_Argentina",
			amount: 0.00,
			options: &numi18n.NumI18NOptions{
				Locale: "es-AR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "$0",
		},
		{
			name:   "Large_currency_Paraguay",
			amount: 1000000.00,
			options: &numi18n.NumI18NOptions{
				Locale: "es-PY",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "₲1,000,000",
		},
		{
			name:   "Negative_number_Colombia",
			amount: -1500.50,
			options: &numi18n.NumI18NOptions{
				Locale: "es-CO",
			},
			expected: "-1,500.5",
		},
		{
			name:   "Negative_currency_Costa_Rica",
			amount: -750.25,
			options: &numi18n.NumI18NOptions{
				Locale: "es-CR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "-₡750.25",
		},
		{
			name:   "Complex_decimal_Venezuela",
			amount: 12345.6789,
			options: &numi18n.NumI18NOptions{
				Locale: "es-VE",
			},
			expected: "12,345.6789",
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

// ========== Additional Language Tests ==========

func TestToWords_German_Locales(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "German_basic_whole_number",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "de-DE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "einhundertdreiundzwanzig Euro",
		},
		{
			name:   "German_number_with_decimals",
			amount: 45.67,
			options: &numi18n.NumI18NOptions{
				Locale: "de-DE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "fünfundvierzig Euro und siebenundsechzig Cent",
		},
		{
			name:   "German_Austria_currency",
			amount: 250.50,
			options: &numi18n.NumI18NOptions{
				Locale: "de-AT",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "zweihundertfünfzig Euro und fünfzig Cent",
		},
		{
			name:   "German_Switzerland_currency",
			amount: 100.25,
			options: &numi18n.NumI18NOptions{
				Locale: "de-CH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "einhundert Franken und fünfundzwanzig Rappen",
		},
		{
			name:   "German_negative_number",
			amount: -75.30,
			options: &numi18n.NumI18NOptions{
				Locale: "de-DE",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
				},
			},
			expected: "minus fünfundsiebzig Euro und dreißig Cent",
		},
		{
			name:   "German_zero_amount",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "de-DE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "null Euro",
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

func TestToWords_Dutch_Locales(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Dutch_basic_whole_number",
			amount: 234,
			options: &numi18n.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "tweehonderdvierendertig euro",
		},
		{
			name:   "Dutch_number_with_decimals",
			amount: 78.45,
			options: &numi18n.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "achtenzeventig euro en vijfenveertig cent",
		},
		{
			name:   "Dutch_Belgium_currency",
			amount: 150.75,
			options: &numi18n.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "honderdvijftig euro en vijfenzeventig cent",
		},
		{
			name:   "Dutch_negative_number",
			amount: -50.25,
			options: &numi18n.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
				},
			},
			expected: "min vijftig euro en vijfentwintig cent",
		},
		{
			name:   "Dutch_zero_amount",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "nl-NL",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "nul euro",
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

func TestToWords_Russian_Locales(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Russian_basic_whole_number",
			amount: 156,
			options: &numi18n.NumI18NOptions{
				Locale: "ru-RU",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "один сто пятьдесят шесть рублей",
		},
		{
			name:   "Russian_number_with_decimals",
			amount: 89.60,
			options: &numi18n.NumI18NOptions{
				Locale: "ru-RU",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "восемьдесят девять рублей и шестьдесят копеек",
		},
		{
			name:   "Russian_world_locale",
			amount: 300.15,
			options: &numi18n.NumI18NOptions{
				Locale: "ru-001",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "три сто рублей и пятнадцать копеек",
		},
		{
			name:   "Russian_Belarus_locale",
			amount: 200.50,
			options: &numi18n.NumI18NOptions{
				Locale: "ru-BY",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: " Копеек",
		},
		{
			name:   "Russian_negative_number",
			amount: -125.75,
			options: &numi18n.NumI18NOptions{
				Locale: "ru-RU",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
				},
			},
			expected: "минус один сто двадцать пять рублей и семьдесят пять копеек",
		},
		{
			name:   "Russian_zero_amount",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "ru-RU",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "ноль рублей",
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

func TestToWords_Hindi_Indian_Locales(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Hindi_basic_whole_number",
			amount: 567,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "पांच एक सौ साठ सात रुपये",
		},
		{
			name:   "Hindi_number_with_decimals",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "एक एक सौ बीस तीन रुपये और चालीस पांच पैसे",
		},
		{
			name:   "Tamil_basic_number",
			amount: 890,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "எட்டு நூறு தொண்ணூறு ரூபாய்கள்",
		},
		{
			name:   "Gujarati_basic_number",
			amount: 456,
			options: &numi18n.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "ચાર એક સો પચાસ છ રૂપિયા",
		},
		{
			name:   "Kannada_basic_number",
			amount: 789,
			options: &numi18n.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "ಏಳು ಒಂದು ನೂರು ಎಂಬತ್ತು ಒಂಬತ್ತು ರೂಪಾಯಿಗಳು",
		},
		{
			name:   "Malayalam_basic_number",
			amount: 345,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "മൂന്ന് നൂറ് നാൽപ്പത് അഞ്ച് രൂപകൾ",
		},
		{
			name:   "Hindi_negative_number",
			amount: -250.30,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
				},
			},
			expected: "ऋण दो एक सौ पचास रुपये और तीस पैसे",
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

func TestToWords_Chinese_Korean_Locales(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Chinese_basic_whole_number",
			amount: 678,
			options: &numi18n.NumI18NOptions{
				Locale: "zh-CN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "六 一百 七十 八元",
		},
		{
			name:   "Chinese_number_with_decimals",
			amount: 234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "zh-CN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "二 一百 三十 四元和五十六角",
		},
		{
			name:   "Chinese_Hong_Kong_currency",
			amount: 890.25,
			options: &numi18n.NumI18NOptions{
				Locale: "zh-HK",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "八 一百 九十港元同二十五仙",
		},
		{
			name:   "Chinese_Taiwan_currency",
			amount: 567.75,
			options: &numi18n.NumI18NOptions{
				Locale: "zh-TW",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "五 一百 六十 七新台幣和七十五分",
		},
		{
			name:   "Korean_basic_whole_number",
			amount: 456,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "사 일백 오십 육 원",
		},
		{
			name:   "Korean_number_with_decimals",
			amount: 789.60,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "칠 일백 팔십 구 원 그리고 육십 전",
		},
		{
			name:   "Chinese_negative_number",
			amount: -150.25,
			options: &numi18n.NumI18NOptions{
				Locale: "zh-CN",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
				},
			},
			expected: "负一 一百 五十元和二十五角",
		},
		{
			name:   "Korean_negative_number",
			amount: -300.50,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
				},
			},
			expected: "마이너스 삼 일백 원 그리고 오십 전",
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

func TestToWords_Portuguese_Brazilian_Locales(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Portuguese_Brazil_basic_number",
			amount: 345,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-BR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "três cem quarenta cinco Reais",
		},
		{
			name:   "Portuguese_Brazil_with_decimals",
			amount: 178.90,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-BR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "um cem setenta oito Reais e noventa centavos",
		},
		{
			name:   "Portuguese_Portugal_basic_number",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-PT",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "dois cem cinquenta seis Euros",
		},
		{
			name:   "Portuguese_Angola_basic_number",
			amount: 890,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-AO",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "oito cem noventa Kwanzas",
		},
		{
			name:   "Portuguese_Mozambique_basic_number",
			amount: 567,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-MZ",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "cinco cem sessenta sete Meticais",
		},
		{
			name:   "Portuguese_negative_number",
			amount: -75.50,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-BR",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
				},
			},
			expected: "menos setenta e cinco Reais e cinquenta centavos",
		},
		{
			name:   "Portuguese_zero_amount",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "pt-BR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "zero Reais",
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

func TestToWords_Arabic_Middle_Eastern_Locales(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Arabic_Saudi_basic_number",
			amount: 234,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-SA",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "اثنان مئة ثلاثون أربعة ريالات سعودية",
		},
		{
			name:   "Arabic_UAE_basic_number",
			amount: 567,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-AE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "خمسة مئة ستون سبعة درهم",
		},
		{
			name:   "Arabic_Egypt_basic_number",
			amount: 345,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-EG",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "ثلاثة مئة أربعون خمسة جنيهات",
		},
		{
			name:   "Arabic_Kuwait_basic_number",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-KW",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "واحد مئة عشرون ثلاثة دنانير",
		},
		{
			name:   "Arabic_Jordan_basic_number",
			amount: 789,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-JO",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "سبعة مئة ثمانون تسعة دنانير",
		},
		{
			name:   "Arabic_with_decimals",
			amount: 156.75,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-SA",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "واحد مئة خمسون ستة ريالات سعودية و سبعون خمسة هللات",
		},
		{
			name:   "Arabic_negative_number",
			amount: -100.50,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-EG",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
				},
			},
			expected: "ناقص مئة جنيهات و خمسون قروش",
		},
		{
			name:   "Persian_Iran_basic_number",
			amount: 456,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "چهار صد پنجاه شش ریال",
		},
		{
			name:   "Hebrew_Israel_basic_number",
			amount: 678,
			options: &numi18n.NumI18NOptions{
				Locale: "he-IL",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "שש מאה שבעים שמונה שקלים",
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

func TestToWords_Nordic_Scandinavian_Locales(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Swedish_basic_number",
			amount: 234,
			options: &numi18n.NumI18NOptions{
				Locale: "sv-SE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "två etthundra trettio fyra Kronor",
		},
		{
			name:   "Norwegian_basic_number",
			amount: 567,
			options: &numi18n.NumI18NOptions{
				Locale: "nb-NO",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Five Hundred Sixty Seven Dollars",
		},
		{
			name:   "Danish_basic_number",
			amount: 345,
			options: &numi18n.NumI18NOptions{
				Locale: "da-DK",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Tre Hundrede Fyrre Fem Kroner",
		},
		{
			name:   "Finnish_basic_number",
			amount: 456,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Neljä Sata Viisikymmentä Kuusi Euroa",
		},
		{
			name:   "Icelandic_basic_number",
			amount: 789,
			options: &numi18n.NumI18NOptions{
				Locale: "is-IS",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Sjö Hundrað Áttatíu Níu Krónur",
		},
		{
			name:   "Swedish_with_decimals",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "sv-SE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "ett etthundra tjugo tre Kronor och fyrtiofem Öre",
		},
		{
			name:   "Norwegian_negative_number",
			amount: -200.75,
			options: &numi18n.NumI18NOptions{
				Locale: "nb-NO",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
				},
			},
			expected: "Minus Two Hundred Dollars And Seventy Five Cents",
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

func TestToWords_Polish_Eastern_European_Locales(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Polish_basic_number",
			amount: 345,
			options: &numi18n.NumI18NOptions{
				Locale: "pl-PL",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "trzy sto czterdzieści pięć Złote",
		},
		{
			name:   "Czech_basic_number",
			amount: 567,
			options: &numi18n.NumI18NOptions{
				Locale: "cs-CZ",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Pět Sto Šedesát Sedm Koruny",
		},
		{
			name:   "Slovak_basic_number",
			amount: 234,
			options: &numi18n.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "dva sto tridsať štyri Eurá",
		},
		{
			name:   "Hungarian_basic_number",
			amount: 789,
			options: &numi18n.NumI18NOptions{
				Locale: "hu-HU",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Hét Egy száz Nyolcvan Kilenc Forint",
		},
		{
			name:   "Croatian_basic_number",
			amount: 456,
			options: &numi18n.NumI18NOptions{
				Locale: "hr-HR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Četiri Sto Pedeset Šest Eura",
		},
		{
			name:   "Polish_with_decimals",
			amount: 123.50,
			options: &numi18n.NumI18NOptions{
				Locale: "pl-PL",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "jeden sto dwadzieścia trzy Złote i pięćdziesiąt Grosze",
		},
		{
			name:   "Czech_negative_number",
			amount: -150.75,
			options: &numi18n.NumI18NOptions{
				Locale: "cs-CZ",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
				},
			},
			expected: "mínus Jedna Sto Padesát Koruny a Sedmdesát Pět Haléře",
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

func TestToWords_African_Locales(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Afrikaans_basic_number",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Honderd Drie En Twintig Rande",
		},
		{
			name:   "Amharic_basic_number",
			amount: 456,
			options: &numi18n.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "አራት መቶ ሃምሳ ስድስት ብር",
		},
		{
			name:   "Hausa_basic_number",
			amount: 234,
			options: &numi18n.NumI18NOptions{
				Locale: "ha-NG",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Biyu Dari daya Talatin Hudu Naira",
		},
		{
			name:   "Swahili_basic_number",
			amount: 567,
			options: &numi18n.NumI18NOptions{
				Locale: "sw-TZ",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Five Hundred Sixty Seven Dollars",
		},
		{
			name:   "Yoruba_basic_number",
			amount: 345,
			options: &numi18n.NumI18NOptions{
				Locale: "yo-NG",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Three Hundred Forty Five Dollars",
		},
		{
			name:   "Malagasy_basic_number",
			amount: 789,
			options: &numi18n.NumI18NOptions{
				Locale: "mg-MG",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "fito zato valopolo sivy Ariary",
		},
		{
			name:   "Afrikaans_with_decimals",
			amount: 89.50,
			options: &numi18n.NumI18NOptions{
				Locale: "af-ZA",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "Nege En Tagtig Rande En Vyftig Sente",
		},
		{
			name:   "Amharic_negative_number",
			amount: -100.25,
			options: &numi18n.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
				},
			},
			expected: "አልፎ መቶ ብር እና ሃያ አምስት ሚልስ",
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

func TestToWords_Southeast_Asian_Locales(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Thai_basic_number",
			amount: 234,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "สอง หนึ่งร้อย สามสิบ สี่บาท",
		},
		{
			name:   "Vietnamese_basic_number",
			amount: 567,
			options: &numi18n.NumI18NOptions{
				Locale: "vi-VN",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "năm một trăm sáu mươi bảy đồng",
		},
		{
			name:   "Indonesian_basic_number",
			amount: 345,
			options: &numi18n.NumI18NOptions{
				Locale: "id-ID",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Tiga Ratus Empat Puluh Lima Rupiah",
		},
		{
			name:   "Malay_basic_number",
			amount: 456,
			options: &numi18n.NumI18NOptions{
				Locale: "ms-MY",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "empat seratus lima puluh enam Ringgit",
		},
		{
			name:   "Javanese_basic_number",
			amount: 789,
			options: &numi18n.NumI18NOptions{
				Locale: "jv-ID",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Pitu Satus Wolung Puluh Sanga Rupiah",
		},
		{
			name:   "Khmer_basic_number",
			amount: 234,
			options: &numi18n.NumI18NOptions{
				Locale: "km-KH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "ពីរ មួយ​រយ សាម​ស៊ប បួន រៀល",
		},
		{
			name:   "Lao_basic_number",
			amount: 567,
			options: &numi18n.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "ຫ້າ ຮ້ອຍ ຫົກສິບ ເຈັດ Kip",
		},
		{
			name:   "Thai_with_decimals",
			amount: 123.50,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "หนึ่ง หนึ่งร้อย ยี่สิบ สามบาทและห้าสิบสตางค์",
		},
		{
			name:   "Vietnamese_negative_number",
			amount: -200.25,
			options: &numi18n.NumI18NOptions{
				Locale: "vi-VN",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
				},
			},
			expected: "âm hai trăm đồng và hai mươi lăm xu",
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

func TestToWords_Central_Asian_Caucasus_Locales(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Turkish_basic_number",
			amount: 234,
			options: &numi18n.NumI18NOptions{
				Locale: "tr-TR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "iki yüz otuz dört lira",
		},
		{
			name:   "Azerbaijani_basic_number",
			amount: 567,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Beş Yüz Altmış Yeddi Manat",
		},
		{
			name:   "Georgian_basic_number",
			amount: 345,
			options: &numi18n.NumI18NOptions{
				Locale: "ka-GE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "სამი Ერთი Ასი ორმოცი ხუთი ლარი",
		},
		{
			name:   "Armenian_basic_number",
			amount: 456,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Չորս Հարյուր Հիսուն Վեց Դրամ",
		},
		{
			name:   "Kazakh_basic_number",
			amount: 789,
			options: &numi18n.NumI18NOptions{
				Locale: "kk-KZ",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Жеті Бір Жүз Сексен Тоғыз Теңге",
		},
		{
			name:   "Kyrgyz_basic_number",
			amount: 234,
			options: &numi18n.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Эки Бир жүз Отуз Төрт Сом",
		},
		{
			name:   "Uzbek_basic_number",
			amount: 567,
			options: &numi18n.NumI18NOptions{
				Locale: "uz-UZ",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "besh bir yuz oltmish yetti So'm",
		},
		{
			name:   "Turkish_with_decimals",
			amount: 123.50,
			options: &numi18n.NumI18NOptions{
				Locale: "tr-TR",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "bir yüz yirmi üç lira ve elli kuruş",
		},
		{
			name:   "Georgian_negative_number",
			amount: -200.25,
			options: &numi18n.NumI18NOptions{
				Locale: "ka-GE",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
				},
			},
			expected: "მინუს ორი Ერთი Ასი ლარი და ოცი ხუთი თეთრი",
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

func TestToWords_Celtic_Basque_Locales(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Irish_basic_number",
			amount: 234,
			options: &numi18n.NumI18NOptions{
				Locale: "ga-IE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Dó Céad Tríocha Ceithre Eorónna",
		},
		{
			name:   "Welsh_basic_number",
			amount: 567,
			options: &numi18n.NumI18NOptions{
				Locale: "cy-GB",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Pum Can Chwedeg Saith Puntiau",
		},
		{
			name:   "Scottish_Gaelic_basic_number",
			amount: 345,
			options: &numi18n.NumI18NOptions{
				Locale: "gd-GB",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Trì Ceud Dà fhichead Còig Puinnd Sasannach",
		},
		{
			name:   "Basque_basic_number",
			amount: 456,
			options: &numi18n.NumI18NOptions{
				Locale: "eu-ES",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Lau Ehun Berrogeita Hamar Sei Euroak",
		},
		{
			name:   "Galician_basic_number",
			amount: 789,
			options: &numi18n.NumI18NOptions{
				Locale: "gl-ES",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Sete Cen Oitenta Nove Euros",
		},
		{
			name:   "Irish_with_decimals",
			amount: 123.50,
			options: &numi18n.NumI18NOptions{
				Locale: "ga-IE",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "Aon Céad Fiche Trí Eorónna Agus Caoga Ceinteannaí",
		},
		{
			name:   "Welsh_negative_number",
			amount: -200.25,
			options: &numi18n.NumI18NOptions{
				Locale: "cy-GB",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
				},
			},
			expected: "minus Dau Can Puntiau a Ugain Pum Ceiniogau",
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

func TestToWords_Constructed_Minor_Locales(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Esperanto_basic_number",
			amount: 234,
			options: &numi18n.NumI18NOptions{
				Locale: "eo-001",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Du Cent Tridek Kvar Eŭroj",
		},
		{
			name:   "Greenlandic_basic_number",
			amount: 567,
			options: &numi18n.NumI18NOptions{
				Locale: "kl-GL",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Tallimat Ataaseq hundrede Arfersanik Arfineq-marluk Danske kroner",
		},
		{
			name:   "Luxembourgish_basic_number",
			amount: 345,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "dräi honnert véierzeg fënnef Euroen",
		},
		{
			name:   "Maltese_basic_number",
			amount: 456,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "erbgħa mija ħamsin sitta Ewro",
		},
		{
			name:   "Faroese_basic_number",
			amount: 789,
			options: &numi18n.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
				},
			},
			expected: "Sjey Eitt hundrað Áttati Níggju Krónur",
		},
		{
			name:   "Esperanto_with_decimals",
			amount: 123.50,
			options: &numi18n.NumI18NOptions{
				Locale: "eo-001",
				WordDetails: &numi18n.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "Unu Cent Dudek Tri Eŭroj Kaj Kvindek Centoj",
		},
		{
			name:   "Maltese_negative_number",
			amount: -100.25,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
				},
			},
			expected: "inqas Mija waħda Ewro u għoxrin u ħamsa Ċenteżmi",
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
