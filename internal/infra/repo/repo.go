package repo

import (
	"github.com/dizzrt/dauth/internal/infra/repo/impl/client"
	"github.com/dizzrt/dauth/internal/infra/repo/impl/identity"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	// identity_impls
	identity.NewUserRepoImpl,
	identity.NewRoleRepoImpl,
	identity.NewUserRoleAssociationRepoImpl,

	// client_impls
	client.NewClientRepoImpl,
	client.NewScopeRepoImpl,
	client.NewClientScopeAssociationRepoImpl,
)
