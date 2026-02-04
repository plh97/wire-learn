package core

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/google/wire"
	"github.com/plh97/wire-learn/model"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

type myWriter struct {
	file  *os.File
	mutex sync.Mutex
	currentDate string
}

func (m *myWriter) Write(b []byte) (n int, err error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	now := time.Now().Format("2006-01-02")
	if m.currentDate == now {
		return m.file.Write(b)
	}
	if m.file != nil {
		m.file.Close()
	}
	os.MkdirAll("logs", 0755)
	fileName := fmt.Sprintf("logs/log_%s.log", now)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	m.file = file
	m.currentDate = now
	return file.Write(b)
}

var CommonProvider = wire.NewSet(NewDB, NewLogger)

func NewDB() *gorm.DB {
	DB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&model.User{}, &model.Video{})

	return DB
}

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
)

func EncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	// Custom level encoder to add colors
	var color string
	switch level {
	case zapcore.DebugLevel:
		color = Cyan
	case zapcore.InfoLevel:
		color = Green
	case zapcore.WarnLevel:
		color = Yellow
	case zapcore.ErrorLevel:
		color = Red
	case zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		color = Purple
	default:
		color = Blue
	}
	enc.AppendString(fmt.Sprintf("%s%s\033[m", color, level.String()))
}

func NewLogger() *zap.Logger {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = EncodeLevel
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		zapcore.NewMultiWriteSyncer(os.Stdout, zapcore.AddSync(&myWriter{})),
		zapcore.InfoLevel,
		// fileCore,
		// consoleCore,
	)
	logger := zap.New(core, zap.AddCaller())
	for i:=0; i<10; i++ {
		logger.Sugar().Info("Logger initialized %d/10", i+1)
		time.Sleep(time.Second)
	}
	return logger
}
