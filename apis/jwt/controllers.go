package jwt

import (
	"os"
	"server/apis/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type request struct {
	UserID string `json:"user_id"`
	Token  string `json:"token,omitempty"`
}

var (
	m *model
)

type model struct {
	subject       string
	issuer        string
	audience      jwt.ClaimStrings
	expiry        time.Duration
	signingMethod jwt.SigningMethod
	jwtSecret     []byte
}

func Init() (err error) {
	err = godotenv.Load("./.env")
	if err != nil {
		return
	}
	m = &model{}
	secret := os.Getenv("JWT_SECRET")
	m.jwtSecret = []byte(secret)
	m.subject = os.Getenv("SUBJECT")
	m.issuer = os.Getenv("ISSUER")
	audience := os.Getenv("AUDIENCE")
	m.audience = append(m.audience, audience)
	signinMethod := os.Getenv("SIGNING_METHOD")
	switch signinMethod {
	default:
		m.signingMethod = jwt.SigningMethodHS256
	}
	expiry := os.Getenv("EXPIRY")
	duration, err := time.ParseDuration(expiry)
	if err != nil {
		return
	}
	if duration.Seconds() <= 100 {
		duration = 1800 * time.Second
	}
	m.expiry = duration
	return
}

func generate(c *fiber.Ctx) (err error) {
	bytes := c.Request().Body()
	req := request{}
	err = json.Decode(bytes, &req)
	if err != nil {
		return c.JSON(err)
	}
	token, err := m.CreateToken(req.UserID)
	if err != nil {
		return c.JSON(err)
	}
	data := json.Map{
		"user_id": req.UserID,
		"token":   token,
	}
	return c.JSON(data)
}

func validate(c *fiber.Ctx) (err error) {
	bytes := c.Request().Body()
	req := request{}
	err = json.Decode(bytes, &req)
	if err != nil {
		return c.JSON(err)
	}
	ok, err := VerifyToken(req.UserID, req.Token)
	if err != nil {
		return c.JSON(err)
	}
	data := json.Map{
		"verified": ok,
	}
	return c.JSON(data)
}

func (m *model) CreateToken(userID string) (signedToken string, err error) {
	claims := jwt.RegisteredClaims{
		ID:       userID,
		Issuer:   m.issuer,
		Subject:  m.subject,
		Audience: m.audience,
		ExpiresAt: &jwt.NumericDate{
			Time: time.Now().Add(m.expiry),
		},
		IssuedAt: &jwt.NumericDate{
			Time: time.Now().UTC(),
		},
	}
	token := jwt.NewWithClaims(m.signingMethod, claims)
	signedToken, err = token.SignedString(m.jwtSecret)
	if err != nil {
		return
	}
	return
}

func VerifyToken(id string, signedToken string) (ok bool, err error) {
	claims := jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(signedToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return m.jwtSecret, nil
	})
	if !token.Valid {
		return
	}
	if claims.Issuer != m.issuer {
		return
	}
	if claims.Subject != m.subject {
		return
	}
	if len(m.audience) != len(claims.Audience) {
		return
	}
	for i := 0; i < len(m.audience); i++ {
		if m.audience[i] != claims.Audience[i] {
			return
		}
	}
	if claims.ID != id {
		return
	}
	ok = true
	return
}
