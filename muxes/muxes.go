package muxes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type (
	//User example model
	User struct {
		ID      int
		Name    string
		Surname string
	}
)

//Users is the user array for the project
type Users = []User

//Serve initializes a basic Mux server
func Serve(db *sql.DB) *http.ServeMux {
	log.Print("Server started at http://127.0.0.1:3000 port.")
	mux := http.NewServeMux()

	mux.HandleFunc("/newUser", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			http.NotFound(w, req)
			return
		}
		var newUser = convertRequestToUser(req)
		var lastInsertID int
		err := db.QueryRow("INSERT INTO users(name,surname) VALUES($1,$2) returning id;", newUser.Name, newUser.Surname).Scan(&lastInsertID)
		checkErr(err)
		fmt.Println("last inserted id =", lastInsertID)

		okStatus(w)
		log.Printf("New User %s %s added successfully.", newUser.Name, newUser.Surname)
		json.NewEncoder(w).Encode(newUser)

		return
	})

	mux.HandleFunc("/getAll", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			http.NotFound(w, req)
			return
		}
		okStatus(w)
		json.NewEncoder(w).Encode(getAllUsers(db))
		log.Printf("All users listed successfully.")

		return
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, req *http.Request) {
		var method = req.Method
		if method != "GET" && method != "DELETE" && method != "PUT" {
			http.NotFound(w, req)
			return
		}

		id, err := strconv.Atoi(strings.TrimPrefix(req.URL.Path, "/users/"))
		if err != nil {
			panic(err)
		}

		okStatus(w)

		if method == "GET" {
			rows, err := db.Query("SELECT * FROM users where id=$1", id)
			checkErr(err)
			user := User{}
			for rows.Next() {
				err = rows.Scan(&user.ID, &user.Name, &user.Surname)
				checkErr(err)
			}
			log.Printf("The user who is id = %d listed successfully.", id)
			json.NewEncoder(w).Encode(user)
			return
		}

		if method == "DELETE" {
			stmt, err := db.Prepare("delete from users where id=$1")
			checkErr(err)
			_, err = stmt.Exec(id)
			checkErr(err)
			log.Printf("The user who is id = %d deleted successfully.", id)
			json.NewEncoder(w).Encode(nil)
			return
		}

		if method == "PUT" {
			user := convertRequestToUser(req)
			stmt, err := db.Prepare("update users set name=$1, surname=$2 where id=$4")
			checkErr(err)
			_, err = stmt.Exec(user.Name, user.Surname, id)
			checkErr(err)
			log.Printf("The user who is id = %d updated successfully.", id)
			json.NewEncoder(w).Encode(user)
			return
		}

		json.NewEncoder(w).Encode(nil)
		return
	})

	return mux
}

func getAllUsers(db *sql.DB) Users {
	rows, err := db.Query("SELECT * FROM users")
	checkErr(err)
	user := User{}
	allUsers := Users{}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Surname)
		allUsers = append(allUsers, user)
		checkErr(err)
	}

	return allUsers
}

func okStatus(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	return
}

func convertRequestToUser(req *http.Request) User {
	body, err := ioutil.ReadAll(req.Body)
	checkErr(err)
	var newUser User
	err = json.Unmarshal(body, &newUser)
	checkErr(err)

	return newUser
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
