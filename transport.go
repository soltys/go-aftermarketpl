package aftermarketpl

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Aftermarketpl struct {
	url    string
	key    string
	secret string
}

func New(key, secret string) *Aftermarketpl {
	return &Aftermarketpl{
		key:    key,
		secret: secret,
		url:    "https://json.aftermarket.pl",
	}
}

func NewCustomUrl(key, secret, url string) *Aftermarketpl {
	return &Aftermarketpl{
		key:    key,
		secret: secret,
		url:    url,
	}
}

func (a *Aftermarketpl) Send(command string, params interface{}) ([]byte, error) {
	requestURL := a.url + command

	requestBody, err := json.Marshal(params)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
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
