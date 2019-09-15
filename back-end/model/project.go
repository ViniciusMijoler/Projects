package model

//Project struct
type Project struct {
	ID             int64  `json:"id"`
	Status         int64  `json:"status"`
	DtCadastro     string `json:"dt_cadastro"`
	DtAtualizacao  string `json:"dt_atualizacao"`
	Nome           string `json:"nome"`
	IDEmpresa      int64  `json:"id_empresa"`
	PalavrasChaves string `json:"palavras_chaves"`
	AreaProjeto    string `json:"area_projeto"`
	DataLimite     string `json:"data_limite"`
	Descricao      string `json:"descricao"`
}
