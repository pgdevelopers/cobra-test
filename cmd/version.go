package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd("1.0"))
}

func versionCmd(message string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Hugo",
		Long:  `All software has versions. This is Hugo's`,
		Run: func(cmd *cobra.Command, args []string) {
			// fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
			fmt.Fprintf(cmd.OutOrStdout(), message)
		},
	}
}
