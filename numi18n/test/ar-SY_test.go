package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_ArabicSyria_Numbers(t *testing.T) {
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
				Locale: "ar-SY",
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
				Locale: "ar-SY",
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
				Locale: "ar-SY",
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
				Locale: "ar-SY",
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
				Locale: "ar-SY",
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
				Locale: "ar-SY",
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
				Locale: "ar-SY",
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
				Locale: "ar-SY",
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
				Locale: "ar-SY",
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
				Locale: "ar-SY",
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

func TestToWords_ArabicSyria_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One lira",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-SY",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "واحد ليرة سورية",
		},
		{
			name:   "Multiple liras",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-SY",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "خمسة ليرات سورية",
		},
		{
			name:   "Zero liras",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-SY",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "صفر ليرات سورية",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-SY",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "مليون ليرات سورية",
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

func TestToWords_ArabicSyria_EdgeCases(t *testing.T) {
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
				Locale: "ar-SY",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "صفر ليرات سورية و واحد قرش",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-SY",
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
				Locale: "ar-SY",
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
				Locale: "ar-SY",
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
				Locale: "ar-SY",
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
				Locale: "ar-SY",
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
