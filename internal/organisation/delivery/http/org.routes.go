package http

func (handler *Handler) Route() {
	org := handler.router.Group("/orgs")
	{
		org.POST("", handler.Create())
		org.GET("/:id", handler.RetrieveByID())
	}
}
