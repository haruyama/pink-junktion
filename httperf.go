package main

import (
	"fmt"
)

const (
	NUMBER_OF_SESSION_ENTRIES = 1000
	NUMBER_OF_SESSION         = 10

//    PATH_OF_ELASTICSEARCH     = "/access_info/access_info"
)

func write_httperf_wsesslog(client DBClient) {
	for j := 0; j < NUMBER_OF_SESSION; j++ {
		for i := 0; i < NUMBER_OF_SESSION_ENTRIES; i++ {
			fmt.Print(client.HttperfLine())
		}
		fmt.Println()
	}
}
