package metrics

type Metrics interface{}

type Counter interface {
	Inc()
	Add(value float64)
	With(labelValues ...string) Counter
}

type Gauge interface {
	Set(value float64)
	Add(value float64)
	Sub(value float64)
	With(labValues ...string) Gauge
}

type Histogram interface {
	With(labelValues ...string) Histogram
	Observe(value float64)
}

type Summary interface {
	With(labValues ...string) Summary
	Observe(value float64)
}
