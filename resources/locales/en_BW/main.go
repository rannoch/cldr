package en_BW

import "github.com/rannoch/cldr"

var Locale = &cldr.Locale{
	Locale: "en_BW",
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
