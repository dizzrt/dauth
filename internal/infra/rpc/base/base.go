package base

import "github.com/dizzrt/dauth/internal/infra/rpc/base/core"

const (
	_ENDPOINT_DAUTH = "discovery:///dauth"

	CK_DAUTH_IDENTITY_USER = "dauth.identity.user"
)

var clientSet = []core.Client{
	{Key: CK_DAUTH_IDENTITY_USER, Endpoint: _ENDPOINT_DAUTH, Builder: NewUserServiceClient},
}

func init() {
	core.NewClients(clientSet...)
}

func GetClient(key string) any {
	return core.GetClient(key)
}
