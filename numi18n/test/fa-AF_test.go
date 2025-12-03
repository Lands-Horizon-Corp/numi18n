package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_PersianAfghanistan_Numbers(t *testing.T) {
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
				Locale: "fa-AF",
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
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "پنج",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "پانزده",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "سی",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "چهل هفت",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "یک صد",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "دو صد پنجاه شش",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "یک هزار",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "یک میلیون",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "یک میلیون دو صد سی چهار هزار پنج صد شصت هفت",
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

func TestToWords_PersianAfghanistan_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One afghani",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "یک افغانی",
		},
		{
			name:   "Multiple afghanis",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "پنج افغانی",
		},
		{
			name:   "Zero afghanis",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "صفر افغانی",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "یک میلیون افغانی",
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

func TestToWords_PersianAfghanistan_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Afghanis and one pul",
			amount: 5.01,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "پنج افغانی و یک پول",
		},
		{
			name:   "Afghanis and multiple puls",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "پنج افغانی و بیست پنج پول",
		},
		{
			name:   "Only puls",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "صفر افغانی و نود نه پول",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "یک هزار دو صد سی چهار افغانی و پنجاه شش پول",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "یک صد بیست سه و چهل پنج",
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

func TestToWords_PersianAfghanistan_Negative(t *testing.T) {
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
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "منفی پنجاه",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "منفی بیست پنج افغانی و هفتاد پنج پول",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &numi18n.OverrideOptions{
						NegativeWord: "منها",
					},
				},
			},
			expected: "منها یک صد افغانی",
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

func TestToWords_PersianAfghanistan_Formatting(t *testing.T) {
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
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "یک صد بیست سه افغانی",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "یک صد بیست سه افغانی",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "نه صد نود نه فقط",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "پنج صد افغانی فقط",
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

func TestToWords_PersianAfghanistan_CustomCurrency(t *testing.T) {
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
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "ریال",
						Plural: "ریال",
					},
				},
			},
			expected: "یک صد ریال",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "یورو",
						Plural:           "یورو",
						FractionUnitName: "سنت",
						FractionPlural:   "سنت",
					},
				},
			},
			expected: "پنجاه یورو و بیست پنج سنت",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "پوند",
						Plural: "پوند",
					},
				},
			},
			expected: "یک پوند",
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

func TestToWords_PersianAfghanistan_EdgeCases(t *testing.T) {
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
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "صفر افغانی و یک پول",
		},
		{
			name:   "Eleven",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "یازده",
		},
		{
			name:   "Twelve",
			amount: 12,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "دوازده",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "بیست یک",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "یک صد یک",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &numi18n.NumI18NOptions{
				Locale: "fa-AF",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "یک هزار یک",
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
