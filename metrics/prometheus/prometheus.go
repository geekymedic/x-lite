package prometheus

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"git.gmtshenzhen.com/geeky-medic/x-lite/metrics"
	"git.gmtshenzhen.com/geeky-medic/x-lite/version"
)

const (
	NameSpacePrometheus = "geekymedic"
	SubSystemPrometheus = ""
)

var sysLabs = struct {
	nameSpace string
	subSystem string
}{nameSpace: "demo_system", subSystem: "bff_demo"}

func init() {
	idx := strings.Index(version.PRONAME, "system-")
	s := len("system-")
	if idx > 0 {
		idx = idx + s
		sysLabs.nameSpace = strings.ReplaceAll(version.PRONAME[0:idx-1], "-", "_")
		sysLabs.subSystem = strings.ReplaceAll(version.PRONAME[idx:], "-", "_")
	}
}

func StartMetricsServer(addr, path string) error {
	http.Handle(path, promhttp.Handler())
	return http.ListenAndServe(addr, nil)
}

type Counter struct {
	cv  *prometheus.CounterVec
	lvs []string
}

func MustCounter(name string) *Counter {
	return MustCounterWithLabelNames(name)
}

func NewCounter(name string) (*Counter, error) {
	return NewCounterWithLabelNames(name)
}

func MustCounterWithLabelNames(name string, labelNames ...string) *Counter {
	counter, err := NewCounterWithLabelNames(name, labelNames...)
	if err != nil {
		panic(fmt.Sprintf("fail to register Counter(%s_%s_%s) into promethus component: %v", NameSpacePrometheus, SubSystemPrometheus, name, err))
	}
	return counter
}

func NewCounterWithLabelNames(name string, labelNames ...string) (*Counter, error) {
	opts := prometheus.CounterOpts(prometheus.Opts{
		Namespace: NameSpacePrometheus,
		Subsystem: SubSystemPrometheus,
		Name:      name})
	var proLabels = prometheus.Labels{}
	labelNames = append(labelNames, "pro_name")
	var lvs = make([]string, len(labelNames)*2, len(labelNames)*2)
	i := 0
	for _, lab := range labelNames {
		proLabels[lab] = ""
		lvs[i] = lab
		lvs[i+1] = "unkowned"
		i += 2
	}
	counter := &Counter{cv: prometheus.NewCounterVec(opts, labelNames), lvs: lvs}
	err := prometheus.Register(counter.cv)
	if err != nil {
		return nil, err
	}
	return counter, nil
}

func (counter *Counter) Add(delta float64) {
	counter.cv.With(makeLabelsKeyValues(counter.lvs...)).Add(delta)
}

func (counter *Counter) Inc() {
	counter.cv.With(makeLabelsKeyValues(counter.lvs...)).Inc()
}

func (counter *Counter) With(labValues ...string) metrics.Counter {
	labValues = append(labValues, fmt.Sprintf("%s_%s", sysLabs.nameSpace, sysLabs.subSystem))
	return &Counter{cv: counter.cv, lvs: fillLabels(counter.lvs, labValues...)}
}

type Gauge struct {
	gv  *prometheus.GaugeVec
	lvs []string
}

func MustGague(name string) *Gauge {
	return MustGagueWithLabelNames(name)
}

func NewGague(name string) (*Gauge, error) {
	return NewGagueWithLabelNames(name)
}

func MustGagueWithLabelNames(name string, labelNames ...string) *Gauge {
	counter, err := NewGagueWithLabelNames(name, labelNames...)
	if err != nil {
		panic(fmt.Sprintf("fail to register Gague(%s_%s_%s) into promethus component: %v", NameSpacePrometheus, SubSystemPrometheus, name, err))
	}
	return counter
}

func NewGagueWithLabelNames(name string, labelNames ...string) (*Gauge, error) {
	opts := prometheus.GaugeOpts(prometheus.Opts{
		Namespace: NameSpacePrometheus,
		Subsystem: SubSystemPrometheus,
		Name:      name})
	var proLabels = prometheus.Labels{}
	labelNames = append(labelNames, "pro_name")
	var lvs = make([]string, len(labelNames)*2)
	i := 0
	for _, lab := range labelNames {
		proLabels[lab] = ""
		lvs[i] = lab
		lvs[i+1] = "unkowned"
		i += 2
	}

	gague := &Gauge{gv: prometheus.NewGaugeVec(opts, labelNames), lvs: lvs}
	err := prometheus.Register(gague.gv)
	if err != nil {
		return nil, err
	}
	return gague, nil
}

func (gague *Gauge) Set(value float64) {
	gague.gv.With(makeLabelsKeyValues(gague.lvs...)).Set(value)
}

func (gague *Gauge) Add(delta float64) {
	gague.gv.With(makeLabelsKeyValues(gague.lvs...)).Add(delta)
}

func (gague *Gauge) Sub(value float64) {
	gague.gv.With(makeLabelsKeyValues(gague.lvs...)).Sub(value)
}

