package service

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/felipeksw/goexpert-fullcycle-multithreading/internal/adapter/webclient"
	"github.com/felipeksw/goexpert-fullcycle-multithreading/internal/dto"
)

func GetLocalidadeByBrasilapi(http webclient.HttpRequest, cep string, c chan<- dto.LocalidadeDto) {
	/*
		Tempo médio de resposta da ViaCEP foi de aproximadamente 550ms
		Tempo médio de resposta da BrasilAPI variou de 60ms a 400ms
	*/
	//time.Sleep(1600 * time.Millisecond)

	req, err := http.Request(fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep))
	if err != nil {
		slog.Error("[brasilapi webserver client]", "error", err.Error())
		c <- dto.LocalidadeDto{Erro: err}
	}

	var a dto.BrasilApiDto
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
