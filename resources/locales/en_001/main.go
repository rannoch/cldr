package en_001

import "github.com/rannoch/cldr"

var Locale = &cldr.Locale{
	Locale: "en_001",
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
