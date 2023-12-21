package logger

import "go.uber.org/zap"

func NewLogger() *zap.SugaredLogger {
	log, err := zap.NewProduction()

	defer log.Sync()

	if err != nil {
		panic(err)
	}

	return log.Sugar()
}
