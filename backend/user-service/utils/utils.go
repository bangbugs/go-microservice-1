package utils

import (
	"errors"
	"time"

	"github.com/bangbugs/go-microservice-1/user-service/common"
	jwt_lib "github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

// SdtClaims defines the custom claims
type SdtClaims struct {
	Name string `json:"name"`
	Role string `json:"role"`
	jwt_lib.StandardClaims
}

type Utils struct{}

// GeneratesJWT -> generates jwt token from given information
func (u *Utils) GenerateJWT(name string, role string) (string, error) {
	claims := SdtClaims{
		name,
		role,
		jwt_lib.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    common.Config.Issuer,
		},
	}

	token := jwt_lib.NewWithClaims(jwt_lib.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(common.Config.JwtSecretPassword))

	return tokenString, err
}

func (u *Utils) ValidateObjectID(id string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New(common.ErrNotObjectIndex)
	}
	return nil
}
