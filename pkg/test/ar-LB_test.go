package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestArLBNumbers(t *testing.T) {
	tests := []struct {
		input    float64
		expected string
	}{
		{0, "صفر"},
		{1, "واحد"},
		{5, "خمسة"},
		{11, "أحد عشر"},
		{12, "اثنا عشر"},
		{15, "خمسة عشر"},
		{21, "عشرون واحد"},
		{30, "ثلاثون"},
		{47, "أربعون سبعة"},
		{100, "مئة"},
		{101, "واحد مئة واحد"},
		{256, "اثنان مئة خمسون ستة"},
		{1000, "ألف"},
		{1001, "واحد ألف واحد"},
		{1000000, "مليون"},
		{1234567, "واحد مليون اثنان مئة ثلاثون أربعة ألف خمسة مئة ستون سبعة"},
	}

	for _, test := range tests {
		result := (&pkg.NumI18NOptions{
			Locale: "ar-LB",
			WordDetails: &pkg.WordDetails{
				Capitalize: true,
			},
		}).ToWords(test.input)

		if result != test.expected {
			t.Errorf("For input %g, expected '%s' but got '%s'", test.input, test.expected, result)
		}
	}
}

func TestArLBCurrency(t *testing.T) {
	tests := []struct {
		input    float64
		expected string
	}{
		{0, "صفر ليرات لبنانية"},
		{1, "واحد ليرة لبنانية"},
		{5, "خمسة ليرات لبنانية"},
		{1000000, "مليون ليرات لبنانية"},
	}

	for _, test := range tests {
		result := (&pkg.NumI18NOptions{
			Locale: "ar-LB",
			WordDetails: &pkg.WordDetails{
				Currency:   true,
				Capitalize: true,
			},
		}).ToWords(test.input)

		if result != test.expected {
			t.Errorf("For input %g, expected '%s' but got '%s'", test.input, test.expected, result)
		}
	}
}

func TestArLBDecimalWithCurrency(t *testing.T) {
	tests := []struct {
		input    float64
		expected string
	}{
		{0.01, "صفر ليرات لبنانية و واحد قرش"},
		{0.99, "صفر ليرات لبنانية و تسعون تسعة قروش"},
		{5.01, "خمسة ليرات لبنانية و واحد قرش"},
		{5.25, "خمسة ليرات لبنانية و عشرون خمسة قروش"},
		{1234.56, "واحد ألف اثنان مئة ثلاثون أربعة ليرات لبنانية و خمسون ستة قروش"},
	}

	for _, test := range tests {
		result := (&pkg.NumI18NOptions{
			Locale: "ar-LB",
			WordDetails: &pkg.WordDetails{
				Currency:   true,
				Decimal:    true,
				Capitalize: true,
			},
		}).ToWords(test.input)

		if result != test.expected {
			t.Errorf("For input %g, expected '%s' but got '%s'", test.input, test.expected, result)
		}
	}
}

func TestArLBDecimalWithoutCurrency(t *testing.T) {
	result := (&pkg.NumI18NOptions{
		Locale: "ar-LB",
		WordDetails: &pkg.WordDetails{
			Decimal:    true,
			Capitalize: true,
		},
	}).ToWords(123.45)

	expected := "واحد مئة عشرون ثلاثة و أربعون خمسة"
	if result != expected {
		t.Errorf("For decimal without currency, expected '%s' but got '%s'", expected, result)
	}
}

func TestArLBNegativeNumbers(t *testing.T) {
	result := (&pkg.NumI18NOptions{
		Locale: "ar-LB",
		WordDetails: &pkg.WordDetails{
			NegativeWord: true,
			Capitalize:   true,
		},
	}).ToWords(-50)

	expected := "ناقص خمسون"
	if result != expected {
		t.Errorf("For negative number, expected '%s' but got '%s'", expected, result)
	}
}

func TestArLBNegativeCurrencyWithDecimal(t *testing.T) {
	result := (&pkg.NumI18NOptions{
		Locale: "ar-LB",
		WordDetails: &pkg.WordDetails{
			Currency:     true,
			Decimal:      true,
			NegativeWord: true,
			Capitalize:   true,
		},
	}).ToWords(-25.75)

	expected := "ناقص عشرون خمسة ليرات لبنانية و سبعون خمسة قروش"
	if result != expected {
		t.Errorf("For negative currency with decimal, expected '%s' but got '%s'", expected, result)
	}
}

func TestArLBNumbersWithoutCapitalization(t *testing.T) {
	tests := []struct {
		input    float64
		expected string
	}{
		{0, "صفر"},
		{1, "واحد"},
		{5, "خمسة"},
		{100, "مئة"},
	}

	for _, test := range tests {
		result := (&pkg.NumI18NOptions{
			Locale: "ar-LB",
			WordDetails: &pkg.WordDetails{
				Capitalize: false,
			},
		}).ToWords(test.input)

		if result != test.expected {
			t.Errorf("For input %g without capitalization, expected '%s' but got '%s'", test.input, test.expected, result)
		}
	}
}

