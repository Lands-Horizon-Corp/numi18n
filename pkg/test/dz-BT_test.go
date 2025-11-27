package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_Dzongkha_Numbers(t *testing.T) {
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
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "སོ་རོ་",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ལྔ་",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ལྔ་གཉིས་",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "སུམ་བརྒྱ་",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "བཞི་བརྒྱ་ བདུན་",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "སུང་",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "གཉིས་ སུང་ ལྔ་བརྒྱ་ དྲུག་",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "མི་ལ་",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "མི་ལི་ཨན་",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "གཅིག་ མི་ལི་ཨན་ གཉིས་ སུང་ སུམ་བརྒྱ་ བཞི་ མི་ལ་ ལྔ་ སུང་ དགུ་བརྒྱ་ བདུན་",
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

func TestToWords_Dzongkha_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One Ngultrum",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "གཅིག་ Ngultrum",
		},
		{
			name:   "Multiple Ngultrum",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ལྔ་ Ngultrum",
		},
		{
			name:   "Zero Ngultrum",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "སོ་རོ་ Ngultrum",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "མི་ལི་ཨན་ Ngultrum",
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

func TestToWords_Dzongkha_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Ngultrum and one Chhertum",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ལྔ་ Ngultrum དང་ གཅིག་ Chhertum",
		},
		{
			name:   "Ngultrum and multiple Chhertum",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ལྔ་ Ngultrum དང་ ཉེ་བརྒྱ་ ལྔ་ Chhertum",
		},
		{
			name:   "Only Chhertum",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "སོ་རོ་ Ngultrum དང་ དང་ལྔ་དུས་དེབ་ དགུ་ Chhertum",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "གཅིག་ མི་ལ་ གཉིས་ སུང་ སུམ་བརྒྱ་ བཞི་ Ngultrum དང་ ལྔ་བརྒྱ་ དྲུག་ Chhertum",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "གཅིག་ སུང་ ཉེ་བརྒྱ་ སུམ་ དང་ བཞི་བརྒྱ་ ལྔ་",
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

func TestToWords_Dzongkha_Negative(t *testing.T) {
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
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "མིང་མེད་ ལྔ་བརྒྱ་",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "མིང་མེད་ ཉེ་བརྒྱ་ ལྔ་ Ngultrum དང་ བདུན་བརྒྱ་ ལྔ་ Chhertum",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "ཁ་དོ་",
					},
				},
			},
			expected: "ཁ་དོ་ སུང་ Ngultrum",
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

func TestToWords_Dzongkha_Formatting(t *testing.T) {
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
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "གཅིག་ སུང་ ཉེ་བརྒྱ་ སུམ་ NGULTRUM",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "གཅིག་ སུང་ ཉེ་བརྒྱ་ སུམ་ ngultrum",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "དགུ་ སུང་ དང་ལྔ་དུས་དེབ་ དགུ་ ཡང་ན་",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "ལྔ་ སུང་ Ngultrum ཡང་ན་",
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

func TestToWords_Dzongkha_CustomCurrency(t *testing.T) {
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
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "རིལ་མོ་",
						Plural: "རིལ་མོ་",
					},
				},
			},
			expected: "སུང་ རིལ་མོ་",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "དངུལ་",
						Plural:           "དངུལ་",
						FractionUnitName: "ཚུགས་",
						FractionPlural:   "ཚུགས་",
					},
				},
			},
			expected: "ལྔ་བརྒྱ་ དངུལ་ དང་ ཉེ་བརྒྱ་ ལྔ་ ཚུགས་",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "སྐར་མ་",
						Plural: "སྐར་མ་",
					},
				},
			},
			expected: "གཅིག་ སྐར་མ་",
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

func TestToWords_Dzongkha_EdgeCases(t *testing.T) {
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
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "སོ་རོ་ Ngultrum དང་ གཅིག་ Chhertum",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "གཅིག་གཅིག་",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ཉེ་གཉིས་",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ཉེ་བརྒྱ་ གཅིག་",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "གཅིག་ སུང་ གཅིག་",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "dz-BT",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "གཅིག་ མི་ལ་ གཅིག་",
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
