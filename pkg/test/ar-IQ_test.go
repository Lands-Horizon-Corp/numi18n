package pkg

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_ArabicIQ_Numbers(t *testing.T) {
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
				Locale: "ar-IQ",
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
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "خمسة",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "خمسة عشر",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ثلاثون",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "أربعون سبعة",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "مئة",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "اثنان مئة خمسون ستة",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ألف",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "مليون",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "واحد مليون اثنان مئة ثلاثون أربعة ألف خمسة مئة ستون سبعة",
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

func TestToWords_ArabicIQ_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One dinar",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "واحد دينار",
		},
		{
			name:   "Multiple dinars",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "خمسة دنانير",
		},
		{
			name:   "Zero dinars",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "صفر دنانير",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "مليون دنانير",
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

func TestToWords_ArabicIQ_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Dinars and one fils",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "خمسة دنانير و واحد فلس",
		},
		{
			name:   "Dinars and multiple fils",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "خمسة دنانير و عشرون خمسة فلوس",
		},
		{
			name:   "Only fils",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "صفر دنانير و تسعون تسعة فلوس",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "واحد ألف اثنان مئة ثلاثون أربعة دنانير و خمسون ستة فلوس",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "واحد مئة عشرون ثلاثة و أربعون خمسة",
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

func TestToWords_ArabicIQ_Negative(t *testing.T) {
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
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ناقص خمسون",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ناقص عشرون خمسة دنانير و سبعون خمسة فلوس",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "سالب",
					},
				},
			},
			expected: "سالب مئة دنانير",
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

func TestToWords_ArabicIQ_Formatting(t *testing.T) {
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
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "واحد مئة عشرون ثلاثة دنانير",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "واحد مئة عشرون ثلاثة دنانير",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "تسعة مئة تسعون تسعة فقط",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "خمسة مئة دنانير فقط",
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

func TestToWords_ArabicIQ_CustomCurrency(t *testing.T) {
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
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "ريال",
						Plural: "ريالات",
					},
				},
			},
			expected: "مئة ريالات",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "درهم",
						Plural:           "دراهم",
						FractionUnitName: "قرش",
						FractionPlural:   "قروش",
					},
				},
			},
			expected: "خمسون دراهم و عشرون خمسة قروش",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "جنيه",
						Plural: "جنيهات",
					},
				},
			},
			expected: "واحد جنيه",
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

func TestToWords_ArabicIQ_EdgeCases(t *testing.T) {
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
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "صفر دنانير و واحد فلس",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "أحد عشر",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "اثنا عشر",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "عشرون واحد",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "واحد مئة واحد",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "ar-IQ",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "واحد ألف واحد",
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
