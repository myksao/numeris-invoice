package item

import "github.com/gin-gonic/gin"

type Handler interface {
	Create() gin.HandlerFunc
	RetrieveByOutletID() gin.HandlerFunc
	RetrieveByID() gin.HandlerFunc
}
