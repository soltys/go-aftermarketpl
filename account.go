package aftermarketpl

type accountIDResponse struct {
	Data string `json:"data"`
}

// AccountID Returns the current user identifier.
func (a *Aftermarketpl) AccountID() (string, error) {
	d := accountIDResponse{}
	err := a.Do("/account/id", emptyRequest{}, &d)
	return d.Data, err
}

type accountCurrencyResponse struct {
	Data string `json:"data"`
}

// AccountCurrency Returns the current user currency.
func (a *Aftermarketpl) AccountCurrency() (string, error) {
	d := accountCurrencyResponse{}
	err := a.Do("/account/currency", emptyRequest{}, &d)
	return d.Data, err
}

type accountBalanceResponse struct {
	Data float64 `json:"data"`
}

// AccountBalance Returns the current user balance.
func (a *Aftermarketpl) AccountBalance() (float64, error) {
	d := accountBalanceResponse{}
	err := a.Do("/account/balance", emptyRequest{}, &d)
	return d.Data, err
}
