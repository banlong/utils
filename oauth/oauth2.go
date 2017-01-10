package main

import (
	"io/ioutil"
	"log"
	"strings"
	"net/http"
	"encoding/json"
	"fmt"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	//"github.com/dgrijalva/jwt-go/request"
	"github.com/dgrijalva/jwt-go/request"
)


//RSA KEYS AND INITIALISATION
const (
	privKeyPath = "oauth/keys/app.rsa"	// openssl genrsa -out app.rsa 1024
	pubKeyPath = "oauth/keys//app.rsa.pub"	// openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var verifyKey, signKey []byte


func initKeys(){
	var err error

	signKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}

	verifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}
}



//STRUCT DEFINITIONS


type UserCredentials struct {
	Username	string  `json:"username"`
	Password	string	`json:"password"`
}

type User struct {
	ID		int 	`json:"id"`
	Name		string  `json:"name"`
	Username	string  `json:"username"`
	Password	string	`json:"password"`
}

type Response struct {
	Data	string	`json:"data"`
}

type Token struct {
	Token 	string    `json:"token"`
}



//SERVER ENTRY POINT


func StartServer(){

	//PUBLIC ENDPOINTS
	http.HandleFunc("/login", LoginHandler)

	//PROTECTED ENDPOINTS
	http.Handle("/resource", negroni.New(
		negroni.HandlerFunc(ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(ProtectedHandler)),
	))

	log.Println("Now listening...8000")
	http.ListenAndServe(":8000", nil)
}

func main() {

	initKeys()
	StartServer()
}



func ProtectedHandler(w http.ResponseWriter, r *http.Request){
	response := Response{"Gained access to protected resource"}
	JsonResponse(response, w)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var user UserCredentials

	//decode request into UserCredentials struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Error in request")
		return
	}

	fmt.Println(user.Username, user.Password)

	//validate user credentials
	if strings.ToLower(user.Username) != "banlong" {
		if user.Password != "1111" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Println("Error logging in")
			fmt.Fprint(w, "Invalid credentials")
			return
		}
	}

	//create a rsa 256 signer
	signer:= jwt.New(jwt.SigningMethodHS256)

	//set claims
	claims := make(jwt.MapClaims)
	claims["iss"] = "admin"
	claims["CustomUserInfo"] = struct {
		Name string
		Role string
	}{user.Username, "Member"}
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	signer.Claims = claims



	tokenString, err := signer.SignedString(signKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		log.Printf("Error signing token: %v\n", err)
		return
	}

	//create a token instance using the token string
	response := Token{tokenString}
	JsonResponse(response, w)

}



//AUTH TOKEN VALIDATION
func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	// sample token string taken from the New example
	token, _ := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error){

		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signKey, nil
	})

	if token.Valid{
		fmt.Fprint(w, "Token is valid")
		next(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Token is not valid")
		return
	}
}



//HELPER FUNCTIONS


func JsonResponse(response interface{}, w http.ResponseWriter) {

	json, err :=  json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
	//w.WriteHeader(http.StatusOK)
}