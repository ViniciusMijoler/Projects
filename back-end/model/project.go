package model

import (
	"database/sql"
	"strings"
)

//Project struct
type Project struct {
	CodigoProject  int64  `json:"codigo"`
	Nome           string `json:"nome"`
	Cpf            string `json:"cpf"`
	DataNascimento string `json:"dataNascimento"`
	Telefone       string `json:"telefone"`
	Senha          string `json:"senha"`
	StatusProject  int64  `json:"status"`
}

//InsertProject ...
func (c *Project) InsertProject(db *sql.DB) error {
	statement, err := db.Prepare(`INSERT INTO CIDADAO (NOME, CPF, DATANASCIMENTO, TELEFONE, STATUS)
								VALUES
								(?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	res, err := statement.Exec(c.Nome, c.Cpf, c.DataNascimento, c.Telefone, c.StatusProject)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	c.CodigoProject = id
	return nil
}

//UpdateProject ...
func (c *Project) UpdateProject(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE CIDADAO
									SET NOME = ?,
										CPF = ?,
										DATANASCIMENTO = ?,
										TELEFONE = ?, 
									WHERE CODIGO = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(c.Nome, c.Cpf, c.DataNascimento, c.Telefone, c.StatusProject, c.CodigoProject)

	return err
}

//GetProject ...
func (c *Project) GetProject(db *sql.DB) error {
	err := db.QueryRow(`SELECT CODIGO, NOME, CPF, DATANASCIMENTO, TELEFONE, STATUS
					FROM CIDADAO
					WHERE CODIGO =  ?`, c.CodigoProject).Scan(&c.CodigoProject, &c.Nome, &c.Cpf, &c.DataNascimento, &c.Telefone, &c.StatusProject)
	if err != nil {
		return err
	}

	return err
}

//GetProject ...
func (c *Project) GetProjects(db *sql.DB) ([]Project, error) {
	var values []interface{}
	var where []string

	if c.CodigoProject != 0 {
		where = append(where, "CODIGO = ?")
		values = append(values, c.CodigoProject)
	}

	if c.Nome != "" {
		where = append(where, "DESCRICAO = ?")
		values = append(values, c.Nome)
	}

	rows, err := db.Query(`SELECT CODIGO, NOME, CPF, DATANASCIMENTO, TELEFONE, STATUS
					FROM CIDADAO
					WHERE STATUS = 1 `+strings.Join(where, " AND "), values...)

	if err != nil {
		return nil, err
	}

	projects := []Project{}
	defer rows.Close()
	for rows.Next() {
		var cid Project
		if err = rows.Scan(&cid.CodigoProject, &cid.Nome, &cid.Cpf, &cid.DataNascimento, &cid.Telefone, &cid.StatusProject); err != nil {
			return nil, err
		}
		projects = append(projects, cid)
	}
	return projects, nil
}

//DeleteProject ...
func (c *Project) DeleteProject(db *sql.DB) error {
	statement, err := db.Prepare(`UPDATE CIDADAO
									SET STATUS = 1
									WHERE CODIGO = ?`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(0, c.CodigoProject)

	return err
}
