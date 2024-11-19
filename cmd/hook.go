/*
Copyright © 2024 Andrew Mueller aiglinski414@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Uh-little-less-dum/dum-commit/internal"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var HookCmd = &cobra.Command{
	Use:   "hook",
	Short: "A conventional commits utility",
	Long:  `A conventional commits utility that gathers various options from the root package.json file or a file passed in as an option with --config=/my/config.json`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data := internal.GetGitCommit()
		commitMsg := fmt.Sprintf("%s(%s): %s\n\n%s", data[internal.CommitTypeKey], data[internal.ScopeKey], data[internal.MsgKey], data[internal.DetailsKey])
		err := os.WriteFile(args[0], []byte(commitMsg), 0777)
		if err != nil {
			log.Fatal(err)
		}
		log.Info("Your commit was applied successfully!")
	},
}

func init() {
	RootCmd.AddCommand(HookCmd)
}
