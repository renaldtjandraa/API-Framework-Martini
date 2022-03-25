package main

import (
	controllers "Martini/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

func main() {
	m := martini.Classic()

	// m.Get("/users/getAllUsers", controllers.Authenticate(controllers.GetAllUsers, 1))
	// m.Post("/users/insertNewUser", controllers.Authenticate(controllers.InsertNewUser, 0))

	m.Get("/users/getAllUsers", controllers.Authenticate(1), controllers.GetAllUsers)
	m.Post("/users/insertNewUser", controllers.Authenticate(0), controllers.InsertNewUser)
	m.Put("/users/updateUser/:id", controllers.Authenticate(1), controllers.UpdateUser)
	// m.Put("/users/updateUser/:id", controllers.UpdateUser)
	m.Delete("/users/deleteUser/:id", controllers.Authenticate(1), controllers.DeleteUser)

	m.Post("/users/Login", controllers.LoginUser)
	m.Post("/users/Logout", controllers.LogoutUser)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})
	handler := corsHandler.Handler(m)

	http.Handle("/", m)
	fmt.Println("Connect to port 8080")
	m.RunOnAddr(":8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
