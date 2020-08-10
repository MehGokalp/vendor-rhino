package card

type CurrencyNotFoundError struct {
	currencyCode string
}

func (e CurrencyNotFoundError) Error() string {
	return "Given currency not found: " + e.currencyCode
}
