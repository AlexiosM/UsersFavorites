package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Log *zap.Logger
)

func init() {
	logConfig := zap.Config{
		Level:    zap.NewAtomicLevel(),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths: []string{"stdout"},
		//ErrorOutputPaths: []string{},
		InitialFields: map[string]interface{}{},
	}
	var err error
	Log, err = logConfig.Build()
	if err != nil {
		panic(err)
	}
}
