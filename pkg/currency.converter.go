package pkg

// ToFormat converts a number to localized formatted string
// Example: 1234567.89 -> "1,234,567.89" (US) or "1.234.567,89" (DE)
func (op *NumI18NOptions) ToFormat(amount float64) string {
	return ""
}
