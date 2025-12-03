package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_Finnish_Numbers(t *testing.T) {
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
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Nolla",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Viisi",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Viisitoista",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Kolmekymmentä",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Neljäkymmentä Seitsemän",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Sata",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Kaksi Sata Viisikymmentä Kuusi",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tuhat",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Miljoona",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Yksi Miljoona Kaksi Sata Kolmekymmentä Neljä Tuhat Viisi Sata Kuusikymmentä Seitsemän",
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

func TestToWords_Finnish_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One euro",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Yksi Euro",
		},
		{
			name:   "Multiple euros",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Viisi Euroa",
		},
		{
			name:   "Zero euros",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Nolla Euroa",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Miljoona Euroa",
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

func TestToWords_Finnish_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Euros and one cent",
			amount: 5.01,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Viisi Euroa Ja Yksi Sentti",
		},
		{
			name:   "Euros and multiple cents",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Viisi Euroa Ja Kaksikymmentä Viisi Senttiä",
		},
		{
			name:   "Only cents",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Nolla Euroa Ja Yhdeksänkymmentä Yhdeksän Senttiä",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Yksi Tuhat Kaksi Sata Kolmekymmentä Neljä Euroa Ja Viisikymmentä Kuusi Senttiä",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Yksi Sata Kaksikymmentä Kolme Ja Neljäkymmentä Viisi",
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

func TestToWords_Finnish_Negative(t *testing.T) {
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
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Miinus Viisikymmentä",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Miinus Kaksikymmentä Viisi Euroa Ja Seitsemänkymmentä Viisi Senttiä",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &numi18n.OverrideOptions{
						NegativeWord: "Negatiivinen",
					},
				},
			},
			expected: "Negatiivinen Sata Euroa",
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

func TestToWords_Finnish_Formatting(t *testing.T) {
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
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "YKSI SATA KAKSIKYMMENTÄ KOLME EUROA",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "yksi sata kaksikymmentä kolme euroa",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Yhdeksän Sata Yhdeksänkymmentä Yhdeksän Vain",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Viisi Sata Euroa Vain",
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

func TestToWords_Finnish_CustomCurrency(t *testing.T) {
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
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "markka",
						Plural: "markkaa",
					},
				},
			},
			expected: "Sata markkaa",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "dollar",
						Plural:           "dollaria",
						FractionUnitName: "centti",
						FractionPlural:   "centtiä",
					},
				},
			},
			expected: "Viisikymmentä dollaria Ja Kaksikymmentä Viisi centtiä",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "kruunu",
						Plural: "kruunua",
					},
				},
			},
			expected: "Yksi kruunu",
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

func TestToWords_Finnish_EdgeCases(t *testing.T) {
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
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Nolla Euroa Ja Yksi Sentti",
		},
		{
			name:   "Eleven",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Yksitoista",
		},
		{
			name:   "Twelve",
			amount: 12,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Kaksitoista",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Kaksikymmentä Yksi",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Yksi Sata Yksi",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &numi18n.NumI18NOptions{
				Locale: "fi-FI",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Yksi Tuhat Yksi",
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
