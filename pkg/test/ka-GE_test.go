package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_Georgian_Numbers(t *testing.T) {
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
				Locale: "ka-GE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ნული",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "ka-GE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ხუთი",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "ka-GE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Თხუთმეტი",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "ka-GE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ოცდაათი",
		},
		{
			name:   "One hundred",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "ka-GE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ერთი Ასი",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "ka-GE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ერთი Ათასი",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "ka-GE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ერთი Მილიონი",
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

func TestToWords_Georgian_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One lari",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "ka-GE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Ერთი ლარი",
		},
		{
			name:   "Multiple lari",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "ka-GE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Ხუთი ლარი",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "ka-GE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Ერთი Მილიონი ლარი",
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

func TestToWords_Georgian_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Lari and tetri",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "ka-GE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Ხუთი ლარი და ოცი ხუთი თეთრი",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "ka-GE",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Ერთი Ერთი Ასი ოცი სამი და ორმოცი ხუთი",
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

func TestToWords_Georgian_Negative(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Negative number",
			amount: -50,
			options: &pkg.NumI18NOptions{
				Locale: "ka-GE",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Მინუს ორმოცდაათი",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "ka-GE",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Მინუს ოცი ხუთი ლარი და სამოცდაათი ხუთი თეთრი",
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

func TestToWords_Georgian_Formatting(t *testing.T) {
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
				Locale: "ka-GE",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Ცხრა Ერთი Ასი ოთხმოცდაათი ცხრა მხოლოდ",
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

func TestToWords_Georgian_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Eleven",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "ka-GE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Თერთმეტი",
		},
		{
			name:   "Twelve",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "ka-GE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Თორმეტი",
		},
		{
			name:   "Twenty",
			amount: 20,
			options: &pkg.NumI18NOptions{
				Locale: "ka-GE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ოცი",
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
