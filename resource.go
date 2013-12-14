package main

import (
	"fmt"
)

type Resource struct {
	Hostname string
	Port     int32
	Path     string
}

func (r *Resource) GetUrl() string {
	return fmt.Sprintf("http://%s:%d%s", r.Hostname, r.Port, r.Path)
}
