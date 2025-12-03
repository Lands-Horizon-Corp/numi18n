package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_Dhivehi_Numbers(t *testing.T) {
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
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ސިރޯ",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ފިފް",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ފިފްޓީން",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ޓްރިސް",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ފިސްރަ ސެވް",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ސްތޯ",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ޓްނޯ ސްތޯ ފިވެސް ސެސް",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ތިސަން",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "މިލިއަން",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "އަން މިލިއަން ޓްނޯ ސްތޯ ޓްރިސް ފޯރ ތިސަން ފިފް ސްތޯ ސެކްސެތް ސެވް",
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

func TestToWords_Dhivehi_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One Rufiyaa",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "އަން Rufiyaa",
		},
		{
			name:   "Multiple Rufiyaa",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ފިފް Rufiyaa",
		},
		{
			name:   "Zero Rufiyaa",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ސިރޯ Rufiyaa",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "މިލިއަން Rufiyaa",
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

func TestToWords_Dhivehi_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Rufiyaa and one Laari",
			amount: 5.01,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ފިފް Rufiyaa އަންޑް އަން Laari",
		},
		{
			name:   "Rufiyaa and multiple Laari",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ފިފް Rufiyaa އަންޑް ޓްވިން ފިފް Laari",
		},
		{
			name:   "Only Laari",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ސިރޯ Rufiyaa އަންޑް ނަވެސް ނަވް Laari",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "އަން ތިސަން ޓްނޯ ސްތޯ ޓްރިސް ފޯރ Rufiyaa އަންޑް ފިވެސް ސެސް Laari",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "އަން ސްތޯ ޓްވިން ތްރީ އަންޑް ފިސްރަ ފިފް",
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

func TestToWords_Dhivehi_Negative(t *testing.T) {
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
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "މިނަސް ފިވެސް",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "މިނަސް ޓްވިން ފިފް Rufiyaa އަންޑް ސެބެތް ފިފް Laari",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &numi18n.OverrideOptions{
						NegativeWord: "ނެގެޓިވް",
					},
				},
			},
			expected: "ނެގެޓިވް ސްތޯ Rufiyaa",
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

func TestToWords_Dhivehi_Formatting(t *testing.T) {
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
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "އަން ސްތޯ ޓްވިން ތްރީ RUFIYAA",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "އަން ސްތޯ ޓްވިން ތްރީ rufiyaa",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "ނަވް ސްތޯ ނަވެސް ނަވް ނުވަތަ",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "ފިފް ސްތޯ Rufiyaa ނުވަތަ",
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

func TestToWords_Dhivehi_CustomCurrency(t *testing.T) {
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
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "Mark",
						Plural: "Mark",
					},
				},
			},
			expected: "ސްތޯ Mark",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "Dollar",
						Plural:           "Dollar",
						FractionUnitName: "Cent",
						FractionPlural:   "Cent",
					},
				},
			},
			expected: "ފިވެސް Dollar އަންޑް ޓްވިން ފިފް Cent",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "Pfund",
						Plural: "Pfund",
					},
				},
			},
			expected: "އަން Pfund",
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

func TestToWords_Dhivehi_EdgeCases(t *testing.T) {
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
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ސިރޯ Rufiyaa އަންޑް އަން Laari",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "އެލްވް",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ޓްވޮލްވް",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ޓްވިން އަން",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "އަން ސްތޯ އަން",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &numi18n.NumI18NOptions{
				Locale: "dv-MV",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "އަން ތިސަން އަން",
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
