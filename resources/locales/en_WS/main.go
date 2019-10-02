package en_WS

import "github.com/rannoch/cldr"

var Locale = &cldr.Locale{
	Locale: "en_WS",
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
