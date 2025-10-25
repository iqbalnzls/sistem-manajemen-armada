package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `json:"app" mapstructure:"app"`
	Database DatabaseConfig `json:"database" mapstructure:"database"`
	RabbitMQ RabbitMqConfig `json:"rabbitmq" mapstructure:"rabbitmq"`
	MQTT     MQTTConfig     `json:"mqtt" mapstructure:"mqtt"`
}

type AppConfig struct {
	Name string `json:"name" mapstructure:"name"`
	Port int32  `json:"port" mapstructure:"port"`
}
type DatabaseConfig struct {
	Host               string `json:"host" mapstructure:"host"`
	Username           string `json:"username" mapstructure:"username"`
	Password           string `json:"password" mapstructure:"password"`
	Port               int32  `json:"port" mapstructure:"port"`
	Name               string `json:"name" mapstructure:"name"`
	Schema             string `json:"schema" mapstructure:"schema"`
	MaxIdleConnections int    `json:"maxIdleConnections" mapstructure:"maxIdleConnections"`
	MaxOpenConnections int    `json:"maxOpenConnections" mapstructure:"maxOpenConnections"`
	DebugMode          bool   `json:"debugMode" mapstructure:"debugMode"`
}

type RabbitMqConfig struct {
	Host     string         `json:"host" mapstructure:"host"`
	Port     int32          `json:"port" mapstructure:"port"`
	User     string         `json:"user" mapstructure:"user"`
	Pass     string         `json:"pass" mapstructure:"pass"`
	Exchange ExchangeConfig `json:"exchange" mapstructure:"exchange"`
	Queue    QueueConfig    `json:"queue" mapstructure:"queue"`
}

type ExchangeConfig struct {
	Name       string `json:"name" mapstructure:"name"`
	Type       string `json:"type" mapstructure:"type"`
	Durable    bool   `json:"durable" mapstructure:"durable"`
	AutoDelete bool   `json:"autoDelete" mapstructure:"autoDelete"`
}

type QueueConfig struct {
	Name       string `json:"name" mapstructure:"name"`
	Durable    bool   `json:"durable" mapstructure:"durable"`
	AutoDelete bool   `json:"autoDelete" mapstructure:"autoDelete"`
	Exclusive  bool   `json:"exclusive" mapstructure:"exclusive"`
	RoutingKey string `json:"routingKey" mapstructure:"routingKey"`
}

type MQTTConfig struct {
	Broker               string `json:"broker" mapstructure:"broker"`
	Port                 int    `json:"port" mapstructure:"port"`
	ClientID             string `json:"clientId" mapstructure:"clientId"`
	Username             string `json:"username" mapstructure:"username"`
	Password             string `json:"password" mapstructure:"password"`
	TopicPattern         string `json:"topicPattern" mapstructure:"topicPattern"`
	QoS                  byte   `json:"qos" mapstructure:"qos"`
	CleanSession         bool   `json:"cleanSession" mapstructure:"cleanSession"`
	AutoReconnect        bool   `json:"autoReconnect" mapstructure:"autoReconnect"`
	ConnectRetryDelay    int    `json:"connectRetryDelay" mapstructure:"connectRetryDelay"`
	MaxReconnectAttempts int    `json:"maxReconnectAttempts" mapstructure:"maxReconnectAttempts"`
}

func NewConfig(configPath string) *Config {
	fmt.Println("Load Config.... ")

	viper.SetConfigFile(configPath)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return &config
}

func (c *Config) AppAddress() string {
	return fmt.Sprintf(":%v", c.App.Port)
}
