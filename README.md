# numi18n

<div align="center">
  <img src="https://raw.githubusercontent.com/Lands-Horizon-Corp/numi18n/master/assets/icon.png" alt="numi18n Logo" width="200"/>
</div>

<p align="center">
  <a href="https://pkg.go.dev/github.com/Lands-Horizon-Corp/numi18n"><img src="https://pkg.go.dev/badge/github.com/Lands-Horizon-Corp/numi18n.svg" alt="Go Reference"></a>
  <a href="https://goreportcard.com/report/github.com/Lands-Horizon-Corp/numi18n"><img src="https://goreportcard.com/badge/github.com/Lands-Horizon-Corp/numi18n" alt="Go Report Card"></a>
  <a href="https://github.com/Lands-Horizon-Corp/numi18n/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License"></a>
  <a href="https://github.com/Lands-Horizon-Corp/numi18n/releases"><img src="https://img.shields.io/github/release/Lands-Horizon-Corp/numi18n.svg" alt="Release"></a>
</p>

<p align="center">
  <strong>🌍 Internationalized Number-to-Words Conversion for Go</strong>
</p>

---

## Introduction

Convert numbers (including decimals) into words with multi-locale and currency support. Ideal for invoicing, e-commerce, financial applications, and internationalization. This Go library provides comprehensive support for over 100 locales with accurate linguistic formatting.

## Installation

```bash
go get github.com/Lands-Horizon-Corp/numi18n
```

## Dependencies

