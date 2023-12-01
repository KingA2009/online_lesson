package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

var mySigningKey = []byte("mysupersecretphrese")

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super secret information")
}

func isAuthorized(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Header["Token"] != nil {
			token, err := jwt.Parse(request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(writer, err.Error())
			}

			if token.Valid {
				endpoint(writer, request)
			}
		} else {
			fmt.Fprintf(writer, "Not Authorized")
		}
	})
}

func handleRequest() {
	http.Handle("/", isAuthorized(homePage))

	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	fmt.Println("My Simple Server")
	handleRequest()
}
