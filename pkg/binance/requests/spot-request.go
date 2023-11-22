package requests

import (
	"net/url"
)

type SpotRequest struct {
	host *url.URL
}

func NewSpotRequest(host string) (*SpotRequest, error) {
	urlHost, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	return &SpotRequest{
		host: urlHost,
	}, nil
}
