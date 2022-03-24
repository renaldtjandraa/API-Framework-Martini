package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
)

//untuk yang array
func SendUsersResponse(w http.ResponseWriter, message string, status int, data []User) {
	var response UsersResponse
	response.Status = status
	response.Message = message
	response.Data = data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//untuk yg bukan array
func SendUserResponse(w http.ResponseWriter, message string, status int, data User) {
	var response UserResponse
	response.Status = status
	response.Message = message
	response.Data = data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//untuk error response
func SendResponse(w http.ResponseWriter, message string, status int) {
	var response MessageResponse
	response.Status = status
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	query := "SELECT * from users "

	rows, errQuery := db.Query(query)
	if errQuery != nil {
		SendResponse(w, "Error Querry", 400)
	}

	var user User
	var users []User

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Email, &user.Password); err != nil {
			SendResponse(w, "Internal Error", 400)
		} else {
			users = append(users, user)
		}
	}
	SendUsersResponse(w, "Request Success", 200, users)
}

func InsertNewUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		SendResponse(w, "Internal Error", 400)
		return
	}

	name := r.FormValue("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	address := r.Form.Get("address")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	result, errQuery := db.Exec("INSERT INTO users(name, age, address , email , password) VALUES (?,?,?,?,?)",
		name,
		age,
		address,
		email,
		password,
	)
	var user User
	temp, _ := result.LastInsertId()
	user.ID = int(temp)
	user.Name = name
	user.Age = age
	user.Address = address
	user.Email = email
	user.Password = password

	if errQuery != nil {
		SendResponse(w, "Querry Error", 400)
		return
	}
	SendUserResponse(w, "Insert Success", 200, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request, params martini.Params) {

	db := Connect()
	defer db.Close()

	UserID := params["id"]

	err := r.ParseForm()
	if err != nil {
		SendResponse(w, "Internal Error", 400)
		return
	}

	name := r.FormValue("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	address := r.Form.Get("address")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	result, errQuery := db.Exec("UPDATE users SET name=?, age=?, address=? , email=? , password=? WHERE id=?",
		name,
		age,
		address,
		email,
		password,
		UserID,
	)

	rowAffected, _ := result.RowsAffected()
	var user User
	user.ID, _ = strconv.Atoi(UserID)
	user.Name = name
	user.Age = age
	user.Address = address
	user.Email = email
	user.Password = password

	if errQuery != nil {
		SendResponse(w, "Query Error", 400)
		return
	} else {
		if rowAffected == 0 {
			SendResponse(w, "No Row Affected", 400)
			return
		}
	}

	SendUserResponse(w, "Insert Success", 200, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, params martini.Params) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		SendResponse(w, "Internal Error", 400)
		return
	}

	UserID := params["id"]

	result, errQuery := db.Exec("DELETE FROM users WHERE ID=?",
		UserID,
	)

	rowAffected, _ := result.RowsAffected()

	if errQuery != nil {
		SendResponse(w, "Querry Error!", 400)
		return
	} else {
		if rowAffected == 0 {
			SendResponse(w, "No Row Affected!", 400)
			return
		}
	}

	SendResponse(w, "Delete Success", 200)
}
