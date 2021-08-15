package service

import (
	greetersvcv1 "account/internal/service/greeter/v1"
	rbacsvc "account/internal/service/rbac/v1"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	greetersvcv1.NewGreeterService,
	rbacsvc.NewV1Service,
)
