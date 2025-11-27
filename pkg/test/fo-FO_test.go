package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_FaroeseIslands_Numbers(t *testing.T) {
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
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Null",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Fimm",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Fimtan",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tríati",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Fyrti Sjey",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Eitt hundrað",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tvey Eitt hundrað Hálvtríss Seks",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Eitt túsund",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ein million",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Eitt Ein million Tvey Eitt hundrað Tríati Fýra Eitt túsund Fimm Eitt hundrað Seksti Sjey",
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

func TestToWords_FaroeseIslands_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One krone",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Eitt Króna",
		},
		{
			name:   "Multiple kroner",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Fimm Krónur",
		},
		{
			name:   "Zero kroner",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Null Krónur",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Ein million Krónur",
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

func TestToWords_FaroeseIslands_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Kroner and one oyra",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Fimm Krónur Og Eitt Oyra",
		},
		{
			name:   "Kroner and multiple oyrur",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Fimm Krónur Og Tuttugu Fimm Oyrur",
		},
		{
			name:   "Only oyrur",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Null Krónur Og Níti Níggju Oyrur",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Eitt Eitt túsund Tvey Eitt hundrað Tríati Fýra Krónur Og Hálvtríss Seks Oyrur",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Eitt Eitt hundrað Tuttugu Trý Og Fyrti Fimm",
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

func TestToWords_FaroeseIslands_Negative(t *testing.T) {
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
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Mínus Hálvtríss",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Mínus Tuttugu Fimm Krónur Og Sjeyti Fimm Oyrur",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "Negativ",
					},
				},
			},
			expected: "Negativ Eitt hundrað Krónur",
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

func TestToWords_FaroeseIslands_Formatting(t *testing.T) {
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
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "EITT EITT HUNDRAÐ TUTTUGU TRÝ KRÓNUR",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "eitt eitt hundrað tuttugu trý krónur",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Níggju Eitt hundrað Níti Níggju Bara",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Fimm Eitt hundrað Krónur Bara",
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

func TestToWords_FaroeseIslands_CustomCurrency(t *testing.T) {
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
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "euro",
						Plural: "evrur",
					},
				},
			},
			expected: "Eitt hundrað evrur",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "dollar",
						Plural:           "dollarar",
						FractionUnitName: "sent",
						FractionPlural:   "sentur",
					},
				},
			},
			expected: "Hálvtríss dollarar Og Tuttugu Fimm sentur",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "pund",
						Plural: "pund",
					},
				},
			},
			expected: "Eitt pund",
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

func TestToWords_FaroeseIslands_EdgeCases(t *testing.T) {
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
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Null Krónur Og Eitt Oyra",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ellivu",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tólv",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tuttugu Eitt",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Eitt Eitt hundrað Eitt",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "fo-FO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Eitt Eitt túsund Eitt",
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
