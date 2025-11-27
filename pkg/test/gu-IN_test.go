package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_Gujarati_Numbers(t *testing.T) {
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
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "શૂન્ય",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "પાંચ",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "પંદર",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ત્રીસ",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ચાલીસ સાત",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "એક સો",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "બે એક સો પચાસ છ",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "એક હજાર",
		},
		{
			name:   "One lakh (exact mapping)",
			amount: 100000,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "એક લાખ",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "એક લાખ",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "એક એક લાખ બે એક સો ત્રીસ ચાર એક હજાર પાંચ એક સો સાઠ સાત",
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

func TestToWords_Gujarati_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One rupee",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "એક રૂપિયો",
		},
		{
			name:   "Multiple rupees",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "પાંચ રૂપિયા",
		},
		{
			name:   "Zero rupees",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "શૂન્ય રૂપિયા",
		},
		{
			name:   "Large amount",
			amount: 100000,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "એક લાખ રૂપિયા",
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

func TestToWords_Gujarati_Decimals(t *testing.T) {
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
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "પાંચ રૂપિયા અને એક પૈસો",
		},
		{
			name:   "Rupees and multiple paisa",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "પાંચ રૂપિયા અને વીસ પાંચ પૈસા",
		},
		{
			name:   "Only paisa",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "શૂન્ય રૂપિયા અને નેવું નવ પૈસા",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "એક એક હજાર બે એક સો ત્રીસ ચાર રૂપિયા અને પચાસ છ પૈસા",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "એક એક સો વીસ ત્રણ અને ચાલીસ પાંચ",
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

func TestToWords_Gujarati_Negative(t *testing.T) {
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
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ઋણ પચાસ",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ઋણ વીસ પાંચ રૂપિયા અને સિત્તેર પાંચ પૈસા",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "નકારાત્મક",
					},
				},
			},
			expected: "નકારાત્મક એક સો રૂપિયા",
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

func TestToWords_Gujarati_Formatting(t *testing.T) {
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
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "એક એક સો વીસ ત્રણ રૂપિયા",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "એક એક સો વીસ ત્રણ રૂપિયા",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "નવ એક સો નેવું નવ માત્ર",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "પાંચ એક સો રૂપિયા માત્ર",
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

func TestToWords_Gujarati_CustomCurrency(t *testing.T) {
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
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "ડોલર",
						Plural: "ડોલર",
					},
				},
			},
			expected: "એક સો ડોલર",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "યુરો",
						Plural:           "યુરો",
						FractionUnitName: "સેન્ટ",
						FractionPlural:   "સેન્ટ",
					},
				},
			},
			expected: "પચાસ યુરો અને વીસ પાંચ સેન્ટ",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "પાઉન્ડ",
						Plural: "પાઉન્ડ",
					},
				},
			},
			expected: "એક પાઉન્ડ",
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

func TestToWords_Gujarati_EdgeCases(t *testing.T) {
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
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "શૂન્ય રૂપિયા અને એક પૈસો",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "અગિયાર",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "બાર",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "વીસ એક",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "એક એક સો એક",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "gu-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "એક એક હજાર એક",
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
