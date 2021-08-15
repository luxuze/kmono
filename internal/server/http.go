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
	"github.com/go-kratos/kratos/v2/transport/http"

	greetersvcv1 "account/internal/service/greeter/v1"
	rbacsvcv1 "account/internal/service/rbac/v1"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, gv1 *greetersvcv1.GreeterService, rv1 *rbacsvcv1.V1Service, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			metrics.Server(),
			validate.Validator(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	greeterpbv1.RegisterGreeterHTTPServer(srv, gv1)
	rbacpbv1.RegisterV1HTTPServer(srv, rv1)

	return srv
}
