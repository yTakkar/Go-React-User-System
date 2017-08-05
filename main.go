package main

import (
	R "Go-React-User-System/routes"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func main() {

	router := httprouter.New()

	router.GET("/*_", R.Route)

	router.POST("/user/signup", R.Signup)
	router.POST("/user/login", R.Login)
	router.POST("/user/logout", R.Logout)

	router.POST("/api/get-session", R.GetSession)

	server := negroni.Classic()
	server.UseHandler(router)
	server.Run(":2000")

}
