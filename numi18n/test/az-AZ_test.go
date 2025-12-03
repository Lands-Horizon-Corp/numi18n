package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_AzerbaijaniAzerbaijan_Numbers(t *testing.T) {
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
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Sıfır",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Beş",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "On beş",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Otuz",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Qırx Yeddi",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Bir Yüz",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "İki Yüz Əlli Altı",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Min",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Milyon",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Bir Milyon İki Yüz Otuz Dörd Min Beş Yüz Altmış Yeddi",
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

func TestToWords_AzerbaijaniAzerbaijan_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One manat",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Bir Manat",
		},
		{
			name:   "Multiple manat",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Beş Manat",
		},
		{
			name:   "Zero manat",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Sıfır Manat",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Milyon Manat",
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

func TestToWords_AzerbaijaniAzerbaijan_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Manat and one qəpik",
			amount: 5.01,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Beş Manat Və Bir Qəpik",
		},
		{
			name:   "Manat and multiple qəpik",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Beş Manat Və İyirmi Beş Qəpiklər",
		},
		{
			name:   "Only qəpik",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Sıfır Manat Və Doxsan Doqquz Qəpiklər",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Bir Min İki Yüz Otuz Dörd Manat Və Əlli Altı Qəpiklər",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Bir Yüz İyirmi Üç Və Qırx Beş",
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

func TestToWords_AzerbaijaniAzerbaijan_Negative(t *testing.T) {
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
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Mənfi Əlli",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Mənfi İyirmi Beş Manat Və Yetmiş Beş Qəpiklər",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &numi18n.OverrideOptions{
						NegativeWord: "Neqativ",
					},
				},
			},
			expected: "Neqativ Bir Yüz Manat",
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

func TestToWords_AzerbaijaniAzerbaijan_Formatting(t *testing.T) {
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
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "BIR YÜZ İYIRMI ÜÇ MANAT",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "bir yüz iyirmi üç manat",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Doqquz Yüz Doxsan Doqquz Yalnız",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Beş Yüz Manat Yalnız",
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

func TestToWords_AzerbaijaniAzerbaijan_CustomCurrency(t *testing.T) {
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
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "lirə",
						Plural: "lirə",
					},
				},
			},
			expected: "Bir Yüz lirə",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "dollar",
						Plural:           "dollar",
						FractionUnitName: "sent",
						FractionPlural:   "sent",
					},
				},
			},
			expected: "Əlli dollar Və İyirmi Beş sent",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "avro",
						Plural: "avro",
					},
				},
			},
			expected: "Bir avro",
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

func TestToWords_AzerbaijaniAzerbaijan_EdgeCases(t *testing.T) {
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
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Sıfır Manat Və Bir Qəpik",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "On bir",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "On iki",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "İyirmi Bir",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Bir Yüz Bir",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &numi18n.NumI18NOptions{
				Locale: "az-AZ",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Bir Min Bir",
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
