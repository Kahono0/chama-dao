package routes

import (
	"fmt"
	"net/http"
	"github.com/Kahono0/chama-dao/models"
	"github.com/Kahono0/chama-dao/utils"
	"encoding/json"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	//get user from context
	user := r.Context().Value("user").(models.User)

	fmt.Fprintf(w, "Welcome %s", user.Name)
}

func signUp(w http.ResponseWriter, r *http.Request) {
	///get user from request body
	var usr models.User
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//check if user already exists
	// _, err = utils.GetUserByAddress(usr.Address)
	// if err == nil {
	// 	http.Error(w, "User already exists", http.StatusBadRequest)
	// 	return
	// }

	//save user to db
	result := utils.DB.Create(&usr)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	u := user{
		Username: usr.Name,
		Address: usr.Address,
	}

	//generate jwt
	token, err := generateJWT(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	type ruser struct{
		User models.User `json:"user"`
		Token struct{
			Token string `json:"token"`
		} `json:"token"`
	}

	rstruct := ruser{
		User: usr,
		Token: struct{
			Token string `json:"token"`
		}{
			Token: token,
		},
	}

	//send token to client
	json.NewEncoder(w).Encode(rstruct)

	return
}

func checkUsername(w http.ResponseWriter, r *http.Request){
	//get username from query params
	username := r.URL.Query().Get("username")

	//check if username exists
	var user models.User

	result := utils.DB.First(&user, "name = ?", username)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Username already exists"))
	return
}

func getUser(w http.ResponseWriter, r *http.Request){
	//get username from query params
	username := r.URL.Query().Get("username")

	//check if username exists
	var user models.User

	result := utils.DB.First(&user, "name = ?", username)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	//send user to client
	json.NewEncoder(w).Encode(user)
	return
}
