package logger

import (
	"github.com/hhq163/logger/core"
	"go.uber.org/zap"
)

const _oddNumberErrMsg = "Ignored key without a value."

type Config struct {
	// Level is the minimum enabled logging level. Note that this is a dynamic
	// level, so calling Config.Level.SetLevel will atomically change the log
	// level of all loggers descended from this config.
	Level core.Level `json:"level" yaml:"level"`
	// Development puts the logger in development mode, which changes the
	// behavior of DPanicLevel and takes stacktraces more liberally.
	Development bool `json:"development" yaml:"development"`
	// DisableCaller stops annotating logs with the calling function's file
	// name and line number. By default, all logs are annotated.
	DisableCaller bool `json:"disableCaller" yaml:"disableCaller"`
	// DisableStacktrace completely disables automatic stacktrace capturing. By
	// default, stacktraces are captured for WarnLevel and above logs in
	// development and ErrorLevel and above in production.
	DisableStacktrace bool `json:"disableStacktrace" yaml:"disableStacktrace"`
	// Encoding sets the logger's encoding. Valid values are "json" and "console"
	Encoding string `json:"encoding" yaml:"encoding"`
	// EncoderConfig sets options for the chosen encoder. See
	// zapcore.EncoderConfig for details.
	EncoderConfig EncoderConfig `json:"encoderConfig" yaml:"encoderConfig"`
	// OutputPaths is a list of URLs or file paths to write logging output to.
	// See Open for details.
	OutputPaths []string `json:"outputPaths" yaml:"outputPaths"`
	// InitialFields is a collection of fields to add to the root logger.
	InitialFields map[string]interface{} `json:"initialFields" yaml:"initialFields"`

	// 用来跳到父级 caller 调用者，一般设置为 1 即可
	CallerSkip int
	// 底层日志库的 conf 暴露出来，方便一些特殊操作
	zapBaseConf *zap.Config
}

func NewProductionConfig(fields ...string) *Config {
	genInitialFields(fields)
	return &Config{
		Level:         InfoLevel,
		Development:   false,
		Encoding:      "json",
		EncoderConfig: EncoderConfig{TimeFormat: TimeFormat_TimeStamp},
		OutputPaths:   []string{"stdout"},
		CallerSkip:    1,
		InitialFields: genInitialFields(fields),
	}
}

func NewDevelopmentConfig(fields ...string) *Config {
	return &Config{
		Level:         DebugLevel,
		Development:   true,
		Encoding:      "console",
		EncoderConfig: EncoderConfig{TimeFormat: TimeFormat_ISO8601},
		OutputPaths:   []string{"stdout"},
		CallerSkip:    1,
		InitialFields: genInitialFields(fields),
	}
}

func NewDefaultConfig(fields ...string) *Config {
	return NewProductionConfig(fields...)
}

func genInitialFields(args []string) map[string]interface{} {
	if len(args)%2 != 0 {
		Fatalf("InitialFields invalid: key %s without a value", args[len(args)-1])
	}
	fields := make(map[string]interface{})
	for i := 0; i < len(args); {
		if i == len(args)-1 {
			DPanic(_oddNumberErrMsg)
			break
		}
		fields[args[i]] = args[i+1]
		i += 2
	}
	return fields
}
