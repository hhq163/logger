package logger

import (
	"context"
	"fmt"

	"github.com/hhq163/jaeger_helper"
	"github.com/hhq163/logger/core"
	"github.com/opentracing/opentracing-go"
	jLog "github.com/opentracing/opentracing-go/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type MyLogger struct {
	conf *Config
	base *zap.SugaredLogger
	span opentracing.Span
}

func (s *MyLogger) SetLevel(level core.Level) {
	// TODO 关联顶层 config 和内部的底层 config
	s.conf.Level = level
	s.conf.zapBaseConf.Level.SetLevel(zapcore.Level(level))
}

func (s *MyLogger) With(args ...interface{}) Logger {
	newLogger := *s
	newLogger.base = newLogger.base.With(args...)
	return &newLogger
}

func (s *MyLogger) Debug(args ...interface{}) {
	if s.span != nil {
		s.span.LogFields(
			jLog.String("level", "debug"),
			jLog.String("event", fmt.Sprint(args...)),
		)
	}
	s.base.Debug(args...)
}

func (s *MyLogger) Info(args ...interface{}) {
	if s.span != nil {
		s.span.LogFields(
			jLog.String("level", "info"),
			jLog.String("event", fmt.Sprint(args...)),
		)
	}
	s.base.Info(args...)
}

func (s *MyLogger) Warn(args ...interface{}) {
	if s.span != nil {
		s.span.LogFields(
			jLog.String("level", "warn"),
			jLog.String("event", fmt.Sprint(args...)),
		)
	}
	s.base.Warn(args...)
}

func (s *MyLogger) Error(args ...interface{}) {
	if s.span != nil {
		s.span.LogFields(
			jLog.String("level", "error"),
			jLog.String("event", fmt.Sprint(args...)),
		)
	}
	s.base.Error(args...)
}

func (s *MyLogger) DPanic(args ...interface{}) {
	if s.span != nil {
		s.span.LogFields(
			jLog.String("level", "dpanic"),
			jLog.String("event", fmt.Sprint(args...)),
		)
	}
	s.base.DPanic(args...)
}

func (s *MyLogger) Panic(args ...interface{}) {
	if s.span != nil {
		s.span.LogFields(
			jLog.String("level", "panic"),
			jLog.String("event", fmt.Sprint(args...)),
		)
	}
	s.base.Panic(args...)
}

func (s *MyLogger) Fatal(args ...interface{}) {
	if s.span != nil {
		s.span.LogFields(
			jLog.String("level", "fatal"),
			jLog.String("event", fmt.Sprint(args...)),
		)
	}
	s.base.Fatal(args...)
}

func (s *MyLogger) Debugf(template string, args ...interface{}) {
	if s.span != nil {
		s.span.LogFields(
			jLog.String("level", "debug"),
			jLog.String("event", fmt.Sprintf(template, args...)),
		)
	}
	s.base.Debugf(template, args...)
}

func (s *MyLogger) Infof(template string, args ...interface{}) {
	if s.span != nil {
		s.span.LogFields(
			jLog.String("level", "info"),
			jLog.String("event", fmt.Sprintf(template, args...)),
		)
	}
	s.base.Infof(template, args...)
}

func (s *MyLogger) Warnf(template string, args ...interface{}) {
	if s.span != nil {
		s.span.LogFields(
			jLog.String("level", "warn"),
			jLog.String("event", fmt.Sprintf(template, args...)),
		)
	}
	s.base.Warnf(template, args...)

}

func (s *MyLogger) Errorf(template string, args ...interface{}) {
	if s.span != nil {
		s.span.LogFields(
			jLog.String("level", "error"),
			jLog.String("event", fmt.Sprintf(template, args...)),
		)
	}
	s.base.Errorf(template, args...)
}

func (s *MyLogger) DPanicf(template string, args ...interface{}) {
	if s.span != nil {
		s.span.LogFields(
			jLog.String("level", "panic"),
			jLog.String("event", fmt.Sprintf(template, args...)),
		)
	}
	s.base.DPanicf(template, args...)
}

func (s *MyLogger) Panicf(template string, args ...interface{}) {
	if s.span != nil {
		s.span.LogFields(
			jLog.String("level", "panic"),
			jLog.String("event", fmt.Sprintf(template, args...)),
		)
	}
	s.base.Panicf(template, args...)
}

func (s *MyLogger) Fatalf(template string, args ...interface{}) {
	if s.span != nil {
		s.span.LogFields(
			jLog.String("level", "fatal"),
			jLog.String("event", fmt.Sprintf(template, args...)),
		)
	}
	s.base.Fatalf(template, args...)
}

