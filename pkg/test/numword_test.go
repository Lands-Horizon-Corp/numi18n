package pkg

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_EnglishUS(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Basic whole number",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "One Hundred Twenty Three Dollars",
		},
		{
			name:   "Number with decimals",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &pkg.WordDetails{
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
			options: &pkg.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &pkg.WordDetails{
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
			options: &pkg.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Zero Dollars",
		},
		{
			name:   "With override options",
			amount: 100.50,
			options: &pkg.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
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
			options: &pkg.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &pkg.WordDetails{
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
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Basic whole number",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency: true,
				},
			},
			expected: "百二十三円",
		},
		{
			name:   "Number with decimals",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "百二十三円と四十五銭",
		},
		{
			name:   "Large number",
			amount: 10000,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency: true,
				},
			},
			expected: "一万円",
		},
		{
			name:   "Negative number",
			amount: -50,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
				},
			},
			expected: "マイナス五十円",
		},
		{
			name:   "Zero amount",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency: true,
				},
			},
			expected: "零円",
		},
		{
			name:   "With override options",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency: true,
					OverrideOptions: &pkg.OverrideOptions{
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
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Basic whole number",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &pkg.WordDetails{
					Currency: true,
				},
			},
			expected: "Isang daan Dalawampu Tatlo Piso",
		},
		{
			name:   "Number with decimals",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &pkg.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "Isang daan Dalawampu Tatlo Piso at Apatnapu Lima Sentimo",
		},
		{
			name:   "Single peso",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &pkg.WordDetails{
					Currency: true,
				},
			},
			expected: "Isa Peso",
		},
		{
			name:   "Negative number",
			amount: -50.25,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &pkg.WordDetails{
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
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &pkg.WordDetails{
					Currency: true,
				},
			},
			expected: "Zero Piso",
		},
		{
			name:   "With only flag",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &pkg.WordDetails{
					Currency: true,
					Only:     true,
				},
			},
			expected: "Lima Daan Piso lamang",
		},
		{
			name:   "With override options",
			amount: 100.50,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &pkg.WordDetails{
					Currency: true,
					Decimal:  true,
					OverrideOptions: &pkg.OverrideOptions{
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
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Very small decimal",
			amount: 0.01,
			options: &pkg.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &pkg.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "Zero Dollars And One Cent",
		},
		{
			name:   "Large number",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &pkg.WordDetails{
					Currency: true,
				},
			},
			expected: "One Million Dollars",
		},
		{
			name:   "No locale fallback to US",
			amount: 100,
			options: &pkg.NumI18NOptions{
				WordDetails: &pkg.WordDetails{
					Currency: true,
				},
			},
			expected: "One Hundred Dollars",
		},
		{
			name:   "Currency matching fallback",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Currency: "USD",
				WordDetails: &pkg.WordDetails{
					Currency: true,
				},
			},
			expected: "One Hundred Dollars",
		},
		{
			name:   "Uppercase formatting",
			amount: 50,
			options: &pkg.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "FIFTY DOLLARS",
		},
		{
			name:   "Lowercase formatting",
			amount: 50,
			options: &pkg.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &pkg.WordDetails{
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
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Custom negative word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "negative",
					},
				},
			},
			expected: "negative One Hundred Dollars",
		},
		{
			name:   "All overrides combined",
			amount: -123.45,
			options: &pkg.NumI18NOptions{
				Locale: "en-US",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
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
