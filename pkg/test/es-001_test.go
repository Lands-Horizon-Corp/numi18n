package test

import (
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func TestToWords_SpanishWorld_Numbers(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Cero",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "es-001",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Cero",
		},
		{
			name:   "Un dígito",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "es-001",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Cinco",
		},
		{
			name:   "Adolescentes",
			amount: 15,
			options: &pkg.NumI18NOptions{
				Locale: "es-001",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Quince",
		},
		{
			name:   "Decenas",
			amount: 30,
			options: &pkg.NumI18NOptions{
				Locale: "es-001",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Treinta",
		},
		{
			name:   "Número compuesto",
			amount: 47,
			options: &pkg.NumI18NOptions{
				Locale: "es-001",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Cuarenta Siete",
		},
		{
			name:   "Cien exacto",
			amount: 100,
			options: &pkg.NumI18NOptions{
				Locale: "es-001",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Cien",
		},
		{
			name:   "Centenas con resto",
			amount: 256,
			options: &pkg.NumI18NOptions{
				Locale: "es-001",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Dos Ciento Cincuenta Seis",
		},
		{
			name:   "Mil",
			amount: 1000,
			options: &pkg.NumI18NOptions{
				Locale: "es-001",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Mil",
		},
		{
			name:   "Un millón",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "es-001",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Un Millón",
		},
		{
			name:   "Número complejo grande",
			amount: 12345,
			options: &pkg.NumI18NOptions{
				Locale: "es-001",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Doce Mil Tres Ciento Cuarenta Cinco",
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

func TestToWords_SpanishWorld_Currency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Un euro",
			amount: 1,
			options: &pkg.NumI18NOptions{
				Locale: "es-001",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Uno Euro",
		},
		{
			name:   "Múltiples euros",
			amount: 5,
			options: &pkg.NumI18NOptions{
				Locale: "es-001",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Cinco Euros",
		},
		{
			name:   "Cero euros",
			amount: 0,
			options: &pkg.NumI18NOptions{
				Locale: "es-001",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Cero Euros",
		},
		{
			name:   "Cantidad grande",
			amount: 1000000,
			options: &pkg.NumI18NOptions{
				Locale: "es-001",
				WordDetails: &pkg.WordDetails{
					Currency:   true,
					Capitalize: true,
				},
			},
			expected: "Un Millón Euros",
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

func TestToWords_SpanishWorld_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		options  *pkg.NumI18NOptions
		expected string
	}{
		{
			name:   "Once (caso especial)",
			amount: 11,
			options: &pkg.NumI18NOptions{
				Locale: "es-001",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Once",
		},
		{
			name:   "Doce (caso especial)",
			amount: 12,
			options: &pkg.NumI18NOptions{
				Locale: "es-001",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Doce",
		},
		{
			name:   "Veintiuno",
			amount: 21,
			options: &pkg.NumI18NOptions{
				Locale: "es-001",
				WordDetails: &pkg.WordDetails{
					Capitalize: true,
				},
			},
			expected: "Veintiuno",
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
