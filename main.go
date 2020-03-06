package main

import (
	"flag"

	"github.com/byrdapp/timestamp/timestamp"
)

var (
	t     int64
	etype string
)

func main() {
	flag.Int64Var(&t, "value", 0, "parse a int64 (millis) timestamp")
	flag.StringVar(&etype, "epoch", "millis", "parse a int64 (millis) timestamp")
	flag.Parse()
	timestamp.Parse(t, etype)
}
