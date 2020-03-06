package parser

import (
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("parse millis", func(t *testing.T) {
		tests := []struct {
			val    int64
			format string
		}{
			{1583496449000, "06-03-2020T12:07:29Z"},
		}

	})
}
