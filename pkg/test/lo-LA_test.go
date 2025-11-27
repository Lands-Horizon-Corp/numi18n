package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_LaoLaos_Numbers(t *testing.T) {
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
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ສູນ",
		},
		{
			name:   "One",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ຫນຶ່ງ",
		},
		{
			name:   "Five",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ຫ້າ",
		},
		{
			name:   "Ten",
			amount: 10,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ສິບ",
		},
		{
			name:   "Eleven",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ສິບເອັດ",
		},
		{
			name:   "Twenty",
			amount: 20,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ຊາວ",
		},
		{
			name:   "Twenty-one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ຊາວເອັດ",
		},
		{
			name:   "Thirty-five",
			amount: 35,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ສາມສິບຫ້າ",
		},
		{
			name:   "Ninety-nine",
			amount: 99,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ເກົ້າສິບເກົ້າ",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ນຶ່ງຮ້ອຍ",
		},
		{
			name:   "One hundred and fifty",
			amount: 150,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ຫນຶ່ງ ຮ້ອຍ ຫ້າສິບ",
		},
		{
			name:   "Five hundred",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ຫ້າຮ້ອຍ",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ພັນ",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ລ້ານ",
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

func TestToWords_LaoLaos_NegativeNumbers(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Negative one",
			amount: -1,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "ລົບ ຫນຶ່ງ",
		},
		{
			name:   "Negative five",
			amount: -5,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "ລົບ ຫ້າ",
		},
		{
			name:   "Negative twenty",
			amount: -20,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "ລົບ ຊາວ",
		},
		{
			name:   "Negative one hundred",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "ລົບ ນຶ່ງຮ້ອຍ",
		},
		{
			name:   "Negative one thousand",
			amount: -1000,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "ລົບ ພັນ",
		},
		{
			name:   "Negative one million",
			amount: -1000000,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Capitalize:   true,
					NegativeWord: true,
				},
			},
			expected: "ລົບ ລ້ານ",
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

func TestToWords_LaoLaos_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Zero Kip",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ສູນ Kip",
		},
		{
			name:   "One Kip",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ຫນຶ່ງ Kip",
		},
		{
			name:   "Five Kip",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ຫ້າ Kip",
		},
		{
			name:   "Twenty Kip",
			amount: 20,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ຊາວ Kip",
		},
		{
			name:   "One hundred Kip",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ນຶ່ງຮ້ອຍ Kip",
		},
		{
			name:   "One thousand Kip",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ພັນ Kip",
		},
		{
			name:   "One million Kip",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ລ້ານ Kip",
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

func TestToWords_LaoLaos_DecimalCurrency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Zero point five",
			amount: 0.5,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ສູນ Kip ແລະ ຫ້າສິບ Att",
		},
		{
			name:   "One point twenty",
			amount: 1.20,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ຫນຶ່ງ Kip ແລະ ຊາວ Att",
		},
		{
			name:   "Five point seventy-five",
			amount: 5.75,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ຫ້າ Kip ແລະ ເຈັດສິບຫ້າ Att",
		},
		{
			name:   "Twenty point fifty",
			amount: 20.50,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ຊາວ Kip ແລະ ຫ້າສິບ Att",
		},
		{
			name:   "One hundred point ninety-nine",
			amount: 100.99,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ນຶ່ງຮ້ອຍ Kip ແລະ ເກົ້າສິບເກົ້າ Att",
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

func TestToWords_LaoLaos_NegativeCurrency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Negative one Kip",
			amount: -1,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ລົບ ຫນຶ່ງ Kip",
		},
		{
			name:   "Negative five Kip",
			amount: -5,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ລົບ ຫ້າ Kip",
		},
		{
			name:   "Negative twenty point five",
			amount: -20.50,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ລົບ ຊາວ Kip ແລະ ຫ້າສິບ Att",
		},
		{
			name:   "Negative one hundred Kip",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ລົບ ນຶ່ງຮ້ອຍ Kip",
		},
		{
			name:   "Negative one hundred point ninety-nine",
			amount: -100.99,
			options: &pkg.NumI18NOptions{
				Locale: "lo-LA",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ລົບ ນຶ່ງຮ້ອຍ Kip ແລະ ເກົ້າສິບເກົ້າ Att",
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
