package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_TswanaBotswana_Numbers(t *testing.T) {
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
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Lefela",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tlhano",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Lesome le tlhano",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Masome a tharo",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Masome a ne le supa",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Lekgolo",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Pedi lekgolo masome a tlhano thataro",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Sekete",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Miliyone",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Nngwe miliyone pedi lekgolo masome a tharo ne sekete tlhano lekgolo masome a thataro le masome supa",
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

func TestToWords_TswanaBotswana_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One pula",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Nngwe Pula",
		},
		{
			name:   "Multiple pulas",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Tlhano Dipula",
		},
		{
			name:   "Zero pulas",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Lefela Dipula",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Miliyone Dipula",
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

func TestToWords_TswanaBotswana_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Pulas and one thebe",
			amount: 5.01,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Tlhano Dipula le nngwe Thebe",
		},
		{
			name:   "Pulas and multiple thebes",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Tlhano Dipula le masome a mabedi le tlhano Dithebe",
		},
		{
			name:   "Only thebes",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Lefela Dipula le masome a robongwe le robongwe Dithebe",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Nngwe sekete pedi lekgolo masome a tharo ne Dipula le masome a tlhano le thataro Dithebe",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Nngwe lekgolo masome a mabedi tharo le masome a ne le tlhano",
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

func TestToWords_TswanaBotswana_Negative(t *testing.T) {
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
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Tlosa masome a tlhano",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Tlosa masome a mabedi le tlhano Dipula le masome a supa le tlhano Dithebe",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &numi18n.OverrideOptions{
						NegativeWord: "Kgaratlhang",
					},
				},
			},
			expected: "Kgaratlhang Lekgolo Dipula",
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

func TestToWords_TswanaBotswana_Formatting(t *testing.T) {
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
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "NNGWE LEKGOLO MASOME A MABEDI THARO DIPULA",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "nngwe lekgolo masome a mabedi tharo dipula",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Robongwe lekgolo masome a robongwe robongwe fela",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Makgolo a tlhano Dipula fela",
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

func TestToWords_TswanaBotswana_CustomCurrency(t *testing.T) {
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
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "rand",
						Plural: "dirands",
					},
				},
			},
			expected: "Lekgolo dirands",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "dollar",
						Plural:           "didollar",
						FractionUnitName: "cent",
						FractionPlural:   "dicent",
					},
				},
			},
			expected: "Masome a tlhano didollar le masome a mabedi le tlhano dicent",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "pound",
						Plural: "dipounds",
					},
				},
			},
			expected: "Nngwe pound",
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

func TestToWords_TswanaBotswana_EdgeCases(t *testing.T) {
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
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Lefela Dipula le nngwe Thebe",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Lesome le nngwe",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Lesome le pedi",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Masome a mabedi le nngwe",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Nngwe lekgolo nngwe",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &numi18n.NumI18NOptions{
				Locale: "tn-BW",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Nngwe sekete nngwe",
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
