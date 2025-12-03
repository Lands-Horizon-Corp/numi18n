package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_LuxembourgishLuxembourg_Numbers(t *testing.T) {
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
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Null",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Fënnef",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Fënnefzéng",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Drësseg",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Siwwenanvéierzeg",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Een Honnert",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Zwou honnert fofzeg sechs",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Dausend",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Millioun",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Een Millioun zwou honnert drësseg véier dausend fënnef honnert sechzeg siwen",
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

func TestToWords_LuxembourgishLuxembourg_Currency(t *testing.T) {
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
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Een Euro",
		},
		{
			name:   "Multiple euros",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Fënnef Euroen",
		},
		{
			name:   "Zero euros",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Null Euroen",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Millioun Euroen",
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

func TestToWords_LuxembourgishLuxembourg_Decimals(t *testing.T) {
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
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Fënnef Euroen an een Cent",
		},
		{
			name:   "Euros and multiple cents",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Fënnef Euroen an fënnefandzwanzeg Centen",
		},
		{
			name:   "Only cents",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Null Euroen an nénganóngzeg Centen",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Een dausend zwou honnert drësseg véier Euroen an sechsanfofzeg Centen",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Een honnert zwanzeg dräi an fënnefanvéierzeg",
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

func TestToWords_LuxembourgishLuxembourg_Negative(t *testing.T) {
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
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Minus fofzeg",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Minus fënnefandzwanzeg Euroen an fënnefanséwenzeg Centen",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &numi18n.OverrideOptions{
						NegativeWord: "Negativ",
					},
				},
			},
			expected: "Negativ Een Honnert Euroen",
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

func TestToWords_LuxembourgishLuxembourg_Formatting(t *testing.T) {
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
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "EEN HONNERT ZWANZEG DRÄI EUROEN",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "een honnert zwanzeg dräi euroen",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Néng honnert nongzeg néng nëmmen",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Fënnefhonnert Euroen nëmmen",
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

func TestToWords_LuxembourgishLuxembourg_CustomCurrency(t *testing.T) {
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
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "Dollar",
						Plural: "Dollaren",
					},
				},
			},
			expected: "Een Honnert Dollaren",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "Pond",
						Plural:           "Pond",
						FractionUnitName: "Penny",
						FractionPlural:   "Pence",
					},
				},
			},
			expected: "Fofzeg Pond an fënnefandzwanzeg Pence",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "Franc",
						Plural: "Francen",
					},
				},
			},
			expected: "Een Franc",
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

func TestToWords_LuxembourgishLuxembourg_EdgeCases(t *testing.T) {
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
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Null Euroen an een Cent",
		},
		{
			name:   "Eleven",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Eelef",
		},
		{
			name:   "Twelve",
			amount: 12,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Zwielef",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Eenandzwanzeg",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Een honnert een",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &numi18n.NumI18NOptions{
				Locale: "lb-LU",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Een dausend een",
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
