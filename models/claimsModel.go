package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	CustomerId    string `json:"customerId,omitempty"`
	CustomerName  string `json:"customerName,omitempty"`
	CustomerEmail string `json:"customerEmail,omitempty"`
	jwt.StandardClaims
}

const (
	ip = "192.168.0.107"
)

func (claim Token) Valid() error {
	var now = time.Now().Unix()
	if claim.VerifyExpiresAt(now, true) && claim.VerifyIssuer(ip, true) {
		return nil
	}
	return fmt.Errorf("Token is invalid")
}

func (claim Token) VerifyAudience(origin string) bool {
	return strings.Compare(claim.Audience, origin) == 0
}
