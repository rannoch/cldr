package mgo

import "github.com/rannoch/cldr"

var Locale = &cldr.Locale{
	Locale: "mgo",
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
