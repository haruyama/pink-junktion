package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	DEFAULT_HOST = "localhost"
	DEFAULT_PORT = 8983
	DEFAULT_PATH = "/solr/update/json"
)

type SolrClient struct {
	resource Resource
}

func (info *SolrClient) GetPVEntryJson() []byte {
	json, _ := json.Marshal([]PVEntry{GetPVEntry()})
	return json
}

func (info *SolrClient) Post() {
	json := info.GetPVEntryJson()
	resp, err := http.Post(info.resource.GetUrl(), "application/json", bytes.NewReader(json))
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(json))
		fmt.Println(resp)
	}
}

func (info *SolrClient) HttperfLine() string {
	json := info.GetPVEntryJson()
	return fmt.Sprintf("%s method=POST contents='%s'\n", info.resource.Path, strings.Replace(string(json), "'", "\\'", -1))
}

func NewSolrClient() *SolrClient {
	resource := Resource{DEFAULT_HOST, DEFAULT_PORT, DEFAULT_PATH}
	return &SolrClient{resource}
}
