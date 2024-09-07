package http

import (
	"errors"
	"invoice/internal/variant"
	"invoice/internal/variant/domain"
	"invoice/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Handler struct {
	Repo      variant.Repo
	validator *validator.Validate
	router    *gin.RouterGroup
	logger    *zap.Logger
}

func NewHandler(router *gin.RouterGroup, repo variant.Repo, logger *zap.Logger) *Handler {
	return &Handler{
		Repo:      repo,
		validator: validator.New(),
		router:    router,
		logger:    logger,
	}
}

func (h *Handler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var org domain.VariantReq
		if err := c.ShouldBindJSON(&org); err != nil {
			h.logger.Error("Error binding JSON", zap.Error(err))
			utils.StandardResponse(c, http.StatusBadRequest, utils.CustomValidationError(err))
			return
		}

		if err := h.validator.Struct(org); err != nil {
			h.logger.Error("Error validating JSON", zap.Error(err))
			utils.StandardResponse(c, http.StatusBadRequest, utils.CustomValidationError(err))
			return
		}

		id, err := h.Repo.Create(c, &org)

		if err != nil {
			h.logger.Error("Error creating organisation", zap.Error(err))
			utils.StandardResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.StandardResponse(c, http.StatusCreated, id)
	}
}

func (h *Handler) RetrieveByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			h.logger.Error("Error retrieving organisation", zap.Error(errors.New("ID is required")))
			utils.StandardResponse(c, http.StatusBadRequest, "ID is required")
			return
		}

		vart, err := h.Repo.RetrieveByID(c, id)
		if err != nil {
			h.logger.Error("Error retrieving organisation", zap.Error(err))
			utils.StandardResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if vart == nil {
			utils.StandardResponse(c, http.StatusOK, []string{})
			return
		}

		utils.StandardResponse(c, http.StatusOK, vart)

	}
}

func (h *Handler) RetrieveByItemID() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req domain.VariantItemReq

		if err := c.ShouldBindUri(&req); err != nil {
			h.logger.Error("Error binding JSON", zap.Error(err))
			utils.StandardResponse(c, http.StatusBadRequest, utils.CustomValidationError(err))
			return
		}

		var page utils.PaginationReq

		if err := c.ShouldBindQuery(&page); err != nil {
			h.logger.Error("Error binding JSON", zap.Error(err))
			utils.StandardResponse(c, http.StatusBadRequest, utils.CustomValidationError(err))
			return
		}

		org, err := h.Repo.RetrieveByItemID(c, &req, &page)
		if err != nil {
			h.logger.Error("Error retrieving organisation", zap.Error(err))
			utils.StandardResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if org == nil {
			utils.StandardResponse(c, http.StatusOK, []string{})
			return
		}

		utils.StandardResponse(c, http.StatusOK, org)

	}
}

func (h *Handler) RetrieveMeasureByVariantID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			h.logger.Error("Error retrieving variant item measure", zap.Error(errors.New("ID is required")))
			utils.StandardResponse(c, http.StatusBadRequest, "ID is required")
			return
		}

		measure, err := h.Repo.RetrieveMeasureByVariantID(c, id)
		if err != nil {
			h.logger.Error("Error retrieving variant item", zap.Error(err))
			utils.StandardResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if measure == nil {
			utils.StandardResponse(c, http.StatusOK, []string{})
			return
		}

		utils.StandardResponse(c, http.StatusOK, measure)

	}
}
