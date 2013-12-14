package main

import (
	"flag"
	"math/rand"
	"time"
)

const (
	NUMBER_OF_COROUTINE = 10
	NUMBER_OF_LOOP      = 1000
	RATE_PER_SEC        = 1000
	URL_ELASTICSEARCH   = "http://localhost:9200/access_info/access_info"
)

func NewDBClient(db string) DBClient {
	if db == "elasticsearch" || db == "es" {
		return NewESClient()
	}
	return NewSolrClient()
}

func post(client DBClient, c chan int) {
	throttle := time.Tick(1e9 / RATE_PER_SEC)
	for i := 0; i < NUMBER_OF_LOOP; i++ {
		<-throttle
		client.Post()
	}
	c <- 1
}

var (
	httperf bool
	db      string
)

func init() {
	flag.BoolVar(&httperf, "httperf", false, "httperf wsesslog mode")
	flag.StringVar(&db, "db", "solr", "db type")
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	flag.Parse()

	client := NewDBClient(db)

	if httperf {
		WriteHttperfWsesslog(client)
		return
	}

	c := make(chan int, NUMBER_OF_COROUTINE)
	for i := 0; i < NUMBER_OF_COROUTINE; i++ {
		go post(client, c)
	}
	for i := 0; i < NUMBER_OF_COROUTINE; i++ {
		<-c
	}
}
