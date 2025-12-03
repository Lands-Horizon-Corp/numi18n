# Number to Words (Go)

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
import "github.com/Lands-Horizon-Corp/numi18n/pkg"
```

### Simple Number to Words

```go
package main

import (
    "fmt"
    "github.com/Lands-Horizon-Corp/numi18n/pkg"
)

func main() {
    // Simple number conversion
    options := &pkg.NumI18NOptions{
        Locale: "en-US",
        WordDetails: &pkg.WordDetails{
            Capitalize: true,
        },
    }
    
    result := options.ToWords(123)
    fmt.Println(result) // "One Hundred Twenty Three"
    
    // Number with decimals
    decimalOptions := &pkg.NumI18NOptions{
        Locale: "en-US", 
        WordDetails: &pkg.WordDetails{
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
currencyOptions := &pkg.NumI18NOptions{
    Locale: "en-US",
    WordDetails: &pkg.WordDetails{
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
japaneseOptions := &pkg.NumI18NOptions{
    Locale: "ja-JP",
    WordDetails: &pkg.WordDetails{
        Currency:   true,
        Decimal:    true,
        Capitalize: true,
    },
}

words := japaneseOptions.ToWords(1234.56)
fmt.Println(words) // "千二百三十四円と五十六銭"

// Spanish (Mexico)
spanishOptions := &pkg.NumI18NOptions{
    Locale: "es-MX", 
    WordDetails: &pkg.WordDetails{
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
customOptions := &pkg.NumI18NOptions{
    Locale: "en-US",
    WordDetails: &pkg.WordDetails{
        Currency:   true,
        Decimal:    true,
        Capitalize: true,
        OverrideOptions: &pkg.OverrideOptions{
            CurrencyName:       "Bitcoin",
            CurrencyPluralName: "Bitcoins", 
            CurrencySymbol:     "₿",
            FractionalName:     "Satoshi",
            FractionalPluralName: "Satoshis",
        },
    },
}

result := customOptions.ToWords(1.50)
fmt.Println(result) // "One Bitcoin And Fifty Satoshis"
```

### Roman Numerals

```go
// Convert to Roman numerals
roman := pkg.ToRoman(1994)
fmt.Println(roman) // "MCMXCIV"

// Convert from Roman numerals
number := pkg.FromRoman("MCMXCIV") 
fmt.Println(number) // 1994

// Validate Roman numerals
isValid := pkg.IsValid("MCMXCIV")
fmt.Println(isValid) // true

isValid = pkg.IsValid("IIII")
fmt.Println(isValid) // false
```

## Configuration Options

### NumI18NOptions

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `Locale` | `string` | `"en-US"` | Locale code for language/region selection |
| `WordDetails` | `*WordDetails` | `nil` | Detailed configuration for word conversion |

### WordDetails Options

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `Currency` | `bool` | `false` | Convert as currency with currency units |
| `Decimal` | `bool` | `false` | Include decimal/fractional parts in words |
| `Capitalize` | `bool` | `false` | Capitalize the first letter of output |
| `UpperCase` | `bool` | `false` | Convert entire output to uppercase |
| `LowerCase` | `bool` | `false` | Convert entire output to lowercase |  
| `NegativeWord` | `bool` | `false` | Include negative indicator for negative numbers |
| `OverrideOptions` | `*OverrideOptions` | `nil` | Custom currency and formatting options |

### OverrideOptions

| Option | Type | Description |
|--------|------|-------------|
| `CurrencyName` | `string` | Override currency name (singular) |
| `CurrencyPluralName` | `string` | Override currency name (plural) |
| `CurrencySymbol` | `string` | Override currency symbol |
| `FractionalName` | `string` | Override fractional unit name (singular) |
| `FractionalPluralName` | `string` | Override fractional unit name (plural) |
| `NegativeWord` | `string` | Override negative word indicator |

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

`golang` `number-to-words` `internationalization` `i18n` `currency` `localization` `decimal` `shopspring-decimal` `multi-locale` `financial` `invoice` `converter` `words` `numbers` `formatting`

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.