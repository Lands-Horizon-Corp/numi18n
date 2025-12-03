package main

import (
	"fmt"

	"github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func main() {
	fmt.Println("=== Number to Words and Formatting Sample ===")
	fmt.Println()

	// Sample numbers to demonstrate conversion
	sampleSimpleNumber := 42.0
	sampleLargeNumber := 1234.56
	sampleCurrencyAmount := 999.99
	sampleMillionAmount := 1000000.00

	// === English (US) Examples ===
	fmt.Println("🇺🇸 English (United States) - en-US:")
	fmt.Println("-----------------------------------")

	// Convert simple number to words (no currency, no decimals)
	englishSimpleWordOptions := &pkg.NumI18NOptions{
		Locale: "en-US",
		WordDetails: &pkg.WordDetails{
			Capitalize: true,
		},
	}
	englishSimpleNumberInWords := englishSimpleWordOptions.ToWords(sampleSimpleNumber)
	fmt.Printf("Simple number (%g): %s\n", sampleSimpleNumber, englishSimpleNumberInWords)

	// Convert large number with decimals to words
	englishDecimalWordOptions := &pkg.NumI18NOptions{
		Locale: "en-US",
		WordDetails: &pkg.WordDetails{
			Capitalize: true,
			Decimal:    true, // Include decimal part in words
		},
	}
	englishLargeNumberInWords := englishDecimalWordOptions.ToWords(sampleLargeNumber)
	fmt.Printf("Large number with decimals (%g): %s\n", sampleLargeNumber, englishLargeNumberInWords)

	// Format numbers with currency (both words and formatted numbers)
	englishCurrencyWordOptions := &pkg.NumI18NOptions{
		Locale: "en-US",
		WordDetails: &pkg.WordDetails{
			Currency:   true, // Include currency in output
			Decimal:    true, // Include decimal part
			Capitalize: true,
		},
	}
	englishCurrencyInWords := englishCurrencyWordOptions.ToWords(sampleCurrencyAmount)
	fmt.Printf("Currency in words ($%.2f): %s\n", sampleCurrencyAmount, englishCurrencyInWords)

	englishCurrencyFormatted := englishCurrencyWordOptions.ToFormat(sampleCurrencyAmount)
	fmt.Printf("Currency formatted ($%.2f): %s\n", sampleCurrencyAmount, englishCurrencyFormatted)

	englishMillionFormatted := englishCurrencyWordOptions.ToFormat(sampleMillionAmount)
	fmt.Printf("Million formatted ($%.2f): %s\n", sampleMillionAmount, englishMillionFormatted)

	fmt.Println()

	// === Japanese Examples ===
	fmt.Println("🇯🇵 Japanese (Japan) - ja-JP:")
	fmt.Println("----------------------------")

	// Convert simple number to words in Japanese
	japaneseSimpleWordOptions := &pkg.NumI18NOptions{
		Locale: "ja-JP",
		WordDetails: &pkg.WordDetails{
			Capitalize: true,
		},
	}
	japaneseSimpleNumberInWords := japaneseSimpleWordOptions.ToWords(sampleSimpleNumber)
	fmt.Printf("Simple number (%g): %s\n", sampleSimpleNumber, japaneseSimpleNumberInWords)

	// Convert large number with decimals to words in Japanese
	japaneseDecimalWordOptions := &pkg.NumI18NOptions{
		Locale: "ja-JP",
		WordDetails: &pkg.WordDetails{
			Capitalize: true,
			Decimal:    true, // Include decimal part in Japanese
		},
	}
	japaneseLargeNumberInWords := japaneseDecimalWordOptions.ToWords(sampleLargeNumber)
	fmt.Printf("Large number with decimals (%g): %s\n", sampleLargeNumber, japaneseLargeNumberInWords)

	// Format numbers with currency in Japanese (Yen)
	japaneseCurrencyWordOptions := &pkg.NumI18NOptions{
		Locale: "ja-JP",
		WordDetails: &pkg.WordDetails{
			Currency:   true, // Include Japanese Yen currency
			Decimal:    true, // Include decimal part (sen)
			Capitalize: true,
		},
	}
	japaneseCurrencyInWords := japaneseCurrencyWordOptions.ToWords(sampleCurrencyAmount)
	fmt.Printf("Currency in words (¥%.2f): %s\n", sampleCurrencyAmount, japaneseCurrencyInWords)

	japaneseCurrencyFormatted := japaneseCurrencyWordOptions.ToFormat(sampleCurrencyAmount)
	fmt.Printf("Currency formatted (¥%.2f): %s\n", sampleCurrencyAmount, japaneseCurrencyFormatted)

	japaneseMillionFormatted := japaneseCurrencyWordOptions.ToFormat(sampleMillionAmount)
	fmt.Printf("Million formatted (¥%.2f): %s\n", sampleMillionAmount, japaneseMillionFormatted)

	fmt.Println()

	// === Philippines Examples ===
	fmt.Println("🇵🇭 Philippines (Filipino) - ph-PH:")
	fmt.Println("----------------------------------")

	// Convert simple number to words in Filipino
	philippinesSimpleWordOptions := &pkg.NumI18NOptions{
		Locale: "ph-PH",
		WordDetails: &pkg.WordDetails{
			Capitalize: true,
		},
	}
	philippinesSimpleNumberInWords := philippinesSimpleWordOptions.ToWords(sampleSimpleNumber)
	fmt.Printf("Simple number (%g): %s\n", sampleSimpleNumber, philippinesSimpleNumberInWords)

	// Convert large number with decimals to words in Filipino
	philippinesDecimalWordOptions := &pkg.NumI18NOptions{
		Locale: "ph-PH",
		WordDetails: &pkg.WordDetails{
			Capitalize: true,
			Decimal:    true, // Include decimal part in Filipino
		},
	}
	philippinesLargeNumberInWords := philippinesDecimalWordOptions.ToWords(sampleLargeNumber)
	fmt.Printf("Large number with decimals (%g): %s\n", sampleLargeNumber, philippinesLargeNumberInWords)

	// Format numbers with currency in Filipino (Peso)
	philippinesCurrencyWordOptions := &pkg.NumI18NOptions{
		Locale: "ph-PH",
		WordDetails: &pkg.WordDetails{
			Currency:   true, // Include Philippine Peso currency
			Decimal:    true, // Include decimal part (sentimo)
			Capitalize: true,
		},
	}
	philippinesCurrencyInWords := philippinesCurrencyWordOptions.ToWords(sampleCurrencyAmount)
	fmt.Printf("Currency in words (₱%.2f): %s\n", sampleCurrencyAmount, philippinesCurrencyInWords)

	philippinesCurrencyFormatted := philippinesCurrencyWordOptions.ToFormat(sampleCurrencyAmount)
	fmt.Printf("Currency formatted (₱%.2f): %s\n", sampleCurrencyAmount, philippinesCurrencyFormatted)

	philippinesMillionFormatted := philippinesCurrencyWordOptions.ToFormat(sampleMillionAmount)
	fmt.Printf("Million formatted (₱%.2f): %s\n", sampleMillionAmount, philippinesMillionFormatted)

	fmt.Println()

	// === Additional Examples with Different Formatting ===
	fmt.Println("💰 Special Formatting Examples:")
	fmt.Println("------------------------------")

	// Test number for comparison across locales
	testNumberForComparison := 123.45

	// English: Number vs Currency formatting comparison
	englishNumberFormatOptions := &pkg.NumI18NOptions{
		Locale: "en-US",
		WordDetails: &pkg.WordDetails{
			Decimal:    true, // Include decimals but no currency
			Capitalize: true,
		},
	}
	englishNumberFormattedText := englishNumberFormatOptions.ToFormat(testNumberForComparison)
	englishNumberInPlainWords := englishNumberFormatOptions.ToWords(testNumberForComparison)

	englishCurrencyFormattedText := englishCurrencyWordOptions.ToFormat(testNumberForComparison)
	englishCurrencyInPlainWords := englishCurrencyWordOptions.ToWords(testNumberForComparison)

	fmt.Printf("English - Number format: %s\n", englishNumberFormattedText)
	fmt.Printf("English - Number words: %s\n", englishNumberInPlainWords)
	fmt.Printf("English - Currency format: %s\n", englishCurrencyFormattedText)
	fmt.Printf("English - Currency words: %s\n", englishCurrencyInPlainWords)

	// Japanese: Number vs Currency formatting comparison
	japaneseNumberFormatOptions := &pkg.NumI18NOptions{
		Locale: "ja-JP",
		WordDetails: &pkg.WordDetails{
			Decimal:    true, // Include decimals but no currency
			Capitalize: true,
		},
	}
	japaneseNumberFormattedText := japaneseNumberFormatOptions.ToFormat(testNumberForComparison)
	japaneseNumberInPlainWords := japaneseNumberFormatOptions.ToWords(testNumberForComparison)

	japaneseCurrencyFormattedText := japaneseCurrencyWordOptions.ToFormat(testNumberForComparison)
	japaneseCurrencyInPlainWords := japaneseCurrencyWordOptions.ToWords(testNumberForComparison)

	fmt.Printf("Japanese - Number format: %s\n", japaneseNumberFormattedText)
	fmt.Printf("Japanese - Number words: %s\n", japaneseNumberInPlainWords)
	fmt.Printf("Japanese - Currency format: %s\n", japaneseCurrencyFormattedText)
	fmt.Printf("Japanese - Currency words: %s\n", japaneseCurrencyInPlainWords)

	// Filipino: Number vs Currency formatting comparison
	filipinoNumberFormatOptions := &pkg.NumI18NOptions{
		Locale: "ph-PH",
		WordDetails: &pkg.WordDetails{
			Decimal:    true, // Include decimals but no currency
			Capitalize: true,
		},
	}
	filipinoNumberFormattedText := filipinoNumberFormatOptions.ToFormat(testNumberForComparison)
	filipinoNumberInPlainWords := filipinoNumberFormatOptions.ToWords(testNumberForComparison)

	filipinoCurrencyFormattedText := philippinesCurrencyWordOptions.ToFormat(testNumberForComparison)
	filipinoCurrencyInPlainWords := philippinesCurrencyWordOptions.ToWords(testNumberForComparison)

	fmt.Printf("Filipino - Number format: %s\n", filipinoNumberFormattedText)
	fmt.Printf("Filipino - Number words: %s\n", filipinoNumberInPlainWords)
	fmt.Printf("Filipino - Currency format: %s\n", filipinoCurrencyFormattedText)
	fmt.Printf("Filipino - Currency words: %s\n", filipinoCurrencyInPlainWords)

	fmt.Println()
	fmt.Println("=== Sample Complete ===")
}
