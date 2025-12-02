package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_SomaliSomalia_Numbers(t *testing.T) {
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
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Eber",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Shan",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Shan iyo toban",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Soddonka",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Afaratan iyo toddoba",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Boqol",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Laba boqol kontan lix",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Kun",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Milyan",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Kow milyan laba boqol soddonka afar kun shan boqol lixdatan toddoba",
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

func TestToWords_SomaliSomalia_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One shilling",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Kow Shilin Soomaaliyeed",
		},
		{
			name:   "Multiple shillings",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Shan Shilin Soomaaliyeed",
		},
		{
			name:   "Zero shillings",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Eber Shilin Soomaaliyeed",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Milyan Shilin Soomaaliyeed",
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

func TestToWords_SomaliSomalia_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Shillings and one cent",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Shan Shilin Soomaaliyeed iyo kow Cent",
		},
		{
			name:   "Shillings and multiple cents",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Shan Shilin Soomaaliyeed iyo labaatan iyo shan Cent",
		},
		{
			name:   "Only cents",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Eber Shilin Soomaaliyeed iyo sagaashan iyo sagaal Cent",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Kow boqol labaatan saddex iyo afaratan iyo shan",
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

func TestToWords_SomaliSomalia_Negative(t *testing.T) {
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
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Naqsi kontan",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Naqsi labaatan iyo shan Shilin Soomaaliyeed iyo todobaatan iyo shan Cent",
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

func TestToWords_SomaliSomalia_Formatting(t *testing.T) {
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
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "KOW BOQOL LABAATAN SADDEX SHILIN SOOMAALIYEED",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "kow boqol labaatan saddex shilin soomaaliyeed",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Sagaal boqol sagaashan sagaal kaliya",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "so-SO",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Shan boqol Shilin Soomaaliyeed kaliya",
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