func (gague *Gauge) With(labValues ...string) metrics.Gauge {
	labValues = append(labValues, fmt.Sprintf("%s_%s", sysLabs.nameSpace, sysLabs.subSystem))
	return &Gauge{gv: gague.gv, lvs: fillLabels(gague.lvs, labValues...)}
}

type Histogram struct {
	hv  *prometheus.HistogramVec
	lvs []string
}

func MustHistogram(name string, buckets []float64) *Histogram {
	return MustHistogramWithLabelNames(name, buckets)
}

func NewHistogram(name string, buckets []float64) (*Histogram, error) {
	return NewHistogramWithLabelNames(name, buckets)
}

func MustHistogramWithLabelNames(name string, buckets []float64, labelNames ...string) *Histogram {
	hv, err := NewHistogramWithLabelNames(name, buckets, labelNames...)
	if err != nil {
		panic(fmt.Sprintf("fail to register Histogram(%s_%s_%s) into promethus component: %v", NameSpacePrometheus, SubSystemPrometheus, name, err))
	}
	return hv
}

func NewHistogramWithLabelNames(name string, buckets []float64, labelNames ...string) (*Histogram, error) {
	opts := prometheus.HistogramOpts{
		Namespace: NameSpacePrometheus,
		Subsystem: SubSystemPrometheus,
		Name:      name,
		Buckets:   buckets}
	var proLabels = prometheus.Labels{}
	labelNames = append(labelNames, "pro_name")
	var lvs = make([]string, len(labelNames)*2)
	i := 0
	for _, lab := range labelNames {
		proLabels[lab] = ""
		lvs[i] = lab
		lvs[i+1] = "unkowned"
		i += 2
	}

	hv := &Histogram{hv: prometheus.NewHistogramVec(opts, labelNames), lvs: lvs}
	err := prometheus.Register(hv.hv)
	if err != nil {
		return nil, err
	}
	return hv, nil
}

func (h *Histogram) Observe(value float64) {
	h.hv.With(makeLabelsKeyValues(h.lvs...)).Observe(value)
}

func (h *Histogram) With(labValues ...string) *Histogram {
	labValues = append(labValues, fmt.Sprintf("%s_%s", sysLabs.nameSpace, sysLabs.subSystem))
	return &Histogram{hv: h.hv, lvs: fillLabels(h.lvs, labValues...)}
}

type Summary struct {
	s   *prometheus.SummaryVec
	lvs []string
}

func MustSummay(name string, objectives map[float64]float64) *Summary {
	return MustSummaryWithLabelNames(name, objectives)
}

func NewSummary(name string, objectives map[float64]float64) (*Summary, error) {
	return NewSummaryWithLabelNames(name, objectives)
}

func MustSummaryWithLabelNames(name string, objectives map[float64]float64, labelNames ...string) *Summary {
	s, err := NewSummaryWithLabelNames(name, objectives, labelNames...)
	if err != nil {
		panic(fmt.Sprintf("fail to register Histogram(%s_%s_%s) into promethus component: %v", NameSpacePrometheus, SubSystemPrometheus, name, err))
	}
	return s
}

func NewSummaryWithLabelNames(name string, objectives map[float64]float64, labelNames ...string) (*Summary, error) {
	opts := prometheus.SummaryOpts{
		Namespace:  NameSpacePrometheus,
		Subsystem:  SubSystemPrometheus,
		Objectives: objectives,
		Name:       name}
	var proLabels = prometheus.Labels{}
	labelNames = append(labelNames, "pro_name")
	var lvs = make([]string, len(labelNames)*2)
	i := 0
	for _, lab := range labelNames {
		proLabels[lab] = ""
		lvs[i] = lab
		lvs[i+1] = "unkowned"
		i += 2
	}

	s := &Summary{s: prometheus.NewSummaryVec(opts, labelNames), lvs: lvs}
	err := prometheus.Register(s.s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Summary) With(labValues ...string) *Summary {
	labValues = append(labValues, fmt.Sprintf("%s_%s", sysLabs.nameSpace, sysLabs.subSystem))
	return &Summary{s: s.s, lvs: fillLabels(s.lvs, labValues...)}
}

func (s *Summary) Observe(value float64) {
	s.s.With(makeLabelsKeyValues(s.lvs...)).Observe(value)
}

func fillLabels(lvs []string, labelsValue ...string) []string {
	var (
		min       = len(lvs) / 2
		newLabels = make([]string, len(lvs))
	)
	if len(labelsValue) < min {
		for i := 0; i < min-len(labelsValue); i++ {
			labelsValue = append(labelsValue, "unkowned")
		}
	}
	for i := 0; i < min; i++ {
		labelKey := lvs[i*2]
		newLabels[i*2] = labelKey
		newLabels[i*2+1] = labelsValue[i]
	}
	return newLabels
}

func makeLabelsKeyValues(lvs ...string) prometheus.Labels {
	labels := prometheus.Labels{}
	for i := 0; i < len(lvs); i += 2 {
		labels[lvs[i]] = lvs[i+1]
	}
	return labels
}
