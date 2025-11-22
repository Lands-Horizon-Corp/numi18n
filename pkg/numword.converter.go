package pkg

// ToWords converts a number to its word representation
// Example: 123.45 -> "one hundred twenty-three and forty-five hundredths"
func (op *NumI18NOptions) ToWords(amount float64) string {
	return ""
}

// ToOrdinal converts a number to its ordinal representation
// Example: 123 -> "123rd", 1 -> "1st", 2 -> "2nd"
func (op *NumI18NOptions) ToOrdinal(amount float64) string {
	return ""
}

// ToOrdinalWords converts a number to its ordinal word representation
// Example: 123 -> "one hundred twenty-third", 1 -> "first", 2 -> "second"
func (op *NumI18NOptions) ToOrdinalWords(amount float64) string {
	return ""
}

// ToFormat converts a number to localized formatted string
// Example: 1234567.89 -> "1,234,567.89" (US) or "1.234.567,89" (DE)
func (op *NumI18NOptions) ToFormat(amount float64) string {
	return ""
}
