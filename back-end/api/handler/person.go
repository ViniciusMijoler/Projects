package handler

import (
	"encoding/json"
	"net/http"
	"projects/back-end/api/db"
	"projects/back-end/model"
	"projects/back-end/util"
)

//InsertPerson ...
func InsertPerson(w http.ResponseWriter, r *http.Request) {
	var t util.App
	var dm model.Person
	var d db.DB
	err := d.Connection()
	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, "Banco de Dados offline", "")
		return
	}
	db := d.DB
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dm); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}
	defer r.Body.Close()
	err = dm.InsertPerson(db)
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Erro ao inserir Situacao", "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, dm, 0, 0)
}

//UpdatePerson ...
// func UpdatePerson(w http.ResponseWriter, r *http.Request) {
// 	var dm model.Person
// 	var t util.App
// 	var d db.DB
// 	err := d.Connection()
// 	if err != nil {
// 		log.Printf("[handler/UpdatePerson] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
// 		return
// 	}
// 	db := d.DB
// 	defer db.Close()

// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		t.ResponseWithError(w, http.StatusBadRequest, "Invalid id", "")
// 		return
// 	}

// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&dm); err != nil {
// 		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", "")
// 		return
// 	}
// 	defer r.Body.Close()
// 	dm.ID = int64(id)
// 	if err := dm.UpdatePerson(db); err != nil {
// 		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
// 		return
// 	}
// 	t.ResponseWithJSON(w, http.StatusOK, dm, 0, 0)
// }

// //DeletePerson ...
// func DeletePerson(w http.ResponseWriter, r *http.Request) {
// 	var dm model.Person
// 	var d db.DB
// 	var t util.App
// 	err := d.Connection()
// 	if err != nil {
// 		log.Printf("[handler/DeletePerson -  Erro ao tentar abrir conexao. Erro: %s", err.Error())
// 		return
// 	}
// 	db := d.DB
// 	defer db.Close()

// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		t.ResponseWithError(w, http.StatusBadRequest, "Invalid id", "")
// 		return
// 	}

// 	dm.CodigoPerson = int64(id)
// 	if err := dm.DeletePerson(db); err != nil {
// 		log.Printf("[handler/DeletePerson -  Erro ao tentar deletar Cidadao. Erro: %s", err.Error())
// 		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
// 		return
// 	}
// 	t.ResponseWithJSON(w, http.StatusOK, dm, 0, 0)
// }

// //GetPerson ...
// func GetPerson(w http.ResponseWriter, r *http.Request) {
// 	var dm model.Person
// 	var t util.App
// 	var d db.DB
// 	err := d.Connection()
// 	if err != nil {
// 		log.Printf("[handler/GetDemanda] -  Erro ao tentar abrir conexao. Erro: %s", err.Error())
// 		return
// 	}
// 	db := d.DB
// 	defer db.Close()

// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		t.ResponseWithError(w, http.StatusBadRequest, "Invalid id", "")
// 		return
// 	}

// 	dm.ID = int64(id)
// 	err = dm.GetPerson(db)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			log.Printf("[handler/GetPerson -  Nao ha Person com este ID.")
// 			t.ResponseWithError(w, http.StatusInternalServerError, "Nao ha cidadao com este ID.", err.Error())
// 		} else {
// 			log.Printf("[handler/GetPerson -  Erro ao tentar buscar Person. Erro: %s", err.Error())
// 			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
// 		}
// 		return
// 	}
// 	t.ResponseWithJSON(w, http.StatusOK, dm, 0, 0)
// }

// //GetPersons ...
// func GetPersons(w http.ResponseWriter, r *http.Request) {
// 	var dm model.Person
// 	var t util.App
// 	var d db.DB
// 	err := d.Connection()
// 	if err != nil {
// 		log.Printf("[handler/GetDemandas] -  Erro ao tentar abrir conexao. Erro: %s", err.Error())
// 		return
// 	}
// 	db := d.DB
// 	defer db.Close()

// 	id, _ := strconv.Atoi(r.FormValue("id"))
// 	nome := r.FormValue("nome")

// 	dm.CodigoPerson = int64(id)
// 	dm.Nome = nome

// 	persons, err := dm.GetPersons(db)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			log.Printf("[handler/GetPersons -  Nao ha Cidadao com este ID.")
// 			t.ResponseWithError(w, http.StatusInternalServerError, "Nao ha Cidadao cadastrado.", err.Error())
// 		} else {
// 			log.Printf("[handler/GetPersons -  Erro ao tentar buscar Cidadao. Erro: %s", err.Error())
// 			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
// 		}
// 		return
// 	}
// 	t.ResponseWithJSON(w, http.StatusOK, persons, 0, 0)
// }
