package service

import (
	"encoding/json"
	"log/slog"

	"github.com/felipeksw/goexpert-fullcycle-multithreading/internal/adapter/webclient"
	"github.com/felipeksw/goexpert-fullcycle-multithreading/internal/dto"
)

func GetLocalidade(http webclient.HttpRequest, url string, c chan<- dto.LocalidadeDto, a dto.Localidade) {

	req, err := http.Request(url)
	if err != nil {
		slog.Error("[brasilapi webserver client]", "error", err.Error())
		c <- dto.LocalidadeDto{Erro: err}
	}

	//var a dto.BrasilApiDto
	err = req.Do(func(p []byte) error {
		err = json.Unmarshal(p, &a)
		return err
	})
	if err != nil {
		slog.Error("[brasilapi do]", "error", err.Error())
		c <- dto.LocalidadeDto{Erro: err}
	}
	slog.Debug("[struct]", "BrasilApiDto", a)

	c <- *a.ToLocalidade()
}
