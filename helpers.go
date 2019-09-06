package PayKassaSCI

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func sendRequest(param map[string]string) ([]byte, int, error) {

	form := url.Values{}
	for key, value := range param {
		form.Add(key, value)
	}

	// Create request
	req, err := http.NewRequest("POST", PayKassaURL, strings.NewReader(form.Encode()))
	if err != nil {
		return []byte{}, 0, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	// Send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, 0, err
	}
	defer resp.Body.Close()

	// Get form resp
	data, err := ioutil.ReadAll(resp.Body)
	return data, resp.StatusCode, err
}

func handlerResp(data []byte, statusCode int, resp interface{}) error {
	var err error

	// Check status code
	if statusCode != http.StatusOK {
		if statusCode/100 == 4 {
			if err = json.Unmarshal(data, &resp); err != nil {
				return err
			}
		}

		return errors.New("response with " + strconv.Itoa(statusCode) + " status code")
	}

	// Unmarshal data
	if err = json.Unmarshal(data, &resp); err != nil {
		return err
	}

	return err
}
