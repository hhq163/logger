// +build !unittest

package logger

import (
	"context"
	"go.uber.org/zap"
	"testing"
)

func TestGlobalDefaultLogger(t *testing.T) {
	Debug("test")
	Debugf("%s", "test")
	Debugw("", "k", "test")
	Info("test")
	Infof("%s", "test")
	Infow("", "k", "test")
	Warn("test")
	Warnf("%s", "test")
	Warnw("", "k", "test")
	Error("test")
	Errorf("%s", "test")
	Errorw("", "k", "test")

	//DPanic("test")
	//DPanicf("%s", "test")
	//DPanicw("", "k", "test")
	//Panic("test")
	//Panicf("%s", "test")
	//Panicw("", "k", "test")
	//Fatal("test")
	//Fatalf("%s", "test")
	//Fatalw("", "k", "test")

	Debug("默认级别，不会输出")
	SetLevel(DebugLevel)
	Debug("修改级别后的输出")
}

func TestNewDefaultLogger(t *testing.T) {
	l := NewDefaultLogger()
	l.Debug("test")
	l.Debugf("%s", "test")
	l.Debugw("", "k", "test")
	l.Info("test")
	l.Infof("%s", "test")
	l.Infow("", "k", "test")
	l.Warn("test")
	l.Warnf("%s", "test")
	l.Warnw("", "k", "test")
	l.Error("test")
	l.Errorf("%s", "test")
	l.Errorw("", "k", "test")
	l.DPanic("test")
	l.DPanicf("%s", "test")
	l.DPanicw("", "k", "test")
	//l.Panic("test")
	//l.Panicf("%s", "test")
	//l.Panicw("", "k", "test")
	//l.Fatal("test")
	//l.Fatalf("%s", "test")
	//l.Fatalw("", "k", "test")
}

func TestNewSpeakinLogger(t *testing.T) {

	prodConf := NewDevelopmentConfig(SvcName, "TestNewSpeakinLogger")
	l := NewSpeakinLogger(prodConf)
	l.Info("customary prod info test")
	l.Infof("customary prod infof %s", "test")
	l.Infow("customary prod infow", "k", "test")

	devConf := NewDevelopmentConfig(SvcName, "TestNewSpeakinLogger")
	l2 := NewSpeakinLogger(devConf)
	l2.Info("customary dev info test")
	l2.Infof("customary dev infof %s", "test")
	l2.Infow("customary dev infow", "k", "test")
}

func TestChangeLevel(t *testing.T) {
	conf := NewProductionConfig(SvcName, "TestChangeLevel")
	conf.Level = DebugLevel
	l := NewSpeakinLogger(conf)
	l.Debug("必须输出")
}

func TestDynamicChangeLevel(t *testing.T) {
	conf := NewProductionConfig(SvcName, "TestDynamicChangeLevel")
	l := NewSpeakinLogger(conf)
	l.Debug("test debug 1")
	l.Info("level:", conf.Level)

	l.SetLevel(DebugLevel)
	l.Debug("test debug 2")
	l.Info("level:", conf.Level)

}

func TestWithSpan(t *testing.T) {
	conf := NewProductionConfig(SvcName, "TestWithSpan")
	l := NewSpeakinLogger(conf)
	l.Info("没有 span 时输出")
	l = l.WithSpan(context.Background())
	l.Info("有 span 时输出")
}

func TestZapLogger(t *testing.T) {
	conf := zap.NewProductionConfig()
	l, _ := conf.Build()
	l.Debug("test debug 1")
	conf.Level.SetLevel(zap.DebugLevel)
	l.Debug("test debug 2")

	l.Core().Enabled(zap.InfoLevel)
	l.Debug("test debug 3")

}

func TestCommonLogger(t *testing.T) {
	Debug("debug test, will not be print")
	Info("info test")

	config := NewDevelopmentConfig()
	SetConfig(config)

	Debug("debug test")
	Error("error test")

	config.DisableCaller = true
	config.DisableStacktrace = true
	SetConfig(config)
	Error("error test without stacktrace and caller")
}

func TestWithCtx(t *testing.T) {
	ctx := context.WithValue(context.Background(), TraceID, "123")
	l.WithCtx(ctx).Info("test context with trace")
}
