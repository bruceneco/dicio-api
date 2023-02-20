package bootstrap

import (
	"github.com/bruceneco/dicio-api/commands"
	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use:              "dicio-api",
	Short:            "Scraper for Dicio website",
	TraverseChildren: true,
}

type App struct {
	*cobra.Command
}

func NewApp() App {
	cmd := App{
		Command: rootCMD,
	}
	cmd.AddCommand(commands.GetSubCommands(CommonModules)...)
	return cmd
}

var RootApp = NewApp()
