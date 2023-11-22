package sender

import (
	"net/http"
	"strconv"
	"strings"
)

type IPLimitWeight struct {
	Second int
	Minute int
	Hour   int
	Day    int

	LastIpWeight int
}

func IpUsedWeight(header http.Header) IPLimitWeight {
	var res IPLimitWeight
	for name, values := range header {
		key := strings.ToUpper(name)
		if key == "X-MBX-USED-WEIGHT" {
			for _, val := range values {
				if currentWeightInMinute, _ := strconv.Atoi(val); currentWeightInMinute > 0 {
					res.Minute = currentWeightInMinute
				}
			}
		}
		if strings.Contains(key, "X-MBX-USED-WEIGHT-") {
			dur := strings.Replace(key, "X-MBX-USED-WEIGHT-", "", 1)
			parseWeightLimit(dur, values, &res)
		}
		if strings.Contains(key, "X-SAPI-USED-IP-WEIGHT-") {
			dur := strings.Replace(key, "X-SAPI-USED-IP-WEIGHT-", "", 1)
			parseWeightLimit(dur, values, &res)
		}
	}
	return res
}

func parseWeightLimit(dur string, values []string, res *IPLimitWeight) {
	if strings.Contains(dur, "S") {
		weightSec, _ := strconv.Atoi(strings.Replace(dur, "S", "", 1))
		if weightSec > 0 {
			res.Second = weightSec
			res.LastIpWeight = firstWeightVal(values)
		}
	}
	if strings.Contains(dur, "M") {
		weightMin, _ := strconv.Atoi(strings.Replace(dur, "M", "", 1))
		if weightMin > 0 {
			res.Minute = weightMin
			res.LastIpWeight = firstWeightVal(values)
		}

	}
	if strings.Contains(dur, "H") {
		weightHour, _ := strconv.Atoi(strings.Replace(dur, "H", "", 1))
		if weightHour > 0 {
			res.Hour = weightHour
			res.LastIpWeight = firstWeightVal(values)
		}

	}
	if strings.Contains(dur, "D") {
		weightDay, _ := strconv.Atoi(strings.Replace(dur, "D", "", 1))
		if weightDay > 0 {
			res.Day = weightDay
			res.LastIpWeight = firstWeightVal(values)
		}

	}
}

func firstWeightVal(values []string) int {
	var weight int
	for _, val := range values {
		weightVal, _ := strconv.Atoi(val)
		if weightVal > 0 {
			weight = weightVal
			break
		}
	}
	return weight
}
