package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_KoreanSouthKorea_Numbers(t *testing.T) {
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
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "영",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "오",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "십오",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "삼십",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "사십 칠",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "일백",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "이 일백 오십 육",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "일천",
		},
		{
			name:   "Ten thousand (exact mapping)",
			amount: 10000,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "일만",
		},
		{
			name:   "One hundred million (exact mapping)",
			amount: 100000000,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "일억",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "일 백만 이 일백 삼십 사 일천 오 일백 육십 칠",
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

func TestToWords_KoreanSouthKorea_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One won",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "일 원",
		},
		{
			name:   "Multiple won",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "오 원",
		},
		{
			name:   "Zero won",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "영 원",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "백만 원",
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

func TestToWords_KoreanSouthKorea_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Won and one jeon",
			amount: 5.01,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "오 원 그리고 일 전",
		},
		{
			name:   "Won and multiple jeon",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "오 원 그리고 이십 오 전",
		},
		{
			name:   "Only jeon",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "영 원 그리고 구십 구 전",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "일 일천 이 일백 삼십 사 원 그리고 오십 육 전",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "일 일백 이십 삼 그리고 사십 오",
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

func TestToWords_KoreanSouthKorea_Negative(t *testing.T) {
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
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "마이너스 오십",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "마이너스 이십 오 원 그리고 칠십 오 전",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &numi18n.OverrideOptions{
						NegativeWord: "음수",
					},
				},
			},
			expected: "음수 일백 원",
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

func TestToWords_KoreanSouthKorea_Formatting(t *testing.T) {
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
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "일 일백 이십 삼 원",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "일 일백 이십 삼 원",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "구 일백 구십 구만",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "오 일백 원만",
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

func TestToWords_KoreanSouthKorea_CustomCurrency(t *testing.T) {
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
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "달러",
						Plural: "달러",
					},
				},
			},
			expected: "일백 달러",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "유로",
						Plural:           "유로",
						FractionUnitName: "센트",
						FractionPlural:   "센트",
					},
				},
			},
			expected: "오십 유로 그리고 이십 오 센트",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "파운드",
						Plural: "파운드",
					},
				},
			},
			expected: "일 파운드",
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

func TestToWords_KoreanSouthKorea_EdgeCases(t *testing.T) {
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
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "영 원 그리고 일 전",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "십일",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "십이",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "이십 일",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "일 일백 일",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "일 일천 일",
		},
		{
			name:   "Ten thousand one",
			amount: 10001,
			options: &numi18n.NumI18NOptions{
				Locale: "ko-KR",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "십 일천 일",
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
