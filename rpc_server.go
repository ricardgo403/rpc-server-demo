package rpcserver

import (
	"github.com/ricardgo403/rpc-server-demo/internal/config"
	rpcserver "github.com/ricardgo403/rpc-server-demo/internal/server"
	"go.uber.org/zap"
)

func Start(logger *zap.SugaredLogger, cfg config.RPCServerCfg) {
	server := rpcserver.NewRPCServer(logger, cfg)
	err := server.Register()
	if err != nil {
		logger.Errorf("Failed to register RPC server: %s", err)
		return
	}

	err = server.Serve()
	if err != nil {
		logger.Errorf("Failed to start RPC server: %s", err)
		return
	}

}
