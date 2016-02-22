package service

import (
  "fmt"
  "strconv"
  "sort"
  "net/http"
  "net/url"
  "github.com/lunchtime-labs/redchannels/blacklist"
  "github.com/lunchtime-labs/redchannels/Godeps/_workspace/src/github.com/gin-gonic/gin"
)

type Config struct {
  Env           string
  Port          int
  LoaderIoToken string
  BlacklistPath string
  UpstreamUrl   string
}

type BlacklistService struct {
}

var (
	Blacklist blacklist.Blacklist
	err       error
)

func (s *BlacklistService) Run(config Config) error {
  router := gin.New()
  router.Use(gin.Logger())

  if config.Env == "production" {
    gin.SetMode(gin.ReleaseMode)
  }

  Blacklist, err = blacklist.New(config.BlacklistPath)
  if err != nil {
    return err
  }

  // routes
  if config.LoaderIoToken != "" {
    // loader.io verification
    router.GET("loaderio-" + config.LoaderIoToken + "/", func(c *gin.Context) {
      c.String(http.StatusOK, "loaderio-" + config.LoaderIoToken)
    })
  }

  router.GET("/", func(c *gin.Context) {
    hasOrigin := false
    for k,v := range c.Request.Header {
      if k == "Origin" && len(v) > 0 {
        hasOrigin = true
      }
    }

    if hasOrigin {
      isAllowed, err := isOriginAllowed(c.Request.Header.Get("Origin"))
      if err != nil {
        fmt.Println(err)
        return
      }

      if isAllowed {
        if config.UpstreamUrl != "" {
          // TODO: verify status code (307) is correct and merge/append `c.Request.URL.Query()` with/to `config.UpstreamUrl`
          c.Redirect(http.StatusTemporaryRedirect, config.UpstreamUrl)
        }
      } else {
        c.AbortWithStatus(http.StatusBadRequest)
      }
    }
  })

  router.Run(":" + strconv.Itoa(config.Port))
  return nil
}

func isOriginAllowed(origin string) (bool, error) {
  originUrl, err := url.Parse(origin)
  if err != nil || originUrl.Host == "" {
    fmt.Println("Malformed Origin: " + origin)
    return false, err
  }

  i := sort.Search(Blacklist.Len(), func(i int) bool {
    return Blacklist[i].Name >= originUrl.Host
  })

  if i < Blacklist.Len() && Blacklist[i].Name == originUrl.Host {
    fmt.Println("Origin: " + originUrl.Host + " FOUND in Blacklist at line " + strconv.Itoa(i) + ", aborting ...")
    return false, nil
  } else {
    fmt.Println("Origin: " + originUrl.Host + " NOT FOUND in Blacklist, redirecting ...")
    return true, nil
  }
}
