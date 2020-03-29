package logger

import (
	"context"
)

/*
直接复用 jaeger 传递下来的 _uuid，所以不用再额外处理了
为保持向前兼容，这两个接口保留
*/

func WithTraceID(ctx context.Context, traceID string) context.Context {
	return ctx
	//if _traceID := GetTraceID(ctx); _traceID != "" {
	//	return ctx
	//}
	//
	//md := metadata.Pairs(TraceID, traceID)
	//return metadata.NewOutgoingContext(ctx, md)
}

func GetTraceID(ctx context.Context) (traceID string) {
	if id, ok := ctx.Value(TraceID).(string); ok {
		return id
	}
	return ""
	//md, ok := metadata.FromIncomingContext(ctx)
	//if !ok {
	//	return ""
	//}
	//return md.Get(TraceID)[0]
}
