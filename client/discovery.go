package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hashicorp/consul/api"
)

type ConsulClient struct {
	URL    string
	Token  string
	Client *http.Client
}

func NewConsulClient(baseURL, token string) *ConsulClient {
	cc := ConsulClient{
		baseURL,
		token,
		&http.Client{
			Timeout: time.Second * 10,
		},
	}
	return &cc
}

func (cc *ConsulClient) GetBaseURL(serviceName string) (string, int, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/health/service/%s?passing=true", cc.URL, serviceName), nil)
	if err != nil {
		return "", 0, err
	}

	resp, err := cc.Client.Do(req)
	if err != nil {
		return "", 0, err
	}

	defer resp.Body.Close()
	var out []*api.ServiceEntry
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(string(bodyBytes))

	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		return "", 0, err
	}

	for _, v := range out {
		return "http://" + v.Node.Address, v.Service.Port, nil
	}

	return "", 0, errors.New("Couldn't find the service")
}
