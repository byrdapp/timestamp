package main

import "github.com/byrdapp/timestamp/parser"

var (
	t     int64
	etype string
)

func main() {
	parser.Parse(t, etype)
}
