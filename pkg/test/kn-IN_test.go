package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_KannadaIndia_Numbers(t *testing.T) {
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
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ಶೂನ್ಯ",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ಐದು",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ಹದಿನೈದು",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ಮೂವತ್ತು",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ನಲವತ್ತು ಏಳು",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ಒಂದು ನೂರು",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ಎರಡು ಒಂದು ನೂರು ಐವತ್ತು ಆರು",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ಒಂದು ಸಾವಿರ",
		},
		{
			name:   "One lakh (exact mapping)",
			amount: 100000,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ಒಂದು ಲಕ್ಷ",
		},
		{
			name:   "One crore (exact mapping)",
			amount: 10000000,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ಒಂದು ಕೋಟಿ",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ಒಂದು ಹತ್ತು ಲಕ್ಷ ಎರಡು ಒಂದು ನೂರು ಮೂವತ್ತು ನಾಲ್ಕು ಒಂದು ಸಾವಿರ ಐದು ಒಂದು ನೂರು ಅರುವತ್ತು ಏಳು",
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

func TestToWords_KannadaIndia_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One rupee",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ಒಂದು ರೂಪಾಯಿ",
		},
		{
			name:   "Multiple rupees",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ಐದು ರೂಪಾಯಿಗಳು",
		},
		{
			name:   "Zero rupees",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ಶೂನ್ಯ ರೂಪಾಯಿಗಳು",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ಹತ್ತು ಲಕ್ಷ ರೂಪಾಯಿಗಳು",
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

func TestToWords_KannadaIndia_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Rupees and one paisa",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ಐದು ರೂಪಾಯಿಗಳು ಮತ್ತು ಒಂದು ಪೈಸೆ",
		},
		{
			name:   "Rupees and multiple paise",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ಐದು ರೂಪಾಯಿಗಳು ಮತ್ತು ಇಪ್ಪತ್ತು ಐದು ಪೈಸೆಗಳು",
		},
		{
			name:   "Only paise",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ಶೂನ್ಯ ರೂಪಾಯಿಗಳು ಮತ್ತು ತೊಂಬತ್ತು ಒಂಬತ್ತು ಪೈಸೆಗಳು",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ಒಂದು ಒಂದು ಸಾವಿರ ಎರಡು ಒಂದು ನೂರು ಮೂವತ್ತು ನಾಲ್ಕು ರೂಪಾಯಿಗಳು ಮತ್ತು ಐವತ್ತು ಆರು ಪೈಸೆಗಳು",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ಒಂದು ಒಂದು ನೂರು ಇಪ್ಪತ್ತು ಮೂರು ಮತ್ತು ನಲವತ್ತು ಐದು",
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

func TestToWords_KannadaIndia_Negative(t *testing.T) {
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
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ಮೈನಸ್ ಐವತ್ತು",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ಮೈನಸ್ ಇಪ್ಪತ್ತು ಐದು ರೂಪಾಯಿಗಳು ಮತ್ತು ಎಪ್ಪತ್ತು ಐದು ಪೈಸೆಗಳು",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "ಋಣಾತ್ಮಕ",
					},
				},
			},
			expected: "ಋಣಾತ್ಮಕ ಒಂದು ನೂರು ರೂಪಾಯಿಗಳು",
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

func TestToWords_KannadaIndia_Formatting(t *testing.T) {
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
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "ಒಂದು ಒಂದು ನೂರು ಇಪ್ಪತ್ತು ಮೂರು ರೂಪಾಯಿಗಳು",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "ಒಂದು ಒಂದು ನೂರು ಇಪ್ಪತ್ತು ಮೂರು ರೂಪಾಯಿಗಳು",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "ಒಂಬತ್ತು ಒಂದು ನೂರು ತೊಂಬತ್ತು ಒಂಬತ್ತು ಮಾತ್ರ",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "ಐದು ಒಂದು ನೂರು ರೂಪಾಯಿಗಳು ಮಾತ್ರ",
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

func TestToWords_KannadaIndia_CustomCurrency(t *testing.T) {
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
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "ಡಾಲರ್",
						Plural: "ಡಾಲರ್\u200cಗಳು",
					},
				},
			},
			expected: "ಒಂದು ನೂರು ಡಾಲರ್\u200cಗಳು",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "ಯುರೋ",
						Plural:           "ಯುರೋಗಳು",
						FractionUnitName: "ಸೆಂಟ್",
						FractionPlural:   "ಸೆಂಟ್\u200cಗಳು",
					},
				},
			},
			expected: "ಐವತ್ತು ಯುರೋಗಳು ಮತ್ತು ಇಪ್ಪತ್ತು ಐದು ಸೆಂಟ್\u200cಗಳು",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "ಪೌಂಡ್",
						Plural: "ಪೌಂಡ್\u200cಗಳು",
					},
				},
			},
			expected: "ಒಂದು ಪೌಂಡ್",
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

func TestToWords_KannadaIndia_EdgeCases(t *testing.T) {
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
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ಶೂನ್ಯ ರೂಪಾಯಿಗಳು ಮತ್ತು ಒಂದು ಪೈಸೆ",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ಹನ್ನೊಂದು",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ಹನ್ನೆರಡು",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ಇಪ್ಪತ್ತು ಒಂದು",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ಒಂದು ಒಂದು ನೂರು ಒಂದು",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "kn-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ಒಂದು ಒಂದು ಸಾವಿರ ಒಂದು",
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
