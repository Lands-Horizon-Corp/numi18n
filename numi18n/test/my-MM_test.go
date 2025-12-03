package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_Myanmar_Numbers(t *testing.T) {
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
				Locale: "my-MM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "သုည",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "my-MM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ငါး",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "my-MM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ဆယ့်ငါး",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "my-MM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "သုံးဆယ်",
		},
		{
			name:   "Hundreds",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "my-MM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ငါးရာ",
		},
		{
			name:   "Thousands",
			amount: 5000,
			options: &numi18n.NumI18NOptions{
				Locale: "my-MM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ငါးထောင်",
		},
		{
			name:   "Large number",
			amount: 123456,
			options: &numi18n.NumI18NOptions{
				Locale: "my-MM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "တ တရာ နှစ်ဆယ် သုံး တထောင် လေး တရာ ငါးဆယ် ခြောက်",
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

func TestToWords_Myanmar_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One kyat",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "my-MM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "တ ကျပ်",
		},
		{
			name:   "Multiple kyat",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "my-MM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "တရာ ကျပ်",
		},
		{
			name:   "Large currency amount",
			amount: 5000,
			options: &numi18n.NumI18NOptions{
				Locale: "my-MM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ငါးထောင် ကျပ်",
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

func TestToWords_Myanmar_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Simple decimal",
			amount: 12.34,
			options: &numi18n.NumI18NOptions{
				Locale: "my-MM",
				WordDetails: &numi18n.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ဆယ့်နှစ် နှင့် သုံးဆယ့်လေး",
		},
		{
			name:   "Decimal with currency",
			amount: 25.50,
			options: &numi18n.NumI18NOptions{
				Locale: "my-MM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "နှစ်ဆယ့်ငါး ကျပ် နှင့် ငါးဆယ် ပြား",
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

func TestToWords_Myanmar_Negative(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Negative number",
			amount: -42,
			options: &numi18n.NumI18NOptions{
				Locale: "my-MM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "လေးဆယ့်နှစ်",
		},
		{
			name:   "Negative currency",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "my-MM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "တရာ ကျပ်",
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

func TestToWords_Myanmar_Formatting(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Without capitalization",
			amount: 42,
			options: &numi18n.NumI18NOptions{
				Locale: "my-MM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: false,
				},
			},
			expected: "လေးဆယ့်နှစ်",
		},
		{
			name:   "With only suffix",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "my-MM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
					Only:       true,
				},
			},
			expected: "တရာ သာ",
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

func TestToWords_Myanmar_CustomCurrency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Custom currency",
			amount: 50,
			options: &numi18n.NumI18NOptions{
				Locale: "my-MM",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "ဒေါ်လာ",
						Plural: "ဒေါ်လာ",
					},
				},
			},
			expected: "ငါးဆယ် ဒေါ်လာ",
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

func TestToWords_Myanmar_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Very large number",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "my-MM",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "သိန်း",
		},
		{
			name:   "Decimal precision",
			amount: 1.23,
			options: &numi18n.NumI18NOptions{
				Locale: "my-MM",
				WordDetails: &numi18n.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "တ နှင့် နှစ်ဆယ့်သုံး",
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
