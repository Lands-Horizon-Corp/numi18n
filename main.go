package main

import (
	"fmt"

	"github.com/Lands-Horizon-Corp/numi18n/pkg/locale"
)

func main() {
	fmt.Println("=== Ordinal Mapping System Demo ===")

	// Test English ordinals
	fmt.Println("\n🇺🇸 English Ordinals (en-US):")
	testOrdinals(locale.USLocale, []int64{1, 2, 3, 21, 22, 23, 101, 102, 103})

	// Test French ordinals
	fmt.Println("\n🇫🇷 French Ordinals (fr-001):")
	testOrdinals(locale.FR001Locale, []int64{1, 2, 3, 4, 5})

	// Test Arabic ordinals
	fmt.Println("\n🇸🇦 Arabic Ordinals (ar-SA):")
	testOrdinals(locale.ARSALocale, []int64{1, 2, 3, 4, 5})
}

func testOrdinals(loc locale.NumI18NLocale, numbers []int64) {
	for _, num := range numbers {
		word := loc.FormatOrdinalWord(num)
		numeric := loc.FormatOrdinalNumeric(num)
		fmt.Printf("  %d -> Word: \"%s\" | Numeric: \"%s\"\n", num, word, numeric)
	}
}
