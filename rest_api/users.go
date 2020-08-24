package main


import (
	"encoding/json"
	"log"
	"net/http"
)

/*
* Get All Data
*/
func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	var users Users
	var arr_user []Users
	var response Response

	db := connect()
	defer db.Close()

	rows, err := db.Query("Select id,username,password from tbl_login")
	
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.Username, &users.Password);err != nil {
			log.Fatal(err.Error())
		} else {
			arr_user = append(arr_user, users)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

/*
* Insert Data
*/
func insertUsers(w http.ResponseWriter, r *http.Request) {
	// var users Users
	var arr_user []Users
	var response Response

	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	name := r.Form.Get("name")
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	level := r.Form.Get("level")

	_, err = db.Exec("INSERT INTO tbl_login (name, username, password, level) value(?, ?, ?, ?)",
		name,
		username,
		password,
		level,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 200
	response.Message = "Berhasil Menambahkan Data"
	response.Data = arr_user
	log.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

/*
* Update Data
*/
func updateUsers(w http.ResponseWriter, r *http.Request) {

	var arr_user []Users
	var response Response

	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	id := r.Form.Get("id")
	name := r.Form.Get("name")
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	level := r.Form.Get("level")

	_, err = db.Exec("UPDATE tbl_login set name = ?, username = ?, password = ?, level = ? where id = ?",
		name,
		username,
		password,
		level,
		id,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Update Data"
	response.Data = arr_user
	log.Print("Update data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

/*
* Delete Data
*/
func deleteUsers(w http.ResponseWriter, r *http.Request) {
	var response Response
	
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	id := r.Form.Get("id")

	_, err = db.Exec("DELETE FROM tbl_login WHERE id = $id" , 
		id,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Delete Data"
	log.Print("Delete data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
