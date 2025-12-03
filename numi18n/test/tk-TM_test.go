package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_TurkmenTurkmenistan_Numbers(t *testing.T) {
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
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Nol",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Bäş",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "On bäş",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
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
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Kyrk ýedi",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Bir ýüz",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Iki bir ýüz elli alty",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Bir müň",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Bir million",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Bir bir million iki bir ýüz otuz dört bir müň bäş bir ýüz altmyş ýedi",
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

func TestToWords_TurkmenTurkmenistan_Currency(t *testing.T) {
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
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Bir Manat",
		},
		{
			name:   "Multiple manats",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Bäş Manat",
		},
		{
			name:   "Zero manats",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Nol Manat",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Bir million Manat",
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

func TestToWords_TurkmenTurkmenistan_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Manats and one tenge",
			amount: 5.01,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Bäş Manat we bir Tenge",
		},
		{
			name:   "Manats and multiple tenges",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Bäş Manat we ýigrimi bäş Tenge",
		},
		{
			name:   "Only tenges",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Nol Manat we togsan dokuz Tenge",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Bir bir müň iki bir ýüz otuz dört Manat we elli alty Tenge",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Bir bir ýüz ýigrimi üç we kyrk bäş",
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

func TestToWords_TurkmenTurkmenistan_Negative(t *testing.T) {
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
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Minus elli",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Minus ýigrimi bäş Manat we ýetmiş bäş Tenge",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &numi18n.OverrideOptions{
						NegativeWord: "Otrisatel",
					},
				},
			},
			expected: "Otrisatel Bir ýüz Manat",
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

func TestToWords_TurkmenTurkmenistan_Formatting(t *testing.T) {
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
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "BIR BIR ÝÜZ ÝIGRIMI ÜÇ MANAT",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "bir bir ýüz ýigrimi üç manat",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Dokuz bir ýüz togsan dokuz diňe",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Bäş ýüz Manat diňe",
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

func TestToWords_TurkmenTurkmenistan_CustomCurrency(t *testing.T) {
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
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "som",
						Plural: "somlar",
					},
				},
			},
			expected: "Bir ýüz somlar",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "dollar",
						Plural:           "dollarlar",
						FractionUnitName: "sent",
						FractionPlural:   "sentler",
					},
				},
			},
			expected: "Elli dollarlar we ýigrimi bäş sentler",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "funt",
						Plural: "funtlar",
					},
				},
			},
			expected: "Bir funt",
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

func TestToWords_TurkmenTurkmenistan_EdgeCases(t *testing.T) {
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
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Nol Manat we bir Tenge",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
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
				Locale: "tk-TM",
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
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ýigrimi bir",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Bir bir ýüz bir",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &numi18n.NumI18NOptions{
				Locale: "tk-TM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Bir bir müň bir",
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
