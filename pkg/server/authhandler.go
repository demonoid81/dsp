package server

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

type user struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type authClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func getToken(user string) (string, error) {
	signingKey := []byte("keymaker")
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()
	claims := authClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte("keymaker")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := verifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}
		name := claims.(jwt.MapClaims)["Username"].(string)
		r.Header.Set("name", name)
		next.ServeHTTP(w, r)
	})
}

func (s *Server) loginHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		username := r.FormValue("username")
		password := r.FormValue("password")

		if len(username) == 0 || len(password) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Please provide name and password to obtain the token"))
			return
		}

		var user user

		collection := s.mongo.MongoClient.Database(s.cfg.MongoDatabase).Collection("users")
		err := collection.FindOne(ctx, bson.D{{"username", username}}).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Invalid username: " + err.Error()))
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Name and password do not match"))
			return
		}

		token, err := getToken(user.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error generating JWT token: " + err.Error()))
		} else {
			w.Header().Set("Authorization", "Bearer "+token)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(token))
		}
	}
}

func (s *Server) registerHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")
		var user user

		collection := s.mongo.MongoClient.Database(s.cfg.MongoDatabase).Collection("users")
		err := collection.FindOne(ctx, bson.D{{"username", username}}).Decode(&user)

		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				hash, err := bcrypt.GenerateFromPassword([]byte(password), 5)

				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("Error While Hashing Password, Try Again: " + err.Error()))
					return
				}

				user.Username = username
				user.Password = string(hash)

				_, err = collection.InsertOne(ctx, user)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("Error While Creating User, Try Again: " + err.Error()))
					return
				}
				w.WriteHeader(http.StatusOK)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error While Creating User, Try Again: " + err.Error()))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Username already exists: " + err.Error()))
		return
	}
}
