package test

import (
	"fmt"
	"testing"

	"github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func TestToRoman(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{
			name:     "Zero",
			input:    0,
			expected: "",
		},
		{
			name:     "Negative number",
			input:    -5,
			expected: "",
		},
		{
			name:     "Single digits",
			input:    1,
			expected: "I",
		},
		{
			name:     "Single digits - Four",
			input:    4,
			expected: "IV",
		},
		{
			name:     "Single digits - Five",
			input:    5,
			expected: "V",
		},
		{
			name:     "Single digits - Nine",
			input:    9,
			expected: "IX",
		},
		{
			name:     "Tens - Ten",
			input:    10,
			expected: "X",
		},
		{
			name:     "Tens - Forty",
			input:    40,
			expected: "XL",
		},
		{
			name:     "Tens - Fifty",
			input:    50,
			expected: "L",
		},
		{
			name:     "Tens - Ninety",
			input:    90,
			expected: "XC",
		},
		{
			name:     "Hundreds - One hundred",
			input:    100,
			expected: "C",
		},
		{
			name:     "Hundreds - Four hundred",
			input:    400,
			expected: "CD",
		},
		{
			name:     "Hundreds - Five hundred",
			input:    500,
			expected: "D",
		},
		{
			name:     "Hundreds - Nine hundred",
			input:    900,
			expected: "CM",
		},
		{
			name:     "Thousands - One thousand",
			input:    1000,
			expected: "M",
		},
		{
			name:     "Complex - Forty two",
			input:    42,
			expected: "XLII",
		},
		{
			name:     "Complex - One hundred twenty three",
			input:    123,
			expected: "CXXIII",
		},
		{
			name:     "Complex - Four hundred ninety four",
			input:    494,
			expected: "CDXCIV",
		},
		{
			name:     "Complex - Year 1994",
			input:    1994,
			expected: "MCMXCIV",
		},
		{
			name:     "Complex - Year 2024",
			input:    2024,
			expected: "MMXXIV",
		},
		{
			name:     "Large - 3999 (max traditional)",
			input:    3999,
			expected: "MMMCMXCIX",
		},
		{
			name:     "Extended - Four thousand",
			input:    4000,
			expected: "M_V",
		},
		{
			name:     "Extended - Five thousand",
			input:    5000,
			expected: "_V",
		},
		{
			name:     "Extended - Ten thousand",
			input:    10000,
			expected: "_X",
		},
		{
			name:     "Extended - One hundred thousand",
			input:    100000,
			expected: "_C",
		},
		{
			name:     "Extended - One million",
			input:    1000000,
			expected: "_M",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numi18n.ToRoman(tt.input)
			if result != tt.expected {
				t.Errorf("ToRoman(%d) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFromRoman(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "Invalid Roman numeral",
			input:    "IIII",
			expected: 0,
		},
		{
			name:     "Invalid Roman numeral - VV",
			input:    "VV",
			expected: 0,
		},
		{
			name:     "Single digits - I",
			input:    "I",
			expected: 1,
		},
		{
			name:     "Single digits - IV",
			input:    "IV",
			expected: 4,
		},
		{
			name:     "Single digits - V",
			input:    "V",
			expected: 5,
		},
		{
			name:     "Single digits - IX",
			input:    "IX",
			expected: 9,
		},
		{
			name:     "Tens - X",
			input:    "X",
			expected: 10,
		},
		{
			name:     "Tens - XL",
			input:    "XL",
			expected: 40,
		},
		{
			name:     "Tens - L",
			input:    "L",
			expected: 50,
		},
		{
			name:     "Tens - XC",
			input:    "XC",
			expected: 90,
		},
		{
			name:     "Hundreds - C",
			input:    "C",
			expected: 100,
		},
		{
			name:     "Hundreds - CD",
			input:    "CD",
			expected: 400,
		},
		{
			name:     "Hundreds - D",
			input:    "D",
			expected: 500,
		},
		{
			name:     "Hundreds - CM",
			input:    "CM",
			expected: 900,
		},
		{
			name:     "Thousands - M",
			input:    "M",
			expected: 1000,
		},
		{
			name:     "Complex - XLII",
			input:    "XLII",
			expected: 42,
		},
		{
			name:     "Complex - CXXIII",
			input:    "CXXIII",
			expected: 123,
		},
		{
			name:     "Complex - CDXCIV",
			input:    "CDXCIV",
			expected: 494,
		},
		{
			name:     "Complex - MCMXCIV",
			input:    "MCMXCIV",
			expected: 1994,
		},
		{
			name:     "Complex - MMXXIV",
			input:    "MMXXIV",
			expected: 2024,
		},
		{
			name:     "Large - MMMCMXCIX",
			input:    "MMMCMXCIX",
			expected: 3999,
		},
		{
			name:     "Case insensitive - lowercase",
			input:    "xlii",
			expected: 42,
		},
		{
			name:     "Case insensitive - mixed",
			input:    "McmXcIv",
			expected: 1994,
		},
		{
			name:     "Extended - M_V",
			input:    "M_V",
			expected: 4000,
		},
		{
			name:     "Extended - _V",
			input:    "_V",
			expected: 5000,
		},
		{
			name:     "Extended - _X",
			input:    "_X",
			expected: 10000,
		},
		{
			name:     "Extended - _C",
			input:    "_C",
			expected: 100000,
		},
		{
			name:     "Extended - _M",
			input:    "_M",
			expected: 1000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numi18n.FromRoman(tt.input)
			if result != tt.expected {
				t.Errorf("FromRoman(%s) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid - I",
			input:    "I",
			expected: true,
		},
		{
			name:     "Valid - IV",
			input:    "IV",
			expected: true,
		},
		{
			name:     "Valid - XLII",
			input:    "XLII",
			expected: true,
		},
		{
			name:     "Valid - MCMXCIV",
			input:    "MCMXCIV",
			expected: true,
		},
		{
			name:     "Valid - MMMCMXCIX",
			input:    "MMMCMXCIX",
			expected: true,
		},
		{
			name:     "Valid - Extended _V",
			input:    "_V",
			expected: true,
		},
		{
			name:     "Valid - Extended _M",
			input:    "_M",
			expected: true,
		},
		{
			name:     "Invalid - IIII",
			input:    "IIII",
			expected: false,
		},
		{
			name:     "Invalid - VV",
			input:    "VV",
			expected: false,
		},
		{
			name:     "Invalid - LL",
			input:    "LL",
			expected: false,
		},
		{
			name:     "Invalid - DD",
			input:    "DD",
			expected: false,
		},
		{
			name:     "Invalid - VX",
			input:    "VX",
			expected: false,
		},
		{
			name:     "Invalid - LC",
			input:    "LC",
			expected: false,
		},
		{
			name:     "Invalid - DM",
			input:    "DM",
			expected: false,
		},
		{
			name:     "Invalid - XM",
			input:    "XM",
			expected: false,
		},
		{
			name:     "Invalid - IM",
			input:    "IM",
			expected: false,
		},
		{
			name:     "Invalid - ABC",
			input:    "ABC",
			expected: false,
		},
		{
			name:     "Invalid - random text",
			input:    "HELLO",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numi18n.IsValid(tt.input)
			if result != tt.expected {
				t.Errorf("IsValid(%s) = %t, want %t", tt.input, result, tt.expected)
			}
		})
	}
}

func TestRomanRoundTrip(t *testing.T) {
	// Test that converting to Roman and back gives the same number
	testNumbers := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 24, 27, 30, 31, 37, 42, 49, 50,
		60, 69, 70, 80, 90, 99, 100,
		101, 111, 123, 200, 300, 400, 500,
		600, 700, 800, 900, 999, 1000,
		1001, 1234, 1500, 1984, 1994, 2000, 2024,
		3000, 3500, 3999, 4000, 5000, 9000, 10000,
		50000, 100000, 500000, 1000000,
	}

	for _, num := range testNumbers {
		t.Run(fmt.Sprintf("Number_%d", num), func(t *testing.T) {
			roman := numi18n.ToRoman(num)
			if roman == "" && num > 0 {
				t.Errorf("ToRoman(%d) returned empty string", num)
				return
			}

			result := numi18n.FromRoman(roman)
			if result != num {
				t.Errorf("Round trip failed for %d: ToRoman(%d) = %s, FromRoman(%s) = %d",
					num, num, roman, roman, result)
			}
		})
	}
}
