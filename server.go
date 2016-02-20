package redchannels

import (
	"github.com/gin-gonic/gin"
)

type Server interface {
  Start(port string, env string) *gin.Engine
}

func Start(port string, env string) *gin.Engine {
  server := gin.New()
  server.Use(gin.Logger())

  if env == "production" {
    gin.SetMode(gin.ReleaseMode)
  }

  server.Run(":" + port)

  return server
}
