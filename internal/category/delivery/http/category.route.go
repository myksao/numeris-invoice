package http

func (handler *Handler) Route() {
	category := handler.router.Group("/categories")
	{
		category.POST("", handler.Create())
		category.GET("/outlet/:id", handler.RetrieveByOutletID())
		category.GET("/:id", handler.RetrieveByID())

	}
}
