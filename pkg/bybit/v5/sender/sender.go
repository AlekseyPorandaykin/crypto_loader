package sender

import (
	"net/http"
	"time"
)

type Action struct {
	Name      string
	Value     string
	Timestamp time.Time
}

type Response struct {
	HttpResp     *http.Response
	Actions      []Action
	WaitDuration time.Duration
}

func (r *Response) AddAction(name, value string) {
	r.Actions = append(r.Actions, Action{Name: name, Value: value, Timestamp: time.Now()})
}
func (r *Response) AddActionWithWait(name string, wait time.Duration) {
	r.Actions = append(r.Actions, Action{Name: name, Value: wait.String(), Timestamp: time.Now()})
	r.WaitDuration += wait
}

func NewResponse(httpResp *http.Response) Response {
	return Response{HttpResp: httpResp, Actions: make([]Action, 0)}
}

type Sender interface {
	Send(req *http.Request) (Response, error)
}
