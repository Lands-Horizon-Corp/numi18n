package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_NorthernSamiNorway_Numbers(t *testing.T) {
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
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Nolla",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Viđa",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Viđanuppelohkái",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Golbmalogi",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Njeljelogi ja čieža",
		},
		{
			name:   "One hundred",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Čuođi",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Guokte čuođi viđalogi guhtta",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Duhát",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Miljonat",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Okta miljonat guokte čuođi golbmalogi njealje duhát viđa čuođi guhttálogi čieža",
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

func TestToWords_NorthernSamiNorway_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One Norwegian krone",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Okta Norgga kruvdno",
		},
		{
			name:   "Multiple Norwegian kroner",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Viđa Norgga kruvdnot",
		},
		{
			name:   "Zero Norwegian kroner",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Nolla Norgga kruvdnot",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Miljonat Norgga kruvdnot",
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

func TestToWords_NorthernSamiNorway_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Kroner and one øre",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Viđa Norgga kruvdnot ja okta Øre",
		},
		{
			name:   "Kroner and multiple øre",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Viđa Norgga kruvdnot ja guoktelogi ja viđa Øret",
		},
		{
			name:   "Only øre",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Nolla Norgga kruvdnot ja ovccilogi ja ovcci Øret",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Okta duhát guokte čuođi golbmalogi njealje Norgga kruvdnot ja viđalogi ja guhtta Øret",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Okta čuođi guoktelogi golbma ja njeljelogi ja viđa",
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

func TestToWords_NorthernSamiNorway_Negative(t *testing.T) {
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
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Miinus viđalogi",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Miinus guoktelogi ja viđa Norgga kruvdnot ja čiežalogi ja viđa Øret",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "negatiiva",
					},
				},
			},
			expected: "Negatiiva Čuođi Norgga kruvdnot",
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

func TestToWords_NorthernSamiNorway_Formatting(t *testing.T) {
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
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "OKTA ČUOĐI GUOKTELOGI GOLBMA NORGGA KRUVDNOT",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "okta čuođi guoktelogi golbma norgga kruvdnot",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Ovcci čuođi ovccilogi ovcci dušše",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Viđa čuođi Norgga kruvdnot dušše",
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

func TestToWords_NorthernSamiNorway_CustomCurrency(t *testing.T) {
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
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "euro",
						Plural: "euro",
					},
				},
			},
			expected: "Čuođi euro",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "euro",
						Plural:           "euro",
						FractionUnitName: "sentti",
						FractionPlural:   "sentti",
					},
				},
			},
			expected: "Viđalogi euro ja guoktelogi ja viđa sentti",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "dollár",
						Plural: "dollára",
					},
				},
			},
			expected: "Okta dollár",
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

func TestToWords_NorthernSamiNorway_EdgeCases(t *testing.T) {
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
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Nolla Norgga kruvdnot ja okta Øre",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Oktnuppelohkái",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Guoktenuppelohkái",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Guoktelogi ja okta",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Okta čuođi okta",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "se-NO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Okta duhát okta",
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
