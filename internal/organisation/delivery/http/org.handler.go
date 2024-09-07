package http

import (
	"fmt"
	"invoice/internal/organisation"
	domain "invoice/internal/organisation/domain"
	"invoice/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Handler struct {
	Repo      organisation.Repo
	validator *validator.Validate
	router    *gin.RouterGroup
	logger    *zap.Logger
}

func NewHandler(router *gin.RouterGroup, repo organisation.Repo, logger *zap.Logger) *Handler {
	return &Handler{
		Repo:      repo,
		validator: validator.New(),
		router:    router,
		logger:    logger,
	}
}

func (h *Handler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var org domain.OrganisationReq
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
			if err.Error() == "duplicate key value violates unique constraint" {
				utils.StandardResponse(c, http.StatusConflict, "Organisation already exists")
				return
			}
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
			utils.StandardResponse(c, http.StatusBadRequest, "ID is required")
			return
		}
		org, err := h.Repo.RetrieveByID(c, id)
		if err != nil {
			if err.Error() == "Organisation not found" {
				utils.StandardResponse(c, http.StatusNotFound, fmt.Sprintf("Organisation with ID %s not found", id))
				return
			}
			h.logger.Error("Error retrieving organisation", zap.Error(err))
			utils.StandardResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if org == nil {
			utils.StandardResponse(c, http.StatusNotFound, fmt.Sprintf("Organisation with ID %s not found", id))
			return
		}

		utils.StandardResponse(c, http.StatusOK, org)

	}

}
