package http

func (handler *Handler) Route() {
	item := handler.router.Group("/items")
	{
		item.POST("", handler.Create())
		item.GET("/outlet/:id", handler.RetrieveByOutletID())
		item.GET("/:id", handler.RetrieveByID())

	}
}
