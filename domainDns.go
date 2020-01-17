package aftermarketpl

type domainDNSListRequest struct {
	Name string `url:"name"`
}

// DomainDNSList returns the list of domain DNS entries.
func (a *Aftermarketpl) DomainDNSList(name string) error {
	response, err := a.Send("/domain/dns/list", domainDNSListRequest{Name: name})
	
	if err != nil {
		return err
	}

	return nil
}
