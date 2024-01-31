package logger

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logg *zap.Logger
var logFile *os.File

var ErrLog *log.Logger
var InfoLog *log.Logger

func init() {
	var err error

	ErrLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	logFile, err = os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		ErrLog.Printf("Error opening log file: %+v", err)
	}

	writers := []zapcore.WriteSyncer{
		zapcore.AddSync(os.Stdout),
		zapcore.AddSync(logFile),
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stackTrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writers...),
		zap.NewAtomicLevelAt(zap.InfoLevel),
	)

	logg = zap.New(core, zap.AddCallerSkip(1))
	defer logg.Sync()

	zap.ReplaceGlobals(logg)

	if err != nil {
		ErrLog.Printf("Error initializing logger: %v", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()
}

func cleanup() {
	if logFile != nil {
		logFile.Close()
	}
}

func Info(message string, fields ...zap.Field) {
	logg.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	logg.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	logg.Error(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	logg.Warn(message, fields...)
}
