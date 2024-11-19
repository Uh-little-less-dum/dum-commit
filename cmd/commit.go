/*
Copyright © 2024 Andrew Mueller aiglinski414@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Uh-little-less-dum/dum-commit/internal"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var CommitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Create commit command",
	Long:  `Creates a commit command instead of working inside of a git hook.`,
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data := internal.GetGitCommit()
		gitCmd := exec.Command("git", "commit", "-m", fmt.Sprintf("%s(%s): %s", data[internal.CommitTypeKey], data[internal.ScopeKey], data[internal.MsgKey]), "-m", data[internal.DetailsKey])
		gitCmd.Stdout = os.Stdout
		err := gitCmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		gitCmd.Stdout = os.Stdout
		log.Info("Your commit was applied successfully!")
	},
}

func init() {
	RootCmd.AddCommand(CommitCmd)
}
