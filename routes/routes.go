package routes

import (
	"context"
	"fmt"
	"strings"
	"net/http"

	//import jwt
	"github.com/Kahono0/chama-dao/models"
	"github.com/Kahono0/chama-dao/utils"
	jwt "github.com/dgrijalva/jwt-go"
)

const SECRET = "secret"

type user struct{
	Username string `json:"username"`
	Address string `json:"address"`
}

func generateJWT(u user) (string, error) {
	//create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
		"address": u.Address,
	})

	//sign token
	tokenString, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func decodeJWT(tokenString string) (jwt.MapClaims, error) {
	//decode token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func enableCors(w *http.ResponseWriter){
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}


func authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//check if user is authorized
		//if not, return 401
		token := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]
		fmt.Println(token)
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Empty token"))
			return
		}

		//decode token
		claims, err := decodeJWT(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		fmt.Println(claims)

		//check if user exists in using address
		var user models.User
		user, err = utils.GetUserByAddress(claims["address"].(string))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("User not found"))
			return
		}

		//set current user as context
		ctx := r.Context()
		ctx = context.WithValue(ctx, "user", user)
		r = r.WithContext(ctx)

		next(w, r)
	}
}

//cors middleware
func cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if(*r).Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(http.StatusOK)
		} else {
			next(w, r)
		}
	}
}


func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	u := user{
		"admin",
		"0x123",
	}

	token, err := generateJWT(u)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(token)
	//signup
	mux.HandleFunc("/signup", cors(signUp))

	//check if username exists
	mux.HandleFunc("/check", cors(checkUsername))

	//get user by username
	mux.HandleFunc("/user", cors(getUser))

	//add proposal

	//middleware
	mux.HandleFunc("/", cors(authorization(homePage)))

	return mux
}
