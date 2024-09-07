package http

func (handler *Handler) Route() {
	currency := handler.router.Group("/currencies")
	{
		currency.POST("", handler.Create())
		currency.GET("", handler.Retrieve())

	}
}
