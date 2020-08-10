package card

type CurrencyNotFoundError struct {
	currencyCode string
}

func (e CurrencyNotFoundError) Error() string {
	return "Given currency not found: " + e.currencyCode
}

type CardNotFoundError struct {
	reference string
}

func (e CardNotFoundError) Error() string {
	return "Card not found with given reference: " + e.reference
}
