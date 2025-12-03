package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_ArabicQatar_Numbers(t *testing.T) {
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
				Locale: "ar-QA",
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
				Locale: "ar-QA",
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
				Locale: "ar-QA",
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
				Locale: "ar-QA",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ثلاثون",
		},
		{
			name:   "Compound numbers",
			amount: 34,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-QA",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ثلاثون أربعة",
		},
		{
			name:   "Hundreds",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-QA",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "مئة",
		},
		{
			name:   "Thousands",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-QA",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ألف",
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

func TestToWords_ArabicQatar_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One riyal",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-QA",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "واحد ريال قطري",
		},
		{
			name:   "Multiple riyals",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-QA",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "خمسة ريالات قطرية",
		},
		{
			name:   "Zero riyals",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-QA",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "صفر ريالات قطرية",
		},
		{
			name:   "Large amount",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "ar-QA",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ألف ريالات قطرية",
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
