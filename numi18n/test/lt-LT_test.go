package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_LithuanianLithuania_Numbers(t *testing.T) {
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
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Nulis",
		},
		{
			name:   "One",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Vienas",
		},
		{
			name:   "Five",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Penki",
		},
		{
			name:   "Ten",
			amount: 10,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Dešimt",
		},
		{
			name:   "Eleven",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Vienuolika",
		},
		{
			name:   "Twenty",
			amount: 20,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Dvidešimt",
		},
		{
			name:   "Twenty-one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Dvidešimt vienas",
		},
		{
			name:   "Thirty-five",
			amount: 35,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Trisdešimt penki",
		},
		{
			name:   "Ninety-nine",
			amount: 99,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Devyniasdešimt devyni",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Vienas šimtas",
		},
		{
			name:   "One hundred and fifty",
			amount: 150,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Vienas šimtas penkiasdešimt",
		},
		{
			name:   "Five hundred",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Penki šimtai",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tūkstantis",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Milijonas",
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

func TestToWords_LithuanianLithuania_NegativeNumbers(t *testing.T) {
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
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "Minus vienas",
		},
		{
			name:   "Negative five",
			amount: -5,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "Minus penki",
		},
		{
			name:   "Negative twenty",
			amount: -20,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "Minus dvidešimt",
		},
		{
			name:   "Negative one hundred",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "Minus Vienas šimtas",
		},
		{
			name:   "Negative one thousand",
			amount: -1000,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "Minus tūkstantis",
		},
		{
			name:   "Negative one million",
			amount: -1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "Minus milijonas",
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

func TestToWords_LithuanianLithuania_Currency(t *testing.T) {
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
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Nulis Eurai",
		},
		{
			name:   "One euro",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Vienas Euras",
		},
		{
			name:   "Five euros",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Penki Eurai",
		},
		{
			name:   "Twenty euros",
			amount: 20,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Dvidešimt Eurai",
		},
		{
			name:   "One hundred euros",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Vienas šimtas Eurai",
		},
		{
			name:   "One thousand euros",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Tūkstantis Eurai",
		},
		{
			name:   "One million euros",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Currency:   true,
				},
			},
			expected: "Milijonas Eurai",
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

func TestToWords_LithuanianLithuania_DecimalCurrency(t *testing.T) {
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
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Decimal:    true,
					Currency:   true,
				},
			},
			expected: "Nulis Eurai ir penkiasdešimt Centai",
		},
		{
			name:   "One point twenty",
			amount: 1.20,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Decimal:    true,
					Currency:   true,
				},
			},
			expected: "Vienas Euras ir dvidešimt Centai",
		},
		{
			name:   "Five point seventy-five",
			amount: 5.75,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Decimal:    true,
					Currency:   true,
				},
			},
			expected: "Penki Eurai ir septyniasdešimt penki Centai",
		},
		{
			name:   "Twenty point fifty",
			amount: 20.50,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Decimal:    true,
					Currency:   true,
				},
			},
			expected: "Dvidešimt Eurai ir penkiasdešimt Centai",
		},
		{
			name:   "One hundred point ninety-nine",
			amount: 100.99,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Decimal:    true,
					Currency:   true,
				},
			},
			expected: "Vienas šimtas Eurai ir devyniasdešimt devyni Centai",
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

func TestToWords_LithuanianLithuania_NegativeCurrency(t *testing.T) {
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
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
					Currency:     true,
				},
			},
			expected: "Minus vienas Euras",
		},
		{
			name:   "Negative five euros",
			amount: -5,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
					Currency:     true,
				},
			},
			expected: "Minus penki Eurai",
		},
		{
			name:   "Negative twenty point five",
			amount: -20.50,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
					Decimal:      true,
					Currency:     true,
				},
			},
			expected: "Minus dvidešimt Eurai ir penkiasdešimt Centai",
		},
		{
			name:   "Negative one hundred euros",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
					Currency:     true,
				},
			},
			expected: "Minus Vienas šimtas Eurai",
		},
		{
			name:   "Negative one hundred point ninety-nine",
			amount: -100.99,
			options: &numi18n.NumI18NOptions{
				Locale: "lt-LT",
				WordDetails: &numi18n.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
					Decimal:      true,
					Currency:     true,
				},
			},
			expected: "Minus Vienas šimtas Eurai ir devyniasdešimt devyni Centai",
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
