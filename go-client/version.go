package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func (client *amberClient) GetAmberVersion() (*Version, error) {
	url := fmt.Sprintf("%s/appraisal/v1/version", client.cfg.Url)

	newRequest := func() (*http.Request, error) {
		return http.NewRequest(http.MethodGet, url, nil)
	}

	var headers = map[string]string{
		headerXApiKey: client.cfg.ApiKey,
	}

	var ver Version
	processResponse := func(resp *http.Response) error {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Errorf("Failed to read body from %s: %s", url, err)
		}

		if err = json.Unmarshal(body, &ver); err != nil {
			return errors.Errorf("Failed to devode json from %s: %s", url, err)
		}

		return nil
	}

	if err := doRequest(client.cfg.TlsCfg, newRequest, nil, headers, processResponse); err != nil {
		return nil, err
	}

	return &ver, nil
}