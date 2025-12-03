package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_XhosaZA_Numbers(t *testing.T) {
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
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Iqanda",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Isihlanu",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ishumi elinesihlanu",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Amashumi amathathu",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Amashumi amane nesixhenxe",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ikhulu",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Isibini ikhulu amashumi amahlanu isithandathu",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Iwaka",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Isigidi sexhenxe",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Inye isigidi sexhenxe isibini ikhulu amashumi amathathu isine iwaka isihlanu ikhulu amashumi amathandathu isixhenxe",
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

func TestToWords_XhosaZA_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One rand",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Inye iRandi",
		},
		{
			name:   "Multiple rands",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Isihlanu iiRandi",
		},
		{
			name:   "Zero rands",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Iqanda iiRandi",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Isigidi sexhenxe iiRandi",
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

func TestToWords_XhosaZA_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Rands and one cent",
			amount: 5.01,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Isihlanu iiRandi no inye icenti",
		},
		{
			name:   "Rands and multiple cents",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Isihlanu iiRandi no amashumi amabini nesihlanu iicenti",
		},
		{
			name:   "Only cents",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Iqanda iiRandi no amashumi asithoba nesithoba iicenti",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Inye iwaka isibini ikhulu amashumi amathathu isine iiRandi no amashumi amahlanu nesithandathu iicenti",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Inye ikhulu amashumi amabini isithathu no amashumi amane nesihlanu",
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

func TestToWords_XhosaZA_Negative(t *testing.T) {
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
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Minus amashumi amahlanu",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Minus amashumi amabini nesihlanu iiRandi no amashumi asixhenxe nesihlanu iicenti",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &numi18n.OverrideOptions{
						NegativeWord: "ayikho",
					},
				},
			},
			expected: "Ayikho Ikhulu iiRandi",
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

func TestToWords_XhosaZA_Formatting(t *testing.T) {
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
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "INYE IKHULU AMASHUMI AMABINI ISITHATHU IIRANDI",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "inye ikhulu amashumi amabini isithathu iirandi",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Ithoba ikhulu amashumi asithoba ithoba kuphela",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Amakhulu amahlanu iiRandi kuphela",
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

func TestToWords_XhosaZA_CustomCurrency(t *testing.T) {
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
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "ipula",
						Plural: "iipula",
					},
				},
			},
			expected: "Ikhulu iipula",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "iyuro",
						Plural:           "iiyuro",
						FractionUnitName: "icenti",
						FractionPlural:   "iicenti",
					},
				},
			},
			expected: "Amashumi amahlanu iiyuro no amashumi amabini nesihlanu iicenti",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "ipawundi",
						Plural: "iipawundi",
					},
				},
			},
			expected: "Inye ipawundi",
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

func TestToWords_XhosaZA_EdgeCases(t *testing.T) {
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
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Iqanda iiRandi no inye icenti",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ishumi elinanye",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ishumi elinesibini",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Amashumi amabini nesinye",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Inye ikhulu inye",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &numi18n.NumI18NOptions{
				Locale: "xh-ZA",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Inye iwaka inye",
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
