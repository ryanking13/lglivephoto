package lglivephoto

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger
var Atom zap.AtomicLevel

func init() {
	Atom = zap.NewAtomicLevel()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = ""

	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),
		Atom,
	))
	defer logger.Sync()

	Atom.SetLevel(zap.InfoLevel)
	Logger = logger.Sugar()
}
