package main

import (
	"encoding/json"
)

const (
	SOLR_DEFAULT_HOST = "localhost"
	SOLR_DEFAULT_PORT = 8983
	SOLR_DEFAULT_PATH = "/solr/update/json"
)

type SolrClient struct {
	*Resource
}

func (client *SolrClient) GetPVEntryJson() []byte {
	json, _ := json.Marshal([]PVEntry{GetPVEntry()})
	return json
}

func (client *SolrClient) Post() {
	PostJsonViaHttp(client.GetUrl(), client.GetPVEntryJson())
}

func (client *SolrClient) HttperfLine() string {
	return HttperfLine(client.Path, client.GetPVEntryJson())
}

func NewSolrClient() *SolrClient {
	resource := Resource{SOLR_DEFAULT_HOST, SOLR_DEFAULT_PORT, SOLR_DEFAULT_PATH}
	return &SolrClient{&resource}
}
