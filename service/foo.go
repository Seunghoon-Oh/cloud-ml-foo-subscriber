package service

import (
	"encoding/json"
	"fmt"

	"github.com/Seunghoon-Oh/cloud-ml-foo-subscriber/network"
	circuit "github.com/rubyist/circuitbreaker"
)

var cb *circuit.Breaker
var httpClient *circuit.HTTPClient

func SetupFooCircuitBreaker() {
	httpClient, cb = network.GetHttpClient()
}

func CreateFoo() {
	if cb.Ready() {
		resp, err := httpClient.Post("http://cloud-ml-foo-manager.cloud-ml-foo:8082/foo", "", nil)
		if err != nil {
			fmt.Println(err)
			cb.Fail()
			return
		}
		cb.Success()
		defer resp.Body.Close()
		rsData := network.ResponseData{}
		json.NewDecoder(resp.Body).Decode(&rsData)
		fmt.Println(rsData.Data)
		return
	}
}
