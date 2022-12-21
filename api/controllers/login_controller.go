package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/smhtkn/testpostgre/api/auth"
	"github.com/smhtkn/testpostgre/api/models"
	"github.com/smhtkn/testpostgre/api/responses"
	"github.com/smhtkn/testpostgre/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	client := models.Client{}
	err = json.Unmarshal(body, &client)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	client.Prepare()
	err = client.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(client.Email, client.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	client := models.Client{}

	err = server.DB.Debug().Model(models.Client{}).Where("email = ?", email).Take(&client).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(client.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(client.ID)
}
