package aftermarketpl

//Aftermarketpl structs holds information neede to perform API request
type Aftermarketpl struct {
	url    string
	key    string
	secret string
}

//New creates new instace of Aftermarketpl struct with default URL
func New(key, secret string) *Aftermarketpl {
	return &Aftermarketpl{
		key:    key,
		secret: secret,
		url:    "https://json.aftermarket.pl",
	}
}

//NewCustomURL creates new instace of Aftermarketpl struct with custom URL
func NewCustomURL(key, secret, url string) *Aftermarketpl {
	return &Aftermarketpl{
		key:    key,
		secret: secret,
		url:    url,
	}
}

//Do executes action with given request and ouput it into response
func (a *Aftermarketpl) Do(actionName string, request, response interface{}) error {
	requestURL, err := a.createActionURL(actionName)
	if err != nil {
		return err
	}

	requestBody, err := encodeRequest(request)
	if err != nil {
		return err
	}

	responseData, err := a.send(requestURL, requestBody)
	if err != nil {
		return err
	}

	err = parseResponse(responseData, &response)
	if err != nil {
		return err
	}

	return nil
}
