package utils

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"strings"
	"sugar/models"
	"sugar/store"
	"time"
)

func SignToken(userID string) (string, error) {
	privateBytes := []byte(strings.TrimSpace(store.State.Config.Keys.Private))

	block, _ := pem.Decode(privateBytes)
	if block != nil {
		privateBytes = block.Bytes
	}

	signingKey, err := x509.ParsePKCS1PrivateKey(privateBytes)
	if err != nil {
		return "", fmt.Errorf("unable to load private key: %v", err)
	}

	// Create the Claims
	claims := models.JWTToken{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Hour * 48),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signed, err := token.SignedString(signingKey)
	if err != nil {
		logrus.Error("could not sign: ", err)
		return "", err
	}

	return signed, err
}

func VerifyAndParseToken(signedData string) (*models.JWTToken, error) {
	var tokenStruct models.JWTToken

	pubKey := []byte(store.State.Config.Keys.Public)

	block, _ := pem.Decode(pubKey)
	if block != nil {
		pubKey = block.Bytes
	}

	verificationKey, err := x509.ParsePKIXPublicKey(pubKey)
	if err != nil {
		return nil, fmt.Errorf("unable to load public key: %s", err)
	}

	token, err := jwt.Parse(signedData, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return verificationKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not parse signed data: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tokenStruct.UserID = claims["userId"].(string)
	} else {
		return nil, errors.New("invalid token")
	}

	return &tokenStruct, nil
}
