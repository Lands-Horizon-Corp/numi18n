package formatter

import "github.com/Lands-Horizon-Corp/numi18n/pkg/locale"

// convertToWordsJapanese converts number to words with Japanese-specific formatting (no spaces)
func convertToWordsJapanese(number int64, targetLocale locale.NumI18NLocale) string {
	if number == 0 {
		return GetWordForNumber(0, targetLocale)
	}

	result := ""

	// Handle large numbers first
	if number >= 100000000 { // 億 (oku)
		oku := number / 100000000
		if oku > 0 {
			result += convertToWordsJapanese(oku, targetLocale) + GetWordForNumber(100000000, targetLocale)
		}
		number %= 100000000
	}

	if number >= 10000 { // 万 (man)
		man := number / 10000
		if man > 0 {
			if man == 1 {
				result += GetWordForNumber(10000, targetLocale) // Just "万" for 10000
			} else {
				result += convertToWordsJapanese(man, targetLocale) + GetWordForNumber(10000, targetLocale)
			}
		}
		number %= 10000
	}

	if number >= 1000 {
		thousands := number / 1000
		if thousands > 0 {
			if thousands == 1 {
				result += GetWordForNumber(1000, targetLocale) // Just "千" for 1000
			} else {
				result += convertToWordsJapanese(thousands, targetLocale) + GetWordForNumber(1000, targetLocale)
			}
		}
		number %= 1000
	}

	if number >= 100 {
		hundreds := number / 100
		if hundreds > 0 {
			if hundreds == 1 {
				result += GetWordForNumber(100, targetLocale) // Just "百" for 100
			} else {
				result += convertToWordsJapanese(hundreds, targetLocale) + GetWordForNumber(100, targetLocale)
			}
		}
		number %= 100
	}

	if number >= 10 {
		tens := number / 10
		if tens > 0 {
			if tens == 1 {
				result += GetWordForNumber(10, targetLocale) // Just "十" for 10
			} else {
				result += convertToWordsJapanese(tens, targetLocale) + GetWordForNumber(10, targetLocale)
			}
		}
		number %= 10
	}

	if number > 0 {
		result += GetWordForNumber(number, targetLocale)
	}

	return result
}

// JapaneseFormatter handles Japanese-specific formatting
type JapaneseFormatter struct{}

func (f *JapaneseFormatter) FormatNumber(number int64, targetLocale locale.NumI18NLocale) string {
	return convertToWordsJapanese(number, targetLocale)
}

func (f *JapaneseFormatter) FormatCurrency(result string, wholePart int64, currencyName, currencyPlural string) string {
	if wholePart == 1 {
		return result + currencyName
	}
	return result + currencyPlural
}

func (f *JapaneseFormatter) FormatFractional(result, fractionalWords string, andText string) string {
	return result + andText + fractionalWords
}

func (f *JapaneseFormatter) FormatFractionalCurrency(result string, fractionalValue int64, fractionName, fractionPlural string) string {
	if fractionalValue == 1 {
		return result + fractionName
	}
	return result + fractionPlural
}

func (f *JapaneseFormatter) FormatNegative(result, negativeWord string) string {
	return negativeWord + result
}
