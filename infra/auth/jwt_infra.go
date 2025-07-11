package auth

import (
    "errors"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "github.com/sapaude/go_sapaude_backend_admin/conf"
)

// JWTInfra JWT认证设施
type JWTInfra struct {
    secretKey     string
    tokenDuration time.Duration
}

func NewJWTInfra() *JWTInfra {

    jwtCfg := conf.AppConfig.JWT
    secret := jwtCfg.Secret
    duration := time.Duration(jwtCfg.ExpirationMinutes * 60)

    return &JWTInfra{
        secretKey:     secret,
        tokenDuration: duration,
    }
}

type Claims struct {
    UserID string `json:"user_id"`
    Role   string `json:"role"`
    jwt.RegisteredClaims
}

func (j *JWTInfra) GenerateToken(userID string, role string) (string, error) {
    claims := &Claims{
        UserID: userID,
        Role:   role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.tokenDuration)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(j.secretKey))
}

func (j *JWTInfra) ValidateToken(tokenStr string) (string, string, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(j.secretKey), nil
    })
    if err != nil || !token.Valid {
        return "", "", errors.New("invalid token")
    }

    claims, ok := token.Claims.(*Claims)
    if !ok {
        return "", "", errors.New("invalid claims")
    }
    return claims.UserID, claims.Role, nil
}
