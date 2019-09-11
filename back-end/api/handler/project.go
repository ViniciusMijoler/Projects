package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"projects/back-end/api/db"
	"projects/back-end/model"
	"projects/back-end/util"

	"github.com/gorilla/mux"
)

//InsertProject ...
func InsertProject(w http.ResponseWriter, r *http.Request) {
	var t util.App
	var dm model.Project
	var d db.DB
	err := d.Connection()
	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, "Banco de Dados estÃ¡ down", "")
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
	err = dm.InsertProject(db)
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Erro ao inserir Situacao", "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, dm, 0, 0)
}

//UpdateProject ...
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	var dm model.Project
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/UpdateProject] -  Erro ao tentar abrir conexÃ£o. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid id", "")
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dm); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", "")
		return
	}
	defer r.Body.Close()
	dm.CodigoProject = int64(id)
	if err := dm.UpdateProject(db); err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, dm, 0, 0)
}

//DeleteProject ...
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	var dm model.Project
	var d db.DB
	var t util.App
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/DeleteProject -  Erro ao tentar abrir conexÃ£o. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid id", "")
		return
	}

	dm.CodigoProject = int64(id)
	if err := dm.DeleteProject(db); err != nil {
		log.Printf("[handler/DeleteProject -  Erro ao tentar deletar CidadÃ£o. Erro: %s", err.Error())
		t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, dm, 0, 0)
}

//GetProject ...
func GetProject(w http.ResponseWriter, r *http.Request) {
	var dm model.Project
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetDemanda] -  Erro ao tentar abrir conexÃ£o. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid id", "")
		return
	}

	dm.CodigoProject = int64(id)
	err = dm.GetProject(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetProject -  NÃ£o hÃ¡ Project com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "NÃ£o hÃ¡ cidadÃ£o com este ID.", err.Error())
		} else {
			log.Printf("[handler/GetProject -  Erro ao tentar buscar Project. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, dm, 0, 0)
}

//GetProjects ...
func GetProjects(w http.ResponseWriter, r *http.Request) {
	var dm model.Project
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetDemandas] -  Erro ao tentar abrir conexÃ£o. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	id, _ := strconv.Atoi(r.FormValue("id"))
	nome := r.FormValue("nome")

	dm.CodigoProject = int64(id)
	dm.Nome = nome

	projects, err := dm.GetProjects(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetProjects -  Nao ha Cidadao com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Nao ha Cidadao cadastrado.", err.Error())
		} else {
			log.Printf("[handler/GetProjects -  Erro ao tentar buscar Cidadao. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, projects, 0, 0)
}
