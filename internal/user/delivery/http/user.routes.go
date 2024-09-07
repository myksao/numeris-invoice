package http

func (handler *Handler) Route() {
	user := handler.router.Group("/users")
	{
		user.POST("", handler.Create())
		user.GET("/outlet/:id", handler.RetrieveByOutletID())
		user.GET("/:id", handler.RetrieveByID())

	}
}
