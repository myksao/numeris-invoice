package http

func (handler *Handler) Route() {
	customer := handler.router.Group("/customers")
	{
		customer.POST("", handler.Create())
		customer.GET("/outlet/:id", handler.RetrieveByOutletID())
		customer.GET("/:id", handler.RetrieveByID())

	}
}
