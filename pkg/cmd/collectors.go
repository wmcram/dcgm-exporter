package cmd

import (
	"bytes"
	"flag"
	"log/slog"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"github.com/urfave/cli/v2"
	"github.com/wmcram/dcgm-exporter/internal/pkg/collector"
	"github.com/wmcram/dcgm-exporter/internal/pkg/dcgmprovider"
	"github.com/wmcram/dcgm-exporter/internal/pkg/hostname"
	"github.com/wmcram/dcgm-exporter/internal/pkg/logging"
	"github.com/wmcram/dcgm-exporter/internal/pkg/nvmlprovider"
	"github.com/wmcram/dcgm-exporter/internal/pkg/prerequisites"
	"github.com/wmcram/dcgm-exporter/internal/pkg/registry"
	"github.com/wmcram/dcgm-exporter/internal/pkg/server"
)

// DcgmCollector is intended for external use as a prometheus.Collector
type DcgmCollector struct {
	s *server.MetricsServer
	r *registry.Registry
}


func (d *DcgmCollector) Describe(ch chan<- *prometheus.Desc) {

}

// Collect grabs metrics from the dcgm-exporter by parsing the buffer from the server's Render method
func (d *DcgmCollector) Collect(ch chan<- prometheus.Metric) {
	metricGroups, err := d.r.Gather()
	if err != nil {
		slog.Error("Error gathering metrics", slog.String(logging.ErrorKey, err.Error()))
		return
	}

	var buf bytes.Buffer
	err = d.s.Render(&buf, metricGroups)
	if err != nil {
		slog.Error("Error rendering metrics", slog.String(logging.ErrorKey, err.Error()))
		return
	}

	decoder := expfmt.NewDecoder(&buf, expfmt.NewFormat(expfmt.TypeTextPlain))
	
	for {
		var mf dto.MetricFamily
		if err := decoder.Decode(&mf); err != nil {
			break
		}

		for _, m := range mf.Metric {
			var labelNames, labelValues []string
			for _, lp := range m.Label {
				labelNames = append(labelNames, lp.GetName())
				labelValues = append(labelValues, lp.GetValue())
			}

			desc := prometheus.NewDesc(mf.GetName(), mf.GetHelp(), labelNames, nil)

			var value float64
			var valueType prometheus.ValueType
			switch mf.GetType() {
			case dto.MetricType_COUNTER:
				value = m.GetCounter().GetValue()
				valueType = prometheus.CounterValue
			case dto.MetricType_GAUGE:
				value = m.GetGauge().GetValue()
				valueType = prometheus.GaugeValue
			case dto.MetricType_UNTYPED:
				value = m.GetUntyped().GetValue()
				valueType = prometheus.UntypedValue
			default:
				slog.Error("Unknown metric type", slog.String("metric", mf.GetName()), slog.String("type", mf.GetType().String()))
				continue
			}

			metric, err := prometheus.NewConstMetric(
				desc,
				valueType,
				value,
				labelValues...
			)
			if err == nil {
				ch <- metric
			}
		}
	}
}


func GetCollector() (*DcgmCollector, error) {
	app := NewApp()
	
	set := flag.NewFlagSet("dcgm", flag.PanicOnError)

	set.String(CLIGPUDevices, "f", "")
	set.String(CLISwitchDevices, "f", "")
	set.String(CLICPUDevices, "f", "")
	set.String(CLIDCGMLogLevel, "NONE", "")
	set.Bool(CLIDumpEnabled, false, "")
	set.String(CLIDumpDirectory, "/tmp/dcgm-exporter-debug", "")
	set.Int(CLIDumpRetention, 24, "")
	set.Bool(CLIDumpCompression, true, "")

	c := cli.NewContext(app, set, nil)

	config, err := contextToConfig(c)
	if err != nil {
		return nil, err
	}

	err = prerequisites.Validate()
	if err != nil {
		return nil, err
	}

	dcgmprovider.Initialize(config)
	dcgmCleanup := dcgmprovider.Client().Cleanup

	nvmlprovider.Initialize()
	nvmlCleanup := nvmlprovider.Client().Cleanup

	slog.Info("DCGM successfully initialized!")
	slog.Info("NVML provider successfully initialized!")

	fillConfigMetricGroups(config)

	cs := getCounters(config)

	deviceWatchListManager := startDeviceWatchListManager(cs, config)

	hostname, err := hostname.GetHostname(config)
	if err != nil {
		nvmlCleanup()
		dcgmCleanup()
		return nil, err
	}

	cf := collector.InitCollectorFactory(cs, deviceWatchListManager, hostname, config)
	r := registry.NewRegistry()
	for _, entityCollector := range cf.NewCollectors() {
		r.Register(entityCollector)
	}

	ch := make(chan string, 10)
	server, _, err := server.NewMetricsServer(config, ch, deviceWatchListManager, r)
	if err != nil {
		return nil, err
	}

	return &DcgmCollector{server, r}, nil
}
