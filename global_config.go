package mgcache

import "sync"

var (
	defaultMetricCollector = NewEmptyCollector()
	defaultStoreType       = "store_type_unknown"

	prometheusMetricCollector         = NewMetricCollector()
	initPrometheusMetricCollectorOnce sync.Once
)

func SetDefaultMetricCollector(collector IMetricCollector) {
	defaultMetricCollector = collector
}

func GetPrometheusMetricCollector() IMetricCollector {
	initPrometheusMetricCollectorOnce.Do(func() {
		prometheusMetricCollector = NewMetricCollector()
	})
	return prometheusMetricCollector
}
