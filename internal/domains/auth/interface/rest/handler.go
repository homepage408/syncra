package rest

import (
	"net/http"

	constant "github.com/homepage408/syncra/internal/shared/constant"
	request "github.com/homepage408/syncra/pkg/request"
	response "github.com/homepage408/syncra/pkg/response"
	validation "github.com/homepage408/syncra/pkg/validation"

	"github.com/gin-gonic/gin"
	"github.com/homepage408/syncra/internal/domains/auth/domain/entity"
	"github.com/homepage408/syncra/internal/domains/auth/usecase"
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

	var userAgent = c.GetHeader("User-Agent")

	var reqeust request.LoginUserRequest
	if err := c.ShouldBindJSON(&reqeust); err != nil {
		formatErrors := validation.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: "Validation error",
			Errors:  formatErrors,
		})
		return
	}

	data, err := h.service.Login(c.Request.Context(), reqeust.Email, reqeust.Password, userAgent)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: "Failed to login user",
			Errors:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Code:    http.StatusOK,
		Success: true,
		Message: "User logged in successfully",
		Data:    data,
	})
}

func (h *Handler) Register(c *gin.Context) {
	var reqeust request.RegisterUserRequest
	if err := c.ShouldBindJSON(&reqeust); err != nil {
		formatErrors := validation.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: "Validation error",
			Errors:  formatErrors,
		})
		return
	}

	data, err := h.service.Register(c.Request.Context(), &entity.User{
		Email:        reqeust.Email,
		Username:     reqeust.Username,
		FullName:     reqeust.FullName,
		PasswordHash: reqeust.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: "Failed to register user",
			Errors:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Code:    http.StatusOK,
		Success: true,
		Message: "User registered successfully",
		Data:    data,
	})
}

func (h *Handler) GetActiveSessions(c *gin.Context) {
	tokenVal, exist := c.Get(constant.TokenContextKey)
	if !exist {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Message: "Token Not Found",
			Code:    http.StatusUnauthorized,
			Success: false,
		})
		return
	}

	token := tokenVal.(string)
	data, err := h.service.GetAcctiveSessions(c.Request.Context(), token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusUnauthorized,
			Success: false,
		})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Code:    http.StatusOK,
		Success: true,
		Message: "User Sessions",
		Data:    data,
	})
}

func (h *Handler) RefreshToken(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"message": "refresh endpoint"})
}
