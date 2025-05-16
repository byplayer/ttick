package ttick

import (
	"github.com/urfave/cli/v2"
)

func NewApp(name string) *cli.App {
    app := cli.NewApp()
    app.Version = "0.0.1"

    return app
}
