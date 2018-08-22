package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	Version    string
	BuildStamp string
	SHA        string
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ipcheck",
	Long:  `All software has versions. ipcheck has a build timestamp and a git SHA too!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%#v %#v %#v\n", Version, BuildStamp, SHA)
	},
}
