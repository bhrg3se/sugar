package models

import "github.com/golang-jwt/jwt"

type Config struct {
	Database struct {
		User         string `toml:"user"`
		Password     string `toml:"password"`
		Host         string `toml:"host"`
		Port         string `toml:"port"`
		Name         string `toml:"name"`
		SSL          bool   `toml:"ssl"`
		CaCertPath   string `toml:"caCertPath"`
		UserCertPath string `toml:"userCertPath"`
		UserKeyPath  string `toml:"userKeyPath"`
	} `toml:"database"`

	Logging struct {
		Level string `toml:"logging"`
	} `toml:"logging"`

	Server struct {
		Listen string `toml:"listen"`
	} `toml:"server"`

	Acme struct {
		ApiKey string `toml:"apiKey"`
		ApiURL string `toml:"apiURL"`
	} `toml:"acme"`

	Keys struct {
		Private string `toml:"private"`
		Public  string `toml:"public"`
	} `toml:"keys"`
}

type JWTToken struct {
	jwt.StandardClaims
	UserID string `json:"userId"`
}

type Door struct {
	ID           string `json:"id"`
	AcmeDeviceID string `json:"acmeDeviceID,omitempty"`
	Name         string `json:"name"`
}

type Permission struct {
	ID     string `json:"id"`
	DoorID string `json:"doorID"`
	UserID string `json:"userID"`
}
