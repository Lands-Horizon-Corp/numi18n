package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_MarathiIndia_Numbers(t *testing.T) {
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
				Locale: "mr-IN",
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
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "पाच",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "पंधरा",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "तीस",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "सत्तेचाळीस",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "एक शंभर",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "दोन एकशे पन्नास सहा",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "एक हजार",
		},
		{
			name:   "One lakh",
			amount: 100000,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "एक लाख",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "एक दहा लाख दोन एकशे तीस चार एक हजार पाच एकशे साठ सात",
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

func TestToWords_MarathiIndia_Currency(t *testing.T) {
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
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "एक रुपया",
		},
		{
			name:   "Multiple rupees",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "पाच रुपये",
		},
		{
			name:   "Zero rupees",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "शून्य रुपये",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "दहा लाख रुपये",
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

func TestToWords_MarathiIndia_Decimals(t *testing.T) {
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
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "पाच रुपये आणि एक पैसा",
		},
		{
			name:   "Rupees and multiple paise",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "पाच रुपये आणि पंचवीस पैसे",
		},
		{
			name:   "Only paise",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "शून्य रुपये आणि नव्याण्णव पैसे",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "एक एक हजार दोन एकशे तीस चार रुपये आणि छप्पन्न पैसे",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "एक एकशे वीस तीन आणि पंचेचाळीस",
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

func TestToWords_MarathiIndia_Negative(t *testing.T) {
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
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "वजा पन्नास",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "वजा पंचवीस रुपये आणि पंचाहत्तर पैसे",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &numi18n.OverrideOptions{
						NegativeWord: "ऋण",
					},
				},
			},
			expected: "ऋण एक शंभर रुपये",
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

func TestToWords_MarathiIndia_Formatting(t *testing.T) {
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
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "एक एकशे वीस तीन रुपये",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "एक एकशे वीस तीन रुपये",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "नऊ एकशे नव्वद नऊ फक्त",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "पाचशे रुपये फक्त",
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

func TestToWords_MarathiIndia_CustomCurrency(t *testing.T) {
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
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "डॉलर",
						Plural: "डॉलर",
					},
				},
			},
			expected: "एक शंभर डॉलर",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "युरो",
						Plural:           "युरो",
						FractionUnitName: "सेंट",
						FractionPlural:   "सेंट",
					},
				},
			},
			expected: "पन्नास युरो आणि पंचवीस सेंट",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "पाऊंड",
						Plural: "पाऊंड",
					},
				},
			},
			expected: "एक पाऊंड",
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

func TestToWords_MarathiIndia_EdgeCases(t *testing.T) {
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
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "शून्य रुपये आणि एक पैसा",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "अकरा",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "बारा",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "एकवीस",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "एक एकशे एक",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &numi18n.NumI18NOptions{
				Locale: "mr-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "एक एक हजार एक",
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
