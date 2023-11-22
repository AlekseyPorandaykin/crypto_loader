package sender

import (
	"net/http"
	"time"
)

// Personal - используется для лимитирования запросов от одного клиента
type Personal struct {
	sender Sender
	locker *Locker
}

func NewPersonal(sender Sender) *Personal {
	return &Personal{sender: sender, locker: NewLocker()}
}

func (p *Personal) Send(req *http.Request, weight int) (*http.Response, SenderError) {
	p.locker.SyncDelay(time.Second)
	resp, err := p.sender.Send(req, weight)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *Personal) Close() {
	p.sender.Close()
}
