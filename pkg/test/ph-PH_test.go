package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_Filipino_Numbers(t *testing.T) {
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
				Locale: "ph-PH",
			},
			expected: "Zero",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
			},
			expected: "Lima",
		},
		{
			name:   "Ten",
			amount: 10,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
			},
			expected: "Sampu",
		},
		{
			name:   "Eleven",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
			},
			expected: "Labing-isa",
		},
		{
			name:   "Twenty",
			amount: 20,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
			},
			expected: "Dalawampu",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
			},
			expected: "Dalawampu Isa",
		},
		{
			name:   "One hundred (special case)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
			},
			expected: "Isang daan",
		},
		{
			name:   "One hundred twenty three",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
			},
			expected: "Isang daan Dalawampu Tatlo",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
			},
			expected: "Isang Libo",
		},
		{
			name:   "Ten thousand",
			amount: 10000,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
			},
			expected: "Sampu Libo",
		},
		{
			name:   "One hundred thousand",
			amount: 100000,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
			},
			expected: "Isang daan Libo",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
			},
			expected: "Isang Milyun",
		},
		{
			name:   "Complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
			},
			expected: "Isang Milyun Dalawa Daan Tatlong pu Apat Libo Lima Daan Animnapu Pito",
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

func TestToWords_Filipino_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One peso",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &pkg.WordDetails{
					Currency: true,
				},
			},
			expected: "Isa Peso",
		},
		{
			name:   "Five pesos",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &pkg.WordDetails{
					Currency: true,
				},
			},
			expected: "Lima Piso",
		},
		{
			name:   "One hundred pesos",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &pkg.WordDetails{
					Currency: true,
				},
			},
			expected: "Isang daan Piso",
		},
		{
			name:   "One thousand pesos",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &pkg.WordDetails{
					Currency: true,
				},
			},
			expected: "Isang Libo Piso",
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

func TestToWords_Filipino_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Zero point five (fraction as percentage)",
			amount: 0.5,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &pkg.WordDetails{
					Decimal: true,
				},
			},
			expected: "Zero at Limampu", // 0.5 = 0 and 50/100
		},
		{
			name:   "One point two five (fraction as percentage)",
			amount: 1.25,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &pkg.WordDetails{
					Decimal: true,
				},
			},
			expected: "Isa at Dalawampu Lima", // 1.25 = 1 and 25/100
		},
		{
			name:   "Three point one four (fraction as percentage)",
			amount: 3.14,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &pkg.WordDetails{
					Decimal: true,
				},
			},
			expected: "Tatlo at Labing-apat", // 3.14 = 3 and 14/100
		},
		{
			name:   "Ten point zero one (fraction as percentage)",
			amount: 10.01,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &pkg.WordDetails{
					Decimal: true,
				},
			},
			expected: "Sampu at Isa", // 10.01 = 10 and 1/100
		},
		{
			name:   "Hundred point five (fraction as percentage)",
			amount: 100.5,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &pkg.WordDetails{
					Decimal: true,
				},
			},
			expected: "Isang daan at Limampu", // 100.5 = 100 and 50/100
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

func TestToWords_Filipino_Negative(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Negative one",
			amount: -1,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
				},
			},
			expected: "minus Isa",
		},
		{
			name:   "Negative five",
			amount: -5,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
				},
			},
			expected: "minus Lima",
		},
		{
			name:   "Negative hundred",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
				},
			},
			expected: "minus Isang daan",
		},
		{
			name:   "Negative thousand",
			amount: -1000,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
				},
			},
			expected: "minus Isang Libo",
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

func TestToWords_Filipino_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Large number",
			amount: 999999999,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
			},
			expected: "Siyam Daan Siyamnapu Siyam Milyun Siyam Daan Siyamnapu Siyam Libo Siyam Daan Siyamnapu Siyam",
		},
		{
			name:   "Small decimal with decimal flag",
			amount: 0.01,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
				WordDetails: &pkg.WordDetails{
					Decimal: true,
				},
			},
			expected: "Zero at Isa", // 0.01 = 0 and 1/100
		},
		{
			name:   "Round million",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
			},
			expected: "Isang Milyun",
		},
		{
			name:   "Eleven hundred",
			amount: 1100,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
			},
			expected: "Isang Libo Isang daan",
		},
		{
			name:   "Twelve hundred",
			amount: 1200,
			options: &pkg.NumI18NOptions{
				Locale: "ph-PH",
			},
			expected: "Isang Libo Dalawa Daan",
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
