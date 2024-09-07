package organisation

import "github.com/gin-gonic/gin"

type Delivery interface {
	Create() gin.HandlerFunc
	RetrieveByID() gin.HandlerFunc
}
