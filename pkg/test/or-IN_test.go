package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_Odia_Numbers(t *testing.T) {
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
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ଶୂନ୍ୟ",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ପାଞ୍ଚ",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ପନ୍ଦର",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ତିରିଶ",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ଚାଳିଶ ସାତ",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ଏକ ଶହ",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ଦୁଇ ଏକ ଶହ ପଚାଶ ଛଅ",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ଏକ ହଜାର",
		},
		{
			name:   "One hundred thousand (exact mapping)",
			amount: 100000,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ଏକ ଲକ୍ଷ",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ଏକ ଦଶ ଲକ୍ଷ ଦୁଇ ଏକ ଶହ ତିରିଶ ଚାରି ଏକ ହଜାର ପାଞ୍ଚ ଏକ ଶହ ଷାଠିଏ ସାତ",
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

func TestToWords_Odia_Currency(t *testing.T) {
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
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ଏକ ଭାରତୀୟ ଟଙ୍କା",
		},
		{
			name:   "Multiple rupees",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ପାଞ୍ଚ ଭାରତୀୟ ଟଙ୍କା",
		},
		{
			name:   "Zero rupees",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ଶୂନ୍ୟ ଭାରତୀୟ ଟଙ୍କା",
		},
		{
			name:   "Large amount",
			amount: 100000,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ଏକ ଲକ୍ଷ ଭାରତୀୟ ଟଙ୍କା",
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

func TestToWords_Odia_Decimals(t *testing.T) {
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
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ପାଞ୍ଚ ଭାରତୀୟ ଟଙ୍କା ଏବଂ ଏକ ପଇସା",
		},
		{
			name:   "Rupees and multiple paisa",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ପାଞ୍ଚ ଭାରତୀୟ ଟଙ୍କା ଏବଂ କୋଡ଼ିଏ ପାଞ୍ଚ ପଇସା",
		},
		{
			name:   "Only paisa",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ଶୂନ୍ୟ ଭାରତୀୟ ଟଙ୍କା ଏବଂ ନବେ ନଅ ପଇସା",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ଏକ ଏକ ହଜାର ଦୁଇ ଏକ ଶହ ତିରିଶ ଚାରି ଭାରତୀୟ ଟଙ୍କା ଏବଂ ପଚାଶ ଛଅ ପଇସା",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ଏକ ଏକ ଶହ କୋଡ଼ିଏ ତିନି ଏବଂ ଚାଳିଶ ପାଞ୍ଚ",
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

func TestToWords_Odia_Negative(t *testing.T) {
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
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ମାଇନସ ପଚାଶ",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ମାଇନସ କୋଡ଼ିଏ ପାଞ୍ଚ ଭାରତୀୟ ଟଙ୍କା ଏବଂ ସତୁରି ପାଞ୍ଚ ପଇସା",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "ଋଣ",
					},
				},
			},
			expected: "ଋଣ ଏକ ଶହ ଭାରତୀୟ ଟଙ୍କା",
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

func TestToWords_Odia_Formatting(t *testing.T) {
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
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "ଏକ ଏକ ଶହ କୋଡ଼ିଏ ତିନି ଭାରତୀୟ ଟଙ୍କା",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "ଏକ ଏକ ଶହ କୋଡ଼ିଏ ତିନି ଭାରତୀୟ ଟଙ୍କା",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "ନଅ ଏକ ଶହ ନବେ ନଅ କେବଳ",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "ପାଞ୍ଚ ଶହ ଭାରତୀୟ ଟଙ୍କା କେବଳ",
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

func TestToWords_Odia_CustomCurrency(t *testing.T) {
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
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "ଡଲାର",
						Plural: "ଡଲାର",
					},
				},
			},
			expected: "ଏକ ଶହ ଡଲାର",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "ୟୁରୋ",
						Plural:           "ୟୁରୋ",
						FractionUnitName: "ସେଣ୍ଟ",
						FractionPlural:   "ସେଣ୍ଟ",
					},
				},
			},
			expected: "ପଚାଶ ୟୁରୋ ଏବଂ କୋଡ଼ିଏ ପାଞ୍ଚ ସେଣ୍ଟ",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "ପାଉଣ୍ଡ",
						Plural: "ପାଉଣ୍ଡ",
					},
				},
			},
			expected: "ଏକ ପାଉଣ୍ଡ",
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

func TestToWords_Odia_EdgeCases(t *testing.T) {
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
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ଶୂନ୍ୟ ଭାରତୀୟ ଟଙ୍କା ଏବଂ ଏକ ପଇସା",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ଏଗାର",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ବାର",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "କୋଡ଼ିଏ ଏକ",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ଏକ ଏକ ଶହ ଏକ",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "or-IN",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ଏକ ଏକ ହଜାର ଏକ",
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
