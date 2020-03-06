package main

import (
	"github.com/byrdapp/timestamp"
)

var (
	t     int64
	etype string
)

func main() {
	timestamp.Parse(t, etype)
}
