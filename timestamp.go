package timestamp

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
)

type EpochType string

const (
	DKFormat           = "02-01-2006T15:04:05Z07:00"
	Millis   EpochType = "millis"
	Seconds  EpochType = "seconds"
)

func Parse(t int64, etype string) {
	parse(t, etype)
}

func parse(t int64, etype string) {
	flag.Int64Var(&t, "value", 0, "parse a int64 (millis) timestamp")
	flag.StringVar(&etype, "epoch", "millis", "parse a int64 (millis) timestamp")
	flag.Parse()

	if t <= 0 {
		flag.PrintDefaults()
		log.Fatalln("value has zero or negative value")
	}

	var stamp Timestamp
	switch EpochType(etype) {
	case Millis:
		stamp = parseMillis(t)
	case Seconds:
		stamp = parseSeconds(t)
	default:
		stamp = parseMillis(t)
	}

	logger.Printf("%T => %v", stamp, stamp.Format())
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

func parseMillis(v int64) Timestamp {
	return Timestamp(time.Unix(0, v*int64(time.Millisecond))).UTC()
}

func parseSeconds(v int64) Timestamp {
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
