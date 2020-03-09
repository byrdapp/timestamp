package parser

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"
)

func TestJSONUnmarshall(t *testing.T) {
	t.Run("parse millis", func(t *testing.T) {
		testformat := []struct {
			val []byte
			// expected Timestamp
		}{
			{[]byte("99999999999")},
			{[]byte("1583496449000")},
		}

		for _, v := range testformat {
			var ts Timestamp
			if err := json.Unmarshal(v.val, &ts); err != nil {
				log.Fatalf("jsonerr: %v", err)
			}
			log.Println(ts.Format())
		}
	})
}
func TestParser(t *testing.T) {
	t.Run("parse millis", func(t *testing.T) {
		testformat := []struct {
			val      int64
			expected Timestamp
		}{
			{1583496449000, parseMillis(1583496449000)},
			{99999999999, parseSeconds(99999999999)},
		}

		for _, v := range testformat {
			stamp := Parse(v.val, "millis")
			if stamp != v.expected {
				t.Fail()
				return
			}

			if !reflect.DeepEqual(stamp, v.expected) {
				t.Fail()
				return
			}
			log.Println("success")
		}

	})
}
