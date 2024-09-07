package http

import (
	"invoice/internal/currency"
	"invoice/internal/currency/domain"
	"invoice/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Handler struct {
	Repo      currency.Repo
	validator *validator.Validate
	router    *gin.RouterGroup
	logger    *zap.Logger
}

func NewHandler(router *gin.RouterGroup, repo currency.Repo, logger *zap.Logger) *Handler {
	return &Handler{
		Repo:      repo,
		validator: validator.New(),
		router:    router,
		logger:    logger,
	}
}

func (h *Handler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var curr domain.CurrencyReq
		if err := c.ShouldBindJSON(&curr); err != nil {
			h.logger.Error("Error binding JSON", zap.Error(err))
			utils.StandardResponse(c, http.StatusBadRequest, utils.CustomValidationError(err))
			return
		}

		if err := h.validator.Struct(curr); err != nil {
			h.logger.Error("Error validating JSON", zap.Error(err))
			utils.StandardResponse(c, http.StatusBadRequest, utils.CustomValidationError(err))
			return
		}

		id, err := h.Repo.Create(c, &curr)

		if err != nil {
			if err.Error() == "duplicate key value violates unique constraint" {
				utils.StandardResponse(c, http.StatusConflict, "currency already exists")
				return
			}
			h.logger.Error("Error creating currency", zap.Error(err))
			utils.StandardResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.StandardResponse(c, http.StatusCreated, id)
	}
}

func (h *Handler) Retrieve() gin.HandlerFunc {
	return func(c *gin.Context) {
		var page utils.PaginationReq

		if err := c.ShouldBindQuery(&page); err != nil {
			h.logger.Error("Error getting pagination", zap.Error(err))
			utils.StandardResponse(c, http.StatusBadRequest, utils.CustomValidationError(err))
			return
		}

		currencies, err := h.Repo.Retrieve(c, &page)
		if err != nil {
			h.logger.Error("Error retrieving currencies", zap.Error(err))
			utils.StandardResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if len(currencies) == 0 {
			utils.StandardResponse(c, http.StatusNotFound, "No currencies found")
			return
		}

		utils.StandardResponse(c, http.StatusOK, currencies)
	}
}
