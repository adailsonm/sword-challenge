package bootstrap

import (
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
	return cmd
}

var RootApp = NewApp()
