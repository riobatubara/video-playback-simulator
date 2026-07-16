package utils

import (
	"os"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = ""
	encoderConfig.LevelKey = ""
	encoderConfig.CallerKey = ""
	encoderConfig.MessageKey = "msg"

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.Lock(zapcore.AddSync(zapcore.Lock(os.Stderr))), // Thread-safe standard error output
		zap.InfoLevel,
	)

	logger = zap.New(core)
}

// LogEvent preserves your exact signature and output format
func LogEvent(index int, tsclient int64, sessionID, event, value string) {
	// Constructing the exact string pattern using fast string interpolation
	// data[%s]:: tsclient: %s, sessid: %s, event: %s, value: %s
	msg := "data[" + strconv.Itoa(index) + "]:: tsclient: " + strconv.FormatInt(tsclient, 10) +
		", sessid: " + sessionID + ", event: " + event + ", value: " + value

	logger.Info(msg)
}
