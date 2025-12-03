package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_Maltese_Numbers(t *testing.T) {
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
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Żero",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ħamsa",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ħmistax",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tletin",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Erbgħin u sebgħa",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Mija waħda",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tnejn mija ħamsin sitta",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Elf",
		},
		{
			name:   "One million",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Miljun",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Wieħed miljun tnejn mija tletin erbgħa elf ħamsa mija sittin sebgħa",
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

func TestToWords_Maltese_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One euro",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Wieħed Ewro",
		},
		{
			name:   "Multiple euros",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Ħamsa Ewro",
		},
		{
			name:   "Zero euros",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Żero Ewro",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Miljun Ewro",
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

func TestToWords_Maltese_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Euros and one centezimu",
			amount: 5.01,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Ħamsa Ewro u wieħed Ċenteżmu",
		},
		{
			name:   "Euros and multiple centezimi",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Ħamsa Ewro u għoxrin u ħamsa Ċenteżmi",
		},
		{
			name:   "Only centezimi",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Żero Ewro u disgħin u disgħa Ċenteżmi",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Wieħed elf tnejn mija tletin erbgħa Ewro u ħamsin u sitta Ċenteżmi",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Wieħed mija għoxrin tlieta u erbgħin u ħamsa",
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

func TestToWords_Maltese_Negative(t *testing.T) {
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
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Inqas ħamsin",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Inqas għoxrin u ħamsa Ewro u sebgħin u ħamsa Ċenteżmi",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &numi18n.OverrideOptions{
						NegativeWord: "negatif",
					},
				},
			},
			expected: "Negatif Mija waħda Ewro",
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

func TestToWords_Maltese_Formatting(t *testing.T) {
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
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "WIEĦED MIJA GĦOXRIN TLIETA EWRO",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "wieħed mija għoxrin tlieta ewro",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Disgħa mija disgħin disgħa biss",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "Ħames mija Ewro biss",
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

func TestToWords_Maltese_CustomCurrency(t *testing.T) {
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
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "dinar",
						Plural: "dinari",
					},
				},
			},
			expected: "Mija waħda dinari",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "lira",
						Plural:           "liri",
						FractionUnitName: "mil",
						FractionPlural:   "mil",
					},
				},
			},
			expected: "Ħamsin liri u għoxrin u ħamsa mil",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "dollaru",
						Plural: "dollari",
					},
				},
			},
			expected: "Wieħed dollaru",
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

func TestToWords_Maltese_EdgeCases(t *testing.T) {
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
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Żero Ewro u wieħed Ċenteżmu",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Ħdax",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Tnax",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Għoxrin u wieħed",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Wieħed mija wieħed",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Wieħed elf wieħed",
		},
		{
			name:   "Two hundred (mitejn)",
			amount: 200,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Mitejn",
		},
		{
			name:   "Two thousand (elfejn)",
			amount: 2000,
			options: &numi18n.NumI18NOptions{
				Locale: "mt-MT",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Elfejn",
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
