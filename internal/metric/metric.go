package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	PriceLoaded = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "crypto_loader",
		Name:      "price_loaded",
		Help:      "The total number prices loaded",
	}, []string{"exchange"})

	PriceLoadDuration = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "crypto_loader",
		Name:      "price_load_duration",
		Help:      "The total duration prices loaded in ms",
	}, []string{"exchange"})

	PriceSaved = promauto.NewCounter(prometheus.CounterOpts{
		Name: "price_saved",
		Help: "The total number prices saved to storage",
	})

	Errors = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "crypto_loader",
		Name:      "errors",
		Help:      "The total errors",
	}, []string{"action", "source"})
)
