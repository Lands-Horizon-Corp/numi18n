package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_QuechuaPeru_Numbers(t *testing.T) {
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
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ch'usaq",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Pichqa",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Chunka pichqaniyuq",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Kimsa chunka",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tawa chunka qanchisniyuq",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Pachak",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Iskay pachak pichqa chunka suqta",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Waranqa",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Waranqa waranqa",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Huk waranqa waranqa iskay pachak kimsa chunka tawa waranqa pichqa pachak suqta chunka qanchis",
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

func TestToWords_QuechuaPeru_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One sol",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Huk Sol",
		},
		{
			name:   "Multiple soles",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Pichqa Soles",
		},
		{
			name:   "Zero soles",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Ch'usaq Soles",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Waranqa waranqa Soles",
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

func TestToWords_QuechuaPeru_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Soles and one céntimo",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Pichqa Soles wan huk Céntimo",
		},
		{
			name:   "Soles and multiple céntimos",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Pichqa Soles wan iskay chunka pichqaniyuq Céntimos",
		},
		{
			name:   "Only céntimos",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Ch'usaq Soles wan isqun chunka isqunniyuq Céntimos",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Huk waranqa iskay pachak kimsa chunka tawa Soles wan pichqa chunka suqtaniyuq Céntimos",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Huk pachak iskay chunka kimsa wan tawa chunka pichqaniyuq",
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

func TestToWords_QuechuaPeru_Negative(t *testing.T) {
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
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Ayqiq pichqa chunka",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Ayqiq iskay chunka pichqaniyuq Soles wan qanchis chunka pichqaniyuq Céntimos",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "Mana allin",
					},
				},
			},
			expected: "Mana allin Pachak Soles",
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

func TestToWords_QuechuaPeru_Formatting(t *testing.T) {
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
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "HUK PACHAK ISKAY CHUNKA KIMSA SOLES",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "huk pachak iskay chunka kimsa soles",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Isqun pachak isqun chunka isqun sapalla",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Pichqa pachak Soles sapalla",
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

func TestToWords_QuechuaPeru_CustomCurrency(t *testing.T) {
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
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "peso",
						Plural: "pesos",
					},
				},
			},
			expected: "Pachak pesos",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "euro",
						Plural:           "euros",
						FractionUnitName: "centavo",
						FractionPlural:   "centavos",
					},
				},
			},
			expected: "Pichqa chunka euros wan iskay chunka pichqaniyuq centavos",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "inti",
						Plural: "intis",
					},
				},
			},
			expected: "Huk inti",
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

func TestToWords_QuechuaPeru_EdgeCases(t *testing.T) {
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
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Ch'usaq Soles wan huk Céntimo",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Chunka hukniyuq",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Chunka iskayniyuq",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Iskay chunka hukniyuq",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Huk pachak huk",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "qu-PE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Huk waranqa huk",
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
