package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)


var rootCmd = &cobra.Command{
	Use:   "name-hub",
	Short: "A list of available usernames in GitHub.",
	Long:  `A list of available usernames in GitHub.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}