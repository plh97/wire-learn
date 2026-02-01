package core

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logEncoder struct {
	zapcore.Encoder
	errFile     *os.File
	file        *os.File
	currentDate string
}

const (
	blueColor   = "\033[34m"
	yellowColor = "\033[33m"
	redColor    = "\033[31m"
	resetColor  = "\033[0m"
)

func myEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var color string
	switch level {
	case zapcore.DebugLevel:
		color = blueColor
	case zapcore.WarnLevel:
		color = yellowColor
	case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		color = redColor
	default:
		color = resetColor
	}
	enc.AppendString(color + level.String() + resetColor)
}

func NewLogger() *zap.Logger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    myEncodeLevel,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(
		consoleEncoder,
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	return zap.New(core)
}
