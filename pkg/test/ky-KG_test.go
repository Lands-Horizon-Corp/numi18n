package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_KyrgyzKyrgyzstan_Numbers(t *testing.T) {
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
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Нөл",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Беш",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Он беш",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Отуз",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Кырк Жети",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Бир жүз",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Эки Бир жүз Элүү Алты",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Бир миң",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Бир миллион",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Бир Бир миллион Эки Бир жүз Отуз Төрт Бир миң Беш Бир жүз Алтымыш Жети",
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

func TestToWords_KyrgyzKyrgyzstan_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One som",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Бир Сом",
		},
		{
			name:   "Multiple som",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Беш Сом",
		},
		{
			name:   "Zero som",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Нөл Сом",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Бир миллион Сом",
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

func TestToWords_KyrgyzKyrgyzstan_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Som and one tyiyn",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Беш Сом Жана Бир Тыйын",
		},
		{
			name:   "Som and multiple tyiyn",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Беш Сом Жана Жыйырма Беш Тыйын",
		},
		{
			name:   "Only tyiyn",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Нөл Сом Жана Токсон Тогуз Тыйын",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Бир Бир миң Эки Бир жүз Отуз Төрт Сом Жана Элүү Алты Тыйын",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Бир Бир жүз Жыйырма Үч Жана Кырк Беш",
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

func TestToWords_KyrgyzKyrgyzstan_Negative(t *testing.T) {
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
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Минус Элүү",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Минус Жыйырма Беш Сом Жана Жетимиш Беш Тыйын",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "Терс",
					},
				},
			},
			expected: "Терс Бир жүз Сом",
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

func TestToWords_KyrgyzKyrgyzstan_Formatting(t *testing.T) {
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
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "БИР БИР ЖҮЗ ЖЫЙЫРМА ҮЧ СОМ",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "бир бир жүз жыйырма үч сом",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Тогуз Бир жүз Токсон Тогуз Гана",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Беш Бир жүз Сом Гана",
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

func TestToWords_KyrgyzKyrgyzstan_CustomCurrency(t *testing.T) {
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
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "доллар",
						Plural: "доллар",
					},
				},
			},
			expected: "Бир жүз доллар",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "евро",
						Plural:           "евро",
						FractionUnitName: "цент",
						FractionPlural:   "цент",
					},
				},
			},
			expected: "Элүү евро Жана Жыйырма Беш цент",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "рубль",
						Plural: "рубль",
					},
				},
			},
			expected: "Бир рубль",
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

func TestToWords_KyrgyzKyrgyzstan_EdgeCases(t *testing.T) {
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
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Нөл Сом Жана Бир Тыйын",
		},
		{
			name:   "Eleven",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Он бир",
		},
		{
			name:   "Twelve",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Он эки",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Жыйырма Бир",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Бир Бир жүз Бир",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "ky-KG",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Бир Бир миң Бир",
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
