package stub

import (
	"context"
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/config"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/driver/cluster/balancer/conn"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/driver/cluster/endpoint"
	"github.com/ydb-platform/ydb-go-sdk/v3/trace"
)

type configStub struct {
	*config.Config
}

func Config(c *config.Config) conn.Config {
	return &configStub{c}
}

func (c configStub) RequestTimeout() time.Duration {
	return c.Config.RequestTimeout
}

func (c configStub) OperationTimeout() time.Duration {
	return c.Config.OperationTimeout
}

func (c configStub) OperationCancelAfter() time.Duration {
	return c.Config.OperationCancelAfter
}

func (c configStub) Meta(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func (c configStub) Trace(context.Context) trace.Driver {
	return c.Config.Trace
}

func (c configStub) Pessimize(context.Context, endpoint.Addr) error {
	return nil
}

func (c configStub) StreamTimeout() time.Duration {
	return c.Config.StreamTimeout
}

func (c configStub) GrpcConnectionPolicy() *conn.GrpcConnectionPolicy {
	return (*conn.GrpcConnectionPolicy)(c.Config.GrpcConnectionPolicy)
}
