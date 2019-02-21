package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type AmqpConfig struct {
	Address      string `toml:"address"`
	ExchangeName string `toml:"exchange_name"`
	ExchangeType string `toml:"exchange_type"`
	AutoDelete   bool   `toml:"autoDelete"`
	Durable      bool   `toml:"durable"`
	QueueName    string `toml:"queue_name"`
	ToutingKey   string `toml:"routing_key"`
}

func InitAmqpConfig() *AmqpConfig {
	var amqpConfig *AmqpConfig

	if _, err := toml.DecodeFile("./config/amqp.toml", &amqpConfig); err != nil {
		fmt.Println(err)
		return amqpConfig
	}
	if amqpConfig == nil {
		fmt.Println("配置文件出错！")
		return amqpConfig
	}
	return amqpConfig
}
