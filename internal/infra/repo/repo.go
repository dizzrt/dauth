package repo

import (
	"github.com/dizzrt/dauth/internal/infra/repo/impl/auth"
	"github.com/dizzrt/dauth/internal/infra/repo/impl/identity"
	"github.com/dizzrt/dauth/internal/infra/repo/impl/sp"
	"github.com/dizzrt/dauth/internal/infra/repo/impl/token"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	// identity_impls
	identity.NewUserRepoImpl,
	identity.NewRoleRepoImpl,
	identity.NewUserRoleAssociationRepoImpl,

	// sp_impls
	sp.NewServiceProviderRepoImpl,
	sp.NewScopeRepoImpl,
	sp.NewSPScopeAssociationRepoImpl,

	// token_impls
	token.NewTokenBlacklistRepoImpl,

	// auth_impls
	auth.NewAuthorizationCodeRepoImpl,
)
