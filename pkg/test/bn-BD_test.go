package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_BengaliBangladesh_Numbers(t *testing.T) {
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
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "শূন্য",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "পাঁচ",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "পনের",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ত্রিশ",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "চল্লিশ সাত",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "একশত",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "দুই শত পঞ্চাশ ছয়",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "হাজার",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "মিলিয়ন",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "এক মিলিয়ন দুই শত ত্রিশ চার হাজার পাঁচ শত ষাট সাত",
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

func TestToWords_BengaliBangladesh_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "One taka",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "এক টাকা",
		},
		{
			name:   "Multiple taka",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "পাঁচ টাকা",
		},
		{
			name:   "Zero taka",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "শূন্য টাকা",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "মিলিয়ন টাকা",
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

func TestToWords_BengaliBangladesh_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Taka and one paisa",
			amount: 5.01,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "পাঁচ টাকা এবং এক পয়সা",
		},
		{
			name:   "Taka and multiple paisa",
			amount: 5.25,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "পাঁচ টাকা এবং বিশ পাঁচ পয়সা",
		},
		{
			name:   "Only paisa",
			amount: 0.99,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "শূন্য টাকা এবং নব্বই নয় পয়সা",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "এক হাজার দুই শত ত্রিশ চার টাকা এবং পঞ্চাশ ছয় পয়সা",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "এক শত বিশ তিন এবং চল্লিশ পাঁচ",
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

func TestToWords_BengaliBangladesh_Negative(t *testing.T) {
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
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "মাইনাস পঞ্চাশ",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "মাইনাস বিশ পাঁচ টাকা এবং সত্তর পাঁচ পয়সা",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &pkg.OverrideOptions{
						NegativeWord: "ঋণাত্মক",
					},
				},
			},
			expected: "ঋণাত্মক একশত টাকা",
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

func TestToWords_BengaliBangladesh_Formatting(t *testing.T) {
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
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "এক শত বিশ তিন টাকা",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "এক শত বিশ তিন টাকা",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "নয় শত নব্বই নয় শুধুমাত্র",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "পাঁচ শত টাকা শুধুমাত্র",
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

func TestToWords_BengaliBangladesh_CustomCurrency(t *testing.T) {
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
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "রুপি",
						Plural: "রুপি",
					},
				},
			},
			expected: "একশত রুপি",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:             "ইউরো",
						Plural:           "ইউরো",
						FractionUnitName: "সেন্ট",
						FractionPlural:   "সেন্ট",
					},
				},
			},
			expected: "পঞ্চাশ ইউরো এবং বিশ পাঁচ সেন্ট",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &pkg.OverrideOptions{
						Name:   "পাউন্ড",
						Plural: "পাউন্ড",
					},
				},
			},
			expected: "এক পাউন্ড",
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

func TestToWords_BengaliBangladesh_EdgeCases(t *testing.T) {
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
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "শূন্য টাকা এবং এক পয়সা",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "এঘার",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "বার",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "বিশ এক",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "এক শত এক",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &pkg.NumI18NOptions{
				Locale: "bn-BD",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "এক হাজার এক",
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
