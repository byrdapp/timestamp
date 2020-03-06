# Hrello

## Time stamp is by default in UTC and parses milliseconds (can be configured, see below). Input value must be an int64

### Run cmd

$ go build && ./timestamp -value={int64}

### For seconds

$ go build && ./timestamp -value={int64} -epoch=seconds
