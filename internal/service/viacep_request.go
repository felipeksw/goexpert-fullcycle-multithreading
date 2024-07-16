package service

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/felipeksw/goexpert-fullcycle-multithreading/internal/adapter/webclient"
	"github.com/felipeksw/goexpert-fullcycle-multithreading/internal/dto"
)

func GetLocalidadeByViacep(http webclient.HttpRequest, cep string, c chan<- dto.LocalidadeDto) {

	req, err := http.Request(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		slog.Error("[viacep webserver client]", "error", err.Error())
		c <- dto.LocalidadeDto{Erro: err}
	}

	var a dto.ViaCepDto
	err = req.Do(func(p []byte) error {
		err = json.Unmarshal(p, &a)
		return err
	})
	if err != nil {
		slog.Error("[viacep do]", "error", err.Error())
		c <- dto.LocalidadeDto{Erro: err}
	}
	slog.Debug("[struct]", "ViaCepDto", a)

	c <- *a.ToLocalidade()
}
