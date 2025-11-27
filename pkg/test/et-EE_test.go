package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_Estonian_Numbers(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Null",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Null",
		},
		{
			name:   "Üks",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Üks",
		},
		{
			name:   "Viis",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Viis",
		},
		{
			name:   "Kümme",
			amount: 10,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Kümme",
		},
		{
			name:   "Üksteist",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Üksteist",
		},
		{
			name:   "Viisteist",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Viisteist",
		},
		{
			name:   "Kakskümmend",
			amount: 20,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Kakskümmend",
		},
		{
			name:   "Kakskümmend viis",
			amount: 25,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Kakskümmend Viis",
		},
		{
			name:   "Kolmkümmend",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Kolmkümmend",
		},
		{
			name:   "Ükssada",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Üks Sada",
		},
		{
			name:   "Ükssada üksteist",
			amount: 111,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Üks Sada Üksteist",
		},
		{
			name:   "Kakssada",
			amount: 200,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Kaks Sada",
		},
		{
			name:   "Tuhat",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Üks Tuhat",
		},
		{
			name:   "Tuhat ükssada kakskümmend kolm",
			amount: 1123,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Üks Tuhat Üks Sada Kakskümmend Kolm",
		},
		{
			name:   "Miljon",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Üks Miljon",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.options.ToWords(test.amount)
			if result != test.expected {
				t.Errorf("Expected '%s' but got '%s'", test.expected, result)
			}
		})
	}
}

func TestToWords_Estonian_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Null eurot",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Null Eurot",
		},
		{
			name:   "Üks euro",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Üks Euro",
		},
		{
			name:   "Kaks eurot",
			amount: 2,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Kaks Eurot",
		},
		{
			name:   "Viis eurot",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Viis Eurot",
		},
		{
			name:   "Kümme eurot",
			amount: 10,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Kümme Eurot",
		},
		{
			name:   "Ükssada eurot",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Üks Sada Eurot",
		},
		{
			name:   "Tuhat eurot",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Üks Tuhat Eurot",
		},
		{
			name:   "Miljon eurot",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Üks Miljon Eurot",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.options.ToWords(test.amount)
			if result != test.expected {
				t.Errorf("Expected '%s' but got '%s'", test.expected, result)
			}
		})
	}
}

func TestToWords_Estonian_Decimals(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Üks euro viiskümmend senti",
			amount: 1.50,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Üks Euro Ja Viiskümmend Senti",
		},
		{
			name:   "Kaks eurot kakskümmend viis senti",
			amount: 2.25,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Kaks Eurot Ja Kakskümmend Viis Senti",
		},
		{
			name:   "Kümme eurot üheksakümmend üheksa senti",
			amount: 10.99,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Kümme Eurot Ja Üheksakümmend Üheksa Senti",
		},
		{
			name:   "Ükssada eurot üks sent",
			amount: 100.01,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Decimal:    true,
					Capitalize: true,
				},
			},
			expected: "Üks Sada Eurot Ja Üks Sent",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.options.ToWords(test.amount)
			if result != test.expected {
				t.Errorf("Expected '%s' but got '%s'", test.expected, result)
			}
		})
	}
}

func TestToWords_Estonian_Negatives(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Miinus üks",
			amount: -1,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Miinus Üks",
		},
		{
			name:   "Miinus viis eurot",
			amount: -5,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Miinus Viis Eurot",
		},
		{
			name:   "Miinus ükssada eurot",
			amount: -100,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Miinus Üks Sada Eurot",
		},
		{
			name:   "Miinus kaks eurot viiskümmend senti",
			amount: -2.50,
			options: &pkg.NumI18NOptions{
				Locale: "et-EE",
				WordDetails: &pkg.WordDetails{
					Currency:     true,
					Decimal:      true,
					NegativeWord: true,
					Capitalize:   true,
				},
			},
			expected: "Miinus Kaks Eurot Ja Viiskümmend Senti",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.options.ToWords(test.amount)
			if result != test.expected {
				t.Errorf("Expected '%s' but got '%s'", test.expected, result)
			}
		})
	}
}
