package routes

import (
	"main/controllers"

	"github.com/gorilla/mux"
)

func SetUserRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/api").Subrouter()
	subRoute.HandleFunc("/register", controllers.Register).Methods("POST")
	subRoute.HandleFunc("/login", controllers.Login).Methods("POST")
	subRoute.HandleFunc("/getUser/{id}", controllers.GetUser).Methods("GET")
	subRoute.HandleFunc("/getAllUser", controllers.GetAllUser).Methods("GET")
	subRoute.HandleFunc("/logout/{id}", controllers.Logout).Methods("POST")
	subRoute.HandleFunc("/addQuestion", controllers.AddQuestion).Methods("POST")
	subRoute.HandleFunc("/getQuestions", controllers.GetQuestions).Methods("GET")
	subRoute.HandleFunc("/getAnswer/{id}", controllers.GetAnswer).Methods("GET")
	subRoute.HandleFunc("/getResult", controllers.GetResult).Methods("POST")
}
