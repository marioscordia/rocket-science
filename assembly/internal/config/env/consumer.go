package env

import (
	"github.com/IBM/sarama"
	"github.com/caarlos0/env/v11"
)

type consumerEnvConfig struct {
	Topic   string `env:"UFO_RECORDED_TOPIC_NAME,required"`
	GroupID string `env:"UFO_RECORDED_CONSUMER_GROUP_ID,required"`
}

type consumerConfig struct {
	raw *consumerEnvConfig
}

func NewConsumerConfig() (*consumerConfig, error) {
	var raw consumerEnvConfig
	if err := env.Parse(&raw); err != nil {
		return nil, err
	}
	return &consumerConfig{raw: &raw}, nil
}

func (cfg *consumerConfig) Topic() string {
	return cfg.raw.Topic
}

func (cfg *consumerConfig) GroupID() string {
	return cfg.raw.GroupID
}

func (cfg *consumerConfig) Config() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V4_1_0_0
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	return config
}
