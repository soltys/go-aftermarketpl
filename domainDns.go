package aftermarketpl

type domainDNSListRequest struct {
	Name string `url:"name"`
}

type domainDNSListResponse struct {
	Data []DomainDNSListEntry `json:"data"`
}

//DomainDNSListEntry represents DNS entry
type DomainDNSListEntry struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	EntryID int    `json:"entryId"`
	Type    string `json:"type"`
}

// DomainDNSList returns the list of domain DNS entries.
func (a *Aftermarketpl) DomainDNSList(name string) ([]DomainDNSListEntry, error) {
	d := domainDNSListResponse{}
	err := a.Do("/domain/get", domainDNSListRequest{Name: name}, &d)
	return d.Data, err
}
