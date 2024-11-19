package init_utils

import (
	"github.com/Uh-little-less-dum/dum-commit/internal/findup"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InitCmd(cmd *cobra.Command) func() {
	return func() {
		initLogger()
		v := viper.GetViper()
		// configFile flag
		cmd.Flags().StringP("config", "c", "", "dum-commit --config=/path/to/my/config.json")
		err := v.BindPFlag("config", cmd.Flags().Lookup("config"))
		handleError(err)
		// scope flag
		cmd.Flags().StringP("scope", "s", "", "dum-commit --scope=myScope")
		err = v.BindPFlag("scope", cmd.Flags().Lookup("scope"))
		handleError(err)
		// type flag
		cmd.Flags().StringP("type", "t", "", "dum-commit --type=myCommitType")
		err = v.BindPFlag("type", cmd.Flags().Lookup("type"))
		handleError(err)
		// config file must be nested in dum-commit, even in the mornal config file.

		pkgJson, err := findup.Find("package.json")
		handleError(err)
		configPath := v.GetString("config")
		log.Info(pkgJson)
		if configPath != "" {
			v.SetConfigFile(configPath)
		} else if pkgJson != "" {
			v.SetConfigFile(pkgJson)
		}
		v.SetConfigType("json")
		if err := v.ReadInConfig(); err != nil {
			log.Debug("Can't read config:", err)
		}
		generalHeight := v.GetInt("dum-commit.height")
		if v.GetInt("dum-commit.type.height") == 0 {
			if generalHeight != 0 {
				v.Set("dum-commit.type.height", generalHeight)
			} else {
				v.Set("dum-commit.type.height", 8)
			}
		}
		if v.GetInt("dum-commit.scope.height") == 0 {
			if generalHeight != 0 {
				v.Set("dum-commit.scope.height", generalHeight)
			} else {
				v.Set("dum-commit.scope.height", 8)
			}
		}
	}
}
