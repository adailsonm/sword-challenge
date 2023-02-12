package bootstrap

import (
	"github.com/adailsonm/sword-challenge/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "challenge-sword",
	Short:            "challenge sword api constructor backend",
	TraverseChildren: true,
}

type App struct {
	*cobra.Command
}

func NewApp() App {
	cmd := App{
		Command: rootCmd,
	}
	cmd.AddCommand(commands.GetSubCommands(CommonModules)...)

	return cmd
}

var RootApp = NewApp()
