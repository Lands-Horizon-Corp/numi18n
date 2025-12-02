package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_SindhiPakistan_Numbers(t *testing.T) {
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
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "صفر",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "پنج",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "پندريهه",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ٽيهه",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "چاليهه ۽ ست",
		},
		{
			name:   "One hundred",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "سئو",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ٻه سئو پنجاهه ڇهه",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "هزار",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "هڪ لک",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "هڪ هڪ لک ٻه سئو ٽيهه چار هزار پنج سئو سٺ ست",
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

func TestToWords_SindhiPakistan_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One Pakistani rupee",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "هڪ پاڪستاني رپيو",
		},
		{
			name:   "Multiple Pakistani rupees",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "پنج پاڪستاني رپيا",
		},
		{
			name:   "Zero Pakistani rupees",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "صفر پاڪستاني رپيا",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "هڪ لک پاڪستاني رپيا",
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

func TestToWords_SindhiPakistan_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Rupees and one paisa",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "پنج پاڪستاني رپيا ۽ هڪ پيسو",
		},
		{
			name:   "Rupees and multiple paisas",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "پنج پاڪستاني رپيا ۽ ويهه ۽ پنج پيسا",
		},
		{
			name:   "Only paisas",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "صفر پاڪستاني رپيا ۽ نوي ۽ نو پيسا",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "هڪ هزار ٻه سئو ٽيهه چار پاڪستاني رپيا ۽ پنجاهه ۽ ڇهه پيسا",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "هڪ سئو ويهه ٽي ۽ چاليهه ۽ پنج",
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

func TestToWords_SindhiPakistan_Negative(t *testing.T) {
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
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "گهٽ پنجاهه",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "گهٽ ويهه ۽ پنج پاڪستاني رپيا ۽ ستر ۽ پنج پيسا",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "منفي",
					},
				},
			},
			expected: "منفي سئو پاڪستاني رپيا",
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

func TestToWords_SindhiPakistan_Formatting(t *testing.T) {
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
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "هڪ سئو ويهه ٽي پاڪستاني رپيا",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "هڪ سئو ويهه ٽي پاڪستاني رپيا",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "نو سئو نوي نو فقط",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "پنج سئو پاڪستاني رپيا فقط",
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

func TestToWords_SindhiPakistan_CustomCurrency(t *testing.T) {
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
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "درهم",
						Plural: "درهم",
					},
				},
			},
			expected: "سئو درهم",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "يورو",
						Plural:           "يورو",
						FractionUnitName: "سينٽ",
						FractionPlural:   "سينٽ",
					},
				},
			},
			expected: "پنجاهه يورو ۽ ويهه ۽ پنج سينٽ",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "پائونڊ",
						Plural: "پائونڊ",
					},
				},
			},
			expected: "هڪ پائونڊ",
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

func TestToWords_SindhiPakistan_EdgeCases(t *testing.T) {
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
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "صفر پاڪستاني رپيا ۽ هڪ پيسو",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "يارهن",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ٻارهه",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ويهه ۽ هڪ",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "هڪ سئو هڪ",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "sd-PK",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "هڪ هزار هڪ",
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
