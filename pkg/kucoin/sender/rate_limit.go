package sender

import "time"

const DurationRateLimit = 30 * time.Second

const (
	SpotVip0       = 3000
	FuturesVip0    = 2000
	ManageMentVip0 = 2000
	PublicVip0     = 2000
)

func parseHeader() {
	//http-code = 429 -rate limit is exceeded,  wait 30s

	//gw-ratelimit-limit - total resource pool quota
	//gw-ratelimit-remaining - resource pool remaining quota
	//gw-ratelimit-reset - resource pool quota reset countdown (milliseconds)
}
