package rpcserver

import (
	"github.com/ricardgo403/rpc-server-demo/internal/config"
	"github.com/ricardgo403/rpc-server-demo/internal/zaplogger"
	"go.uber.org/zap"
	"net"
	"net/rpc"
	"strconv"
	"strings"
)

type RPCServer struct {
	logger *zap.SugaredLogger
	cfg    config.RPCServerCfg
}

func (serverPtr *RPCServer) Negate(i int64, reply *int64) error {
	serverPtr.logger.Info("Negate called")
	*reply = -i
	return nil
}

func NewRPCServer(logger *zap.SugaredLogger, cfg config.RPCServerCfg) *RPCServer {
	if logger == nil {
		logger = zaplogger.NewLogger()
	}
	return &RPCServer{
		logger: logger,
		cfg:    cfg,
	}
}

func (serverPtr *RPCServer) Register() error {
	err := rpc.Register(serverPtr)
	if err != nil {
		return err
	}
	return nil
}

func (serverPtr *RPCServer) Serve() error {
	portStr := strconv.Itoa(serverPtr.cfg.Port)
	serverAddr := strings.Join([]string{serverPtr.cfg.Addr, ":", portStr}, "")
	ln, err := net.Listen(serverPtr.cfg.Type, serverAddr)
	if err != nil {
		serverPtr.logger.Error(err)
		return err
	}
listenerLoop:
	for {
		c, err := ln.Accept()
		if err != nil {
			serverPtr.logger.Error(err)
			break listenerLoop
		}
		go rpc.ServeConn(c)
	}
	return nil
}
