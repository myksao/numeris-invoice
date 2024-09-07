package note

import "github.com/gin-gonic/gin"

type Handler interface {
	Create() gin.HandlerFunc
	RetrieveByEntityID() gin.HandlerFunc
}
