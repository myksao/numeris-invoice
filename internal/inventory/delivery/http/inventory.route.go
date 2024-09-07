package http

func (handler *Handler) Route() {
	invoice := handler.router.Group("/inventory")
	{
		invoice.POST("/process", handler.Process())
	}
}
