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

	m.Get("/users/getAllUsers", controllers.GetAllUsers)
	m.Post("/users/insertNewUser", controllers.InsertNewUser)
	m.Put("/users/updateUser/:id", controllers.UpdateUser)
	m.Delete("/users/deleteUser/:id", controllers.DeleteUser)

	http.Handle("/", m)
	fmt.Println("Connect to port 8080")
	m.RunOnAddr(":8080")
}
