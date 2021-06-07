package lglivephoto

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger
var atom zap.AtomicLevel

func init() {
	atom = zap.NewAtomicLevel()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = ""

	_logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),
		atom,
	))
	defer _logger.Sync()

	atom.SetLevel(zap.InfoLevel)
	logger = _logger.Sugar()
}
