package bs_Cyrl

import "github.com/rannoch/cldr"

var Locale = &cldr.Locale{
	Locale: "bs_Cyrl",
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
