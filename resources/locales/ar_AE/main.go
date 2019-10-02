package ar_AE

import "github.com/rannoch/cldr"

var Locale = &cldr.Locale{
	Locale: "ar_AE",
	Number: cldr.Number{
		Symbols:    symbols,
		Formats:    formats,
		Currencies: currencies,
	},
	Calendar:   calendar,
	PluralRule: pluralRule,
}

func init() {
	cldr.RegisterLocale(Locale)
}
