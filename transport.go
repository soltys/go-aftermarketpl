package aftermarketpl

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
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

type aftermarketResponse struct {
	Ok     int    `json:"ok"`
	Status int    `json:"status"`
	Error  string `json:"error"`
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
func (a *Aftermarketpl) send(requestURL string, requestBody io.Reader) ([]byte, error) {
	request, err := http.NewRequest("POST", requestURL, requestBody)

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

//Do executes action with given request and ouput it into response
func (a *Aftermarketpl) Do(actionName string, request, response interface{}) error {
	requestURL := a.url + actionName

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

func encodeRequest(request interface{}) (io.Reader, error) {
	urlValues, err := query.Values(request)
	if err != nil {
		return nil, err
	}

	return bytes.NewBufferString(urlValues.Encode()), nil
}

func parseResponse(response []byte, v interface{}) error {
	afResponse := aftermarketResponse{}
	err := json.Unmarshal(response, &afResponse)

	if err != nil {
		errors.New("Response is not typical Aftermarket.pl response")
	} else {
		if afResponse.Ok == 0 {
			return errors.New(afResponse.Error)
		}
	}

	err = json.Unmarshal(response, &v)
	if err != nil {
		return err
	}
	return nil
}
