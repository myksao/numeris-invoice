package http

func (handler *Handler) Route() {
	variant := handler.router.Group("/variants")
	{
		variant.POST("", handler.Create())
		variant.GET("/item/:id", handler.RetrieveByItemID())
		variant.GET("/:id", handler.RetrieveByID())
		variant.GET("/:id/measure", handler.RetrieveMeasureByVariantID())

	}
}
