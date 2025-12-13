package jwt

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	token_api "github.com/dizzrt/dauth/api/gen/token"
	"github.com/dizzrt/dauth/internal/domain/token/entity"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func getTestJWTManager() *jwtManager {
	return &jwtManager{
		algorithm: jwt.SigningMethodHS256.Alg(),
		issuer:    "dauth",
		secret:    []byte("test-secret"),
	}
}

func getToken() jwt.Claims {
	now := time.Now()
	expiresAt := now.Add(24 * time.Hour)

	ssoToken := &entity.SSOToken{
		BaseToken: entity.BaseToken{
			RegisteredClaims: jwt.RegisteredClaims{
				ID:        uuid.NewString(),
				Issuer:    "dauth",
				Subject:   "dauth-sso",
				Audience:  jwt.ClaimStrings{"dauth"},
				IssuedAt:  jwt.NewNumericDate(now),
				NotBefore: jwt.NewNumericDate(now),
				ExpiresAt: jwt.NewNumericDate(expiresAt),
			},
			UID:  10000,
			Type: token_api.Token_TokenType_SSO,
		},
	}

	return ssoToken
}

func TestJwt(t *testing.T) {
	jm := getTestJWTManager()

	token := getToken()
	tokenStr, err := jm.Sign(context.Background(), token, nil)
	if err != nil {
		t.Errorf("sign sso token failed: %v", err)
		return
	}

	t.Logf("sso token: %s", tokenStr)

	var ssoToken entity.SSOToken
	err = jm.Verify(context.Background(), tokenStr, nil, &ssoToken)
	if err != nil {
		t.Errorf("verify sso token failed: %v", err)
		return
	}

	if ssoToken.Type != token_api.Token_TokenType_SSO {
		t.Errorf("token type not sso: %v", ssoToken.Type)
		return
	}

	jb, _ := json.Marshal(ssoToken)
	t.Logf("sso token: %s", string(jb))
}
