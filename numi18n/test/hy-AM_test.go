package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_Armenian_Numbers(t *testing.T) {
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
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Զրո",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Հինգ",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Տասնհինգ",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Երեսուն",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Քառասուն Յոթ",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Մեկ Հարյուր",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Երկու Հարյուր Հիսուն Վեց",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Մեկ Հազար",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Մեկ Միլիոն",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Մեկ Միլիոն Երկու Հարյուր Երեսուն Չորս Հազար Հինգ Հարյուր Վաթսուն Յոթ",
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

func TestToWords_Armenian_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One dram",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Մեկ Դրամ",
		},
		{
			name:   "Multiple drams",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Հինգ Դրամ",
		},
		{
			name:   "Zero drams",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Զրո Դրամ",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Մեկ Միլիոն Դրամ",
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

func TestToWords_Armenian_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Drams and one luma",
			amount: 5.01,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Հինգ Դրամ Եվ Մեկ Լումա",
		},
		{
			name:   "Drams and multiple lumas",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Հինգ Դրամ Եվ Քսան Հինգ Լումա",
		},
		{
			name:   "Only lumas",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Զրո Դրամ Եվ Իննսուն Ինը Լումա",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Մեկ Հազար Երկու Հարյուր Երեսուն Չորս Դրամ Եվ Հիսուն Վեց Լումա",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Մեկ Հարյուր Քսան Երեք Եվ Քառասուն Հինգ",
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

func TestToWords_Armenian_Negative(t *testing.T) {
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
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Մինուս Հիսուն",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Մինուս Քսան Հինգ Դրամ Եվ Յոթանասուն Հինգ Լումա",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &numi18n.OverrideOptions{
						NegativeWord: "Բացասական",
					},
				},
			},
			expected: "Բացասական Մեկ Հարյուր Դրամ",
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

func TestToWords_Armenian_Formatting(t *testing.T) {
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
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "ՄԵԿ ՀԱՐՅՈՒՐ ՔՍԱՆ ԵՐԵՔ ԴՐԱՄ",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "մեկ հարյուր քսան երեք դրամ",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Ինը Հարյուր Իննսուն Ինը Միայն",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Հինգ Հարյուր Դրամ Միայն",
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

func TestToWords_Armenian_CustomCurrency(t *testing.T) {
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
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "եվրո",
						Plural: "եվրո",
					},
				},
			},
			expected: "Մեկ Հարյուր եվրո",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "դոլար",
						Plural:           "դոլար",
						FractionUnitName: "ցենտ",
						FractionPlural:   "ցենտ",
					},
				},
			},
			expected: "Հիսուն դոլար Եվ Քսան Հինգ ցենտ",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "ֆունտ",
						Plural: "ֆունտ",
					},
				},
			},
			expected: "Մեկ ֆունտ",
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

func TestToWords_Armenian_EdgeCases(t *testing.T) {
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
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Զրո Դրամ Եվ Մեկ Լումա",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Տասնմեկ",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Տասներկու",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Քսան Մեկ",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Մեկ Հարյուր Մեկ",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &numi18n.NumI18NOptions{
				Locale: "hy-AM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Մեկ Հազար Մեկ",
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
