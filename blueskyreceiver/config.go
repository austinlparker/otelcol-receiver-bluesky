package blueskyreceiver

// Config defines configuration for kubernetes events receiver.
type Config struct {
	BGS string `mapstructure:"bgs"`
}
