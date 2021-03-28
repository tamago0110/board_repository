package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type JWT struct {
	Token string `json:"token"`
}

type Auth struct {
	UserID string
	Iss    string
}

type Dummy struct {
	UserID string
}

func CreateToken(userID string, tokenSecret, tokenIss string) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = jwt.MapClaims{
		"user": userID,
		"iss":  tokenIss,
	}

	var secretKey = tokenSecret

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Parse(signedString string, tokenSecret string) (*Auth, error) {
	token, err := jwt.Parse(signedString, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})

	if err != nil {
		return nil, err
	}
	if token == nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}
	userID, ok := claims["user"].(string)
	if !ok {
		return nil, err
	}
	iss, ok := claims["iss"].(string)
	if !ok {
		return nil, err
	}

	return &Auth{UserID: userID, Iss: iss}, nil
}

func Authz(tokenSecret, tokenIss string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		auth, err := Parse(tokenString, tokenSecret)
		if err != nil {
			c.String(http.StatusUnauthorized, "You are not authorized user.")
			return
		}

		if auth.Iss != tokenIss {
			c.String(http.StatusUnauthorized, "You are not authorized user.")
			return
		}

		c.Set("userID", auth.UserID)
		c.Next()
	}
}

func dummyFunc(c *gin.Context) {
	var dummy Dummy

	userID, ok := c.Get("userID")
	if !ok {
		c.String(http.StatusUnauthorized, "You are not authorized user.")
		return
	}
	strUserID := fmt.Sprintf("%v", userID)

	dummy.UserID = strUserID
	c.JSON(http.StatusOK, dummy)
}
