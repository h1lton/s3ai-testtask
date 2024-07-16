package rest

import (
	"net/http"
	"s3ai-testtask/internal/domain/interfaces"
	"s3ai-testtask/internal/infrastructure/config"
	"s3ai-testtask/internal/presentation/rest/handler"
)

func NewServer(
	atmService interfaces.ATMService,
	cfg *config.Config,
) (*http.Server, error) {
	h := handler.New(atmService)
	r, err := NewRouter(h)
	if err != nil {
		return nil, err
	}

	srv := &http.Server{
		Addr:    cfg.Server.Addr,
		Handler: r,
	}

	return srv, nil
}
