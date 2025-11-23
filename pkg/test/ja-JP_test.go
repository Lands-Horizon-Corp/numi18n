package pkg

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_Japanese_Numbers(t *testing.T) {
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
				Locale: "ja-JP",
			},
			expected: "零",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "五",
		},
		{
			name:   "Ten",
			amount: 10,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "十",
		},
		{
			name:   "Eleven",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "十一",
		},
		{
			name:   "Twenty",
			amount: 20,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "二十",
		},
		{
			name:   "Twenty three",
			amount: 23,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "二十三",
		},
		{
			name:   "One hundred",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "百",
		},
		{
			name:   "One hundred twenty three",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "百二十三",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "千",
		},
		{
			name:   "Ten thousand (ichiman)",
			amount: 10000,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "一万",
		},
		{
			name:   "One hundred thousand",
			amount: 100000,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "十万",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "百万",
		},
		{
			name:   "One hundred million (ichioku)",
			amount: 100000000,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "一億",
		},
		{
			name:   "Complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "百二十三万四千五百六十七",
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

func TestToWords_Japanese_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One yen",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency: true,
				},
			},
			expected: "一円",
		},
		{
			name:   "Five yen",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency: true,
				},
			},
			expected: "五円",
		},
		{
			name:   "Zero yen",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency: true,
				},
			},
			expected: "零円",
		},
		{
			name:   "One hundred yen",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency: true,
				},
			},
			expected: "百円",
		},
		{
			name:   "One thousand yen",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency: true,
				},
			},
			expected: "千円",
		},
		{
			name:   "Ten thousand yen",
			amount: 10000,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency: true,
				},
			},
			expected: "一万円",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency: true,
				},
			},
			expected: "百万円",
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

func TestToWords_Japanese_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Yen with sen (old fractional unit)",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "五円と二十五銭",
		},
		{
			name:   "Only sen (fractional)",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "零円と九十九銭",
		},
		{
			name:   "Complex amount with decimals",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "千二百三十四円と五十六銭",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Decimal: true,
				},
			},
			expected: "百二十三と四十五",
		},
		{
			name:   "JPY precision (no decimals by default)",
			amount: 123.99,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency: true,
				},
			},
			expected: "百二十三円",
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

func TestToWords_Japanese_Negative(t *testing.T) {
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
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
				},
			},
			expected: "マイナス五十",
		},
		{
			name:   "Negative currency",
			amount: -1000,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
				},
			},
			expected: "マイナス千円",
		},
		{
			name:   "Negative with decimals",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
				},
			},
			expected: "マイナス二十五円と七十五銭",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "負",
					},
				},
			},
			expected: "負百円",
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

func TestToWords_Japanese_Formatting(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Only: true,
				},
			},
			expected: "九百九十九のみ",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency: true,
					Only:     true,
				},
			},
			expected: "五百円のみ",
		},
		{
			name:   "Uppercase (no effect in Japanese)",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "百二十三円",
		},
		{
			name:   "Lowercase (no effect in Japanese)",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "百二十三円",
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

func TestToWords_Japanese_CustomCurrency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Custom currency name (Dollar in Japanese)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "ドル",
						Plural: "ドル",
					},
				},
			},
			expected: "百ドル",
		},
		{
			name:   "Custom currency with decimals (Euro)",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency: true,
					Decimal:  true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "ユーロ",
						Plural:           "ユーロ",
						FractionUnitName: "セント",
						FractionPlural:   "セント",
					},
				},
			},
			expected: "五十ユーロと二十五セント",
		},
		{
			name:   "Custom currency (Peso)",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "ペソ",
						Plural: "ペソ",
					},
				},
			},
			expected: "千ペソ",
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

func TestToWords_Japanese_EdgeCases(t *testing.T) {
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
				Locale: "ja-JP",
				WordDetails: &pkg.WordDetails{
					Currency: true,
					Decimal:  true,
				},
			},
			expected: "零円と一銭",
		},
		{
			name:   "Ten (special case in Japanese)",
			amount: 10,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "十",
		},
		{
			name:   "One hundred (special case)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "百",
		},
		{
			name:   "One thousand (special case)",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "千",
		},
		{
			name:   "Ten thousand (ichiman)",
			amount: 10000,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "一万",
		},
		{
			name:   "One hundred million (ichioku)",
			amount: 100000000,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "一億",
		},
		{
			name:   "Numbers with repetitive digits",
			amount: 1111,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "千百十一",
		},
		{
			name:   "Large Japanese number",
			amount: 12345678,
			options: &pkg.NumI18NOptions{
				Locale: "ja-JP",
			},
			expected: "千二百三十四万五千六百七十八",
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
