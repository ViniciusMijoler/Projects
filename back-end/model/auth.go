package model

import "database/sql"

//User struct
type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

//LoginResponse struct
type LoginResponse struct {
	UserName string `json:"user_name"`
	IDPessoa string `json:"id_pessoa"`
	Token    string `json:"token"`
}

//Login ...
func (u *User) Login(db *sql.DB) (interface{}, error) {
	// verify user credentials
	var user LoginResponse
	err := db.QueryRow(`SELECT user_name, id_pessoa, MD5(CONCAT(password, id, 'project')) as token
					FROM usuario
					WHERE user_name = $1 AND password = MD5($2)`, u.UserName, u.Password).Scan(&user.UserName, &user.IDPessoa, &user.Token)
	if err != nil {
		return nil, err
	}
	return user, nil
}
