package http

func (handler *Handler) Route() {
	org := handler.router.Group("/outlets")
	{
		org.POST("", handler.Create())
		org.GET("/:id", handler.RetrieveByID())
		org.GET("/org/:id", handler.RetrieveByOrgID())
	}
}
