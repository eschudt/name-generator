package client

import (
	"encoding/json"
	"github.com/hashicorp/consul/api"
	"io/ioutil"
	"net/http"
	"time"
)

type ConsulClient struct {
	BaseURL string
	Port    string
	Token   string
	Client  *http.Client
}

func NewConsulClient(baseURL, port, token string) *ConsulClient {
	cc := ConsulClient{
		baseURL,
		port,
		token,
		&http.Client{
			Timeout: time.Second * 10,
		},
	}
	return &cc
}

func (cc *ConsulClient) GetBaseURL(serviceName string) (url string, port int) {

	req, err := http.NewRequest("GET", cc.BaseURL+"/v1/health/service/"+serviceName+"?passing=true", nil)
	if err != nil {
		return "", 0
	}

	resp, err := cc.Client.Do(req)
	if err != nil {
		return "", 0
	}

	defer resp.Body.Close()
	var out []*api.ServiceEntry
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		return "", 0
	}

	for _, v := range out {
		return v.Node.Address, v.Service.Port
	}

	return "", 0
}
