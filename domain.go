package aftermarketpl

import "log"

import "encoding/json"

type domainGetRequest struct {
	Name string `url:"name"`
}

type domainGetResponse struct {
	Name         string `json:"name"`
	AutoRenew    bool   `json:"autorenew"`
	CostCurrency string `json:"costCurrency"`
}

// DomainGet  returns the domain information on user account.
func (a *Aftermarketpl) DomainGet(name string) error {
	response, err := a.Send("/domain/get", domainGetRequest{Name: name})

	if err != nil {
		return err
	}

	d := aftermarketResponse{
		Data: domainGetResponse{},
	}
	err = json.Unmarshal(response, &d)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Print(d.Data.(domainGetResponse).Name)
		log.Print(d.Data.(domainGetResponse).AutoRenew)
		log.Print(d.Data.(domainGetResponse).CostCurrency)
	}

	return nil
}
