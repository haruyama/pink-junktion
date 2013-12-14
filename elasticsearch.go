package main

import (
	"encoding/json"
)

const (
	ES_DEFAULT_HOST = "localhost"
	ES_DEFAULT_PORT = 9200
	ES_DEFAULT_PATH = "/access/pv"
)

type ESClient struct {
	*Resource
}

func (client *ESClient) GetPVEntryJson() []byte {
	json, _ := json.Marshal(GetPVEntry())
	return json
}

func (client *ESClient) Post() {
	PostJsonViaHttp(client.GetUrl(), client.GetPVEntryJson())
}

func (client *ESClient) HttperfLine() string {
	return HttperfLine(client.Path, client.GetPVEntryJson())
}

func NewESClient() *ESClient {
	resource := Resource{ES_DEFAULT_HOST, ES_DEFAULT_PORT, ES_DEFAULT_PATH}
	return &ESClient{&resource}
}
