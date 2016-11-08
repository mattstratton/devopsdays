package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version string
	Build   string
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Probably Fine",
	Long:  `All software has versions. This is ours.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version: ", Version)
		fmt.Println("Build Time: ", Build)
	},
}
