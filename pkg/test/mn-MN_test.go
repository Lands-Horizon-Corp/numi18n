package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_MongolianMongolia_Numbers(t *testing.T) {
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
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Тэг",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Таван",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Арван таван",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Гучин",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Дөчин долоон",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Нэг зуу",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Хоёр нэг зуу тавин зургаан",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Нэг мянга",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Сая",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Нэгэн сая хоёр нэг зуу гучин дөрвөн нэг мянга таван нэг зуу жаран долоон",
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

func TestToWords_MongolianMongolia_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One togrog",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Нэгэн Төгрөг",
		},
		{
			name:   "Multiple togrogs",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Таван Төгрөг",
		},
		{
			name:   "Zero togrogs",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Тэг Төгрөг",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Сая Төгрөг",
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

func TestToWords_MongolianMongolia_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Togrogs and one mongo",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Таван Төгрөг ба нэгэн Мөнгө",
		},
		{
			name:   "Togrogs and multiple mongos",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Таван Төгрөг ба хорин таван Мөнгө",
		},
		{
			name:   "Only mongos",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Тэг Төгрөг ба ерэн есөн Мөнгө",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Нэгэн нэг мянга хоёр нэг зуу гучин дөрвөн Төгрөг ба тавин зургаан Мөнгө",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Нэгэн нэг зуу хорин гурван ба дөчин таван",
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

func TestToWords_MongolianMongolia_Negative(t *testing.T) {
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
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Хасах тавин",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Хасах хорин таван Төгрөг ба далан таван Мөнгө",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "сөрөг",
					},
				},
			},
			expected: "Сөрөг Нэг зуу Төгрөг",
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

func TestToWords_MongolianMongolia_Formatting(t *testing.T) {
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
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "НЭГЭН НЭГ ЗУУ ХОРИН ГУРВАН ТӨГРӨГ",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "нэгэн нэг зуу хорин гурван төгрөг",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Есөн нэг зуу ерэн есөн зөвхөн",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Таван зуу Төгрөг зөвхөн",
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

func TestToWords_MongolianMongolia_CustomCurrency(t *testing.T) {
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
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "доллар",
						Plural: "доллар",
					},
				},
			},
			expected: "Нэг зуу доллар",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
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
			expected: "Тавин евро ба хорин таван цент",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "фунт",
						Plural: "фунт",
					},
				},
			},
			expected: "Нэгэн фунт",
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

func TestToWords_MongolianMongolia_EdgeCases(t *testing.T) {
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
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Тэг Төгрөг ба нэгэн Мөнгө",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Арван нэгэн",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Арван хоёр",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Хорин нэгэн",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Нэгэн нэг зуу нэгэн",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "mn-MN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Нэгэн нэг мянга нэгэн",
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
