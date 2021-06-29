package metrics

import (
	"github.com/prometheus/client_golang/prometheus"

	httpMetrics "github.com/storybuilder/storybuilder/channels/http/metrics"
)

// Register registers declared metrics.
//
// This is the central location to register metrics from different
// layers of the service.
func Register() {
	prometheus.MustRegister(httpMetrics.HTTPReqDuration)
}
