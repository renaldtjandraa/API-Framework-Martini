package main

import (
	controllers "Martini/controllers"
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	m := martini.Classic()

	m.Get("/users/getAllUsers", controllers.Authenticate(controllers.GetAllUsers, 1))
	m.Post("/users/insertNewUser", controllers.Authenticate(controllers.InsertNewUser, 0))
	m.Put("/users/updateUser/:id", controllers.Authenticate(controllers.InsertNewUser, 1))
	m.Delete("/users/updateUser/:id", controllers.DeleteUser)

	m.Post("/users/Login", controllers.LoginUser)
	m.Post("/users/Logout", controllers.LogoutUser)

	http.Handle("/", m)
	fmt.Println("Connect to port 8080")
	m.RunOnAddr(":8080")
}
