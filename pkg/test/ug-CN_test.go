package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_UyghurChina_Numbers(t *testing.T) {
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
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "نۆل",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "بەش",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ئون بەش",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ئوتتۇز",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "قىرىق يەتتە",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "بىر يۈز",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ئىككى بىر يۈز ئەللىك ئالتە",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "بىر مىڭ",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "بىر مىليون",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "بىر بىر مىليون ئىككى بىر يۈز ئوتتۇز تۆت بىر مىڭ بەش بىر يۈز ئاتمىش يەتتە",
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

func TestToWords_UyghurChina_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One yuan",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "بىر يۈەن",
		},
		{
			name:   "Multiple yuan",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "بەش يۈەن",
		},
		{
			name:   "Zero yuan",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "نۆل يۈەن",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "بىر مىليون يۈەن",
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

func TestToWords_UyghurChina_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Yuan and one jiao",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "بەش يۈەن ۋە بىر جىياو",
		},
		{
			name:   "Yuan and multiple jiao",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "بەش يۈەن ۋە يىگىرمە بەش جىياو",
		},
		{
			name:   "Only jiao",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "نۆل يۈەن ۋە توقسان توققۇز جىياو",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "بىر بىر مىڭ ئىككى بىر يۈز ئوتتۇز تۆت يۈەن ۋە ئەللىك ئالتە جىياو",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "بىر بىر يۈز يىگىرمە ئۈچ ۋە قىرىق بەش",
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

func TestToWords_UyghurChina_Negative(t *testing.T) {
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
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "مىنۇس ئەللىك",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "مىنۇس يىگىرمە بەش يۈەن ۋە يەتمىش بەش جىياو",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "مەنپىي",
					},
				},
			},
			expected: "مەنپىي بىر يۈز يۈەن",
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

func TestToWords_UyghurChina_Formatting(t *testing.T) {
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
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "بىر بىر يۈز يىگىرمە ئۈچ يۈەن",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "بىر بىر يۈز يىگىرمە ئۈچ يۈەن",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "توققۇز بىر يۈز توقسان توققۇز پەقەت",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "بەش يۈز يۈەن پەقەت",
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

func TestToWords_UyghurChina_CustomCurrency(t *testing.T) {
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
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "دوللار",
						Plural: "دوللار",
					},
				},
			},
			expected: "بىر يۈز دوللار",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "ئېۋرو",
						Plural:           "ئېۋرو",
						FractionUnitName: "سەنت",
						FractionPlural:   "سەنت",
					},
				},
			},
			expected: "ئەللىك ئېۋرو ۋە يىگىرمە بەش سەنت",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "پاۋند",
						Plural: "پاۋند",
					},
				},
			},
			expected: "بىر پاۋند",
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

func TestToWords_UyghurChina_EdgeCases(t *testing.T) {
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
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "نۆل يۈەن ۋە بىر جىياو",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ئون بىر",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ئون ئىككى",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "يىگىرمە بىر",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "بىر بىر يۈز بىر",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "ug-CN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "بىر بىر مىڭ بىر",
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