This library uses only one external dependency:
- **[github.com/shopspring/decimal](https://github.com/shopspring/decimal)** - For precise decimal arithmetic and formatting

## Usage

### Basic Import

```go
import "github.com/Lands-Horizon-Corp/numi18n/numi18n"
```

### Simple Number to Words

```go
package main

import (
    "fmt"
    "github.com/Lands-Horizon-Corp/numi18n/numi18n"
)

func main() {
    // Simple number conversion
    options := &numi18n.NumI18NOptions{
        Locale: "en-US",
        WordDetails: &numi18n.WordDetails{
            Capitalize: true,
        },
    }
    
    result := options.ToWords(123)
    fmt.Println(result) // "One Hundred Twenty Three"
    
    // Number with decimals
    decimalOptions := &numi18n.NumI18NOptions{
        Locale: "en-US", 
        WordDetails: &numi18n.WordDetails{
            Decimal:    true,
            Capitalize: true,
        },
    }
    
    result = decimalOptions.ToWords(123.45)
    fmt.Println(result) // "One Hundred Twenty Three And Forty Five"
}
```

### Currency Conversion

```go
// Convert to currency format
currencyOptions := &numi18n.NumI18NOptions{
    Locale: "en-US",
    WordDetails: &numi18n.WordDetails{
        Currency:   true,
        Decimal:    true, 
        Capitalize: true,
    },
}

words := currencyOptions.ToWords(452.36)
fmt.Println(words) // "Four Hundred Fifty Two Dollars And Thirty Six Cents"

formatted := currencyOptions.ToFormat(452.36)
fmt.Println(formatted) // "$452.36"
```

### Multi-Locale Examples

```go
// Japanese
japaneseOptions := &numi18n.NumI18NOptions{
    Locale: "ja-JP",
    WordDetails: &numi18n.WordDetails{
        Currency:   true,
        Decimal:    true,
        Capitalize: true,
    },
}

words := japaneseOptions.ToWords(1234.56)
fmt.Println(words) // "千二百三十四円と五十六銭"

// Spanish (Mexico)
spanishOptions := &numi18n.NumI18NOptions{
    Locale: "es-MX", 
    WordDetails: &numi18n.WordDetails{
        Currency:   true,
        Decimal:    true,
        Capitalize: true,
    },
}

words = spanishOptions.ToWords(999.99)
fmt.Println(words) // "Novecientos Noventa y Nueve Pesos con Noventa y Nueve Centavos"
```

### Custom Currency Override

```go
customOptions := &numi18n.NumI18NOptions{
    Locale: "en-US",
    WordDetails: &numi18n.WordDetails{
        Currency:   true,
        Decimal:    true,
        Capitalize: true,
        OverrideOptions: &numi18n.OverrideOptions{
            Name:             "Bitcoin",
            Plural:           "Bitcoins", 
            Symbol:           "₿",
            FractionUnitName: "Satoshi",
            FractionPlural:   "Satoshis",
            FractionSymbol:   "₿",
            NegativeWord:     "minus",
        },
    },
}

result := customOptions.ToWords(1.50)
fmt.Println(result) // "One Bitcoin And Fifty Satoshis"
```

### Roman Numerals

```go
// Convert to Roman numerals
roman := numi18n.ToRoman(1994)
fmt.Println(roman) // "MCMXCIV"

// Convert from Roman numerals
number := numi18n.FromRoman("MCMXCIV") 
fmt.Println(number) // 1994

// Validate Roman numerals
isValid := numi18n.IsValid("MCMXCIV")
fmt.Println(isValid) // true

isValid = numi18n.IsValid("IIII")
fmt.Println(isValid) // false
```

### Get Available Locales

```go
// Usage example
availableLocales := numi18n.Locales()
fmt.Printf("Total locales available: %d\n", len(availableLocales))

// Iterate through available locales
for _, locale := range availableLocales {
    fmt.Printf("Locale: %s - %s\n", locale.NumI18Identifier.Locale, locale.NumI18Identifier.CountryName)
}
```

## Configuration Options

### NumI18NOptions

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `CountryName` | `string` | `""` | Full country name, e.g., "Philippines" |
| `Currency` | `string` | `""` | Currency code, e.g., "PHP", "USD" |
| `ISO3166Alpha2` | `string` | `""` | ISO 3166-1 alpha-2 code, e.g., "PH" |
| `ISO3166Alpha3` | `string` | `""` | ISO 3166-1 alpha-3 code, e.g., "PHL" |
| `ISO3166Numeric` | `string` | `""` | ISO 3166 numeric code, e.g., "608" |
| `Locale` | `string` | `"en-US"` | Locale string, e.g., "en-PH", "fr-FR" |
| `Timezone` | `string` | `""` | Timezone, e.g., "Asia/Manila" |
| `Language` | `string` | `""` | Language code, e.g., "en", "fr", "ja" |
| `WordDetails` | `*WordDetails` | `nil` | Detailed configuration for word conversion |
| `OrdinalDetails` | `*OrdinalDetails` | `nil` | Configuration for ordinal number conversion |

### WordDetails Options

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `Currency` | `bool` | `false` | Include currency words, e.g., "one hundred pesos" |
| `Decimal` | `bool` | `false` | Include decimals/fractions in words, e.g., "fifty-six centavos" |
| `Only` | `bool` | `false` | Output only words without numbers or symbols |
| `Capitalize` | `bool` | `false` | Capitalize the first letter of the resulting words |
| `Uppercase` | `bool` | `false` | Convert entire output to uppercase |
| `Lowercase` | `bool` | `false` | Convert entire output to lowercase |
| `NegativeWord` | `bool` | `false` | Use the negative word instead of a negative sign, e.g., "minus one" |
| `OverrideOptions` | `*OverrideOptions` | `nil` | Custom currency and formatting options |

### OverrideOptions

| Option | Type | Description |
|--------|------|-------------|
| `Name` | `string` | Currency name override (singular), e.g., "peso" |
| `Plural` | `string` | Currency plural override, e.g., "pesos" |
| `Symbol` | `string` | Currency symbol override, e.g., "₱" |
| `FractionUnitName` | `string` | Fractional unit name override, e.g., "centavo" |
| `FractionPlural` | `string` | Fractional plural override, e.g., "centavos" |
| `FractionSymbol` | `string` | Fraction symbol override, e.g., "¢" |
| `NegativeWord` | `string` | Custom negative word override, e.g., "minus" |

## Supported Locales

### 🌍 **100+ Supported Locales**

| Locale Code | Language | Country/Region |
|-------------|----------|----------------|
| **Africa** | | |
| `af-ZA` | Afrikaans | South Africa |
| `am-ET` | Amharic | Ethiopia |
| `ha-NG` | Hausa | Nigeria |
| `sw-KE` | Swahili | Kenya |
| `xh-ZA` | Xhosa | South Africa |
| `zu-ZA` | Zulu | South Africa |
| **Americas** | | |
| `en-CA` | English | Canada |
| `en-US` | English | United States |
| `es-AR` | Spanish | Argentina |
| `es-BO` | Spanish | Bolivia |
| `es-CL` | Spanish | Chile |
| `es-CO` | Spanish | Colombia |
| `es-CR` | Spanish | Costa Rica |
| `es-CU` | Spanish | Cuba |
| `es-DO` | Spanish | Dominican Republic |
| `es-EC` | Spanish | Ecuador |
| `es-GT` | Spanish | Guatemala |
| `es-HN` | Spanish | Honduras |
| `es-MX` | Spanish | Mexico |
| `es-NI` | Spanish | Nicaragua |
| `es-PA` | Spanish | Panama |
| `es-PE` | Spanish | Peru |
| `es-PY` | Spanish | Paraguay |
| `es-SV` | Spanish | El Salvador |
| `es-UY` | Spanish | Uruguay |
| `es-VE` | Spanish | Venezuela |
| `fr-CA` | French | Canada |
| `pt-BR` | Portuguese | Brazil |
| **Asia** | | |
| `ar-AE` | Arabic | UAE |
| `ar-BH` | Arabic | Bahrain |
| `ar-EG` | Arabic | Egypt |
| `ar-IQ` | Arabic | Iraq |
| `ar-JO` | Arabic | Jordan |
| `ar-KW` | Arabic | Kuwait |
| `ar-LB` | Arabic | Lebanon |
| `ar-OM` | Arabic | Oman |
| `ar-QA` | Arabic | Qatar |
| `ar-SA` | Arabic | Saudi Arabia |
| `ar-SY` | Arabic | Syria |
| `bn-BD` | Bengali | Bangladesh |
| `bn-IN` | Bengali | India |
| `zh-CN` | Chinese | China |
| `zh-HK` | Chinese | Hong Kong |
| `zh-SG` | Chinese | Singapore |
| `zh-TW` | Chinese | Taiwan |
| `hi-IN` | Hindi | India |
| `gu-IN` | Gujarati | India |
| `kn-IN` | Kannada | India |
| `ml-IN` | Malayalam | India |
| `mr-IN` | Marathi | India |
| `ta-IN` | Tamil | India |
| `ur-IN` | Urdu | India |
| `id-ID` | Indonesian | Indonesia |
| `ja-JP` | Japanese | Japan |
| `ko-KR` | Korean | South Korea |
| `ms-MY` | Malay | Malaysia |
| `my-MM` | Myanmar | Myanmar |
| `ne-NP` | Nepali | Nepal |
| `ph-PH` | Filipino | Philippines |
| `th-TH` | Thai | Thailand |
| `tr-TR` | Turkish | Turkey |
| `ur-PK` | Urdu | Pakistan |
| `vi-VN` | Vietnamese | Vietnam |
| **Europe** | | |
| `be-BY` | Belarusian | Belarus |
| `bg-BG` | Bulgarian | Bulgaria |
| `bs-BA` | Bosnian | Bosnia and Herzegovina |
| `ca-ES` | Catalan | Spain |
| `cs-CZ` | Czech | Czech Republic |
| `cy-GB` | Welsh | United Kingdom |
| `da-DK` | Danish | Denmark |
| `de-AT` | German | Austria |
| `de-CH` | German | Switzerland |
| `de-DE` | German | Germany |
| `el-GR` | Greek | Greece |
| `en-GB` | English | United Kingdom |
| `en-IE` | English | Ireland |
| `es-ES` | Spanish | Spain |
| `et-EE` | Estonian | Estonia |
| `eu-ES` | Basque | Spain |
| `fi-FI` | Finnish | Finland |
| `fo-FO` | Faroese | Faroe Islands |
| `fr-BE` | French | Belgium |
| `fr-CH` | French | Switzerland |
| `fr-FR` | French | France |
| `fr-LU` | French | Luxembourg |
| `ga-IE` | Irish | Ireland |
| `gd-GB` | Scottish Gaelic | United Kingdom |
| `gl-ES` | Galician | Spain |
| `hr-HR` | Croatian | Croatia |
| `hu-HU` | Hungarian | Hungary |
| `is-IS` | Icelandic | Iceland |
| `it-CH` | Italian | Switzerland |
| `it-IT` | Italian | Italy |
| `lb-LU` | Luxembourgish | Luxembourg |
| `lt-LT` | Lithuanian | Lithuania |
| `lv-LV` | Latvian | Latvia |
| `mk-MK` | Macedonian | North Macedonia |
| `nl-BE` | Dutch | Belgium |
| `nl-NL` | Dutch | Netherlands |
| `nn-NO` | Norwegian Nynorsk | Norway |
| `no-NO` | Norwegian | Norway |
| `pl-PL` | Polish | Poland |
| `pt-PT` | Portuguese | Portugal |
| `rm-CH` | Romansh | Switzerland |
| `ro-RO` | Romanian | Romania |
| `ru-BY` | Russian | Belarus |
| `ru-RU` | Russian | Russia |
| `sk-SK` | Slovak | Slovakia |
| `sl-SI` | Slovenian | Slovenia |
| `sq-AL` | Albanian | Albania |
| `sr-RS` | Serbian | Serbia |
| `sv-FI` | Swedish | Finland |
| `sv-SE` | Swedish | Sweden |
| `uk-UA` | Ukrainian | Ukraine |
| **Oceania** | | |
| `en-AU` | English | Australia |
| `en-NZ` | English | New Zealand |
| `mi-NZ` | Māori | New Zealand |

## Examples

For complete examples, see [`example/en-us_sample.go`](example/en-us_sample.go) which demonstrates:

- Basic number to words conversion
- Currency formatting across multiple locales
- Decimal handling and formatting
- Custom currency configuration
- Multi-locale comparison

## Testing

Run the comprehensive test suite:

```bash
go test -v ./...
```

The library includes extensive tests covering:
- All supported locales
- Currency conversion accuracy
- Decimal precision handling
- Roman numeral conversion
- Edge cases and error handling

## Keywords

`go` `golang` `number-to-words` `internationalization` `i18n` `l10n` `currency` `localization` `decimal` `shopspring-decimal` `multi-locale` `financial` `invoice` `converter` `words` `numbers` `formatting` `unicode` `emoji` `flags` `countries` `language` `numeral` `roman-numerals` `ordinal` `cardinal` `money` `accounting` `business` `enterprise` `library` `package` `module`

## Go Package Information

```go
// Package numi18n provides internationalized number-to-words conversion
// with support for 100+ locales, currencies, and custom formatting.
//
// Features:
//   - Convert numbers to words in multiple languages
//   - Currency formatting with locale-specific rules  
//   - Decimal and fraction handling
//   - Roman numeral conversion
//   - Ordinal number support
//   - Custom currency overrides
//   - Unicode emoji support for country flags
//
// Example:
//   options := &numi18n.NumI18NOptions{
//       Locale: "en-US",
//       WordDetails: &numi18n.WordDetails{
//           Currency: true,
//           Decimal: true,
//           Capitalize: true,
//       },
//   }
//   result := options.ToWords(1234.56)
//   // Output: "One Thousand Two Hundred Thirty Four Dollars And Fifty Six Cents"
package numi18n
```

## Tags

- `#golang`
- `#go`
- `#internationalization`
- `#i18n`
- `#localization` 
- `#currency`
- `#numbers`
- `#words`
- `#conversion`
- `#formatting`
- `#multilingual`
- `#financial`
- `#accounting`
- `#invoice`
- `#decimal`
- `#unicode`
- `#emoji`
- `#countries`
- `#languages`

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.