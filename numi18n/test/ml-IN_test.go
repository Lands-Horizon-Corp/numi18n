package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_MalayalamIndia_Numbers(t *testing.T) {
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
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "പൂജ്യം",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "അഞ്ച്",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "പതിനഞ്ച്",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "മുപ്പത്",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "നാൽപ്പത്തേഴ്",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ഒരു നൂറ്",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "രണ്ട് നൂറ് അൻപത് ആറ്",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ആയിരം",
		},
		{
			name:   "One lakh (exact mapping)",
			amount: 100000,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ലക്ഷം",
		},
		{
			name:   "One crore (exact mapping)",
			amount: 10000000,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "പത്ത് ലക്ഷം",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ഒന്ന് ലക്ഷം രണ്ട് നൂറ് മുപ്പത് നാല് ആയിരം അഞ്ച് നൂറ് അറുപത് ഏഴ്",
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

func TestToWords_MalayalamIndia_Currency(t *testing.T) {
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
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ഒന്ന് രൂപ",
		},
		{
			name:   "Multiple rupees",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "അഞ്ച് രൂപകൾ",
		},
		{
			name:   "Zero rupees",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "പൂജ്യം രൂപകൾ",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ലക്ഷം രൂപകൾ",
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

func TestToWords_MalayalamIndia_Decimals(t *testing.T) {
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
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "അഞ്ച് രൂപകൾ ഉം ഒന്ന് പൈസ",
		},
		{
			name:   "Rupees and multiple paise",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "അഞ്ച് രൂപകൾ ഉം ഇരുപത്തഞ്ച് പൈസകൾ",
		},
		{
			name:   "Only paise",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "പൂജ്യം രൂപകൾ ഉം തൊണ്ണൂറ്റൊമ്പത് പൈസകൾ",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ഒന്ന് ആയിരം രണ്ട് നൂറ് മുപ്പത് നാല് രൂപകൾ ഉം അൻപത്താറ് പൈസകൾ",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ഒന്ന് നൂറ് ഇരുപത് മൂന്ന് ഉം നാൽപ്പത്തഞ്ച്",
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

func TestToWords_MalayalamIndia_Negative(t *testing.T) {
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
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "കുറച്ച് അൻപത്",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "കുറച്ച് ഇരുപത്തഞ്ച് രൂപകൾ ഉം എഴുപത്തഞ്ച് പൈസകൾ",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &numi18n.OverrideOptions{
						NegativeWord: "നെഗറ്റീവ്",
					},
				},
			},
			expected: "നെഗറ്റീവ് ഒരു നൂറ് രൂപകൾ",
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

func TestToWords_MalayalamIndia_Formatting(t *testing.T) {
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
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "ഒന്ന് നൂറ് ഇരുപത് മൂന്ന് രൂപകൾ",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "ഒന്ന് നൂറ് ഇരുപത് മൂന്ന് രൂപകൾ",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "ഒമ്പത് നൂറ് തൊണ്ണൂറ് ഒമ്പത് മാത്രം",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "അഞ്ഞൂറ് രൂപകൾ മാത്രം",
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

func TestToWords_MalayalamIndia_CustomCurrency(t *testing.T) {
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
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "ഡോളർ",
						Plural: "ഡോളർസ്",
					},
				},
			},
			expected: "ഒരു നൂറ് ഡോളർസ്",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "യൂറോ",
						Plural:           "യൂറോസ്",
						FractionUnitName: "സെന്റ്",
						FractionPlural:   "സെന്റുകൾ",
					},
				},
			},
			expected: "അൻപത് യൂറോസ് ഉം ഇരുപത്തഞ്ച് സെന്റുകൾ",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "പൗണ്ട്",
						Plural: "പൗണ്ടുകൾ",
					},
				},
			},
			expected: "ഒന്ന് പൗണ്ട്",
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

func TestToWords_MalayalamIndia_EdgeCases(t *testing.T) {
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
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "പൂജ്യം രൂപകൾ ഉം ഒന്ന് പൈസ",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "പതിനൊന്ന്",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "പന്ത്രണ്ട്",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ഇരുപത്തൊന്ന്",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ഒന്ന് നൂറ് ഒന്ന്",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &numi18n.NumI18NOptions{
				Locale: "ml-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ഒന്ന് ആയിരം ഒന്ന്",
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
