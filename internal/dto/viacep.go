package dto

type ViaCepDto struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (v *ViaCepDto) ToLocalidade() *LocalidadeDto {
	return &LocalidadeDto{
		Cep:        v.Cep,
		Estado:     v.Uf,
		Cidade:     v.Localidade,
		Logradouro: v.Logradouro,
		Servico:    "viacep",
		Erro:       nil,
	}
}
