package ar_LB

import "github.com/rannoch/cldr"

var Locale = &cldr.Locale{
	Locale: "ar_LB",
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
