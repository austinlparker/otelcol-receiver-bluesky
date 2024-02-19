package blueskyreceiver

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

type blueskyreceiver struct {
	config          *Config

// newReceiver creates the Kubernetes events receiver with the given configuration.
func newReceiver(
	set receiver.CreateSettings,
	config *Config,
	consumer consumer.Logs,
) (receiver.Logs, error) {

	return &blueskyreceiver{
		config:       config,
	}, nil
}

func (kr *blueskyreceiver) Start(ctx context.Context, _ component.Host) error {
	

	return nil
}

func (kr *blueskyreceiver) Shutdown(context.Context) error {
	
	return nil
}