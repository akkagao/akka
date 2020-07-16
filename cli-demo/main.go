package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	env := ""
	app := cli.NewApp()
	app.Name = "cli-demo"
	app.Usage = "cli test demo"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "env,e",
			Usage:       "指定环境",
			Value:       "dev",
			Destination: &env,
		},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Println("env", env)
		return nil
	}
	app.Run(os.Args)
}
