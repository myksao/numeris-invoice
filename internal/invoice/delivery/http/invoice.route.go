package http

func (handler *Handler) Route() {
	invoice := handler.router.Group("/invoices")
	{
		invoice.POST("", handler.Create())
		invoice.PATCH("/:id", handler.UpdateStatus())
		invoice.GET("/:id", handler.RetrieveByID())
		invoice.GET("/:id/boqs", handler.FetchBoqs())
		invoice.GET("/summary/:id", handler.RetrieveSummary())
		invoice.PATCH("/:id/boqs", handler.CreateInvoiceBoq())
	}
}
