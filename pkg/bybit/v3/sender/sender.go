package sender

import (
	"net/http"
	"time"
)

type Sender interface {
	Send(req *http.Request) (*http.Response, error)
}

func timestamp() int64 {
	return time.Now().UnixNano() / 1000000
}
