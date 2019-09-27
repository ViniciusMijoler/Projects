package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"projects/back-end/api/db"
	"projects/back-end/model"
	"projects/back-end/util"
)

//Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	var t util.App
	var d db.DB
	var u model.User

	err := d.Connection()
	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, "Banco de Dados offline", "")
		return
	}
	db := d.DB
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}
	defer r.Body.Close()
	user, err := u.Login(db)
	if err != nil {
		if err == sql.ErrNoRows {
			t.ResponsePostWithJSON(w, http.StatusOK, nil)
		} else {
			t.ResponseWithError(w, http.StatusInternalServerError, "Erro ao validar Usuário", err.Error())
		}
	}
	t.ResponsePostWithJSON(w, http.StatusOK, user)
}
