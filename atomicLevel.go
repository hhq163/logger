package logger

import (
	"logger/v3/core"
	"go.uber.org/atomic"
)

const (
	DebugLevel  = core.DebugLevel
	InfoLevel   = core.InfoLevel
	WarnLevel   = core.WarnLevel
	ErrorLevel  = core.ErrorLevel
	DPanicLevel = core.DPanicLevel
	PanicLevel  = core.PanicLevel
	FatalLevel  = core.FatalLevel
)

type AtomicLevel struct {
	l *atomic.Int32
}

func NewAtomicLevel() AtomicLevel {
	return AtomicLevel{
		l: atomic.NewInt32(int32(InfoLevel)),
	}
}

func NewAtomicLevelAt(l core.Level) AtomicLevel {
	a := NewAtomicLevel()
	a.SetLevel(l)
	return a
}

func (lvl AtomicLevel) Level() core.Level {
	return core.Level(int8(lvl.l.Load()))
}

func (lvl AtomicLevel) SetLevel(l core.Level) {
	lvl.l.Store(int32(l))
}

func (lvl AtomicLevel) String() string {
	return lvl.Level().String()
}

func (lvl *AtomicLevel) UnmarshalText(text []byte) error {
	if lvl.l == nil {
		lvl.l = &atomic.Int32{}
	}

	var l core.Level
	if err := l.UnmarshalText(text); err != nil {
		return err
	}

	lvl.SetLevel(l)
	return nil
}

func (lvl AtomicLevel) MarshalText() (text []byte, err error) {
	return lvl.Level().MarshalText()
}
