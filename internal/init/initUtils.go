package init_utils

import (
	"github.com/Uh-little-less-dum/go-utils/pkg/fs/findup"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InitCmd(MyCmd *cobra.Command) func() {
	return func() {
		log.Info("What the flying fuck...")
		v := viper.GetViper()

		// configFile flag
		MyCmd.PersistentFlags().StringP("config", "c", "", "dum-commit --config=/path/to/my/config.json")
		err := v.BindPFlag("config", MyCmd.PersistentFlags().Lookup("config"))
		handleError(err)

		// defaults flag - Apply sensible defaults to force a list instead of inputs.
		MyCmd.PersistentFlags().BoolP("defaults", "d", false, "Apply sensible defaults to force a list of options instead of a text input.")
		err = v.BindPFlag("defaults", MyCmd.PersistentFlags().Lookup("defaults"))
		handleError(err)
		// scope flag
		MyCmd.PersistentFlags().StringP("scope", "s", "", "dum-commit --scope=myScope")
		err = v.BindPFlag("scope", MyCmd.PersistentFlags().Lookup("scope"))
		handleError(err)
		// type flag
		MyCmd.PersistentFlags().StringP("type", "t", "", "dum-commit --type=myCommitType")
		err = v.BindPFlag("type", MyCmd.PersistentFlags().Lookup("type"))
		handleError(err)

		// config file must be nested in dum-commit, even in the mornal config file.
		pkgJson, err := findup.Find("package.json")
		log.Info("pkgJson", pkgJson)
		configPath := v.GetString("config")
		if (err == nil) && (pkgJson != "") {
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
}
