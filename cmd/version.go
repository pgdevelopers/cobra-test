package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(abcCmd)
}

var abcCmd = &cobra.Command{
	Use:   "abc",
	Short: "abc123",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
		fmt.Fprintf(cmd.OutOrStdout(), "abc test")
	},
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
