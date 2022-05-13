package utils

import "github.com/golang-jwt/jwt/v4"

func CreateToken(claims jwt.Claims, jwtKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenRaw, err := token.SignedString([]byte(jwtKey))
	return tokenRaw, err
}

func ParseAndValidToken(tokenStr string, claims jwt.Claims, jwtKey string) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, nil
}