func TestArLBCurrencyWithoutCapitalization(t *testing.T) {
	result := (&pkg.NumI18NOptions{
		Locale: "ar-LB",
		WordDetails: &pkg.WordDetails{
			Currency:   true,
			Capitalize: false,
		},
	}).ToWords(1)

	expected := "واحد ليرة لبنانية"
	if result != expected {
		t.Errorf("For currency without capitalization, expected '%s' but got '%s'", expected, result)
	}
}

func TestArLBCustomCurrency(t *testing.T) {
	result := (&pkg.NumI18NOptions{
		Locale: "ar-LB",
		WordDetails: &pkg.WordDetails{
			Currency:   true,
			Capitalize: true,
			OverrideOptions: &pkg.OverrideOptions{
				Name:   "دولار",
				Plural: "دولارات",
			},
		},
	}).ToWords(5)

	expected := "خمسة دولارات"
	if result != expected {
		t.Errorf("For custom currency, expected '%s' but got '%s'", expected, result)
	}
}

func TestArLBCustomCurrencyWithDecimal(t *testing.T) {
	result := (&pkg.NumI18NOptions{
		Locale: "ar-LB",
		WordDetails: &pkg.WordDetails{
			Currency:   true,
			Decimal:    true,
			Capitalize: true,
			OverrideOptions: &pkg.OverrideOptions{
				Name:             "دولار",
				Plural:           "دولارات",
				FractionUnitName: "سنت",
				FractionPlural:   "سنت",
			},
		},
	}).ToWords(5.25)

	expected := "خمسة دولارات و عشرون خمسة سنت"
	if result != expected {
		t.Errorf("For custom currency with decimal, expected '%s' but got '%s'", expected, result)
	}
}

func TestArLBFormattingEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:  "Zero with currency",
			input: 0,
			options: &pkg.NumI18NOptions{
				Locale: "ar-LB",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "صفر ليرات لبنانية",
		},
		{
			name:  "One cent",
			input: 0.01,
			options: &pkg.NumI18NOptions{
				Locale: "ar-LB",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "صفر ليرات لبنانية و واحد قرش",
		},
		{
			name:  "Large number",
			input: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "ar-LB",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "مليون",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.options.ToWords(test.input)
			if result != test.expected {
				t.Errorf("For test '%s', expected '%s' but got '%s'", test.name, test.expected, result)
			}
		})
	}
}

func TestArLBSpecialNumbers(t *testing.T) {
	tests := []struct {
		input    float64
		expected string
	}{
		{11, "أحد عشر"},
		{12, "اثنا عشر"},
		{13, "ثلاثة عشر"},
		{19, "تسعة عشر"},
		{20, "عشرون"},
		{21, "عشرون واحد"},
		{99, "تسعون تسعة"},
		{101, "واحد مئة واحد"},
		{1001, "واحد ألف واحد"},
	}

	for _, test := range tests {
		result := (&pkg.NumI18NOptions{
			Locale: "ar-LB",
			WordDetails: &pkg.WordDetails{
				Capitalize: true,
			},
		}).ToWords(test.input)

		if result != test.expected {
			t.Errorf("For special number %g, expected '%s' but got '%s'", test.input, test.expected, result)
		}
	}
}

func TestArLBLargeNumbers(t *testing.T) {
	tests := []struct {
		input    float64
		expected string
	}{
		{1000, "ألف"},
		{2000, "اثنان ألف"},
		{10000, "عشرة ألف"},
		{100000, "واحد مئة ألف"},
		{1000000, "مليون"},
	}

	for _, test := range tests {
		result := (&pkg.NumI18NOptions{
			Locale: "ar-LB",
			WordDetails: &pkg.WordDetails{
				Capitalize: true,
			},
		}).ToWords(test.input)

		if result != test.expected {
			t.Errorf("For large number %g, expected '%s' but got '%s'", test.input, test.expected, result)
		}
	}
}

func TestArLBComprehensiveCurrency(t *testing.T) {
	tests := []struct {
		input    float64
		expected string
	}{
		{1, "واحد ليرة لبنانية"},
		{2, "اثنان ليرات لبنانية"},
		{5, "خمسة ليرات لبنانية"},
		{10, "عشرة ليرات لبنانية"},
		{100, "مئة ليرات لبنانية"},
	}

	for _, test := range tests {
		result := (&pkg.NumI18NOptions{
			Locale: "ar-LB",
			WordDetails: &pkg.WordDetails{
				Currency:   true,
				Capitalize: true,
			},
		}).ToWords(test.input)

		if result != test.expected {
			t.Errorf("For currency amount %g, expected '%s' but got '%s'", test.input, test.expected, result)
		}
	}
}

func TestArLBNegativeFormatting(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:  "Negative basic",
			input: -100,
			options: &pkg.NumI18NOptions{
				Locale: "ar-LB",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ناقص مئة",
		},
		{
			name:  "Negative currency",
			input: -50,
			options: &pkg.NumI18NOptions{
				Locale: "ar-LB",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "ناقص خمسون ليرات لبنانية",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.options.ToWords(test.input)
			if result != test.expected {
				t.Errorf("For test '%s', expected '%s' but got '%s'", test.name, test.expected, result)
			}
		})
	}
}