func (s *MyLogger) Debugw(msg string, keysAndValues ...interface{}) {
	if s.span != nil {
		s.span.LogFields(
			jLog.String("level", "debug"),
			jLog.String("event", fmt.Sprintf("%s(%v)", msg, keysAndValues)),
		)
	}
	s.base.Debugw(msg, keysAndValues...)
}

func (s *MyLogger) Infow(msg string, keysAndValues ...interface{}) {
	if s.span != nil {
		s.span.LogFields(
			jLog.String("level", "info"),
			jLog.String("event", fmt.Sprintf("%s(%v)", msg, keysAndValues)),
		)
	}
	s.base.Infow(msg, keysAndValues...)
}

func (s *MyLogger) Warnw(msg string, keysAndValues ...interface{}) {
	if s.span != nil {
		s.span.LogFields(
			jLog.String("level", "warn"),
			jLog.String("event", fmt.Sprintf("%s(%v)", msg, keysAndValues)),
		)
	}
	s.base.Warnw(msg, keysAndValues...)
}

func (s *MyLogger) Errorw(msg string, keysAndValues ...interface{}) {
	if s.span != nil {
		s.span.LogFields(
			jLog.String("level", "error"),
			jLog.String("event", fmt.Sprintf("%s(%v)", msg, keysAndValues)),
		)
	}
	s.base.Errorw(msg, keysAndValues...)
}

func (s *MyLogger) DPanicw(msg string, keysAndValues ...interface{}) {
	if s.span != nil {
		s.span.LogFields(
			jLog.String("level", "panic"),
			jLog.String("event", fmt.Sprintf("%s(%v)", msg, keysAndValues)),
		)
	}
	s.base.DPanicw(msg, keysAndValues...)
}

func (s *MyLogger) Panicw(msg string, keysAndValues ...interface{}) {
	if s.span != nil {
		s.span.LogFields(
			jLog.String("level", "panic"),
			jLog.String("event", fmt.Sprintf("%s(%v)", msg, keysAndValues)),
		)
	}
	s.base.Panicw(msg, keysAndValues...)
}

func (s *MyLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	if s.span != nil {
		s.span.LogFields(
			jLog.String("level", "fatal"),
			jLog.String("event", fmt.Sprintf("%s(%v)", msg, keysAndValues)),
		)
	}
	s.base.Fatalw(msg, keysAndValues...)
}

func (s *MyLogger) Sync() error {
	return s.base.Sync()
}

// Deprecated: please use WithCtx
func (s *MyLogger) WithSpan(ctx context.Context) Logger {
	return s.WithCtx(ctx)
}

func (s *MyLogger) WithCtx(ctx context.Context) Logger {
	var newLogger Logger = s
	if span := jaeger_helper.SpanFromContext(ctx); span != nil {
		sInstance := *s
		sInstance.span = span
		newLogger = &sInstance
	}
	if tID := GetTraceID(ctx); tID != "" {
		newLogger = newLogger.With(TraceID, tID)
	}

	return newLogger
}

func NewDefaultLogger() Logger {
	return NewMyLogger(NewProductionConfig())
}

func NewMyLogger(config *Config) Logger {
	if config == nil {
		config = NewProductionConfig()
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	if config.EncoderConfig.TimeFormat == TimeFormat_ISO8601 {
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}
	if config.EncoderConfig.EncodeCaller == FullPathCaller {
		encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	}

	zapConfig := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.Level(config.Level)),
		Development:       config.Development,
		DisableCaller:     config.DisableCaller,
		DisableStacktrace: config.DisableStacktrace,
		Sampling:          &zap.SamplingConfig{Initial: 100, Thereafter: 100},
		Encoding:          config.Encoding,
		EncoderConfig:     encoderConfig,
		OutputPaths:       config.OutputPaths,
		InitialFields:     config.InitialFields,
	}

	zapLogger, err := zapConfig.Build()
	if err != nil {
		fmt.Println("Build zap logger error!", err)
		return nil
	}
	zapLogger = zapLogger.WithOptions(zap.AddCallerSkip(config.CallerSkip))
	config.zapBaseConf = zapConfig

	// info, ok := debug.ReadBuildInfo()
	// if ok == false {
	// 	fmt.Println("use logger without go module")
	// } else {
	// 	if ReplaceModuleVersion != "" {
	// 		zapLogger = zapLogger.With(zap.String(LabelModuleName, info.Main.Path+" "+ReplaceModuleVersion))
	// 	} else {
	// 		zapLogger = zapLogger.With(zap.String(LabelModuleName, info.Main.Path+" "+info.Main.Version))
	// 	}
	// }

	spkLogger := &MyLogger{
		base: zapLogger.Sugar(),
		conf: config,
	}

	return spkLogger
}
