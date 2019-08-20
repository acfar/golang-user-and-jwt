package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"refactorGolang/database"
	"refactorGolang/models"
	"refactorGolang/utils"
	"golang.org/x/crypto/bcrypt"
)


var user models.User
var connect database.Connection

func (c Controller)  Signup() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		var error models.Error
		db:=connect.Connect()

		_=json.NewDecoder(r.Body).Decode(&user)

		if user.Username==""{
			error.Message="username required"
			utils.RespondWithError(w,http.StatusBadRequest,error)
			return
		}

		if user.Password==""{
			error.Message="password required"
			utils.RespondWithError(w,http.StatusBadRequest,error)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password),10)

		if err !=nil{
			log.Fatal(err)
		}
		user.Password=string(hash)

		_, err = db.Exec("insert into users(username, password) values(?,?)", user.Username, user.Password, )
		if err!=nil{
			error.Message="Server Error"
			utils.RespondWithError(w,http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-type","application/json" )
		utils.RespondJson(w, user)
		json.NewEncoder(w).Encode(user)
	}
}

func (c Controller) Login () http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var user models.User
		var jwt models.JWT
		var error models.Error
		db:=connect.Connect()

		json.NewDecoder(r.Body).Decode(&user)

		if user.Username==""{
			error.Message = "Username is missing"
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}
		 if user.Password==""{
		 	error.Message="password is missing"
		 	utils.RespondWithError(w, http.StatusBadRequest, error)
		 }
		password := user.Password
		row :=db.QueryRow("select * from users where username=?", user.Username)
		err :=row.Scan(&user.Id, &user.Username, &user.Password)
		if err!=nil{
			if err == sql.ErrNoRows {
				error.Message = "The user does not exist"
				utils.RespondWithError(w, http.StatusBadRequest, error)
				return
			} else {
				log.Fatal(err)
			}
		}
		hashedPassword := user.Password
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
		if err != nil {
			error.Message = "Invalid Password"
			utils.RespondWithError(w, http.StatusUnauthorized, error)
			return
		}

		token, err := utils.GenerateToken(user)

		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		jwt.Token = token

		utils.RespondJson(w, jwt)

	}

}
