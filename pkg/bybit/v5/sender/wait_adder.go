package sender

import (
	"net/http"
	"strconv"
	"time"
)

// 600 запросов разрешено в 5 секундный интервал, но у некоторых запросов лимит маленький, и можно упереться в него раньше блокера.
// Поэтому ставим лимит по самому крайнему ограничению = 5 запросов в секунду
var (
	intervalBetweenRequests = time.Second / 5
	intervalLongWait        = 11 * time.Minute
	intervalMediumWait      = 30 * time.Second
	intervalShortWait       = time.Second / 5
)

var limitRequests = 10
var criticalLimitRequests = 5

type WaitAdder struct {
	sender Sender
}

func NewWaitAdder(sender Sender) Sender {
	return &WaitAdder{sender: sender}
}

func (s *WaitAdder) Send(req *http.Request) (Response, error) {
	resp, err := s.sender.Send(req)
	if err != nil {
		return resp, err
	}
	resp.AddActionWithWait("Add common interval between requests", intervalBetweenRequests)
	if resp.HttpResp.StatusCode == http.StatusForbidden {
		resp.AddActionWithWait("Status 403 Forbidden, add wait interval", intervalLongWait)
		return resp, err
	}
	limitStatusStr := resp.HttpResp.Header.Get("X-Bapi-Limit-Status")
	limitStatus, _ := strconv.Atoi(limitStatusStr) //your remaining requests for current endpoint
	resp.AddAction("Bapi Limit Status", resp.HttpResp.Header.Get("X-Bapi-Limit-Status"))
	resp.AddAction("Bapi Limit reset timestamp", resp.HttpResp.Header.Get("X-Bapi-Limit-Reset-Timestamp"))
	resp.AddAction("Bapi Limit", resp.HttpResp.Header.Get("X-Bapi-Limit"))
	if limitStatus < limitRequests && limitStatusStr != "" {
		now := time.Now().In(time.UTC)
		nextRequest := now.Add(1 * time.Second)
		resetTimestamp, _ := strconv.Atoi(resp.HttpResp.Header.Get("X-Bapi-Limit-Reset-Timestamp")) //the timestamp indicating when your request limit resets if you have exceeded your rate_limit. Otherwise, this is just the current timestamp.
		if resetTimestamp > 0 {
			nextRequest = time.UnixMilli(int64(resetTimestamp)).In(time.UTC)
		}
		diff := nextRequest.Sub(now)
		if diff <= 0 {
			diff = intervalShortWait
		}
		waitDuration := time.Duration(limitRequests-limitStatus) * diff
		if limitStatus < criticalLimitRequests {
			waitDuration = intervalMediumWait
		}
		resp.AddActionWithWait("Limit less threshold, add wait interval", waitDuration)
	}
	if resp.HttpResp.StatusCode != http.StatusOK {
		resp.AddActionWithWait("Status response not ok, add wait interval", intervalMediumWait)
		return resp, err
	}
	return resp, nil
}
