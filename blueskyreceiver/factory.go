package blueskyreceiver

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

// NewFactory creates a factory for k8s_cluster receiver.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		component.ReceiverType("blueskyreceiver"),
		createDefaultConfig,
		receiver.WithLogs(createLogsReceiver, component.AlphaStabilityLevel),)
}

func createDefaultConfig() component.Config {
	return &Config{
		BGS: "wss://bsky.network"
	}
}

func createLogsReceiver(
	_ context.Context,
	params receiver.CreateSettings,
	cfg component.Config,
	consumer consumer.Logs,
) (receiver.Logs, error) {
	rCfg := cfg.(*Config)

	return newReceiver(params, rCfg, consumer)
}