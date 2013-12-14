package main

type DBClient interface {
	Post()
	HttperfLine() string
}
