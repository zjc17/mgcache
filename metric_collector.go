package mgcache

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespaceCache = "mgcache"
)

var (
	// gaugeVecLabelSet set value into cache
	gaugeVecLabelSet = prometheus.Labels{"type": "set"}
	// gaugeVecLabelHit get with cache hit
	gaugeVecLabelHit = prometheus.Labels{"type": "hit"}
	// gaugeVecLabelMiss get with cache miss
	gaugeVecLabelMiss = prometheus.Labels{"type": "miss"}
)

const (
	metricValueCacheHit  = "hit"
	metricValueCacheMiss = "miss"
	metricValueCacheSet  = "set"
)

type (

	// IMetricCollector is a wrapper for prometheus.Collector
	IMetricCollector interface {
		CacheHit(serviceID string, storeType string)
		CacheMiss(serviceID string, storeType string)
		CacheSet(serviceID string, storeType string)
	}

	metricCollector struct {
		cacheBehaviorMetric *prometheus.GaugeVec
	}

	emptyCollector struct {
	}
)

func NewMetricCollector() IMetricCollector {

	cacheBehaviorMetric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "mgcache_behavior",
			Namespace: namespaceCache,
			Help:      "This represent the number of items in cache",
		},
		[]string{"service", "store", "metric"},
	)
	defaultPrometheusRegisterer.MustRegister(cacheBehaviorMetric)

	return &metricCollector{
		cacheBehaviorMetric: cacheBehaviorMetric,
	}
}

func NewEmptyCollector() IMetricCollector {
	return &emptyCollector{}
}

func (m metricCollector) CacheHit(serviceID string, storeType string) {
	m.cacheBehaviorMetric.WithLabelValues(serviceID, storeType, metricValueCacheHit).Inc()
}

func (m metricCollector) CacheMiss(serviceID string, storeType string) {
	m.cacheBehaviorMetric.WithLabelValues(serviceID, storeType, metricValueCacheMiss).Inc()
}

func (m metricCollector) CacheSet(serviceID string, storeType string) {
	m.cacheBehaviorMetric.WithLabelValues(serviceID, storeType, metricValueCacheSet).Inc()
}

func (e emptyCollector) CacheHit(serviceID string, storeType string) {
}

func (e emptyCollector) CacheMiss(serviceID string, storeType string) {
}

func (e emptyCollector) CacheSet(serviceID string, storeType string) {
}
