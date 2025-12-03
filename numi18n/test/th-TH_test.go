package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToWords_ThaiThailand_Numbers(t *testing.T) {
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
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ศูนย์",
		},
		{
			name:   "Single digit",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ห้า",
		},
		{
			name:   "Teens",
			amount: 15,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "สิบห้า",
		},
		{
			name:   "Tens",
			amount: 30,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "สามสิบ",
		},
		{
			name:   "Compound number",
			amount: 47,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "สี่สิบเจ็ด",
		},
		{
			name:   "One hundred (exact mapping)",
			amount: 100,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "หนึ่งร้อย",
		},
		{
			name:   "Hundreds with remainder",
			amount: 256,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "สอง หนึ่งร้อย ห้าสิบ หก",
		},
		{
			name:   "One thousand",
			amount: 1000,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "หนึ่งพัน",
		},
		{
			name:   "One million (exact mapping)",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "หนึ่งล้าน",
		},
		{
			name:   "Large complex number",
			amount: 1234567,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "หนึ่ง หนึ่งล้าน สอง หนึ่งร้อย สามสิบ สี่ หนึ่งพัน ห้า หนึ่งร้อย หกสิบ เจ็ด",
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

func TestToWords_ThaiThailand_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "One baht",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "หนึ่งบาท",
		},
		{
			name:   "Multiple bahts",
			amount: 5,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ห้าบาท",
		},
		{
			name:   "Zero bahts",
			amount: 0,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "ศูนย์บาท",
		},
		{
			name:   "Large amount",
			amount: 1000000,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "หนึ่งล้านบาท",
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

func TestToWords_ThaiThailand_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *numi18n.NumI18NOptions
		expected string
	}{
		{
			name:   "Bahts and one satang",
			amount: 5.01,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ห้าบาทและหนึ่งสตางค์",
		},
		{
			name:   "Bahts and multiple satangs",
			amount: 5.25,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ห้าบาทและยี่สิบห้าสตางค์",
		},
		{
			name:   "Only satangs",
			amount: 0.99,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ศูนย์บาทและเก้าสิบเก้าสตางค์",
		},
		{
			name:   "Complex amount",
			amount: 1234.56,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "หนึ่ง หนึ่งพัน สอง หนึ่งร้อย สามสิบ สี่บาทและห้าสิบหกสตางค์",
		},
		{
			name:   "Decimal without currency",
			amount: 123.45,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "หนึ่ง หนึ่งร้อย ยี่สิบ สามและสี่สิบห้า",
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

func TestToWords_ThaiThailand_Negative(t *testing.T) {
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
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ลบห้าสิบ",
		},
		{
			name:   "Negative currency",
			amount: -25.75,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ลบยี่สิบห้าบาทและเจ็ดสิบห้าสตางค์",
		},
		{
			name:   "Negative with custom word",
			amount: -100,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
					OverrideOptions: &numi18n.OverrideOptions{
						NegativeWord: "ติดลบ",
					},
				},
			},
			expected: "ติดลบหนึ่งร้อยบาท",
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

func TestToWords_ThaiThailand_Formatting(t *testing.T) {
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
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Uppercase: true,
				},
			},
			expected: "หนึ่ง หนึ่งร้อย ยี่สิบ สามบาท",
		},
		{
			name:   "Lowercase",
			amount: 123,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Currency:  true,
					Lowercase: true,
				},
			},
			expected: "หนึ่ง หนึ่งร้อย ยี่สิบ สามบาท",
		},
		{
			name:   "Only flag",
			amount: 999,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "เก้า หนึ่งร้อย เก้าสิบ เก้า เท่านั้น",
		},
		{
			name:   "Only flag with currency",
			amount: 500,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Only:       true,
					Capitalize: true,
				},
			},
			expected: "ห้าร้อยบาท เท่านั้น",
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

func TestToWords_ThaiThailand_CustomCurrency(t *testing.T) {
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
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "เปโซ",
						Plural: "เปโซ",
					},
				},
			},
			expected: "หนึ่งร้อยเปโซ",
		},
		{
			name:   "Custom currency with decimals",
			amount: 50.25,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:             "ยูโร",
						Plural:           "ยูโร",
						FractionUnitName: "เซ็นต์",
						FractionPlural:   "เซ็นต์",
					},
				},
			},
			expected: "ห้าสิบยูโรและยี่สิบห้าเซ็นต์",
		},
		{
			name:   "Single custom currency",
			amount: 1,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Capitalize: true,
					OverrideOptions: &numi18n.OverrideOptions{
						Name:   "ปอนด์",
						Plural: "ปอนด์",
					},
				},
			},
			expected: "หนึ่งปอนด์",
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

func TestToWords_ThaiThailand_EdgeCases(t *testing.T) {
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
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "ศูนย์บาทและหนึ่งสตางค์",
		},
		{
			name:   "Eleven (special case)",
			amount: 11,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "สิบเอ็ด",
		},
		{
			name:   "Twelve (special case)",
			amount: 12,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "สิบสอง",
		},
		{
			name:   "Twenty one",
			amount: 21,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "ยี่สิบเอ็ด",
		},
		{
			name:   "One hundred one",
			amount: 101,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "หนึ่ง หนึ่งร้อย หนึ่ง",
		},
		{
			name:   "One thousand one",
			amount: 1001,
			options: &numi18n.NumI18NOptions{
				Locale: "th-TH",
				WordDetails: &numi18n.WordDetails{
					Capitalize: true,
				},
			},
			expected: "หนึ่ง หนึ่งพัน หนึ่ง",
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
