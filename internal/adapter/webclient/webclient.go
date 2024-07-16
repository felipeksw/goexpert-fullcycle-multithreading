package webclient

import (
	"context"
	"io"
	"log/slog"
	"net/http"
)

type HttpRequest struct {
	client  *http.Client
	request *http.Request
	ctx     context.Context
}

func NewHttpRequest(ctx context.Context) HttpRequest {
	slog.Debug("[NewHttpRequest created]")
	return HttpRequest{
		request: nil,
		client:  http.DefaultClient,
		ctx:     ctx,
	}
}

func (h *HttpRequest) Request(url string) (*HttpRequest, error) {

	slog.Debug("[string]", "url", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		slog.Error("[http.NewRequest Failed]", "error", err.Error())
		return nil, err
	}
	slog.Debug("[http.NewRequest Created]", "req", *req)

	if h.ctx != nil {
		req = req.WithContext(h.ctx)
		slog.Debug("[Context Added]")
	}

	return &HttpRequest{
		request: req,
		client:  h.client,
		ctx:     h.ctx,
	}, nil
}

func (h *HttpRequest) Do(ret func([]byte) error) error {

	resp, err := h.client.Do(h.request)
	if err != nil {
		slog.Error("[http.Client.Do failed]", "error", err.Error())
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	defer func() {
		body = nil
	}()
	if err != nil {
		slog.Error("[io.ReadAll failed]", "error", err.Error())
		return err
	}

	return ret(body)
}
