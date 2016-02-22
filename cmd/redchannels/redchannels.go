package main

import (
  "os"
  "log"
  "github.com/lunchtime-labs/redchannels/Godeps/_workspace/src/github.com/codegangsta/cli"
  "github.com/lunchtime-labs/redchannels/service"
)

func main() {
  app := cli.NewApp()
	app.Name = "redchannels"
	app.Usage = "CLI interface to the `redchannels` microservice"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
    cli.IntFlag{
      Name: "port, p",
      Value: 5000,
      Usage: "Port to bind RedChannels HTTP server to",
      EnvVar: "PORT",
    },
    cli.StringFlag{
      Name: "env",
      Value: "development",
      Usage: "RedChannels application environment",
      EnvVar: "ENV",
    },
    cli.StringFlag{
      Name: "blacklist, b",
      Value: "blacklist.txt",
      Usage: "Path to RedChannels blacklist file",
      EnvVar: "BLACKLIST_PATH",
    },
    cli.StringFlag{
      Name: "loaderio-verify-token",
      Value: "",
      Usage: "Loader.io verification token",
      EnvVar: "LOADERIO_VERIFY_TOKEN",
    },
	}

	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "Run the RedChannels http server",
			Action: func(c *cli.Context) {
        config := service.Config{
          Port: c.GlobalInt("port"),
          Env: c.GlobalString("env"),
          Blacklist: c.GlobalString("blacklist"),
          LoaderIoToken: c.GlobalString("loaderio-verify-token"),
        }

				service := service.BlacklistService{}
				if err := service.Run(config); err != nil {
					log.Fatal(err)
				}
			},
		},
	}

  app.Run(os.Args)
}
