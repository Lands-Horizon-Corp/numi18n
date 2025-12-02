package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_AlbanianAlbania_Numbers(t *testing.T) {
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
				Locale: "sq-AL",
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
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Pesë",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Pesëmbëdhjetë",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tridhjetë",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Dyzet e shtatë",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Njëqind",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Dy njëqind pesëdhjetë gjashtë",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Një mijë",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Një milion",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Një një milion dy njëqind tridhjetë katër një mijë pesë njëqind gjashtëdhjetë shtatë",
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

func TestToWords_AlbanianAlbania_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One lek",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Një Lek",
		},
		{
			name:   "Multiple lekë",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Pesë Lekë",
		},
		{
			name:   "Zero lekë",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Zero Lekë",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Një milion Lekë",
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

func TestToWords_AlbanianAlbania_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Lekë and one qindarkë",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Pesë Lekë dhe një Qindarkë",
		},
		{
			name:   "Lekë and multiple qindarka",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Pesë Lekë dhe njëzet e pesë Qindarka",
		},
		{
			name:   "Only qindarka",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Zero Lekë dhe nëntëdhjetë e nëntë Qindarka",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Një njëqind njëzet tre dhe dyzet e pesë",
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

func TestToWords_AlbanianAlbania_Negative(t *testing.T) {
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
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Minus pesëdhjetë",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Minus njëzet e pesë Lekë dhe shtatëdhjetë e pesë Qindarka",
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

func TestToWords_AlbanianAlbania_Formatting(t *testing.T) {
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
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "NJË NJËQIND NJËZET TRE LEKË",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "një njëqind njëzet tre lekë",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Nëntë njëqind nëntëdhjetë nëntë vetëm",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "sq-AL",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Pesëqind Lekë vetëm",
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
