package database

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var (
	CountQueryDB      *prometheus.CounterVec
	QueryDurationDB   *prometheus.CounterVec
	ErrorQueryDB      *prometheus.CounterVec
	TotalErrorQueryDB prometheus.Counter
)

func Init(namespace string) {
	CountQueryDB = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: "database",
		Name:      "db_query_count",
		Help:      "Count queries to db",
	}, []string{"query_name"})

	QueryDurationDB = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: "database",
		Name:      "db_query_duration",
		Help:      "Duration execute query(in ms)",
	}, []string{"query_name"})
	TotalErrorQueryDB = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: "database",
		Name:      "db_query_errors",
		Help:      "Count error in db",
	})
	ErrorQueryDB = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: "database",
		Name:      "db_query_errors",
		Help:      "Count error in db",
	}, []string{"query_name"})
}

func WithExecuteMetric(name string, start time.Time) {
	DurationExecuteQueryDB(name, time.Since(start))
	IncCountQueryDB(name)
}

func IncCountQueryDB(name string) {
	CountQueryDB.WithLabelValues(name).Inc()
}

func DurationExecuteQueryDB(name string, duration time.Duration) {
	QueryDurationDB.WithLabelValues(name).Add(float64(duration.Milliseconds()))
}

func IncErrorQueryDB(name string) {
	TotalErrorQueryDB.Inc()
	ErrorQueryDB.WithLabelValues(name).Inc()
}
