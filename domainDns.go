package aftermarketpl
import "log"
type domainDNSListRequest struct {
	name string `url:"name"`
}

// DomainDNSList returns the list of domain DNS entries.
func (a *Aftermarketpl) DomainDNSList(name string) error {
	response, err := a.Send("/domain/dns/list", domainDNSListRequest{name: name})
	log.Println(string(response))
	if err != nil {
		return err
	}

	return nil
}
