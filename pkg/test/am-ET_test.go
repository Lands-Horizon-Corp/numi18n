package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_AmharicET_Numbers(t *testing.T) {
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
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ዜሮ",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "አምስት",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "አስራ አምስት",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ሰላሳ",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "አርባ ሰባት",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "መቶ",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ሁለት መቶ ሃምሳ ስድስት",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ሺህ",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ሚሊዮን",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "አንድ ሚሊዮን ሁለት መቶ ሰላሳ አራት ሺህ አምስት መቶ ስድስት ዐሥር ሰባት",
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

func TestToWords_AmharicET_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One birr",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "አንድ ብር",
		},
		{
			name:   "Multiple birr",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "አምስት ብር",
		},
		{
			name:   "Zero birr",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ዜሮ ብር",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ሚሊዮን ብር",
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

func TestToWords_AmharicET_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Birr and one mills",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "አምስት ብር እና አንድ ሚልስ",
		},
		{
			name:   "Birr and multiple mills",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "አምስት ብር እና ሃያ አምስት ሚልስ",
		},
		{
			name:   "Only mills",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ዜሮ ብር እና ዘጠና ዘጠኝ ሚልስ",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "አንድ ሺህ ሁለት መቶ ሰላሳ አራት ብር እና ሃምሳ ስድስት ሚልስ",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "አንድ መቶ ሃያ ሶስት እና አርባ አምስት",
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

func TestToWords_AmharicET_Negative(t *testing.T) {
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
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "አልፎ ሃምሳ",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "አልፎ ሃያ አምስት ብር እና ሰባ አምስት ሚልስ",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "ናይቭ",
					},
				},
			},
			expected: "ናይቭ መቶ ብር",
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

func TestToWords_AmharicET_Formatting(t *testing.T) {
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
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "አንድ መቶ ሃያ ሶስት ብር",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "አንድ መቶ ሃያ ሶስት ብር",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "ዘጠኝ መቶ ዘጠና ዘጠኝ ብቻ",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "አምስት መቶ ብር ብቻ",
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

func TestToWords_AmharicET_CustomCurrency(t *testing.T) {
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
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "ዶላር",
						Plural: "ዶላር",
					},
				},
			},
			expected: "መቶ ዶላር",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "ዩሮ",
						Plural:           "ዩሮ",
						FractionUnitName: "ሳንት",
						FractionPlural:   "ሳንት",
					},
				},
			},
			expected: "ሃምሳ ዩሮ እና ሃያ አምስት ሳንት",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "ፓውንድ",
						Plural: "ፓውንድ",
					},
				},
			},
			expected: "አንድ ፓውንድ",
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

func TestToWords_AmharicET_EdgeCases(t *testing.T) {
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
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ዜሮ ብር እና አንድ ሚልስ",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "አስራ አንድ",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "አስራ ሁለት",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ሃያ አንድ",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "አንድ መቶ አንድ",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "am-ET",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "አንድ ሺህ አንድ",
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
