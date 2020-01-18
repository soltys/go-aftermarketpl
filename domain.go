package aftermarketpl

type domainGetRequest struct {
	Name string `url:"name"`
}

//DomainGetDataResponse represents domain information on user account.
type DomainGetDataResponse struct {
	Name         string `json:"name"`
	AutoRenew    bool   `json:"autorenew"`
	CostCurrency string `json:"costCurrency"`
}

type domainGetResponse struct {
	Data DomainGetDataResponse `json:"data"`
}

// DomainGet returns the domain information on user account.
func (a *Aftermarketpl) DomainGet(name string) (*DomainGetDataResponse, error) {
	d := domainGetResponse{}
	err := a.Do("/domain/get", domainGetRequest{Name: name}, &d)
	return &d.Data, err
}
