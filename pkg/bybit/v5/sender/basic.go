package sender

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type Basic struct {
	httpClient *http.Client
}

func NewBasic() *Basic {
	return &Basic{httpClient: http.DefaultClient}
}

func (s *Basic) WithHttpClient(httpClient *http.Client) {
	s.httpClient = httpClient
}

func (s *Basic) Send(req *http.Request) (*http.Response, error) {
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error execute http request")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("incorrect status code=%d (%s)", resp.StatusCode, resp.Status)
	}
	return resp, nil
}
