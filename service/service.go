package service

import (
  "fmt"
  "strconv"
  "net/http"
  "github.com/lunchtime-labs/redchannels/blacklist"
  "github.com/lunchtime-labs/redchannels/Godeps/_workspace/src/github.com/gin-gonic/gin"
)

type Config struct {
  Env           string
  Port          int
  LoaderIoToken string
  Blacklist     string
}

type BlacklistService struct {
}

func (s *BlacklistService) Run(config Config) error {
  router := gin.New()
  router.Use(gin.Logger())

  if config.Env == "production" {
    gin.SetMode(gin.ReleaseMode)
  }

  _, err := blacklist.PopulateBlacklist(config.Blacklist)
  if err != nil {
    return err
  }

  // routes
  router.GET("/", enumerateHeaders)
  // loader.io verification
  router.GET("loaderio-" + config.LoaderIoToken + "/", func(c *gin.Context) {
    c.String(http.StatusOK, "loaderio-" + config.LoaderIoToken)
  })
  router.Run(":" + strconv.Itoa(config.Port))

  return nil
}

func enumerateHeaders(context *gin.Context) {
  for name, value := range context.Request.Header {
    fmt.Println(name, value)
  }
}
