/*
Copyright © 2024 Andrew Mueller aiglinski414@gmail.com
*/
package cmd

import (
	"os"

	init_utils "github.com/Uh-little-less-dum/dum-commit/internal/init"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "dum-commit",
	Short: "A conventional commits utility",
	Long:  `A conventional commits utility that gathers various options from the root package.json file or a file passed in as an option with --config=/my/config.json`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	init_utils.InitCmd(RootCmd)()
}
