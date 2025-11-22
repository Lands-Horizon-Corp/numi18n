package src

func (op NumI18NOptions) CurrencyDetails() NumI18NCurrency {
	return NumI18NCurrency{
		CountryName:    op.CountryName,
		Currency:       op.Currency,
		ISO3166Alpha2:  op.ISO3166Alpha2,
		ISO3166Alpha3:  op.ISO3166Alpha3,
		ISO3166Numeric: op.ISO3166Numeric,
		Locale:         op.Locale,
		Timezone:       op.Timezone,
		Language:       op.Language,
	}
}

func CurrencyList() []NumI18NCurrency {
	return []NumI18NCurrency{}
}
