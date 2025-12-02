package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_SlovakSlovakia_Numbers(t *testing.T) {
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
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Nula",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Päť",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Pätnásť",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tridsať",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Štyridsať sedem",
		},
		{
			name:   "One hundred",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Sto",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Dva sto päťdesiat šesť",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tisíc",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Milión",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Jeden milión dva sto tridsať štyri tisíc päť sto šesťdesiat sedem",
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

func TestToWords_SlovakSlovakia_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One euro",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Jeden Euro",
		},
		{
			name:   "Multiple euros",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Päť Eurá",
		},
		{
			name:   "Zero euros",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Nula Eurá",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Milión Eurá",
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

func TestToWords_SlovakSlovakia_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Euros and one cent",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Päť Eurá a jeden Cent",
		},
		{
			name:   "Euros and multiple cents",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Päť Eurá a dvadsať päť Centy",
		},
		{
			name:   "Only cents",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Nula Eurá a deväťdesiat deväť Centy",,
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Jeden tisíc dva sto tridsať štyri Eurá a päťdesiat šesť Centy",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Jeden sto dvadsať tri a štyridsať päť",,
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

func TestToWords_SlovakSlovakia_Negative(t *testing.T) {
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
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Mínus päťdesiat",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Mínus dvadsať päť Eurá a sedemdesiat päť Centy",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "Záporné",
					},
				},
			},
			expected: "Záporné Sto Eurá",
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

func TestToWords_SlovakSlovakia_Formatting(t *testing.T) {
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
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "JEDEN STO DVADSAŤ TRI EURÁ",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "jeden sto dvadsať tri eurá",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Deväť Sto Deväťdesiat Deväť Len",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Päťsto Eurá Len",
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

func TestToWords_SlovakSlovakia_CustomCurrency(t *testing.T) {
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
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "koruna",
						Plural: "koruny",
					},
				},
			},
			expected: "Sto Koruny",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "dolár",
						Plural:           "dolárov",
						FractionUnitName: "cent",
						FractionPlural:   "centov",
					},
				},
			},
			expected: "Päťdesiat dolárov a dvadsať päť centov",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "libra",
						Plural: "libier",
					},
				},
			},
			expected: "Jeden Libra",
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

func TestToWords_SlovakSlovakia_EdgeCases(t *testing.T) {
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
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Nula Eurá a jeden Cent",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Jedenásť",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Dvanásť",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Dvadsať Jeden",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Jeden Sto Jeden",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "sk-SK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Jeden Tisíc Jeden",
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
