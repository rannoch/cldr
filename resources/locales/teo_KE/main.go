package teo_KE

import "github.com/rannoch/cldr"

var Locale = &cldr.Locale{
	Locale: "teo_KE",
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
