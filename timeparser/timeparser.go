package timeparser

import (
	"errors"
	"fmt"
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

type Timestamp time.Time

func New(t int64) (Timestamp, error) {
	var stamp Timestamp
	return stamp.parse(t)
}

func (t Timestamp) parse(val int64) (Timestamp, error) {
	if val <= 0 {
		return t, errors.New(fmt.Sprintf("value: %v has zero or negative value", t))
	}

	if val <= MaxSecondsCap {
		t = parseSeconds(val)
	} else {
		t = parseMillis(val)
	}
	if t.IsZero() {
		log.Printf("time has zero or wrong value: %v", val)
	}
	return t, nil
}

func (t Timestamp) FormatDKTime() string {
	return time.Time(t).Format(DKFormat)
}
func (t Timestamp) IsZero() bool {
	return time.Time(t).IsZero()
}
func (t Timestamp) StringDKTime() string {
	return t.FormatDKTime()
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
	if intval <= MaxSecondsCap {
		*t = parseSeconds(intval)
	} else {
		*t = parseMillis(intval)
	}
	return nil
}
