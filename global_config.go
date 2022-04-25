package mgcache

import (
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

var (
	defaultMetricCollector      = NewEmptyCollector()
	defaultStoreType            = "store_type_unknown"
	defaultPrometheusRegisterer = prometheus.DefaultRegisterer

	// should be used only by GetPrometheusMetricCollector
	prometheusMetricCollector         IMetricCollector
	initPrometheusMetricCollectorOnce sync.Once
)

func SetDefaultMetricCollector(collector IMetricCollector) {
	defaultMetricCollector = collector
}

func SetDefaultPrometheusRegisterer(registerer prometheus.Registerer) {
	defaultPrometheusRegisterer = registerer
}

func GetPrometheusMetricCollector() IMetricCollector {
	initPrometheusMetricCollectorOnce.Do(func() {
		prometheusMetricCollector = NewMetricCollector()
	})
	return prometheusMetricCollector
}
