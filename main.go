package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/eschudt/name-generator/client"
)

type Hello struct {
	Message string
}

type Name struct {
	FirstName string
	LastName  string
}

type Age struct {
	Age      int
	YearBorn int
}

type NameAge struct {
	Name Name
	Age  Age
}

var (
	netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	consulClient *client.ConsulClient
)

func main() {
	url := os.Getenv("CONSUL_HTTP_ADDR")
	listenPort := os.Getenv("NOMAD_PORT_http")
	consulClient = client.NewConsulClient(url, "")

	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/name", nameHandler)
	mux.HandleFunc("/nameage", nameageHandler)
	http.ListenAndServe(":"+listenPort, mux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	data := Hello{"Hello World"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(resp)
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	data := Name{"John", "Smith"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(resp)
}

func nameageHandler(w http.ResponseWriter, r *http.Request) {
	address, port, err := consulClient.GetBaseURL("age-generator")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s:%d/age", address, port), nil)
		if err != nil {
			fmt.Printf("Creating request error: %s", err.Error())
		}
		result, err := netClient.Do(req)
		if err != nil {
			fmt.Printf("Do request error: %s", err.Error())
		}

		var actual Age
		bodyBytes, err := ioutil.ReadAll(result.Body)
		if err != nil {
			fmt.Printf("Read body error: %s", err.Error())
		}
		json.Unmarshal(bodyBytes, &actual)

		data := NameAge{
			Name: Name{
				FirstName: "John",
				LastName:  "Smith",
			},
			Age: actual,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		w.Write(resp)
	}
}
