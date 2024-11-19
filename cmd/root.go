/*
Copyright © 2024 Andrew Mueller aiglinski414@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Uh-little-less-dum/dum-commit/internal"
	init_utils "github.com/Uh-little-less-dum/dum-commit/internal/init"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dum-commit",
	Short: "A conventional commits utility",
	Long:  `A conventional commits utility that gathers various options from the root package.json file or a file passed in as an option with --config=/my/config.json`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		form := internal.NewCommitForm()
		err := form.Run()
		if err != nil {
			log.Fatal(err)
		}
		commitMsg := fmt.Sprintf("%s(%s): %s\n\n%s", form.GetString(internal.CommitTypeKey), form.GetString(internal.ScopeKey), form.GetString(internal.MsgKey), form.GetString(internal.DetailsKey))
		err = os.WriteFile(args[0], []byte(commitMsg), 0777)
		if err != nil {
			log.Fatal(err)
		}
		log.Info("Your commit was applied successfully!")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// TODO: Add flags for various inputs.
func init() {
	cobra.OnInitialize(init_utils.InitCmd(rootCmd))
}
