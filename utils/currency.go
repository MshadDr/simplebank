package utils

// constatn value for available currencies
const(
	USD = "USD"
	CAD = "CAD"
	EUR = "EUR"
)

// if the currency is supported it returns true
func IsSuppertedCurrency( currency string ) bool {
	switch currency {
	case USD,CAD, EUR:
		return true
	}
	return false
}
