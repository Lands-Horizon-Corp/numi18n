package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_Hindi_Numbers(t *testing.T) {
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
				Locale: "hi-IN",
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
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "पांच",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "पंद्रह",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
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
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "चालीस सात",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "एक सौ",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "दो एक सौ पचास छह",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "एक हजार",
		},
		{
			name:   "One lakh (exact mapping)",
			amount: 100000,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "एक लाख",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "दस लाख",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "एक दस लाख दो एक सौ तीस चार एक हजार पांच एक सौ साठ सात",
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

func TestToWords_Hindi_Currency(t *testing.T) {
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
				Locale: "hi-IN",
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
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "पांच रुपये",
		},
		{
			name:   "Zero rupees",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
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
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "दस लाख रुपये",
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

func TestToWords_Hindi_Decimals(t *testing.T) {
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
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "पांच रुपये और एक पैसा",
		},
		{
			name:   "Rupees and multiple paisa",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "पांच रुपये और बीस पांच पैसे",
		},
		{
			name:   "Only paisa",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "शून्य रुपये और नब्बे नौ पैसे",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "एक एक हजार दो एक सौ तीस चार रुपये और पचास छह पैसे",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "एक एक सौ बीस तीन और चालीस पांच",
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

func TestToWords_Hindi_Negative(t *testing.T) {
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
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ऋण पचास",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ऋण बीस पांच रुपये और सत्तर पांच पैसे",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &numi18n.OverrideOptions{
						NegativeWord: "नकारात्मक",
					},
				},
			},
			expected: "नकारात्मक एक सौ रुपये",
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

func TestToWords_Hindi_Formatting(t *testing.T) {
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
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "एक एक सौ बीस तीन रुपये",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "एक एक सौ बीस तीन रुपये",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "नौ एक सौ नब्बे नौ केवल",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "पांच एक सौ रुपये केवल",
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

func TestToWords_Hindi_CustomCurrency(t *testing.T) {
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
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "डॉलर",
						Plural: "डॉलर",
					},
				},
			},
			expected: "एक सौ डॉलर",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "यूरो",
						Plural:           "यूरो",
						FractionUnitName: "सेंट",
						FractionPlural:   "सेंट",
					},
				},
			},
			expected: "पचास यूरो और बीस पांच सेंट",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "पाउंड",
						Plural: "पाउंड",
					},
				},
			},
			expected: "एक पाउंड",
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

func TestToWords_Hindi_EdgeCases(t *testing.T) {
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
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "शून्य रुपये और एक पैसा",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ग्यारह",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "बारह",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "बीस एक",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "एक एक सौ एक",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &numi18n.NumI18NOptions{
				Locale: "hi-IN",
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
