package main

import (
  "os"
  "log"
  "github.com/lunchtime-labs/redchannels"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

  env := os.Getenv("ENV")
	if env == "" {
		log.Fatal("ENV must be set")
	}

  redchannels.Start(port, env)
}
