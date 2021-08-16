package server

import (
	greeterpbv1 "account/api/helloworld/v1"
	rbacpbv1 "account/api/rbac/v1"
	"account/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	greetersvcv1 "account/internal/service/greeter/v1"
	rbacsvcv1 "account/internal/service/rbac/v1"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, gv1 *greetersvcv1.GreeterService, rv1 *rbacsvcv1.V1Service, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			metrics.Server(),
			validate.Validator(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	greeterpbv1.RegisterGreeterServer(srv, gv1)
	rbacpbv1.RegisterV1Server(srv, rv1)
	return srv
}
