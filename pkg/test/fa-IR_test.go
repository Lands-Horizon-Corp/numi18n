package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_PersianIran_Numbers(t *testing.T) {
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
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "صفر",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "پنج",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "پانزده",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "سی",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "چهل هفت",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "یک صد",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "دو صد پنجاه شش",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "یک هزار",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "یک میلیون",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "یک میلیون دو صد سی چهار هزار پنج صد شصت هفت",
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

func TestToWords_PersianIran_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One rial",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "یک ریال",
		},
		{
			name:   "Multiple rials",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "پنج ریال",
		},
		{
			name:   "Zero rials",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "صفر ریال",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "یک میلیون ریال",
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

func TestToWords_PersianIran_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Rials and one dinar",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "پنج ریال و یک دینار",
		},
		{
			name:   "Rials and multiple dinars",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "پنج ریال و بیست پنج دینار",
		},
		{
			name:   "Only dinars",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "صفر ریال و نود نه دینار",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "یک هزار دو صد سی چهار ریال و پنجاه شش دینار",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "یک صد بیست سه و چهل پنج",
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

func TestToWords_PersianIran_Negative(t *testing.T) {
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
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "منفی پنجاه",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "منفی بیست پنج ریال و هفتاد پنج دینار",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "منها",
					},
				},
			},
			expected: "منها یک صد ریال",
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

func TestToWords_PersianIran_Formatting(t *testing.T) {
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
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "یک صد بیست سه ریال",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "یک صد بیست سه ریال",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "نه صد نود نه فقط",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "پنج صد ریال فقط",
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

func TestToWords_PersianIran_CustomCurrency(t *testing.T) {
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
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "تومان",
						Plural: "تومان",
					},
				},
			},
			expected: "یک صد تومان",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "یورو",
						Plural:           "یورو",
						FractionUnitName: "سنت",
						FractionPlural:   "سنت",
					},
				},
			},
			expected: "پنجاه یورو و بیست پنج سنت",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "درهم",
						Plural: "درهم",
					},
				},
			},
			expected: "یک درهم",
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

func TestToWords_PersianIran_EdgeCases(t *testing.T) {
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
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "صفر ریال و یک دینار",
		},
		{
			name:   "Eleven",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "یازده",
		},
		{
			name:   "Twelve",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "دوازده",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "بیست یک",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "یک صد یک",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "fa-IR",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "یک هزار یک",
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
