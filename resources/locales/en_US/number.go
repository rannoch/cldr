package en_US

import "github.com/rannoch/cldr"

var (
	symbols = cldr.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"}
	formats = cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00;(¤#,##0.00)", Percent: "#,##0%"}
)
