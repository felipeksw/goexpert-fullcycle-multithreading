package dto

type Localidade interface {
	ToLocalidade() *LocalidadeDto
}

type LocalidadeDto struct {
	Cep        string
	Estado     string
	Cidade     string
	Logradouro string
	Servico    string
	Erro       error
}
