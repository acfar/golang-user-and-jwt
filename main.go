package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"refactorGolang/controller"
	"refactorGolang/utils"
)


func main(){
	controller := controller.Controller{}

	router := mux.NewRouter()
	router.HandleFunc("/getData", utils.TokenVerifyMiddleWare(controller.GetData())).Methods("GET")
	router.HandleFunc("/login", controller.Login()).Methods("GET")
	router.HandleFunc("/signup", controller.Signup()).Methods("POST")
	http.Handle("/",router)
	log.Fatal(http.ListenAndServe(":1234",router))

}
