package trace

import (
	"context"
	"strings"

	"github.com/google/uuid"
)

// YunxinTraceId 云信TraceId管理器
type YunxinTraceId struct{}

var (
	traceContextKey = "yunxin_trace_id"
)

// GetFromContext 从context中获取trace-id
func GetFromContext(ctx context.Context) string {
	if traceId, ok := ctx.Value(traceContextKey).(string); ok {
		return traceId
	}
	return ""
}

// SetToContext 设置trace-id到context
func SetToContext(ctx context.Context, traceId string) context.Context {
	return context.WithValue(ctx, traceContextKey, traceId)
}

func Gen() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
