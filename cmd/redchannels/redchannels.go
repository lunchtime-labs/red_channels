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

	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "Run the RedChannels http server",
      Flags: []cli.Flag{
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
        cli.StringFlag{
          Name: "upstream",
          Value: "",
          Usage: "Url to which non-blacklisted requests are redirected",
          EnvVar: "UPSTREAM_URL",
        },
      },
			Action: func(c *cli.Context) {
        config := service.Config{
          Port: c.Int("port"),
          Env: c.String("env"),
          BlacklistPath: c.String("blacklist"),
          LoaderIoToken: c.String("loaderio-verify-token"),
          UpstreamUrl: c.String("upstream"),
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
