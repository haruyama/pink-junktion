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

func NewDBClient() DBClient {
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

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	httperf := flag.Bool("httperf", false, "write httperf wsesslog")
	flag.Parse()

	client := NewDBClient()

	if *httperf {
		write_httperf_wsesslog(client)
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
