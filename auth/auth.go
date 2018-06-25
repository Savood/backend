package auth

import (
	"os"
	"github.com/dgrijalva/jwt-go"
	"git.dhbw.chd.cx/savood/backend/models"
	"fmt"
	"log"
	"github.com/globalsign/mgo/bson"
)

// GetAuthorizer produces Authorizer function
func GetAuthorizer(secret *string) (func(string) (*models.Principal, error)) {

	if secret == nil {
		s := os.Getenv("JWT_SECRET")
		secret = &s
		log.Printf("secret parameter nil, getting env variable JWT_SECRET")
	}

	authFunc := func(tokenString string) (*models.Principal, error) {

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(*secret), nil
		})
		if err != nil {
			log.Printf("jwt parsing failed: %+v", err)
			return nil, err
		}

		var principal models.Principal

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			log.Printf("claims: %+v", claims)

			if claims["userid"] == nil || claims["email"] == nil ||
				claims["userid"].(string) == "" || claims["email"].(string) == "" {

				return nil, fmt.Errorf("claim parsing failed")

			}

			principal = models.Principal{
				Userid: bson.ObjectIdHex(claims["userid"].(string)),
				Email:  claims["email"].(string),
			}

		} else {
			return nil, err
		}

		return &principal, nil

	}

	return authFunc

}
