package zaplogger

import (
	"go.uber.org/zap"
)

// NewLogger Creates a new instance of Uber Zap
func NewLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	sugarLogger := logger.Sugar()
	return sugarLogger
}

// CloseLogger Flushes any pending logs
func CloseLogger(logger *zap.SugaredLogger) {
	err := logger.Sync()
	if err != nil {
		panic("failed to sync logger: " + err.Error())
	}
}
