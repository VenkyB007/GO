package api

import (
	"github.com/gin-gonic/gin"
)

// Method http Method in enum
type Method int

const (
	GET Method = iota
	POST
	PUT
	DELETE
	PATCH
)

// Operations contains the parameters of every endpoint exposed by studentapp
type Operations struct {
	Method     string
	Path       string
	Handler    gin.HandlerFunc
	Middleware []gin.HandlerFunc
}
