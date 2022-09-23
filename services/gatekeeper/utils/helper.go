package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/segmentio/ksuid"
)

type Payload struct {
	Email    string
	EID      string
	IsActive bool
}

func CreateUniqueId() string {
	id := ksuid.New()
	return id.String()
}

func GenerateJWT(payload *Payload) string {
	v, err := time.ParseDuration(os.Getenv("JWT_EXPIRE"))
	if err != nil {
		panic("invalid time duration")
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":       time.Now().Add(v).Unix(),
		"email":     payload.Email,
		"e_id":      payload.EID,
		"is_active": payload.IsActive,
	})

	token, err := t.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		panic(err)
	}
	return token
}

func parse(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

// Verify verifies the jwt token against the secret
func Verify(token string) (*Payload, error) {
	parsed, err := parse(token)

	if err != nil {
		return nil, err
	}

	// Parsing token claims
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	email, ok := claims["email"].(string)
	Eid, ok := claims["e_id"].(string)
	IsActive, ok := claims["is_active"].(bool)
	if !ok {
		return nil, errors.New("something went wrong")
	}

	return &Payload{
		Email:    email,
		EID:      Eid,
		IsActive: IsActive,
	}, nil
}
