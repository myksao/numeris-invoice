package http

func (handler *Handler) Route() {
	bank_accounts := handler.router.Group("/bank-accounts")
	{
		bank_accounts.POST("", handler.Create())
		bank_accounts.GET("/outlet/:id", handler.RetrieveByOutletID())
		bank_accounts.GET("/:id", handler.RetrieveByID())

	}
}
