package cmd

import (
	"github.com/3ks/NameHub/gen"
	"github.com/spf13/cobra"
)

var (
	initCmd *cobra.Command
)

func init() {
	// init
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Generate all username.",
		Long:  `Generate all username.`,
		Run: func(cmd *cobra.Command, args []string) {
			gen.Username()
		},
	}

	rootCmd.AddCommand(initCmd)
}
