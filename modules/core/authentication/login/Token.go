package login

import (
	"back-end/modules/database"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"time"
)

var signingKey = []byte(os.Getenv("TOKEN_SECRET_KEY"))

func GenerateToken(user database.Login) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorised"] = true
	claims["user_id"] = user.Id
	claims["type_id"] = user.UserTypeId
	claims["first_name"] = user.FirstName
	claims["last_name"] = user.LastName
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()

	tokenString, err := token.SignedString([]byte("mid-wife-login-auth"))

	if err == nil {
		return tokenString
	}
	return ""
}

func RefreshToken(claims jwt.MapClaims) string {
	user := database.Login{}
	user.Id = claims["user_id"].(int)
	user.FirstName = claims["first_name"].(string)
	user.LastName = claims["last_name"].(string)
	return GenerateToken(user)
}

func DecodeToken(r *http.Request) (*jwt.Token, error) {
	return jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error")
		}
		return signingKey, nil
	})
}