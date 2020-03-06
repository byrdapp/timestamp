package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"time"
)

type Logger interface {
	Printf(format string, v ...interface{})
	Fatal(...interface{})
}

var (
	logger Logger = log.New(os.Stdout, "", 0)
	t      int64
	etype  string
)

type EpochType string

const (
	DKFormat           = "02-01-2006T15:04:05Z07:00"
	Millis   EpochType = "millis"
	Seconds  EpochType = "seconds"
)

func main() {
	flag.Int64Var(&t, "value", 0, "parse a int64 (millis) timestamp")
	flag.StringVar(&etype, "epoch", "millis", "parse a int64 (millis) timestamp")
	flag.Parse()

	logger.Printf("%v", t)
	if t <= 0 {
		flag.PrintDefaults()
		log.Fatalln("value has zero or negative value")
	}

	var stamp Timestamp
	switch EpochType(etype) {
	case Millis:
		stamp = ParseMillis(t)
	case Seconds:
		stamp = ParseSeconds(t)
	default:
		stamp = ParseMillis(t)
	}

	logger.Printf("%v : %T", stamp.Format(), stamp)
	if stamp.IsZero() {
		log.Fatalln("time has zero or wrong value")
	}
}

type Timestamp time.Time

func (t Timestamp) Format() string {
	return time.Time(t).Format(DKFormat)
}
func (t Timestamp) IsZero() bool {
	return time.Time(t).IsZero()
}
func (t Timestamp) String() string {
	return time.Time(t).String()
}
func (t Timestamp) UnixNano() int64 {
	return time.Time(t).UnixNano()
}
func (t Timestamp) UTC() Timestamp {
	return Timestamp(time.Time(t).UTC())
}

func ParseMillis(v int64) Timestamp {
	return Timestamp(time.Unix(0, v*int64(time.Millisecond))).UTC()
}

func ParseSeconds(v int64) Timestamp {
	return Timestamp(time.Unix(v, 0)).UTC()
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	millis, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	*t = Timestamp(time.Unix(0, millis))
	return nil
}
