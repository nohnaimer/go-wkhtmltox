package main

import (
	"log"
	"os"

	"github.com/gogap/config"

	"github.com/urfave/cli"
	"go-wkhtmltox/server"

	_ "go-wkhtmltox/wkhtmltox/fetcher/data"
	_ "go-wkhtmltox/wkhtmltox/fetcher/http"
)

func main() {

	var err error

	defer func() {
		if err != nil {
			log.Printf("[go-wkhtmltox]: %s\n", err.Error())
		}
	}()

	app := cli.NewApp()

	app.Usage = "A server for wkhtmltopdf and wkhtmltoimage"

	app.Commands = cli.Commands{
		cli.Command{
			Name:    "run",
			Usage:   "run go-wkhtmltox service",
			Action:  run,
			Aliases: []string{"c"},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "config",
					Usage: "config filename",
					Value: "app.conf",
				},
				&cli.StringFlag{
					Name:  "cwd",
					Usage: "change work dir",
				},
			},
		},
	}

	err = app.Run(os.Args)
}

func run(ctx *cli.Context) (err error) {

	cwd := ctx.String("cwd")
	if len(cwd) != 0 {
		err = os.Chdir(cwd)
	}

	if err != nil {
		return
	}

	configFile := ctx.String("config")

	conf := config.NewConfig(
		config.ConfigFile(configFile),
	)

	srv, err := server.New(conf)

	if err != nil {
		return
	}

	err = srv.Run()

	return
}
