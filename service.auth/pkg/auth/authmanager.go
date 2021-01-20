package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/0B1t322/distanceLearningWebSite/pkg/models/user"
	"github.com/dgrijalva/jwt-go/v4"
)

type TokenParser interface {
	ParseToken(accessToken string, signingKey []byte) (TokenInfo, error)
}

/*
AuthManager manage authentication
can create tokens and parse them

	methods:
		CreateToken
		ParseToken
*/
type AuthManager struct {
	// secret key  for hash tokens
	signingKey []byte
	// some  hash salt but know not use
	hashSalt string
	// duration for tokens
	expireDuration time.Duration
}

// NewAuthManager create a authmanager
// 	params:
//		signingKey - key for hashing JWT
// 		hashSalt - salt for unhash password (now not use)
// 		expireDuration - duration of JWT token
func NewAuthManager(
	signingKey 		[]byte,
	hashSalt 		string,
	expireDuration 	time.Duration,
) *AuthManager {
	return &AuthManager{
		signingKey:     signingKey,
		hashSalt:       hashSalt,
		expireDuration: expireDuration,
	}
}

/*
CreateToken create a JWT token with information about user - username and role
	params:
		user - a model of user
*/
func (am *AuthManager) CreateToken(u *user.User) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&AuthClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: jwt.At(time.Now().Add(am.expireDuration)),
				IssuedAt:  jwt.At(time.Now()),
			},
			TokenInfo: TokenInfo{
				UID:      fmt.Sprint(u.ID),
				Username: u.Username,
				Role:     u.Role,
			},
		})

	return token.SignedString(am.signingKey)
}

/*
ParseToken parse a token and give information about user with interface TokenInfo
	params:
		token - JWT  token
*/
func (am *AuthManager) ParseToken(token string) (*TokenInfo, error) {
	return ParseToken(token, am.signingKey)
}

/*
ParseToken parse a token and give information about user with interface TokenInfo
	params:
		accessToken - JWT token
		signingKey - a secretKey that hash token
*/
func ParseToken(accessToken string, signingKey []byte) (*TokenInfo, error) {
	token, err := jwt.ParseWithClaims(
		accessToken, &AuthClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexcpeted signing method: %v", t.Header["alg"])
			}

			return signingKey, nil
		})

	if err != nil {
		if checkExpired(err) {
			return nil, ErrTokenExpire
		}

		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return &claims.TokenInfo, nil
	}

	return nil, ErrInvalidToken
}

func checkExpired(err error) bool {
	if strings.Contains(err.Error(), ErrTokenExpire.Error()) {
		return true
	}

	return false
}
