// +build !unittest

package logger

import (
	"context"
	"testing"

	"go.uber.org/zap"
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
}

func TestNewMyLogger(t *testing.T) {

	cfg := logger.NewDevelopmentConfig()
	cfg.OutputPaths = append(cfg.OutputPaths, "access_log.txt")//指定输出文件名
	l = logger.NewMyLogger(cfg)

	l.Info("customary prod info test")
	l.Infof("customary prod infof %s", "test")
	l.Infow("customary prod infow", "k", "test")


	devConf := NewDevelopmentConfig(SvcName, "TestNewMyLogger")
	devConf.Encoding = "json"		//指定输出格式为json，key=>val有利于ELK搜集
	devConf.OutputPaths = append(cfg.OutputPaths, "access_log.txt")//指定输出文件名
	l2 := NewMyLogger(devConf)
	l2.Info("key1","customary dev info test")
}

func TestChangeLevel(t *testing.T) {
	conf := NewProductionConfig(SvcName, "TestChangeLevel")
	conf.Level = DebugLevel
	l := NewMyLogger(conf)
	l.Debug("必须输出")
}

func TestDynamicChangeLevel(t *testing.T) {
	conf := NewProductionConfig(SvcName, "TestDynamicChangeLevel")
	l := NewMyLogger(conf)
	l.Debug("test debug 1")
	l.Info("level:", conf.Level)

	l.SetLevel(DebugLevel)
	l.Debug("test debug 2")
	l.Info("level:", conf.Level)

}

func TestWithSpan(t *testing.T) {
	conf := NewProductionConfig(SvcName, "TestWithSpan")
	l := NewMyLogger(conf)
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
