package aftermarketpl

type domainGetRequest struct {
	Name string `url:"name"`
}

//DomainGet  returns the domain information on user account.
func (a *Aftermarketpl) DomainGet(name string) error {
	_, err := a.Send("/domain/get", domainGetRequest{Name: name})

	if err != nil {
		return err
	}

	return nil
}
