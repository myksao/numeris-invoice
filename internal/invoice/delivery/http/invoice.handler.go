package http

import (
	"invoice/internal/invoice"
	"invoice/internal/invoice/domain"
	"invoice/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Handler struct {
	Repo      invoice.Repo
	BoqRepo   invoice.BoqRepo
	validator *validator.Validate
	router    *gin.RouterGroup
	logger    *zap.Logger
}

func NewHandler(router *gin.RouterGroup, repo invoice.Repo, logger *zap.Logger, rboq invoice.BoqRepo) *Handler {
	return &Handler{
		Repo:      repo,
		validator: validator.New(),
		router:    router,
		logger:    logger,
		BoqRepo:   rboq,
	}
}

func (h *Handler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var curr domain.InvoiceReq
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
			h.logger.Error("Error creating organisation", zap.Error(err))
			utils.StandardResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.StandardResponse(c, http.StatusCreated, id)
	}
}

func (h *Handler) UpdateStatus() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		if id == "" {
			h.logger.Error("Error retrieving ID")
			c.JSON(400, gin.H{"error": "ID is required"})
			return
		}

		var curr domain.InvoiceStatusReq

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

		_, err := h.Repo.UpdateStatus(c, &curr, &id)

		if err != nil {
			h.logger.Error("Error updating status", zap.Error(err))
			utils.StandardResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.StandardResponse(c, http.StatusOK, id)
	}
}

func (h *Handler) RetrieveByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			utils.StandardResponse(c, http.StatusBadRequest, "ID is required")
			return
		}
		curr, err := h.Repo.RetrieveByID(c, id)
		if err != nil {
			h.logger.Error("Error retrieving organisation", zap.Error(err))
			utils.StandardResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if curr == nil {
			utils.StandardResponse(c, http.StatusOK, []string{})
			return
		}

		utils.StandardResponse(c, http.StatusOK, curr)

	}
}

func (h *Handler) FetchBoqs() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			utils.StandardResponse(c, http.StatusBadRequest, "ID is required")
			return
		}
		boqs, err := h.Repo.FetchBoqs(c, id)
		if err != nil {
			h.logger.Error("Error retrieving Boqs", zap.Error(err))
			utils.StandardResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if boqs == nil {
			utils.StandardResponse(c, http.StatusOK, []string{})
			return
		}

		utils.StandardResponse(c, http.StatusOK, boqs)

	}
}

func (h *Handler) RetrieveSummary() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			utils.StandardResponse(c, http.StatusBadRequest, "ID is required")
			return
		}
		summary, err := h.Repo.RetrieveSummary(c, id)
		if err != nil {
			h.logger.Error("Error retrieving summary", zap.Error(err))
			utils.StandardResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.StandardResponse(c, http.StatusOK, summary)

	}
}

func (h *Handler) CreateInvoiceBoq() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			utils.StandardResponse(c, http.StatusBadRequest, "Invoice ID is required")
			return
		}
		var boq []*domain.UpdateInvoiceBoqReq
		if err := c.ShouldBindJSON(&boq); err != nil {
			h.logger.Error("Error binding JSON", zap.Error(err))
			utils.StandardResponse(c, http.StatusBadRequest, utils.CustomValidationError(err))
			return
		}

		var boqm struct {
			Boq []*domain.UpdateInvoiceBoqReq `json:"boq" validate:"required"`
		}

		boqm.Boq = boq

		if err := h.validator.Struct(boqm); err != nil {
			h.logger.Error("Error validating JSON", zap.Error(err))
			utils.StandardResponse(c, http.StatusBadRequest, utils.CustomValidationError(err))
			return
		}

		res, err := h.BoqRepo.Create(c, id, boq)
		if err != nil {
			h.logger.Error("Error creating invoice Boq", zap.Error(err))
			utils.StandardResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.StandardResponse(c, http.StatusCreated, res.ID)
	}
}
