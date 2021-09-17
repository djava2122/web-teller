package service

import (
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"gitlab.pactindo.com/ebanking/common/redis"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

// Claims .
type Claims struct {
	Core           string `json:"core"`
	TellerID       string `json:"tellerID"`
	TellerPass     string `json:"tellerPass"`
	CoCode         string `json:"coCode"`
	branchCode     string `json:"branchCode"`
	branchName     string `json:"branchName"`
	TillCoCode     string `json:"tillCoCode"`
	CompanyCode    string `json:"companyCode"`
	BeginBalance   string `json:"beginBalance"`
	CurrentBalance string `json:"currentBalance"`
	jwt.StandardClaims
}

func createToken(claims *Claims) (string, error) {

	claims.StandardClaims = jwt.StandardClaims{
		// In JWT, the expiry time is expressed as unix milliseconds
		ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)

	keyRedis := "webteller-session-" + claims.Core + claims.TellerID
	err = redis.DB.Set(keyRedis, tokenString, 30*time.Minute).Err()
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func getClaims(token string) (*Claims, error) {
	// Initialize a new instance of `Claims`
	claims := &Claims{}
	tokens := strings.Fields(token)

	if len(tokens) != 2 {
		return claims, errors.New("Unauthorized")
	}

	// Parse the JWT string and store the result in `claims`.
	tkn, err := jwt.ParseWithClaims(tokens[1], claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return claims, errors.New("Unauthorized")
		}
		return claims, errors.New("BadRequest")
	}
	if !tkn.Valid {
		return claims, errors.New("Unauthorized")
	}

	keyRedis := "webteller-session-" + claims.Core + claims.TellerID
	r, err := redis.DB.Get(keyRedis).Result()
	if err != nil || r == "" {
		return nil, errors.New("session expire")
	}

	if r != tokens[1] {
		return nil, errors.New("Unauthorized")
	}

	redis.DB.Expire(keyRedis, 30*time.Minute)

	return claims, nil
}
