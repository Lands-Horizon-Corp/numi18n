package locale

type Converter interface {
	ToWords(num int) string
}

var locales = map[string]Converter{}

func RegisterLocale(code string, conv Converter) {
	locales[code] = conv
}

func GetLocale(code string) Converter {
	return locales[code]
}
