package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"time"

	"github.com/felipeksw/goexpert-fullcycle-multithreading/internal/adapter/webclient"
	"github.com/felipeksw/goexpert-fullcycle-multithreading/internal/dto"
	"github.com/felipeksw/goexpert-fullcycle-multithreading/internal/service"
)

func main() {

	slog.SetLogLoggerLevel(slog.LevelDebug)

	cep := os.Args[1]
	var re = regexp.MustCompile(`^[0-9]{8}$`)
	if !re.MatchString(cep) {
		panic("o cep deve conter 8 digitos numéricos")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	http := webclient.NewHttpRequest(ctx)

	ch := make(chan dto.LocalidadeDto, 2)

	//go service.GetLocalidadeByViacep(http, cep, ch)
	//go service.GetLocalidadeByBrasilapi(http, cep, ch)

	go service.GetLocalidade(http, fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep), ch, &dto.BrasilApiDto{})
	go service.GetLocalidade(http, fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep), ch, &dto.ViaCepDto{})

	select {
	case localidade := <-ch:
		fmt.Printf("--------\n")
		fmt.Printf("Serviço: %s \n", localidade.Servico)
		fmt.Printf("Endereço: %s, %s/%s - %s \n", localidade.Logradouro, localidade.Cidade, localidade.Estado, localidade.Cep)
		fmt.Printf("--------\n")
		return
	case <-ctx.Done():
		slog.Error("[request timeout]", "error", ctx.Err())
		return
	}
}
