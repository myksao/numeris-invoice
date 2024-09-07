package http

import (
	"errors"
	"invoice/internal/category"
	"invoice/internal/category/domain"
	"invoice/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Handler struct {
	Repo      category.Repo
	validator *validator.Validate
	router    *gin.RouterGroup
	logger    *zap.Logger
}

func NewHandler(router *gin.RouterGroup, repo category.Repo, logger *zap.Logger) *Handler {
	return &Handler{
		Repo:      repo,
		validator: validator.New(),
		router:    router,
		logger:    logger,
	}
}

func (h *Handler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var org domain.CategoryReq
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
			h.logger.Error("Error retrieving category", zap.Error(errors.New("ID is required")))
			utils.StandardResponse(c, http.StatusBadRequest, "ID is required")
			return
		}

		category, err := h.Repo.RetrieveByID(c, id)
		if err != nil {
			h.logger.Error("Error retrieving category", zap.Error(err))
			utils.StandardResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if category == nil {
			utils.StandardResponse(c, http.StatusNotFound, "Category not found")
			return
		}

		utils.StandardResponse(c, http.StatusOK, category)
	}
}

func (h *Handler) RetrieveByOutletID() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req domain.CategoryOutletReq

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

		customers, err := h.Repo.RetrieveByOutletID(c, &req, &page)
		if err != nil {
			h.logger.Error("Error retrieving categories", zap.Error(err))
			utils.StandardResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if len(customers) == 0 {
			utils.StandardResponse(c, http.StatusOK, []string{})
			return
		}

		utils.StandardResponse(c, http.StatusOK, customers)
	}
}
