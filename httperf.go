package main

import (
	"fmt"
	"strings"
)

const (
	NUMBER_OF_SESSION_ENTRIES = 1000
	NUMBER_OF_SESSION         = 10
)

func HttperfLine(path string, json []byte) string {
	return fmt.Sprintf("%s method=POST contents='%s'\n", path, strings.Replace(string(json), "'", "\\'", -1))
}

func WriteHttperfWsesslog(client DBClient) {
	for j := 0; j < NUMBER_OF_SESSION; j++ {
		for i := 0; i < NUMBER_OF_SESSION_ENTRIES; i++ {
			fmt.Print(client.HttperfLine())
		}
		fmt.Println()
	}
}
