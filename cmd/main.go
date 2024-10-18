package main

import (
	"fmt"
	rpcserver "github.com/ricardgo403/rpc-server-demo"
	"github.com/ricardgo403/rpc-server-demo/internal/config"
	"github.com/ricardgo403/rpc-server-demo/internal/zaplogger"
)

func main() {
	logger := zaplogger.NewLogger()
	defer zaplogger.CloseLogger(logger)

	cfg, err := config.LoadConfigsFromEnv(logger)
	if err != nil {
		panic(err)
	}

	logger.Info(cfg)

	rpcserver.Start(logger, cfg.RPCServerCfg)

	var input string
	_, err = fmt.Scanln(&input)
	if err != nil {
		return
	}
}
