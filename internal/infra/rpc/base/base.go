package base

import "github.com/dizzrt/dauth/internal/infra/rpc/base/core"

const (
	_ENDPOINT_DAUTH = "discovery:///dauth"

	CK_DAUTH_IDENTITY_USER    = "dauth.identity.user"
	CK_DAUTH_SERVICE_PROVIDER = "dauth.sp"
	CK_DAUTH_TOKEN            = "dauth.token"
)

var clientSet = []core.Client{
	{Key: CK_DAUTH_IDENTITY_USER, Endpoint: _ENDPOINT_DAUTH, Builder: NewUserServiceClient},
	{Key: CK_DAUTH_SERVICE_PROVIDER, Endpoint: _ENDPOINT_DAUTH, Builder: NewServiceProviderServiceClient},
	{Key: CK_DAUTH_TOKEN, Endpoint: _ENDPOINT_DAUTH, Builder: NewTokenServiceClient},
}

func init() {
	core.NewClients(clientSet...)
}

func GetClient(key string) any {
	return core.GetClient(key)
}
