package aftermarketpl

type domainDNSListRequest struct {
	name string `url:"name"`
}

// DomainDNSList returns the list of domain DNS entries.
func (a *Aftermarketpl) DomainDNSList(name string) error {
	_, err := a.Send("/domain/dns/list", domainDNSListRequest{name: name})

	if err != nil {
		return err
	}

	return nil
}
