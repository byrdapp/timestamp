package main

import (
	"flag"
	"fmt"

	"github.com/byrdapp/timestamp/timeparser"
)

var (
	t     int64
	etype string
)

func main() {
	flag.Int64Var(&t, "value", 0, "parse a int64 (millis) timestamp")
	flag.StringVar(&etype, "epoch", "millis", "parse a int64 (millis) timestamp")
	flag.Parse()
	ts := timeparser.New(t)
	fmt.Println(ts.StringDKTime())
}
