package ttick

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/byplayer/ttick/internal/util"
	"github.com/rkoesters/xdg/basedir"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

var (
	version    = "0.0.1-dev"
	configPath = filepath.Join(basedir.ConfigHome, appName)
)

const (
	appName    = "ttick"
	appUsage   = "ticktick CLI Client"
	configName = "config"
	configType = "json"
)

func AppBeforeHook(c *cli.Context) error {
	viper.SetConfigType(configType)
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("ttick") // uppercased automatically by viper
	viper.AutomaticEnv()

	configFile := filepath.Join(configPath, configName+"."+configType)
	if err := util.AssureExists(configFile); err != nil {
		return err
	}

	var token string
	if err := viper.ReadInConfig(); err != nil {
		if _, isConfigNotFoundError := err.(viper.ConfigFileNotFoundError); !isConfigNotFoundError {
			// config file was found but could not be read => not recoverable
			return err
		} else if !viper.IsSet("token") {
			// config file not found and token missing (not provided via another source,
			// such as environment variables) => ask interactively for token and store it in config file.
			fmt.Printf("Input API Token: ")
			fmt.Scan(&token)
			viper.Set("token", token)
			buf, err := json.MarshalIndent(viper.AllSettings(), "", "  ")
			if err != nil {
				panic(fmt.Errorf("fatal error config file: %s", err))
			}
			err = os.WriteFile(configFile, buf, 0600)
			if err != nil {
				panic(fmt.Errorf("fatal error config file: %s", err))
			}
		}
	}

	return nil
}

func NewApp(name string) *cli.App {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appUsage
	app.Version = version
	app.EnableBashCompletion = true

	app.Before = AppBeforeHook

	return app
}
