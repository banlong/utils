package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
)
// using asymmetric crypto/RSA keys
// location of the files used for signing and verification
const (
	privKeyPath = "oauth/keys/app.rsa.ppk" // openssl genrsa -out app.rsa 1024
	pubKeyPath = "oauth/keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)
// verify key and sign key
var (
	verifyKey, signKey []byte
)
//struct User for parsing login credentials
type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
type Token struct {
	Token string `json:"token"`
}
type Response struct {
	Text string `json:"text"`
}


// read the key files before starting http handlers
func init() {
	var err error
	signKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}
	verifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/auth", authHandler).Methods("POST")
	server := &http.Server{
		Addr: ":9000",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()

}

// reads the login credentials, checks them and creates JWT the token
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	//decode into User struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error in request body")
		return
	}

	// validate user credentials
	if user.UserName != "banlong" && user.Password != "1111" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Wrong info")
		return
	}

	// create a signer for rsa 256
	//t := jwt.New(jwt.GetSigningMethod("RS256"))
	t := jwt.New(jwt.SigningMethodRS512)

	// set our claims
	claims := make(jwt.MapClaims)
	claims["iss"] = "admin"
	claims["CustomUserInfo"] = struct {
		Name string
		Role string
	}{user.UserName, "Member"}
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	t.Claims = claims

	tokenString, err := t.SignedString(signKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Sorry, error while Signing Token!")
		log.Printf("Token Signing error: %v\n", err)
		return
	}
	response := Token{tokenString}
	jsonResponse(response, w)
}

// only accessible with a valid token
func authHandler(w http.ResponseWriter, r *http.Request) {
	token, err := request.ParseFromRequest(r, request.OAuth2Extractor, keyLookupFunc)
	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError: // something was wrong during the validation
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, "Token Expired, get a new one.")
				return
			default:
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "Error while Parsing Token!")
				log.Printf("ValidationError error: %+v\n", vErr.Errors)
				return
			}
		default: // something else went wrong
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error while Parsing Token!")
			log.Printf("Token parse error: %v\n", err)
			return
		}
	}


	if token.Valid {
		response := Response{"Authorized to the system"}
		jsonResponse(response, w)
	} else {
		response := Response{"Invalid token"}
		jsonResponse(response, w)
	}
}

func jsonResponse(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func keyLookupFunc(token *jwt.Token) (interface{}, error) {
	// since we only use one private key to sign the tokens,
	// we also only use its public counter part to verify
	return verifyKey, nil
}

