package dto

type BrasilApiDto struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func (b *BrasilApiDto) ToLocalidade() *LocalidadeDto {
	return &LocalidadeDto{
		Cep:        b.Cep,
		Estado:     b.State,
		Cidade:     b.City,
		Logradouro: b.Street,
		Servico:    "brasilapi",
		Erro:       nil,
	}
}
