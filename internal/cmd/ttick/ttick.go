package ttick

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

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

	defaultShortDateFormat     = "06/01/02(Mon)"
	defaultShortDatetimeFormat = "06/01/02(Mon) 15:04"
)

func CheckConfigPermission(configFile string) error {
	if exists, _ := util.Exists(configFile); exists {
		// Ensure that the config file has permission 0600, because it contains
		// the API token and should only be read by the user.
		// This is only necessary iff the config file exists, which may not be the case
		// when config is loaded from environment variables.
		fi, err := os.Lstat(configFile)
		if err != nil {
			return fmt.Errorf("fatal error config file: %s", err)
		}
		if runtime.GOOS != "windows" && fi.Mode().Perm() != 0600 {
			return fmt.Errorf("config file has wrong permissions(%o). \nMake sure to give permissions 600 to file %s",
				fi.Mode().Perm(), configFile)
		}
	}

	return nil
}

func loadConfig(c *cli.Context) error {
	viper.SetDefault("ShortDateFormat", defaultShortDateFormat)
	viper.SetDefault("ShortDatetimeFormat", defaultShortDatetimeFormat)
	viper.SetDefault("color", true)

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

	if err := CheckConfigPermission(configFile); err != nil {
		return err
	}

	config := &Config{
		AccessToken:    viper.GetString("token"),
		DebugMode:      c.Bool("debug"),
		Color:          viper.GetBool("color"),
		DateFormat:     viper.GetString("ShortDateFormat"),
		DateTimeFormat: viper.GetString("ShortDatetimeFormat")}

	return nil
}

func AppBeforeHook(c *cli.Context) error {
	if err := loadConfig(c); err != nil {
		return err
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
