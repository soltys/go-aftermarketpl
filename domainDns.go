package aftermarketpl

type domainNameRequest struct {
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
	err := a.Do("/domain/dns/list", domainNameRequest{Name: name}, &d)
	return d.Data, err
}

type domainDNSListCountResponse struct {
	Data int `json:"data"`
}

// DomainDNSListCount returns the number of domain DNS entries.
func (a *Aftermarketpl) DomainDNSListCount(name string) (int, error) {
	d := domainDNSListCountResponse{}
	err := a.Do("/domain/dns/list/count", domainNameRequest{Name: name}, &d)
	return d.Data, err
}

type domainDNSRemoveRequest struct {
	EntryID int    `url:"entryId"`
	Name    string `url:"name"`
}

type domainDNSRemoveResponse struct {
	Data bool `json:"data"`
}

//DomainDNSRemove Remove DNS entry for domain.
func (a *Aftermarketpl) DomainDNSRemove(name string, entryID int) (bool, error) {
	request := domainDNSRemoveRequest{Name: name, EntryID: entryID}
	d := domainDNSRemoveResponse{}
	err := a.Do("/domain/dns/remove", request, &d)
	return d.Data, err
}

const (
	RecordTypeA     string = "A"
	RecordTypeAAAA  string = "AAAA"
	RecordTypeCNAME string = "CNAME"
	RecordTypeMX    string = "MX"
	RecordTypeTXT   string = "TXT"
	RecordTypeNS    string = "NS"
	RecordTypeCAA   string = "CAA"
	RecordTypeSRV   string = "SRV"
)

type domainDNSAddRequest struct {
	Name       string `url:"name"`
	Value      string `url:"value"`
	RecordType string `url:"type"`
}

type domainDNSAddResponse struct {
	Data int `json:"data"`
}

//DomainDNSAdd adds new DNS entry for domain
func (a *Aftermarketpl) DomainDNSAdd(name, value, recordType string) (int, error) {
	request := domainDNSAddRequest{Name: name, Value: value, RecordType: recordType}
	d := domainDNSAddResponse{}
	err := a.Do("/domain/dns/add", request, &d)
	return d.Data, err
}
