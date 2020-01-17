package aftermarketpl

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

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

//Send is basic function to send struct by request
func (a *Aftermarketpl) Send(command string, params interface{}) ([]byte, error) {
	requestURL := a.url + command

	requestBody, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", requestURL, bytes.NewBufferString(requestBody.Encode()))

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	request.SetBasicAuth(a.key, a.secret)

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
