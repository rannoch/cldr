package fr_BI

import "github.com/rannoch/cldr"

var Locale = &cldr.Locale{
	Locale: "fr_BI",
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
