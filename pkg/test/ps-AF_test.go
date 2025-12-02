package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_Pashto_Numbers(t *testing.T) {
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
				Locale: "ps-AF",
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
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "پنځه",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "پنځلس",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "دیرش",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "څلویښت اووه",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "یو سل",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "دوه یو سل پنځوس شپږ",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "یو زره",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "یو میلیون",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "یو یو میلیون دوه یو سل دیرش څلور یو زره پنځه یو سل شپیته اووه",
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

func TestToWords_Pashto_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One afghani",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "یو افغانۍ",
		},
		{
			name:   "Multiple afghani",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "پنځه افغانۍ",
		},
		{
			name:   "Zero afghani",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "صفر افغانۍ",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "یو میلیون افغانۍ",
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

func TestToWords_Pashto_Fractional(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Simple decimal",
			amount: 1.5,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "یو او پنځوس",
		},
		{
			name:   "Complex decimal",
			amount: 123.456,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "یو یو سل شل درې او څلویښت پنځه",
		},
		{
			name:   "Decimal with currency",
			amount: 10.25,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "لس افغانۍ او شل پنځه پول",
		},
		{
			name:   "Zero point something",
			amount: 0.75,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "صفر افغانۍ او اویا پنځه پول",
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

func TestToWords_Pashto_Negative(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Negative integer",
			amount: -5,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "پنځه",
		},
		{
			name:   "Negative decimal",
			amount: -10.50,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "لس او پنځوس",
		},
		{
			name:   "Negative currency",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "یو سل افغانۍ",
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

func TestToWords_Pashto_Formatting(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Lowercase",
			amount: 25,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Capitalize: false,
				},
			},
			expected: "شل پنځه",
		},
		{
			name:   "Capitalized",
			amount: 25,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "شل پنځه",
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

func TestToWords_Pashto_CustomCurrency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Custom currency symbol",
			amount: 100.50,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "ډالر",
						Plural:           "ډالره",
						FractionUnitName: "سیټ",
						FractionPlural:   "سیټونه",
					},
				},
			},
			expected: "یو سل ډالره او پنځوس سیټونه",
		},
		{
			name:   "Zero with custom currency",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "ډالر",
						Plural: "ډالره",
					},
				},
			},
			expected: "صفر ډالره",
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

func TestToWords_Pashto_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Very large number",
			amount: 999999999,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "نهه یو سل نوی نهه یو میلیون نهه یو سل نوی نهه یو زره نهه یو سل نوی نهه",
		},
		{
			name:   "One billion",
			amount: 1000000000,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "یو ملیون",
		},
		{
			name:   "Complex fraction currency",
			amount: 1234.99,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "یو یو زره دوه یو سل دیرش څلور افغانۍ او نوی نهه پول",
		},
		{
			name:   "Exactly one pul (sub-unit)",
			amount: 0.01,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "صفر افغانۍ او یو پول",
		},
		{
			name:   "Eleven (teens boundary)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "یولس",
		},
		{
			name:   "Twenty-one (compound boundary)",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "شل یو",
		},
		{
			name:   "Tiny fractional amount",
			amount: 0.001,
			options: &pkg.NumI18NOptions{
				Locale: "ps-AF",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "صفر",
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
