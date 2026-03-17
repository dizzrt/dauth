package jwt

import (
	"context"
	"testing"
	"time"

	"github.com/dizzrt/dauth/internal/conf"
	"github.com/dizzrt/dauth/internal/infra/cache/impl"
	"github.com/dizzrt/dauth/internal/infra/foundation"
	"github.com/rs/xid"
)

func getJWTManager() JWTManager {
	appConfig := conf.GetAppConfig()
	redisClient, _, err := foundation.NewRedisClient(appConfig)
	if err != nil {
		panic(err)
	}

	tokenCache := impl.NewTokenCacheImpl(redisClient)
	return NewJWTManager(appConfig, tokenCache)
}

func TestJWTManager(t *testing.T) {
	jwtManager := getJWTManager()

	now := time.Now()
	expiresAt := now.Add(24 * time.Hour)
	idToken := &IDToken{
		Token: Token{
			RegisteredClaims: RegisteredClaims{
				ID:        xid.New().String(),
				Issuer:    "dauth",
				Subject:   "dauth-sso",
				Audience:  ClaimStrings{"dauth"},
				IssuedAt:  NewNumericDate(now),
				NotBefore: NewNumericDate(now),
				ExpiresAt: NewNumericDate(expiresAt),
			},
			Type: TokenType_ID,
		},
		UID: 10000,
		SID: "123456",
	}

	ctx := context.Background()
	token, err := jwtManager.Sign(ctx, idToken, nil)
	if err != nil {
		t.Errorf("sign token failed: %v", err)
		return
	}

	t.Logf("sign token success, token: %s", token)

	var baseToken Token
	err = jwtManager.Verify(ctx, token, &baseToken, nil)
	if err != nil {
		t.Errorf("verify token failed: %v", err)
		return
	}

	t.Logf("verify token success, baseToken: %+v", baseToken)

	var idTokenNew IDToken
	err = jwtManager.Verify(ctx, token, &idTokenNew, nil)
	if err != nil {
		t.Errorf("verify token failed: %v", err)
		return
	}

	t.Logf("verify token success, idTokenNew: %+v", idTokenNew)
}
