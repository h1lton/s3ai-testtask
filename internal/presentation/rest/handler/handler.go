package handler

import (
	"s3ai-testtask/internal/domain/interfaces"
)

type Handler struct {
	atmService interfaces.ATMService
}

func New(
	atmService interfaces.ATMService,
) *Handler {
	return &Handler{
		atmService: atmService,
	}
}
