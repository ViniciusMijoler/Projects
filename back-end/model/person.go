package model

import (
	"database/sql"
	"time"
)

//Address struct
type Address struct {
	ID            int64  `json:"id"`
	Status        int64  `json:"status"`
	DtCadastro    string `json:"dt_cadastro"`
	DtAtualizacao string `json:"dt_atualizacao"`
	Rua           string `json:"rua"`
	Numero        string `json:"numero"`
	Complemento   string `json:"complemento"`
	Bairro        string `json:"bairro"`
	Cidade        string `json:"cidade"`
	Estado        int64  `json:"estado"`
	Pais          int64  `json:"pais"`
}

//Person struct
type Person struct {
	ID            int64   `json:"id"`
	Status        int64   `json:"status"`
	DtCadastro    string  `json:"dt_cadastro"`
	DtAtualizacao string  `json:"dt_atualizacao"`
	TipoPessoa    int64   `json:"tipo_pessoa"`
	Nome          string  `json:"nome"`
	Apelido       string  `json:"apelido"`
	Email         string  `json:"email"`
	Telefone      string  `json:"telefone"`
	Celular       string  `json:"celular"`
	Facebook      string  `json:"facebook"`
	Twitter       string  `json:"twitter"`
	Instagram     string  `json:"instagram"`
	Linkedin      string  `json:"linkedin"`
	Endereco      Address `json:"endereco"`
}

//InsertPerson ...
func (c *Person) InsertPerson(db *sql.DB) error {
	dateNow := time.Now()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	{
		statement, err := tx.Prepare(`INSERT INTO endereco (id, status, dt_cadastro, rua, numero, complemento, bairro, cidade, estado, pais)
									VALUES (DEFAULT, DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8)
									RETURNING id`)
		if err != nil {
			tx.Rollback()
			return err
		}
		err = statement.QueryRow(dateNow, c.Endereco.Rua, c.Endereco.Numero, c.Endereco.Complemento,
			c.Endereco.Bairro, c.Endereco.Cidade, c.Endereco.Estado, c.Endereco.Pais).Scan(&c.Endereco.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	{
		statement, err := tx.Prepare(`INSERT INTO pessoa (id, status, dt_cadastro, tipo_pessoa, nome, apelido, email, 
															telefone, celular, facebook, twitter, instagram, linkedin, id_endereco)
									VALUES (DEFAULT, DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
									RETURNING id`)
		if err != nil {
			tx.Rollback()
			return err
		}
		err = statement.QueryRow(dateNow, c.TipoPessoa, c.Nome, c.Apelido, c.Email, c.Telefone,
			c.Celular, c.Facebook, c.Twitter, c.Instagram, c.Linkedin, c.Endereco.ID).Scan(&c.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

//UpdatePerson ...
// func (c *Person) UpdatePerson(db *sql.DB) error {
// 	statement, err := db.Prepare(`UPDATE CIDADAO
// 									SET NOME = ?,
// 										CPF = ?,
// 										DATANASCIMENTO = ?,
// 										TELEFONE = ?,
// 									WHERE CODIGO = ?`)

// 	if err != nil {
// 		return err
// 	}

// 	_, err = statement.Exec(c.Nome, c.Cpf, c.DataNascimento, c.Telefone, c.StatusPerson, c.CodigoPerson)

// 	return err
// }

//GetPerson ...
// func (c *Person) GetPerson(db *sql.DB) error {
// 	err := db.QueryRow(`SELECT CODIGO, NOME, CPF, DATANASCIMENTO, TELEFONE, STATUS
// 					FROM CIDADAO
// 					WHERE CODIGO =  ?`, c.CodigoPerson).Scan(&c.CodigoPerson, &c.Nome, &c.Cpf, &c.DataNascimento, &c.Telefone, &c.StatusPerson)
// 	if err != nil {
// 		return err
// 	}

// 	return err
// }

//GetPersons ...
// func (c *Person) GetPersons(db *sql.DB) ([]Person, error) {
// 	var values []interface{}
// 	var where []string

// 	if c.CodigoPerson != 0 {
// 		where = append(where, "CODIGO = ?")
// 		values = append(values, c.CodigoPerson)
// 	}

// 	if c.Nome != "" {
// 		where = append(where, "DESCRICAO = ?")
// 		values = append(values, c.Nome)
// 	}

// 	rows, err := db.Query(`SELECT CODIGO, NOME, CPF, DATANASCIMENTO, TELEFONE, STATUS
// 					FROM CIDADAO
// 					WHERE STATUS = 1 `+strings.Join(where, " AND "), values...)

// 	if err != nil {
// 		return nil, err
// 	}

// 	projects := []Person{}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var cid Person
// 		if err = rows.Scan(&cid.CodigoPerson, &cid.Nome, &cid.Cpf, &cid.DataNascimento, &cid.Telefone, &cid.StatusPerson); err != nil {
// 			return nil, err
// 		}
// 		projects = append(projects, cid)
// 	}
// 	return projects, nil
// }

//DeletePerson ...
// func (c *Person) DeletePerson(db *sql.DB) error {
// 	statement, err := db.Prepare(`UPDATE CIDADAO
// 									SET STATUS = 1
// 									WHERE CODIGO = ?`)

// 	if err != nil {
// 		return err
// 	}

// 	_, err = statement.Exec(0, c.CodigoPerson)

// 	return err
// }
