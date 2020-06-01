package cmd

import (
	"github.com/3ks/NameHub/gen"
	"github.com/spf13/cobra"
)

var (
	clean *cobra.Command
)

func init() {
	// init
	clean = &cobra.Command{
		Use:   "clean",
		Short: "Generate all username.",
		Long:  `Generate all username.`,
		Run: func(cmd *cobra.Command, args []string) {
			gen.Clean()
		},
	}

	rootCmd.AddCommand(clean)
}
