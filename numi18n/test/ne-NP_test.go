package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_Nepali_Numbers(t *testing.T) {
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
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "शून्य",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "पाँच",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "पन्ध्र",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "तिस",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "सत्चालिस",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "एक सय",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "दुई एक सय पचास छ",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "एक हजार",
		},
		{
			name:   "One lakh (hundred thousand)",
			amount: 100000,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "एक लाख",
		},
		{
			name:   "Large complex number",
			amount: 123456,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "एक एक सय बिस तीन एक हजार चार एक सय पचास छ",
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

func TestToWords_Nepali_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One rupee",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "एक रुपैया",
		},
		{
			name:   "Multiple rupees",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "पाँच रुपैया",
		},
		{
			name:   "Zero rupees",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "शून्य रुपैया",
		},
		{
			name:   "Large amount",
			amount: 100000,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "एक लाख रुपैया",
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

func TestToWords_Nepali_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Rupees and one paisa",
			amount: 5.01,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "पाँच रुपैया र एक पैसा",
		},
		{
			name:   "Rupees and multiple paisa",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "पाँच रुपैया र पच्चिस पैसा",
		},
		{
			name:   "Only paisa",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "शून्य रुपैया र उनान्सय पैसा",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "एक एक हजार दुई एक सय तिस चार रुपैया र छप्पन्न पैसा",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "एक एक सय बिस तीन र पैंतालिस",
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

func TestToWords_Nepali_Negative(t *testing.T) {
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
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "माइनस पचास",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "माइनस पच्चिस रुपैया र पचहत्तर पैसा",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &numi18n.OverrideOptions{
						NegativeWord: "नकारात्मक",
					},
				},
			},
			expected: "नकारात्मक एक सय रुपैया",
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

func TestToWords_Nepali_Formatting(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Without capitalization",
			amount: 42,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Capitalize: false,
				},
			},
			expected: "बयालिस",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "नौ एक सय नब्बे नौ मात्र",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "पाँच सय रुपैया मात्र",
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

func TestToWords_Nepali_CustomCurrency(t *testing.T) {
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
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "डलर",
						Plural: "डलर",
					},
				},
			},
			expected: "एक सय डलर",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "युरो",
						Plural:           "युरो",
						FractionUnitName: "सेन्ट",
						FractionPlural:   "सेन्ट",
					},
				},
			},
			expected: "पचास युरो र पच्चिस सेन्ट",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "पाउन्ड",
						Plural: "पाउन्ड",
					},
				},
			},
			expected: "एक पाउन्ड",
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

func TestToWords_Nepali_EdgeCases(t *testing.T) {
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
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "शून्य रुपैया र एक पैसा",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "एघार",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "बाह्र",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "एकाइस",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "एक एक सय एक",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "एक एक हजार एक",
		},
		{
			name:   "One lakh system",
			amount: 150000,
			options: &numi18n.NumI18NOptions{
				Locale: "ne-NP",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "एक एक सय पचास एक हजार",
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
