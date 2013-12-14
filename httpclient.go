package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func PostJsonViaHttp(url string, json []byte) {
	_, err := http.Post(url, "application/json", bytes.NewReader(json))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, string(json))
	}
}
