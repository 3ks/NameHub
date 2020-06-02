package cmd

import (
	"github.com/3ks/NameHub/gen"
	"github.com/spf13/cobra"
)

var (
	startCmd *cobra.Command
)

func init() {
	// start
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start detect username by GitHub API.",
		Long:  `Start detect username by GitHub API.`,
		Run: func(cmd *cobra.Command, args []string) {
			gen.Username()
		},
	}

	rootCmd.AddCommand(startCmd)
}
