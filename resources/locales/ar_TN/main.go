package ar_TN

import "github.com/rannoch/cldr"

var Locale = &cldr.Locale{
	Locale: "ar_TN",
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
