package logger

import (
	"go.uber.org/zap"
	"net/url"
	"os"
)

func newWinFileSink(u *url.URL) (zap.Sink, error) {
	// https://github.com/uber-go/zap/issues/621
	// Remove leading slash left by url.Parse()
	return os.OpenFile(u.Path[1:], os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
}

func init() {
	zap.RegisterSink("winfile", newWinFileSink)
}
