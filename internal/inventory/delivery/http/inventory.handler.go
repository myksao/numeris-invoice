package http

import (
	"invoice/internal/inventory"
	"invoice/internal/inventory/domain"
	"invoice/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Handler struct {
	Repo      inventory.Repo
	validator *validator.Validate
	router    *gin.RouterGroup
	logger    *zap.Logger
}

func NewHandler(router *gin.RouterGroup, repo inventory.Repo, logger *zap.Logger) *Handler {
	return &Handler{
		Repo:      repo,
		validator: validator.New(),
		router:    router,
		logger:    logger,
	}
}

func (h *Handler) Process() gin.HandlerFunc {
	return func(c *gin.Context) {
		var curr domain.InventoryProcess
		if err := c.ShouldBindJSON(&curr); err != nil {
			h.logger.Error("Error binding JSON", zap.Error(err))
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if err := h.validator.Struct(curr); err != nil {
			h.logger.Error("Error validating JSON", zap.Error(err))
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		id, err := h.Repo.Process(c, &curr)

		if err != nil {
			h.logger.Error("Error creating organisation", zap.Error(err))
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		utils.StandardResponse(c, http.StatusCreated, id)
	}
}
