package pkg

// WordToNum converts written numbers back to their numeric representation
// Example: "one hundred twenty-three" -> 123.0, "forty-five" -> 45.0, "first" -> 1.0
func (op *NumI18NOptions) WordToNum(amount string) float64 {
	return 0
}

// FormatToNum converts localized formatted number strings back to their numeric representation
// Example: "1,234,567.89" -> 1234567.89 (US), "1.234.567,89" -> 1234567.89 (DE), "1 234 567,89" -> 1234567.89 (FR)
func (op *NumI18NOptions) FormatToNum(amount string) float64 {
	return 0
}
