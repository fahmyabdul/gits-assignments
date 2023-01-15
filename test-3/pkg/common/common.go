package common

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func VerifyJwt(jwtToken string) error {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("g1tSv3r1fy#!"), nil
	})

	if err != nil {
		return err
	}

	if claims["name"] != "GITS-TEST" {
		return fmt.Errorf("unauthorized")
	}

	return nil
}
