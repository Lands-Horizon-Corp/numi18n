package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_LatvianLatvia_Numbers(t *testing.T) {
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
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Nulle",
		},
		{
			name:   "One",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Viens",
		},
		{
			name:   "Five",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Pieci",
		},
		{
			name:   "Ten",
			amount: 10,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Desmit",
		},
		{
			name:   "Eleven",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Vienpadsmit",
		},
		{
			name:   "Twenty",
			amount: 20,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Divdesmit",
		},
		{
			name:   "Twenty-one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Divdesmit viens",
		},
		{
			name:   "Thirty-five",
			amount: 35,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Trīsdesmit pieci",
		},
		{
			name:   "Ninety-nine",
			amount: 99,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Deviņdesmit deviņi",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Viens simts",
		},
		{
			name:   "One hundred and fifty",
			amount: 150,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Viens simts piecdesmit",
		},
		{
			name:   "Five hundred",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Pieci simti",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tūkstotis",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Miljons",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToWords(tt.amount)
			if result != tt.expected {
				t.Errorf("ToWords(%f) = %q, want %q", tt.amount, result, tt.expected)
			}
		})
	}
}

func TestToWords_LatvianLatvia_NegativeNumbers(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Negative one",
			amount: -1,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "Mīnus viens",
		},
		{
			name:   "Negative five",
			amount: -5,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "Mīnus pieci",
		},
		{
			name:   "Negative twenty",
			amount: -20,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "Mīnus divdesmit",
		},
		{
			name:   "Negative one hundred",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "Mīnus Viens simts",
		},
		{
			name:   "Negative one thousand",
			amount: -1000,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "Mīnus tūkstotis",
		},
		{
			name:   "Negative one million",
			amount: -1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "Mīnus miljons",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToWords(tt.amount)
			if result != tt.expected {
				t.Errorf("ToWords(%f) = %q, want %q", tt.amount, result, tt.expected)
			}
		})
	}
}

func TestToWords_LatvianLatvia_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Zero euros",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Nulle Eiro",
		},
		{
			name:   "One euro",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Viens Eiro",
		},
		{
			name:   "Five euros",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Pieci Eiro",
		},
		{
			name:   "Twenty euros",
			amount: 20,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Divdesmit Eiro",
		},
		{
			name:   "One hundred euros",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Viens simts Eiro",
		},
		{
			name:   "One thousand euros",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Tūkstotis Eiro",
		},
		{
			name:   "One million euros",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Miljons Eiro",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToWords(tt.amount)
			if result != tt.expected {
				t.Errorf("ToWords(%f) = %q, want %q", tt.amount, result, tt.expected)
			}
		})
	}
}

func TestToWords_LatvianLatvia_DecimalCurrency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Zero point five",
			amount: 0.5,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Decimal:    true,
					Currency:   true,
				},
			},
			expected: "Nulle Eiro un piecdesmit Centi",
		},
		{
			name:   "One point twenty",
			amount: 1.20,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Decimal:    true,
					Currency:   true,
				},
			},
			expected: "Viens Eiro un divdesmit Centi",
		},
		{
			name:   "Five point seventy-five",
			amount: 5.75,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Decimal:    true,
					Currency:   true,
				},
			},
			expected: "Pieci Eiro un septiņdesmit pieci Centi",
		},
		{
			name:   "Twenty point fifty",
			amount: 20.50,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Decimal:    true,
					Currency:   true,
				},
			},
			expected: "Divdesmit Eiro un piecdesmit Centi",
		},
		{
			name:   "One hundred point ninety-nine",
			amount: 100.99,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Decimal:    true,
					Currency:   true,
				},
			},
			expected: "Viens simts Eiro un deviņdesmit deviņi Centi",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToWords(tt.amount)
			if result != tt.expected {
				t.Errorf("ToWords(%f) = %q, want %q", tt.amount, result, tt.expected)
			}
		})
	}
}

func TestToWords_LatvianLatvia_NegativeCurrency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Negative one euro",
			amount: -1,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
					Currency:     true,
				},
			},
			expected: "Mīnus viens Eiro",
		},
		{
			name:   "Negative five euros",
			amount: -5,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
					Currency:     true,
				},
			},
			expected: "Mīnus pieci Eiro",
		},
		{
			name:   "Negative twenty point five",
			amount: -20.50,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
					Decimal:      true,
					Currency:     true,
				},
			},
			expected: "Mīnus divdesmit Eiro un piecdesmit Centi",
		},
		{
			name:   "Negative one hundred euros",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
					Currency:     true,
				},
			},
			expected: "Mīnus Viens simts Eiro",
		},
		{
			name:   "Negative one hundred point ninety-nine",
			amount: -100.99,
			options: &numi18n.NumI18NOptions{
				Locale: "lv-LV",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
					Decimal:      true,
					Currency:     true,
				},
			},
			expected: "Mīnus Viens simts Eiro un deviņdesmit deviņi Centi",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.options.ToWords(tt.amount)
			if result != tt.expected {
				t.Errorf("ToWords(%f) = %q, want %q", tt.amount, result, tt.expected)
			}
		})
	}
}
