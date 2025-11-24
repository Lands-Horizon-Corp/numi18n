package pkg

// ToFormat converts a number to localized formatted string
// Example: 1234567.89 -> "1,234,567.89" (US) or "1.234.567,89" (DE)
func (op *NumI18NOptions) ToFormat(amount float64) string {
	return ""
}

// FormatToNum converts localized formatted number strings back to their numeric representation
// Example: "1,234,567.89" -> 1234567.89 (US), "1.234.567,89" -> 1234567.89 (DE), "1 234 567,89" -> 1234567.89 (FR)
func (op *NumI18NOptions) FormatToNum(amount string) float64 {
	return 0
}
