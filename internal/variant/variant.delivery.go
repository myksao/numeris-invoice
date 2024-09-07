package variant

import "github.com/gin-gonic/gin"

type Handler interface {
	Create() gin.HandlerFunc
	RetrieveByID() gin.HandlerFunc
	RetrieveByItemID() gin.HandlerFunc
}
