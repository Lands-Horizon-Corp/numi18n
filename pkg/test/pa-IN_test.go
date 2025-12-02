package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_Punjabi_Numbers(t *testing.T) {
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
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ਸਿਫ਼ਰ",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ਪੰਜ",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ਪੰਦਰਾਂ",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ਤੀਹ",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ਸਤਤਾਲੀ",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ਇੱਕ ਸੌ",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ਦੋ ਇੱਕ ਸੌ ਪੰਜਾਹ ਛੇ",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ਇੱਕ ਹਜ਼ਾਰ",
		},
		{
			name:   "One hundred thousand (exact mapping)",
			amount: 100000,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ਇੱਕ ਲੱਖ",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ਇੱਕ ਦਸ ਲੱਖ ਦੋ ਇੱਕ ਸੌ ਤੀਹ ਚਾਰ ਇੱਕ ਹਜ਼ਾਰ ਪੰਜ ਇੱਕ ਸੌ ਸੱਠ ਸੱਤ",
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

func TestToWords_Punjabi_Currency(t *testing.T) {
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
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ਇੱਕ ਭਾਰਤੀ ਰੁਪਇਆ",
		},
		{
			name:   "Multiple rupees",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ਪੰਜ ਭਾਰਤੀ ਰੁਪਏ",
		},
		{
			name:   "Zero rupees",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ਸਿਫ਼ਰ ਭਾਰਤੀ ਰੁਪਏ",
		},
		{
			name:   "Large amount",
			amount: 100000,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ਇੱਕ ਲੱਖ ਭਾਰਤੀ ਰੁਪਏ",
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

func TestToWords_Punjabi_Decimals(t *testing.T) {
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
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ਪੰਜ ਭਾਰਤੀ ਰੁਪਏ ਅਤੇ ਇੱਕ ਪੈਸਾ",
		},
		{
			name:   "Rupees and multiple paisa",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ਪੰਜ ਭਾਰਤੀ ਰੁਪਏ ਅਤੇ ਪੱਚੀਸ ਪੈਸੇ",
		},
		{
			name:   "Only paisa",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ਸਿਫ਼ਰ ਭਾਰਤੀ ਰੁਪਏ ਅਤੇ ਨਿੰਨਾਣਵੇਂ ਪੈਸੇ",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ਇੱਕ ਇੱਕ ਹਜ਼ਾਰ ਦੋ ਇੱਕ ਸੌ ਤੀਹ ਚਾਰ ਭਾਰਤੀ ਰੁਪਏ ਅਤੇ ਛਪਨ ਪੈਸੇ",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ਇੱਕ ਇੱਕ ਸੌ ਵੀਹ ਤਿੰਨ ਅਤੇ ਪੰਤਾਲੀ",
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

func TestToWords_Punjabi_Negative(t *testing.T) {
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
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ਮਾਇਨਸ ਪੰਜਾਹ",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ਮਾਇਨਸ ਪੱਚੀਸ ਭਾਰਤੀ ਰੁਪਏ ਅਤੇ ਪਚਿਹੱਤਰ ਪੈਸੇ",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "ਰਿਣ",
					},
				},
			},
			expected: "ਰਿਣ ਇੱਕ ਸੌ ਭਾਰਤੀ ਰੁਪਏ",
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

func TestToWords_Punjabi_Formatting(t *testing.T) {
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
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "ਇੱਕ ਇੱਕ ਸੌ ਵੀਹ ਤਿੰਨ ਭਾਰਤੀ ਰੁਪਏ",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "ਇੱਕ ਇੱਕ ਸੌ ਵੀਹ ਤਿੰਨ ਭਾਰਤੀ ਰੁਪਏ",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "ਨੌ ਇੱਕ ਸੌ ਨੱਬੇ ਨੌ ਸਿਰਫ਼",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "ਪੰਜ ਸੌ ਭਾਰਤੀ ਰੁਪਏ ਸਿਰਫ਼",
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

func TestToWords_Punjabi_CustomCurrency(t *testing.T) {
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
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "ਡਾਲਰ",
						Plural: "ਡਾਲਰ",
					},
				},
			},
			expected: "ਇੱਕ ਸੌ ਡਾਲਰ",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "ਯੂਰੋ",
						Plural:           "ਯੂਰੋ",
						FractionUnitName: "ਸੈਂਟ",
						FractionPlural:   "ਸੈਂਟ",
					},
				},
			},
			expected: "ਪੰਜਾਹ ਯੂਰੋ ਅਤੇ ਪੱਚੀਸ ਸੈਂਟ",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "ਪਾਊਂਡ",
						Plural: "ਪਾਊਂਡ",
					},
				},
			},
			expected: "ਇੱਕ ਪਾਊਂਡ",
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

func TestToWords_Punjabi_EdgeCases(t *testing.T) {
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
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ਸਿਫ਼ਰ ਭਾਰਤੀ ਰੁਪਏ ਅਤੇ ਇੱਕ ਪੈਸਾ",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ਗਿਆਰਾਂ",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ਬਾਰਾਂ",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ਇੱਕੀਸ",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ਇੱਕ ਇੱਕ ਸੌ ਇੱਕ",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "pa-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ਇੱਕ ਇੱਕ ਹਜ਼ਾਰ ਇੱਕ",
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
