package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_Arabic001_Numbers(t *testing.T) {
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
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "صفر",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "خمسة",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "خمسة عشر",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ثلاثون",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "أربعون سبعة",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "مئة",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "اثنان مئة خمسون ستة",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ألف",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "مليون",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
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

func TestToWords_Arabic001_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One dollar",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "واحد دولار",
		},
		{
			name:   "Multiple dollars",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "خمسة دولارات",
		},
		{
			name:   "Zero dollars",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "صفر دولارات",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "مليون دولارات",
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

func TestToWords_Arabic001_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Dollars and one cent",
			amount: 5.01,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "خمسة دولارات و واحد سنت",
		},
		{
			name:   "Dollars and multiple cents",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "خمسة دولارات و عشرون خمسة سنتات",
		},
		{
			name:   "Only cents",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "صفر دولارات و تسعون تسعة سنتات",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "واحد ألف اثنان مئة ثلاثون أربعة دولارات و خمسون ستة سنتات",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
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

func TestToWords_Arabic001_Negative(t *testing.T) {
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
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ناقص خمسون",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ناقص عشرون خمسة دولارات و سبعون خمسة سنتات",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &numi18n.OverrideOptions{
						NegativeWord: "سالب",
					},
				},
			},
			expected: "سالب مئة دولارات",
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

func TestToWords_Arabic001_Formatting(t *testing.T) {
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
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "واحد مئة عشرون ثلاثة دولارات",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "واحد مئة عشرون ثلاثة دولارات",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "تسعة مئة تسعون تسعة فقط",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "خمسة مئة دولارات فقط",
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

func TestToWords_Arabic001_CustomCurrency(t *testing.T) {
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
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
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
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "يورو",
						Plural:           "يوروات",
						FractionUnitName: "سنت",
						FractionPlural:   "سنتات",
					},
				},
			},
			expected: "خمسون يوروات و عشرون خمسة سنتات",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
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

func TestToWords_Arabic001_EdgeCases(t *testing.T) {
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
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "صفر دولارات و واحد سنت",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "أحد عشر",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "اثنا عشر",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "عشرون واحد",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "واحد مئة واحد",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-001",
				WordDetails: &numi18n.WordDetails{
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
