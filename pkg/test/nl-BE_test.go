package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_DutchBelgian_Numbers(t *testing.T) {
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
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Nul",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Vijf",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Vijftien",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Dertig",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Zevenenveertig",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Honderd",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tweehonderdzesenvijftig",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Duizend",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Miljoen",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Een miljoen twee honderd dertig vier duizend vijf honderd zestig zeven",
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

func TestToWords_DutchBelgian_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One euro",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Een euro",
		},
		{
			name:   "Multiple euros",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Vijf euro",
		},
		{
			name:   "Zero euros",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Nul euro",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Miljoen euro",
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

func TestToWords_DutchBelgian_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Euros and one cent",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Vijf euro en een cent",
		},
		{
			name:   "Euros and multiple cents",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Vijf euro en vijfentwintig cent",
		},
		{
			name:   "Only cents",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Nul euro en negenennegentig cent",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Duizendtweehonderdvierendertig euro en zesenvijftig cent",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Honderddrieentwintig en vijfenveertig",
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

func TestToWords_DutchBelgian_Negative(t *testing.T) {
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
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Min vijftig",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Min vijfentwintig euro en vijfenzeventig cent",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "Negatief",
					},
				},
			},
			expected: "Negatief honderd euro",
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

func TestToWords_DutchBelgian_Formatting(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Without capitalization",
			amount: 42,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Capitalize: false,
				},
			},
			expected: "tweeenveertig",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Negenhonderdnegenennegentig alleen",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Vijfhonderd euro alleen",
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

func TestToWords_DutchBelgian_CustomCurrency(t *testing.T) {
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
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "dollar",
						Plural: "dollar",
					},
				},
			},
			expected: "Honderd dollar",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "pond",
						Plural:           "pond",
						FractionUnitName: "pence",
						FractionPlural:   "pence",
					},
				},
			},
			expected: "Vijftig pond en vijfentwintig pence",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "frank",
						Plural: "frank",
					},
				},
			},
			expected: "Een frank",
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

func TestToWords_DutchBelgian_EdgeCases(t *testing.T) {
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
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Nul euro en een cent",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Elf",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Twaalf",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Eenentwintig",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Honderdeen",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Duizendeen",
		},
		{
			name:   "Belgian number style",
			amount: 83,
			options: &pkg.NumI18NOptions{
				Locale: "nl-BE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Drieentachtig",
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
