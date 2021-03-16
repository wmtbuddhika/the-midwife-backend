package user

import (
	"back-end/modules/database"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	result := database.GetAllUsers()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(result)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user = database.Login{}
	data, err := ioutil.ReadAll(r.Body)

	if err == nil {
		err := json.Unmarshal(data, &user)
		if err == nil {
			result := database.CreateUser(user)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			json.NewEncoder(w).Encode(result)
		}
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user = database.Login{}
	data, err := ioutil.ReadAll(r.Body)

	if err == nil {
		err := json.Unmarshal(data, &user)
		if err == nil {
			result := database.UpdateUser(user)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			json.NewEncoder(w).Encode(result)
		}
	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user = database.Login{}
	data, err := ioutil.ReadAll(r.Body)

	if err == nil {
		err := json.Unmarshal(data, &user)
		if err == nil {
			result := database.DeleteUser(user)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			json.NewEncoder(w).Encode(result)
		}
	}
}