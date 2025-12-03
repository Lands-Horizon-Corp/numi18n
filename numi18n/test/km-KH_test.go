package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_Khmer_Numbers(t *testing.T) {
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
				Locale: "km-KH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "សូន្យ",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "km-KH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ប្រាំ",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "km-KH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "សាម\u200bស៊ប",
		},
		{
			name:   "One hundred",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "km-KH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "មួយ\u200bរយ",
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

func TestToWords_Khmer_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One riel",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "km-KH",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "មួយ រៀល",
		},
		{
			name:   "Multiple riel",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "km-KH",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ប្រាំ រៀល",
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

func TestToWords_Khmer_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Riel and sen",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "km-KH",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ប្រាំ រៀល និង ម្ភៃ ប្រាំ សេន",
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

func TestToWords_Khmer_Negative(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Negative number",
			amount: -50,
			options: &numi18n.NumI18NOptions{
				Locale: "km-KH",
				WordDetails: &numi18n.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ដក ហា\u200bស៊ប",
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

func TestToWords_Khmer_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Twenty",
			amount: 20,
			options: &numi18n.NumI18NOptions{
				Locale: "km-KH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ម្ភៃ",
		},
		{
			name:   "Ten",
			amount: 10,
			options: &numi18n.NumI18NOptions{
				Locale: "km-KH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ដប់",
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
