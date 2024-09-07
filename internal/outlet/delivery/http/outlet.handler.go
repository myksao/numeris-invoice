package http

import (
	"fmt"
	"invoice/internal/outlet"
	"invoice/internal/outlet/domain"
	"invoice/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Handler struct {
	Repo      outlet.Repo
	validator *validator.Validate
	router    *gin.RouterGroup
	logger    *zap.Logger
}

func NewHandler(router *gin.RouterGroup, repo outlet.Repo, logger *zap.Logger) *Handler {
	return &Handler{
		Repo:      repo,
		validator: validator.New(),
		router:    router,
		logger:    logger,
	}
}

func (h *Handler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var org domain.OutletReq
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
				utils.StandardResponse(c, http.StatusConflict, "Outlet already exists")
				return
			}
			h.logger.Error("Error creating outlet", zap.Error(err))
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
			h.logger.Error("Error retrieving organisation", zap.Error(err))
			utils.StandardResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if org == nil {
			utils.StandardResponse(c, http.StatusNotFound, fmt.Sprintf("Outlet with ID %s not found", id))
			return
		}

		utils.StandardResponse(c, http.StatusOK, org)

	}

}

func (h *Handler) RetrieveByOrgID() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		if id == "" {
			utils.StandardResponse(c, http.StatusBadRequest, "ID is required")
			return
		}

		type Request struct {
			Offset string `form:"offset" binding:"omitempty,number,gte=0"`
			Limit  string `form:"limit" binding:"omitempty,number,gte=1,required_with=offset"`
		}

		var req Request

		if err := c.ShouldBindQuery(&req); err != nil {
			h.logger.Error("Error binding JSON", zap.Error(err))
			utils.StandardResponse(c, http.StatusBadRequest, utils.CustomValidationError(err))
			return
		}

		org, err := h.Repo.RetrieveByOrgID(c, id, req.Limit, req.Offset)
		if err != nil {
			h.logger.Error("Error retrieving organisation", zap.Error(err))
			utils.StandardResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if org == nil {
			utils.StandardResponse(c, http.StatusNotFound, fmt.Sprintf("Org with ID %s not found", id))
			return
		}

		utils.StandardResponse(c, http.StatusOK, org)

	}

}
