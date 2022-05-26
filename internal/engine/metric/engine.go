package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	EngineModelEvaluateCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: METRIC_NS,
			Subsystem: "engine",
			Name:      "model_evaluate_count",
			Help:      "风控模型评估次数",
		}, []string{"guid"})
)
