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

type aftermarketResponse struct {
	Ok     int    `json:"ok"`
	Status int    `json:"status"`
	Error  string `json:"error"`
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
		return errors.New("Response is not typical Aftermarket.pl response")
	}

	if afResponse.Ok == 0 {
		return errors.New(afResponse.Error)
	}

	err = json.Unmarshal(response, &v)
	if err != nil {
		return err
	}
	return nil
}

func (a *Aftermarketpl) createActionURL(actionURL string) (string, error) {
	return a.url + actionURL, nil
}
