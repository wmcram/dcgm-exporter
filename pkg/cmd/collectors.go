package cmd

import (
	"log/slog"

	"github.com/urfave/cli/v2"
	"github.com/wmcram/dcgm-exporter/internal/pkg/collector"
	"github.com/wmcram/dcgm-exporter/internal/pkg/dcgmprovider"
	"github.com/wmcram/dcgm-exporter/internal/pkg/hostname"
	"github.com/wmcram/dcgm-exporter/internal/pkg/nvmlprovider"
	"github.com/wmcram/dcgm-exporter/internal/pkg/prerequisites"
)

func GetCollectors() ([]collector.EntityCollectorTuple, error) {
	app := NewApp()
	c := cli.NewContext(app, nil, nil)


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
	return cf.NewCollectors(), nil
}