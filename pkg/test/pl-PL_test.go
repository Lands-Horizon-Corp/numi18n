package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_Polish_Numbers(t *testing.T) {
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
				Locale: "pl-PL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Zero",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "pl-PL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Pięć",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "pl-PL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Piętnaście",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "pl-PL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Trzydzieści",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "pl-PL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Czterdzieści siedem",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "pl-PL",
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
				Locale: "pl-PL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Dwa sto pięćdziesiąt sześć",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "pl-PL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tysiąc",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "pl-PL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Milion",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "pl-PL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Jeden milion dwa sto trzydzieści cztery tysiąc pięć sto sześćdziesiąt siedem",
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

func TestToWords_Polish_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One zloty",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "pl-PL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Jeden Złoty",
		},
		{
			name:   "Multiple zloty",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "pl-PL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Pięć Złote",
		},
		{
			name:   "Zero zloty",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "pl-PL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Zero Złote",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "pl-PL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Milion Złote",
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
