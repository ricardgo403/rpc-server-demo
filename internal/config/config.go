package config

import (
	"github.com/caarlos0/env/v11"
	"go.uber.org/zap"
)

type Config struct {
	RPCServerCfg RPCServerCfg           `envPrefix:"RPC_SERVER"`
	configs      map[string]interface{} `env:"configs"  envDefault:""`
}

type RPCServerCfg struct {
	Type string `env:"TYPE"  envDefault:"tcp"`
	Addr string `env:"ADDR"  envDefault:"127.0.0.1"`
	Port int    `env:"PORT"  envDefault:"9999"`
}

func NewConfig() *Config {
	return &Config{
		RPCServerCfg: RPCServerCfg{},
		configs:      make(map[string]interface{}),
	}
}

func LoadConfigsFromEnv(logger *zap.SugaredLogger) (*Config, error) {

	var cfg = NewConfig()
	if err := env.Parse(cfg); err != nil {
		logger.Error("Failed to parse configs: ", err)
	}
	logger.Infof("Configs: %+v", cfg)

	return cfg, nil
}
