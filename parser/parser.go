package parser

import (
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

const (
	DKFormat      = "02-01-2006T15:04:05Z07:00"
	MaxSecondsCap = int64(9999999999)
)

func Parse(t int64) Timestamp {
	return parse(t)
}

func parse(t int64) Timestamp {
	if t <= 0 {
		log.Printf("value: %v has zero or negative value", t)
	}

	var stamp Timestamp

	if t <= MaxSecondsCap {
		logger.Printf("is seconds!")
		stamp = parseSeconds(t)
	} else {
		logger.Printf("is millis!")
		stamp = parseMillis(t)
	}
	if stamp.IsZero() {
		log.Printf("time has zero or wrong value: %v", t)
	}
	return stamp
}

type Timestamp time.Time

func (t Timestamp) Format() string {
	return time.Time(t).Format(DKFormat)
}
func (t Timestamp) IsZero() bool {
	return time.Time(t).IsZero()
}
func (t Timestamp) String() string {
	return t.Format()
}
func (t Timestamp) Unix() int64 {
	return time.Time(t).Unix()
}

// func (t Timestamp) UnixNano() int64 {
// 	return time.Time(t).UnixNano()
// }
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
	intval, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	if intval <= time.Second.Microseconds() {
		logger.Printf("is seconds!")
		*t = parseSeconds(intval)
	} else {
		logger.Printf("is millis!")
		*t = parseMillis(intval)
	}
	return nil
}
