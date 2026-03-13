package handler

import (
	"fmt"

	"github.com/dizzrt/dauth/api/gen/authn"
	"github.com/dizzrt/dauth/api/gen/identity"
	"github.com/dizzrt/ellie/transport/grpc"
	"github.com/dizzrt/ellie/transport/http"
	"github.com/google/wire"
)

type ServiceRegistrar struct {
	Identity *IdentityHandler
	Authn    *AuthnHandler
}

func (r *ServiceRegistrar) Register(srv any) {
	switch v := srv.(type) {
	case *grpc.Server:
		identity.RegisterUserServiceServer(v, r.Identity)
		authn.RegisterAuthnServiceServer(v, r.Authn)
	case *http.Server:
		identity.RegisterUserServiceHTTPServer(v, r.Identity)
		authn.RegisterAuthnServiceHTTPServer(v, r.Authn)
	default:
		panic(fmt.Sprintf("unexpected server type %T", srv))
	}
}

var ProviderSet = wire.NewSet(
	NewIdentityHandler,
	NewAuthnHandler,
	wire.Struct(new(ServiceRegistrar), "*"),
)
