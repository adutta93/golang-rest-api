package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"mongo-golang-api/controllers"
	// "mongo-golang-api/models"
	"net/http"
)

func main() {
	route := httprouter.New()
	userController := controllers.NewUserController(getSession())
	route.GET("/api/user/:id", userController.GetUser)
	route.POST("/api/user/", userController.CreateUser)
	route.DELETE("/api/user/:id", userController.DeleteUser)
	err := http.ListenAndServe(":8080", route)
	if err != nil {
		panic(err)
	}
	fmt.Print("Server started")
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017/go-mongo")
	if err != nil {
		panic(err)
	}
	fmt.Print("MongoDB server started")
	return session
}
