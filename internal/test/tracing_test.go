package test

import (
	"fmt"
	"github.com/liuzhaomax/maxblog-stats/internal/core"
	"testing"
)

func TestTraceID(t *testing.T) {
	fmt.Println(core.TraceID())
	fmt.Println(core.SpanID())
}
