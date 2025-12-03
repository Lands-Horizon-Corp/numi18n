package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_Yiddish001_Numbers(t *testing.T) {
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
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "נול",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "פֿינף",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "פֿופֿצן",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "דרײַסיק",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "זיבן און פֿערציק",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "הונדערט",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "צוויי איין הונדערט פֿופֿציק זעקס",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "איין טויזנט",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "איין מיליאָן",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "איין איין מיליאָן צוויי איין הונדערט דרײַסיק פֿיר איין טויזנט פֿינף איין הונדערט זעכציק זיבן",
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

func TestToWords_Yiddish001_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One dollar",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "איין דאָלער",
		},
		{
			name:   "Multiple dollars",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "פֿינף דאָלערס",
		},
		{
			name:   "Zero dollars",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "נול דאָלערס",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "איין מיליאָן דאָלערס",
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

func TestToWords_Yiddish001_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Dollars and one cent",
			amount: 5.01,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "פֿינף דאָלערס און איין צענט",
		},
		{
			name:   "Dollars and multiple cents",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "פֿינף דאָלערס און פֿינף און צוואַנציק צענט",
		},
		{
			name:   "Only cents",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "נול דאָלערס און נײַנצן און אַכציק צענט",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "איין איין טויזנט צוויי איין הונדערט דרײַסיק פֿיר דאָלערס און זעקס און פֿופֿציק צענט",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "איין איין הונדערט צוואַנציק דרײַ און פֿינף און פֿערציק",
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

func TestToWords_Yiddish001_Negative(t *testing.T) {
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
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "מינוס פֿופֿציק",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "מינוס פֿינף און צוואַנציק דאָלערס און פֿינף און זיבעציק צענט",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &numi18n.OverrideOptions{
						NegativeWord: "נעגאַטיוו",
					},
				},
			},
			expected: "נעגאַטיוו הונדערט דאָלערס",
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

func TestToWords_Yiddish001_Formatting(t *testing.T) {
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
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "איין איין הונדערט צוואַנציק דרײַ דאָלערס",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "איין איין הונדערט צוואַנציק דרײַ דאָלערס",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "נײַן איין הונדערט נײַנציק נײַן בלויז",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "פֿינף הונדערט דאָלערס בלויז",
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

func TestToWords_Yiddish001_CustomCurrency(t *testing.T) {
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
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "שעקל",
						Plural: "שעקלען",
					},
				},
			},
			expected: "הונדערט שעקלען",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "יורו",
						Plural:           "יוראָס",
						FractionUnitName: "צענט",
						FractionPlural:   "צענטן",
					},
				},
			},
			expected: "פֿופֿציק יוראָס און פֿינף און צוואַנציק צענטן",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "פּונט",
						Plural: "פּונטן",
					},
				},
			},
			expected: "איין פּונט",
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

func TestToWords_Yiddish001_EdgeCases(t *testing.T) {
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
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "נול דאָלערס און איין צענט",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "עלף",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "צוועלף",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "איין און צוואַנציק",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "איין איין הונדערט איין",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &numi18n.NumI18NOptions{
				Locale: "yi-001",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "איין איין טויזנט איין",
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
