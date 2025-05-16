package ttick

import (
	"github.com/urfave/cli/v2"
)

var (
    version = "0.0.1-dev"
)

func NewApp(name string) *cli.App {
    app := cli.NewApp()
    app.Version = version

    return app
}
