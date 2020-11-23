package alfred

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	handler *gin.Engine
}

func NewHttpHandler() *HttpHandler {
	handler := gin.New()
	handler.Use(gin.Recovery(), GinLogger())

	return &HttpHandler{
		handler: gin.New(),
	}

}

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		end := time.Now()
		if raw != "" {
			path = path + "?" + raw
		}

		log.Printf("%s %s %s %s %d %s %d %s", c.Request.Host, c.ClientIP(), c.Request.Method, path, c.Writer.Status(), end.Sub(start), c.Writer.Size(), c.Errors.ByType(gin.ErrorTypePrivate).String())
		/*
			a.Infof(
				//path,
				c.Request.Host,
				"client", c.ClientIP(),
				"method", c.Request.Method,
				"path", path,
				"status", c.Writer.Status(),
				"latency", end.Sub(start),
				"size", c.Writer.Size(),
				"error", c.Errors.ByType(gin.ErrorTypePrivate).String(),
			)
		*/

	}
}

func (h *Handler) version(c *gin.Context) {
	c.JSON(200, gin.H{"version": version})
	//m.Infof("hello %s", c.Param("name"))
}

func (h *Handler) nodes(c *gin.Context) {
	c.JSON(200, gin.H{"nodes": version})
	//m.Infof("hello %s", c.Param("name"))
}

func (h *Handler) node(c *gin.Context) {
	c.JSON(200, gin.H{"node": c.Param("id")})
	//m.Infof("hello %s", c.Param("name"))
}
