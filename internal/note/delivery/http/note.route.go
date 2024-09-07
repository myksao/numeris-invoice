package http

func (handler *Handler) Route() {
	customer := handler.router.Group("/notes")
	{
		customer.POST("", handler.Create())
		customer.GET("/entity/:entity", handler.RetrieveByEntity())
		customer.GET("/entity/:entity/:id", handler.RetrieveByEntityID())
		customer.GET("/:id", handler.RetrieveByID())

	}
}
