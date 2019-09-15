package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"projects/back-end/api/db"
	"projects/back-end/model"
	"projects/back-end/util"
	"strconv"

	"github.com/gorilla/mux"
)

//InsertProject ...
func InsertProject(w http.ResponseWriter, r *http.Request) {
	var t util.App
	var p model.Project
	var d db.DB
	err := d.Connection()
	if err != nil {
		t.ResponseWithError(w, http.StatusInternalServerError, "Banco de Dados offline", "")
		return
	}
	db := d.DB
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}
	defer r.Body.Close()
	// err = p.InsertProject(db)
	if err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Erro ao inserir Situacao", "")
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, p, 0, 0)
}

//UpdateProject ...
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	var t util.App
	var p model.Project
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/UpdateProject] -  Erro ao tentar abrir conexão. Erro: %s", err.Error())
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
	if err := decoder.Decode(&p); err != nil {
		t.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload", "")
		return
	}
	defer r.Body.Close()
	p.ID = int64(id)
	// if err := p.UpdateProject(db); err != nil {
	// 	t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
	// 	return
	// }
	t.ResponseWithJSON(w, http.StatusOK, p, 0, 0)
}

//DeleteProject ...
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	var t util.App
	var p model.Project
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/DeleteProject -  Erro ao tentar abrir conexao. Erro: %s", err.Error())
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

	p.ID = int64(id)
	// if err := p.DeleteProject(db); err != nil {
	// 	log.Printf("[handler/DeleteProject -  Erro ao tentar deletar Cidadao. Erro: %s", err.Error())
	// 	t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
	// 	return
	// }
	t.ResponseWithJSON(w, http.StatusOK, p, 0, 0)
}

//GetProject ...
func GetProject(w http.ResponseWriter, r *http.Request) {
	var p model.Project
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetDemanda] -  Erro ao tentar abrir conexao. Erro: %s", err.Error())
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

	p.ID = int64(id)
	// err = p.GetProject(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetProject -  Nao ha Projeto com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Nao ha cidadao com este ID.", err.Error())
		} else {
			log.Printf("[handler/GetProject -  Erro ao tentar buscar Projeto. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, p, 0, 0)
}

//GetProjects ...
func GetProjects(w http.ResponseWriter, r *http.Request) {
	var p model.Project
	var t util.App
	var d db.DB
	err := d.Connection()
	if err != nil {
		log.Printf("[handler/GetDemandas] -  Erro ao tentar abrir conexao. Erro: %s", err.Error())
		return
	}
	db := d.DB
	defer db.Close()

	id, _ := strconv.Atoi(r.FormValue("id"))
	nome := r.FormValue("nome")

	p.ID = int64(id)
	p.Nome = nome

	projects := []model.Project{}
	// projects, err := p.GetProjects(db)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[handler/GetProjects -  Nao ha Projeto com este ID.")
			t.ResponseWithError(w, http.StatusInternalServerError, "Nao ha Projeto cadastrado.", err.Error())
		} else {
			log.Printf("[handler/GetProjects -  Erro ao tentar buscar Projeto. Erro: %s", err.Error())
			t.ResponseWithError(w, http.StatusInternalServerError, err.Error(), "")
		}
		return
	}
	t.ResponseWithJSON(w, http.StatusOK, projects, 0, 0)
}
