package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/homepage408/syncra/internal/domains/auth/application/usecase"
	"github.com/homepage408/syncra/pkg/logger"
)

type Handler struct {
	service *usecase.Service
	log     logger.Logger
}

func New(service *usecase.Service, log logger.Logger) *Handler {
	return &Handler{service: service, log: log}
}

func (h *Handler) Login(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"message": "login endpoint"})
}

func (h *Handler) Register(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"message": "register endpoint"})
}

func (h *Handler) RefreshToken(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"message": "refresh endpoint"})
}
