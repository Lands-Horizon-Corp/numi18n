package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_TamilIndia_Numbers(t *testing.T) {
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
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "பூஜ்யம்",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ஐந்து",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "பதினைந்து",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "முப்பது",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "நாற்பத்து ஏழு",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "நூறு",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "இரண்டு நூறு ஐம்பது ஆறு",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ஒரு ஆயிரம்",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ஒரு பத்து லட்சம்",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ஒன்று ஒரு பத்து லட்சம் இரண்டு நூறு முப்பது நான்கு ஒரு ஆயிரம் ஐந்து நூறு அறுபது ஏழு",
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

func TestToWords_TamilIndia_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One rupee",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ஒன்று ரூபாய்",
		},
		{
			name:   "Multiple rupees",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ஐந்து ரூபாய்கள்",
		},
		{
			name:   "Zero rupees",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "பூஜ்யம் ரூபாய்கள்",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ஒரு பத்து லட்சம் ரூபாய்கள்",
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

func TestToWords_TamilIndia_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Rupees and one paisa",
			amount: 5.01,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ஐந்து ரூபாய்கள் மற்றும் ஒன்று பைசா",
		},
		{
			name:   "Rupees and multiple paisa",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ஐந்து ரூபாய்கள் மற்றும் இருபத்து ஐந்து பைசாக்கள்",
		},
		{
			name:   "Only paisa",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "பூஜ்யம் ரூபாய்கள் மற்றும் தொண்ணூற்று ஒன்பது பைசாக்கள்",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ஒன்று ஒரு ஆயிரம் இரண்டு நூறு முப்பது நான்கு ரூபாய்கள் மற்றும் ஐம்பத்து ஆறு பைசாக்கள்",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ஒன்று நூறு இருபது மூன்று மற்றும் நாற்பத்து ஐந்து",
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

func TestToWords_TamilIndia_Negative(t *testing.T) {
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
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "கழித்தல் ஐம்பது",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "கழித்தல் இருபத்து ஐந்து ரூபாய்கள் மற்றும் எழுபத்து ஐந்து பைசாக்கள்",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &numi18n.OverrideOptions{
						NegativeWord: "எதிர்",
					},
				},
			},
			expected: "எதிர் நூறு ரூபாய்கள்",
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

func TestToWords_TamilIndia_Formatting(t *testing.T) {
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
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "ஒன்று நூறு இருபது மூன்று ரூபாய்கள்",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "ஒன்று நூறு இருபது மூன்று ரூபாய்கள்",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "ஒன்பது நூறு தொண்ணூறு ஒன்பது மட்டும்",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "ஐந்நூறு ரூபாய்கள் மட்டும்",
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

func TestToWords_TamilIndia_CustomCurrency(t *testing.T) {
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
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "டாலர்",
						Plural: "டாலர்கள்",
					},
				},
			},
			expected: "நூறு டாலர்கள்",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "யூரோ",
						Plural:           "யூரோக்கள்",
						FractionUnitName: "சென்ட்",
						FractionPlural:   "சென்ட்கள்",
					},
				},
			},
			expected: "ஐம்பது யூரோக்கள் மற்றும் இருபத்து ஐந்து சென்ட்கள்",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "பவுண்ட்",
						Plural: "பவுண்ட்கள்",
					},
				},
			},
			expected: "ஒன்று பவுண்ட்",
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

func TestToWords_TamilIndia_EdgeCases(t *testing.T) {
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
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "பூஜ்யம் ரூபாய்கள் மற்றும் ஒன்று பைசா",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "பதினொன்று",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "பன்னிரண்டு",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "இருபத்து ஒன்று",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ஒன்று நூறு ஒன்று",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &numi18n.NumI18NOptions{
				Locale: "ta-IN",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ஒன்று ஒரு ஆயிரம் ஒன்று",
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
