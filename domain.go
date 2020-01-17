package aftermarketpl

import "log"

type domainGetRequest struct {
	Name string `url:"name"`
}

// DomainGet  returns the domain information on user account.
func (a *Aftermarketpl) DomainGet(name string) error {
	response, err := a.Send("/domain/get", domainGetRequest{Name: name})

	if err != nil {
		return err
	}

	log.Printf("%s", response)

	return nil
}
