package util

// / Constants for supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

func IsSUpportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
